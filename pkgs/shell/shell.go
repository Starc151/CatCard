package shell

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/tealeg/xlsx"
)

type card struct {
	id int
	picCard *canvas.Image
	descriptCard []*canvas.Text
}

type lib struct {
	numRowsXls int
	numPic int
	card
}

type Shell struct{
	window	fyne.Window
	entry *widget.Entry
	btn *widget.Button
	lib
}

func (sh *Shell) libData() {
	xlFile, _ := xlsx.OpenFile("./slu/data.xlsx")
	sh.numRowsXls = len(xlFile.Sheets[0].Rows)
	t, _ := os.ReadDir("pics")
	sh.numPic = len(t)-1
}

func (sh *Shell) build() {
	sh.menu()
	sh.setContent()
	sh.libData()
}

func (sh *Shell) setContent() {
	// searchBox := sh.searchBox()
	sh.pic()

	reverseBtn := sh.button("Обратная сторона", 437, 10, 450, sh.reverse)
	preBtn := sh.button("Назад", 210, 10, 510, sh.preCard)
	if sh.id == 2 {
		preBtn.Disable()
	}
	nextBtn := 	sh.button("Вперед", 210, 237, 510, sh.nextCard)
	if sh.id == max(sh.numPic, sh.numRowsXls) {
		nextBtn.Disable()
	}
	cont := container.NewWithoutLayout(
		// searchBox,
		sh.descript(),
		sh.picCard,
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
