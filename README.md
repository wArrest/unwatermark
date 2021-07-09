# 抖音去水印

## install
~~~
go get -u github.com/wArrest/unwatermark
~~~
## usage
参考测试代码
返回结果是一个key为原始url，value为真实源地址的map
~~~golang
func TestGetResult(t *testing.T){
  u:="https://v.douyin.com/epjE8r9/"
  d := NewDouYin([]string{u})
  res:=d.GetResults()
  fmt.Println(res[u])
}
~~~

> 仅用于学习使用，禁止用于商业用途。
