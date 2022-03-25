package search

import (
	"fmt"
	"github.com/go-errors/errors"
	"github.com/typesense/typesense-go/typesense/api"
)

func Perform(term string) error {
	client, err := getTypesenseClient()
	if err != nil {
		return errors.Wrap(err, 0)
	}

	res, err := client.Collection(nftCollectionName).Documents().Search(&api.SearchCollectionParams{
		Q:       term,
		QueryBy: "nftId",
	})

	if err != nil {
		return errors.Wrap(err, 0)
	}

	fmt.Printf("Search done in %dms\n", *res.SearchTimeMs)
	fmt.Printf("Number of hits: %d\n", *res.Found)

	for _, hit := range *res.Hits {
		for f, v := range *hit.Document {
			fmt.Printf("Field: %q, value: %q\n", f, v)
		}
		fmt.Println("")
	}

	return nil
}
