package unwatermark

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type DouYinJsonData struct {
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
	BaseMedia
}

func (d DouYin) GetRealLink(txt string) (string, error) {
	vid, err := d.findVid(txt)
	if err != nil {
		return "", err
	}
	// 通过这个接口获取视频信息，其中包括带有水印的链接
	url1 := "https://www.iesdouyin.com/web/api/v2/aweme/iteminfo/?item_ids=" + vid + ""
	client := http.Client{}
	resp, _ := client.Do(NewReq(url1, nil))
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	data := DouYinJsonData{}
	json.Unmarshal(respBody, &data)
	if len(data.ItemList) == 0 {
		return "", errors.New("no content")
	}
	videoLink := "https://aweme.snssdk.com/aweme/v1/play/?video_id=" + data.ItemList[0].Video.Vid + "&ratio=720p&line=0"
	resp, err = client.Do(NewReq(videoLink, nil))
	if err != nil {
		return "", err
	}
	return resp.Request.URL.String(), nil
}

func (d DouYin) findVid(txt string) (string, error) {
  txt = SimpleTxt(txt)
  //如果是网页链接，直接取出来
  if strings.Contains(txt, "/video") {
    vUrl, err := url.Parse(txt)
    if err != nil {
      return "", err
    }
    return path.Base(vUrl.Path), nil
  }
  //如果是短链接，需要去请求一次
  client := http.Client{}
  resp, err := client.Do(NewReq(txt, nil))
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()
  vidstr, _ := path.Split(resp.Request.URL.String())
  vid := path.Base(vidstr)
  return vid, nil
}
