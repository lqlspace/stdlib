[json with go](http://polyglot.ninja/golang-json/)

## caution
```cassandraql
1、json只序列化导出的字段；

2、如果不想导出字段A被json，A的tag里被置成"-"；  

3、如果要对外提供小写字段，使用golang的tag；

4、对于指针变量来说，Marshal时传递p或&p均可，Unmarshal时如果传递了未初始化的p，
则error，传&p则不会；

```
