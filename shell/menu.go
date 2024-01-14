package shell

import (
	"bufio"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func (sh *Shell) menu() {
	mainMenu := fyne.NewMainMenu(
		// sh.menuItems("Каталоги", "catalogs"),
		sh.menuItems("Действия", "actions"),
		sh.menuItems("Справка", "help"),
	)
	sh.window.SetMainMenu(mainMenu)
}

func (sh *Shell) menuItems(nameMenu, fileMenu string) *fyne.Menu {
	file, _ := os.Open("shell/menuItems/" + fileMenu + ".txt")
	defer file.Close()
	menuItems := []*fyne.MenuItem{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		switch scanner.Text() {
		case "Контакты":
			menuItems = append(menuItems, fyne.NewMenuItem(scanner.Text(), func() {
				sh.contacts()
			}))
		case "Посмотреть правку":
			menuItems = append(menuItems, fyne.NewMenuItem(scanner.Text(), func() {
				sh.help()
			}))
		case "О программе":
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
	sh.showInformation(		"Контакты", "А вы с какой целью, собственно, интересуетесь?",)
}

func (sh *Shell) help(){
	vBox := container.NewVBox()
	vBox.Resize(fyne.NewSize(240, 490))
	layout := container.NewWithoutLayout()

	file, _ := os.Open("shell/help/briefInformation.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := canvas.Text{}
		txt.TextSize = 20
		txt.Text = scanner.Text()
		vBox.Add(&txt)
	}
	layout.Add(vBox)
	sh.showCustom("Справка", "ok", layout)
}

