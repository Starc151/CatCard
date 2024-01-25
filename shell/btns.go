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
	n := 2
	for k := 0; k < sh.lenCatalog/4+1; k++ {
		temp := [3]int{}
		for i := 0; i < 3; i++ {
			temp[i] = n
			n++
		}
		allCatalog = append(allCatalog, temp)
	}

	

	fmt.Println(allCatalog)
}

// sh.lenCatalog/3+0