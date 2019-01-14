package main

import (
	"time"

	"github.com/codemodus/clip"
)

type globalConf struct {
	fs      *clip.FlagSet
	verbose bool
}

func newGlobalConf(name string) *globalConf {
	c := globalConf{
		fs: clip.NewFlagSet(name),
	}

	c.fs.BoolVar(&c.verbose, "v", c.verbose, "enable logging")

	return &c
}

type createFnConf struct {
	fs     *clip.FlagSet
	name   string
	region string
	role   string
}

func newCreateFnConf(vendor string) *createFnConf {
	c := createFnConf{
		fs: clip.NewFlagSet("create"),
	}

	switch vendor {
	case aws:
		c.fs.StringVar(&c.name, "name", c.name, "function name")
		c.fs.StringVar(&c.region, "region", c.region, "target region")
		c.fs.StringVar(&c.role, "role", c.role, "role arn")
	case gcp:
	default:
		panic("mistakes were made")
	}

	return &c
}

type updateFnConf struct {
	fs      *clip.FlagSet
	name    string
	timeout time.Duration
}

func newUpdateFnConf(vendor string) *updateFnConf {
	c := updateFnConf{
		fs: clip.NewFlagSet("update"),
	}

	switch vendor {
	case aws:
		c.fs.StringVar(&c.name, "name", c.name, "function name")
		c.fs.DurationVar(&c.timeout, "timeout", c.timeout, "timeout duration")
	case gcp:
	default:
		panic("mistakes were made")
	}

	return &c
}
