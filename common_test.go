package unwatermark

import (
  "fmt"
  "mvdan.cc/xurls/v2"
  "testing"
)

func TestSimpleCode(t *testing.T){
  url:="4.12 usr:/ %%深夜看球的正确姿势 %%抖音美食创作人 和林先生一起看球的正确打开方式～  https://v.douyin.com/ecEnu38/ 复制此链接，打开Dou音搜索，直接观看视频！"
  xurlsStrict := xurls.Strict()
  output := xurlsStrict.FindAllString(url, -1)
  fmt.Println(output)
}
