package folders

import (
	"github.com/go-errors/errors"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/conf"
	"os"
)

func CreateProjectFolderStructure(p conf.Project) error {
	if _, err := os.Stat(p.GetOutDir()); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(p.GetOutDir(), 0755); err != nil {
			return errors.Wrap(err, 0)
		}
	}

	nftsRawDir := p.GetNftsRawDir()
	if _, err := os.Stat(nftsRawDir); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(nftsRawDir, 0755); err != nil {
			return errors.Wrap(err, 0)
		}
	}

	imgRawDir := p.GetImgRawDir()
	if _, err := os.Stat(imgRawDir); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(imgRawDir, 0755); err != nil {
			return errors.Wrap(err, 0)
		}
	}

	imgMinDir := p.GetImgMinDir()
	if _, err := os.Stat(imgMinDir); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(imgMinDir, 0755); err != nil {
			return errors.Wrap(err, 0)
		}
	}

	nftsDir := p.GetNftOutDir()
	if _, err := os.Stat(nftsDir); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(nftsDir, 0755); err != nil {
			return errors.Wrap(err, 0)
		}
	}

	return nil
}
