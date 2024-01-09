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
	if sh.card.id == 2 && text == "Назад"{
		sh.btn.Disable()
	}
	return sh.btn
}
func (sh *Shell) back() {
	if sh.card.id > 2 {
		sh.card.id--
		sh.buildShel()
	}
}

func (sh *Shell) forward() {
	sh.card.id++
	sh.buildShel()
}