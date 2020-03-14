package main

import (
	"encoding/json"
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/fanux/robot/issue"
	"github.com/fanux/robot/processor/pay"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// This example shows the minimal code needed to get a restful.WebService working.
//
// GET http://localhost:8080/hello

func main() {
	issue.Regist("/pay", &pay.Pay{})
	ws := new(restful.WebService)
	ws.Route(ws.GET("/pay").To(payHandle))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func payHandle(req *restful.Request, resp *restful.Response) {
	b,err := ioutil.ReadAll(req.Request.Body)
	event := &issue.IssueCommentEvent{}
	value,err := url.ParseQuery(string(b))
	eventstr := []byte(value.Get("payload"))
	if len(eventstr) == 0 {
		return
	}
	err = json.Unmarshal(eventstr,event)
	if err != nil {
		fmt.Printf("decode event error : %s",err)
		return
	}
	fmt.Printf("repo name is : %s body: %s/n",*event.Repo.FullName, *event.Comment.Body)
	config := issue.NewConfig("", "")
	err = issue.Process(config, *event)
	if err != nil {
		fmt.Printf("promote error %s",err)
	}
	io.WriteString(resp, "world")
}