package serializers

import (
	"encoding/json"
	"fmt"
)

// Resp : struct
type Resp struct {
	Result interface{} `json:"result"`
	Error  interface{} `json:"error"`
}

func ConvertSerializer(data interface{}, output interface{}) error {
	marshal, err := json.Marshal(data)
	if err != nil {
		return err
	}
	json.Unmarshal(marshal, &output)

	fmt.Println(output)

	return nil
}
