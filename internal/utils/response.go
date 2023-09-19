package util

type response struct {
	status int
	msg    string
	data   interface{}
	err    interface{}
}

func Response(status int, msg string, data, err interface{}) response {

	return response{
		status: status,
		msg:    msg,
		data:   data,
		err:    err,
	}
}
