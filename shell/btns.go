package shell

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	// "fyne.io/fyne/v2/widget"
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
	numPic := 2
	cont := container.NewWithoutLayout(sh.searchBox())
	gridColumns := container.NewGridWithColumns(4)
	for i := 0; i < sh.lenCatalog; i++{
		pic := sh.pic(fmt.Sprintf("pics/%s/%d.jpg", sh.catalogName, numPic))
		pic.SetMinSize(fyne.NewSize(195, 109))
		contCard := container.New(
			layout.NewStackLayout(),
			sh.button(
				"", 195, 109,
				0, 0, func() {fmt.Println(numPic)},
			),
			pic,
		)
		contCard.Resize(fyne.NewSize(195, 109))
		gridColumns.Add(container.NewWithoutLayout(contCard))
		numPic++
	}
	
	vScroll := container.NewVScroll(gridColumns)
	vScroll.Resize(fyne.NewSize(890, 480))
	vScroll.Move(fyne.NewPos(10, 60))

	cont.Add(vScroll)
	sh.setContent(cont)
}