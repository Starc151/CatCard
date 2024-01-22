package shell

import (
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/xuri/excelize/v2"
)

type card struct {
	id int
	picCard *canvas.Image
	descriptCard []*canvas.Text
}

type catalog struct {
	xlFile *excelize.File
	catalogName string
	lenCatalog int
	card
}

type Shell struct{
	window	fyne.Window
	entry *widget.Entry
	btn *widget.Button
	catalog
}

func NewShell(app fyne.App) {
	sh := Shell{}
	sh.window = app.NewWindow("Каталога почтовых карточек с техникой")
	sh.xlFile, _ = excelize.OpenFile("slu/data.xlsx")
	sh.id = 2
	sh.catalogName = sh.xlFile.GetSheetList()[0]
	sh.menu()
	sh.setContent()
	sh.window.Resize(fyne.NewSize(905, 650))
	sh.window.SetFixedSize(true)
	sh.window.CenterOnScreen()
	sh.window.Show()
}

func (sh *Shell) getlenCatalog() {
	rows, _ := sh.xlFile.GetRows(sh.catalogName)
	numsPics := 0
	numObjcts, _ := os.ReadDir("pics/" + sh.catalogName)
	for  _, isJgp := range numObjcts {
		if !isJgp.IsDir() {
			if filepath.Ext(isJgp.Name()) == ".jpg" {
				numsPics++
			}
		}
	}
	sh.lenCatalog = max(numsPics, len(rows))
}

func (sh *Shell) setContent() {
	sh.getlenCatalog()
	searchBox := sh.searchBox()
	sh.pic()

	reverseBtn := sh.button("Обратная сторона", 437, 10, 450, sh.reverse)
	preBtn := sh.button("Назад", 210, 10, 510, sh.preCard)
	if sh.id <= 2 {
		preBtn.Disable()
	}
	nextBtn := 	sh.button("Вперед", 210, 237, 510, sh.nextCard)
	if sh.id >= sh.lenCatalog {
		nextBtn.Disable()
	}
	cont := container.NewWithoutLayout(
		searchBox,
		sh.descript(),
		sh.picCard,
		reverseBtn,
		preBtn,
		nextBtn,		
	)
	sh.window.SetContent(cont)
}

