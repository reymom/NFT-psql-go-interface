package model

import "encoding/json"

type genericNFTs interface {
	NFTMusician | NFTCoin
}

func NFTOfTypeFromBytes[T genericNFTs](bytes []byte) (*T, error) {
	var nft T
	e := json.Unmarshal(bytes, &nft)
	if e != nil {
		return nil, e
	}
	return &nft, nil
}

//content for nft of type NFTTypeIdMusician

type MusicianStatusId uint8

func (t MusicianStatusId) Validate() error {
	if t < lastMusicianStatusId {
		return nil
	}
	return errUnknownMusicianStatusId
}

type NFTMusician struct {
	Name   string           `json:"name"`
	Status MusicianStatusId `json:"status"`
}

func (c *NFTMusician) GetNFTTypeId() NFTTypeId {
	return NFTTypeIdMusician
}

func (c *NFTMusician) ToJson() ([]byte, error) {
	return json.Marshal(c)
}

//content for nft of type NFTTypeIdCoin

type NFTCoin struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Value   uint   `json:"value"`
	Active  bool   `json:"active"`
}

func (c *NFTCoin) GetNFTTypeId() NFTTypeId {
	return NFTTypeIdMusician
}

func (c *NFTCoin) ToJson() ([]byte, error) {
	return json.Marshal(c)
}
