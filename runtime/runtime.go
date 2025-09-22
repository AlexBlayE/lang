package runtime

import "errors"

type MemManager interface {
	Get(string) Type
	Set(string, Type)
	PushScope()
	PopScope()
}

type Type interface {
	Debug()
}

type Statement interface {
	Visit(StmVisitor) error
}

type Expression interface {
	Visit(ExprVisitor) Type
}

type StmVisitor interface {
	Accept(Statement) error
}

type ExprVisitor interface {
	Accept(Expression) Type
}

type StmDispatcher func(Statement, MemManager, ExprDispatcher) StmVisitor

type ExprDispatcher func(Expression, MemManager, StmDispatcher) ExprVisitor

type Runtime struct {
	StmDispatcher
	ExprDispatcher
	MemManager
}

func (r *Runtime) RunProgram(program []Statement) error {
	r.PushScope()
	defer r.PopScope()

	for _, stm := range program {
		v := r.StmDispatcher(stm, r.MemManager, r.ExprDispatcher)
		if v == nil {
			return errors.New("need a valid dispatcher")
		}

		err := stm.Visit(v)
		if err != nil {
			return err
		}
	}

	return nil
}
