package intrf

type Ierror interface {
	Checkerror(w Iwriteresponse, mess string, err error)
}
