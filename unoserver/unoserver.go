package unoserver

import (
	"context"
	"fmt"
	"os/exec"
)

var (
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
	Executable: DefaultLibreOfficeExecutable,
}

func New() *Unoserver {
	return &Unoserver{}
}

func Default() *Unoserver {
	return &Unoserver{
		Host:       "127.0.0.1",
		Port:       "2002",
		Executable: DefaultLibreOfficeExecutable,
	}
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

func SetUserInstallation(userInstallation string) {
	unoserver.SetUserInstallation(userInstallation)
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

func (u *Unoserver) SetUserInstallation(userInstallation string) {
	u.UserInstallation = fmt.Sprintf("file://%s", userInstallation)
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

	cmd := exec.CommandContext(ctx, u.Executable, args...)

	return cmd
}
