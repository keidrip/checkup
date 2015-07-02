package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["assertPageSource"] = assertPageSource
}

func assertPageSource(_url interface{}) {

	url := _url.(string)

	fmt.Print("[assertPageSource]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}