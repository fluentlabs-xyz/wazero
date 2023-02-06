package wasm

import (
	"encoding/binary"
	"github.com/tetratelabs/wazero/api"
)

type TraceMemoryInstance struct {
	del     api.Memory
	changes []api.MemoryChangeInfo
}

func NewTraceMemory(del api.Memory) api.Memory {
	return &TraceMemoryInstance{del: del}
}

func (m *TraceMemoryInstance) Definition() api.MemoryDefinition {
	return m.del.Definition()
}

func (m *TraceMemoryInstance) Size() uint32 {
	return m.del.Size()
}

func (m *TraceMemoryInstance) Grow(deltaPages uint32) (previousPages uint32, ok bool) {
	return m.del.Grow(deltaPages)
}

func (m *TraceMemoryInstance) ReadByte(offset uint32) (byte, bool) {
	return m.del.ReadByte(offset)
}

func (m *TraceMemoryInstance) ReadUint16Le(offset uint32) (uint16, bool) {
	return m.del.ReadUint16Le(offset)
}

func (m *TraceMemoryInstance) ReadUint32Le(offset uint32) (uint32, bool) {
	return m.del.ReadUint32Le(offset)
}

func (m *TraceMemoryInstance) ReadFloat32Le(offset uint32) (float32, bool) {
	return m.del.ReadFloat32Le(offset)
}

func (m *TraceMemoryInstance) ReadUint64Le(offset uint32) (uint64, bool) {
	return m.del.ReadUint64Le(offset)
}

func (m *TraceMemoryInstance) ReadFloat64Le(offset uint32) (float64, bool) {
	return m.del.ReadFloat64Le(offset)
}

func (m *TraceMemoryInstance) Read(offset, byteCount uint32) ([]byte, bool) {
	return m.del.Read(offset, byteCount)
}

func (m *TraceMemoryInstance) WriteByte(offset uint32, v byte) bool {
	m.changes = append(m.changes, api.MemoryChangeInfo{Offset: offset, Value: []byte{v}})
	return m.del.WriteByte(offset, v)
}

func (m *TraceMemoryInstance) WriteUint16Le(offset uint32, v uint16) bool {
	buffer := make([]byte, 2)
	binary.LittleEndian.PutUint16(buffer, v)
	m.changes = append(m.changes, api.MemoryChangeInfo{Offset: offset, Value: buffer})
	return m.del.WriteUint16Le(offset, v)
}

func (m *TraceMemoryInstance) WriteUint32Le(offset, v uint32) bool {
	buffer := make([]byte, 2)
	binary.LittleEndian.PutUint32(buffer, v)
	m.changes = append(m.changes, api.MemoryChangeInfo{Offset: offset, Value: buffer})
	return m.del.WriteUint32Le(offset, v)
}

func (m *TraceMemoryInstance) WriteFloat32Le(offset uint32, v float32) bool {
	panic("not supported yet")
	return m.del.WriteFloat32Le(offset, v)
}

func (m *TraceMemoryInstance) WriteUint64Le(offset uint32, v uint64) bool {
	buffer := make([]byte, 2)
	binary.LittleEndian.PutUint64(buffer, v)
	m.changes = append(m.changes, api.MemoryChangeInfo{Offset: offset, Value: buffer})
	return m.del.WriteUint64Le(offset, v)
}

func (m *TraceMemoryInstance) WriteFloat64Le(offset uint32, v float64) bool {
	panic("not supported yet")
	return m.del.WriteFloat64Le(offset, v)
}

func (m *TraceMemoryInstance) Write(offset uint32, v []byte) bool {
	m.changes = append(m.changes, api.MemoryChangeInfo{Offset: offset, Value: v})
	return m.del.Write(offset, v)
}

func (m *TraceMemoryInstance) WriteString(offset uint32, v string) bool {
	m.changes = append(m.changes, api.MemoryChangeInfo{Offset: offset, Value: []byte(v)})
	return m.del.WriteString(offset, v)
}

func (m *TraceMemoryInstance) RawBuffer() []byte {
	return m.del.RawBuffer()
}

func (m *TraceMemoryInstance) PageSize() (result uint32) {
	return m.del.PageSize()
}

func (m *TraceMemoryInstance) PeekMemoryChanges(drop bool) []api.MemoryChangeInfo {
	result := m.changes
	if drop {
		m.changes = nil
	}
	return result
}

func (m *TraceMemoryInstance) GetMin() uint32 {
	return m.del.GetMin()
}

func (m *TraceMemoryInstance) GetMax() uint32 {
	return m.del.GetMax()
}
