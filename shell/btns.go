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
	allCatalog := [][3]int{} //*canvas.Image{}
	
	for i := 0; i < sh.lenCatalog/3+1; i++{
		allCatalog = append(allCatalog, [3]int{1,2,3})
	}
	

	fmt.Println(allCatalog)
}

// sh.lenCatalog/3+0