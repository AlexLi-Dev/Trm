// Author: Lutong.li
package changedatastyle

import (
	"encoding/json"
)

func StructToMap(s interface{}) (map[string]string, error) {
	j, _ := json.Marshal(s)

	m := make(map[string]string)
	err := json.Unmarshal(j, &m)
	if err != nil {
		return m, err
	}
	return m, err

}
