package ipfsutil_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"github.com/thanhpp/scm/pkg/ipfsutil"
)

const (
	envFile      = ".env"
	prjIDEnv     = "PROJECT_ID"
	prjSecretEnv = "PROJECT_SECRET"
)

func init() {
	if err := godotenv.Load(envFile); err != nil {
		log.Println("load env file", envFile, "err", err)
	}
}

func TestUploadToIPFS(t *testing.T) {
	projectID := os.Getenv(prjIDEnv)
	projectSecret := os.Getenv(prjSecretEnv)

	if len(projectID)*len(projectSecret) == 0 {
		t.Error("empty project data")
		return
	}

	client, err := ipfsutil.NewIPFSInfuraClient(projectID, projectSecret)
	require.NoError(t, err)

	var (
		ctx      = context.Background()
		filePath = "/home/thanhpp/go/src/github.com/thanhpp/scm-tool/pkg/ipfsutil/ipfsutil.go"
	)

	res, err := client.UploadFile(ctx, filePath)
	require.NoError(t, err)
	t.Logf("upload completed %+v", res.String())
}
