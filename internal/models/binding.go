package models

type Data struct {
	Set []Bindings `json:"bindings,omitempty"`
}
type Bindings struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

/*
{
	"bindings":[
		{
			"key": "string key",
			"value": "string value"
		},
		{
			"key": "string key",
			"value": "string value"
		}
	]
}
*/
