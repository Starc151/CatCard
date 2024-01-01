package shell

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Shell struct{
	window	fyne.Window
	card struct {
		pic *canvas.Image
		descript *canvas.Text
		note *canvas.Text
	}
	btn *widget.Button
	search	struct {
		field *widget.Entry
		btn *widget.Button
	}
}

func (sh *Shell) sizeShell() {
	sh.window.Resize(fyne.NewSize(925, 650))
	// sh.window.SetFixedSize(true)
}

func NewShell(app fyne.App) {
	sh := Shell{}
	sh.window = app.NewWindow("Каталога почтовых карточек с техникой")
	sh.getMenu()
	sh.searchBox()
	sh.getPic()
	sh.getDescript()
	sh.getNote()
	
	cont := container.NewWithoutLayout(
		sh.search.field,
		sh.search.btn,
		sh.card.pic,
		sh.card.descript,
		sh.button("Prev", 210, 10),
		sh.button("Next", 210, 237),
		sh.button("Oborot", 437, 470),
		sh.card.note,
	)
	
	sh.window.SetContent(cont)
	sh.sizeShell()
	sh.window.Show()
}
