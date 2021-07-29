package daemon

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func init() {
	cmd := exec.Command(os.Args[0], flag.Args()...)
	if err := cmd.Start(); err != nil {
		fmt.Printf("start %s failed, error: %v\n", os.Args[0], err)
		os.Exit(1)
	}
	fmt.Printf("%s [PID] %d running...\n", os.Args[0], cmd.Process.Pid)
	os.Exit(0)
}
