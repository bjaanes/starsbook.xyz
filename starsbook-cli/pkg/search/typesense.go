package search

import (
	"github.com/go-errors/errors"
	"github.com/typesense/typesense-go/typesense"
	"os"
)

func getTypesenseClient() (*typesense.Client, error) {
	apiKey := os.Getenv("TYPESENSE_API_KEY")

	if apiKey == "" {
		return nil, errors.New("TYPESENSE_API_KEY env variable missing")
	}

	return typesense.NewClient(
		typesense.WithServer("https://starsbook-typesense.gjermund.tech:30220"),
		typesense.WithAPIKey(apiKey)), nil
}
