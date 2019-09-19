package mgrun

import (
	"context"
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/urso/magetools/clitool"
)

type Compile func(context context.Context, opts ...clitool.ArgOpt) error

type mgCompile struct {
	m *Mage
}

func makeCompile(m *Mage) Compile {
	mc := &mgCompile{m}
	return mc.Do
}

func (Compile) TargetDir(path string) clitool.ArgOpt {
	return clitool.ExtraIf("todir", path)
}

func (Compile) TargetBin(name string) clitool.ArgOpt {
	return clitool.ExtraIf("name", name)
}

func (Compile) OS(name string) clitool.ArgOpt {
	return clitool.Combine(
		clitool.ExtraIf("goos", name),
		clitool.FlagIf("-goos", name),
	)
}

func (Compile) Arch(name string) clitool.ArgOpt {
	return clitool.Combine(
		clitool.ExtraIf("goarch", name),
		clitool.FlagIf("-goarch", name),
	)
}

func (mc *mgCompile) Do(context context.Context, opts ...clitool.ArgOpt) error {
	args := clitool.CreateArgs(opts...)

	goos := args.GetExtraDefault("goos", runtime.GOOS)
	goarch := args.GetExtraDefault("goarch", runtime.GOARCH)
	targetName := args.GetExtraDefault("name", fmt.Sprintf("mage-%v-%v", goos, goarch))
	to := path.Join(args.GetExtraDefault("todir", "."), targetName)

	args.Flags = append(args.Flags,
		clitool.CreateArgs(
			clitool.Flag("-f", ""),
			clitool.Flag("-compile", to)).Flags...)

	return mc.m.Exec(context, args, os.Stdout, os.Stderr)
}
