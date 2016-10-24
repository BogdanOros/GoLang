package tools

import (
"os/exec"
"os"
	"strconv"
)

func ReadKey() int{
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()

	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)
	return int(b[0])
}

func ClearScreen() error{
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func GoToXY (x int, y int) error {
	cmd := exec.Command("tput","cup", strconv.Itoa(x), strconv.Itoa(y))
	cmd.Stdout = os.Stdout
	return cmd.Run()
}


func tput(args ...string) error {
	cmd := exec.Command("tput", args...)
	cmd.Stdout = os.Stdout
	return cmd.Run();
}
