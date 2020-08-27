package operations

//type BuildRequestFunc func() error
//type BuildTransportFunc func() error
//type BuildSendFunc func() error
//type BuildResponseFunc func() error

//Esta interfaz engloba los pasos básicos de cualquier operacion o sub-operacion dentro de una integracion.
type IntegrationFlow interface {
	BuildRequest() error
	BuildTransport() error
	SendRequest() error
	BuildResponse() error

	//BuildRequest() error
	//SetBuildRequest(f BuildRequestFunc)
	//
	//BuildTransport() error
	//SetBuildTransport(f BuildTransportFunc)
	//
	//Send() error
	//SetSend(f BuildSendFunc)
	//
	//BuildResponse() error
	//SetBuildResponse(f BuildResponseFunc)
}
