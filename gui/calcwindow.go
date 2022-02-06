package gui

import (
	"strconv"
	"strings"

	"gioui.org/font/gofont"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/text"

	"gioui.org/unit"
	"gioui.org/widget/material"

	"github.com/mklhmnn/rpn-calc/internal"
)

type CalcWindow struct {
	theme      *material.Theme
	calculator *internal.Stack
	list       layout.List
	input      string
}

func NewCalcWindow() *CalcWindow {
	var this = &CalcWindow{}
	this.theme = material.NewTheme((gofont.Collection()))
	this.calculator = internal.NewStack()
	this.list = layout.List{
		Axis:        layout.Vertical,
		ScrollToEnd: true,
	}
	return this
}

func (this *CalcWindow) Render(gtx layout.Context) {
	lines := createInput()
	this.calculator.Foreach(func(v float64) {
		lines = append(lines, strconv.FormatFloat(v, 'g', 9, 64))
	})
	if len(this.input) > 0 {
		lines = append(lines, this.input+"|")
	} else if len(lines) == 0 {
		lines = append(lines, "?")
	}

	layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.SE.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return this.list.Layout(gtx, len(lines), func(gtx layout.Context, idx int) layout.Dimensions {
				line := lines[idx]
				title := material.H2(this.theme, line)
				title.Alignment = text.End
				return title.Layout(gtx)
			})
		})
	})
}

func createInput() []string {
	return make([]string, 0, 10)
}

func (this *CalcWindow) HandleKey(s string) bool {
	switch s {
	case "M", "N":
		return redrawOrBeep(this.negate())
	case "S": {
		this.finish()
		return redrawOrBeep(this.calculator.Sqrt())
	}
	case "X": {
		this.finish()
		return redrawOrBeep(this.calculator.Invert())
	}
	case key.NameLeftArrow:
		return this.swap()
	case key.NameRightArrow:
		return this.swap()
	case key.NameEnter:
		return this.enter()
	case key.NameReturn:
		return this.enter()
	case key.NameDeleteBackward:
		return this.backspace()
	case key.NameEscape:
		return this.delete()
	case key.NameDeleteForward:
		return this.delete()
	}
	return false
}

func (this *CalcWindow) HandleText(s string) bool {
	if len(s) == 1 {
		chr := s[0]
		if '0' <= chr && chr <= '9' {
			this.input = this.input + s
			return true
		}
		if chr == '.' {
			if !strings.Contains(this.input, s) {
				this.input = this.input + s
				return true
			}
			return false
		}
		if chr == '+' {
			this.finish()
			return redrawOrBeep(this.calculator.Add())
		}
		if chr == '-' {
			this.finish()
			return redrawOrBeep(this.calculator.Substract())
		}
		if chr == '*' {
			this.finish()
			return redrawOrBeep(this.calculator.Multiply())
		}
		if chr == '/' {
			this.finish()
			return redrawOrBeep(this.calculator.Divide())
		}
	}
	return false
}

func redrawOrBeep(ok bool) bool {
	// todo beep if !ok
	return ok
}

func (this *CalcWindow) enter() bool {
	return redrawOrBeep(this.finish() || this.calculator.Duplicate())
}

func (this *CalcWindow) backspace() bool {
	var len = len(this.input)
	var ok = len > 0
	if ok {
		this.input = this.input[0 : len-1]
	}
	return redrawOrBeep(ok)
}

func (this *CalcWindow) delete() bool {
	return redrawOrBeep(this.finish() || this.calculator.Drop())
}

func (this *CalcWindow) finish() bool {
	if len(this.input) == 0 {
		return false
	}

	value, error := strconv.ParseFloat(this.input, 64)
	if error == nil {
		this.calculator.Push(value)
	}
	return this.clear()
}

func (this *CalcWindow) swap() bool {
	this.finish()
	return redrawOrBeep(this.calculator.Swap())
}

func (this *CalcWindow) clear() bool {
	this.input = ""
	return true
}

func (this *CalcWindow) negate() bool {
	if len(this.input) == 0 {
		return this.calculator.Negate()
	}

	if this.input[0] == '-' {
		this.input = this.input[1:]
	} else {
		this.input = "-" + this.input
	}
	return true
}
