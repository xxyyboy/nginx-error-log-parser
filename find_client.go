package main

import (
	"regexp"
)

var clientIPRegex = regexp.MustCompile(`client[: ]+(?P<IP>[0-9a-f\.:]{7,39})`)

func findClient(entry *nginxErrorEntry) {
	if ok := clientIPRegex.MatchString(entry.Message); ok {
		matched := clientIPRegex.FindStringSubmatch(entry.Message)
		entry.Client = matched[1]
		entry.Msg = replaceMatched(entry.Msg, matched[0])
	}
}
