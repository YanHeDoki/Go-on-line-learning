package dao

import (
	"goadmin/dao/mysql"
	"goadmin/utils"
	"time"
)

func CountRegister(day string) (int, error) {
	db := mysql.GetDb()
	sqlstr := "select count(*) from	ucenter_member where DATE(gmt_create)=? "
	count := 0
	err := db.Get(&count, sqlstr, day)
	if err != nil {
		return 0, err
	}
	return count, err
}

func InsertStatistics(day string, sd *StatisticsDaily) error {
	db := mysql.GetDb()
	sqlstr := "delete from statistics_daily where DATE(date_calculated)=? "
	_, err := db.Exec(sqlstr, day)
	if err != nil {
		return err
	}
	date := time.Now()
	datestr := date.Format("2006-01-02 15:04:05")
	sqlstr = "insert into statistics_daily (id,date_calculated,register_num,login_num,video_view_num,course_num,gmt_create,gmt_modified) values(?,?,?,?,?,?,?,?)"
	_, err = db.Exec(sqlstr, utils.GenID(), sd.DateCalculated, sd.RegisterNum, sd.LoginNum, sd.VideoViewNum, sd.CourseNum, datestr, datestr)
	if err != nil {
		return err
	}
	return nil
}

func GetShowData(datetype, start, end string) ([]*StatisticsDaily, error) {
	db := mysql.GetDb()
	sqlstr := "select date_calculated ," + datetype + " from statistics_daily where DATE(date_calculated) between ? and ? "
	stalist := []*StatisticsDaily{}
	err := db.Select(&stalist, sqlstr, start, end)
	if err != nil {
		return nil, err
	}
	return stalist, err
}
