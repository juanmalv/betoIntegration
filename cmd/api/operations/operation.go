package operations

type BuildRequestFunc func() error
type BuildTransportFunc func() error
type BuildSendFunc func() error

//Esta interfaz debe estar adaptada a las necesidades de la Integracion y debe ser expandida en casos comunes a la misma
type Operation interface {
	BuildRequest() error
	SetBuildRequest(f BuildRequestFunc)

	BuildTransport() error
	SetBuildTransport(f BuildTransportFunc)

	SetSend(f BuildSendFunc)
	Send() error

	BuildResponse()
}
