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
	}
	btn *widget.Button
	search	struct {
		field *widget.Entry
		btn *widget.Button
	}
}
func (sh *Shell) buildShel() {
	sh.getMenu()
	sh.searchBox()
	sh.getPic()
	sh.getDescript()
}
func NewShell(app fyne.App) {
	sh := Shell{}
	sh.window = app.NewWindow("Каталога почтовых карточек с техникой")
	sh.buildShel()
	
	cont := container.NewWithoutLayout(
		sh.search.field,
		sh.search.btn,
		sh.card.pic,
		sh.card.descript,
		sh.button("Обратная сторона", 437, 10, 450),
		sh.button("Вперед", 210, 10, 510),
		sh.button("Назад", 210, 237, 510),
	)
	
	sh.window.Resize(fyne.NewSize(905, 650))
	sh.window.SetFixedSize(true)
	sh.window.CenterOnScreen()
	sh.window.SetContent(cont)
	sh.window.Show()
}
