#!/bin/bash
go build igo.go board.go

# change this variable to 1 if tmux isn't your thing, and
# you would rather just build the application
tmuxPref=0

buildResult=$(echo $?)

if [[ $buildResult -ne 0 ]]; then
    exit 1
fi

if [[ $tmuxPref == 1 ]]; then
    exit 0
fi

tmux split-window -h
tmux select-pane -t 1
tmux send-keys './igo'
tmux select-pane -t 0
tmux send-keys './igo'