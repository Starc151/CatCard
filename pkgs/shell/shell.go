package shell

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Shell struct{
	window	fyne.Window
	search	struct {
		field *widget.Entry
		btn *widget.Button
	}
}

func (sh *Shell) sizeShell(){
	sh.window.Resize(fyne.NewSize(950, 750))
	sh.window.SetFixedSize(true)
}

func NewShell(app fyne.App) {
	sh := Shell{}
	sh.window = app.NewWindow("Каталога почтовых карточек с техникой")
	sh.getMenu()
	
	cont := container.NewWithoutLayout(sh.searchBox()...)
	sh.window.SetContent(cont)
	sh.sizeShell()
	sh.window.Show()
}
