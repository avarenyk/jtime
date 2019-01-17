package jtime

import (
	"strings"
	"time"
)

type Duration struct {
	time.Duration
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *Duration) UnmarshalJSON(p []byte) error {
	duration, err := time.ParseDuration(strings.Trim(string(p), `"`))
	if err != nil {
		return err
	}

	d.Duration = duration

	return nil
}
