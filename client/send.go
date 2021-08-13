package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Send(to,method string,m interface{}) []byte {
	s,_:=json.Marshal(m)
	r,err:=http.Post("http://"+to+"/"+method,"application/json",strings.NewReader(string(s)))
	if err!=nil{
		log.Println(err)
	}
	defer r.Body.Close()
	content,err:=ioutil.ReadAll(r.Body)
	if err!=nil {
		log.Println(err)
	}
	return content
}
