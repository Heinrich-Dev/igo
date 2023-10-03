#!/bin/bash
# Assumes script is run in host pane in tmux
go build

buildResult=$(echo $?)
if [$buildResult != 0]; then
    exit 1
fi
# tmux isnt running, create a session and add a new window
tmuxResult=$(pgrep tmux)
if ["$tmuxResult" == ""]; then
    tmux new -s igo 'go run .' \; split-window -h './igo'
else
    tmux select-pane -t 0
    ./igo
    tmux select-pane -t 1
    ./igo
fi