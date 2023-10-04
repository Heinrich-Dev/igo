#!/bin/bash
go build

# change this variable to 1 if tmux isn't your thing, and
# you would rather just build the application
tmuxPref=0

if [[ $tmuxPref == 1 ]]; then
    exit 0
fi
buildResult=$(echo $?)

if [[ $buildResult -ne 0 ]]; then
    exit 1
fi

tmuxOpen=$(pgrep tmux)

if [[ "$tmuxOpen" == "" ]]; then
    tmux new -s igo './igo' \; split-window -h './igo' \; select-pane -t 0
fi