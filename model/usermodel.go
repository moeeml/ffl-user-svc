package model

import (
	"database/sql"
	"fmt"
	"strings"

	sx "github.com/jmoiron/sqlx"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	userFieldNames          = builderx.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	UserModel interface {
		Insert(data User) (sql.Result, error)
		FindOne(id int64) (*User, error)
		FindOneByAccount(account string) (*User, error)
		FindByIds(ids []int64) (*[]User, error)
		Update(data User) error
		Delete(id int64) error
	}

	defaultUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	User struct {
		Id         int64  `db:"id"`
		Account    string `db:"account"`
		Password   string `db:"password"`
		Name       string `db:"name"`
		Avatar     string `db:"avatar"`
		CreateTime int64  `db:"create_time"`
		UpdateTime int64  `db:"update_time"`
		Status     string `db:"status"`
	}
)

func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &defaultUserModel{
		conn:  conn,
		table: "`user`",
	}
}

func (m *defaultUserModel) Insert(data User) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Account, data.Password, data.Name, data.Avatar, data.Status)
	return ret, err
}

func (m *defaultUserModel) FindOne(id int64) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	var resp User
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindByIds(ids []int64) (*[]User, error) {
	query := fmt.Sprintf("select %s from %s where `id` in (?) limit 1", userRows, m.table)
	var resp []User

	q, args, _ := sx.In(query, ids)
	err := m.conn.QueryRows(&resp, q, args...)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByAccount(account string) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `account` = ? limit 1", userRows, m.table)
	var resp User
	err := m.conn.QueryRow(&resp, query, account)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Update(data User) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Account, data.Password, data.Name, data.Avatar, data.Status, data.Id)
	return err
}

func (m *defaultUserModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
