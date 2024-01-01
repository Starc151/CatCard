package shell

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)
func (sh *Shell) searchBox() {
	sh.search.field = &widget.Entry{}
	sh.search.field.SetPlaceHolder("Поиск...")
	sh.search.field.Resize(fyne.NewSize(830, 40))
	sh.search.field.Move(fyne.NewPos(10, 0))

	sh.search.btn = &widget.Button{}
	icon, _ := fyne.LoadResourceFromPath("img/search.png")
	sh.search.btn.SetIcon(icon)
	sh.search.btn.Resize(fyne.NewSize(40, 40))
	sh.search.btn.Move(fyne.NewPos(850, 0))
}