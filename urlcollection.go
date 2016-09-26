package crawler

import (
    "sync"
    "github.com/danielemoraschi/go-sitemap-common/http"
)

type UrlCollectionInterface interface {
    Data() []http.HttpResource
}

type UrlCollection struct {
    sync.RWMutex
    data []http.HttpResource
}

func (s *UrlCollection) AddList(values []http.HttpResource) {
    s.Lock()
    s.data = append(s.data, values...)
    s.Unlock()
}

func (s *UrlCollection) Add(values http.HttpResource) {
    s.Lock()
    s.data = append(s.data, values)
    s.Unlock()
}

func (s *UrlCollection) Reset() {
    s.Lock()
    s.data = []http.HttpResource{}
    s.Unlock()
}

func (s *UrlCollection) Data() []http.HttpResource {
    s.RLock()
    n := s.data
    s.RUnlock()
    return n
}

func (s *UrlCollection) Count() int {
    s.RLock()
    n := len(s.data)
    s.RUnlock()
    return n
}

func (s *UrlCollection) RemoveDuplicatesUnordered() {
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
