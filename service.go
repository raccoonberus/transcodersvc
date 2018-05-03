package transcodersvc

import (
	"strings"
	"github.com/racoonberus/transcodersvc/file"
	"fmt"
)

type VideoConvertRule interface {
	GetName() string
	IsSupports(path, name, extension, vcodec, acodec string) bool
	GetCommand(outputFormat, inFile, outFile, vcodec, acodec string) string
}

type VideoConvertContext struct {
	rules []VideoConvertRule
}

func (ctx *VideoConvertContext) Add(rule VideoConvertRule) {
	ctx.rules = append(ctx.rules, rule)
}

func (ctx VideoConvertContext) GetCmd(filename string, outputFormat string) (string, error) {
	parts := strings.Split(filename, ".")
	name, extension := parts[0], parts[1]
	vcodec, err := file.GetVideoCodec(filename)
	acodec := ""
	for _, rule := range ctx.rules {
		if err != nil {
			return "", err
		}

		outFile := fmt.Sprintf("/bucket/output/%s.mp4", name)

		if rule.IsSupports(filename, name, extension, vcodec, "") {
			return rule.GetCommand(outputFormat,filename, outFile, vcodec, acodec), nil
		}
	}

	return "",
		fmt.Errorf("unsupported video format - filename: %s ; vcodec - %s ; acodec - %s",
			filename,
			vcodec,
			acodec,
		)
}
