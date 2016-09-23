package template

import (
    "github.com/danielemoraschi/go-sitemap-common"
)

type TemplateInterface interface {
    Generate(urlList crawler.Urls) (string, error)
}
