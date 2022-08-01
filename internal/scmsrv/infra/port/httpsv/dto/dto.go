package dto

import "time"

const MyTimeLayout = "15:04 02-01-2006"

type MyTime time.Time

func (m *MyTime) UnmarshalJSON(data []byte) error {
	t, err := time.Parse(MyTimeLayout, string(data))
	if err != nil {
		return err
	}

	*m = MyTime(t)

	return nil
}

func (m MyTime) MarshalJSON() ([]byte, error) {
	str := m.Time().Format(MyTimeLayout)

	return []byte(str), nil
}

func (m MyTime) Time() time.Time {
	return time.Time(m)
}
