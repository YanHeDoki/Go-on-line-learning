package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"
)

var secret = []byte("雪下的是盐")

//获取随机的字符串
func GetRandstring() string {

	char := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charArr := strings.Split(char, "")
	charlen := len(charArr)
	ran := rand.New(rand.NewSource(time.Now().Unix()))
	var rchar string = ""
	for i := 1; i <= 6; i++ {
		rchar = rchar + charArr[ran.Intn(charlen)]
	}
	return rchar
}

//注册时候返回md5加密和随机的盐
func EncryptPassword(pwd string) string {
	h := md5.New()
	h.Write(secret)
	return hex.EncodeToString(h.Sum([]byte(pwd)))
}

//
//func SCronSta() {
//	for {
//		d := time.Now()
//		if d.Hour() == 1 {
//			s := d.AddDate(0, 0, -1)
//			day := s.Format("2006-01-02 15")
//			count, _ := dao.CountRegister(day)
//			sd := &dao.StatisticsDaily{
//				DateCalculated: day,
//				RegisterNum:    count,
//				LoginNum:       rand.Intn(200),
//				VideoViewNum:   rand.Intn(200),
//				CourseNum:      rand.Intn(200),
//			}
//			dao.InsertStatistics(day, sd)
//		} else {
//			time.Sleep(time.Second * 5)
//		}
//	}
//}
