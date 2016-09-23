package policy

import (
    "regexp"
    _ "fmt"
)

var /* const */ pattern =
    regexp.MustCompile(`.*(\\.(css|js|gif|jpg|jpeg|png|mp3|zip|gz|tar|7z))`)

type ValidExtensionPolicy struct {
    PolicyInterface
}

func (p *ValidExtensionPolicy) ShouldVisit(u string) bool {
    match := pattern.MatchString(u)
    if match {
        //fmt.Printf("Match: %s\n", u)
    }
    return !match
}

func ValidExtensionPolicyFactory() PolicyInterface {
    return &ValidExtensionPolicy{}
}
