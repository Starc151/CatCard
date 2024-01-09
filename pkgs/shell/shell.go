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
		id int
		pic *canvas.Image
		descript []*canvas.Text
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

	cont := container.NewWithoutLayout(
		sh.search.field,
		sh.search.btn,
		sh.card.pic,
		sh.getDescript(sh.card.id),
		sh.button("Обратная сторона", 437, 10, 450, nil),
		sh.button("Назад", 210, 10, 510, sh.back),
		sh.button("Вперед", 210, 237, 510, sh.forward),
	)
	sh.window.SetContent(cont)
}
func NewShell(app fyne.App) {
	sh := Shell{}
	sh.card.id = 2
	sh.window = app.NewWindow("Каталога почтовых карточек с техникой")
	sh.buildShel()
	

	
	sh.window.Resize(fyne.NewSize(905, 650))
	sh.window.SetFixedSize(true)
	sh.window.CenterOnScreen()
	sh.window.Show()
}
