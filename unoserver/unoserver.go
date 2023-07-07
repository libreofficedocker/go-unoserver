package unoserver

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"
)

var (
	DefaultContextTimeout = 0 * time.Minute
)

var (
	ContextTimeout       = DefaultContextTimeout
	OoSetupConnectionURL = "socket,host=%s,port=%s,tcpNoDelay=1;urp;StarOffice.ComponentContext"
	OoSetupFlags         = []string{
		"--headless",
		"--invisible",
		"--nocrashreport",
		"--nodefault",
		"--nologo",
		"--nofirststartwizard",
		"--norestore",
	}
)

var unoserver = &Unoserver{
	Host:       "127.0.0.1",
	Port:       "2002",
	Executable: "libreoffice",
}

func New() *Unoserver {
	return &Unoserver{}
}

func SetExecutable(executable string) {
	unoserver.SetExecutable(executable)
}

func SetInterface(host string) {
	unoserver.SetInterface(host)
}

func SetPort(port string) {
	unoserver.SetPort(port)
}

func SetContextTimeout(timeout time.Duration) {
	unoserver.SetContextTimeout(timeout)
}

func Command(opts ...string) *exec.Cmd {
	return unoserver.Command(opts...)
}

func CommandContext(ctx context.Context, opts ...string) *exec.Cmd {
	return unoserver.CommandContext(ctx, opts...)
}

type Unoserver struct {
	Host             string
	Port             string
	Executable       string
	UserInstallation string
}

func (u *Unoserver) SetExecutable(executable string) {
	u.Executable = executable
}

func (u *Unoserver) SetInterface(host string) {
	u.Host = host
}

func (u *Unoserver) SetPort(port string) {
	u.Port = port
}

func (u *Unoserver) SetContextTimeout(timeout time.Duration) {
	ContextTimeout = timeout
}

func (u *Unoserver) Command(opts ...string) *exec.Cmd {
	var args = []string{}

	if u.Host == "" {
		u.Host = "127.0.0.1"
	}

	connection := fmt.Sprintf(OoSetupConnectionURL, u.Host, u.Port)

	args = append(args, OoSetupFlags...)
	args = append(args, fmt.Sprintf("--accept=%s", connection))

	if u.UserInstallation != "" {
		args = append(args, fmt.Sprintf("-env:UserInstallation=%s", u.UserInstallation))
	}

	log.Println("Running: ", u.Executable, args)

	cmd := exec.Command(u.Executable, args...)

	return cmd
}

func (u *Unoserver) CommandContext(ctx context.Context, opts ...string) *exec.Cmd {
	var args = []string{}

	if u.Host == "" {
		u.Host = "127.0.0.1"
	}

	connection := fmt.Sprintf(OoSetupConnectionURL, u.Host, u.Port)

	args = append(args, OoSetupFlags...)
	args = append(args, fmt.Sprintf("--accept=%s", connection))

	if u.UserInstallation != "" {
		args = append(args, fmt.Sprintf("-env:UserInstallation=%s", u.UserInstallation))
	}

	log.Println("Running: ", u.Executable, args)

	cmd := exec.CommandContext(ctx, u.Executable, args...)

	return cmd
}
