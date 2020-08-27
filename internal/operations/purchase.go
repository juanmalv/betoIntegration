package operations

import (
	"github.com/juanmalv/betoIntegration/internal/services/channel"
	"github.com/juanmalv/betoIntegration/internal/services/requests"
	"github.com/juanmalv/betoIntegration/internal/services/responses"
	"github.com/juanmalv/betoIntegration/internal/services/transport"
)

type PurchaseOperation struct {
	request        requests.PurchaseRequest
	transport      transport.Transport
	channelService channel.Channel
	response       responses.PurchaseResponse
}

func (p *PurchaseOperation) BuildRequest () error {
	if p.request.IsMock(){
		return nil
	}

	return p.request.BuildRequest()
}

func (p *PurchaseOperation) BuildTransport() error {
	if p.transport.IsMock(){
		return nil
	}

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

	return nil
}

func (p *PurchaseOperation) BuildResponse() error{
	if p.response.IsMock(){
		return nil
	}

	return p.BuildResponse()
}

func Purchase() interface{}{
	//Si queremos testear de acá para abajo, podemos llamar a execute pasandole un objeto con lo que queremos
	//mockear ya cargado
	var purchase PurchaseOperation
	return execute(&purchase)
}


