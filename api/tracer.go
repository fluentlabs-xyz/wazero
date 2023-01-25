package api

import "fmt"

type Tracer interface {
	LogState(pc uint64, opcode fmt.Stringer)
}
