package shell

import (
	"bufio"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func (sh *Shell) getMenu() {
	mainMenu := fyne.NewMainMenu(
		getMenuItems("Каталоги", "catalogs", sh.window),
		getMenuItems("Действия", "actions", sh.window),
	)
	sh.window.SetMainMenu(mainMenu)
}

func getMenuItems(nameMenu, fileMenu string, w fyne.Window) *fyne.Menu {
	file, _ := os.Open("./pkgs/shell/" + fileMenu + ".txt")
	defer file.Close()
	menuItems := []*fyne.MenuItem{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		switch scanner.Text() {
		case "Контакты":
			menuItems = append(menuItems, fyne.NewMenuItem(scanner.Text(), func() {contacts(w)}))
		default:
			menuItems = append(menuItems, fyne.NewMenuItem(scanner.Text(), nil))
		}
	}
	return fyne.NewMenu(nameMenu, menuItems...)
}

func contacts(w fyne.Window) {
	dialog.ShowInformation("Контакты", "А вы с какой целью, собственно, интересуетесь?", w)
}

