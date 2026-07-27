// Author: Lutong.li
package changedatastyle

import (
	"encoding/json"
	"fmt"
)

func StructToMap(s interface{}) (map[string]string, error) {
	fmt.Printf("StructToMap %+v\n", s)
	j, _ := json.Marshal(s)

	m := make(map[string]string)
	err := json.Unmarshal(j, &m)
	if err != nil {
		return m, err
	}
	return m, err

}
