#mustache.goku

mustache.goku is a [mustache](https://github.com/hoisie/mustache) like 
template engine for [goku](https://github.com/QLeelulu/goku)

##Usage

use in goku like this

    te := NewMustacheTemplateEngine()
    te.UseCache = true

    sc := &goku.ServerConfig{
        TemplateEnginer: te,
    }
    goku.CreateServer(routeTable, nil, sc)

