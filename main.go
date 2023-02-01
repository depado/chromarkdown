package main

import (
	"bytes"
	"html/template"
	"os"

	"github.com/Depado/bfchroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/gobuffalo/packr"
	bf "github.com/russross/blackfriday/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GlobCSS is a byte slice containing the style CSS of the renderer
var GlobCSS template.CSS

// render takes a []byte input and runs the mardown render (with the bfchroma
// plugin enabled and with default values)
func render(input []byte) []byte {
	// Flags and extensions setup for blackfriday
	var exts = bf.NoIntraEmphasis | bf.Tables | bf.FencedCode | bf.Autolink |
		bf.Strikethrough | bf.SpaceHeadings | bf.BackslashLineBreak |
		bf.DefinitionLists | bf.Footnotes
	var flags = bf.UseXHTML | bf.Smartypants | bf.SmartypantsFractions |
		bf.SmartypantsDashes | bf.SmartypantsLatexDashes
	if !viper.GetBool("no-toc") {
		flags = flags | bf.TOC
	}

	// Setup the renderer
	r := bfchroma.NewRenderer(
		bfchroma.WithoutAutodetect(),
		bfchroma.Extend(bf.NewHTMLRenderer(bf.HTMLRendererParameters{Flags: flags})),
		bfchroma.Style(viper.GetString("theme")),
		bfchroma.ChromaOptions(html.WithClasses(true)),
	)

	// GlobalCSS component
	if GlobCSS == "" && r.Formatter.Classes {
		b := new(bytes.Buffer)
		if err := r.Formatter.WriteCSS(b, r.Style); err != nil {
			logrus.WithError(err).Warning("Couldn't write CSS")
		}
		GlobCSS = template.CSS(b.String())
	}

	// Run the renderer
	return bf.Run(
		input,
		bf.WithRenderer(r),
		bf.WithExtensions(exts),
	)
}

var rootCmd = &cobra.Command{
	Use:   "chromarkdown [input file]",
	Short: "Chromarkdown is a simple markdown-to-html renderer",
	Long: `Chromarkdown uses a combination of blackfriday and chroma to render an input markdown file.
It generates standalone HTML files that includes fonts, a grid system and extra CSS.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var fd *os.File
		var t *template.Template
		var in []byte

		box := packr.NewBox("./templates")
		s, err := box.FindString("index.tmpl")
		if err != nil {
			logrus.WithError(err).Fatal("Unable to find template")
		}
		if t, err = template.New("output").Parse(s); err != nil {
			logrus.WithError(err).Fatal("Couldn't parse template")
		}

		if fd, err = os.Create(viper.GetString("output")); err != nil {
			logrus.WithError(err).Fatal("Couldn't create file")
		}
		defer fd.Close() // nolint: errcheck

		if in, err = os.ReadFile(args[0]); err != nil {
			logrus.WithError(err).Fatal("Couldn't read in.md")
		}
		err = t.ExecuteTemplate(fd, "output", map[string]interface{}{
			"title":    viper.GetString("title"),
			"rendered": template.HTML(string(render(in))), // nolint: gas
			"css":      GlobCSS,
		})
		if err != nil {
			logrus.WithError(err).Fatal("Couldn't execute template")
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("output", "o", "out.html", "specify the path of the output HTML")
	rootCmd.PersistentFlags().StringP("title", "t", "Ouput", "Specify the title of the HTML page")
	rootCmd.PersistentFlags().String("theme", "monokai", "Specify the theme for syntax highlighting")
	rootCmd.PersistentFlags().Bool("no-toc", false, "Disable the table of content")
	if err := viper.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		logrus.WithError(err).Fatal("Couldn't bind flags")
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Fatal("Couldn't run root command")
	}
}
