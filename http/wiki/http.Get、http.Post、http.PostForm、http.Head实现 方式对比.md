## http.GET、http.Post、http.PostForm、http.Head实现方式对比
这四个函数均是对DefaultClient对象接口的封装：
```cassandraql
// http.Get，只需要指定URL即可
func Get(url string) (resp *Response, err error) {
	return DefaultClient.Get(url)
}

// http.Post，入参有三个，其中第二个参数常见如（application/json;charset=utf8），第三入参为io.Reader接口
func Post(url, contentType string, body io.Reader) (resp *Response, err error) {
	return DefaultClient.Post(url, contentType, body)
}

// http.PostForm，入参有两个，因为contentType有默认值，data也会进行转换
func PostForm(url string, data url.Values) (resp *Response, err error) {
	return DefaultClient.PostForm(url, data)
}

// http.Head，一个URL入参，返回头部信息
func Head(url string) (resp *Response, err error) {
	return DefaultClient.Head(url)
}
```
DefaultClient对象是一个默认的全局Client对象：
```cassandraql
var DefaultClient = &Client{}
```
DefaultClient对象四个方法的实现如下：
```cassandraql
func (c *Client) Get(url string) (resp *Response, err error) {
	req, err := NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *Client) Post(url, contentType string, body io.Reader) (resp *Response, err error) {
	req, err := NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return c.Do(req)
}

func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error) {
	return c.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}

func (c *Client) Head(url string) (resp *Response, err error) {
	req, err := NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}
```
可见，Client对象工作方式：  
- 1.调用NewRequest创建一个 http.Request;  
- 2.调用client.Do函数从事请求操作；
