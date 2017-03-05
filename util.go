package common

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

//IsHex check if string is hex string
func IsHex(str string) (bool, error) {
	return regexp.MatchString("^([0-9A-Fa-f])+$", str)
}

//ToBytes convert string (hex/base64) to byte array
func ToBytes(str string, l ...int) ([]byte, error) {
	r, err := IsHex(str)

	if r && err == nil {
		if len(l) > 0 {
			n := l[0] * 2
			for i := len(str); i < n; i++ {
				str = "0" + str
			}
		}
		if len(str)%2 != 0 {
			str = "0" + str
		}
		return hex.DecodeString(str)
	}
	return base64.StdEncoding.DecodeString(str)
}

//ToHex convert byte to hex string
func ToHex(b []byte) string {
	return hex.EncodeToString(b)
}

//ToBase64 convert byte to base64 string
func ToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

//HexToInt convert hex string to integer
func HexToInt(str string) (uint64, error) {
	r, err := IsHex(str)

	if r && err == nil {
		if len(str)%2 != 0 {
			str = "0" + str
		}
		return strconv.ParseUint(str, 16, 64)
	}
	return 0, TraceError("Not a hex string")
}

//IntToHex convert integer to even hex string, default is 64bit (8 bytes) hex formats (l = 8)
func IntToHex(amt int, l ...int) string {
	h := fmt.Sprintf("%x", amt)
	b := 8
	if len(l) > 0 {
		b = l[0]
	}

	n := b * 2
	for i := len(h); i < n; i++ {
		h = "0" + h
	}
	return h
}

//IntToBytes convert integer to byte array , default is 64bit bytes (l = 8)
func IntToBytes(amt uint64, l ...int) ([]byte, error) {
	h := fmt.Sprintf("%x", amt)
	b := 8
	if len(l) > 0 {
		b = l[0]
	}

	return ToBytes(h, b)
}

//FromJSON convert json string to golang object
func FromJSON(data string, v interface{}) error {
	return json.Unmarshal([]byte(data), v)
}

//ToJSON output to json format
func ToJSON(obj interface{}) string {
	b, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(b)
}

//ToByteJSON output to json byte
func ToByteJSON(obj interface{}) []byte {
	b, err := json.Marshal(obj)
	if err != nil {
		return nil
	}
	return b
}

//FromBase64 convert string base64 to byte array
func FromBase64(str string) []byte {
	r, e := base64.StdEncoding.DecodeString(str)
	if e != nil {
		return nil
	}
	return r
}

//ByteToInt convert 8 bytes array to uint64
func ByteToInt(b []byte) uint64 {
	if len(b) > 8 {
		return 0
	}
	h := ToHex(b)
	i, e := HexToInt(h)
	if e != nil {
		return 0
	}
	return i
}

//TraceError trace error
func TraceError(err interface{}) error {
	if err == nil {
		return nil
	}
	s := fmt.Sprint(err)
	if s == "" || len(s) == 0 {
		return nil
	}

	_, fn, line, _ := runtime.Caller(1)
	//fname := filepath.Base(fn)

	if strings.HasPrefix(s, "[ERROR]") {
		return fmt.Errorf("%v \n %s:%d ", err, fn, line)
		//return errors.New(s)
	}

	return fmt.Errorf("[ERROR]: %s:%d %v ", fn, line, err)
}

//ToString convert value to string representation
func ToString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

//ParseInt convert string number to integer (uint64)
func ParseInt(str string) uint64 {
	r, e := strconv.ParseUint(str, 10, 64)
	if e != nil {
		return 0
	}
	return r
}

//SpliceString array
func SpliceString(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
