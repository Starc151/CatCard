package shell

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func (sh *Shell) pic() {
	sh.picCard = &canvas.Image{}
	if _, err := os.Stat(fmt.Sprintf("pics/%d.jpg", sh.id)); err != nil{
		if os.IsNotExist(err) {
			sh.picCard.File = "img/noPhoto.jpg"
		}
	} else {
		sh.picCard.File = fmt.Sprintf("pics/%d.jpg", sh.id)
	}
	sh.picCard.Resize(fyne.NewSize(437, 310))
	sh.picCard.Move(fyne.NewPos(10, 100))
}