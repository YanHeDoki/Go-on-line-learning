package utils

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"testing"
)

func TestExcleCols(t *testing.T) {
	f, err := excelize.OpenFile("C:/Users/YanHe/Desktop/WD/test.xlsx")
	if err != nil {
		fmt.Println("open file err:", err)
	}
	defer f.Close()
	rs, _ := f.Rows("sheet1")
	i := 0
	stumap := make(map[string][]string, 10)
	for rs.Next() {
		if i == 0 {
			i++
			continue
		}
		row, err := rs.Columns()
		if err != nil {
			fmt.Println("Excel Read Columns err:", err)
		}
		if len(row) < 2 {
			fmt.Println("excel err:", err)
		}
		if _, ok := stumap[row[0]]; ok {
			stumap[row[0]] = append(stumap[row[0]], row[1])
		} else {
			stumap[row[0]] = make([]string, 0, 3)
			stumap[row[0]] = append(stumap[row[0]], row[1])
		}
	}
	fmt.Println(stumap)
}
