## io.Reader和io.Writer和io.Closer 定义
io.Reader和io.Writer是两个接口，分别为调用者提供数据和输出数据，io.Closer接口则用来关闭对象，定义如下：
```cassandraql
// 此处p用来接收Reader中的数据，类型为字节数组
type Reader interface {
	Read(p []byte) (n int, err error)
}

// 此处p是写入Writer的数据源，类型为字节数组
type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}
```

## bytes.Reader定义 (注：没有bytes.Writer定义)
bytes.Reader是一个结构体，实现了io.Reader（只读）、io.Seek等，内部有字节数组类型的缓存，以及索引
```cassandraql
type Reader struct {
	s        []byte
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}
```

## bytes.Buffer定义
bytes.Buffer是一个结构体，实现了io.Reader和io.Writer方法
```cassandraql
type Buffer struct {
	buf      []byte // contents are the bytes buf[off : len(buf)]
	off      int    // read at &buf[off], write at &buf[len(buf)]
	lastRead readOp // last read operation, so that Unread* can work correctly.
}
```

## 字符串类型实现io.Reader/io.Writer接口
字符串有个专门的struct存放要Read的string:
```cassandraql
type Reader struct {
	s        string
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}
```
其中s为数据源，其初始化使用strings.NewReader函数：
```cassandraql
func NewReader(s string) *Reader { return &Reader{s, 0, -1} }
```
Read方法实现如下：
```cassandraql
func (r *Reader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}
```

## bytes实现io.Reader/io.Writer接口
bytes及字节数组,有个专门的结构体来存放要Read的bytes:
```cassandraql
type Reader struct {
	s        []byte
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}
```
其中s为数据源，其初始化使用bytes.NewReader函数：
```cassandraql
func NewReader(b []byte) *Reader { return &Reader{b, 0, -1} }
```
Read方法实现如下：
```cassandraql
func (r *Reader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}
```

