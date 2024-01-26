package shell

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func (sh *Shell) nextCard() {
	sh.id++
	sh.setContent()
}

func (sh *Shell) preCard() {
	sh.id--
	sh.setContent()
}

func (sh *Shell) reverse() {
	path := fmt.Sprintf("pics/%s/back/%d.jpg", sh.catalogName, sh.id)
	reversePic := canvas.NewImageFromFile(checkPic(path))
	reversePic.SetMinSize(fyne.NewSize(437, 310))
	sh.showCustom("", "ok", reversePic)
}

func (sh *Shell) showAllCatalog() {
	allCatalog := [][]*canvas.Image{}
	numPic := 2
	for i := 0; i < sh.lenCatalog/3+1; i++ {
		row := []*canvas.Image{}
		for r := 0; r < 3; r++ {
			if numPic <= sh.lenCatalog + 1 {
				picCard := sh.pic(fmt.Sprintf("pics/%s/%d.jpg", sh.catalogName, numPic))
				picCard.Resize(fyne.NewSize(218, 155))
				row = append(row, picCard)
			}
			numPic++
		}
		allCatalog = append(allCatalog, row)
	}
	fmt.Println(allCatalog)
}