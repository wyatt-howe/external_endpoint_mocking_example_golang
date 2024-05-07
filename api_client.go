package main

import "fmt"

type Options struct {
	APIOptions error
	AppID      string
}

type Client struct {
	options Options
}

func (objectself *Client) accessExternalEndpoint() (string, error) {
	var data, err = fmt.Println("Phoned home!", objectself)
	return fmt.Sprintf("{'externaldata':'%x'}", data), err
}
