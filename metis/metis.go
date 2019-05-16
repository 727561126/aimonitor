package almetis

import "fmt"
import "http"

func A_Post(){
    p, err := http.PostForm("https://www.pgyer.com/apiv2/app/view", url.Values{"appKey": {"62c99290f0cb2c567cb153c1fba75d867e"},
        "_api_key": {"584f29517115df2034348b0c06b3dc57"}, "buildKey": {"22d4944d06354c8dcfb16c4285d04e41"}})
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(body))


}
