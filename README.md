basicly this is a tmux session for the system administrators that will diagnose problems and offer solutions 

Purpose:
Mainly when a admin enters to a host system, to understand the running system the engineer runs some commands to see the states of the processes.
This screens will be doing that without any further configuration need.

Features:
1. Sessions are by default recorded. during the installation you need to ask where to record the sessions (Local or Remote)
2. Suggestions to Engineers for the running commands,
3. Analyse Online/Static for the running systems
4. Kubernetes stack..



optional (TODO):
9. mutliple users can interact the same session, like tmate (later implemented)
. 


Screen1. (left)
Main Screen The administrator is running on this screen.

Screen2. (right-1)
Suggestion screen
suggestions to make some changes, whenever the admin is writing a command the screen should show mostly used similar command completions
like if the user says " # lvcreate .. " on the screen1, the suggestion screen should tell the rest of the command "# lvcreate ............"

Screen3. (right-2)
analyse screen
what kind of tuning can be done, like reeup disk space, increase kernel parameter, change mysql password, the number of inodes is too small, 
STATIC has only defined suggestions, like backup etc, change some paramters if it runs PHP, nginx and so on..





Static Analyse
some best practice executions can be done by default. that can be done also via the CURL, but we need to get some information about the destination system. We will read the OS, total memory usage, installed applications, open ports, running services, error log, system log about errors, hypervisor data.. and request a static file for that analyse..


Dynamic (online) Analyse
first we will use a simple bash programm to take the parameters to run on the host systems, later we need to write a cross-compiled go program to get the values from the Knowledgebase.. the bash runner (1st)

Online anaylse will send some queries to the destination api, and the results will create execution routines. Like we will have some static Templates that the queries will fill the required informations and run them to get the bash scripts to be executed locally..




Remote driven push process will be created to run on the tmux session to analyse the problems...
--> later will be defined more..

a bash script will remotely create XYZ.sh files for simple iterations..