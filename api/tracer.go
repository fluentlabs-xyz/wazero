package api

type Tracer interface {
	LogState(pc uint64, opcode uint16)
}
