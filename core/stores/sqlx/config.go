package sqlx

import "errors"

var (
	errEmptyDatasource = errors.New("empty datasource")
	errEmptyDriverName = errors.New("empty driver name")
)

type SqlConf struct {
	DataSource string
	DriverName string   `json:",default=mysql"`
	Replicas   []string `json:",optional"`
	Policy     string   `json:",default=round-robin,options=round-robin|random"`
}

func (sc SqlConf) Validate() error { _ = "STUB: not implemented"; return nil }
