package template

var (
	Imports = `import (
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"github.com/yilefreedom/go-zero/core/stores/cache"
	"github.com/yilefreedom/go-zero/core/stores/sqlc"
	"github.com/yilefreedom/go-zero/core/stores/sqlx"
	"github.com/yilefreedom/go-zero/core/stringx"
	"github.com/yilefreedom/go-zero/tools/goctl/model/sql/builderx"
)
`
	ImportsNoCache = `import (
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"github.com/yilefreedom/go-zero/core/stores/sqlc"
	"github.com/yilefreedom/go-zero/core/stores/sqlx"
	"github.com/yilefreedom/go-zero/core/stringx"
	"github.com/yilefreedom/go-zero/tools/goctl/model/sql/builderx"
)
`
)
