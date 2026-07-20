package model

import (
	"database/sql"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames        = builder.RawFieldNames(&User{})
	userRows              = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet = strings.Join(stringx.Remove(userFieldNames,
		"`id`", "`create_time`", "`update_time`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames,
		"`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	UserModel interface {
		Insert(data User) (sql.Result, error)
		FindOne(id int64) (*User, error)
		FindOneByUser(user string) (*User, error)
		FindOneByMobile(mobile string) (*User, error)
		FindOneByName(name string) (*User, error)
		Update(data User) error
		Delete(id int64) error
	}

	defaultUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	User struct {
		ID         int64     `db:"id"`
		User       string    `db:"user"`
		Name       string    `db:"name"`
		Password   string    `db:"password"`
		Mobile     string    `db:"mobile"`
		Gender     string    `db:"gender"`
		Nickname   string    `db:"nickname"`
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func NewUserModel(conn sqlx.SqlConn) UserModel { _ = "STUB: not implemented"; return *new(UserModel) }

func (m *defaultUserModel) Insert(data User) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (m *defaultUserModel) FindOne(id int64) (*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *defaultUserModel) FindOneByUser(user string) (*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *defaultUserModel) FindOneByMobile(mobile string) (*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *defaultUserModel) FindOneByName(name string) (*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *defaultUserModel) Update(data User) error { _ = "STUB: not implemented"; return nil }

func (m *defaultUserModel) Delete(id int64) error { _ = "STUB: not implemented"; return nil }
