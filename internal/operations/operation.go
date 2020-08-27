package operations

import (
	"fmt"
)

type Operation interface {
	execute(operation OperationFlow) //Deberia recibir trx y ctx
}

type OperationFlow interface {
	BuildRequest() error
	BuildTransport() error
	SendRequest() error
	BuildResponse() error
}

func execute(operation OperationFlow) error{
	err := operation.BuildRequest()
	if err != nil{
		return fmt.Errorf("Error building request!")
	}

	err = operation.BuildTransport()
	if err != nil{
		fmt.Errorf("Error building transport!")
	}

	err = operation.SendRequest()
	if err != nil{
		fmt.Errorf("Error in channel service")
	}

	err = operation.BuildResponse()
	if err != nil{
		fmt.Errorf("Error building response")
	}

	return nil
}