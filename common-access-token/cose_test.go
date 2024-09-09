package main

import (
	"testing"
)

//func TestMACer(t *testing.T) {
//	assert := assert.New(t)
//
//	k := key.Key{
//		iana.KeyParameterKty: iana.KeyTypeSymmetric,
//		iana.KeyParameterAlg: iana.AlgorithmHMAC_256_64,
//	}
//
//	k[iana.SymmetricKeyParameterK] = key.GetRandomBytes(32)
//	macer, err := hmac.New(k)
//	require.NoError(t, err)
//
//	tag, err := macer.MACCreate([]byte("hello world"))
//	require.NoError(t, err)
//
//	assert.NoError(macer.MACVerify([]byte("hello world"), tag))
//	assert.ErrorContains(macer.MACVerify([]byte("hello world 1"), tag), "invalid MAC")
//
//	k.SetOps(iana.KeyOperationMacVerify)
//	_, err = macer.MACCreate([]byte("hello world"))
//	assert.ErrorContains(err, "invalid key_ops")
//
//	k.SetOps(iana.KeyOperationMacCreate)
//	assert.ErrorContains(macer.MACVerify([]byte("hello world"), tag), "invalid key_ops")
//
//	k.SetOps(iana.KeyOperationMacVerify, iana.KeyOperationMacCreate)
//	tag, err = macer.MACCreate([]byte("hello world 1"))
//	require.NoError(t, err)
//	assert.NoError(macer.MACVerify([]byte("hello world 1"), tag))
//
//	k[iana.SymmetricKeyParameterK] = key.GetRandomBytes(32)
//	assert.ErrorContains(macer.MACVerify([]byte("hello world 1"), tag), "invalid MAC")
//}
//
//func TestExampleMac0Message(t *testing.T) {
//	// load key
//	k := key.Key{
//		iana.KeyParameterKty:        iana.KeyTypeSymmetric,
//		iana.KeyParameterKid:        []byte("our-secret"),
//		iana.KeyParameterAlg:        iana.AlgorithmHMAC_256_256,
//		iana.SymmetricKeyParameterK: key.Base64Bytesify("hJtXIZ2uSN5kbQfbtTNWbpdmhkV8FJG-Onbc6mxCcYg"),
//	}
//
//	macer, err := k.MACer()
//	if err != nil {
//		panic(err)
//	}
//
//	// create a COSE_Mac0 message
//	obj := &cose.Mac0Message[[]byte]{
//		Unprotected: cose.Headers{},
//		Payload:     []byte("This is the content."),
//	}
//
//	// compute MAC
//	err = obj.Compute(macer, nil)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("Tag: %x\n", obj.Tag())
//	// Tag: 726043745027214f
//
//	// encode COSE_Mac0 message
//	coseData, err := cbor.Marshal(obj)
//	if err != nil {
//		panic(err)
//	}
//
//	// decode a COSE_Mac0 message
//	var obj3 cose.Mac0Message[[]byte]
//	cbor.Unmarshal(coseData, &obj3)
//	if err != nil {
//		panic(err)
//	}
//
//	// verify MAC
//	err = obj3.Verify(macer, nil)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("Payload: %s\n", string(obj3.Payload))
//	// Payload: This is the content.
//	fmt.Printf("Tag: %x\n", obj3.Tag())
//	// Tag: 726043745027214f
//
//	// or verify and decode a COSE_Mac0 message
//	obj2, err := cose.VerifyMac0Message[[]byte](macer, coseData, nil)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("Payload: %s\n", string(obj2.Payload))
//	// Payload: This is the content.
//	fmt.Printf("Tag: %x\n", obj2.Tag())
//	// Tag: 726043745027214f
//	size := len(obj2.Tag()) > 0
//	assert.Equal(t, size, true)
//	// Output:
//	// Tag: 726043745027214f
//	// Payload: This is the content.
//	// Tag: 726043745027214f
//	// Payload: This is the content.
//	// Tag: 726043745027214f
//}

func TestSomething(t *testing.T) {
	ExampleMac0Message1()
	//v := 1
	//assert.Equal(t, v, 1)
}
