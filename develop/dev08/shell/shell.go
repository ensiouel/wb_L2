package shell

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	processpkg "github.com/shirou/gopsutil/process"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

func Run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(dir + ">")

		lineBytes, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		line := string(lineBytes)
		if line == "" {
			continue
		}

		if strings.Contains(line, "|") {
			err = ExecPipe(line)
			if err != nil {
				fmt.Println(err)
			}

			continue
		}

		err = Exec(line)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func Exec(line string) error {
	args := strings.Split(line, " ")

	command := args[0]
	args = args[1:]

	switch command {
	case "cd":
		err := cd(args)
		if err != nil {
			return fmt.Errorf("cd: %s", err)
		}

		return nil
	case "pwd":
		err := pwd()
		if err != nil {
			return fmt.Errorf("pwd: %s", err)
		}

		return nil
	case "echo":
		err := echo(args)
		if err != nil {
			return fmt.Errorf("echo: %s", err)
		}

		return nil
	case "kill":
		err := kill(args)
		if err != nil {
			return fmt.Errorf("kill: %s", err)
		}

		return nil
	case "ps":
		err := ps()
		if err != nil {
			return fmt.Errorf("ps: %s", err)
		}

		return nil
	case "exec":
		path, err := exec.LookPath(args[0])
		if err != nil {
			return err
		}

		return syscall.Exec(path, args[1:], os.Environ())
	case "exit":
		os.Exit(0)
	default:
		return fmt.Errorf("unknown command: %s", command)
	}

	return nil
}

func cd(args []string) error {
	if len(args) == 0 {
		path, err := filepath.Abs(".")
		if err != nil {
			return err
		}

		fmt.Println(path)
		return nil
	}

	return os.Chdir(args[0])
}

func pwd() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Println(dir)
	return nil
}

func echo(args []string) error {
	argsLen := len(args)

	for i, arg := range args {
		fmt.Print(arg)
		if i < argsLen-1 {
			fmt.Print(" ")
		}
	}

	fmt.Println()
	return nil
}

func kill(args []string) error {
	if len(args) == 0 {
		return errors.New("no pid specified")
	}

	pid, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	return process.Kill()
}

func ps() error {
	processes, err := processpkg.Processes()
	if err != nil {
		return err
	}

	fmt.Println("NAME\tPID")

	for _, process := range processes {
		name, err := process.Name()
		if err != nil {
			return err
		}

		fmt.Printf("%s\t%d\n", name, process.Pid)
	}

	return err
}

func ExecPipe(line string) error {
	pipeArgs := strings.Split(line, "|")

	var buffer bytes.Buffer
	for i := 0; i < len(pipeArgs); i++ {
		args := strings.Split(pipeArgs[i], " ")

		command := args[0]
		args = args[1:]

		cmd := exec.Command(command, args...)
		cmd.Stdin = bytes.NewReader(buffer.Bytes())
		buffer.Reset()
		cmd.Stdout = &buffer

		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	return nil
}
