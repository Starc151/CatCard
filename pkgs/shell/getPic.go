package shell

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func (sh *Shell) getPic() {
	sh.card.pic = &canvas.Image{}
	sh.card.pic.File = "pics/test.jpg"
	sh.card.pic.Resize(fyne.NewSize(437, 310))
	sh.card.pic.Move(fyne.NewPos(10, 50))
}