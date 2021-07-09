package unwatermark

import (
  "fmt"
  "testing"
)

func TestDouYin_findVid(t *testing.T) {
  //https://www.douyin.com/video/6977379955709578526?previous_page=app_code_link
  d := NewDouYin([]string{})

  vid,_ := d.findVid("https://v.douyin.com/epYRLk5/")
  if vid != "6977379955709578526" {
    t.Errorf("解析错误:%v", vid)
  }
  vid,_ = d.findVid("https://www.douyin.com/video/6977379955709578526?previous_page=app_code_link")
  if vid != "6977379955709578526" {
    t.Errorf("解析错误:%v", vid)
  }
}

func TestDouYin_findVideoLink(t *testing.T) {
  d := NewDouYin([]string{})
  d.findVideoLink("6977379955709578526")
}

func TestGetResult(t *testing.T){
  u:="https://v.douyin.com/epYRLk5/"
  d := NewDouYin([]string{u})
  res:=d.GetResults()
  fmt.Println(res[u])
}
