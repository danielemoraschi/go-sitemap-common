package crawler

import (
    "fmt"
    "sync"
    "github.com/danielemoraschi/go-sitemap-common/policy"
    "github.com/danielemoraschi/go-sitemap-common/http"
    "github.com/danielemoraschi/go-sitemap-common/http/fetcher"
    "github.com/danielemoraschi/go-sitemap-common/parser"
)


type Urls struct {
    sync.RWMutex
    data []http.HttpResource
}

func (s *Urls) AddList(values []http.HttpResource) {
    s.Lock()
    s.data = append(s.data, values...)
    s.Unlock()
}

func (s *Urls) Add(values http.HttpResource) {
    s.Lock()
    s.data = append(s.data, values)
    s.Unlock()
}

func (s *Urls) Reset() {
    s.Lock()
    s.data = []http.HttpResource{}
    s.Unlock()
}

func (s *Urls) Data() []http.HttpResource {
    s.RLock()
    n := s.data
    s.RUnlock()
    return n
}

func (s *Urls) Count() int {
    s.RLock()
    n := len(s.data)
    s.RUnlock()
    return n
}

func (s *Urls) RemoveDuplicatesUnordered() {
    encountered := map[string]http.HttpResource{}
    // Create a map of all unique elements.
    for v, el := range s.data {
        encountered[s.data[v].String()] = el
    }

    s.Reset()
    for _, el := range encountered {
        s.Add(el)
    }
}


// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(semaphore chan bool, urlsList *Urls, wg *sync.WaitGroup, url string, depth int,
    fetcher fetcher.FetcherInterface, parser parser.ParserInterface, policies []policy.PolicyInterface) {

    defer func() { <-semaphore }()
    defer wg.Done()

    policy.UrlAllowedByPolicies(policies, url)

    fmt.Printf("Visiting: %s\n", url)

    urlResource, err := http.HttpResourceFactory(url, "")
    if err != nil {
        fmt.Println("HttpResource URL error: ", err)
        return
    }

    _, err = fetcher.Fetch(&urlResource)
    if err != nil {
        fmt.Println("Fetch error: ", err)
        return
    }

    urls, err := parser.Parse(urlResource)
    if err != nil {
        fmt.Println("Parse error: ", err)
        return
    }

    //fmt.Printf("Found: %s\n", urls)

    urlsList.AddList(urls)

    if depth <= 0 {
        return
    }

    for _, u := range urls {
        if policy.UrlAllowedByPolicies(policies, u.String()) {
            wg.Add(1)
            go Crawl(semaphore, urlsList, wg, u.String(), depth-1, fetcher, parser, policies)
        }
    }

    return
}
