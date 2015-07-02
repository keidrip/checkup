package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["storeCurrentUrl"] = storeCurrentUrl
}

func storeCurrentUrl(_url interface{}) {

	url := _url.(string)

	fmt.Print("[storeCurrentUrl]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}