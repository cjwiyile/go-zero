package template

var Error = `package {{.pkg}}

import "github.com/yileCJW/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound
`
