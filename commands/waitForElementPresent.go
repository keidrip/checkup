package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["waitForElementPresent"] = waitForElementPresent
}

func waitForElementPresent(_url interface{}) {

	url := _url.(string)

	fmt.Print("[waitForElementPresent]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}