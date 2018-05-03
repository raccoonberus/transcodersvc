package convert

import (
	"fmt"
	"github.com/racoonberus/transcodersvc/util"
)

type ThreeGPPRule struct{}

func (r ThreeGPPRule) GetName() string {
	return "3GPP/3GPP2"
}

func (r ThreeGPPRule) IsSupports(path, name, extension, vcodec, acodec string) bool {
	return util.InSlice(extension, []interface{}{"3gp", "3g2"}) &&
		util.InSlice(vcodec, []interface{}{"mpeg4", "h263"})
}

func (r ThreeGPPRule) GetCommand(outputFormat, inFile, outFile, vcodec, acodec string) string {
	if "h263" == vcodec {
		return fmt.Sprintf(
			"ffmpeg -i %s -c:v libx264 -c:a aac -strict experimental %s",
			inFile,
			outFile,
		)
	}
	if "mpeg4" == vcodec {
		return fmt.Sprintf(
			"ffmpeg -i %s -vcodec copy -acodec copy %s",
			inFile,
			outFile,
		)
	}
	return ""
}
