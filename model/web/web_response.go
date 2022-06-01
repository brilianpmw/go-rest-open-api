package web

type WebResponse struct {
	Code    int
	Status  string
	Data    interface{}
	Message string
}
