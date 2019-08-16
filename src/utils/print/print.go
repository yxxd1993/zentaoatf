package printUtils

import (
	"encoding/json"
	"flag"
	"fmt"
	logUtils "github.com/easysoft/zentaoatf/src/utils/log"
	stringUtils "github.com/easysoft/zentaoatf/src/utils/string"
	"github.com/easysoft/zentaoatf/src/utils/vari"
	"github.com/fatih/color"
)

func PrintUsage(flagSet flag.FlagSet) {
	PrintUsageWithSpaceLine(flagSet, true)
}
func PrintUsageWithSpaceLine(flagSet flag.FlagSet, spaceLine bool) {
	prefix := ""
	if spaceLine {
		prefix = "\n"
	}

	fmt.Printf("%s %s \n", prefix, color.CyanString(flagSet.Name()))
	flagSet.PrintDefaults()
}

func PrintSample() {
	fmt.Printf("\nSample to use: \n")
	fmt.Printf("TODO... \n")
}

func PrintToSide(msg string) {
	if !vari.RunFromCui {
		fmt.Println(msg)
		return
	}
	slideView, _ := vari.Cui.View("side")
	slideView.Clear()

	fmt.Fprintln(slideView, msg)
}

func PrintToMainNoScroll(msg string) {
	if !vari.RunFromCui {
		fmt.Println(msg)
		return
	}
	mainView, _ := vari.Cui.View("main")
	mainView.Clear()

	fmt.Fprintln(mainView, msg)
}

func PrintToCmd(msg string) {
	if !vari.RunFromCui {
		fmt.Println(msg)
		return
	}
	cmdView, _ := vari.Cui.View("cmd")
	_, _ = fmt.Fprintln(cmdView, msg)

	logUtils.AdjustOrigin("cmd")
}
func PrintStructToCmd(obj interface{}) {
	str := stringUtils.StructToStr(obj)
	PrintToCmd(str)
}

func ClearSide() {
	slideView, _ := vari.Cui.View("side")
	slideView.Clear()
}

func PrintUnicode(str []byte) {
	var a interface{}
	err := json.Unmarshal(str, &a)
	if err == nil {
		if !vari.RunFromCui {
			fmt.Println(a)
		}
	} else {
		PrintToCmd(string(str))
	}

}
