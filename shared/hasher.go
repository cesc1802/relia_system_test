package shared

type Hasher interface {
	Hash(data string) string
}
