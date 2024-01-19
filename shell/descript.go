package shell

import (
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

	vBox.Add(noDiscript)
	
	if (sh.id <= len(rows)) && (len(rows[sh.id-1]) > 0) {
		vBox.RemoveAll()
		for numCell, cell := range rows[sh.id-1] {
			sh.descriptCard[numCell] = &canvas.Text{}
			sh.descriptCard[numCell].TextSize = 20
			sh.descriptCard[numCell].Text = rows[0][numCell]+" "+cell
			if cell != "" {
				sh.descriptCard = append(sh.descriptCard, sh.descriptCard[numCell])
				vBox.Add(sh.descriptCard[numCell])
			}
		}
	}

	layout := container.NewWithoutLayout(vScroll(416, 450, vBox))
	layout.Move(fyne.NewPos(470, 100))
	return layout
}
