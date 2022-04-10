package unwatermark

import (
	"strings"
	"sync"
)

type Media interface {
	//获取批量处理的合集
	GetResult(data []string) []LinkData
	//获取一条真实链接
	GetRealLink(txt string) (string, error)
}
type LinkData struct {
	Index int
	Link  string
}
type BaseMedia struct{}

func (base BaseMedia) GetResult(data []string) []LinkData {
	var (
		result             []LinkData
		linkChan           chan LinkData
		handleWg, resultWg sync.WaitGroup
	)
	//开启一个协程接收处理结果
	go func() {
		resultWg.Add(1)
		for linkData := range linkChan {
			result = append(result, linkData)
		}
		resultWg.Done()
	}()
	//最多10个一起并发处理
	rateLimit := make(chan struct{}, 10)
	for index, txt := range data {
		handleWg.Add(1)
		rateLimit <- struct{}{}
		go func(index int, txt string) {
			defer func() {
				handleWg.Done()
				<-rateLimit
			}()
			link, err := base.GetRealLink(txt)
			if err != nil {
				return
			}
			//将结果塞入通道
			linkChan <- LinkData{
				Index: index,
				Link:  link,
			}
		}(index, txt)
	}
	handleWg.Wait()
	return result
}
func (base BaseMedia) GetRealLink(txt string) (string, error) { return "", nil }

func GetMedia(rawText string) Media {
	if strings.Contains(rawText, "kuaishou") {
		return &Ks{}
	} else if strings.Contains(rawText, "douyin") {
		return &DouYin{}
	} else if strings.Contains(rawText, "cc.oceanengine.com") {
		return &GuLiang{}
	}
	return nil
}
