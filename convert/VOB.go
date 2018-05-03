package convert

import (
	"strings"
	"fmt"
)

type VOBRule struct{}

func (r VOBRule) GetName() string {
	return "VOB"
}

func (r VOBRule) IsSupports(path, name, extension, vcodec, acodec string) bool {
	return strings.ToLower(extension) == "vob"
}

func (r VOBRule) GetCommand(outputFormat, inFile, outFile, vcodec, acodec string) string {
	return fmt.Sprintf(
		"ffmpeg -i %s -vcodec libx264 -acodec aac -strict experimental %s",
		inFile,
		outFile,
	)
}
