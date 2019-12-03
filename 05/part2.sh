#!/bin/bash

read line
minLength=$(echo $line | wc -c)
origLine=$line
for  c in a b c d e f g h i j k l m n o p q r s t u v w x y z; do
	line=$(echo $origLine | sed -e "s/$c//ig")
	lineChanged=1
	while [ $lineChanged -eq 1 ]; do
		newLine=$(echo "$line" | sed -E -e 's/([a-z])/<\U\1\E>/g' -e 's/<([A-Z])>\1//g' -e 's/([A-Z])<\1>//g' -e 's/<([A-Z])>/\l\1\E/g')
		if [ "$newLine" == "$line" ]; then
			lineChanged=0
		else
			line="$newLine"
		fi
	done
	newLen=$(echo $line | wc -c)
	echo "Without letter $c, new len is $newLen"
	if [ $newLen -lt $minLength ]; then
		minLength=$newLen
	fi
done
echo $minLength
