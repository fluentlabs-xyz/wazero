package api

type OpCodeInfo interface {
	String() string
	Code() byte
}

type Tracer interface {
	BeforeState(pc uint64, opcode OpCodeInfo, stack []uint64)
	AfterState(pc uint64, opcode OpCodeInfo, stack []uint64)
}
