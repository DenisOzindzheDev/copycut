package bind

import (
	"encoding/json"
	"io/ioutil"

	"github.com/DenisOzindzheDev/copycut/internal/models"
)

func Store(k, v string) {
	// fmt.Printf("Last Data in cli[board : %s\n", string(data))
	data, _ := ioutil.ReadFile("test.json")
	bindings := models.Data{}
	parsed := json.Unmarshal(data, &bindings)
	_ = parsed

	bind := models.Data{
		Set: []models.Bindings{{
			Key:   k,
			Value: v,
		},
		},
	}
	for _, v := range bindings.Set {

		bind.Set = append(bind.Set, models.Bindings{Key: v.Key, Value: v.Value})
	}
	file, _ := json.MarshalIndent(bind, " ", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)
}
