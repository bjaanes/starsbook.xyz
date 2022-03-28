package nftinfo

import (
	"encoding/json"
	"fmt"
	"github.com/go-errors/errors"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/conf"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

type NFTInfo struct {
	ID          int         `json:"-"`
	Attributes  []Attribute `json:"attributes"`
	Description string      `json:"description"`
	ExternalUrl string      `json:"external_url"`
	Name        string      `json:"name"`
	Image       string      `json:"image"`
}

type Attribute struct {
	TraitType   string      `json:"trait_type"`
	TraitValue  string      `json:"trait_value,omitempty"`
	DisplayType string      `json:"display_type,omitempty"`
	Value       interface{} `json:"value,omitempty"`
}

func (attr Attribute) GetValue() string {
	if attr.TraitValue != "" {
		return attr.TraitValue
	} else {
		return fmt.Sprintf("%v", attr.Value)
	}
}

func GetNFTInfos(project conf.Project) ([]NFTInfo, error) {
	var nfts []NFTInfo
	err := filepath.WalkDir(project.GetNftsRawDir(), func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return errors.Wrap(err, 0)
		}

		if d.IsDir() {
			return nil
		}

		f, err := os.Open(p)
		defer f.Close()
		if err != nil {
			return errors.Wrap(err, 0)
		}

		nftInfo := NFTInfo{}
		err = json.NewDecoder(f).Decode(&nftInfo)
		if err != nil {
			return errors.Wrap(err, 0)
		}

		_, fn := path.Split(p)

		idStr := strings.TrimSuffix(fn, filepath.Ext(fn))
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return errors.Wrap(err, 0)
		}
		nftInfo.ID = id

		nfts = append(nfts, nftInfo)

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return nfts, nil
}
