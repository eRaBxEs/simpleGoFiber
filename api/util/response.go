package util

// Response A map[string]interface{} type for responding to json clients
type Response struct {
	Code   int                    `json:"code"`
	Store  map[string]interface{} `json:"store"`
	Errors ErrMsg                 `json:"errors"`
}

// ErrMsg holds error messages
type ErrMsg map[string]string

// Set stores an a named value in Data
func (s *Response) Set(name string, value interface{}) {
	if s.Store == nil {
		s.Store = map[string]interface{}{}
	}

	s.Store[name] = value
}

// ErrMsg helper for adding error messages
func (s *Response) ErrMsg(name, value string) {
	// create error field if it doesn't exist

	if s.Errors == nil {
		s.Errors = ErrMsg{}
	}

	s.Errors[name] = value
}
