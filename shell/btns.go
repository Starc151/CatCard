package shell

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/xuri/excelize/v2"
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

func (sh *Shell) editCard() {
	gridColumns := container.NewGridWithColumns(1)
	box := container.NewWithoutLayout()
	cont := container.NewWithoutLayout(sh.searchBox())
	vBox := container.NewVBox()
	rows, _ := sh.xlFile.GetRows(sh.catalogName)
	valueNew := map[int]string{}
	entryList := []*widget.Entry{}
	for numCell := 0; numCell < len(rows[0]); numCell++ {
		descript := canvas.NewText(rows[0][numCell], nil)
		descript.Move(fyne.NewPos(20, 0))
		entry := widget.NewEntry()
		if numCell < len(rows[sh.id-1]) {
			entry.Text = rows[sh.id-1][numCell]
		}
		entry.Resize(fyne.NewSize(400, 40))
		entry.Move(fyne.NewPos(400, -30))
		entryList = append(entryList, entry)
		layoutEntry := container.NewWithoutLayout(entryList[numCell])
		layoutDiscript := container.NewWithoutLayout(descript)
		vBox.Add(layoutDiscript)
		vBox.Add(layoutEntry)
	}
	
	gridColumns.Add(box)
	cansel := sh.button("Отмена", 437, 40, 10, 500, sh.showCard)
	accept := 	sh.button("Применить", 416, 40, 470, 500, func() {
		for v, k := range entryList {
			valueNew[v] = k.Text
		}
		sh.editCardAccept(valueNew)
	})
	cont.Add(vScroll(890, 400, 10, 70, vBox))
	cont.Add(cansel)
	cont.Add(accept)
	sh.setContent(cont)
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
	for numPic := 2; numPic <= sh.lenCatalog+1; numPic++ {
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

	cont.Add(vScroll(890, 460, 10, 80, gridColumns))
	sh.setContent(cont)
}

func (sh *Shell) search(request string) {
	result := map[string][]int{}
	for _, listNane := range sh.xlFile.GetSheetList() {
		rows, _ := sh.xlFile.GetRows(listNane)
		for numRow, row := range rows {
			for _, cell := range row {
				lowRequest := strings.ToLower(request)
				lowerCell := strings.ToLower(cell)
				if strings.Contains(lowerCell, lowRequest) && numRow != 0 {
					result[listNane] = append(result[listNane], numRow+1)
					break
				}
			}
		}
	}

	sh.lenCatalog = len(result)
	catalogNameText := canvas.NewText("Результаты поиска:", nil)
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
	cont.Add(vScroll(890, 460, 10, 80, gridColumns))
	if len(result) == 0 || request == "" {
		catalogNameText.Text = "Ничего не найдено"
		cont = container.NewWithoutLayout(sh.searchBox(), catalogName)
	}
	sh.setContent(cont)
}

func (sh *Shell) editCardAccept(valueNew map[int]string) {
	for k, v := range valueNew {
		cell, _ := excelize.CoordinatesToCellName(k+1, sh.id)
		sh.xlFile.SetCellValue(sh.catalogName, cell, v)
	}
	sh.xlFile.Save()
	sh.showCard()	
}