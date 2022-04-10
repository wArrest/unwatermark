package unwatermark

import (
	"fmt"
	"testing"
)

func TestKs(t *testing.T) {
	ks := Ks{}
	vid := ks.findVid("https://www.kuaishou.com/short-video/3xt44pz7jkjydtq?authorId=3xgpxxvtfac655g&streamSource=find&area=homexxbrilliant")
	if vid != "3xt44pz7jkjydtq" {
		t.Errorf("vid 错误,vid:%s", vid)
	}
	vid = ks.findVid("https://www.kuaishou.com/f/X-1cv459h92GQ1Px_A")
	if vid != "3xt44pz7jkjydtq" {
		t.Errorf("vid 错误,vid:%s", vid)
	}
	fmt.Println(ks.GetRealLink("https://www.kuaishou.com/f/X-1cv459h92GQ1Px_A"))
}
