package md5

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
     "sort"
)

// ID is an alias for the MD5 hash represented as a string.
type ID string

func NewID(data map[string]interface{}) ID {
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
     sort.Strings(keys)

	// Concatenate the sorted key-value pairs to form the input data string
	var inputStr string
	for _, k := range keys {
		inputStr += fmt.Sprintf("%s=%v|", k, data[k])
	}

	// Compute the MD5 hash of the input data string
	hasher := md5.New()
	hasher.Write([]byte(inputStr))
	hash := hasher.Sum(nil)

	// Convert the hash to a hexadecimal string and return it as an ID
	return ID(hex.EncodeToString(hash))
}

func NewWithSourceID(data map[string]interface{}, source string) ID {
     keys := make([]string, 0, len(data))
     for k := range data {
          keys = append(keys, k)
     }
     sort.Strings(keys)

     // Concatenate the sorted key-value pairs to form the input data string
     var inputStr string
     for _, k := range keys {
          inputStr += fmt.Sprintf("%s=%v|", k, data[k])
     }
     inputStr += fmt.Sprintf("source=%s", source)

     // Compute the MD5 hash of the input data string
     hasher := md5.New()
     hasher.Write([]byte(inputStr))
     hash := hasher.Sum(nil)

     // Convert the hash to a hexadecimal string and return it as an ID
     return ID(hex.EncodeToString(hash))
}

func NewWithSourceAndServiceID(data map[string]interface{}, source string, service string) ID {
     keys := make([]string, 0, len(data))
     for k := range data {
          keys = append(keys, k)
     }
     sort.Strings(keys)

     // Concatenate the sorted key-value pairs to form the input data string
     var inputStr string
     for _, k := range keys {
          inputStr += fmt.Sprintf("%s=%v|", k, data[k])
     }
     inputStr += fmt.Sprintf("source=%s", source)
     inputStr += fmt.Sprintf("service=%s", service)

     // Compute the MD5 hash of the input data string
     hasher := md5.New()
     hasher.Write([]byte(inputStr))
     hash := hasher.Sum(nil)

     // Convert the hash to a hexadecimal string and return it as an ID
     return ID(hex.EncodeToString(hash))
}

func NewMd5Hash(data string) ID {
     hasher := md5.New()
     hasher.Write([]byte(data))
     hash := hasher.Sum(nil)
     return ID(hex.EncodeToString(hash))
}
