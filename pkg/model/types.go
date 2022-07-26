package model

import "time"

type NFTTypeId int16

func (t NFTTypeId) Validate() error {
	if t < lastNFTTypeId {
		return nil
	}
	return errUnknownNFTTypeId
}

type User struct {
	UserId string
	Name   string
	Mail   string
}

type WrappedNFT struct {
	Id        string
	OwnerId   string
	CreatedAt time.Time
	NFT
}

func ParseNFTContent(t NFTTypeId, bytes []byte) (NFT, error) {
	switch t {
	case NFTTypeIdMusician:
		return NFTOfTypeFromBytes[NFTMusician](bytes)
	case NFTTypeIdCoin:
		return NFTOfTypeFromBytes[NFTCoin](bytes)
	default:
		return nil, errUnknownNFTTypeId
	}
}
