package search

import (
	"fmt"
	"github.com/go-errors/errors"
	"github.com/typesense/typesense-go/typesense/api"
)

func ListApiKeys() error {
	c, err := getTypesenseClient()
	if err != nil {
		return errors.Wrap(err, 0)
	}

	keys, err := c.Keys().Retrieve()
	if err != nil {
		return errors.Wrap(err, 0)
	}

	fmt.Printf("%d keys found:\n", len(keys))
	for _, key := range keys {
		fmt.Printf("Key id: %d\n", key.Id)
		fmt.Printf("Key description: %s\n", *key.Description)
		fmt.Println("---")
	}

	return nil
}

func CreateApiKey(description string) error {
	c, err := getTypesenseClient()
	if err != nil {
		return errors.Wrap(err, 0)
	}

	apiKey, err := c.Keys().Create(&api.ApiKeySchema{
		Actions:     []string{"documents:search"},
		Collections: []string{"*"},
		Description: &description,
	})
	if err != nil {
		return errors.Wrap(err, 0)
	}

	fmt.Printf("Key created: %s\n", apiKey.Value)

	return nil
}

func DeleteApiKey(id int64) error {
	c, err := getTypesenseClient()
	if err != nil {
		return errors.Wrap(err, 0)
	}

	_, err = c.Key(id).Delete()
	if err != nil {
		return errors.Wrap(err, 0)
	}

	fmt.Printf("Key %d deleted\n", id)

	return nil
}
