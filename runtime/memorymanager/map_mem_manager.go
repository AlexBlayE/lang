package memorymanager

import (
	"lang/pkg"
	"lang/runtime"
)

type MapMemManager struct {
	Vars pkg.Stack[map[string]runtime.Type]
}

func (mmm *MapMemManager) Get(key string) runtime.Type {
	length := mmm.Vars.Len()

	for i := length - 1; i >= 0; i-- {
		if val, ok := mmm.Vars.Index(i)[key]; ok {
			return val
		}
	}

	return nil
}

func (mmm *MapMemManager) Set(key string, val runtime.Type) {
	last := mmm.Vars.Len() - 1
	mmm.Vars.Index(last)[key] = val
}

func (mmm *MapMemManager) PushScope() {
	mmm.Vars.Push(make(map[string]runtime.Type))
}

func (mmm *MapMemManager) PopScope() {
	mmm.Vars.Pop()
}
