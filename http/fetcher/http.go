package fetcher

import (
    "log"
    "io/ioutil"
    "github.com/danielemoraschi/go-sitemap-common/http"
)

func (res *HttpResource) Fetch(httpResource http.HttpResource) (body []byte, err error) {
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
