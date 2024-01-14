package shell

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func (sh *Shell) descript() *fyne.Container {
	sh.descriptCard = append(sh.descriptCard, &canvas.Text{})
	vBox := container.NewVBox()
	noDiscript := canvas.NewText("Нет данных", nil)
	noDiscript.TextSize = 20
	firstRow := sh.xlFile.Sheets[0].Rows[0].Cells
	currentRow := sh.xlFile.Sheets[0].Rows[sh.id-1].Cells
	fmt.Println(len(firstRow))
	fmt.Println(len(currentRow))
	if sh.id <= sh.numRowsXls {
	// 	for numCell := 0; numCell <= len(firstRow)-1; numCell++ {
	// 		if len(currentRow) != 0{
	// 			fmt.Println(len(currentRow))
	// 			if currentRow[numCell].String() != ""{
	// 				// fmt.Println(len(firstRow)-1)
	// 				// fmt.Println(numCell)
	// 				sh.descriptCard[numCell] = &canvas.Text{}
	// 				sh.descriptCard[numCell].TextSize = 20
					sh.descriptCard[numCell].Text = firstRow[numCell].String()+" "+currentRow[numCell].String()
					sh.descriptCard = append(sh.descriptCard, sh.descriptCard[numCell])
					vBox.Add(sh.descriptCard[numCell])
	// 			}
	// 		}
	// 	}
	} else {
		vBox.Add(noDiscript)
	}
	layout := container.NewWithoutLayout(vScroll(416, 450, vBox))
	layout.Move(fyne.NewPos(470, 100))
	return layout
}
