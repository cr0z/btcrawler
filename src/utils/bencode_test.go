/**
 * Created by 93201 on 2017/5/10.
 */
package utils

import (
	"testing"
)

func TestBencodeInt(t *testing.T) {
	result := BencodeInt(23)
	t.Log("EncodedInt: " + result)
}

func TestBencodeString(t *testing.T) {
	result := BencodeString("testing")
	t.Log("EncodedString: " + result)
}
func TestBencodeList(t *testing.T) {
	list := BencodedList{}
	list.AppendInt(20)
	list.AppendString("HerpDerp")
	t.Log("EncodedList: " + list.StringValue())
}

// Test BencodedDict
func TestBencodedDict(t *testing.T) {
	dict := BencodedDict{}
	testList := BencodedList{}
	testList.AppendInt(6969)
	dict.AppendInt("testInt", 1337)
	dict.AppendString("testString", "test")
	dict.AppendBencodedList("testDict", testList)
	t.Log(dict.StringValue())
}
