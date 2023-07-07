package unoserver

import (
	"context"
	"time"
)

var (
	DefaultContextTimeout = 0 * time.Minute
)

var (
	ContextTimeout = DefaultContextTimeout
)

var unoserver = &Unoserver{
	Interface:  "127.0.0.1",
	Port:       "2002",
	Executable: "libreoffice",
}

func SetExecutable(executable string) {
	unoserver.SetExecutable(executable)
}

func SetInterface(interf string) {
	unoserver.SetInterface(interf)
}

func SetPort(port string) {
	unoserver.SetPort(port)
}

func SetContextTimeout(timeout time.Duration) {
	unoserver.SetContextTimeout(timeout)
}

func Run(infile string, outfile string, opts ...string) error {
	return unoserver.Run(infile, outfile, opts...)
}

func RunContext(ctx context.Context, infile string, outfile string, opts ...string) error {
	return unoserver.RunContext(ctx, infile, outfile, opts...)
}

type Unoserver struct {
	Interface  string
	Port       string
	Executable string
}

func (u *Unoserver) SetExecutable(executable string) {
	u.Executable = executable
}

func (u *Unoserver) SetInterface(interf string) {
	u.Interface = interf
}

func (u *Unoserver) SetPort(port string) {
	u.Port = port
}

func (u *Unoserver) SetContextTimeout(timeout time.Duration) {
	ContextTimeout = timeout
}

func (u *Unoserver) Run(infile string, outfile string, opts ...string) error {
	return nil
}

func (u *Unoserver) RunContext(ctx context.Context, infile string, outfile string, opts ...string) error {
	return nil
}
