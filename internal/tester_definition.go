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
	},
	ExecutableFileName: "your_program.sh",
	TestCases: []tester_definition.TestCase{
		{
			Slug:     "si0",
			TestFunc: test1,
		},
		{
			Slug:     "xt9",
			TestFunc: test2,
		},
		{
			Slug:     "wd4",
			TestFunc: test3,
			Timeout:  20 * time.Second,
		},
		{
			Slug:     "zz5",
			TestFunc: test4,
			Timeout:  20 * time.Second,
		},
	},
}
