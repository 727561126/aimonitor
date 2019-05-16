package main

import "fmt"
import "aimonitor/utils"
import "aimonitor/aliyun"
//import "github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
import "aimonitor/postgres"

func alitopg(){
    fmt.Println(aliyunrds.AliYun_sj().PerformanceKeys.PerformanceKey)
    db := alpostgres.SqlOpen()
    for _,v := range aliyunrds.AliYun_sj().PerformanceKeys.PerformanceKey {
        for _, b := range v.Values.PerformanceValue {
           fmt.Println(b.Value)
           fmt.Println(b.Date)
           alpostgres.SqlInsert(db, v.Key, b.Value, utils.UtcToBj(b.Date))
        }
    }


}

func main(){
//  ue := new(utils.UtilsErr)
//  fmt.Println(ue.Error())
//  fmt.Println(utils.UtcToBj("2019-05-06T20:48:01Z"))
    
//    fmt.Println(utils.UtcToBj("2019-05-08T00:00:00Z"))
//    fmt.Println(utils.BjToUtc("2019-05-08 00:00:00"))
    ///aliyun_re := &rds.DescribeDBInstancePerformanceResponse{}
    //err := json.Unmarshal([]byte(aliyunrds.AliYun_sj()), &aliyun_re)
    db := alpostgres.SqlOpen()
    fmt.Println(alpostgres.SqlSelect(db))
    
}
