package main

import (
	"encoding/json"
	"fmt"
	"github.com/fanux/robot/issue"
	"github.com/fanux/robot/processor/pay"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func handler(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	event := &issue.IssueCommentEvent{}
	value, err := url.ParseQuery(string(b))
	eventstr := []byte(value.Get("payload"))
	if len(eventstr) == 0 {
		return
	}
	err = json.Unmarshal(eventstr, event)
	if err != nil {
		fmt.Printf("decode event error : %s", err)
		return
	}
	fmt.Printf("repo name is : %s body: %s/n", *event.Repo.FullName, *event.Comment.Body)
	config := issue.NewConfig("", "")
	err = issue.Process(config, *event)
	if err != nil {
		fmt.Printf("promote error %s", err)
	}
}

func payEvent(body []byte) *issue.IssueCommentEvent {
	event := &issue.IssueCommentEvent{}
	json.Unmarshal(body, event)
	config := issue.NewConfig("", "")
	issue.Regist("/pay", &pay.Pay{})
	err := issue.Process(config, *event)
	if err != nil {
		fmt.Printf("promote error %s", err)
		return nil
	}
	return event
}

func main() {
	issue.Regist("/pay", &pay.Pay{})
	fmt.Println("FunctionCompute go runtime inited.")
	http.HandleFunc("/", handler)
	port := os.Getenv("FC_SERVER_PORT")
	if port == "" {
		port = "9000"
	}
	http.ListenAndServe(":"+port, nil)
}
