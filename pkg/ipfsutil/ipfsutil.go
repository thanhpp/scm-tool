package ipfsutil

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	ipfsFiles "github.com/ipfs/go-ipfs-files"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	caopts "github.com/ipfs/interface-go-ipfs-core/options"
	path "github.com/ipfs/interface-go-ipfs-core/path"
)

const (
	infuraAPI = "https://ipfs.infura.io:5001"
)

type IPFSInfuraClient struct {
	projectID     string
	projectSecret string
	client        *httpapi.HttpApi
}

func NewIPFSInfuraClient(projectID, projectSecret string) (*IPFSInfuraClient, error) {
	c := &IPFSInfuraClient{
		projectID:     projectID,
		projectSecret: projectSecret,
	}
	httpClient := http.DefaultClient

	client, err := httpapi.NewURLApiWithClient(infuraAPI, httpClient)
	if err != nil {
		return nil, err
	}
	client.Headers.Add("Authorization", "Basic "+basicAuth(projectID, projectSecret))
	c.client = client

	return c, nil
}

func (c IPFSInfuraClient) UploadFile(ctx context.Context, path string) (path.Resolved, error) {
	stat, err := os.Lstat(path)
	if err != nil {
		return nil, err
	}

	file, err := ipfsFiles.NewSerialFile(path, false, stat)
	if err != nil {
		return nil, err
	}

	res, err := c.client.Unixfs().Add(ctx, file, caopts.Unixfs.Pin(true))
	if err != nil {
		return nil, err
	}
	fmt.Printf("[DEBUG] add res %+v\n", res)

	return res, nil
}

func basicAuth(projectId, projectSecret string) string {
	auth := projectId + ":" + projectSecret
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
