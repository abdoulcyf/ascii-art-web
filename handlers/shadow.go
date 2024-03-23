package handlers

import (
	"fmt"
	cli "main/pkg/cli"
	arr "main/pkg/plot/arr"
	draw "main/pkg/plot/draw"
)

// Function variable to hold ASCII art generator
var GenerateShadowAsciiArt func(text string) string

func ShadowHandler() {
	directory := "."
	//patternFileName := "standard.txt"
	patternFileName := "shadow.txt"
	chLength := 8
	var firstCh byte = ' '
	var lastCh byte = '~'
	//----------------------------------------
	patternContent, errPattern := cli.ReadFileToStr(directory, patternFileName)

	if errPattern != nil {
		fmt.Println("Error : loading pattern file failed")
		return
	}
	//--------------------------------
	patternMap, errPatternMap := arr.ContentToMap(patternContent, chLength, firstCh, lastCh)

	if errPatternMap != nil {
		fmt.Println(errPatternMap)
		return
	}
	//--------------------------------------------
	cliStr, errClistr := cli.ReadCli()

	if errClistr != nil {
		//fmt.Println(errClistr)
		return
	}
	//----------------------------------------------

	finalStr := draw.StrMaker(cliStr, patternMap, chLength)
	fmt.Println(finalStr)

	//Define the ASCII art generator function
	GenerateShadowAsciiArt = func(text string) string {
		return draw.StrMaker(text, patternMap, chLength)
	}
	//-----------------------------------------
}
