#!/bin/bash

read line
lineChanged=1
while [ $lineChanged -eq 1 ]; do
	newLine=$(echo "$line" | sed -E -e 's/([a-z])/<\U\1\E>/g' -e 's/<([A-Z])>\1//g' -e 's/([A-Z])<\1>//g' -e 's/<([A-Z])>/\l\1\E/g')
	if [ "$newLine" == "$line" ]; then
		lineChanged=0
	else
		line="$newLine"
	fi
done
echo $newLine | wc -c
