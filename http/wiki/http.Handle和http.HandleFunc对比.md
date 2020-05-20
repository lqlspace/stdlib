## http.Handle和http.HandleFunc区别：
两者区别在第二个参数  

http.Handle第二个参数是Handler接口，里面包含一个方法，实现该接口的对象均可作为实参：  
```cassandraql
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```  
http.HandleFunc第二个参数是函数变量，函数原型：  
```
handler func(ResponseWriter, *Request)
```
满足条件的函数、对象方法均可。

调用http.HandleFunc时，golang内部会通过适配器转换：
```cassandraql
// HandlerFunc类型是一个适配器，它会将满足func(ResponseWriter, *Reqeust原型)
// 的函数通过HandleFunc(f)形式转换成满足Handler接口的对象，该对象的ServerHTTP方法
// 则直接调用自身
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

  







