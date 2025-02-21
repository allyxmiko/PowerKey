package action

import (
	"PowerKey/action/wol"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

var WakeOnLAN = wol.WakeOnLAN

func Shutdown(username, password, ip string, delay int) error {
	cmd := exec.Command(
		"net",
		"rpc",
		"shutdown",
		"-f",
		"-t", strconv.Itoa(delay),
		"-C", fmt.Sprintf("设备将于%d秒后关闭", delay),
		"-U", username+"%"+password,
		"-I", ip)
	output, err := cmd.Output()
	if err != nil {
		return err
	}
	result := string(output)
	if !strings.Contains(result, "succeeded") {
		return errors.New(result)
	}
	return nil
}
