package down

type Request struct {
	Method  string
	URL     string
	Header  map[string]string
	Content []byte
}
