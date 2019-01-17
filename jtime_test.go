package jtime

import (
	"testing"
	"time"
)

func TestDuration_MarshalJSON(t *testing.T) {
	var d Duration
	d.Duration, _ = time.ParseDuration("1s")
	json, _ := d.MarshalJSON()
	if string(json) != `"1s"` {
		t.Fail()
	}
}

func TestDuration_UnmarshalJSON(t *testing.T) {
	var tests = []struct {
		json     string
		duration string
	}{
		{`"1s"`, "1s"}, {`"1h10m"`, "1h10m"}, {`"1s123ns"`, "1s123ns"},
	}

	for _, test := range tests {
		duration, _ := time.ParseDuration(test.duration)
		jdur := &Duration{}
		jdur.UnmarshalJSON([]byte(test.json))

		if jdur.String() != duration.String() {
			t.Errorf("unmarshal failed, got %s, exptected %s", jdur.String(), duration.String())
		}
	}


	jdur := &Duration{}
	err := jdur.UnmarshalJSON([]byte("1second"))
	if err == nil {
		t.Fail()
	}
}
