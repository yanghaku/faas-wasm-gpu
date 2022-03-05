package controller

import (
	"github.com/openfaas/faas-provider/types"
	"log"
)

type Controller struct {
	SecretsClient SecretsClient
}

func NewController() *Controller {
	log.Printf("Controller instance created")
	return &Controller{
		SecretsClient: NewDefaultMemorySecretsClient(),
	}
}

func (c *Controller) DeployFunc(funcDeployment *types.FunctionDeployment) error {
	log.Printf("deploy function: %+v\n", funcDeployment)
	return nil
}

func (c *Controller) UpdateFunc(funcDeployment *types.FunctionDeployment) error {
	log.Printf("update function: %+v\n", funcDeployment)
	return nil
}

func (c *Controller) DeleteFunc(funcName string) error {
	log.Printf("delete function: %s\n", funcName)
	return nil
}

func (c *Controller) FuncStateList() ([]types.FunctionStatus, error) {
	log.Printf("query state list\n")
	return []types.FunctionStatus{}, nil
}

func (c *Controller) FuncState(funcName string) (*types.FunctionStatus, error) {
	log.Printf("query function state for %s\n", funcName)
	return &types.FunctionStatus{}, nil
}

func (c *Controller) ScaleFunc(funcName string, replicas int32) error {
	log.Printf("scale the deveped function %s to %d\n", funcName, replicas)
	return nil
}
