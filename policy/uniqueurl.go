package policy

import "sync"

type UniqueUrlPolicy struct {
    sync.Mutex
    urlMap map[string]bool
    PolicyInterface
}

func (p UniqueUrlPolicy) ShouldVisit(url string) bool {
    p.Lock()
    _, ok := p.urlMap[url]
    p.urlMap[url] = true
    p.Unlock()

    return !ok
}

func UniqueUrlPolicyFactory() PolicyInterface {
    p := UniqueUrlPolicy{urlMap: make(map[string]bool)}
    return p
}
