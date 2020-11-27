package fx

import "github.com/yilefreedom/go-zero/core/threading"

func Parallel(fns ...func()) {
	group := threading.NewRoutineGroup()
	for _, fn := range fns {
		group.RunSafe(fn)
	}
	group.Wait()
}
