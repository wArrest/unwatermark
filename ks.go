package unwatermark

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type KsBody struct {
	OperationName string          `json:"operationName"`
	Query         string          `json:"query"`
	Variables     KsBodyVariables `json:"variables"`
}
type KsBodyVariables struct {
	Page        string `json:"page"`
	PhotoId     string `json:"photoId"`
	WebPageArea string `json:"webPageArea"`
}
type KsRespBody struct {
	Data struct {
		VisionVideoDetail struct {
			Photo struct {
				PhotoUrl string `json:"photoUrl"`
			} `json:"photo"`
		} `json:"visionVideoDetail"`
	} `json:"data"`
}
type Ks struct {
	BaseMedia
	Did string
}

func (x *Ks) GetRealLink(txt string) (string, error) {
	vid := x.findVid(txt)
	return x.findUrl(vid)
}
func (x *Ks) findVid(url1 string) string {
	//长链接直接返回vid
	if strings.Contains(url1, "short-video") {
		urlParse, _ := url.Parse(url1)
		url1 = urlParse.Path
		return path.Base(url1)
	}
	//短链接需要重定向一次
	client := http.Client{}
	resp, _ := client.Get(url1)
	url2 := resp.Request.URL.String()
	urlParse, _ := url.Parse(url2)
	x.Did = urlParse.Query().Get("ztDid")
	return path.Base(urlParse.Path)
}
func (x *Ks) findUrl(vid string) (string, error) {
	jsonS := fmt.Sprintf("{\"operationName\":\"visionVideoDetail\",\"query\":\"query visionVideoDetail($photoId: String, $type: String, $page: String, $webPageArea: String) {  visionVideoDetail(photoId: $photoId, type: $type, page: $page, webPageArea: $webPageArea) {    status    type    author {      id      name      following      headerUrl      __typename    }    photo {      id      duration      caption      likeCount      realLikeCount      coverUrl      photoUrl      liked      timestamp      expTag      llsid      viewCount      videoRatio      stereoType      croppedPhotoUrl      manifest {        mediaType        businessType        version        adaptationSet {          id          duration          representation {            id            defaultSelect            backupUrl            codecs            url            height            width            avgBitrate            maxBitrate            m3u8Slice            qualityType            qualityLabel            frameRate            featureP2sp            hidden            disableAdaptive            __typename          }          __typename        }        __typename      }      __typename    }    tags {      type      name      __typename    }    commentLimit {      canAddComment      __typename    }    llsid    danmakuSwitch    __typename  }}\",\"variables\":{\"page\":\"detail\",\"photoId\":\"%s\",\"webPageArea\":\"homexxbrilliant\"}}", vid)
	req := NewReq("https://www.kuaishou.com/graphql", []byte(jsonS))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Cookie", fmt.Sprintf("did=%s; Path=/; Domain=.kuaishou.com; HttpOnly; Expires=Tue, 02 Aug 2022 14:11:03 GMT;", x.Did))
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	data := KsRespBody{}
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return "", err
	}
	return data.Data.VisionVideoDetail.Photo.PhotoUrl, nil
}
