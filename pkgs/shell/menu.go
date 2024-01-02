package shell

import (
	"bufio"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func (sh *Shell) getMenu() {
	mainMenu := fyne.NewMainMenu(
		sh.getMenuItems("Каталоги", "catalogs"),
		sh.getMenuItems("Действия", "actions"),
	)
	sh.window.SetMainMenu(mainMenu)
}

func (sh *Shell)getMenuItems(nameMenu, fileMenu string) *fyne.Menu {
	file, _ := os.Open("./pkgs/shell/" + fileMenu + ".txt")
	defer file.Close()
	menuItems := []*fyne.MenuItem{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		switch scanner.Text() {
		case "Контакты":
			menuItems = append(menuItems, fyne.NewMenuItem(scanner.Text(), func() {contacts(sh.window)}))
		default:
			menuItems = append(menuItems, fyne.NewMenuItem(scanner.Text(), nil))
		}
	}
	return fyne.NewMenu(nameMenu, menuItems...)
}

func contacts(w fyne.Window) {
	dialog.ShowInformation("Контакты", "А вы с какой целью, собственно, интересуетесь?", w)
}

