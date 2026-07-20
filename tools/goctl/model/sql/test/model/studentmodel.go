package model

import (
	"database/sql"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	studentFieldNames          = builder.RawFieldNames(&Student{})
	studentRows                = strings.Join(studentFieldNames, ",")
	studentRowsExpectAutoSet   = strings.Join(stringx.Remove(studentFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	studentRowsWithPlaceHolder = strings.Join(stringx.Remove(studentFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheStudentIdPrefix        = "cache#student#id#"
	cacheStudentClassNamePrefix = "cache#student#class#name#"
)

type (
	StudentModel interface {
		Insert(data Student) (sql.Result, error)
		FindOne(id int64) (*Student, error)
		FindOneByClassName(class, name string) (*Student, error)
		Update(data Student) error

		Delete(id int64, className, studentName string) error
	}

	defaultStudentModel struct {
		sqlc.CachedConn
		table string
	}

	Student struct {
		Id         int64           `db:"id"`
		Class      string          `db:"class"`
		Name       string          `db:"name"`
		Age        sql.NullInt64   `db:"age"`
		Score      sql.NullFloat64 `db:"score"`
		CreateTime time.Time       `db:"create_time"`
		UpdateTime sql.NullTime    `db:"update_time"`
	}
)

func NewStudentModel(conn sqlx.SqlConn, c cache.CacheConf) StudentModel {
	_ = "STUB: not implemented"
	return *new(StudentModel)
}

func (m *defaultStudentModel) Insert(data Student) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (m *defaultStudentModel) FindOne(id int64) (*Student, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *defaultStudentModel) FindOneByClassName(class, name string) (*Student, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *defaultStudentModel) Update(data Student) error { _ = "STUB: not implemented"; return nil }

func (m *defaultStudentModel) Delete(id int64, className, studentName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *defaultStudentModel) formatPrimary(primary any) string {
	_ = "STUB: not implemented"
	return ""
}

func (m *defaultStudentModel) queryPrimary(conn sqlx.SqlConn, v, primary any) error {
	_ = "STUB: not implemented"
	return nil
}
