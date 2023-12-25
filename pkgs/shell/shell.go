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

func (sh *Shell) fieldSearch() []fyne.CanvasObject{
	sh.search.field = &widget.Entry{}
	sh.search.field.SetPlaceHolder("Поиск...")
	sh.search.field.Resize(fyne.NewSize(890, 40))

	sh.search.btn = &widget.Button{}
	icon, _ := fyne.LoadResourceFromPath("img/search.png")
	sh.search.btn.SetIcon(icon)
	sh.search.btn.Resize(fyne.NewSize(40, 40))
	sh.search.btn.Move(fyne.NewPos(900, 0))

	search := []fyne.CanvasObject{sh.search.field, sh.search.btn}
	return search
}

func NewShell(app fyne.App) {
	sh := Shell{}
	sh.window = app.NewWindow("Каталога почтовых карточек с техникой")
	sh.getMenu()
	
	cont := container.NewWithoutLayout(sh.fieldSearch()...)
	sh.window.SetContent(cont)
	sh.sizeShell()
	sh.window.Show()
}