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
	nftCollectionName        = "nfts"
	collectionCollectionName = "collections"
)

var (
	nftDefaultSortingField = "nftId"
	nftFields              = []api.Field{
		{
			Name: "id",
			Type: "string",
		},
		{
			Name: "name",
			Type: "string",
		},
		{
			Name: "nftId",
			Type: "int32",
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
			Name: "imageFileName",
			Type: "string",
		},
		{
			Name: "collectionName",
			Type: "string",
		},
		{
			Name: "collectionShortName",
			Type: "string",
		},
		{
			Name: "collectionImageUrl",
			Type: "string",
		},
		{
			Name: "comingSoon",
			Type: "bool",
		},
	}
	collectionDefaultSortingField = ""
	collectionFields              = []api.Field{
		{
			Name: "name",
			Type: "string",
		},
		{
			Name: "collectionId",
			Type: "int32",
		},
		{
			Name: "shortName",
			Type: "string",
		},
		{
			Name: "description",
			Type: "string",
		},
		{
			Name: "externalUrl",
			Type: "string",
		},
		{
			Name: "numberOfNfts",
			Type: "int32",
		},
		{
			Name: "originalMintPrice",
			Type: "int32",
		},
		{
			Name: "imageUrl",
			Type: "string",
		},
		{
			Name: "comingSoon",
			Type: "bool",
		},
	}
)

type nftDocument struct {
	Name                string  `json:"name"`
	ID                  string  `json:"id"`
	NFTID               int32   `json:"nftId"`
	RarityScore         float32 `json:"rarityScore"`
	RarityRank          int32   `json:"rarityRank"`
	ImageFileName       string  `json:"imageFileName"`
	CollectionName      string  `json:"collectionName"`
	CollectionShortName string  `json:"collectionShortName"`
	CollectionImageUrl  string  `json:"collectionImageUrl"`
	ComingSoon          bool    `json:"comingSoon"`
}

type nftCollection struct {
	Name              string `json:"name"`
	CollectionID      int32  `json:"collectionId"`
	ShortName         string `json:"shortName"`
	ID                string `json:"ID"`
	Description       string `json:"description"`
	ExternalUrl       string `json:"externalUrl"`
	NumberOfNfts      int32  `json:"numberOfNfts"`
	OriginalMintPrice int32  `json:"originalMintPrice"`
	ImageUrl          string `json:"imageUrl"`
	ComingSoon        bool   `json:"comingSoon"`
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

	if force && hasCollectionCollection(colls) {
		fmt.Printf("Force, deleting collection %q\n", collectionCollectionName)
		_, err := client.Collection(collectionCollectionName).Delete()
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
			DefaultSortingField: &nftDefaultSortingField,
			Fields:              nftFields,
			SymbolsToIndex:      nil,
			TokenSeparators:     nil,
		})
		if err != nil {
			return errors.Wrap(err, 0)
		}
	}

	if !hasCollectionCollection(colls) {
		fmt.Printf("Missing collections collection, creating %q...\n", collectionCollectionName)
		_, err := client.Collections().Create(&api.CollectionSchema{
			Name:                collectionCollectionName,
			DefaultSortingField: &collectionDefaultSortingField,
			Fields:              collectionFields,
			SymbolsToIndex:      nil,
			TokenSeparators:     nil,
		})
		if err != nil {
			return errors.Wrap(err, 0)
		}
	}

	for i, p := range conf.Projects {
		fmt.Printf("Upserting collection to collection collection  for %q\n", p.Name)
		if _, err := client.Collection(collectionCollectionName).Documents().Upsert(&nftCollection{
			Name:              p.Name,
			CollectionID:      int32(i + 1),
			ID:                p.ShortName,
			ShortName:         p.ShortName,
			Description:       p.Description,
			ExternalUrl:       p.ExternalUrl,
			NumberOfNfts:      int32(p.NumberOfNFTs),
			OriginalMintPrice: int32(p.MintPrice),
			ImageUrl:          fmt.Sprintf("https://starsbook.xyz/%s/projectImage", p.ShortName),
			ComingSoon:        p.ComingSoon,
		}); err != nil {
			return errors.Wrap(err, 0)
		}

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
			NFTID, err := strconv.Atoi(strId)
			if err != nil {
				return errors.Wrap(err, 0)
			}
			id := fmt.Sprintf("%s_%s", p.ShortName, strId)

			documents = append(documents, nftDocument{
				Name:                nft.Name,
				ID:                  id,
				NFTID:               int32(NFTID),
				RarityScore:         nft.RarityScore,
				RarityRank:          int32(nft.RarityRank),
				ImageFileName:       nft.Img,
				CollectionName:      p.Name,
				CollectionShortName: p.ShortName,
				CollectionImageUrl:  fmt.Sprintf("https://starsbook.xyz/%s/projectImage", p.ShortName),
				ComingSoon:          p.ComingSoon,
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

func hasCollectionCollection(colls []*api.CollectionResponse) bool {
	for _, coll := range colls {
		if coll.Name == collectionCollectionName {
			return true
		}
	}

	return false
}
