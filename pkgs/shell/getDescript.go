package shell

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	"github.com/tealeg/xlsx"
)

func (sh *Shell) getDescript(numRow int) *fyne.Container {
	vBox := container.NewVBox()
	excelFileName := "./slu/data.xlsx"
    xlFile, _ := xlsx.OpenFile(excelFileName)
	firstRow := xlFile.Sheets[0].Rows[0].Cells
	for numCell := 0; numCell <= len(firstRow)-1; numCell++ {
		descript := canvas.Text{}
		descript.Text = firstRow[numCell].String()+" "+xlFile.Sheets[0].Rows[numRow].Cells[numCell].String()
		descript.TextSize = 20
		vBox.Add(&descript)
	}
	layout := container.NewWithoutLayout(vBox)
	layout.Move(fyne.NewPos(470, 100))
	return layout
}
