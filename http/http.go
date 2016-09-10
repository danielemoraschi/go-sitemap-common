package http

import (
    "net/url"
    "net/http"
    "log"
    "io/ioutil"
)

type HttpResource struct {
    url *url.URL
    content []byte
}

func (res *HttpResource) Url() string {
    return res.url.String()
}

func (res *HttpResource) SetContent(content []byte) []byte {
    res.content = content
    return res.content
}

func (res *HttpResource) Content() []byte {
    if len(res.content) > 0 {
        return res.content
    }

    response, err := http.Get(res.Url())
    if err != nil {
        log.Fatal(err)
    }

    defer response.Body.Close()

    res.content, err = ioutil.ReadAll(response.Body)

    if err != nil {
        log.Fatal(err)
    }

    return res.content
}

func HttpResourceFactory(baseUrl string) HttpResource {
    parsedUrl, err := url.ParseRequestURI(baseUrl)
    if err != nil {
        panic(err)
    }

    res := HttpResource{url: parsedUrl}
    return res
}