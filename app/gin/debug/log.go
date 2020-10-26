package debug

import (
	"encoding/json"
	"fmt"
)

// LogStruct "github.com/kantaroso/game-information/debug"
func LogStruct(v interface{}) {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(string(jsonBytes))
}
