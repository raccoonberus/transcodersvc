package file

import (
	"fmt"
	"os"
	"strings"

	"github.com/racoonberus/transcodersvc/command"
)

func GetVideoCodec(filename string) (string, error) {
	cmd := fmt.Sprintf(
		"ffprobe -v error -select_streams v:0 -show_entries stream=codec_name -of default=nokey=1:noprint_wrappers=1 /bucket/input/%s | tail -1",
		filename,
	)
	cc, err := command.ShellExec(cmd)
	if err != nil {
		return "", err
	}
	codec := string(cc)
	codec = strings.Trim(codec, "\n\t ")
	return codec, err
}

func Exists(filename string) (bool, error) {
	_, err := os.Stat(filename)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
