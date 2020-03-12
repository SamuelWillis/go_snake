package helpers

import(
	"encoding/json"
	"log"
)

// Dump object to the logs
func Dump(obj interface{}) {
	data, err := json.MarshalIndent(obj, "", "  ")
	if err == nil {
		log.Printf(string(data))
	}
}
