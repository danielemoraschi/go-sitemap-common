package fetcher

import (
    "fmt"
    "github.com/danielemoraschi/go-sitemap-common/http"
)

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    body string
    urls []string
}

func (f fakeFetcher) Fetch(httpResource http.HttpResource) (string, []string, error) {
    if res, ok := f[httpResource.Url()]; ok {
        return res.body, res.urls, nil
    }

    return "", nil, fmt.Errorf("Not found: %s", httpResource.Url())
}

// fetcher is a populated fakeFetcher.
var FakeFetcher = fakeFetcher{
    "http://golang.org/": &fakeResult{
        "The Go Programming Language",
        []string{
            "http://go.com/",
            "http://golang.org/pkg/",
            "http://golang.org/cmd/",
        },
    },
    "http://golang.org/pkg/": &fakeResult{
        "Packages",
        []string{
            "http://golang.org/",
            "http://golang.org/cmd/",
            "http://golang.org/pkg/fmt/",
            "http://golang.org/pkg/os/",
        },
    },
    "http://golang.org/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
            "http://golang.org/pkg/fmt/03",
        },
    },
    "http://golang.org/pkg/os/": &fakeResult{
        "Package os",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
            "http://golang.org/pkg/os/03",
        },
    },
    "http://go.com/": &fakeResult{
        "Go Base",
        []string{
            "http://go.com/re1",
            "http://go.com/re2",
        },
    },
    "http://go.com/re1": &fakeResult{
        "Go R1",
        []string{
            "http://golang.org/",
        },
    },
}
