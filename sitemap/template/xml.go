package template

import (
    "github.com/danielemoraschi/go-sitemap-common/http"
    "encoding/xml"
)

const (
    HEADER  = `<?xml version="1.0" encoding="UTF-8"?>`
    XMLNS   = "http://www.sitemaps.org/schemas/sitemap/0.9"
)


type XMLUrlSet struct {
    TemplateInterface
    XMLName xml.Name    `xml:"urlset"`
    XMLNS string        `xml:"xmlns,attr"`
    Urls []XMLUrl       `xml:"url"`
}


type XMLUrl struct {
    XMLName xml.Name        `xml:"url"`
    Loc string              `xml:"loc"`
    Frequency ChangeFreq    `xml:"changefreq,omitempty"`
    Priority float64        `xml:"priority,omitempty"`
}


func XMLUrlSetFactory() TemplateInterface {
    return &XMLUrlSet{}
}


func (urlSet *XMLUrlSet) Set(data []http.HttpResource) TemplateInterface {
    urlSet.Urls = []XMLUrl{}
    for _, el := range data {
        urlSet.Urls = append(urlSet.Urls, XMLUrl{
            Loc: el.String(),
            Frequency: MONTHLY,
            Priority: 1.0,
        })
    }
    return urlSet
}


// Generate serializes the sitemap URLSet to XML, with the <urlset> xmlns added
// and the XML preamble prepended.
func (urlSet *XMLUrlSet) Generate() (siteMapXML []byte, err error) {
    if len(urlSet.Urls) > MAXURLS {
        err = ErrExceededMaxURLs
        return
    }

    urlSet.XMLNS = XMLNS
    siteMapXML = []byte(HEADER)

    var urlSetXML []byte
    urlSetXML, err = xml.Marshal(urlSet)

    if err == nil {
        siteMapXML = append(siteMapXML, urlSetXML...)
    }

    if len(siteMapXML) > MAXFILESIZE {
        err = ErrExceededMaxFileSize
    }

    return siteMapXML, err
}

