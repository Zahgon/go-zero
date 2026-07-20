package migrate

import (
	"go/ast"
	"go/token"

	"github.com/spf13/cobra"
)

const defaultMigrateVersion = "v1.3.0"

const (
	confirmUnknown = iota
	confirmAll
	confirmIgnore
)

var (
	fset            = token.NewFileSet()
	builderxConfirm = confirmUnknown
)

func migrate(_ *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

func rewriteImport(verbose bool) error { _ = "STUB: not implemented"; return nil }

func rewriteFile(pkgs map[string]*ast.Package, verbose bool) error {
	_ = "STUB: not implemented"
	return nil
}

func writeFile(pkgs []*ast.Package, verbose bool) error { _ = "STUB: not implemented"; return nil }

func replacePkg(file *ast.File) { _ = "STUB: not implemented"; return }

func refactorBuilderx(deprecated, replacement string, fn func(allow bool)) {
	_ = "STUB: not implemented"
	return
}
