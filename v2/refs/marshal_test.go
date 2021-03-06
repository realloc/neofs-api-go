package refs_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/stretchr/testify/require"
)

func TestOwnerID_StableMarshal(t *testing.T) {
	ownerFrom := new(refs.OwnerID)

	t.Run("non empty", func(t *testing.T) {
		ownerFrom.SetValue([]byte("Owner ID"))

		wire, err := ownerFrom.StableMarshal(nil)
		require.NoError(t, err)

		ownerTo := new(refs.OwnerID)
		require.NoError(t, ownerTo.Unmarshal(wire))

		require.Equal(t, ownerFrom, ownerTo)
	})
}

func TestContainerID_StableMarshal(t *testing.T) {
	cnrFrom := new(refs.ContainerID)

	t.Run("non empty", func(t *testing.T) {
		cnrFrom.SetValue([]byte("Container ID"))

		wire, err := cnrFrom.StableMarshal(nil)
		require.NoError(t, err)

		cnrTo := new(refs.ContainerID)
		require.NoError(t, cnrTo.Unmarshal(wire))

		require.Equal(t, cnrFrom, cnrTo)
	})
}

func TestObjectID_StableMarshal(t *testing.T) {
	objectIDFrom := new(refs.ObjectID)

	t.Run("non empty", func(t *testing.T) {
		objectIDFrom.SetValue([]byte("Object ID"))

		wire, err := objectIDFrom.StableMarshal(nil)
		require.NoError(t, err)

		objectIDTo := new(refs.ObjectID)
		require.NoError(t, objectIDTo.Unmarshal(wire))

		require.Equal(t, objectIDFrom, objectIDTo)
	})
}

func TestAddress_StableMarshal(t *testing.T) {
	cid := []byte("Container ID")
	oid := []byte("Object ID")

	addressFrom := generateAddress(cid, oid)

	t.Run("non empty", func(t *testing.T) {
		wire, err := addressFrom.StableMarshal(nil)
		require.NoError(t, err)

		addressTo := new(refs.Address)
		require.NoError(t, addressTo.Unmarshal(wire))

		require.Equal(t, addressFrom, addressTo)
	})
}

func TestChecksum_StableMarshal(t *testing.T) {
	checksumFrom := new(refs.Checksum)

	t.Run("non empty", func(t *testing.T) {
		checksumFrom.SetType(refs.TillichZemor)
		checksumFrom.SetSum([]byte("Homomorphic Hash"))

		wire, err := checksumFrom.StableMarshal(nil)
		require.NoError(t, err)

		checksumTo := new(refs.Checksum)
		require.NoError(t, checksumTo.Unmarshal(wire))

		require.Equal(t, checksumFrom, checksumTo)
	})
}

func TestSignature_StableMarshal(t *testing.T) {
	signatureFrom := generateSignature("Public Key", "Signature")

	t.Run("non empty", func(t *testing.T) {
		wire, err := signatureFrom.StableMarshal(nil)
		require.NoError(t, err)

		signatureTo := new(refs.Signature)
		require.NoError(t, signatureTo.Unmarshal(wire))

		require.Equal(t, signatureFrom, signatureTo)
	})
}

func TestVersion_StableMarshal(t *testing.T) {
	versionFrom := generateVersion(2, 0)

	t.Run("non empty", func(t *testing.T) {
		wire, err := versionFrom.StableMarshal(nil)
		require.NoError(t, err)

		versionTo := new(refs.Version)
		require.NoError(t, versionTo.Unmarshal(wire))

		require.Equal(t, versionFrom, versionTo)
	})
}

func generateSignature(k, v string) *refs.Signature {
	sig := new(refs.Signature)
	sig.SetKey([]byte(k))
	sig.SetSign([]byte(v))

	return sig
}

func generateVersion(maj, min uint32) *refs.Version {
	version := new(refs.Version)
	version.SetMajor(maj)
	version.SetMinor(min)

	return version
}

func generateAddress(bCid, bOid []byte) *refs.Address {
	addr := new(refs.Address)

	cid := new(refs.ContainerID)
	cid.SetValue(bCid)

	oid := new(refs.ObjectID)
	oid.SetValue(bOid)

	return addr
}
