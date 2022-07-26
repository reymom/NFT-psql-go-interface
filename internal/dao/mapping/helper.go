package mapping

import (
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/reymom/NFT-psql-go-interface/pkg/model"
)

func checkVersion(expectedVersion uint8, actualVersion uint8) error {
	if expectedVersion != actualVersion {
		return model.ErrVersionNotSupported
	}
	return nil
}

func timeFromPgTimestamptz(pgT pgtype.Timestamptz, acceptNil bool) (*time.Time, error) {
	if pgT.Status != pgtype.Present {
		var e error
		if !acceptNil {
			e = model.ErrTimestamptzNotPresent
		}
		return nil, e
	}
	ret := pgT.Time.UTC()
	return &ret, nil
}

func pgJsonFromByte(b []byte) pgtype.JSON {
	return pgtype.JSON{
		Bytes:  b,
		Status: pgtype.Present,
	}
}

func jsonFromPgJson(j pgtype.JSON) ([]byte, error) {
	if j.Status != pgtype.Present {
		return nil, model.ErrJsonNotPresent
	}
	return j.Bytes, nil
}

func uuidFromPGUuid(pgUuid pgtype.UUID) (string, error) {
	if pgUuid.Status != pgtype.Present {
		return "", model.ErrUuidNotPresent
	}
	parsedUuid, e := uuid.FromBytes(pgUuid.Bytes[:])
	if e != nil {
		return "", e
	}
	return parsedUuid.String(), nil
}

func pgUuidFromUuid(id string) (pgtype.UUID, error) {
	pgUuid := pgtype.UUID{Status: pgtype.Undefined}
	parsedId, e := uuid.Parse(id)
	if e != nil {
		return pgUuid, e
	}
	b, e := parsedId.MarshalBinary()
	if e != nil {
		return pgUuid, e
	}
	copy(pgUuid.Bytes[:], b)
	pgUuid.Status = pgtype.Present
	return pgUuid, nil
}

func pgInt2FromUint8(i uint8) pgtype.Int2 {
	return pgtype.Int2{
		Int:    int16(i),
		Status: pgtype.Present,
	}
}

func uint8FromPgInt2(i pgtype.Int2) (uint8, error) {
	if i.Status != pgtype.Present {
		return 0, model.ErrInt2NotPresent
	}
	num := i.Int
	if num < 0 || math.MaxUint8 < num {
		return 0, model.ErrInt2NotCastableToUint8
	}
	return uint8(num), nil
}
