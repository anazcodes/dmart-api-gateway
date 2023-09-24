package util

type response struct {
	Status int
	Msg    string
	Data   interface{}
	Err    interface{}
}

func Response(status int, msg string, data, err interface{}) response {

	return response{
		Status: status,
		Msg:    msg,
		Data:   data,
		Err:    err,
	}
}
