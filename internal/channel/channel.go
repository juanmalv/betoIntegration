package channel

import (
	"beto_integration/internal/transport"
	"fmt"
)

type Channel struct {
	ChannelResponse string
}

func (c *Channel)Send(request string, transport transport.Transport, filters ...string) (string,error) {

	if c.IsMock(){
		return c.ChannelResponse, nil
	}

	fmt.Println("Sending request...")
	fmt.Println("Sent!")

	return "Approved", nil
}

func (c *Channel) IsMock()bool{
	return c.ChannelResponse != ""
}

func BuildFilters() string{
	fmt.Println("Hide this!!")

	return "asdf"
}
