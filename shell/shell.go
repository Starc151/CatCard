package shell

import (
	"fmt"
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
}

type catalog struct {
	xlFile      *excelize.File
	catalogName string
	lenCatalog  int
	card
}

type Shell struct {
	window fyne.Window
	entry  *widget.Entry
	btn    *widget.Button
	catalog
}

func NewShell(app fyne.App) {
	sh := Shell{}
	sh.window = app.NewWindow("Каталога почтовых карточек с техникой")
	sh.xlFile, _ = excelize.OpenFile("slu/data.xlsx")
	sh.menu()
	if sh.xlFile == nil {
		sh.showConfirm()
		sh.window.Show()
	} else {
		sh.catalogName = sh.xlFile.GetSheetList()[0]
		sh.showAllCatalog()
	}
	sh.window.Resize(fyne.NewSize(905, 600))
	sh.window.SetFixedSize(true)
	sh.window.CenterOnScreen()
	sh.window.Show()
}

func (sh *Shell) getlenCatalog() {
	rows, _ := sh.xlFile.GetRows(sh.catalogName)
	numsPics := 0
	numObjcts, _ := os.ReadDir("pics/" + sh.catalogName)
	for _, isJgp := range numObjcts {
		if filepath.Ext(isJgp.Name()) == ".jpg" {
			numsPics++
		}
	}
	sh.lenCatalog = max(numsPics, len(rows)-1)
}

func (sh *Shell) emptyCont() *fyne.Container {

	cont := container.NewVBox(
		sh.searchBox(),
		container.NewCenter(
			container.NewVBox(
				canvas.NewText("", nil),
				canvas.NewText("", nil),
				canvas.NewText("", nil),
				canvas.NewText("", nil),
				container.NewCenter(canvas.NewText(
					fmt.Sprintf("Каталог \"%s\" пустой!", sh.catalogName),
					nil),
				),
				container.NewCenter(canvas.NewText("Выберите другой каталог,", nil)),
				container.NewCenter(canvas.NewText("воспользуйтесь поиском", nil)),
				container.NewCenter(canvas.NewText("или заполните exel-файл", nil)),
			),
		),
	)
	return cont
}

func (sh *Shell) showCard() {
	sh.getlenCatalog()

	reverseBtn := sh.button("Обратная сторона", 437, 40, 10, 450, sh.reverse)
	preBtn := sh.button("Назад", 210, 40, 10, 510, sh.preCard)
	if sh.id <= 2 {
		preBtn.Disable()
	}
	editBtn := 	sh.button("Редактировать", 416, 40, 470, 450, sh.editCard)
	nextBtn := sh.button("Вперед", 210, 40, 237, 510, sh.nextCard)
	if sh.id > sh.lenCatalog {
		nextBtn.Disable()
	}
	allCatalog := sh.button(
		fmt.Sprintf("Посмотреть весь каталог \"%s\" ", sh.catalogName),
		880, 40,
		10, 50,
		sh.showAllCatalog,
	)

	cont := container.NewWithoutLayout(
		sh.searchBox(),
		allCatalog,
		vScroll(416, 315, 470, 100, sh.descript()),
		sh.pic(fmt.Sprintf("pics/%s/%d.jpg", sh.catalogName, sh.id)),
		reverseBtn, preBtn, editBtn, nextBtn,
	)
	sh.setContent(cont)
}

func (sh *Shell) setContent(cont fyne.CanvasObject) {
	sh.window.SetContent(cont)
}

func (sh *Shell) showConfirm() {
	btn := sh.button("Закрыть", 100, 50, 0, 0, sh.window.Close)
	txt := canvas.NewText("Файл с данными не найден! Приложение будет закрыто!", nil)
	box := container.NewCenter(container.NewVBox(txt, btn))
	box.Resize(fyne.NewSize(865, 200))
	box.Move(fyne.NewPos(20, 100))

	cont := container.NewWithoutLayout(box)
	sh.setContent(cont)
}