package main

import (
	"encoding/base64"
	"fmt"
	"github.com/fxamacker/cbor/v2"
	"github.com/ldclabs/cose/cose"
	"github.com/ldclabs/cose/cwt"
	"github.com/ldclabs/cose/iana"
	"github.com/ldclabs/cose/key"
	"log"
	"time"
)

func ExampleMac0Message1() {
	// load key
	k := key.Key{
		iana.KeyParameterKty:        iana.KeyTypeSymmetric,
		iana.KeyParameterKid:        []byte("akamai_key_hs256"),
		iana.KeyParameterAlg:        iana.AlgorithmHMAC_256_256,
		iana.SymmetricKeyParameterK: key.HexBytesify("403697de87af64611c1d32a05dab0fe1fcb715a86ab435f1ec99192d79569388"),
	}

	macer, err := k.MACer()
	if err != nil {
		panic(err)
	}

	payload := FormPayload()
	//to add payload
	obj := &cose.Mac0Message[[]byte]{
		Protected: cose.Headers{
			iana.HeaderParameterAlg: iana.AlgorithmHMAC_256_256,
		},
		Unprotected: cose.Headers{
			iana.HeaderParameterKid: k.Kid(),
		},
		Payload: payload.Bytesify(),
	}

	// compute MAC
	err = obj.Compute(macer, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Tag: %x\n", obj.Tag())
	// Tag: 726043745027214f

	cwtData, err := obj.MarshalCBOR()
	if err != nil {
		panic(err)
	}

	log.Println(base64.URLEncoding.EncodeToString(cwtData))
	var obj3 cose.Mac0Message[[]byte]
	err = cbor.Unmarshal(cwtData, &obj3)
	if err != nil {
		panic(err)
	}
	// verify MAC
	err = obj3.Verify(macer, nil)
	if err != nil {
		panic(err)
	}
	log.Println("successfully verified")
	//log.Println(hexToBase64(string(obj2.Tag())))
}

// todo write method to form payload basis request
func FormPayload() cwt.ClaimsMap {
	payload := cwt.ClaimsMap{
		iana.Catu: cwt.ClaimsMap{
			iana.Host: []any{iana.Exact, "ak-mediavod-cat-poc-staging.jiocinema.com"},
			iana.Path: []any{iana.Exact, "/jcvod/video/movie/ccp_gangsofwasseypur2_mv_ott/hindi/fhd/h264/dolby_5point1/1717852550/jiocinemanondrmdash-4b182e5feca04a759a3a3d7734f9e5c3/dash-web-premium-plain-a9b3bb9d2a6941b1b18622616fbad3dd/master.mpd"},
		},
		iana.CatM: []any{"GET"},
		iana.Exp:  time.Now().Add(-24 * time.Hour).Unix(),
		//iana.Catalpn: [][]byte{
		//	[]byte("h2"),
		//	[]byte("h3"),
		//},
		iana.CatR: cwt.ClaimsMap{
			iana.RenewalType:  2,
			iana.ExpExtension: 120,
		},
	}

	log.Println(payload)
	return payload
}

func base64EncodeWithTag(tag int, data string) (string, error) {
	// Convert tag to string (you can adjust the format as needed)
	tagStr := fmt.Sprintf("%d:", tag)

	// Combine tag and data
	combined := tagStr + data

	// Encode the combined string to Base64
	encoded := base64.URLEncoding.EncodeToString([]byte(combined))

	return encoded, nil
}
