package policy

import (
    _ "fmt"
    "net/url"
)

type SameDomainPolicy struct {
    baseURL *url.URL
    PolicyInterface
}

func (p SameDomainPolicy) ShouldVisit(u string) bool {
    urlToCheck, err := url.ParseRequestURI(u)
    if err != nil {
        panic(err)
    }

    ret := p.baseURL.Host == urlToCheck.Host

    //fmt.Printf("Checking: %s vs %s = %v\n", p.baseURL.Host, urlToCheck.Host, ret)
    return ret
}

func SameDomainPolicyFactory(baseUrl string) PolicyInterface {
    URL, err := url.ParseRequestURI(baseUrl)
    if err != nil {
        panic(err)
    }
    p := SameDomainPolicy{baseURL: URL}
    return p
}