package shell

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func (sh *Shell) getDescript() {
	sh.card.descript = &canvas.Text{}
	sh.card.descript.TextSize = 32
	sh.card.descript.Text = "Description"
	sh.card.descript.Resize(fyne.NewSize(437, 310))
	sh.card.descript.Move(fyne.NewPos(470, 100))
}
