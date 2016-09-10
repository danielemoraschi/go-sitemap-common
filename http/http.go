package http

import (
    "net/url"
)

type HttpResource struct {
    url *url.URL
    content []byte
}

func (res HttpResource) String() string {
    return res.url.String()
}

func (res *HttpResource) Url() *url.URL {
    return res.url
}

func (res *HttpResource) SetContent(content []byte) []byte {
    res.content = content
    return res.content
}

func (res *HttpResource) Content() []byte {
    return res.content
}

func HttpResourceFactory(baseUrl string, fragment string) (ret HttpResource, err error) {
    parsedUrl, err := url.ParseRequestURI(baseUrl)
    if err != nil {
        return ret, err
    }

    if fragment != "" {
        parsedFragment, err := url.Parse(fragment)
        if err != nil {
            return ret, err
        }
        parsedUrl = parsedUrl.ResolveReference(parsedFragment)
    }

    return HttpResource{url: parsedUrl}, nil
}