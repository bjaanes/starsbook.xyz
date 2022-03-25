package search

import (
	"fmt"
	"github.com/go-errors/errors"
)

func PrintStatus() error {
	c, err := getTypesenseClient()
	if err != nil {
		return errors.Wrap(err, 0)
	}

	colls, err := c.Collections().Retrieve()
	if err != nil {
		return errors.Wrap(err, 0)
	}

	fmt.Printf("Number of collections: %d", len(colls))
	fmt.Println("")

	for _, coll := range colls {
		fmt.Printf("Collection name: %q\n", coll.Name)
		fmt.Printf("DefaultSortingField: %q\n", *coll.DefaultSortingField)
		fmt.Printf("Number of documents: %d\n", coll.NumDocuments)

		fmt.Println("Fields:")
		for _, f := range coll.Fields {
			fmt.Printf("Field name: %q\n", f.Name)
			fmt.Printf("Field type: %q\n", f.Type)
			fmt.Printf("Field face: %v\n", *f.Facet)
			fmt.Println("---")
		}
		fmt.Println("")
	}

	return nil
}
