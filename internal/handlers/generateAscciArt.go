package handlers

import (
	"fmt"

	"github.com/ediallocyf/asciiartweb/util"
)

// GenerateAsciiArt generates ASCII art based on the provided text and banner type.
func GenerateAsciiArt(text, bannerType string) (string, error) {
	var patternFileName string

	switch bannerType {
	case shadowBanner:
		patternFileName = shadowPatterFileAdrr
	case standardBanner:
		patternFileName = standardPatterFileAdrr
	case thinkerBanner:
		patternFileName = thinkertoyPatterFileAdrr
	default:
		return "", fmt.Errorf("unknown banner type: %s", bannerType)
	}

	chLength := 8
	patternContent, errPattern := util.ReadFileToStr(patternFileName)
	if errPattern != nil {
		return "", fmt.Errorf("error loading pattern file: %v", errPattern)
	}
	patternMap, errPatternMap := util.ContentToMap(patternContent, chLength)
	if errPatternMap != nil {
		return "", fmt.Errorf("error creating pattern map: %v", errPatternMap)
	}
	return util.StrMaker(text, patternMap, chLength), nil
}
