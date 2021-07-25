# 常见视频媒体，去水印小工具

## feature
- [x] 抖音
- [ ] 今日头条
- [ ] 西瓜视频
- [ ] 快手
- [ ] 小红书
- [ ] bilibili
- [ ] 腾讯视频
- [ ] 巨量引擎

## install
~~~
go get -u github.com/wArrest/unwatermark
~~~
## usage
可以参考测试代码，返回结果是一个key为原始url，value为真实源地址的map
~~~golang
func TestGetResult(t *testing.T){
  u:="https://v.douyin.com/epjE8r9/"
  d := NewDouYin([]string{u})
  res:=d.GetResults()
  fmt.Println(res[u])
}
~~~
## remark
> 该项目仅用于学习使用，禁止用于商业用途❗️❗️❗️
