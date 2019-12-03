#!/bin/sh
# use as:
# cat datafile | ./day9.1-clean.sh > cleaneddatafile
perl -p -e 's/!.//g' | egrep -o '<[^>]*>' | wc | awk '{print $3 - $2 - $1 - $1}'
