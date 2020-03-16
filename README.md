# LzhengHttp

> simple httprequests

```shell

go get github.com/6yi/LzhengHttp

```

Get
```go
  import  "github.com/6yi/LzhengHttp"
  
	req:=LzhengHttp.InitHttpRequest()
	resp,err:=req.Request("GET","http://www.baidu.com")

```
Post

```go
  import  "github.com/6yi/LzhengHttp"
  
	req:=LzhengHttp.InitHttpRequest()
	resp,err:=req.Request("POST","http://www.baidu.com")
```

### if you wanna add httpHead or Params,Cookie ,you can use these func

``` go

req.SetHead("Content-Type","application/x-www-form-urlencoded")
req.SetHead("Connection","Keep")

req.SetParams("password","ddd286eb2672e5c8998ccbc1bfc21fa1")

cook1:=http.Cookie{Name:"XSRF-TOKEN",
		Value:"test",
		Path:"/"}
cook2:=http.Cookie{Name:"test2",
		Value:"test2",
		Path:"/"}		
req.SetCookie(cook1)		
req.SetCookie(cook2)	
```


