package mgrun

import (
	"context"
	"os"

	"github.com/urso/magetools/clitool"
)

type Run func(context context.Context, opts ...clitool.ArgOpt) error

type mgRun struct {
	m *Mage
}

func makeRun(m *Mage) Run {
	mr := &mgRun{m}
	return mr.Do
}

func (Run) Target(name string) clitool.ArgOpt { return clitool.Positional(name) }
func (Run) Verbose(b bool) clitool.ArgOpt     { return clitool.BoolFlag("-v", b) }

func (mr *mgRun) Do(context context.Context, opts ...clitool.ArgOpt) error {
	args := clitool.CreateArgs(opts...)
	return mr.m.Exec(context, args, os.Stdout, os.Stderr)
}
