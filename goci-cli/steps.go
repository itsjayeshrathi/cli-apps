package main

import "os/exec"

type step struct {
	name     string
	exe      string
	args     []string
	messsage string
	proj     string
}

func newStep(name, exe, message, proj string, args []string) step {
	return step{
		name:     name,
		exe:      exe,
		args:     args,
		messsage: message,
		proj:     proj,
	}
}

func (s step) execute() (string, error) {
	cmd := exec.Command(s.exe, s.args...)
	cmd.Dir = s.proj
	if err := cmd.Run(); err != nil {
		return "", &stepErr{
			step:  s.name,
			msg:   "falied to execute",
			cause: err,
		}
	}
	return s.messsage, nil
}
