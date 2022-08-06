package rbmq_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/thanhpp/scm/pkg/rbmq"
)

const (
	url   = "amqp://guest:guest@localhost:5672/"
	queue = "test-queue"
)

func newClient() (*rbmq.Client, error) {
	return rbmq.NewClient(url)
}

func TestCreateQueue(t *testing.T) {
	c, err := newClient()
	require.NoError(t, err)

	err = c.CreateQueue(queue)
	require.NoError(t, err)

	err = c.Close()
	require.NoError(t, err)
}

type A struct {
	Val string `json:"va;"`
}

func TestPubishMsg(t *testing.T) {
	c, err := newClient()
	require.NoError(t, err)

	err = c.CreateQueue(queue)
	require.NoError(t, err)

	a := &A{
		Val: "test val",
	}
	data, err := json.Marshal(a)
	require.NoError(t, err)

	err = c.PublishJSONMessage(queue, data)
	require.NoError(t, err)

	err = c.Close()
	require.NoError(t, err)
}

func TestConsum(t *testing.T) {
	c, err := newClient()
	require.NoError(t, err)
	defer func() {
		err = c.Close()
		require.NoError(t, err)
	}()

	err = c.CreateQueue(queue)
	require.NoError(t, err)

	a := &A{
		Val: "test val",
	}
	data, err := json.Marshal(a)
	require.NoError(t, err)

	err = c.PublishJSONMessage(queue, data)
	require.NoError(t, err)

	msg, err := c.GetConsumerChannel(queue)
	require.NoError(t, err)

	tick := time.Tick(time.Second * 5)
	for {
		select {
		case <-tick:
			return
		case message := <-msg:
			aResp := new(A)
			err = json.Unmarshal(message.Body, aResp)
			require.NoError(t, err)

			t.Logf("received: %+v\n", aResp)
		}
	}
}
