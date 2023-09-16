package shared

type IDGenerator interface {
	Generate() string
}
