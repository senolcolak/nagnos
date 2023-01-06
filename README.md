basicly this is a tmux session for the system administrators that will diagnose problems and offer solutions 

when a System Admin/ Developer /DevOps /SRE connects to a system or systems, he tries to run some commands to diagnose the problems and analyse the current situation.

#### Sr-Assistant is trying to ease the situation via creating the following functions.

1. Session Recording
Session is recorded and later you can replay everything that you have done on your local computer.
    -  You can easily create a Summary/Report for a given timeframe.
    -  You can integrate it to a ticket system so the copy/paste operations are done automatically and the details are diagnosed.
    -  Recorded sessions can be send to the Sr-Server for metadata purposes


2. SR-Server
Client(Sr-Assistant) can be used standalone but the main goal is to use it with sr-server. sr-srserver is providing the tmux screen the necessary suggestions and is running some commands that would be helpfull for the operator.
    -  There is a parameter on the ``./srassist/config`` file which is setting to send logs local/remote
    -  There is a parameter on ``./srassist/config`` file which has the api credentials for getting suggestions.

3. Supported environments
Client will support plugins like Docker/Kubernetes/Linux/Network, these plugins will enable the functionalities of these use cases.

## Example use case

#### Screen1. (left) Main Screen The administrator is running on this screen.
user is testing the internet connection
```
$ ping google.com
```
#### Screen2. (right) Suggestion screen
suggestions and analyses, 
meanwhile on the right panel
:white_check_mark: df -h any disk < %80  
:white_check_mark: df -i i node numbers < %80  
:white_check_mark: ip route  
....  

whenever the admin is writing a command the screen should show mostly used similar command completions
like if the user says " # lvcreate .. " on the screen1, the suggestion screen should tell the rest of the command "# lvcreate ............"

#### Screen3. (bottom 3 line) tail analyse screen
this screen tails all the logs  and looks for errors. compares the results with the previous admins and creates interaction requests. Like please check that, than press enter and so on..
````
suggestion #1: increase the following kernel parameter for xxxx then press enter
````
what kind of tuning can be done, like reeup disk space, increase kernel parameter, change mysql password, the number of inodes is too small,  
STATIC has only defined suggestions, like backup etc, change some paramters if it runs PHP, nginx and so on..



## Old notes

### Static Analyse
some best practice executions can be done by default. that can be done also via the CURL, but we need to get some information about the destination system. We will read the OS, total memory usage, installed applications, open ports, running services, error log, system log about errors, hypervisor data.. and request a static file for that analyse..


### Dynamic (online) Analyse
first we will use a simple bash programm to take the parameters to run on the host systems, later we need to write a cross-compiled go program to get the values from the Knowledgebase.. the bash runner (1st)

Online anaylse will send some queries to the destination api, and the results will create execution routines. Like we will have some static Templates that the queries will fill the required informations and run them to get the bash scripts to be executed locally..




Remote driven push process will be created to run on the tmux session to analyse the problems...
--> later will be defined more..

a bash script will remotely create XYZ.sh files for simple iterations..