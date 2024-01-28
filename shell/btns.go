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
	cont := container.NewVBox(sh.searchBox())
	numPic := 2
	lenRow := 3
	lenCatalog := sh.lenCatalog/lenRow + 1
	var yPos float32 = 60
	for i := 0; i < lenCatalog; i++ {
		var xPos float32 = 10
		for r := 0; r <  lenRow; r++ {
			if numPic <= sh.lenCatalog + 1 {
				picCard := sh.pic(fmt.Sprintf("pics/%s/%d.jpg", sh.catalogName, numPic))
				btnCard := sh.button(
					"", 218, 155,
					0, 0, nil,
				)
				row := container.NewWithoutLayout(container.New(
					layout.NewStackLayout(), btnCard, picCard,
				))
				row.Resize(fyne.NewSize(276, 195))
				row.Move(fyne.NewPos(xPos, yPos))
				xPos += 296
				cont.Add(row)
			}
			numPic++
		}
		yPos += 215
	}
	// cont.RemoveAll()
	// // cont = container.NewVBox()
	// var ii float32 = 0
	// for i := 0; i < 100; i++ {
	// 	t := canvas.NewText("jhk", nil)
	// 	t.Move(fyne.NewPos(10, ii))
	// 	ii += 20
	// 	cont.Add(t)
	// }
	// contVBox := container.NewVBox(cont)
	
	sh.setContent(cont)
}