package policy

import "sync"

type UniqueUrlPolicy struct {
    sync.RWMutex
    urlMap map[string]bool
    PolicyInterface
}

func (p *UniqueUrlPolicy) ShouldVisit(url string) bool {
    p.Lock()
    defer p.Unlock()
    _, ok := p.urlMap[url]
    p.urlMap[url] = true
    return !ok
}

func UniqueUrlPolicyFactory() PolicyInterface {
    return &UniqueUrlPolicy{urlMap: make(map[string]bool)}
}
