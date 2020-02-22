package engine

import (
	"errors"
)

func NewMemoryAccess() *memoryAccess {
	return &memoryAccess{
		cells: []int{0},
	}
}

type MemoryAccess interface {
	GetCellValue() int
	SetCellValue(x int)
	IncrementPointer()
	DecrementPointer() error
	IncrementCellValue()
	DecrementCellValue()
}

type memoryAccess struct {
	cells      []int
	currentIdx int
}

func (s *memoryAccess) GetCellValue() int {
	return s.cells[s.currentIdx]
}

func (s *memoryAccess) SetCellValue(x int) {
	s.cells[s.currentIdx] = x
}

func (s *memoryAccess) IncrementPointer() {
	s.currentIdx++
	if len(s.cells) <= s.currentIdx {
		s.cells = append(s.cells, 0)
	}
}

func (s *memoryAccess) DecrementPointer() error {
	if s.currentIdx == 0 {
		return errors.New("cannot decrement zero pointer")
	}

	s.currentIdx--
	return nil
}

func (s *memoryAccess) IncrementCellValue() {
	s.cells[s.currentIdx]++
}

func (s *memoryAccess) DecrementCellValue() {
	s.cells[s.currentIdx]--
}
