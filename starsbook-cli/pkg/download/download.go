package download

import (
	"fmt"
	"github.com/go-errors/errors"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/conf"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/nftinfo"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func NFTs(conf conf.Conf) error {
	for _, project := range conf.Projects {
		fmt.Printf("Downloading nfts files for %q\n", project.Name)

		if err := downloadNFTs(project); err != nil {
			return errors.Wrap(err, 0)
		}
	}

	return nil
}

func ProjectImgs(conf conf.Conf) error {
	for _, project := range conf.Projects {
		fmt.Printf("Downloading project image for %q\n", project.Name)

		if err := downloadProjectImage(project); err != nil {
			return errors.Wrap(err, 0)
		}
	}

	return nil
}

func Imgs(conf conf.Conf) error {
	for _, project := range conf.Projects {
		fmt.Printf("Downloading imgs for %q\n", project.Name)

		if err := downloadImages(project); err != nil {
			return errors.Wrap(err, 0)
		}
	}

	return nil
}

func downloadNFTs(project conf.Project) error {

	errchan := make(chan error)
	defer close(errchan)

	sem := make(chan struct{}, 20)
	download := func(url string, fp string) {
		sem <- struct{}{} // When there's room we can proceed

		fmt.Printf("Downloading %s\n", fp)
		if err := downloadFile(url, fp); err != nil {
			errchan <- errors.Wrap(err, 0)
		} else {
			errchan <- nil
		}
		fmt.Printf("Download %s done\n", fp)

		<-sem // Free room in the channel
	}

	// Expects it to only contain files named by the id of the nfts 1, 2, 3, etc
	numDownloads := 0
	nftOutDir := project.GetNftsRawDir()
	for i := 1; i <= project.NumberOfNFTs; i++ {
		fp := filepath.Join(nftOutDir, strconv.Itoa(i)+".json")

		if _, err := os.Stat(fp); errors.Is(err, os.ErrNotExist) {
			url := fmt.Sprintf("https://ipfs.stargaze.zone/ipfs/%s/%d", project.IPFSBase, i)
			numDownloads++
			go download(url, fp)
		}
	}

	for i := 0; i < numDownloads; i = i + 1 {
		if err := <-errchan; err != nil {
			return errors.Wrap(err, 0)
		}
		fmt.Printf("Download %d/%d done\n", i, numDownloads)
	}

	return nil
}

func downloadImages(project conf.Project) error {
	nftInfos, err := nftinfo.GetNFTInfos(project)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	// CHANNELS :D
	errchan := make(chan error)
	defer close(errchan)

	sem := make(chan struct{}, 20)
	download := func(url string, fp string) {
		sem <- struct{}{} // When there's room we can proceed

		fmt.Printf("Downloading %s\n", fp)
		if err := downloadFile(url, fp); err != nil {
			errchan <- errors.Wrap(err, 0)
		} else {
			errchan <- nil
		}
		fmt.Printf("Download %s done\n", fp)

		<-sem // Free room in the channel
	}

	numDownloads := 0
	for _, nft := range nftInfos {
		_, fn := path.Split(nft.Image)
		fp := filepath.Join(project.GetImgRawDir(), fn)
		if _, err := os.Stat(fp); errors.Is(err, os.ErrNotExist) {
			var url string
			if strings.HasPrefix(nft.Image, "ipfs://") {
				ipfsLink := strings.TrimPrefix(nft.Image, "ipfs://")
				url = fmt.Sprintf("https://ipfs.stargaze.zone/ipfs/%s", ipfsLink)
			} else if strings.HasPrefix(nft.Image, "https://") {
				url = nft.Image
			} else {
				return errors.Errorf("Nft image url %q invalid", nft.Image)
			}

			numDownloads++
			go download(url, fp)
		}
	}

	for i := 0; i < numDownloads; i = i + 1 {
		if err := <-errchan; err != nil {
			return errors.Wrap(err, 0)
		}
		fmt.Printf("Download %d/%d done\n", i, numDownloads)
	}

	return nil
}

func downloadProjectImage(project conf.Project) error {
	url := project.LinkToProjectImage
	fn := filepath.Join(project.GetOutDir(), "projectImage.")

	if _, err := os.Stat(fn); errors.Is(err, os.ErrNotExist) {
		return downloadFile(url, fn)
	}

	return nil
}

func downloadFile(url string, fp string) error {
	resp, err := http.Get(url)
	if err != nil {
		return errors.Wrap(err, 0)
	}
	defer resp.Body.Close()
	out, err := os.Create(fp)
	if err != nil {
		return errors.Wrap(err, 0)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	return nil
}
