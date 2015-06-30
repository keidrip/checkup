package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["open"] = open
}

func open(_url interface{}) {

	url := _url.(string)

	fmt.Println("[open]: " + url)

	err := WD.Get(url)
	errors.Fatal(err)

}