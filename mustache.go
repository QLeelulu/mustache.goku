package mustache

import (
    "github.com/QLeelulu/goku"
    "github.com/hoisie/mustache"
    "io"
)

// MustacheTemplateEngine
type MustacheTemplateEngine struct {
    ExtName       string
    UseCache      bool
    TemplateCache map[string]*mustache.Template
}

// template file ext name, default is ".html"
func (te *MustacheTemplateEngine) Ext() string {
    if te.ExtName == "" {
        return ".html"
    }
    return te.ExtName
}

// return whether the tempalte support layout
func (te *MustacheTemplateEngine) SupportLayout() bool {
    return true
}

func (te *MustacheTemplateEngine) Render(filepath string, layoutPath string, viewData *goku.ViewData, wr io.Writer) {
    var tmpl *mustache.Template
    tmpl = te.getTemplate(filepath)

    var r string
    if te.SupportLayout() && layoutPath != "" {
        layout := te.getTemplate(layoutPath)

        r = tmpl.RenderInLayout(layout, viewData)
    } else {
        r = tmpl.Render(viewData)
    }
    wr.Write([]byte(r))
}

func (te *MustacheTemplateEngine) getTemplate(filepath string) *mustache.Template {
    var tmpl *mustache.Template
    if te.UseCache {
        tmpl = te.TemplateCache[filepath]
    }
    if tmpl == nil {
        var err error
        tmpl, err = mustache.ParseFile(filepath)
        if err != nil {
            panic("MustacheTemplateEngine.Render: parse template \"" + filepath + "\" error, " + err.Error())
        }
        if te.UseCache {
            te.TemplateCache[filepath] = tmpl
        }
    }
    return tmpl
}

func NewMustacheTemplateEngine() *MustacheTemplateEngine {
    return &MustacheTemplateEngine{
        TemplateCache: make(map[string]*mustache.Template),
    }
}
