package config

type ErrInvalidGinMode struct {
	value string
}

func (e ErrInvalidGinMode) Error() string {
	return "invalid gin mode [" + e.value + "]"
}
