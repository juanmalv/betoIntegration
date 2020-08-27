package operations

import "fmt"

type Operation interface {
	execute(operation IntegrationFlow) //Deberia recibir trx y ctx
}


func execute(operation IntegrationFlow) error{
	err := operation.BuildRequest()
	if err != nil{
		return fmt.Errorf("Error!")
	}

}