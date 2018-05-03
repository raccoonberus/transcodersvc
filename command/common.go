package command

import "os/exec"

func ShellExec(cmd string) (string, error) {
	cc, err := exec.Command("sh", "-c", cmd).Output()
	return string(cc), err
}
