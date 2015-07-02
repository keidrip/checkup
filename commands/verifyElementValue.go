package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["verifyElementValue"] = verifyElementValue
}

func verifyElementValue(_url interface{}) {

	url := _url.(string)

	fmt.Print("[verifyElementValue]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}