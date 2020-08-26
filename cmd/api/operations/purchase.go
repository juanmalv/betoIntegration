package operations

import (
	"beto_integration/internal/requests"
	"beto_integration/internal/transport"
)

type PurchaseOperation struct {
	buildRequestFunc   BuildRequestFunc
	buildTransportFunc BuildTransportFunc
	transport          transport.Transport
	buildSendFunc      BuildSendFunc
	response           responses.PurchaseResponse
}

func (p *PurchaseOperation) SetBuildRequest(f BuildRequestFunc) {
	p.BuildRequestFunc = f
}
func (p *PurchaseOperation) BuildRequest() error {
	return p.BuildRequestFunc()
}

func (p *PurchaseOperation) SetBuildTransport(f BuildTransportFunc) {
	p.BuildTransportFunc = f
}
func (p *PurchaseOperation) BuildTransport() error {
	return p.BuildTransportFunc()
}

func (p *PurchaseOperation) SetSend(f BuildSendFunc) {
	p.buildSendFunc = f
}
func (p *PurchaseOperation) Send() error {
	return p.buildSendFunc()
}

func (p *PurchaseOperation) BuildResponse() {
	p.response.BuildResponse()
}
