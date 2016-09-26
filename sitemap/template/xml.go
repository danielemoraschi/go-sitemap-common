package template

import (
    "encoding/xml"
    "github.com/danielemoraschi/go-sitemap-common"
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

    HEADER  = `<?xml version="1.0" encoding="UTF-8"?>`
    XMLNS   = "http://www.sitemaps.org/schemas/sitemap/0.9"

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

type Url struct {
    XMLName xml.Name        `xml:"url"`
    Loc string              `xml:"loc"`
    Frequency ChangeFreq    `xml:"changefreq,omitempty"`
    Priority float64        `xml:"priority,omitempty"`
}

type UrlSet struct {
    TemplateInterface
    XMLName xml.Name    `xml:"urlset"`
    XMLNS string        `xml:"xmlns,attr"`
    Urls []Url          `xml:"url"`
}

func (urlSet *UrlSet) Generate(urlList crawler.Urls) (siteMapXML []byte, err error) {
    if urlList.Count() > MAXURLS {
        err = ErrExceededMaxURLs
        return
    }

    urlSet.XMLNS = XMLNS

    siteMapXML = []byte(HEADER)

    for _, el := range urlList.Data() {
        urlSet.Urls = append(urlSet.Urls, Url{
            Loc: el.String(),
            Frequency: DAILY,
            Priority: 0.5,
        })
    }

    XMLMap, err = xml.MarshalIndent(urlSet, "", "    ")
    if err == nil {
        sitemapXML = append(sitemapXML, urlsetXML...)
    }
}

