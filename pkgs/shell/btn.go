package shell

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (sh *Shell) button(text string, width, x, y float32, f func()) *widget.Button {
	sh.btn = &widget.Button{}
	sh.btn.Text = text
	sh.btn.OnTapped = f
	sh.btn.Resize(fyne.NewSize(width, 40))
	sh.btn.Move(fyne.NewPos(x, y))
	return sh.btn
}
