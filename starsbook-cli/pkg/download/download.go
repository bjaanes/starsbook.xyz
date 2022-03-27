package download

import (
	"fmt"
	"github.com/go-errors/errors"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/conf"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/nftinfo"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func NFTs(conf conf.Conf) error {
	sh, err := getIPFSSHell(conf)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	for _, project := range conf.Projects {
		fmt.Printf("Downloading nfts files for %q\n", project.Name)

		if err := setUpOutDirs(project); err != nil {
			return errors.Wrap(err, 0)
		}

		if err := downloadNFTs(sh, project); err != nil {
			return errors.Wrap(err, 0)
		}
	}

	return nil
}

func ProjectImgs(conf conf.Conf) error {
	for _, project := range conf.Projects {
		fmt.Printf("Downloading project image for %q\n", project.Name)

		if err := setUpOutDirs(project); err != nil {
			return errors.Wrap(err, 0)
		}

		if err := downloadProjectImage(project); err != nil {
			return errors.Wrap(err, 0)
		}
	}

	return nil
}

func Imgs(conf conf.Conf) error {
	sh, err := getIPFSSHell(conf)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	for _, project := range conf.Projects {
		fmt.Printf("Downloading imgs for %q\n", project.Name)

		if err := setUpOutDirs(project); err != nil {
			return errors.Wrap(err, 0)
		}

		if err := downloadImages(sh, project); err != nil {
			return errors.Wrap(err, 0)
		}
	}

	return nil
}

func downloadNFTs(sh *shell.Shell, project conf.Project) error {
	ls, err := sh.List(project.IPFSBase)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	// Expects it to only contain files named by the id of the nfts 1, 2, 3, etc
	nftOutDir := project.GetNftsRawDir()
	for _, f := range ls {
		if _, err := os.Stat(filepath.Join(nftOutDir, f.Hash)); errors.Is(err, os.ErrNotExist) {
			fmt.Printf("Downloading %s (%s)\n", f.Name, f.Hash)
			if err := sh.Get(f.Hash, nftOutDir); err != nil {
				return errors.Wrap(err, 0)
			}
		}
	}

	return nil
}

func downloadImages(sh *shell.Shell, project conf.Project) error {
	nftInfos, err := nftinfo.GetNFTInfos(project)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	for _, nft := range nftInfos {
		_, fn := path.Split(nft.Image)
		fp := filepath.Join(project.GetImgDir(), fn)
		if _, err := os.Stat(fp); errors.Is(err, os.ErrNotExist) {
			fmt.Println("Downloading " + fn)
			block, err := sh.BlockGet(strings.TrimPrefix(nft.Image, "ipfs://"))
			if err != nil {
				return errors.Wrap(err, 0)
			}

			err = os.WriteFile(fp, block, 0644)
			if err != nil {
				return errors.Wrap(err, 0)
			}
		}
	}

	return nil
}

func downloadProjectImage(project conf.Project) error {
	url := project.LinkToProjectImage
	fn := filepath.Join(project.GetOutDir(), "projectImage.")

	if _, err := os.Stat(fn); errors.Is(err, os.ErrNotExist) {
		resp, err := http.Get(url)
		if err != nil {
			return errors.Wrap(err, 0)
		}
		defer resp.Body.Close()
		out, err := os.Create(fn)
		if err != nil {
			return errors.Wrap(err, 0)
		}
		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return errors.Wrap(err, 0)
		}
	}

	return nil
}

func getIPFSSHell(c conf.Conf) (*shell.Shell, error) {
	url := "localhost:5001"
	if c.IPFSUrl != "" {
		url = c.IPFSUrl
	}

	sh := shell.NewShell(url)
	if !sh.IsUp() {
		return nil, errors.Errorf("IPFS not up at %q", url)
	}

	return sh, nil
}

func setUpOutDirs(project conf.Project) error {
	if _, err := os.Stat(project.GetOutDir()); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(project.GetOutDir(), 0755); err != nil {
			return errors.Wrap(err, 0)
		}
	}

	nftsRawDir := project.GetNftsRawDir()
	if _, err := os.Stat(nftsRawDir); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(nftsRawDir, 0755); err != nil {
			return errors.Wrap(err, 0)
		}
	}

	imgDir := project.GetImgDir()
	if _, err := os.Stat(imgDir); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(imgDir, 0755); err != nil {
			return errors.Wrap(err, 0)
		}
	}

	nftsDir := project.GetNftOutDir()
	if _, err := os.Stat(nftsDir); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(nftsDir, 0755); err != nil {
			return errors.Wrap(err, 0)
		}
	}

	return nil
}
