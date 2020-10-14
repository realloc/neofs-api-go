package owner

import (
	"github.com/mr-tron/base58"
	"github.com/nspcc-dev/neo-go/pkg/encoding/address"
	"github.com/nspcc-dev/neofs-api-go/internal"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/pkg/errors"
)

// ID represents v2-compatible owner identifier.
type ID refs.OwnerID

// ErrBadID should be returned when bytes slice isn't Owner.ID size (NEO3WalletSize).
// Notice: if byte slice changed, please, replace error message.
const ErrBadID = internal.Error("owner.ID should be 25 bytes length")

// NewIDFromV2 wraps v2 OwnerID message to ID.
func NewIDFromV2(idV2 *refs.OwnerID) *ID {
	return (*ID)(idV2)
}

// NewID creates and initializes blank ID.
//
// Works similar as NewIDFromV2(new(OwnerID)).
func NewID() *ID {
	return NewIDFromV2(new(refs.OwnerID))
}

// SetNeo3Wallet sets owner identifier value to NEO3 wallet address.
func (id *ID) SetNeo3Wallet(v *NEO3Wallet) {
	(*refs.OwnerID)(id).SetValue(v.Bytes())
}

// ToV2 returns the v2 owner ID message.
func (id *ID) ToV2() *refs.OwnerID {
	return (*refs.OwnerID)(id)
}

func (id *ID) String() string {
	return base58.Encode((*refs.OwnerID)(id).GetValue())
}

func ScriptHashBE(id *ID) ([]byte, error) {
	addr, err := address.StringToUint160(id.String())
	if err != nil {
		return nil, err
	}

	return addr.BytesBE(), nil
}

// NewIDFromNeo3Wallet creates new owner identity from 25-byte neo wallet.
func NewIDFromNeo3Wallet(v *NEO3Wallet) *ID {
	id := NewID()
	id.SetNeo3Wallet(v)

	return id
}

// Parse converts base58 string representation into ID.
func (id *ID) Parse(s string) error {
	data, err := base58.Decode(s)
	if err != nil {
		return errors.Wrap(err, "could not parse owner.ID from string")
	} else if len(data) != NEO3WalletSize {
		return ErrBadID
	}

	(*refs.OwnerID)(id).SetValue(data)

	return nil
}
