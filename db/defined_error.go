package db

type ErrStokBukuHabis struct{ Msg string }

func (e ErrStokBukuHabis) Error() string {
	return e.Msg
}

type ErrStokBukuKurang struct{ Msg string }

func (e ErrStokBukuKurang) Error() string {
	return e.Msg
}

type ErrJenisPesananDuplikat struct{ Msg string }

func (e ErrJenisPesananDuplikat) Error() string {
	return e.Msg
}

type ErrIDBukuTidakTerdaftar struct{ Msg string }

func (e ErrIDBukuTidakTerdaftar) Error() string {
	return e.Msg
}

type ErrUpdateNolData struct{ Msg string }

func (e ErrUpdateNolData) Error() string {
	return e.Msg
}

type ErrUpdateMultipleData struct{ Msg string }

func (e ErrUpdateMultipleData) Error() string {
	return e.Msg
}

type ErrNegativeNumber struct{ Msg string }

func (e ErrNegativeNumber) Error() string {
	return e.Msg
}

type ErrFieldTableNotExisted struct{ Msg string }

func (e ErrFieldTableNotExisted) Error() string {
	return e.Msg
}

type ErrStatusNotAcceptable struct{ Msg string }

func (e ErrStatusNotAcceptable) Error() string {
	return e.Msg
}

type ErrDBTypeNotImplemented struct{ Msg string }

func (e ErrDBTypeNotImplemented) Error() string {
	return e.Msg
}
