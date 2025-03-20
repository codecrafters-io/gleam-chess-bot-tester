package internal

import (
	"time"

	"github.com/codecrafters-io/tester-utils/tester_definition"
)

var testerDefinition = tester_definition.TesterDefinition{
	AntiCheatTestCases: []tester_definition.TestCase{
		{
			Slug:     "anti-cheat-1",
			TestFunc: antiCheatExecute,
		},
		{
			Slug:     "anti-cheat-2",
			TestFunc: antiCheatDeps,
		},
		{
			Slug:     "anti-cheat-3",
			TestFunc: antiCheatFFI, // TODO: Remove FFI, might be a problem
		},
	},
	ExecutableFileName: "your_program.sh",
	TestCases: []tester_definition.TestCase{
		{
			Slug:     "a01",
			TestFunc: test1,
		},
		{
			Slug:     "a02",
			TestFunc: test2,
		},
		{
			Slug:     "a03",
			TestFunc: test3,
			Timeout:  20 * time.Second,
		},
		{
			Slug:     "a04",
			TestFunc: test4,
			Timeout:  20 * time.Second,
		},
	},
}
