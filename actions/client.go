package actions

import "fmt"

// Client an action for a server
type Client struct {
	Name string
}

// Do Action interface implementation
func (c *Client) Do() {
	fmt.Println("Client", c.Name, "doing action.")
}
