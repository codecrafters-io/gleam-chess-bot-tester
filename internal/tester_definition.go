package internal

import (
	"github.com/codecrafters-io/tester-utils/tester_definition"
)

var testerDefinition = tester_definition.TesterDefinition{
	AntiCheatTestCases: []tester_definition.TestCase{},
	ExecutableFileName: "script.sh",
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
		},
		{
			Slug:     "a04",
			TestFunc: test4,
		},
		{
			Slug:     "a05",
			TestFunc: test5,
		},
	},
}
