package valueobject

type Username string

func (this *Username) GetValue() string {
	return string(*this)
}
