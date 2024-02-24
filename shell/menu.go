package shell

import (
	"bufio"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func (sh *Shell) setCatalogs() []*fyne.MenuItem {
	catalogs := ""
	menuCatalogs := []*fyne.MenuItem{}
	for _, sheetName := range sh.xlFile.GetSheetList() {
		catalogs += sheetName+"\n"
		temp := sheetName
		menuCatalogs = append(menuCatalogs, fyne.NewMenuItem(sheetName, func() {sh.setCatalogName(temp)}))
	}
	os.WriteFile("shell/menuItems/catalogs.txt", []byte(catalogs), 0777)
	return menuCatalogs
}

func (sh *Shell) setCatalogName(catalogName string) {
	sh.catalogName = catalogName
	sh.id = 2
	sh.showAllCatalog()
}

func (sh *Shell) menu() {
	if sh.xlFile != nil {
		mainMenu := fyne.NewMainMenu(
			fyne.NewMenu("Справка", sh.menuItems("Справка", "help")...),
			fyne.NewMenu("Каталоги", sh.setCatalogs()...),
		)
		sh.window.SetMainMenu(mainMenu)
	} else {
		mainMenu := fyne.NewMainMenu(
			fyne.NewMenu("Справка", sh.menuItems("Справка", "help")...),
		)
		sh.window.SetMainMenu(mainMenu)
	}
}

func (sh *Shell) menuItems(nameMenu, fileMenu string) []*fyne.MenuItem {
	menuItems := []*fyne.MenuItem{}

	file, _ := os.Open("shell/menuItems/" + fileMenu + ".txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	for scanner.Scan() {
		switch scanner.Text() {
		case "Контакты":
			menuItems = append(menuItems, fyne.NewMenuItem(scanner.Text(), func() {
				sh.contacts()
			}))
		case "Посмотреть справку":
			menuItems = append(menuItems, fyne.NewMenuItem(scanner.Text(), func() {
				sh.help()
			}))
		case "О программе":
			menuItems = append(menuItems, fyne.NewMenuItem(scanner.Text(), func() {
				sh.about()
			}))
		}
	}
	return menuItems
}

func (sh *Shell) contacts() {
	sh.showInformation(		"Контакты", "А вы с какой целью, собственно, интересуетесь?")
}

func (sh *Shell) help(){
	vBox := container.NewVBox()
	layout := container.NewWithoutLayout()

	file, _ := os.Open("shell/help/briefInformation.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	for scanner.Scan() {
		txt := canvas.Text{}
		txt.TextSize = 14
		txt.Text = scanner.Text()
		vBox.Add(&txt)
	}
	layout.Add(vBox)
	sh.showCustom("Справка", "ok", layout)
}

func (sh *Shell) about() {
	sh.showInformation(		"О программе", "Каталог почтовых карточек с различной техникой\n"+
		"Создан по заказу Npc prodaction")
}
