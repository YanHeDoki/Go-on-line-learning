package utils

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"goadmin/constants"
	"mime/multipart"
)

func ReadExcel(ef *multipart.FileHeader) (map[string][]string, error) {

	f, err := ef.Open()
	if err != nil {
		fmt.Println("file open err:", err)
		return nil, err
	}
	file, err := excelize.OpenReader(f)
	if err != nil {
		fmt.Println("excel open err:", err)
	}
	defer file.Close()
	rs, err := file.Rows("sheet1")
	if err != nil {
		fmt.Println("Excel Read Rows Err:", err)
		return nil, err
	}
	defer rs.Close()
	i := 0
	stumap := make(map[string][]string, 10)
	for rs.Next() {
		//第一行是标题不需要 所以跳过
		if i == 0 {
			i++
			continue
		}

		row, err := rs.Columns()
		if err != nil {
			fmt.Println("Excel Read Columns err:", err)
			return nil, err
		}
		if len(row) < 2 {
			fmt.Println("excel err:", err)
			return nil, errors.New(constants.ExcelErr)
		}
		if _, ok := stumap[row[0]]; ok {
			stumap[row[0]] = append(stumap[row[0]], row[1])
		} else {
			stumap[row[0]] = make([]string, 0, 3)
			stumap[row[0]] = append(stumap[row[0]], row[1])
		}
	}

	return stumap, nil
}
