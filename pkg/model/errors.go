package model

type ErrorHandler interface {
	HandleError(e error)
}

type ConstError string

func (c ConstError) Error() string {
	return string(c)
}

//psql store related errors
const (
	ErrFakeWritePool ConstError = "fake write pool supports no connection"
	ErrNilNotAllowed ConstError = "nil not allowed"
)

//type validations errors
const (
	errUnknownNFTTypeId        ConstError = "undefined NFT type"
	errUnknownMusicianStatusId ConstError = "unknown musician state"
)

//errors for psql mapping
const (
	ErrVersionNotSupported    ConstError = "version not supported"
	ErrTimestamptzNotPresent  ConstError = "timestamptz not present"
	ErrJsonNotPresent         ConstError = "json not present"
	ErrUuidNotPresent         ConstError = "uuid not present"
	ErrInt2NotPresent         ConstError = "int2 not present"
	ErrInt2NotCastableToUint8 ConstError = "int2 not castable to uint8"
)
