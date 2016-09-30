package template

import (
    "github.com/danielemoraschi/go-sitemap-common/http"
    "errors"
)

type ChangeFreq string

const (
    // MAXURLS is the maximum allowable number of URLs in a sitemap <urlset>,
    // per http://www.sitemaps.org/protocol.html#index.
    MAXURLS = 50000

    // MAXFILESIZE is the maximum allowable uncompressed size of a sitemap.xml
    // file, per http://www.sitemaps.org/protocol.html#index.
    MAXFILESIZE = 10 * 1024 * 1024

    ALWAYS  ChangeFreq = "always"
    HOURLY  ChangeFreq = "hourly"
    DAILY   ChangeFreq = "daily"
    WEEKLY  ChangeFreq = "weekly"
    MONTHLY ChangeFreq = "monthly"
    YEARLY  ChangeFreq = "yearly"
    NEVER   ChangeFreq = "never"
)

var (
    // ErrExceededMaxURLs is an error indicating that the sitemap has more
    // than the allowable MaxURLs URL entries.
    ErrExceededMaxURLs = errors.New("Exceeded maximum number of URLs in a sitemap <urlset>")

    // ErrExceededMaxFileSize is an error indicating that the sitemap or sitemap
    // index file size exceeds the allowable MaxFileSize byte size.
    ErrExceededMaxFileSize = errors.New("Exceeded maximum file size of a sitemap or sitemap index XML file")
)

type TemplateInterface interface {
    Set(data []http.HttpResource) TemplateInterface
    Generate() ([]byte, error)
}
