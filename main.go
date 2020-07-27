package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type blog struct {
	Categories []string `json:"categories"`
	Tags       []string `json:"tags"`
	Permalink  string   `json:"permalink"`
	Title      string   `json:"title"`
	Contents   string   `json:"contents"`
}

type Blogs []struct {
	collet []blog
}

var (
	githubUserName = "oldthreefeng"
	max            = 10
)

func main() {
	res := make([]blog, 0)
	response, data, errs := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		Get("https://www.fenghong.tech/index.json").EndStruct(&res)
	if errs != nil || http.StatusOK != response.StatusCode {
		log.Fatal(errs, data)
	}
	buf := &bytes.Buffer{}
	buf.WriteString("\n\n")
	cstSh, _ := time.LoadLocation("Asia/Shanghai")
	updated := time.Now().In(cstSh).Format("2006-01-02 15:04:05")
	buf.WriteString("### 我的近期动态\n\n⭐️ Star [个人主页](https://github.com/" + githubUserName + "/" + githubUserName + ") 后会自动更新，最近更新时间：`" + updated + "`\n\n📝")
	for k, v := range res {
		if k > max {
			break
		}
		url := v.Permalink
		title := v.Title
		comma := strings.Split(v.Contents, ".")
		content := comma[0]
		buf.WriteString("* " + " [" + title + "](" + url + ")\n\n" + "  > " + content + "\n")
	}
	buf.WriteString("\n\n")

	fmt.Println(buf.String())

	readme, err := ioutil.ReadFile("README.md")
	if nil != err {
		log.Fatalf("read README.md failed: %s", data)
	}
	startFlag := []byte("<!--events start -->")
	beforeStart := readme[:bytes.Index(readme, startFlag)+len(startFlag)]
	newBeforeStart := make([]byte, len(beforeStart))
	copy(newBeforeStart, beforeStart)
	endFlag := []byte("<!--events end -->")
	afterEnd := readme[bytes.Index(readme, endFlag):]
	newAfterEnd := make([]byte, len(afterEnd))
	copy(newAfterEnd, afterEnd)
	newReadme := append(newBeforeStart, buf.Bytes()...)
	newReadme = append(newReadme, newAfterEnd...)
	if err := ioutil.WriteFile("README.md", newReadme, 0644); nil != err {
		log.Fatalf("write README.md failed: %s", data)
	}
}
