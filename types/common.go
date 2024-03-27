package types

type ErrorMsg struct {
	Level   string `json:"level,omitempty"`
	Route   string `json:"route,omitempty"`
	Trigger string `json:"trigger,omitempty"`
	Desc    string `json:"description,omitempty"`
	ErrCode string `json:"code,omitempty"`
}
type PingResponse struct {
	Network  string `json:"network"`
	Service  string `json:"service"`
	Status   string `json:"status"`
	Version  string `json:"version"`
	Timezone string `json:"timezone"`
}

/*
	type FailedMsg struct {
		Success bool `json:"success"`
		Desc string `json:"description"`
	}

	type SuccessMsg struct {
		Success bool `json:"success"`
		Data interface{} `json:"data"`
	}
*/
type ResponseMsg struct {
	Level   string      `json:"level"`
	Route   string      `json:"route"`
	Trigger string      `json:"trigger"`
	Desc    string      `json:"description"`
	Data    interface{} `json:data`
}
