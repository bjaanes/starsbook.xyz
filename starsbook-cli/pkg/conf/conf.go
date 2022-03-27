package conf

import (
	"encoding/json"
	"github.com/go-errors/errors"
	"os"
	"path/filepath"
)

const RootFolder = "public"
const ImgFolder = "imgs"
const NFTRawFolder = "nfts_raw"
const NFTOutFolder = "nfts"

type Project struct {
	Name                         string   `json:"name"`
	IPFSBase                     string   `json:"ipfsBase"`
	MintPrice                    int      `json:"mintPrice"`
	AttributesToIgnoreForRarity  []string `json:"attributesToIgnoreForRarity"`
	AttributesToIgnoreForDisplay []string `json:"attributesToIgnoreForDisplay"`
	ShortName                    string   `json:"shortName"`
	LinkToProjectImage           string   `json:"linkToProjectImage"`
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

func (p Project) GetImgDir() string {
	return filepath.Join(p.GetOutDir(), ImgFolder)
}
