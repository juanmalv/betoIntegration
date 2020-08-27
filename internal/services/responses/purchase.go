package responses

type PurchaseResponse struct {
	amount int
}

func (p *PurchaseResponse) BuildResponse() error {
	p.amount = 1
	return nil
}

func (p *PurchaseResponse) IsMock() bool {
	return p != nil
}