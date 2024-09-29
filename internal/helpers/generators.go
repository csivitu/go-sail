package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// GenerateDatabaseFile generates the database initialization file
func GenerateDatabaseFile(folderPath string, provider Provider) error {
	filename := filepath.Join(folderPath, "database.go")
	tmpl, err := template.New("database").Parse(`
package initializers

import (
    "fmt"
    {{range .Imports}}
    {{.}}
    {{- end}}
)

var DB {{.DBVariable}}

func ConnectDB(){
    {{.ConnectionCode}}
}
`)
	if err != nil {
		return fmt.Errorf("error parsing database template: %w", err)
	}

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating database file: %w", err)
	}
	defer f.Close()

	data := struct {
		Imports        []string
		ConnectionCode string
		DBVariable     string
	}{
		Imports:        provider.GetImports(),
		ConnectionCode: provider.GetConnectionCode(),
		DBVariable:     provider.GetDBVariable(),
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		return fmt.Errorf("error executing database template: %w", err)
	}

	return nil
}

// GenerateMigrationFile generates the migration file
func GenerateMigrationFile(folderPath string, provider Provider) error {
	filename := filepath.Join(folderPath, "migrations.go")
	tmpl, err := template.New("migration").Parse(`
package initializers

import (
    "fmt"
    {{range .Imports}}
    {{.}}
    {{- end}}
)

func DBMigrate() error {
    {{.MigrationCode}}
    return nil
}
`)
	if err != nil {
		return fmt.Errorf("error parsing migration template: %w", err)
	}

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating migration file: %w", err)
	}
	defer f.Close()

	data := struct {
		Imports       []string
		MigrationCode string
	}{
		Imports:       provider.GetImports(),
		MigrationCode: provider.GetMigrationCode(),
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		return fmt.Errorf("error executing migration template: %w", err)
	}

	return nil
}
