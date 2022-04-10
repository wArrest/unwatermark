package unwatermark

import (
	"fmt"
	"testing"
)

func TestGuliang_GetResults(t *testing.T) {
	gl := GuLiang{}
	fmt.Println(gl.GetRealLink("https://cc.oceanengine.com/inspiration/creative-radar/detail/6990328531382222861?appCode=4&period=3&listType=1"))
}
