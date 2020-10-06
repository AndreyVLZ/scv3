package main

import "strings"

func mustFrameBeInSequence(id string) bool {
	if id != "TXXX" && strings.HasPrefix(id, "T") {
		return false
	}

	switch id {
	case "MCDI", "ETCO", "SYTC", "RVRB", "MLLT", "PCNT", "RBUF", "POSS", "OWNE", "SEEK", "ASPI":
	case "IPLS", "RVAD": // Specific ID3v2.3 frames.
		return false
	}

	return true
}
func main(){
	b := mustFrameBeInSequence("APIC")
	println(b)
}