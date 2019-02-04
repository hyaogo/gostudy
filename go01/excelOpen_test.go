package go01

import (
	"fmt"
	"github.com/Luxurioust/excelize"
	"os/exec"
	"testing"
)

//浏览器打开Excel中的地址
func Test(t *testing.T) {
	xlsx, err := excelize.OpenFile("D:/x.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	rows := xlsx.GetRows("Sheet1")
	i := 0
	for _, row := range rows {
		//控制条数
		if i > 77 {
			break
		}
		for _, cell := range row {
			fmt.Println(cell)
			if cell != "" {
				//cell 地址
				cmd := exec.Command("explorer", cell)
				err := cmd.Start()
				if err != nil {
					fmt.Println(err.Error())
				}
			}

		}
		i++

	}
}
