package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/zchee/xhyve-bindings"
)

func usage() {
	fmt.Println("Usage: ./main <uuid> <iso> <img> <userdata> <vmlinuz> <initrd>")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 7 {
		usage()
	}

	// TODO GOMAXPROCS is necessary?
	cpus := runtime.NumCPU()

	runtime.GOMAXPROCS(cpus)

	//TODO Summarize the argument
	// xhyve like args
	args := strings.Fields("-A -s 0:0,hostbridge -s 31,lpc -l com1 -s 2:0,virtio-net")
	uuid := os.Args[1]
	cpu := os.Args[2]
	memory := os.Args[3]
	iso := os.Args[4]
	img := os.Args[5]
	kexec := os.Args[6]

	d := flag.Bool("d", false, "Whether or not to launch in the background(like a daemon)")
	flag.Parse()

	// When '-d' flag, execute itself the else below command
	if *d {
		// If did not call the gorutine, xhyve does not running
		go func() {
			cmd := exec.Command("sudo", os.Args[0])
			fmt.Print(cmd)
			cmd.Run()
		}()
	} else {
		xhyve.Exec(append(args,
			"-U", fmt.Sprintf("%s", uuid),
			fmt.Sprintf("-c %sM", cpu),
			fmt.Sprintf("-m %sM", memory),
			fmt.Sprintf("-s 3,ahci-cd,%s", iso),
			fmt.Sprintf("-s 4,virtio-blk,%s", img),
			"-f", fmt.Sprintf("%s", kexec))...)
	}
}
