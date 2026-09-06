package parser

import (
	"github.com/zeromicro/go-zero/core/lang"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/parser/api/ast"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/parser/api/importstack"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/parser/api/placeholder"
)

const (
	atServerGroupKey  = "group"
	atServerPrefixKey = "prefix"
)

type API struct {
	Filename      string
	Syntax        *ast.SyntaxStmt
	info          *ast.InfoStmt
	importStmt    []ast.ImportStmt
	TypeStmt      []ast.TypeStmt
	ServiceStmts  []*ast.ServiceStmt
	importManager *importstack.ImportStack
	importSet     map[string]lang.PlaceholderType
}

func convert2API(a *ast.AST, importSet map[string]lang.PlaceholderType, is *importstack.ImportStack) (*API, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (api *API) checkImportStmt() error { _ = "STUB: not implemented"; return nil }

func (api *API) checkInfoStmt() error { _ = "STUB: not implemented"; return nil }

func (api *API) checkServiceStmt() error { _ = "STUB: not implemented"; return nil }

func (api *API) checkTypeStmt() error { _ = "STUB: not implemented"; return nil }

func (api *API) checkTypeDeclareContext() error { _ = "STUB: not implemented"; return nil }

func (api *API) checkTypeContext(declareContext map[string]placeholder.Type) error {
	_ = "STUB: not implemented"
	return nil
}

func (api *API) checkTypeExprContext(declareContext map[string]placeholder.Type, tp ast.DataType) error {
	_ = "STUB: not implemented"
	return nil
}

func (api *API) getAtServerValue(atServer *ast.AtServerStmt, key string) string {
	_ = "STUB: not implemented"
	return ""
}

func (api *API) mergeAPI(in *API) error { _ = "STUB: not implemented"; return nil }

func (api *API) parseImportedAPI(imports []ast.ImportStmt) ([]*API, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (api *API) parseReverse() error { _ = "STUB: not implemented"; return nil }

func (api *API) SelfCheck() error { _ = "STUB: not implemented"; return nil }
