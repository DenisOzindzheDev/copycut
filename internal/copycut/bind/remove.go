package bind

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/DenisOzindzheDev/copycut/internal/models"
	"github.com/nexidian/gocliselect"
	"golang.design/x/clipboard"
)

func Delete(k string) {
	menu := gocliselect.NewMenu("Select binding")
	data, _ := ioutil.ReadFile("test.json")
	bindings := models.Data{}
	parsed := json.Unmarshal(data, &bindings)
	_ = parsed
	saved := []models.Bindings{}
	for _, v := range bindings.Set {
		menu.AddItem(v.Key, v.Value)
		saved = append(saved, models.Bindings{Key: v.Key, Value: v.Value})
	}
	choice := menu.Display()
	fmt.Printf("Choice: %s\n", choice)
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	for k, binding := range saved {
		if binding.Value == choice {
			saved = append(saved[:k], saved[k+1:]...)
		}
	}
	file, _ := json.MarshalIndent(models.Data{Set: saved}, " ", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)
}
