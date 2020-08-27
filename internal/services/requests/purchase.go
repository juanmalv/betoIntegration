package requests

import "encoding/xml"

type PurchaseRequest struct {
	amount int
}

func (p *PurchaseRequest)BuildRequest() error {

	p.amount = 1
	return nil
}

func (p *PurchaseRequest)IsMock() bool {
	return p != nil
}

func (p *PurchaseRequest) ToXML() (string, error) {
	req, err := xml.Marshal(p)

	if err != nil {
		return "", err
	}

	return string(req), nil
}