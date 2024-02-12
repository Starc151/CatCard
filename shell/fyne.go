package shell

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (sh *Shell) button(text string, w, h, x, y float32, f func()) *widget.Button {
	sh.btn = &widget.Button{}
	sh.btn.Text = text
	sh.btn.OnTapped = f
	sh.btn.Resize(fyne.NewSize(w, h))
	sh.btn.Move(fyne.NewPos(x, y))
	return sh.btn
}

func vScroll (w, h, x, y float32, cont *fyne.Container) fyne.CanvasObject {
	vScroll := container.NewVScroll(cont)
	vScroll.Resize(fyne.NewSize(w, h))
	vScroll.Move(fyne.NewPos(x, y))
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

func (sh *Shell) showConfirm() {
	btn := sh.button("Закрыть", 100, 50, 0, 0, sh.window.Close)
	txt := canvas.NewText("Файл с данными не найден! Приложение будет закрыто!", nil)
	box := container.NewCenter(container.NewVBox(txt, btn))
	box.Resize(fyne.NewSize(865, 200))
	box.Move(fyne.NewPos(20, 100))

	cont := container.NewWithoutLayout(box)
	sh.setContent(cont)
}