package utils

import "time"

//UTC时间转换为北京时间
func UtcToBj(utc string) string{
  BJ, _ := time.LoadLocation("Asia/Shanghai") 
  bj, _ := time.Parse("2006-01-02T15:04:05Z", utc)
  
  return bj.In(BJ).Format("2006-01-02 15:04:05")
}

//北京时间转换为UTC时间
func BjToUtc(bj string) string{
  BJ, _ := time.LoadLocation("Asia/Shanghai")
  UTC, _ := time.LoadLocation("")
  uc, _ := time.ParseInLocation("2006-01-02 15:04:05", bj, BJ) 
  
  return uc.In(UTC).Format("2006-01-02T15:04:05Z")
}
