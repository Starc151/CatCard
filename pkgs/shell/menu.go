package shell

import (
	"bufio"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func (sh *Shell) menu() {
	mainMenu := fyne.NewMainMenu(
		sh.menuItems("Каталоги", "catalogs"),
		sh.menuItems("Действия", "actions"),
	)
	sh.window.SetMainMenu(mainMenu)
}

func (sh *Shell) menuItems(nameMenu, fileMenu string) *fyne.Menu {
	file, _ := os.Open("./pkgs/shell/" + fileMenu + ".txt")
	defer file.Close()
	menuItems := []*fyne.MenuItem{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		switch scanner.Text() {
		case "Контакты":
			menuItems = append(menuItems, fyne.NewMenuItem(scanner.Text(), func() {
				sh.contacts()
			}))
		default:
			menuItems = append(menuItems, fyne.NewMenuItem(scanner.Text(), nil))
		}
	}
	return fyne.NewMenu(nameMenu, menuItems...)
}

func (sh *Shell) contacts() {
	dialog.ShowInformation(
		"Контакты",
		"А вы с какой целью, собственно, интересуетесь?",
		sh.window)
}

