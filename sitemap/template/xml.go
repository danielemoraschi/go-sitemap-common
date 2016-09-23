package template

import (
    "encoding/xml"
    "github.com/danielemoraschi/go-sitemap-common"
)

var /* const */ Header =
    xml.Header("<?xml version=\"1.0\" encoding=\"UTF-8\"?>")


type Url struct {
    XMLName xml.Name    `xml:"url"`
    Loc string          `xml:"loc"`
    ChangeFreq string   `xml:"changefreq"`
    Priority string     `xml:"priority"`
}

type UrlSet struct {
    TemplateInterface
    XMLName xml.Name    `xml:"urlset"`
    Urls []Url          `xml:"url"`
}

func (urlSet *UrlSet) Generate(urlList crawler.Urls) (string, error) {
    for _, el := range urlList.Data() {
        urlSet.Urls = append(urlSet.Urls, Url{
            Loc: el.String(),
            ChangeFreq: "dayly",
            Priority: "0.5",
        })
    }

    return xml.MarshalIndent(urlSet, "", "    ")
}

