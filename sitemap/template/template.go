package template

import (
    "github.com/danielemoraschi/go-sitemap-common/http"
)

type TemplateInterface interface {
    Set(data []http.HttpResource) TemplateInterface
    Generate() ([]byte, error)
}
