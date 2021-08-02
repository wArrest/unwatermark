package unwatermark

import (
  "fmt"
  "testing"
)

func TestGuliang_GetResults(t *testing.T) {
  gl:=NewGuliang([]string{"https://cc.oceanengine.com/inspiration/creative-radar/detail/6990328531382222861?appCode=4&period=3&listType=1"})
  fmt.Println(gl.GetResults())
}

func TestGuliang_findUrl(t *testing.T) {
  //gl:=NewGuliang([]string{"https://cc.oceanengine.com/inspiration/creative-radar/detail/6990328531382222861?appCode=4&period=3&listType=1"})
  //gl.findUrl()
}
