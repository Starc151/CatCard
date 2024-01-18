package shell

import (
	"bufio"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func (sh *Shell) setCatalogs() {
	data := ""
	for _, menuItems := range sh.xlFile.GetSheetList() {
		data += menuItems+"\n"
	}
	os.WriteFile("shell/menuItems/catalogs.txt", []byte(data), 0777)
}

func (sh *Shell) menu() {
	mainMenu := fyne.NewMainMenu(
		sh.menuItems("Каталоги", "catalogs"),
		sh.menuItems("Действия", "actions"),
		sh.menuItems("Справка", "help"),
	)
	sh.window.SetMainMenu(mainMenu)
}

func getFile(path string) *bufio.Scanner {
	file, _ := os.Open(path)
	defer file.Close()
	return bufio.NewScanner(file)
}

func (sh *Shell) menuItems(nameMenu, fileMenu string) *fyne.Menu {
	menuItems := []*fyne.MenuItem{}

	scanner := getFile("shell/menuItems/" + fileMenu + ".txt")
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
				sh.about()
			}))
		default:
			menuItems = append(menuItems, fyne.NewMenuItem(scanner.Text(), nil))
		}
	}
	return fyne.NewMenu(nameMenu, menuItems...)
}

func (sh *Shell) contacts() {
	sh.showInformation(		"Контакты", "А вы с какой целью, собственно, интересуетесь?")
}

func (sh *Shell) help(){
	vBox := container.NewVBox()
	vBox.Resize(fyne.NewSize(240, 490))
	layout := container.NewWithoutLayout()

	scanner := getFile("shell/help/briefInformation.txt")
	for scanner.Scan() {
		txt := canvas.Text{}
		txt.TextSize = 20
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
