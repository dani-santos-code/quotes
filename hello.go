package phrase

import (
	"rsc.io/quote/v3"
)

// Get returns a phrase
func Get() (sentence string) {
	sentence = quote.GlassV3()
	return sentence
}

// GetOpt returns a phrase
func GetOpt() (sentence string) {
	sentence = quote.OptV3()
	return sentence
}
