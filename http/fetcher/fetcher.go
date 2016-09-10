package fetcher

import (
	"github.com/danielemoraschi/go-sitemap-common/http"
)

type FetcherInterface interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(res *http.HttpResource) (body []byte, err error)
}
