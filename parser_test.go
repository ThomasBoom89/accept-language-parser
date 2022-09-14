package alp

import (
	"errors"
	"reflect"
	"testing"
)

func Test_Parse(t *testing.T) {
	cases := []struct {
		payload  string
		expected []Language
		err      error
	}{
		{"fr-CH", []Language{{Name: "fr-CH", Quality: 1}}, nil},
		{"fr-CH, fr;q=0.9", []Language{{Name: "fr-CH", Quality: 1}, {Name: "fr", Quality: 0.9}}, nil},
		{"fr-CH, fr;q=0.9, en;q=0.8, de;q=0.7, *;q=0.5", []Language{{Name: "fr-CH", Quality: 1}, {Name: "fr", Quality: 0.9}, {Name: "en", Quality: 0.8}, {Name: "de", Quality: 0.7}, {Name: "*", Quality: 0.5}}, nil},
		{"fr-CH, fr;0.9,", nil, errors.New("quality has wrong format")},
		{"fr-CH, fr;q=a.9,", nil, errors.New("detected quality is not a float")},
	}

	for i, c := range cases {
		languages, err := Parse(c.payload)
		if err != nil && c.err == nil {
			t.Fatal("should not throw err:", err, "| test case:", i, "| payload:", c.payload)
		}
		if c.err != nil && c.err.Error() != err.Error() {
			t.Fatal("should throw error:", c.err.Error(), ",not:", err.Error())
		}
		if !(len(languages) == 0 && len(c.expected) == 0) && !reflect.DeepEqual(languages, c.expected) {
			t.Fatal("language should be equal:", languages, c.expected)
		}
	}
}
