package api

type OpCodeInfo interface {
	String() string
	Code() byte
}

type Tracer interface {
	LogState(pc uint64, opcode OpCodeInfo)
}
