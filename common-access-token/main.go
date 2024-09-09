package main

//func main() {
//	// Create a ed25519 signer key
//	privKey, err := ed25519.GenerateKey()
//	if err != nil {
//		panic(err)
//	}
//	signer, err := privKey.Signer()
//	if err != nil {
//		panic(err)
//	}
//
//	// Create a verifier key
//	pubKey, err := ed25519.ToPublicKey(privKey)
//	if err != nil {
//		panic(err)
//	}
//	verifier, err := pubKey.Verifier()
//	if err != nil {
//		panic(err)
//	}
//
//	// create a claims set
//	claims := cwt.Claims{
//		Issuer:     "ldc:ca",
//		Subject:    "ldc:chain",
//		Audience:   "ldc:txpool",
//		Expiration: 1670123579,
//		CWTID:      []byte{1, 2, 3, 4},
//	}
//
//	// sign with Sign1Message
//	obj := cose.Sign1Message[cwt.Claims]{Payload: claims}
//	cwtData, err := obj.SignAndEncode(signer, nil)
//	if err != nil {
//		panic(err)
//	}
//
//	// decode and verify the cwt
//	obj2, err := cose.VerifySign1Message[cwt.Claims](verifier, cwtData, nil)
//	if err != nil {
//		panic(err)
//	}
//
//	// validate the cwt's claims
//	validator, err := cwt.NewValidator(&cwt.ValidatorOpts{
//		ExpectedIssuer:   "ldc:ca",
//		ExpectedAudience: "ldc:txpool",
//		ClockSkew:        time.Minute,
//	})
//	if err != nil {
//		panic(err)
//	}
//
//	err = validator.Validate(&obj2.Payload)
//	fmt.Printf("Validate Claims: %v\n", err)
//	// Validate Claims: cose/cwt: Validator.Validate: token has expired
//
//	cborData, err := key.MarshalCBOR(obj2.Payload)
//	// cborData, err := cbor.Marshal(myClaims)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("CBOR(%d bytes): %x\n", len(cborData), cborData)
//	// CBOR(44 bytes): a501666c64633a636102696c64633a636861696e036a6c64633a7478706f6f6c041a638c103b074401020304
//}
