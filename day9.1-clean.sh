#!/bin/sh
# use as:
# cat datafile | ./day9.1-clean.sh > cleaneddatafile
perl -p -e 's/!.//g' | perl -p -e 's/<[^>]*>//g'
