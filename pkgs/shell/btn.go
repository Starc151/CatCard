package shell

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (sh *Shell) button(text string, width, x float32) *widget.Button {
	sh.btn = &widget.Button{}
	sh.btn.Text = text
	sh.btn.Resize(fyne.NewSize(width, 40))
	sh.btn.Move(fyne.NewPos(x, 400))
	return sh.btn
}