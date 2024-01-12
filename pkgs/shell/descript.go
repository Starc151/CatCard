package shell

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/tealeg/xlsx"
)

func (sh *Shell) descript( d int) *fyne.Container {
	sh.descriptCard = append(sh.descriptCard, &canvas.Text{})
	vBox := container.NewVBox()
	excelFileName := "./slu/data.xlsx"
    xlFile, _ := xlsx.OpenFile(excelFileName)
	firstRow := xlFile.Sheets[0].Rows[0].Cells
	for numCell := 0; numCell <= len(firstRow)-1; numCell++ {
		descript := canvas.Text{}
		descript.Text = firstRow[numCell].String()+" "+xlFile.Sheets[0].Rows[d-1].Cells[numCell].String()
		descript.TextSize = 20
		sh.descriptCard = append(sh.descriptCard, &descript)
		vBox.Add(sh.descriptCard[numCell])
	}
	fmt.Println(sh.id)
	vScroll := container.NewVScroll(vBox)
	vScroll.Resize(fyne.NewSize(416, 450))
	layout := container.NewWithoutLayout(vScroll)
	layout.Move(fyne.NewPos(470, 100))
	return layout
}