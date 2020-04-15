watch -n 0.1 'tmux capture-pane -e -pt "0.0"| tail -1'
cut -d '#'  -f2 reads the command if the command is inside the list the suggestion is printed..


this will capture the first panel and see if there is a space on the shell, it will make a suggestion for the command from the experts knowledgebase.




*Static*
There are static files that will get the Most Frequently used command completions..
i.e.

INPUT: "ls " after the space is used the pane will make the suggestion for the command..
ls -ltr   <-- DATE This command will list the files inside directory from ascending order in date
ls -lSr   <-- SIZE This command will list the files inside the directory ascending order in Size

grep

*Online*