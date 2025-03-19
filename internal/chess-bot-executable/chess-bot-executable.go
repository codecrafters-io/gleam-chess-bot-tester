package chess_bot_executable

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/codecrafters-io/tester-utils/executable"
	"github.com/codecrafters-io/tester-utils/logger"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

type ChessBotExecutable struct {
	executable *executable.Executable
	logger     *logger.Logger
	args       []string
	cgroupPath string
}

func NewChessBotExecutable(stageHarness *test_case_harness.TestCaseHarness) *ChessBotExecutable {
	b := &ChessBotExecutable{
		executable: stageHarness.NewExecutable(),
		logger:     stageHarness.Logger,
		cgroupPath: "/sys/fs/cgroup/chess-bot",
	}

	if err := b.setupCgroup(); err != nil {
		b.logger.Debugf("Failed to setup cgroup: %v", err)
	}

	stageHarness.RegisterTeardownFunc(func() {
		b.Kill()
		b.cleanupCgroup()
	})

	return b
}

func (b *ChessBotExecutable) setupCgroup() error {
	if err := os.MkdirAll(path.Join(b.cgroupPath, "memory"), 0755); err != nil {
		return fmt.Errorf("failed to create memory cgroup: %v", err)
	}
	if err := os.MkdirAll(path.Join(b.cgroupPath, "cpu"), 0755); err != nil {
		return fmt.Errorf("failed to create cpu cgroup: %v", err)
	}

	if err := os.WriteFile(
		path.Join(b.cgroupPath, "memory/memory.limit_in_bytes"),
		[]byte("268435456"),
		0644,
	); err != nil {
		return fmt.Errorf("failed to set memory limit: %v", err)
	}

	if err := os.WriteFile(
		path.Join(b.cgroupPath, "cpu/cpu.cfs_quota_us"),
		[]byte("50000"),
		0644,
	); err != nil {
		return fmt.Errorf("failed to set CPU quota: %v", err)
	}

	return nil
}

func (b *ChessBotExecutable) cleanupCgroup() error {
	return os.RemoveAll(b.cgroupPath)
}

func (b *ChessBotExecutable) Run(args ...string) error {
	b.args = args
	if len(b.args) == 0 {
		b.logger.Infof("$ ./%s", path.Base(b.executable.Path))
	} else {
		var log string
		log += fmt.Sprintf("$ ./%s", path.Base(b.executable.Path))
		for _, arg := range b.args {
			if strings.Contains(arg, " ") {
				log += " \"" + arg + "\""
			} else {
				log += " " + arg
			}
		}
		b.logger.Infof("%s", log)
	}

	if err := b.executable.Start(b.args...); err != nil {
		return err
	}

	pid := b.executable.Process.Pid
	if err := os.WriteFile(
		path.Join(b.cgroupPath, "memory/cgroup.procs"),
		[]byte(fmt.Sprintf("%d", pid)),
		0644,
	); err != nil {
		b.logger.Debugf("Failed to add process to memory cgroup: %v", err)
	}
	if err := os.WriteFile(
		path.Join(b.cgroupPath, "cpu/cgroup.procs"),
		[]byte(fmt.Sprintf("%d", pid)),
		0644,
	); err != nil {
		b.logger.Debugf("Failed to add process to CPU cgroup: %v", err)
	}

	return nil
}

func (b *ChessBotExecutable) HasExited() bool {
	return b.executable.HasExited()
}

func (b *ChessBotExecutable) Kill() error {
	b.logger.Debugf("Terminating program")
	if err := b.executable.Kill(); err != nil {
		b.logger.Debugf("Error terminating program: '%v'", err)
		return err
	}

	b.logger.Debugf("Program terminated successfully")
	return nil // When does this happen?
}
