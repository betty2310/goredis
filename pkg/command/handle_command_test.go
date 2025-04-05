package command

import (
	"testing"
	"time"
)

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

func TestSetStringCommand(t *testing.T) {
	command := `set name huynh`
	args := SplitArguments(command)
	if args[0] != "set" || args[1] != "name" || args[2] != "huynh" {
		t.Errorf(`command '%v' should return command %v with key %v and value %v. But return %v`, command, "set", "name", "huynh", args)
	}
}

func TestSetStringNotTrimCommand(t *testing.T) {
	command := `set name huynh`
	args := SplitArguments(command)
	if args[0] != "set" || args[1] != "name" || args[2] != "huynh" {
		t.Errorf(`command '%v' should return command %v with key %v and value %v. But return %v`, command, "set", "name", "huynh", args)
	}
}

func TestGetCommand(t *testing.T) {
	command := `get name`
	args := SplitArguments(command)
	if args[0] != "get" || args[1] != "name" {
		t.Errorf(`command '%v' should return command %v with key %v. But return %v`, command, "get", "name", args)
	}
}

func TestSetStringHaveExpInSecondCommand(t *testing.T) {
	command := `set name huynh EXP 3s`
	args := SplitArguments(command)

	if args[0] != "set" || args[1] != "name" || args[2] != "huynh" || args[3] != "EXP" || args[4] != "3s" {
		t.Errorf(`string '%v' should return command %v with key %v and value %v have exp date is %v. But return %v`, command, "set", "name", "huynh", args, args[4])
	}

	minutes, err := time.ParseDuration(args[4])
	if err != nil {
		t.Errorf(`Cannot parse %v to datetime`, args[4])
	}
	threeMinutes := 3 * time.Second
	if minutes != threeMinutes {
		t.Errorf(`Exp %v return wrong time`, args[4])
	}
}
