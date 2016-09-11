package fetcher

import (
    "io/ioutil"
    "net/http"
    myhttp "github.com/danielemoraschi/go-sitemap-common/http"
)

type HttpFetcher struct {
    FetcherInterface
}

func (f HttpFetcher) Fetch(res *myhttp.HttpResource) ([]byte, error) {
    response, err := http.Get(res.String())

    if err != nil {
        return nil, err
    }

    defer response.Body.Close()

    content, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }
    res.SetContent(content)
    return content, nil
}
