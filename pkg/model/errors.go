package model

type ErrorHandler interface {
	HandleError(e error)
}

type ConstError string

func (c ConstError) Error() string {
	return string(c)
}

const (
	ErrFakeWritePool ConstError = "fake write pool supports no connection"
	ErrNilNotAllowed ConstError = "nil not allowed"
)

const (
	errUnknownNFTTypeId        ConstError = "undefined NFT type"
	errUnknownMusicianStatusId ConstError = "unknown musician state"
)
