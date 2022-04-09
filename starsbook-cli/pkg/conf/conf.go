package conf

import (
	"encoding/json"
	"github.com/go-errors/errors"
	"os"
	"path/filepath"
)

const RootFolder = "bucket-assets"
const ImgRawFolder = "imgs_raw"
const ImgCompressedFolder = "imgs_min"
const NFTRawFolder = "nfts_raw"
const NFTOutFolder = "nfts"

type Project struct {
	Name                         string   `json:"name"`
	Description                  string   `json:"description"`
	ExternalUrl                  string   `json:"externalUrl"`
	IPFSBase                     string   `json:"ipfsBase"`
	MintPrice                    int      `json:"mintPrice"`
	NumberOfNFTs                 int      `json:"numberOfNfts"`
	AttributesToIgnoreForRarity  []string `json:"attributesToIgnoreForRarity"`
	AttributesToIgnoreForDisplay []string `json:"attributesToIgnoreForDisplay"`
	RarityOverrideAttribute      string   `json:"rarityOverrideAttribute"`
	ShortName                    string   `json:"shortName"`
	LinkToProjectImage           string   `json:"linkToProjectImage"`
	TraitIdOverride              string   `json:"traitIdOverride"`
	ComingSoon                   bool     `json:"comingSoon"`
	Hidden                       bool     `json:"hidden"`
}

type Conf struct {
	IPFSUrl  string    `json:"ipfsUrl"`
	Projects []Project `json:"projects"`
}

func GetConfig() (Conf, error) {
	cf, err := os.Open("conf.json")
	if err != nil {
		return Conf{}, errors.Wrap(err, 0)
	}
	defer cf.Close()
	conf := Conf{}
	err = json.NewDecoder(cf).Decode(&conf)
	if err != nil {
		return Conf{}, errors.Wrap(err, 0)
	}

	return conf, nil
}

func (p Project) GetOutDir() string {
	return filepath.Join(RootFolder, p.ShortName)
}

func (p Project) GetNftsRawDir() string {
	return filepath.Join(p.GetOutDir(), NFTRawFolder)
}

func (p Project) GetNftOutDir() string {
	return filepath.Join(p.GetOutDir(), NFTOutFolder)
}

func (p Project) GetImgRawDir() string {
	return filepath.Join(p.GetOutDir(), ImgRawFolder)
}

func (p Project) GetImgMinDir() string {
	return filepath.Join(p.GetOutDir(), ImgCompressedFolder)
}
