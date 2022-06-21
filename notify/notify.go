package notify

import "os/exec"

func Notify(title, text string) {
	cmd := exec.Command("notify-send", title, text)
	cmd.Run()
}
