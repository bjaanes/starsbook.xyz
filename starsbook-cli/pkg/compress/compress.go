package compress

import (
	"fmt"
	"github.com/go-errors/errors"
	"github.com/h2non/bimg"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/conf"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/nftinfo"
	"io"
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

	errchan := make(chan error)
	defer close(errchan)

	sem := make(chan struct{}, 100)
	compress := func(fp string, minFp string) {
		sem <- struct{}{} // When there's room we can proceed

		if filepath.Ext(minFp) == ".gif" {
			fmt.Printf("Copying %s\n", fp)
			f, err := os.Open(fp)
			if err != nil {
				errchan <- errors.Wrap(err, 0)
			}
			defer f.Close()

			dest, err := os.Create(minFp)
			if err != nil {
				errchan <- errors.Wrap(err, 0)
			}
			defer dest.Close()

			_, err = io.Copy(dest, f)
			if err != nil {
				errchan <- errors.Wrap(err, 0)
			}
		} else {
			fmt.Printf("Compressing %s\n", fp)
			buffer, err := bimg.Read(fp)
			if err != nil {
				errchan <- errors.Wrap(err, 0)
				return
			}

			newImage, err := bimg.NewImage(buffer).Resize(250, 250)
			if err != nil {
				fmt.Println(fp)
				errchan <- errors.Wrap(err, 0)
				return
			}

			if err := bimg.Write(minFp, newImage); err != nil {
				errchan <- errors.Wrap(err, 0)
				return
			}
			fmt.Printf("Compressing %s done\n", fp)
		}

		errchan <- nil
		<-sem // Free room in the channel
	}

	numCompressions := 0
	for _, nft := range nftInfos {
		_, fn := path.Split(nft.Image)
		fp := filepath.Join(p.GetImgRawDir(), fn)
		minFp := filepath.Join(p.GetImgMinDir(), fn)

		if _, err := os.Stat(minFp); errors.Is(err, os.ErrNotExist) {
			numCompressions++
			go compress(fp, minFp)
		}
	}

	for i := 0; i < numCompressions; i = i + 1 {
		if err := <-errchan; err != nil {
			return errors.Wrap(err, 0)
		}
		fmt.Printf("Compress %d/%d done\n", i, numCompressions)
	}

	return nil
}
