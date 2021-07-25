package unwatermark

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
  "net/url"
  "path"
  "strings"
  "sync"
)

type JsonData struct {
  ItemList []Item `json:"item_list"`
}

type Item struct {
  Video     Video  `json:"video"`
  ForwardId string `json:"forward_id"`
}

type Video struct {
  Vid string `json:"vid"`
}

type DouYin struct {
  urlMap map[string]string
}

func NewDouYin(urls []string) *DouYin {
  urlMap := make(map[string]string)
  for _, url := range urls {
    urlMap[url] = ""
  }
  return &DouYin{
    urlMap: urlMap,
  }
}
func (d *DouYin) GetResults() map[string]string {
  wg := sync.WaitGroup{}
  //最多10个一起并发处理
  c := make(chan int, 10)
  for url1, _ := range d.urlMap {
    wg.Add(1)
    c <- 1
    go func(url1 string) {
      vid, err := d.findVid(url1)
      if err != nil {
        wg.Done()
        <-c
        return
      }
      d.urlMap[url1] = d.findVideoLink(vid)
      wg.Done()
      <-c
    }(url1)
  }
  wg.Wait()
  return d.urlMap
}

func (d *DouYin) findVid(url1 string) (string, error) {
  url1 = SimpleCode(url1)
  //如果是网页链接，直接取出来
  if strings.Contains(url1, "/video") {
    vUrl, err := url.Parse(url1)
    if err != nil {
      return "", err
    }
    return path.Base(vUrl.Path), nil
  }
  //如果是短链接，需要去请求一次
  client := http.Client{}
  resp, err := client.Do(NewReq(url1))
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()

  vidstr, _ := path.Split(resp.Request.URL.String())
  vid := path.Base(vidstr)
  return vid, nil
}
func (d *DouYin) findVideoLink(vid string) string {
  // 通过这个接口获取视频信息，其中包括带有水印的链接
  url := "https://www.iesdouyin.com/web/api/v2/aweme/iteminfo/?item_ids=" + vid + ""
  client := http.Client{}
  resp, _ := client.Do(NewReq(url))
  respBody, _ := ioutil.ReadAll(resp.Body)
  data := JsonData{}
  json.Unmarshal(respBody, &data)
  if len(data.ItemList) == 0 {
    return ""
  }
  videoLink := "https://aweme.snssdk.com/aweme/v1/play/?video_id=" + data.ItemList[0].Video.Vid + "&ratio=720p&line=0"
  resp, err := client.Do(NewReq(videoLink))
  if err!=nil {
    return ""
  }
  return resp.Request.URL.String()
}

