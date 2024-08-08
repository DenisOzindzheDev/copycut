package clipboard

import (
	"encoding/json"
	"io/ioutil"

	"golang.design/x/clipboard"
)

func Store() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	data := clipboard.Read(clipboard.FmtText)
	// fmt.Printf("Last Data in cli[board : %s\n", string(data))
	type item struct {
		ID   int    `json:"id"`
		Data string `json:"data"`
	}

	last := item{
		ID:   1,
		Data: string(data),
	}
	file, _ := json.MarshalIndent(last, " ", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)
}
