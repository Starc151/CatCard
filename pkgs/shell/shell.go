package shell

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Card struct {
	id int
	pic *canvas.Image
	descript []*canvas.Text
}

type Lib struct {
	NumRowsXls int
	NumPic int
	Card
}

type Shell struct{
	window	fyne.Window
	entry *widget.Entry
	btn *widget.Button
}

func (sh *Shell) build() {
	l := Lib{}
	l.NumPic = 5
	l.NumRowsXls = 5
	l.id = 5
	sh.menu()
	searchBox := sh.searchBox()
	backBtn := sh.button("Назад", 210, 10, 510, nil)
	if l.id == 2 {
		backBtn.Disable()
	}
	nextBtn := 	sh.button("Вперед", 210, 237, 510, nil)
	if l.id == max(l.NumPic, l.NumRowsXls) {
		nextBtn.Disable()
	}

	

	cont := container.NewWithoutLayout(
		searchBox,
		sh.button("Обратная сторона", 437, 10, 450, nil),
		backBtn,
		nextBtn,		
	)
	sh.window.SetContent(cont)
}
func NewShell(app fyne.App) {
	sh := Shell{}
	sh.window = app.NewWindow("Каталога почтовых карточек с техникой")
	sh.build()
	sh.window.Resize(fyne.NewSize(905, 650))
	sh.window.SetFixedSize(true)
	sh.window.CenterOnScreen()
	sh.window.Show()
}
