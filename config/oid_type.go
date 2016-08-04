package config

import (
	"crypto/sha256"
	"hash"
	"regexp"
)

var (
	OidTypes = []*OidType{
		NewOidType("sha256", regexp.MustCompile(`\A[[:alnum:]]{64}`)),
	}
)

type OidType struct {
	Name      string
	Validator *regexp.Regexp
}

func NewOidType(name string, validator *regexp.Regexp) *OidType {
	return &OidType{Name: name, Validator: validator}
}

func OidTypeFromConfig(c *Configuration) *OidType {
	var name = c.OidType()
	for _, o := range OidTypes {
		if o.Name == name {
			return o
		}
	}
	return OidTypes[0]

}

func (h *OidType) GetHasher() hash.Hash {
	switch h.Name {
	case "sha256":
		return sha256.New()
	default:
		return nil
	}
}
