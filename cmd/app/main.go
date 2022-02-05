package main

import (
	"fmt"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"

	"github.com/mklhmnn/rpn-calc/gui"
)

func main() {
	w := app.NewWindow(
		app.Title("RPN-Calc"),
		app.Size(unit.Dp(400), unit.Dp(300)),
	)
	go loop(w)
	app.Main()
}

func loop(w *app.Window) {
	var ops op.Ops
	cw := gui.NewCalcWindow()
	for e := range w.Events() {
		switch e := e.(type) {
		case system.DestroyEvent:
			if e.Err != nil {
				log.Fatal(e.Err)
			}
			os.Exit(0)
			return

		// https://pkg.go.dev/gioui.org/io/key#Event
		case key.Event:
			fmt.Print(e.State)
			fmt.Print(": ")
			printHex(e.Name)
			if e.State == key.Press && cw.HandleKey(e.Name, e.Modifiers) {
				w.Invalidate()
			}

		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			cw.Render(gtx)

			e.Frame(gtx.Ops)
		}
	}
}

func printHex(s string) {
	switch s {
	case key.NameLeftArrow:
		fmt.Println("left")
	case key.NameRightArrow:
		fmt.Println("right")
	case key.NameEnter:
		fmt.Println("Enter")
	case key.NameReturn:
		fmt.Println("Return")
	case key.NameEscape:
		fmt.Println("Escape")
	case key.NameDeleteBackward:
		fmt.Println("backspace")
	case key.NameDeleteForward:
		fmt.Println("Del")
	}
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
}
