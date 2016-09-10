package fetcher

import (
    "fmt"
    "github.com/danielemoraschi/go-sitemap-common/http"
    "strings"
)

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    body string
    urls []string
}

func (f fakeFetcher) Fetch(res *http.HttpResource) ([]byte, error) {
    if res, ok := f[res.String()]; ok {
        return []byte(strings.Join(res.urls[:], ",")), nil
    }

    return nil, fmt.Errorf("Not found: %s", res.Url())
}

// fetcher is a populated fakeFetcher.
var FakeFetcher = fakeFetcher{
    "http://golang.org/": &fakeResult{
        "The Go Programming Language",
        []string{
            "<a href='http://go.com/'>link</a>",
            "<a href='http://golang.org/pkg/'>link</a>",
            "<a href='http://golang.org/cmd/'>link</a>",
        },
    },
    "http://golang.org/pkg/": &fakeResult{
        "Packages",
        []string{
            "<a href='http://golang.org/'>link</a>",
            "<a href='http://golang.org/cmd/'>link</a>",
            "<a href='http://golang.org/pkg/fmt/'>link</a>",
            "<a href='http://golang.org/pkg/os/'>link</a>",
        },
    },
    "http://golang.org/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string{
            "<a href='http://golang.org/'>link</a>",
            "<a href='http://golang.org/pkg/'>link</a>",
            "<a href='http://golang.org/pkg/fmt/03'>link</a>",
        },
    },
    "http://golang.org/pkg/os/": &fakeResult{
        "Package os",
        []string{
            "<a href='http://golang.org/'>link</a>",
            "<a href='http://golang.org/pkg/'>link</a>",
            "<a href='http://golang.org/pkg/os/03'>link</a>",
        },
    },
    "http://go.com/": &fakeResult{
        "Go Base",
        []string{
            "<a href='http://go.com/re1'>link</a>",
            "<a href='http://go.com/re2'>link</a>",
        },
    },
    "http://go.com/re1": &fakeResult{
        "Go R1",
        []string{
            "<a href='http://golang.org/'>link</a>",
        },
    },
}
