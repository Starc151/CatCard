package shell

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func (sh *Shell) pic() {
	sh.picCard = &canvas.Image{}
	sh.picCard.File = checkPic(fmt.Sprintf("pics/%s/%d.jpg", sh.catalogName, sh.id))
	sh.picCard.Resize(fyne.NewSize(437, 310))
	sh.picCard.Move(fyne.NewPos(10, 100))
}

func checkPic(path string) string {
	if _, err := os.Stat(path); err != nil{
		if os.IsNotExist(err) {
			return "img/noPhoto.jpg"
		}
	}
	return path
}