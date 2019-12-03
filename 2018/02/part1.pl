#!/usr/bin/perl

my $doubles = 0;
my $triples = 0;
foreach $line (<>) {
	# Sort the string
	my $sorted_line = join('', sort split //, $line);
	# Remove tripled letters
	if ($sorted_line =~ s/([a-z])\1{2}//g) {
		$triples += 1;
	}
	# Remove doubled letters
	if ($sorted_line =~ s/([a-z])\1{1}//g) {
		$doubles += 1;
	}
}
print "$doubles doubles and $triples triples, result is: ";
print $triples * $doubles;
print "\n";
