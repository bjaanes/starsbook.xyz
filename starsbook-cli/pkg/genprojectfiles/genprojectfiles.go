package genprojectfiles

import (
	"encoding/json"
	"fmt"
	"github.com/go-errors/errors"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/attributes"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/conf"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/nftinfo"
	"io/ioutil"
	"math"
	"path"
	"path/filepath"
	"sort"
	"strconv"
)

type MinProjectOutputJson struct {
	Name         string `json:"name"`
	Link         string `json:"link"`
	NumberOfNfts int    `json:"numberOfNfts"`
}

type ProjectOutputJson struct {
	Name         string          `json:"name"`
	Link         string          `json:"link"`
	NumberOfNfts int             `json:"numberOfNfts"`
	LowestScore  float32         `json:"lowestScore"`
	HighestScore float32         `json:"highestScore"`
	NFTs         []NFTOutputJson `json:"nfts"`
}

type NFTOutputJson struct {
	Name        string                `json:"name"`
	ID          int                   `json:"id"`
	Img         string                `json:"img"`
	RarityScore float32               `json:"rarityScore"`
	RarityRank  int                   `json:"rarityRank"`
	Attributes  []AttributeOutputJson `json:"attributes"`
	Prices      PriceOutputJson       `json:"prices"`
}

type PriceOutputJson struct {
	KRV         float32 `json:"KRV"`
	KARV_NOW    float32 `json:"KARV_NOW"`
	KARV_FUTURE float32 `json:"KARV_FUTURE"`
	KHRV        float32 `json:"KHRV"`
	KHRV_NOW    float32 `json:"KHRV_NOW"`
	KHRV_FUTURE float32 `json:"KHRV_FUTURE"`
}

type AttributeOutputJson struct {
	Type        string  `json:"type"`
	Value       string  `json:"value"`
	Rarity      float32 `json:"rarity"`
	RarityScore float32 `json:"rarityScore"`
}

func ProjectFiles(c conf.Conf) error {
	for _, p := range c.Projects {
		if err := generateProjectFile(p); err != nil {
			return errors.Wrap(err, 0)
		}
	}

	return nil
}

func generateProjectFile(p conf.Project) error {
	fmt.Printf("Generate project files for %q\n", p.Name)
	nfts, err := nftinfo.GetNFTInfos(p)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	attributeMap, err := attributes.GenerateAttributeMap(p)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	attrFile, err := json.MarshalIndent(attributeMap, "", " ")
	if err != nil {
		return errors.Wrap(err, 0)
	}
	err = ioutil.WriteFile(filepath.Join(conf.RootFolder, p.ShortName, "attribute_rarity.json"), attrFile, 0644)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	projectOutput, err := GenerateProjectOutput(p, nfts, attributeMap)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	for i, _ := range projectOutput.NFTs {
		nftFile, err := json.MarshalIndent(projectOutput.NFTs[i], "", " ")
		if err != nil {
			errors.Wrap(err, 0)
		}
		if err := ioutil.WriteFile(filepath.Join(p.GetNftOutDir(), strconv.Itoa(projectOutput.NFTs[i].ID)+".json"), nftFile, 0644); err != nil {
			return errors.Wrap(err, 0)
		}
	}

	projectFile, err := json.MarshalIndent(projectOutput, "", " ")
	if err != nil {
		return errors.Wrap(err, 0)
	}
	if err := ioutil.WriteFile(filepath.Join(p.GetOutDir(), "project.json"), projectFile, 0644); err != nil {
		return errors.Wrap(err, 0)
	}

	minProjectOutput := MinProjectOutputJson{
		Name:         projectOutput.Name,
		Link:         projectOutput.Link,
		NumberOfNfts: projectOutput.NumberOfNfts,
	}
	minProjectFile, err := json.MarshalIndent(minProjectOutput, "", " ")
	if err != nil {
		return errors.Wrap(err, 0)
	}
	if err := ioutil.WriteFile(filepath.Join(p.GetOutDir(), "min_project.json"), minProjectFile, 0644); err != nil {
		return errors.Wrap(err, 0)
	}

	return nil
}

func GenerateProjectOutput(p conf.Project, nfts []nftinfo.NFTInfo, attributeMap attributes.AttributeMap) (ProjectOutputJson, error) {
	projectOutput := ProjectOutputJson{
		Name:         p.Name,
		NumberOfNfts: len(nfts),
	}

	lowestScore := float32(math.MaxFloat32)
	highestScore := float32(0)
	for _, nft := range nfts {
		if projectOutput.Link == "" {
			projectOutput.Link = nft.ExternalUrl
		}

		nftOutput := NFTOutputJson{
			Name: nft.Name,
		}

		_, imgFn := path.Split(nft.Image)
		nftOutput.Img = imgFn

		id, err := nft.GetID()
		if err != nil {
			return ProjectOutputJson{}, errors.Wrap(err, 0)
		}
		nftOutput.ID = id

		nftRarityScore := float32(0.0)
		for _, attr := range nft.Attributes {
			attrOut := AttributeOutputJson{
				Type:  attr.TraitType,
				Value: attr.GetValue(),
			}
			numberOfNftsWithTrait := attributeMap[attr.TraitType][attr.GetValue()]
			rarity := float32(numberOfNftsWithTrait) / float32(projectOutput.NumberOfNfts)
			attrOut.Rarity = rarity
			if numberOfNftsWithTrait != 0 {
				attrOut.RarityScore = float32(projectOutput.NumberOfNfts) / float32(numberOfNftsWithTrait)
				nftRarityScore += attrOut.RarityScore
			}

			nftOutput.Attributes = append(nftOutput.Attributes, attrOut)
		}
		nftOutput.RarityScore = nftRarityScore

		if nftRarityScore < lowestScore {
			lowestScore = nftRarityScore
		}
		if nftRarityScore > highestScore {
			highestScore = nftRarityScore
		}

		projectOutput.NFTs = append(projectOutput.NFTs, nftOutput)
	}

	projectOutput.LowestScore = lowestScore
	projectOutput.HighestScore = highestScore

	sort.Slice(projectOutput.NFTs, func(i, j int) bool {
		if projectOutput.NFTs[i].RarityScore == projectOutput.NFTs[j].RarityScore {
			return projectOutput.NFTs[i].ID < projectOutput.NFTs[j].ID // Lower ID slightly better?
		}

		return projectOutput.NFTs[i].RarityScore > projectOutput.NFTs[j].RarityScore
	})

	for i, _ := range projectOutput.NFTs {
		projectOutput.NFTs[i].RarityRank = i + 1
		projectOutput.NFTs[i].Prices = genPrices(projectOutput.NFTs[i], projectOutput, projectOutput.LowestScore, p.MintPrice)
	}

	return projectOutput, nil
}
