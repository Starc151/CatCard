package shell

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func (sh *Shell) nextCard() {
	sh.id++
	sh.showCard()
}

func (sh *Shell) preCard() {
	sh.id--
	sh.showCard()
}

func (sh *Shell) reverse() {
	path := fmt.Sprintf("pics/%s/back/%d.jpg", sh.catalogName, sh.id)
	reversePic := canvas.NewImageFromFile(checkPic(path))
	reversePic.SetMinSize(fyne.NewSize(437, 310))
	sh.showCustom("", "ok", reversePic)
}

func (sh *Shell) showAllCatalog() {
	sh.getlenCatalog()
	catalogNameText := canvas.NewText(sh.catalogName, nil)
	catalogNameText.Alignment = fyne.TextAlignCenter
	catalogName := container.NewHBox(
		container.NewAdaptiveGrid(3,
			catalogNameText,
			catalogNameText,
			catalogNameText,
		),
	)

	
	catalogName.Move(fyne.NewPos(10, 60))
	cont := container.NewWithoutLayout(sh.searchBox(), catalogName)
	if sh.lenCatalog == 0 {
		cont = sh.emptyCont()
	}
	gridColumns := container.NewGridWithColumns(4)
	for numPic := 2; numPic <= sh.lenCatalog+1; numPic++{
		pic := sh.pic(fmt.Sprintf("pics/%s/%d.jpg", sh.catalogName, numPic))
		id := numPic
		pic.SetMinSize(fyne.NewSize(195, 109))
		btn := sh.button(
			"", 195, 109,
			0, 0, func() {
				sh.id = id
				sh.showCard()
			},
		)
		contCard := container.New(
			layout.NewStackLayout(),
			btn,
			pic,
		)
		contCard.Resize(fyne.NewSize(195, 109))
		gridColumns.Add(container.NewWithoutLayout(contCard))
	}
	
	vScroll := container.NewVScroll(gridColumns)
	vScroll.Resize(fyne.NewSize(890, 440))
	vScroll.Move(fyne.NewPos(10, 100))
	cont.Add(vScroll)
	sh.setContent(cont)
}
