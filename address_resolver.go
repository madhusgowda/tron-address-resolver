package TronAddressResolver

import (
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"log"
	"strings"
)

const (
	TRXAddrPrefix = "T"
)

// ResolveAddress resolves a Tron address by decoding its base58 format to hex.
func ResolveAddress(address string) (string, error) {
	if !strings.HasPrefix(address, TRXAddrPrefix) {
		return address, fmt.Errorf("invalid address: %s", address)
	}
	decoded, _, err := base58.CheckDecode(address)
	if err != nil {
		log.Fatalf("Unable to resolve TRX address: %v", address)
		return "", fmt.Errorf(err.Error())
	}
	address = fmt.Sprintf("%x", decoded)
	return "0x" + address, nil
}
