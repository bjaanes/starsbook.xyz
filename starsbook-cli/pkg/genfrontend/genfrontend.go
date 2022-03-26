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

//go:embed templates/projects.ts.tpl
var projects string

func ProjectFiles(c conf.Conf) error {
	projectRoutesTmpl, err := template.New("projectRoutesTmpl").Parse(projectRoutes)
	if err != nil {
		return errors.Wrap(err, 0)
	}
	projectsRoutesFile, err := os.Create(fmt.Sprintf("src/generated/projectRoutes.ts"))
	if err != nil {
		return errors.Wrap(err, 0)
	}
	if err := projectRoutesTmpl.Execute(projectsRoutesFile, c); err != nil {
		return errors.Wrap(err, 0)
	}

	projectsTmpl, err := template.New("projectRoutesTmpl").Parse(projects)
	if err != nil {
		return errors.Wrap(err, 0)
	}
	projectsFile, err := os.Create(fmt.Sprintf("src/generated/projects.ts"))
	if err != nil {
		return errors.Wrap(err, 0)
	}
	if err := projectsTmpl.Execute(projectsFile, c); err != nil {
		return errors.Wrap(err, 0)
	}

	return nil
}
