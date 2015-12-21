package mschapv1

import (
	"encoding/binary"
)

type ChallengeAttr struct {
	VendorId uint32
	VendorType uint8
	VendorLength uint8
	Value []byte
}

type ResponseAttr struct {
	VendorId uint32
	VendorType uint8
	VendorLength uint8
	Ident uint8
	Flags uint8
	LMResponse []byte //24bytes
	NTResponse []byte //24bytes
}

func DecodeResponse(b []byte) ResponseAttr {
	return ResponseAttr{
		VendorId: binary.BigEndian.Uint32(b[0:4]),
		VendorType: b[4],
		VendorLength: b[5],
		Ident: b[6],
		Flags: b[7],
		LMResponse: b[8:32],
		NTResponse: b[33:57],
	}
}

func DecodeChallenge(b []byte) ChallengeAttr {
	return ChallengeAttr{
		VendorId: binary.BigEndian.Uint32(b[0:4]),
		VendorType: b[4],
		VendorLength: b[5],
		Value: b[6:],
	}
}

/*type MSChap2Attr struct {
	VendorId uint32
	VendorType uint8
	VendorLength uint8
	Ident uint8
	Flags uint8
	PeerChallenge []byte //16bytes
	// 8bytes zero
	Response []byte //24bytes
}*/
/*func DecodeMSCHAPv2(b []byte) MSChap2Attr {
	return MSChap2Attr{
		VendorId: binary.BigEndian.Uint32(b[0:4]),
		VendorType: b[4],
		VendorLength: b[5],
		Ident: b[6],
		Flags: b[7],
		PeerChallenge: b[8:24],
		// reserved 24-32
		Response: b[32:],
	}
}*/