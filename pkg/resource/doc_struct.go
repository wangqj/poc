package resource

type KVBody struct {
	Labels    map[string]string `json:"labels"`
	ValueType string            `json:"valueType"`
	Value     string            `json:"value"`
}

type ErrorMsg struct {
	Msg string `json:"msg"`
}
