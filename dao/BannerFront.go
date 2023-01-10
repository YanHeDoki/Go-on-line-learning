package dao

import "goadmin/dao/mysql"

func GetAllBanner() ([]*Banner, error) {
	db := mysql.GetDb()
	sqlstr := "select id,title,image_url,link_url,sort from crm_banner ORDER BY id desc limit 2  "
	res := make([]*Banner, 0, 10)
	err := db.Select(&res, sqlstr)
	if err != nil {
		return nil, err
	}
	return res, nil
}
