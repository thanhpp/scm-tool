package dto_test

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thanhpp/scm/internal/scmsrv/infra/port/httpsv/dto"
	"github.com/thanhpp/scm/pkg/logger"
)

func setLocalTime() error {
	l, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return err
	}

	time.Local = l

	return nil
}

func TestMytime(t *testing.T) {
	var (
		strTime = "17:31 18-09-1999"
		m       = new(dto.MyTime)
	)

	err := setLocalTime()
	require.NoError(t, err)

	err = m.UnmarshalJSON([]byte(strTime))
	require.NoError(t, err)

	t.Log(m.Time().String())

	byteMytime, err := m.MarshalJSON()
	require.NoError(t, err)
	val, err := strconv.Unquote(string(byteMytime))
	require.NoError(t, err)
	assert.Equal(t, strTime, val)
}

func TestMarshalMyTime(t *testing.T) {
	var (
		val = dto.CreateImportTicketReq{
			ReceiveTime: "17:31 18-09-1999",
		}
	)
	logger.SetDefaultLog()

	data, err := json.Marshal(val)
	require.NoError(t, err)

	t.Log(string(data))
}
