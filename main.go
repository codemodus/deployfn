package main // import "github.com/codemodus/deployfn"

import (
	"os"
	"path"

	"github.com/codemodus/clip"
)

const (
	aws = "aws"
	gcp = "gcp"
)

func main() {
	name := path.Base(os.Args[0])

	if err := run(name); err != nil {
		logError(name, err)
		os.Exit(1)
	}
}

func run(name string) error {
	var (
		globalCnf      = newGlobalConf(name)
		awsCreateFnCnf = newCreateFnConf(aws)
		awsUpdateFnCnf = newUpdateFnConf(aws)
		gcpCreateFnCnf = newCreateFnConf(gcp)
		gcpUpdateFnCnf = newUpdateFnConf(gcp)
	)

	cs := clip.NewCommandSet(
		clip.NewCommandNamespace(aws, clip.NewCommandSet(
			clip.NewCommand(
				awsCreateFnCnf.fs, runCreateFnFunc(globalCnf, awsCreateFnCnf), nil,
			),
			clip.NewCommand(
				awsUpdateFnCnf.fs, runUpdateFnFunc(globalCnf, awsUpdateFnCnf), nil,
			),
		)),

		clip.NewCommandNamespace(gcp, clip.NewCommandSet(
			clip.NewCommand(
				gcpCreateFnCnf.fs, runCreateFnFunc(globalCnf, gcpCreateFnCnf), nil,
			),
			clip.NewCommand(
				gcpUpdateFnCnf.fs, runUpdateFnFunc(globalCnf, gcpUpdateFnCnf), nil,
			),
		)),
	)

	app := clip.New(name, globalCnf.fs, cs)

	if err := app.Parse(os.Args); err != nil {
		return app.UsageLongHelp(err)
	}

	return app.Run()
}
