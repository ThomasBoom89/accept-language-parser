package alp

import (
	"errors"
	"strconv"
	"strings"
)

type Language struct {
	Name    string
	Quality float32
}

const (
	HeaderSeparator   = ","
	LanguageSeparator = ";"
	QualitySeparator  = "="
)

func Parse(header string) ([]Language, error) {
	var res []Language
	var quality float32
	var err error
	languages := strings.Split(header, HeaderSeparator)
	for _, language := range languages {
		quality = 1.0
		splittedLanguage := strings.Split(language, LanguageSeparator)
		name := strings.TrimSpace(splittedLanguage[0])
		if len(splittedLanguage) != 1 {
			quality, err = parseQuality(splittedLanguage[1])
			if err != nil {
				return nil, err
			}
		}
		res = append(res, Language{Name: name, Quality: quality})
	}

	return res, nil
}

func parseQuality(value string) (float32, error) {
	splittedQuality := strings.Split(value, QualitySeparator)
	if len(splittedQuality) <= 1 {
		return 0, errors.New("quality has wrong format")
	}
	quality, err := strconv.ParseFloat(splittedQuality[1], 32)
	if err != nil {
		return 0, errors.New("detected quality is not a float")
	}

	return float32(quality), nil
}
