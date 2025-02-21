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
		"-t", strconv.Itoa(delay),
		"-I", ip,
		"-U", username+"%"+password,
		"-C", fmt.Sprintf("此设备将于%d秒后关闭", delay))
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
