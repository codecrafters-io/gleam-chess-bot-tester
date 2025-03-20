package chess_bot_executable

import (
	"fmt"
	"os"
	"path"
	"strconv"
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

func (b *ChessBotExecutable) cleanupCgroup() error {
	b.logger.Debugf("Cleaning up cgroup at %s", b.cgroupPath)
	err := os.RemoveAll(b.cgroupPath)
	if err != nil {
		b.logger.Errorf("Error removing cgroup: %v", err)
		return err
	}
	return nil
}

func (b *ChessBotExecutable) setupCgroupConstraints() error {
	if b.executable == nil || b.executable.Process == nil {
		return fmt.Errorf("executable or process is nil")
	}

	pid := b.executable.Process.Pid
	pidStr := strconv.Itoa(pid) // Convert pid to string

	memoryProcsPath := path.Join(b.cgroupPath, "memory", "cgroup.procs")
	cpuProcsPath := path.Join(b.cgroupPath, "cpu", "cgroup.procs")

	b.logger.Debugf("Adding process %d to memory cgroup: %s", pid, memoryProcsPath)
	if err := os.WriteFile(memoryProcsPath, []byte(pidStr), 0644); err != nil {
		return fmt.Errorf("failed to add process to memory cgroup (%s): %v", memoryProcsPath, err)
	}

	b.logger.Debugf("Adding process %d to CPU cgroup: %s", pid, cpuProcsPath)
	if err := os.WriteFile(cpuProcsPath, []byte(pidStr), 0644); err != nil {
		return fmt.Errorf("failed to add process to CPU cgroup (%s): %v", cpuProcsPath, err)
	}

	return nil
}

func (b *ChessBotExecutable) setupCgroup() error {
	b.logger.Debugf("Setting up cgroup at %s", b.cgroupPath)

	memoryPath := path.Join(b.cgroupPath, "memory")
	cpuPath := path.Join(b.cgroupPath, "cpu")

	b.logger.Debugf("Creating memory cgroup directory: %s", memoryPath)
	if err := os.MkdirAll(memoryPath, 0755); err != nil {
		return fmt.Errorf("failed to create memory cgroup (%s): %v", memoryPath, err)
	}

	b.logger.Debugf("Creating cpu cgroup directory: %s", cpuPath)
	if err := os.MkdirAll(cpuPath, 0755); err != nil {
		return fmt.Errorf("failed to create cpu cgroup (%s): %v", cpuPath, err)
	}

	memoryLimitPath := path.Join(b.cgroupPath, "memory", "memory.limit_in_bytes")
	b.logger.Debugf("Setting memory limit to 268435456 in %s", memoryLimitPath)
	if err := os.WriteFile(memoryLimitPath, []byte("268435456"), 0644); err != nil {
		return fmt.Errorf("failed to set memory limit (%s): %v", memoryLimitPath, err)
	}

	cpuQuotaPath := path.Join(b.cgroupPath, "cpu", "cpu.cfs_quota_us")
	b.logger.Debugf("Setting CPU quota to 50000 in %s", cpuQuotaPath)
	if err := os.WriteFile(cpuQuotaPath, []byte("50000"), 0644); err != nil {
		return fmt.Errorf("failed to set CPU quota (%s): %v", cpuQuotaPath, err)
	}

	return nil
}

func (b *ChessBotExecutable) GetExecutableDirectory() string {
	return path.Dir(b.executable.Path)
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

	if err := b.setupCgroupConstraints(); err != nil {
		return fmt.Errorf("failed to setup cgroup constraints: %v", err)
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
