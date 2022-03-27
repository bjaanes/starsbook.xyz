package search

import (
	"fmt"
	"github.com/go-errors/errors"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/attributes"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/conf"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/genprojectfiles"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/nftinfo"
	"github.com/typesense/typesense-go/typesense/api"
	"strconv"
)

const (
	nftCollectionName = "nfts"
)

var (
	defaultSortingField = "rarityRank"
	fields              = []api.Field{
		{
			Name: "name",
			Type: "string",
		},
		{
			Name: "id",
			Type: "string",
		},
		{
			Name: "nftId",
			Type: "string",
		},
		{
			Name: "rarityScore",
			Type: "float",
		},
		{
			Name: "rarityRank",
			Type: "int32",
		},
		{
			Name: "imageUrl",
			Type: "string",
		},
		{
			Name: "collectionName",
			Type: "string",
		},
		{
			Name: "collectionImageUrl",
			Type: "string",
		},
	}
)

type nftDocument struct {
	Name               string  `json:"name"`
	ID                 string  `json:"id"`
	NFTID              string  `json:"nftId"`
	RarityScore        float32 `json:"rarityScore"`
	RarityRank         int32   `json:"rarityRank"`
	ImageUrl           string  `json:"imageUrl"`
	CollectionName     string  `json:"collectionName"`
	CollectionImageUrl string  `json:"collectionImageUrl"`
}

func UpsertIndexes(conf conf.Conf, force bool) error {

	client, err := getTypesenseClient()
	if err != nil {
		return errors.Wrap(err, 0)
	}

	colls, err := client.Collections().Retrieve()
	if err != nil {
		return errors.Wrap(err, 0)
	}

	if force && hasNftCollection(colls) {
		fmt.Printf("Force, deleting collection %q\n", nftCollectionName)
		_, err := client.Collection(nftCollectionName).Delete()
		if err != nil {
			return errors.Wrap(err, 0)
		}

		colls, err = client.Collections().Retrieve()
		if err != nil {
			return errors.Wrap(err, 0)
		}
	}

	if !hasNftCollection(colls) {
		fmt.Printf("Missing nfts collection, creating %q...\n", nftCollectionName)
		_, err := client.Collections().Create(&api.CollectionSchema{
			Name:                nftCollectionName,
			DefaultSortingField: &defaultSortingField,
			Fields:              fields,
			SymbolsToIndex:      nil,
			TokenSeparators:     nil,
		})
		if err != nil {
			return errors.Wrap(err, 0)
		}
	}

	for _, p := range conf.Projects {
		fmt.Printf("Upserting documents for %q\n", p.Name)
		nfts, err := nftinfo.GetNFTInfos(p)
		if err != nil {
			return errors.Wrap(err, 0)
		}

		attributeMap, err := attributes.GenerateRarityAttributeMap(p)
		if err != nil {
			return errors.Wrap(err, 0)
		}

		projectOutput, err := genprojectfiles.GenerateProjectOutput(p, nfts, attributeMap)
		if err != nil {
			return errors.Wrap(err, 0)
		}

		var documents []interface{}
		for _, nft := range projectOutput.NFTs {
			strId := ""
			if p.TraitIdOverride != "" {
				for _, attr := range nft.Attributes {
					if attr.Type == p.TraitIdOverride {
						strId = attr.Value
					}
				}
			}
			if strId == "" {
				strId = strconv.Itoa(nft.ID)
			}
			id := fmt.Sprintf("%s_%s", p.ShortName, strId)

			documents = append(documents, nftDocument{
				Name:               nft.Name,
				ID:                 id,
				NFTID:              strId,
				RarityScore:        nft.RarityScore,
				RarityRank:         int32(nft.RarityRank),
				ImageUrl:           fmt.Sprintf("https://starsbook.xyz/%s/imgs/%s", p.ShortName, nft.Img),
				CollectionName:     p.Name,
				CollectionImageUrl: fmt.Sprintf("https://starsbook.xyz/%s/projectImage", p.ShortName),
			})
		}

		importAction := "upsert"
		fmt.Printf("Upserting %d documents from %s into collection: %q\n", len(documents), p.Name, nftCollectionName)
		tooLarge := len(documents) > 5000

		for tooLarge {
			documentsToUpsert := documents[0:5000]
			if _, err := client.Collection(nftCollectionName).Documents().Import(documentsToUpsert, &api.ImportDocumentsParams{
				Action: &importAction,
			}); err != nil {
				return errors.Wrap(err, 0)
			}

			documents = documents[5000:]
			tooLarge = len(documents) > 5000
		}

		if _, err := client.Collection(nftCollectionName).Documents().Import(documents, &api.ImportDocumentsParams{
			Action: &importAction,
		}); err != nil {
			return errors.Wrap(err, 0)
		}
	}

	return nil
}

func hasNftCollection(colls []*api.CollectionResponse) bool {
	for _, coll := range colls {
		if coll.Name == nftCollectionName {
			return true
		}
	}

	return false
}
