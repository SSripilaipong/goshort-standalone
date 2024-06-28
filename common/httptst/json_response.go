package httptst

import (
	"encoding/json"
	"fmt"
)

func JsonResponse(reader fmt.Stringer) map[string]any {
	m := make(map[string]any)
	raw := reader.String()
	fmt.Println("raw:", raw)
	_ = json.Unmarshal([]byte(raw), &m)
	return m
}
