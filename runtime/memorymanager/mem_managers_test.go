package memorymanager_test

import (
	"lang/ast/types"
	"lang/pkg"
	"lang/runtime"
	memeorymanager "lang/runtime/memorymanager"
	"testing"
)

// form map_mem_manager.go
func TestMapMemManager(t *testing.T) {
	mmm := memeorymanager.MapMemManager{Vars: pkg.Stack[map[string]runtime.Type]{}}

	mmm.PushScope()

	mmm.Set("var1", &types.Number{Val: 5})
	mmm.Set("var2", &types.Boolean{Val: false})

	v1 := mmm.Get("var1")
	n, ok := v1.(*types.Number)
	if !ok {
		t.Error("Get Map Memory Manager Fail 1")
	}

	if n.Val != 5 {
		t.Error("Invalid Map Memory Value 1")
	}

	v2 := mmm.Get("var2")
	b, ok := v2.(*types.Boolean)
	if !ok {
		t.Error("Get Map Memory Manager Fail 2")
	}

	if b.Val != false {
		t.Error("Invalid Map Memory Value 2")
	}

	mmm.PopScope()

	v1 = mmm.Get("var1")
	if v1 != nil {
		t.Error("Invalid Memory")
	}
}

func TestMapMemManagerMultipleScopes(t *testing.T) {
	mmm := memeorymanager.MapMemManager{Vars: pkg.Stack[map[string]runtime.Type]{}}
	mmm.PushScope()

	mmm.Set("var0", &types.Number{Val: 5})
	mmm.PushScope()
	mmm.Set("var1", &types.Number{Val: 10})
	mmm.PushScope()
	mmm.Set("var2", &types.Number{Val: 10})

	v := mmm.Get("var0")
	if v == nil {
		t.Error("Variable cant be nil")
	}

	num, ok := v.(*types.Number)
	if !ok {
		t.Error("variable isnt a number")
	}

	if num.Val != 5 {
		t.Error("invalid number")
	}

}
