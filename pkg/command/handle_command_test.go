package command

import "testing"

func TestPingCommand(t *testing.T) {
	command := "ping"
	args := SplitArguments(command)
	if args[0] != "ping" || len(args) != 1 {
		t.Errorf(`command %v should return %v and have length 1`, command, command)
	}
}

func TestEchoHelloWithoutQuotationCommand(t *testing.T) {
	command := "echo hello"
	args := SplitArguments(command)
	if args[0] != "echo" || args[1] != "hello" || len(args) != 2 {
		t.Errorf(`command '%v' should return %v and have length 1.But return %v with lenth %v`, command, "hello", args, len(args))
	}
}

func TestEchoHelloWithQuotationCommand(t *testing.T) {
	command := `echo "hello world"`
	args := SplitArguments(command)
	if args[0] != "echo" || args[1] != "hello world" {
		t.Errorf(`command '%v' should return command %v and arg %v. But return %v and %v, also the length is %v`, command, "echo", "hello world", args[0], args[1], len(args))
	}
}
