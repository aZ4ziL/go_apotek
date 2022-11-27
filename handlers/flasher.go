package handlers

var flasher Flasher

type Flasher struct {
	Message string
	Type    string
}

func (f *Flasher) Set(_type, msg string) {
	f.Message = msg
	f.Type = _type
}

func (f *Flasher) Del() {
	f.Message = ""
	f.Type = ""
}
