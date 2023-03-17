package api

type OpCodeInfo interface {
	String() string
	Code() byte
	GetParams() []uint64
	Pc() uint64
}

type OpCodeDrop interface {
	Drop() uint32
}

type MemoryChangeInfo struct {
	Offset uint32
	Value  []byte
}

type Tracer interface {
	GlobalVariable(relativePc uint64, opcode OpCodeInfo, value uint64)

	BeforeState(relativePc uint64, opcode OpCodeInfo, stack []uint64, memory *MemoryChangeInfo)
	AfterState(relativePc uint64, opcode OpCodeInfo, stack []uint64, memory *MemoryChangeInfo)
}
