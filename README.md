#mustache.goku

mustache.goku is a [mustache](https://github.com/hoisie/mustache) like 
template engine for [goku](https://github.com/QLeelulu/goku)

##Installation

To install goku, simply run `go get github.com/QLeelulu/mustache.goku`. To use it in a program, use `import "github.com/QLeelulu/gmustache"`

##Usage

use in goku like this

    te := NewMustacheTemplateEngine()
    te.UseCache = true

    sc := &goku.ServerConfig{
        TemplateEnginer: te,
    }
    goku.CreateServer(routeTable, nil, sc)

