package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/tiborvass/xhyve-bindings"
)

func usage() {
	fmt.Println("Usage: ./main <vmlinuz> <initrd>")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 3 {
		usage()
	}
	vmlinuz := os.Args[1]
	initrd := os.Args[2]
	args := strings.Fields("-m 1G -s 0:0,hostbridge -s 31,lpc -l com1,stdio")
	xhyve.Spawn(append(args, "-f", fmt.Sprintf("kexec,%s,%s,earlyprintk=serial console=ttyS0", vmlinuz, initrd)))
}
