package interfaces

type FoundryIDGenerator interface {
	GetID() [16]byte
}
