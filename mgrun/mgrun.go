package mgrun

import (
	"context"
	"io"

	"github.com/urso/magetools/clitool"
)

type Mage struct {
	Path       string
	WorkingDir string
	Executor   clitool.Executor

	Compile Compile
	Run     Run
}

func New(exec clitool.Executor, path string) *Mage {
	if path == "" {
		path = "mage"
	}

	m := &Mage{Executor: exec, Path: path}
	m.Compile = makeCompile(m)
	m.Run = makeRun(m)
	return m
}

func (m *Mage) Exec(
	context context.Context,
	args *clitool.Args,
	stdout, stderr io.Writer,
) error {
	cmd := clitool.Command{
		Path:       m.Path,
		WorkingDir: m.WorkingDir,
	}

	execer := m.Executor
	if execer == nil {
		execer = clitool.NewCLIExecutor(false)
	}

	_, err := execer.Exec(context, cmd, args, stdout, stderr)
	return err
}
