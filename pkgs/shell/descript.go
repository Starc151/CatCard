package shell

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/tealeg/xlsx"
)

func (sh *Shell) descript() *fyne.Container {
	sh.descriptCard = append(sh.descriptCard, &canvas.Text{})
	vBox := container.NewVBox()
    xlFile, _ := xlsx.OpenFile("./slu/data.xlsx")
	firstRow := xlFile.Sheets[0].Rows[0].Cells
	currentRow := xlFile.Sheets[0].Rows[sh.id-1].Cells
	for numCell := 0; numCell <= len(firstRow)-1; numCell++ {
		if len(currentRow) != 0{
			if currentRow[numCell].String() != ""{
				sh.descriptCard[numCell] = &canvas.Text{}
				sh.descriptCard[numCell].Text = firstRow[numCell].String()+" "+currentRow[numCell].String()
				sh.descriptCard[numCell].TextSize = 20
				sh.descriptCard = append(sh.descriptCard, sh.descriptCard[numCell])
				vBox.Add(sh.descriptCard[numCell])
			}
		}
	}
	vScroll := container.NewVScroll(vBox)
	vScroll.Resize(fyne.NewSize(416, 450))
	layout := container.NewWithoutLayout(vScroll)
	layout.Move(fyne.NewPos(470, 100))
	return layout
}
