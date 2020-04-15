#!/bin/bash
#poem.sh
#select and print on the terminal a poem from the file $HOME/poems.txt
#
POEMS="$HOME/poems.txt"
#
#Find out how many poems there are by setting a variable with the output of the
#grep program. Grep looks for patterns in files. The -c returns a count
#We look to see how many %%s there are
#The back quote (`) is used to tell the shell to run the program.
#
NUMPOEMS=`grep -c "%%" $POEMS`
#
ANUMBER=0
#
#Read and Check that the number is in range and is valid
#Checking input is good practice in shell scripting to avoid problems that
#may stop your program from working
#
#$ANUMBER must be greater than 0 
#and $ANUMBER must be less than or equal to $NUMPOEMS
#Set a loop control variable
GOON=0
while [ $GOON -lt 1 ]
do
	echo "Enter a number between 1 and $NUMPOEMS"
	read ANUMBER
	#check the size of the input is greater than 0
	if [ ${#ANUMBER} -lt 1 ]
	then
		#The construct ${#VariableName} returns the size of the Variable.
		#Obviously if it is less than 1 the user must have 
		#pressed Enter without typing a number so go round the loop again
		continue
	fi
	if [ "$ANUMBER" -gt 0 -a "$ANUMBER" -le "$NUMPOEMS" ]
	then
		#Got a number in the range so set GOON to exit the loop
		GOON=1
	else
		#Out of range so go round loop again
		continue
	fi
done
#
#Now use $ANUMBER to select a poem
#So we first find line in the poem file until the next %% pattern.
#for this we shall use a nifty program called awk
#awk finds patterns in files and performs actions. Awk has its own way of 
#doing things which are different to the shell. Actions are enclosed in
#a single quote ('). each action is enclosed in curly braces ({}) The BEGIN 
#action below tells awk that %% is the record separator. The if tests for the 
#record number (NR) equal to ANUM and only prints each line of that record
#Notice that awk uses == to do its equals tests
#we also have to tell awk about the $ANUMBER variable with -v
#
awk -v ANUM=$ANUMBER 'BEGIN { RS = "%%" }{if( NR == ANUM ) print $0 }' $POEMS