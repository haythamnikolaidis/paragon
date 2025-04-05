package rpc_test

import (
	"paragon/rpc"
	"testing"
)

type EncodingExample struct {
    Testing bool
}

func TestEncode(t *testing.T) {
   expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
   actual := rpc.EncodeMessage(EncodingExample{Testing: true})
   if expected != actual {
       t.Fatalf("Expected: %s, Actual: %s", expected, actual)
   }

}


func TestDecode(t *testing.T) {
    testmsg := "Content-Length: 15\r\n\r\n{\"method\":\"hi\"}"
    expectedLength := 15
    msg, content, err := rpc.DecodeMessage([]byte(testmsg))
    actualLength := len(content)
    if err != nil {
        t.Fatal(err)
    }

    if expectedLength != actualLength {
       t.Fatalf("Expected Length: %d, Actual Length: %d", expectedLength, actualLength)
    }

    if msg != "hi" {
       t.Fatalf("Expected Message: %s, Actual Message: %s", "hi", msg)
    }
}
