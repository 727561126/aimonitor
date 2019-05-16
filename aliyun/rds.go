package aliyunrds

import "github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
import "fmt"
//返回阿里云数据
func AliYun_sj() *rds.DescribeDBInstancePerformanceResponse{
  client, err := rds.NewClientWithAccessKey("cn-beijing", "LTAIrxbNxa14iYmH", "HUhJehjqcSSOKR46JMoPigCehaAcsG")

  request := rds.CreateDescribeDBInstancePerformanceRequest()

  request.EndTime = "2019-05-05T16:00Z"
  request.StartTime = "2019-05-01T08:00Z"
  request.Key = "MySQL_NetworkTraffic"
  request.DBInstanceId = "rdst3j60ei1o0m8jjoi5"

  response, err := client.DescribeDBInstancePerformance(request)
  if err != nil {
    fmt.Print(err.Error())
  }
  return response
}
