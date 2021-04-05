package serializers

import "encoding/json"

// Resp : struct
type Resp struct {
	Result interface{}
	Error  interface{}
}

func ConvertSerializer(data interface{}, output interface{}) error {
	marshal, err := json.Marshal(data)
	if err != nil {
		return err
	}
	json.Unmarshal(marshal, &output)

	return nil
}
