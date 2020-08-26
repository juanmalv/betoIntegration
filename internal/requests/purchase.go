package requests

type PurchaseRequest struct {
	amount int
}

func (p *PurchaseRequest)BuildRequest(){
	p.amount = 1
	return
}