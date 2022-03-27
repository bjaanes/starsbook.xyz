package compress

import (
	"fmt"
	"github.com/go-errors/errors"
	"github.com/h2non/bimg"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/conf"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/nftinfo"
	"os"
	"path"
	"path/filepath"
)

func Images(conf conf.Conf) error {
	for _, project := range conf.Projects {
		fmt.Printf("Compressing imgs for %q\n", project.Name)

		if err := compressImages(project); err != nil {
			return errors.Wrap(err, 0)
		}
	}

	return nil
}

func compressImages(p conf.Project) error {
	nftInfos, err := nftinfo.GetNFTInfos(p)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	for _, nft := range nftInfos {
		_, fn := path.Split(nft.Image)
		fp := filepath.Join(p.GetImgDir(), fn)
		minFp := filepath.Join(p.GetImgDir(), "min_"+fn)

		if _, err := os.Stat(minFp); errors.Is(err, os.ErrNotExist) {
			fmt.Println("Compressing " + fn)

			buffer, err := bimg.Read(fp)
			if err != nil {
				return errors.Wrap(err, 0)
			}

			newImage, err := bimg.NewImage(buffer).Resize(250, 250)
			if err != nil {
				return errors.Wrap(err, 0)
			}

			if err := bimg.Write(minFp, newImage); err != nil {
				return errors.Wrap(err, 0)
			}
		}
	}

	return nil
}
