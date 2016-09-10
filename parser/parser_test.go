package parser

import (
    "testing"
    "fmt"
    "github.com/danielemoraschi/go-sitemap-common/http"
)

func TestRegexParser_Parse(t *testing.T) {
    content :=
        "<a href='www.google.com'>re</a>" +
        "<a href='www.yahoo.com'>re</a>" +
        "<a href='/sub1'>re</a>" +
        "<a href='#sub2'>re</a>" +
        "<a title='sdsds' href=\"http://www.google.com\">re</a>" +
        "<a title='sdsds' href=\"https://www.amazon.com\">re</a>"

    res, _ := http.HttpResourceFactory("http://www.google.com", "")
    res.SetContent([]byte(content))

    urls, _ := HttpParser{}.Parse(res)

    for _, item := range urls {
        fmt.Printf("Got: %v\n", item)
    }
}
