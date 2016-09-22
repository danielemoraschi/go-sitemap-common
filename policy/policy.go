package policy


type PolicyInterface interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    ShouldVisit(url string) bool
}

func UrlAllowedByPolicies(policies []PolicyInterface, url string) bool {
    for i := 0; i < len(policies); i++ {
        if ! policies[i].ShouldVisit(url) {
            return false
        }
    }
    return true
}