package main

import (
	"errors"
	"fmt"
)

func useClient(client *Client) string {
	var result, err = client.accessExternalEndpoint()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("{'finaloutput':%s}", result)
}

func main() {
	var cl = &Client{options: Options{APIOptions: errors.New(""), AppID: "app123"}}

	fmt.Println(useClient(cl))
}
