package api

type OpCodeInfo interface {
	String() string
	Code() byte
	GetParams() []uint64
	Pc() uint64
}

type MemoryChangeInfo struct {
	Offset uint32
	Value  []byte
}

type Tracer interface {
	BeforeState(pc uint64, opcode OpCodeInfo, stack []uint64, memory *MemoryChangeInfo)
	AfterState(pc uint64, opcode OpCodeInfo, stack []uint64, memory *MemoryChangeInfo)
}
