#!/bin/bash

# Define variables
CGROUP_PATH="/sys/fs/cgroup/chess-bot"
MEMORY_LIMIT="268435456"  # 256MB in bytes
CPU_MAX="50000 100000"    # 50% CPU limit (50000/100000)

# Remove existing cgroup if present
if [ -d "$CGROUP_PATH" ]; then
  echo "Removing existing cgroup"
  rmdir "$CGROUP_PATH"
fi

# Create new cgroup
echo "Creating cgroup at $CGROUP_PATH"
mkdir -p "$CGROUP_PATH"

# Enable controllers
echo "Enabling controllers"
echo "+memory +cpu" > /sys/fs/cgroup/cgroup.subtree_control

# Set limits
echo "Setting memory limit to $MEMORY_LIMIT"
echo "$MEMORY_LIMIT" > "$CGROUP_PATH/memory.max"

echo "Setting CPU limit"
echo "$CPU_MAX" > "$CGROUP_PATH/cpu.max"

echo "Cgroup created successfully"

# To add a process to this cgroup, use:
# echo $PID > "$CGROUP_PATH/cgroup.procs"