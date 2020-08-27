package operations

import (
	"beto_integration/internal/channel"
	"beto_integration/internal/requests"
	"beto_integration/internal/responses"
	"beto_integration/internal/transport"
)

type PurchaseOperation struct {
	request requests.PurchaseRequest
	transport transport.Transport
	channelService channel.Channel
	response responses.PurchaseResponse
}

func (p *PurchaseOperation) BuildRequest () error {
	if p.request.IsMock(){
		return nil
	}

	return p.request.BuildRequest()
}

func (p *PurchaseOperation) BuildTransport() error {


	return p.transport.BuildTransport()
}

func (p *PurchaseOperation) SendRequest() error {
	req, err := p.request.ToXML()

	if err != nil{
		return err
	}


	//Si en Send encontramos el campo ya completado (mock), devolvemos la response, evitando ir al provider
	p.channelService.ChannelResponse, err =  p.channelService.Send(req, p.transport, channel.BuildFilters())

	if err != nil{
		return err
	}

	p.channelService.ChannelResponse = channelResponse

	return nil
}

func (p *PurchaseOperation) BuildResponse() error{
	return p.BuildResponse()
}

func Purchase() interface{}{
	//Si queremos testear de acá para abajo, podemos llamar a execute pasandole un objeto con lo que queremos
	//mockear ya cargado
	var purchase PurchaseOperation
	return execute(&purchase)
}



//type PurchaseOperation struct {
//	buildRequestFunc   BuildRequestFunc
//	buildTransportFunc BuildTransportFunc
//	buildSendFunc      BuildSendFunc
//	buildResponseFunc  BuildResponseFunc
//}
//
//func (p *PurchaseOperation) SetBuildRequest(f BuildRequestFunc) {
//	p.buildRequestFunc = f
//}
//func (p *PurchaseOperation) BuildRequest() error {
//	return p.buildRequestFunc()
//}
//
//func (p *PurchaseOperation) SetBuildTransport(f BuildTransportFunc) {
//	p.buildTransportFunc = f
//}
//func (p *PurchaseOperation) BuildTransport() error {
//	return p.buildTransportFunc()
//}
//
//func (p *PurchaseOperation) SetSend(f BuildSendFunc) {
//	p.buildSendFunc = f
//}
//func (p *PurchaseOperation) Send() error {
//	return p.buildSendFunc()
//}
//
//func (p *PurchaseOperation) BuildResponse() {
//	p.buildResponseFunc()
//}


