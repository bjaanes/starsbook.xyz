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
		typesense.WithServer("http://provider.akt.computer:30220"),
		typesense.WithAPIKey(apiKey)), nil
}
