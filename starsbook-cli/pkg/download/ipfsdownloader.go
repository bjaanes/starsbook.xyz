package download

import (
	"github.com/go-errors/errors"
	shell "github.com/ipfs/go-ipfs-api"
	"strings"
)

type protocol int

const (
	ipfs = iota
	https
)

type IPFSDownloader struct {
	proto   protocol
	gateway string
	sh      *shell.Shell
}

func CreateIPFSDownloader(url string) (IPFSDownloader, error) {
	if url == "" {
		return IPFSDownloader{}, errors.New("IPFS url cannot be empty")
	}

	var proto protocol
	if strings.HasPrefix(url, "https") {
		proto = https
	} else {
		proto = ipfs
	}

	var sh *shell.Shell
	if proto == ipfs {
		sh := shell.NewShell(url)
		if !sh.IsUp() {
			return IPFSDownloader{}, errors.Errorf("IPFS not up at %q", url)
		}
	}

	return IPFSDownloader{
		proto:   proto,
		gateway: url,
		sh:      sh,
	}, nil
}
