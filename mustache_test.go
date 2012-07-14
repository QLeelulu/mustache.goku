package mustache

import (
    "bytes"
    "github.com/QLeelulu/goku"
    "os"
    "path"
    "testing"
)

func TestMustacheEngine(t *testing.T) {
    filename := path.Join(path.Join(os.Getenv("PWD"), "test"), "test1.mustache")
    layout := path.Join(path.Join(os.Getenv("PWD"), "test"), "layout.mustache")

    te := NewMustacheTemplateEngine()

    viewData := new(goku.ViewData)
    viewData.Model = "lulu"
    buf := new(bytes.Buffer)
    te.Render(filename, layout, viewData, buf)
    r := buf.String()
    if r != "hello lulu." {
        t.Fatalf("Error: render got %s\n\t", r)
    }
}
