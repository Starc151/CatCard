package shell

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func (sh *Shell) pic(path string) *canvas.Image {
	picCard := canvas.NewImageFromFile(checkPic(path))
	picCard.Resize(fyne.NewSize(437, 310))
	picCard.Move(fyne.NewPos(10, 100))
	return picCard
}

func checkPic(path string) string {
	if _, err := os.Stat(path); err != nil{
		if os.IsNotExist(err) {
			return "img/noPhoto.jpg"
		}
	}
	return path
}