package shell

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)
func (sh *Shell) searchBox() []fyne.CanvasObject{
	sh.search.field = &widget.Entry{}
	sh.search.field.SetPlaceHolder("Поиск...")
	sh.search.field.Resize(fyne.NewSize(890, 40))

	sh.search.btn = &widget.Button{}
	icon, _ := fyne.LoadResourceFromPath("img/search.png")
	sh.search.btn.SetIcon(icon)
	sh.search.btn.Resize(fyne.NewSize(40, 40))
	sh.search.btn.Move(fyne.NewPos(900, 0))

	search := []fyne.CanvasObject{sh.search.field, sh.search.btn}
	return search
}