package providers

type UuidProvider interface {
	Generate() string
}
