package parser

import (
    "github.com/danielemoraschi/go-sitemap-common/http"
    "regexp"
    "fmt"
)

type ParserInterface interface {
    // Parse parse the body of HttpResource and
    // return a slice of URLs found on that page.
    Parse(resource http.HttpResource) ([]http.HttpResource, error)
}

type HttpParser struct {
    ParserInterface
    urls []http.HttpResource
}

var /* const */ pattern =
    regexp.MustCompile(`<a\s[^>]*href=["']((?:\\.|[^("|')\\])*)["']`)


func (p HttpParser) Parse(resource http.HttpResource) ([]http.HttpResource, error) {
    var resources []http.HttpResource
    all := pattern.FindAllSubmatch(resource.Content(), -1)

    for _, items := range all {
        res, err := http.HttpResourceFactory(resource.String(), string(items[1]))
        if err != nil {
            fmt.Printf("Error %v\n", err)
        } else {
            //fmt.Printf("Matched %v\n", res.String())
            resources = append(resources, res)
        }
    }

    p.urls = resources
    return resources, nil
}