package main

import "fmt"

func runCreateFnFunc(gcnf *globalConf, cnf *createFnConf) func() error {
	return func() error {
		fmt.Println("create")
		return nil
	}
}

func runUpdateFnFunc(gcnf *globalConf, cnf *updateFnConf) func() error {
	return func() error {
		fmt.Println("update")
		return nil
	}
}
