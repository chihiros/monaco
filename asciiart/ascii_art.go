package asciiart

import (
	"bytes"
	"embed"
	"fmt"
	"log"
	"os/exec"
	"text/template"
)

//go:embed template/*.tpl
var templates embed.FS

type AsciiArt struct {
	Name string
	Args []string
}

func NewAsciiArt(args []string) *AsciiArt {
	return &AsciiArt{
		Name: "mona_front",
		Args: args,
	}
}

func getAaTemplate() *template.Template {
	funcMap := template.FuncMap{
		"printOverLine": func(args []string) string {
			return printLine("‾", args)
		},
		"printUnderLine": func(args []string) string {
			return printLine("_", args)
		},
		"echo": func(args []string) string {
			cmd := exec.Command("echo", args...)
			var out bytes.Buffer
			cmd.Stdout = &out
			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
			return out.String()
		},
	}

	TEMPLATE_PATH := "template/*.tpl"
	tpl := template.Must(template.New("").Funcs(funcMap).ParseFS(templates, TEMPLATE_PATH))

	return tpl
}

func (a *AsciiArt) Print() {
	tpl := getAaTemplate()
	var buf bytes.Buffer

	err := tpl.ExecuteTemplate(&buf, a.Name, a.Args)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf.String())
}
