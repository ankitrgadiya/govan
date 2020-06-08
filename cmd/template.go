package cmd

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Template body for HTML pages
var templString = `<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
<meta http-equiv="refresh" content="0; url={{.Redirect}}" />
{{- if .Package }}
<meta name="go-import" content="{{.Domain}}/{{.Package}} git {{.Source}}">
{{- end }}
{{- if .Dir }}
<meta name="go-source" content="{{.Domain}}/{{.Package}} _ {{.Dir}} {{.File}}">
{{- end }}
</head>
</html>
`

type templValue struct {
	Redirect string
	Domain   string
	Package  string
	Source   string
	Dir      string
	File     string
}

// generateHTML generates the static HTML for the given package
// TODO: Improve this function to support more VCS and more hosts.
func generateHTML(domain, pkg, source string) ([]byte, error) {
	var (
		dir, file, redirect string
		b                   bytes.Buffer
	)

	if pkg != "" {
		redirect = "https://pkg.go.dev/" + domain + "/" + pkg

		// Add the URL scheme if not already present
		if !strings.HasPrefix(source, "http") {
			source = "https://" + source
		}

		// Deduce go-source paths for the hosting
		switch path := urlMustParse(source); path.Host {
		case "github.com":
			dir = source + "/tree/master/{dir}"
			file = source + "/blob/master/{dir}/{file}#L{line}"
		case "gitlab.com":
			dir = source + "/-/tree/master/{dir}"
			file = source + "/-/blob/master/{dir}/{file}#L{line}"
		}
	} else {
		redirect = "https://" + domain
	}

	t := template.Must(template.New("package").Parse(templString))
	err := t.Execute(&b, &templValue{redirect, domain, pkg, source, dir, file})
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// writeHTML generates the static HTML and writes it to the file.
func writeHTML(domain, pkg, source, filename string) error {
	data, err := generateHTML(domain, pkg, source)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(filename), 0755)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0666)
}
