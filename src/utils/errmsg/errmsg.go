package errmsg

type ErrCode int

const (
	OK    = 200
	ERROR = 400
)

var ErrMsg = map[ErrCode]string{
	OK:    "OK",
	ERROR: "ERROR",
}
