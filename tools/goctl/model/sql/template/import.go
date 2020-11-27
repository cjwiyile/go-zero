package template

var (
	Imports = `import (
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"github.com/yileCJW/go-zero/core/stores/cache"
	"github.com/yileCJW/go-zero/core/stores/sqlc"
	"github.com/yileCJW/go-zero/core/stores/sqlx"
	"github.com/yileCJW/go-zero/core/stringx"
	"github.com/yileCJW/go-zero/tools/goctl/model/sql/builderx"
)
`
	ImportsNoCache = `import (
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"github.com/yileCJW/go-zero/core/stores/sqlc"
	"github.com/yileCJW/go-zero/core/stores/sqlx"
	"github.com/yileCJW/go-zero/core/stringx"
	"github.com/yileCJW/go-zero/tools/goctl/model/sql/builderx"
)
`
)
