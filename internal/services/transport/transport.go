package transport


//Si el transport es igual para todas las operaciones, unificar


type Transport struct{
	header string
}

func (t *Transport) BuildTransport() error {
	t.header = "xml"
	return nil
}

func (t *Transport) IsMock() bool {
	return t != nil
}
