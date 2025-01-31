package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/yilefreedom/go-zero/core/stores/cache"
	"github.com/yilefreedom/go-zero/core/stores/sqlc"
	"github.com/yilefreedom/go-zero/core/stores/sqlx"
	"github.com/yilefreedom/go-zero/core/stringx"
	"github.com/yilefreedom/go-zero/tools/goctl/model/sql/builderx"
)

var (
	bookFieldNames          = builderx.FieldNames(&Book{})
	bookRows                = strings.Join(bookFieldNames, ",")
	bookRowsExpectAutoSet   = strings.Join(stringx.Remove(bookFieldNames, "create_time", "update_time"), ",")
	bookRowsWithPlaceHolder = strings.Join(stringx.Remove(bookFieldNames, "book", "create_time", "update_time"), "=?,") + "=?"

	bookPrefix = "cache#Book#book#"
)

type (
	BookModel struct {
		sqlc.CachedConn
		table string
	}

	Book struct {
		Book  string `db:"book"`  // book name
		Price int64  `db:"price"` // book price
	}
)

func NewBookModel(conn sqlx.SqlConn, c cache.CacheConf, table string) *BookModel {
	return &BookModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      table,
	}
}

func (m *BookModel) Insert(data Book) (sql.Result, error) {
	query := `insert into ` + m.table + ` (` + bookRowsExpectAutoSet + `) values (?, ?)`
	return m.ExecNoCache(query, data.Book, data.Price)
}

func (m *BookModel) FindOne(book string) (*Book, error) {
	bookKey := fmt.Sprintf("%s%v", bookPrefix, book)
	var resp Book
	err := m.QueryRow(&resp, bookKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := `select ` + bookRows + ` from ` + m.table + ` where book = ? limit 1`
		return conn.QueryRow(v, query, book)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *BookModel) Update(data Book) error {
	bookKey := fmt.Sprintf("%s%v", bookPrefix, data.Book)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := `update ` + m.table + ` set ` + bookRowsWithPlaceHolder + ` where book = ?`
		return conn.Exec(query, data.Price, data.Book)
	}, bookKey)
	return err
}

func (m *BookModel) Delete(book string) error {
	bookKey := fmt.Sprintf("%s%v", bookPrefix, book)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := `delete from ` + m.table + ` where book = ?`
		return conn.Exec(query, book)
	}, bookKey)
	return err
}
