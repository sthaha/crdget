package controller

import (
	"github.com/sthaha/k8s-crd-get/pkg/controller/customresourcedefinition"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, customresourcedefinition.Add)
}
