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
	rows, _ := sh.xlFile.GetRows(sh.catalogName)
	firstRow := rows[0]
	
	if sh.id <= len(rows) {
		for numCell, cell := range rows[sh.id-1] {
			// if rows[sh.id-1] != nil {
				sh.descriptCard[numCell] = &canvas.Text{}
				sh.descriptCard[numCell].TextSize = 20
				sh.descriptCard[numCell].Text = firstRow[numCell]+" "+cell
				if cell != "" {
					sh.descriptCard = append(sh.descriptCard, sh.descriptCard[numCell])
					vBox.Add(sh.descriptCard[numCell])
				}
			// } else {
				if len(rows[sh.id-1]) == 0 {
					fmt.Println(len(rows[sh.id-1]))
					// vBox.Add(noDiscript)
				}
			// 	fmt.Println(len(rows[sh.id-1]))
			// 	vBox.Add(noDiscript)
			// }
		}
	} else {
		vBox.Add(noDiscript)
	}

	layout := container.NewWithoutLayout(vScroll(416, 450, vBox))
	layout.Move(fyne.NewPos(470, 100))
	return layout
}
