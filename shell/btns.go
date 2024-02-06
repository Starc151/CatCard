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
	var x float32 = 0 
	mainCont := container.NewWithoutLayout(sh.searchBox())
	l := container.NewGridWithColumns(4)
	for i := 0; i <= sh.lenCatalog; i++{
		c := container.New(
			layout.NewStackLayout(),
			sh.button(
				"", 195, 109,
				0, 0, func() {fmt.Println("asd")},
			),
			sh.pic(fmt.Sprintf("pics/%s/%d.jpg", sh.catalogName, numPic)),
		)
		c.Resize(fyne.NewSize(195, 109))
		c.Move(fyne.NewPos(10, x))
		nl := container.NewWithoutLayout(c)
		l.Add(nl)
		numPic++
		x += 119
	}
	
	vScroll := container.NewVScroll(l)
	vScroll.Resize(fyne.NewSize(890, 480))
	vScroll.Move(fyne.NewPos(10, 60))

	mainCont.Add(vScroll)
	// numPic := 2
	// cont := container.NewWithoutLayout(sh.searchBox())
	// grid := container.NewGridWithColumns(4)
	// grid.Resize(fyne.NewSize(890, 580))
	// for i := 0; i < sh.lenCatalog; i++ {
	// 	card := container.New(
	// 		layout.NewStackLayout(),
	// 		sh.button(
	// 			"", 218, 155,
	// 			0, 0, nil,
	// 		),
	// 		sh.pic(fmt.Sprintf("pics/%s/%d.jpg", sh.catalogName, numPic)),
	// 	)
	// 	card.Resize(fyne.NewSize(195, 109))
	// 	columnsGrid  := container.NewWithoutLayout(card)

	// 	grid.Add(columnsGrid )
	// 	numPic++
	// }

	// vScroll := container.NewVScroll(grid)
	// vScroll.Resize(fyne.NewSize(890, 580))
	// vScroll.Move(fyne.NewPos(10, 50))

	// cont.Add(vScroll)
	sh.setContent(mainCont)
}