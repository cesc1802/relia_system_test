package jwt

type option struct {
	SecretKey []byte
}

type OptionFunc func(o *option)

func WithSecretKey(key []byte) OptionFunc {
	return func(o *option) {
		o.SecretKey = key
	}
}
