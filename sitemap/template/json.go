package template

import (
    "github.com/danielemoraschi/go-sitemap-common/http"
    "encoding/json"
)


type JsonUrlSet struct {
    TemplateInterface   `json:"-"`
    UrlSet []JsonUrl    `json:"urlset"`
}


type JsonUrl struct {
    Url       string        `json:"url"`
    Loc       string        `json:"loc"`
    Frequency ChangeFreq    `json:"changefreq,omitempty"`
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
            Frequency: MONTHLY,
            Priority: 1.0,
        })
    }
    return urlSet
}


// Generate serializes the sitemap URLSet to Json.
func (urlSet *JsonUrlSet) Generate() (siteMapJson []byte, err error) {
    if len(urlSet.UrlSet) > MAXURLS {
        err = ErrExceededMaxURLs
        return
    }

    var urlSetJson []byte
    urlSetJson, err = json.Marshal(urlSet)

    if err == nil {
        siteMapJson = append(siteMapJson, urlSetJson...)
    }

    if len(siteMapJson) > MAXFILESIZE {
        err = ErrExceededMaxFileSize
    }

    return siteMapJson, err
}

