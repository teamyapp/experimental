package idgen

type IDGenerator interface {
	NextID() string
}
