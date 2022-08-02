package dto

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin/binding"
)

const MyTimeLayout = "15:04 02-01-2006"

type MyTime time.Time

func (m *MyTime) UnmarshalJSON(data []byte) error {
	t, err := time.Parse(MyTimeLayout, string(data))
	if err != nil {
		return fmt.Errorf("unmarshal MyTime %s error: %w", string(data), err)
	}

	*m = MyTime(t)

	return nil
}

func (m MyTime) MarshalJSON() ([]byte, error) {
	str := m.Time().Format(MyTimeLayout)

	return json.Marshal(str)
}

func (m MyTime) Time() time.Time {
	return time.Time(m)
}

func (m MyTime) Bind(req *http.Request, obj interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	if err := binding.MapFormWithTag(obj, req.Form, "receive_time"); err != nil {
		return err
	}

	return nil
}
