package shell

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
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

func vScroll (w, h float32, cont *fyne.Container) fyne.CanvasObject {
	vScroll := container.NewVScroll(cont)
	vScroll.Resize(fyne.NewSize(w, h))
	return vScroll
}

func (sh *Shell) showInformation(title, msg string) {
	dialog.ShowInformation(
		title,
		msg,
		sh.window)
}

func (sh *Shell) showCustom(title, btnTxt string, content fyne.CanvasObject) {
	dialog.ShowCustom(
		title,
		btnTxt,
		content,
		sh.window)
}
