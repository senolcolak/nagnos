#!/bin/sh
tmux new-session -d -t 0
tmux selectp -t 0 # select the first (0) pane
tmux splitw -h -p 20 'tail -f /var/log/wifi.log' # split it into two halves
tmux selectp -t 1 # select the new, second (1) pane
tmux splitw -v -p 30 'tail -f /var/log/system.log' # split it into two halves
tmux selectp -t 1 # select bach the second (1) pane
tmux splitw -v -p 50 'tail -f /var/log/install.log'
tmux selectp -t 0 # go back to the first pane
tmux -2 attach-session -d
