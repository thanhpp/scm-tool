package nftsvclient

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/pkg/constx"
	"github.com/thanhpp/scm/pkg/logger"
)

type NFTServiceClientError struct {
	fn  string
	op  string
	err error
}

func (e NFTServiceClientError) Error() string {
	return fmt.Sprintf("pkg: nftsvclient, fn: %s, op: %s, err: %s", e.fn, e.op, e.err)
}

func (e NFTServiceClientError) Unwrap() error {
	return e.err
}

func newError(fn, op string, err error) error {
	return NFTServiceClientError{
		fn:  fn,
		op:  op,
		err: err,
	}
}

var (
	ErrNotOKStatus = fmt.Errorf("nftsvclient: not ok status")
)

type NFTServiceClient struct {
	baseURL    string
	httpClient *http.Client
}

func New(baseURL string) *NFTServiceClient {
	return &NFTServiceClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
			Timeout: constx.DefaultHTTPClientTimeout,
		},
	}
}

func (c *NFTServiceClient) MintSeriNFT(ctx context.Context, serial *entity.Serial) error {
	req := new(ReqCreateSeriNFT)
	if err := req.SetData(serial); err != nil {
		return err
	}

	resp := new(RespCreateSeriNFT)
	url := fmt.Sprintf("%s/mint", c.baseURL)

	if err := c.doRequest(ctx, http.MethodPost, url, req, resp); err != nil {
		return fmt.Errorf("nftsvclient: mint seri nft, err: %w, resp %+v", err, resp)
	}
	logger.Debugw("mint seri nft ok", "tx_hash", resp.Data.TxHash)

	return nil
}

func (c *NFTServiceClient) GetNFTInfoBySeri(ctx context.Context, seri string) (*NFTInfo, error) {
	resp := new(RespGetNFTInfo)

	url := fmt.Sprintf("%s/seri/%s", c.baseURL, seri)
	if err := c.doRequest(ctx, http.MethodGet, url, nil, resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}

func (c *NFTServiceClient) doRequest(
	ctx context.Context, method, url string, body interface{}, resp interface{},
) error {
	var (
		req *http.Request
		err error
	)

	if body != nil {
		bodyData, err := json.Marshal(body)
		if err != nil {
			return newError("doRequest", "json marshal", err)
		}
		logger.Debugw("doRequest", "body", string(bodyData))

		req, err = http.NewRequestWithContext(ctx, method, url, bytes.NewReader(bodyData))
		if err != nil {
			return newError("doRequest", "new request with data", err)
		}
	} else {
		req, err = http.NewRequestWithContext(ctx, method, url, nil)
		if err != nil {
			return newError("doRequest", "new request", err)
		}
	}

	httpResp, err := c.httpClient.Do(req)
	if err != nil {
		return newError("doRequest", "http Do", err)
	}
	defer httpResp.Body.Close()

	if resp != nil {
		bodyData, err := ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return newError("doRequest", "read body data", err)
		}
		logger.Debugw("body data", "data", string(bodyData))

		err = json.Unmarshal(bodyData, resp)
		if err != nil {
			return newError("doRequest", "json decode", err)
		}

		if httpResp.StatusCode != http.StatusOK {
			return ErrNotOKStatus
		}
	}

	return nil
}
