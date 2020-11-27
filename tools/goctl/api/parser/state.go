package parser

import "github.com/yilefreedom/go-zero/tools/goctl/api/spec"

type state interface {
	process(api *spec.ApiSpec) (state, error)
}
