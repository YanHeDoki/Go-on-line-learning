package dao

import (
	"goadmin/dao/mysql"
	"goadmin/utils"
	"time"
)

func GetPageBanner(page, limit int) ([]*Banner, int, error) {
	db := mysql.GetDb()
	totalsql := "select count(1) from crm_banner where is_deleted!=1"
	total := 0
	err := db.Get(&total, totalsql)
	if err != nil {
		return nil, 0, err
	}
	sqlstr := "select id,title,image_url,link_url,sort from crm_banner limit ?,?"
	res := make([]*Banner, 0, 10)
	err = db.Select(&res, sqlstr, (page-1)*limit, limit)
	if err != nil {
		return nil, 0, err
	}
	return res, total, err
}

func GetBannerById(Id string) (*Banner, error) {
	db := mysql.GetDb()
	sqlstr := "select id,title,image_url,link_url,sort from crm_banner where id=?"
	res := &Banner{}
	err := db.Select(&res, sqlstr, Id)
	if err != nil {
		return nil, err
	}
	return res, err
}

func AddBanner(b *Banner) error {
	db := mysql.GetDb()
	sqlstr := "insert into crm_banner(id,title,image_url,link_url,sort,gmt_create,gmt_modified) values(?,?,?,?,?,?,?)"
	date := time.Now()
	datestr := date.Format("2006-01-02 15:04:05")
	_, err := db.Exec(sqlstr, utils.GenID(), b.Title, b.ImageUrl, b.LinkUrl, b.Sort, datestr, datestr)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBanner(b *Banner) error {
	db := mysql.GetDb()
	sqlstr := "update crm_banner set title=?,image_url=?,link_url=?,sort=? where id=?"
	_, err := db.Exec(sqlstr, b.Title, b.ImageUrl, b.LinkUrl, b.Sort, b.Id)
	if err != nil {
		return err
	}
	return nil
}

func DelBanner(Id string) error {
	db := mysql.GetDb()
	sqlstr := "update crm_banner set is_deleted = 1 where id=?"
	_, err := db.Exec(sqlstr, Id)
	if err != nil {
		return err
	}
	return nil
}
