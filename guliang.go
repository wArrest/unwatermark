package unwatermark

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
  "path"
  "sync"
)

type Guliang struct {
	urlMap map[string]string
}

func NewGuliang(urls []string) *Guliang {
	urlMap := make(map[string]string)
	for _, url := range urls {
		urlMap[url] = ""
	}
	return &Guliang{
		urlMap: urlMap,
	}
}
func (g *Guliang) GetResults() map[string]string {
	wg := sync.WaitGroup{}
	//最多10个一起并发处理
	c := make(chan int, 10)
	for url1, _ := range g.urlMap {
		wg.Add(1)
		c <- 1
		go func(url1 string) {
			rUrl, err := g.findUrl(url1)
			if err != nil {
				fmt.Println(err)
				wg.Done()
				<-c
				return
			}
			g.urlMap[url1] = rUrl
			wg.Done()
			<-c
		}(url1)
	}
	wg.Wait()
	return g.urlMap
	return nil
}

func (g *Guliang) findUrl(url1 string) (string, error) {
	//去除换行和空格
	url1 = SimpleCode(url1)
	up, err := url.Parse(url1)
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
