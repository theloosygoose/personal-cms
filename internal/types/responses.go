package types

type ServerResponse struct {
	Message string
	Code    int
}

type ErrorResponse struct {
    //Error code
    Code int

    //Error Message
    Message string
}
