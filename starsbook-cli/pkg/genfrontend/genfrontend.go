package genfrontend

import (
	_ "embed"
	"fmt"
	"github.com/go-errors/errors"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/conf"
	"html/template"
	"os"
)

//go:embed templates/projectRoutes.ts.tpl
var projectRoutes string

func ProjectFiles(c conf.Conf) error {
	tmpl, err := template.New("projectRoutesTmpl").Parse(projectRoutes)
	if err != nil {
		return errors.Wrap(err, 0)
	}

	f, err := os.Create(fmt.Sprintf("src/generated/projectRoutes.ts"))
	if err != nil {
		return errors.Wrap(err, 0)
	}
	if err := tmpl.Execute(f, c); err != nil {
		return errors.Wrap(err, 0)
	}

	return nil
}
