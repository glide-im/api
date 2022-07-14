package tm

type Writer struct {
	Html string
}

func (w *Writer) Write(b []byte) (n int, err error) {
	w.Html += string(b)
	return len(b), nil
}

type Captcha struct {
	Code string
}
