package amiderrors

type ErrorConv struct {
	is_Stopped bool
	Err        error
}

func (ec *ErrorConv) Do(fn func() error) {
	if ec.Err != nil || ec.is_Stopped {
		return
	}
	ec.Err = fn()
}
func (ec *ErrorConv) DoMany(args ...func() error) {
	for _, fn := range args {
		ec.Err = fn()
		if ec.Err != nil || ec.is_Stopped {
			break
		}
	}
}
func (ec *ErrorConv) Stop() {
	ec.is_Stopped = true
}
