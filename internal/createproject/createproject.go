package createproject

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates
var templatesFS embed.FS

// data is available inside every template as {{.Name}} and {{.Module}}.
type data struct {
	Name   string
	Module string
}

func Run(name string) error {
	// 1. Create project folder
	if _, err := os.Stat(name); err == nil {
		return fmt.Errorf("project %q already exists", name)
	}
	if err := os.MkdirAll(name, 0o755); err != nil {
		return fmt.Errorf("could not create project %q: %w", name, err)
	}

	// If anything below fails, remove the half-built project.
	if err := scaffold(name); err != nil {
		_ = os.RemoveAll(name)
		return err
	}
	return nil
}

func scaffold(dir string) error {
	d := data{
		Name:   filepath.Base(dir),
		Module: filepath.Base(dir),
	}

	// 2. go mod init
	if err := goModInit(dir, d.Module); err != nil {
		return err
	}

	// 3 + 4. cmd/server/main.go and internal/* from templates
	return renderTemplates(dir, d)
}

func goModInit(dir, module string) error {
	cmd := exec.Command("go", "mod", "init", module)
	cmd.Dir = dir
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("go mod init failed: %w\n%s", err, out)
	}
	return nil
}

func renderTemplates(dir string, d data) error {
	return fs.WalkDir(templatesFS, "templates", func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		rel := strings.TrimPrefix(strings.TrimPrefix(path, "templates"), "/")
		if rel == "" {
			return nil
		}
		target := filepath.Join(dir, filepath.FromSlash(strings.TrimSuffix(rel, ".tmpl")))

		if entry.IsDir() {
			return os.MkdirAll(target, 0o755)
		}

		content, err := templatesFS.ReadFile(path)
		if err != nil {
			return err
		}
		tmpl, err := template.New(rel).Parse(string(content))
		if err != nil {
			return fmt.Errorf("parse template %s: %w", rel, err)
		}

		f, err := os.Create(target)
		if err != nil {
			return err
		}
		defer f.Close()

		if err := tmpl.Execute(f, d); err != nil {
			return fmt.Errorf("render %s: %w", rel, err)
		}
		return nil
	})
}
