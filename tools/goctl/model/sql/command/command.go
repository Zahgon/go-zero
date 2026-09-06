package command

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/config"
)

var (
	VarStringSrc string

	VarStringDir string

	VarBoolCache bool

	VarBoolIdea bool

	VarStringURL string

	VarStringSliceTable []string

	VarStringStyle string

	VarStringDatabase string

	VarStringSchema string

	VarStringHome string

	VarStringRemote string

	VarStringBranch string

	VarBoolStrict bool

	VarStringSliceIgnoreColumns []string

	VarStringCachePrefix string
)

var errNotMatched = errors.New("sql not matched")

func MysqlDDL(_ *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

func MySqlDataSource(_ *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

func mergeColumns(columns []string) []string { _ = "STUB: not implemented"; return nil }

type pattern map[string]struct{}

func (p pattern) Match(s string) bool { _ = "STUB: not implemented"; return false }

func (p pattern) list() []string { _ = "STUB: not implemented"; return nil }

func parseTableList(tableValue []string) pattern { _ = "STUB: not implemented"; return *new(pattern) }

func PostgreSqlDataSource(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

type ddlArg struct {
	src, dir      string
	cfg           *config.Config
	cache, idea   bool
	database      string
	strict        bool
	ignoreColumns []string
	prefix        string
}

func fromDDL(arg ddlArg) error { _ = "STUB: not implemented"; return nil }

type dataSourceArg struct {
	url, dir      string
	tablePat      pattern
	cfg           *config.Config
	cache, idea   bool
	strict        bool
	ignoreColumns []string
	prefix        string
}

func fromMysqlDataSource(arg dataSourceArg) error { _ = "STUB: not implemented"; return nil }

type pgDataSourceArg struct {
	url, dir      string
	tablePat      pattern
	schema        string
	cfg           *config.Config
	cache, idea   bool
	strict        bool
	ignoreColumns []string
	prefix        string
}

func fromPostgreSqlDataSource(arg pgDataSourceArg) error { _ = "STUB: not implemented"; return nil }
