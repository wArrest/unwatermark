package unwatermark

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
  "path"
)

type GuLiang struct {
  BaseMedia
}

func (g *GuLiang)GetRealLink(txt string) (string, error) {
  //去除换行和空格
  txt = SimpleTxt(txt)
  up, err := url.Parse(txt)
  if err!=nil {
    return "", err
  }
  mid := path.Base(up.Path)
  shortUrl := fmt.Sprintf("https://cc.oceanengine.com/creative_radar_api/v1/material/info?material_id=%s", mid)
  var data struct {
    Data struct {
      Vid string `json:"vid"`
    } `json:"data"`
  }
  client := http.Client{}
  resp, _ := client.Get(shortUrl)
  bb, _ := ioutil.ReadAll(resp.Body)
  json.Unmarshal(bb, &data)
  vUrl := fmt.Sprintf("https://aweme.snssdk.com/aweme/v1/play/?video_id=%s", data.Data.Vid)
  req := NewReq(vUrl, nil)
  resp2, err := client.Do(req)
  if err != nil {
    return "", err
  }
  defer func() {
    resp.Body.Close()
    resp2.Body.Close()
  }()
  return resp2.Request.URL.String(), nil
}
