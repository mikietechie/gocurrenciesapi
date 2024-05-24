package responses

type Error400Response struct {
	Error interface{} `json:"error"`
}
