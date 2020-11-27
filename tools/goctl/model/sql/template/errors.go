package template

var Error = `package {{.pkg}}

import "github.com/yilefreedom/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound
`
