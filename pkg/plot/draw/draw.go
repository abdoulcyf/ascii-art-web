//==================================================
package draw

import (

	// "fmt"
	//"fmt"
	//"fmt"
	"strings"
	//"strconv"
)

/*


 */

func StrMaker(cliStr string, patternMap map[byte][]string, chLength int) string {
	var finalStr string
	//var finelItem string

	cliStrArr := strings.Split(cliStr, "\\n")
	//fmt.Print(">", cliStrArr[1], "<\n")
	//fmt.Println()
	//-------------------------------------------

	for _, strItem := range cliStrArr {
		//finelItem
		if strItem == "" {
			finalStr += "\n"
			//fmt.Println("yes")
			continue
		}
		//-----------------------------------
		for row := 1; row <= chLength; row++ {
			for _, v := range strItem {
				//if v == '\n' {
				//finalStr += "\n"
				//} else {
				finalStr += patternMap[byte(v)][row-1]
				//}

			}
			//finalStr += " "+strconv.Itoa(row)+"\n"
			finalStr += "\n"
		}
		//--------------------------------------
		//finalStr += "\n"
	}

	//----------------------------------------
	//fmt.Println(cliStrArr,len(cliStrArr))

	if ((len(cliStrArr)) == 2) && (cliStrArr[0] == "") && cliStrArr[1] == "" {
		//fmt.Println("yesy") for only "\n"
		finalStr = finalStr[:len(finalStr)-2]
	} else {

		finalStr = finalStr[:len(finalStr)-1]
	}

	return finalStr
}

//------------------------------

//func strMaker_NoEnter()
//==================================================