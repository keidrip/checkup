package commands

import (
	"fmt"
	"unidriver/errors"
)

func init() {
	CommandList["storePageSource"] = storePageSource
}

func storePageSource(_url interface{}) {

	url := _url.(string)

	fmt.Print("[storePageSource]: " + url + " ")
	err := WD.Get(url)
	errors.Fatal(err)
	fmt.Print("SUCCESS")
	fmt.Println("")

}