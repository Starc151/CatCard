package shell

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Card struct {
	id int
	picCard *canvas.Image
	// descript []*canvas.Text
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
	Lib
}

func (sh *Shell) build() {
	sh.NumRowsXls = 5
	sh.menu()
	sh.setContent()
}

func (sh *Shell) setContent() {
	searchBox := sh.searchBox()
	sh.pic()
	reverseBtn := sh.button("Обратная сторона", 437, 10, 450, sh.reverse)
	preBtn := sh.button("Назад", 210, 10, 510, sh.preCard)
	if sh.id == 2 {
		preBtn.Disable()
	}
	nextBtn := 	sh.button("Вперед", 210, 237, 510, sh.nextCard)
	if sh.id == max(sh.NumPic, sh.NumRowsXls) {
		nextBtn.Disable()
	}

	cont := container.NewWithoutLayout(
		searchBox,
		sh.picCard,
		sh.button("Обратная сторона", 437, 10, 450, nil),
		reverseBtn,
		preBtn,
		nextBtn,		
	)
	sh.window.SetContent(cont)
}
func NewShell(app fyne.App) {
	sh := Shell{}
	sh.id = 2
	sh.window = app.NewWindow("Каталога почтовых карточек с техникой")
	sh.build()
	sh.window.Resize(fyne.NewSize(905, 650))
	sh.window.SetFixedSize(true)
	sh.window.CenterOnScreen()
	sh.window.Show()
}
