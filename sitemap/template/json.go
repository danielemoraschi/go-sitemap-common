package template

import (
    "github.com/danielemoraschi/go-sitemap-common/http"
    "github.com/danielemoraschi/go-sitemap-common/sitemap"
    "encoding/json"
)


type JsonUrlSet struct {
    TemplateInterface   `json:"-"`
    UrlSet []JsonUrl    `json:"urlset"`
}


type JsonUrl struct {
    Url       string        `json:"url"`
    Loc       string        `json:"loc"`
    Frequency sitemap.ChangeFreq    `json:"changefreq,omitempty"`
    Priority  float64       `json:"priority,omitempty"`
}


func JsonUrlSetFactory() TemplateInterface {
    return &JsonUrlSet{}
}


func (urlSet *JsonUrlSet) Set(data []http.HttpResource) TemplateInterface {
    urlSet.UrlSet = []JsonUrl{}
    for _, el := range data {
        urlSet.UrlSet = append(urlSet.UrlSet, JsonUrl{
            Loc: el.String(),
            Frequency: sitemap.MONTHLY,
            Priority: 1.0,
        })
    }
    return urlSet
}


// Generate serializes the sitemap URLSet to Json.
func (urlSet *JsonUrlSet) Generate() (siteMapJson []byte, err error) {
    if len(urlSet.UrlSet) > sitemap.MAXURLS {
        err = sitemap.ErrExceededMaxURLs
        return
    }

    var urlSetJson []byte
    urlSetJson, err = json.MarshalIndent(urlSet, "", "    ")

    if err == nil {
        siteMapJson = append(siteMapJson, urlSetJson...)
    }

    if len(siteMapJson) > sitemap.MAXFILESIZE {
        err = sitemap.ErrExceededMaxFileSize
    }

    return siteMapJson, err
}

