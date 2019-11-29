package response

type Response struct {
	AccessType string `json:"charset"`
	Data       string `json:"data"`
	TransCode  string `json:"transCode"`
	MerId      string `json:"merId"`
}

func (s *Response) SetData(data string) *Response {
	s.Data = data
	return s
}

func (s *Response) SetKeyValue(key string, value string) *Response {
	switch key {
	case "accessType":
		s.AccessType = value
	case "transCode":
		s.TransCode = value
	case "data":
		s.Data = value
	case "merId":
		s.MerId = value
	}
	return s
}
