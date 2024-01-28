package shell

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)
func (sh *Shell) searchBox() *fyne.Container {
	sh.entry = &widget.Entry{}
	sh.entry.SetPlaceHolder("Поиск...")
	sh.entry.Resize(fyne.NewSize(830, 40))
	sh.entry.Move(fyne.NewPos(10, 0))

	btn := sh.button(
			"",
			40, 40,
			850, 0,
			nil,
		)
	icon, _ := fyne.LoadResourceFromPath("img/search.png")
	btn.Icon = icon

	return container.NewWithoutLayout(sh.entry, btn)
}