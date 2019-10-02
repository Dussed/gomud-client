package drawing

import (
	"fmt"

	"github.com/dussed/gomud-client/util"
)

type terminalMargins struct {
	top, left, bottom, right int
}

type TerminalWindow struct {
	terminalMargins
	width, height int
}

func DrawDebug() {
	width, height, err := util.GetDimensions()

	if err != nil {
		fmt.Println("err occured", err.Error())
	}

	thisWindow := TerminalWindow{
		width:  width,
		height: height,
	}

	for tbLineDraw := thisWindow.height; tbLineDraw > 0; tbLineDraw-- {

		for lrLineDraw := thisWindow.width; lrLineDraw > 0; lrLineDraw-- {
			fmt.Print(" ")
		}

		fmt.Print("\n")
	}

}
