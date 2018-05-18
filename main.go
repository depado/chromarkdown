package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"

	"github.com/Depado/bfchroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/sirupsen/logrus"
	bf "gopkg.in/russross/blackfriday.v2"
)

var exts = bf.NoIntraEmphasis | bf.Tables | bf.FencedCode | bf.Autolink |
	bf.Strikethrough | bf.SpaceHeadings | bf.BackslashLineBreak |
	bf.DefinitionLists | bf.Footnotes

var flags = bf.UseXHTML | bf.Smartypants | bf.SmartypantsFractions |
	bf.SmartypantsDashes | bf.SmartypantsLatexDashes | bf.TOC

// GlobCSS is a byte slice containing the style CSS of the renderer
var GlobCSS template.CSS

func render(input []byte) []byte {
	r := bfchroma.NewRenderer(
		bfchroma.WithoutAutodetect(),
		bfchroma.Extend(
			bf.NewHTMLRenderer(bf.HTMLRendererParameters{Flags: flags}),
		),
		bfchroma.Style("monokai"),
		bfchroma.ChromaOptions(html.WithClasses()),
	)
	if GlobCSS == "" && r.Formatter.Classes {
		b := new(bytes.Buffer)
		if err := r.Formatter.WriteCSS(b, r.Style); err != nil {
			logrus.WithError(err).Warning("Couldn't write CSS")
		}
		GlobCSS = template.CSS(b.String())
	}
	return bf.Run(
		input,
		bf.WithRenderer(r),
		bf.WithExtensions(exts),
	)
}

func main() {
	var err error
	var fd *os.File
	var t *template.Template
	var in []byte

	if fd, err = os.Create("out.html"); err != nil {
		logrus.WithError(err).Fatal("Couldn't create file")
	}
	defer fd.Close()

	if in, err = ioutil.ReadFile("in.md"); err != nil {
		logrus.WithError(err).Fatal("Couldn't read in.md")
	}
	if t, err = template.ParseFiles("templates/index.tmpl"); err != nil {
		logrus.WithError(err).Fatal("Couldn't parse template")
	}
	err = t.ExecuteTemplate(fd, "index.tmpl", map[string]interface{}{
		"title":    "Output",
		"rendered": template.HTML(string(render(in))),
		"css":      GlobCSS,
	})
	if err != nil {
		logrus.WithError(err).Fatal("Couldn't execute template")
	}
}
