package unwatermark

import (
  "fmt"
  "testing"
)

func TestKs(t *testing.T) {
  ks:=NewKs([]string{"https://www.kuaishou.com/f/X-58XHD4OQYLW1Fp_A"})
  fmt.Println(ks.GetResults())
}

func TestKs_findVid(t *testing.T) {
  ks:=NewKs([]string{""})
  vid:=ks.findVid("https://www.kuaishou.com/short-video/3xt44pz7jkjydtq?authorId=3xgpxxvtfac655g&streamSource=find&area=homexxbrilliant")
  if vid != "3xt44pz7jkjydtq"{
    t.Errorf("vid 错误,vid:%s",vid)
  }
  vid=ks.findVid("https://www.kuaishou.com/f/X-1cv459h92GQ1Px_A")
  if vid != "3xt44pz7jkjydtq"{
    t.Errorf("vid 错误,vid:%s",vid)
  }
}
