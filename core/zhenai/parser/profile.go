package parser

import (
	"regexp"
	"strconv"
	"u2pppw/core/engine"
)

const ageRe = ``
func ParseProfile(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(ageRe)
	match := re.FindSubmatch(contents)
	if match != nil {
		age, err := strconv.Atoi(string(match[1]))
		if err != nil {
			
		}
	}
}