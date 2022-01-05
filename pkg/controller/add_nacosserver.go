package controller

import (
	"github.com/kubernetes/sample-controller/pkg/controller/nacosserver"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, nacosserver.Add)
}
