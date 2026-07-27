// Author: Lutong.li
package returndata

func NewReturnData() returnData {
	return returnData{}
}

type returnData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
