#!/bin/sh
tmux new-session -d -t 'active'
tmux selectp -t 0 # select the first (0) pane
tmux splitw -h -p 50 # split it into two halves

tmux selectp -t 1 # select the new, second (1) pane
tmux splitw -v -p 50 # split it into two halves
tmux selectp -t 0 # go back to the first pane
