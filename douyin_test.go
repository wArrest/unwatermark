package unwatermark

import (
  "fmt"
  "testing"
)

func TestGetVideoLink(t *testing.T) {
  d := DouYin{}
  vid,_ := d.findVid("https://v.douyin.com/epYRLk5/")
  if vid != "6977379955709578526" {
    t.Errorf("解析错误:%v", vid)
  }
  link,_:=d.GetRealLink("https://v.douyin.com/epYRLk5/")
  fmt.Println(link)
}
