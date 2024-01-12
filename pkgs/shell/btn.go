package shell

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
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

func (sh *Shell) nextCard() {
	sh.id++
	sh.setContent()
}

func (sh *Shell) preCard() {
	sh.id--
	sh.setContent()
}

func (sh *Shell) reverse() {
	path := fmt.Sprintf("pics/back/%d.jpg", sh.id)
	reversePic := canvas.NewImageFromFile(checkPic(path))
	reversePic.SetMinSize(fyne.NewSize(437, 310))
	dialog.ShowCustom("", "ok", reversePic, sh.window)
}