// mustache.goku is a mustache like template engine for goku
//
// Installation:
// To install goku, simply run: 
//     go get github.com/QLeelulu/mustache.goku 
// To use it in a program, use: 
//     import "github.com/QLeelulu/gmustache"
//
// Usage:
// 
//    te := NewMustacheTemplateEngine()
//    te.UseCache = true
//
//    sc := &goku.ServerConfig{
//        TemplateEnginer: te,
//    }
//    goku.CreateServer(routeTable, nil, sc)
package mustache
