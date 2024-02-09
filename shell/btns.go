package shell

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func (sh *Shell) nextCard() {
	sh.id++
	sh.showCard()
}

func (sh *Shell) preCard() {
	sh.id--
	sh.showCard()
}

func (sh *Shell) reverse() {
	path := fmt.Sprintf("pics/%s/back/%d.jpg", sh.catalogName, sh.id)
	reversePic := canvas.NewImageFromFile(checkPic(path))
	reversePic.SetMinSize(fyne.NewSize(437, 310))
	sh.showCustom("", "ok", reversePic)
}

func (sh *Shell) showAllCatalog() {
	sh.getlenCatalog()
	catalogNameText := canvas.NewText(sh.catalogName, nil)


	catalogName := container.New(layout.NewCenterLayout(), catalogNameText)
	catalogName.Resize(fyne.NewSize(890, 70))
	catalogName.Move(fyne.NewPos(10, 20))
	cont := container.NewWithoutLayout(sh.searchBox(), catalogName)
	if sh.lenCatalog == 0 {
		cont = sh.emptyCont()
	}
	gridColumns := container.NewGridWithColumns(4)
	for numPic := 2; numPic <= sh.lenCatalog+1; numPic++{
		pic := sh.pic(fmt.Sprintf("pics/%s/%d.jpg", sh.catalogName, numPic))
		id := numPic
		pic.SetMinSize(fyne.NewSize(195, 109))
		btn := sh.button(
			"", 195, 109,
			0, 0, func() {
				sh.id = id
				sh.showCard()
			},
		)
		contCard := container.New(
			layout.NewStackLayout(),
			btn,
			pic,
		)
		contCard.Resize(fyne.NewSize(195, 109))
		gridColumns.Add(container.NewWithoutLayout(contCard))
	}
	
	vScroll := container.NewVScroll(gridColumns)
	vScroll.Resize(fyne.NewSize(890, 460))
	vScroll.Move(fyne.NewPos(10, 80))
	cont.Add(vScroll)
	sh.setContent(cont)
}

func (sh *Shell) search(request string) {
	result := map[string][]int{}
	for _, listNane := range sh.xlFile.GetSheetList() {
		rows, _ :=  sh.xlFile.GetRows(listNane)
		for numRow, row := range rows {
			for _, cell := range row {
				lowRequest := strings.ToLower(request)
				lowerCell := strings.ToLower(cell)
				if strings.Contains(lowerCell, lowRequest) && numRow != 0 {
					result[listNane] = append(result[listNane], numRow + 1)
					break
				}
			}
		}
	}
	
	sh.lenCatalog = len(result)
	catalogNameText := canvas.NewText("Результаты поиска:", nil)
	if len(result) == 0 {
		catalogNameText.Text = "Ничего не найдено"
	}
	catalogName := container.New(layout.NewCenterLayout(), catalogNameText)
	catalogName.Resize(fyne.NewSize(890, 70))
	catalogName.Move(fyne.NewPos(10, 20))
	cont := container.NewWithoutLayout(sh.searchBox(), catalogName)
	
	gridColumns := container.NewGridWithColumns(4)
	for listName, ids := range result {
		catalogName := listName
		for _, numStr := range ids {
			pic := sh.pic(fmt.Sprintf("pics/%s/%d.jpg", listName, numStr))
			id := numStr
			pic.SetMinSize(fyne.NewSize(195, 109))
			btn := sh.button(
				"", 195, 109,
				0, 0, func() {
					sh.catalogName = catalogName
					sh.id = id
					sh.showCard()
				},
			)
			contCard := container.New(
				layout.NewStackLayout(),
				btn,
				pic,
			)
			contCard.Resize(fyne.NewSize(195, 109))
			gridColumns.Add(container.NewWithoutLayout(contCard))
		}
	}
	
	vScroll := container.NewVScroll(gridColumns)
	vScroll.Resize(fyne.NewSize(890, 460))
	vScroll.Move(fyne.NewPos(10, 80))
	cont.Add(vScroll)
	sh.setContent(cont)
}