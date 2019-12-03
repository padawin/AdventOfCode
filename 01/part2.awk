NR == 1 {
	# Retrieve the number of lines in the current file
	# The file to process needs to be passed as argument to awk for this to
	# work, otherwise (eg. read from stdin), the filename will be incorrect
	( "cat " FILENAME " | wc -l" ) | getline NL
	# Initialises the current frequency and the seen frequencies
	c = 0
	seen[c] = 1
}

{
	c += $1
	if (c in seen) {
		print c;
		exit
	}
	else {
		seen[c] = 1
	}
}

# When we reach the end of the file
FNR == NL {
	#  Rewind up to 4 times (file parsed 5 times)
	while (1) {
		ARGC++
		ARGV[ARGIND+1] = FILENAME
		nextfile
	}

}
