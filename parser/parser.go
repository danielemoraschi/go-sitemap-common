package parser

import "github.com/danielemoraschi/go-sitemap-common/http"

type ParserInterface interface {
    // Parse parse the body of HttpResource and
    // return a slice of URLs found on that page.
    Parse(httpResource http.HttpResource) (urls []http.HttpResource, err error)
}

