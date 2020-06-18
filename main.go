package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %[1]s PACKAGENAME (eg: %[1]s main)", os.Args[0])
	}

	outputFileTpl := template.Must(template.New("output-tpl").Parse(`package {{.PackageName}}

var resources map[string][]byte = map[string][]byte{
	{{range .Resources}}
	"{{.FilePath}}": []byte{
		{{range .FileContents}}{{.}},{{end}}
	},
	{{end}}
}
`))

	packageName := os.Args[1]

	tplData := map[string]interface{}{
		"PackageName": packageName,
	}

	resources := make([]map[string]interface{}, 0)

	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			fileContents, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			resources = append(resources, map[string]interface{}{
				"FilePath":     path,
				"FileContents": fileContents,
			})

			return nil
		})
	if err != nil {
		log.Fatal(err)
	}

	tplData["Resources"] = resources

	var outputFileContents bytes.Buffer
	outputFileTpl.Execute(&outputFileContents, tplData)

	ioutil.WriteFile("resources.go", outputFileContents.Bytes(), 0644)
}
