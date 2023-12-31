package shell

import "fyne.io/fyne/v2"

type Shell struct{
	window    fyne.Window
}
func (sh *Shell) sizeShell(){
	sh.window.Resize(fyne.NewSize(950, 750))
}

func NewShell(app fyne.App) {
	sh := Shell{}
	sh.window = app.NewWindow("Каталога почтовых карточек с техникой")

	sh.sizeShell()
	sh.window.Show()
}