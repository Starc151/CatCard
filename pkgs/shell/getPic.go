package shell

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func (sh *Shell) getPic() {
	sh.card.pic = &canvas.Image{}
	if _, err := os.Stat(fmt.Sprintf("pics/%d.jpg", sh.card.id)); err != nil{
		if os.IsNotExist(err) {
			sh.card.pic.File = "img/noPhoto.jpg"
		}
	} else {
		sh.card.pic.File = fmt.Sprintf("pics/%d.jpg", sh.card.id)
	}
	sh.card.pic.Resize(fyne.NewSize(437, 310))
	sh.card.pic.Move(fyne.NewPos(10, 100))
}