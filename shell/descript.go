package shell

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func (sh *Shell) descript() *fyne.Container {
	vBox := container.NewVBox()
	noDiscript := canvas.NewText("Нет данных", nil)
	rows, _ := sh.xlFile.GetRows(sh.catalogName)

	vBox.Add(noDiscript)
	
	if (sh.id <= len(rows)) && (len(rows[sh.id-1]) > 0){
		vBox.RemoveAll()
		for numCell := 0; numCell <  min(len(rows[0]), len(rows[sh.id-1])) ; numCell++ {
			descript := canvas.NewText(rows[0][numCell]+" "+rows[sh.id-1][numCell], nil)
			if rows[sh.id-1][numCell] != "" {
				vBox.Add(descript)
			}
		}
	}

	return vBox
}