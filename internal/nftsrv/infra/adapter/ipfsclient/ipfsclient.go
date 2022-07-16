package ipfsclient

import (
	"context"

	"github.com/thanhpp/scm/pkg/ipfsutil"
)

type IPFSClient struct {
	c *ipfsutil.IPFSInfuraClient
}

func NewIPFSClient(infuraPrjID, infuraPrjSec string) (*IPFSClient, error) {
	c, err := ipfsutil.NewIPFSInfuraClient(
		infuraPrjID, infuraPrjSec,
	)
	if err != nil {
		return nil, err
	}

	return &IPFSClient{c: c}, nil
}

func (c *IPFSClient) UploadFile(ctx context.Context, path string) (string, error) {
	r, err := c.c.UploadFile(ctx, path)
	if err != nil {
		return "", err
	}

	return r.Cid().String(), nil
}
