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
	sh.card.descript.Move(fyne.NewPos(470, 50))
}

func (sh *Shell) getNote() {
	sh.card.note = &canvas.Text{}
	sh.card.note.TextSize = 32
	sh.card.note.Text = "Note"
	sh.card.note.Resize(fyne.NewSize(897, 200))
	sh.card.note.Move(fyne.NewPos(10, 400))
}