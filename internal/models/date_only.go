package models

import "time"

type DateOnly time.Time

func (d *DateOnly) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return err
	}
	*d = DateOnly(t)
	return nil
}

func (d DateOnly) ToTime() time.Time {
	return time.Time(d)
}
