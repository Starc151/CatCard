package shell

import (
	"bufio"
	"os"

	"fyne.io/fyne/v2"
)

func getMenuItems(nameMenu, fileMenu string) *fyne.Menu {
	file, _ := os.Open("./pkgs/shell/" + fileMenu + ".txt")
	defer file.Close()
	menuItems := []*fyne.MenuItem{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		menuItems = append(menuItems, fyne.NewMenuItem(scanner.Text(), func() {}))
	}
	return fyne.NewMenu(nameMenu, menuItems...)
}

func (sh *Shell) getMenu() {
	mainMenu := fyne.NewMainMenu(
		getMenuItems("Каталоги", "catalogs"),
		getMenuItems("Действия", "actions"),
	)
	sh.window.SetMainMenu(mainMenu)
}
