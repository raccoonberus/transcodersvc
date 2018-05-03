package convert

import "fmt"

type AviRule struct{}

func (r AviRule) GetName() string {
	return "AVI"
}

func (r AviRule) IsSupports(path, name, extension, vcodec, acodec string) bool {
	return extension == "avi"
}

func (r AviRule) GetCommand(outputFormat, inFile, outFile, vcodec, acodec string) string {
	return fmt.Sprintf(
		"ffmpeg -y -i %s -vcodec libx264 -acodec libfaac %s",
		inFile,
		outFile,
	)
}
