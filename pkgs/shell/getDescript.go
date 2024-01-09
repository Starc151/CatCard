package shell

import (
	// "fmt"``

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	"github.com/tealeg/xlsx"
)

func (sh *Shell) getDescript(numRow int) *fyne.Container {
	sh.card.descript = append(sh.card.descript, &canvas.Text{})
	vBox := container.NewVBox()
	excelFileName := "./slu/data.xlsx"
    xlFile, _ := xlsx.OpenFile(excelFileName)
	firstRow := xlFile.Sheets[0].Rows[0].Cells
	for numCell := 0; numCell <= len(firstRow)-1; numCell++ {
		if (len(xlFile.Sheets[0].Rows[numRow-1].Cells) != 0) {
			if (xlFile.Sheets[0].Rows[numRow-1].Cells[numCell].String() != "") {
				sh.card.descript[numCell] = &canvas.Text{}
				sh.card.descript[numCell].Text = firstRow[numCell].String()+" "+xlFile.Sheets[0].Rows[numRow-1].Cells[numCell].String()
				sh.card.descript[numCell].TextSize = 20
				sh.card.descript = append(sh.card.descript, sh.card.descript[numCell])
				vBox.Add(sh.card.descript[numCell])
			}
		}
	}
	// fmt.Println(sh.card.id)
	// fmt.Println(len(xlFile.Sheets[0].Rows[numRow-1].Cells))
	vScroll := container.NewVScroll(vBox)
	vScroll.Resize(fyne.NewSize(416, 450))
	layout := container.NewWithoutLayout(vScroll)
	layout.Move(fyne.NewPos(470, 100))
	return layout
}
