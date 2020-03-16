package LzhengHttp

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)
type LzhengHttp struct {
	head map[string]string
	resp *http.Response
	form url.Values
	cookie []http.Cookie
}
func (l *LzhengHttp)SetHead(key,value string)  {
	l.head[key]=value
}
func (l *LzhengHttp)SetCookie(cookie http.Cookie)  {
	l.cookie = append(l.cookie,cookie)
}

func (l *LzhengHttp)SetParms(key,value string){
	if l.form==nil{
		l.form=make(url.Values)
	}
	l.form.Set(key,value)
}

func InitHttpRequest()*LzhengHttp {
	return &LzhengHttp{head:make(map[string]string),resp:nil,form:nil,cookie:make([]http.Cookie,0)}
}

func (l *LzhengHttp) Request(method,Url string,) (*http.Response,error){
	client := &http.Client{}
	req, err := http.NewRequest(method, Url, nil)
	if err != nil {
		return nil,err
	}
	if l.head!=nil{
		for k,v:=range l.head{
			req.Header.Set(k,v)
		}
	}
	if len(l.cookie)!=0{
		for _,v:=range l.cookie{
			req.AddCookie(&v)
		}
	}
	if l.form!=nil{
		req.Body = ioutil.NopCloser(strings.NewReader(l.form.Encode()))
	}
	resp, err := client.Do(req)
	l.resp=resp

	return resp,err
}
