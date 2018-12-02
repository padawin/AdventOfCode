#!/usr/bin/perl

my @candidates;
my $len_lines = 0;
my $count_candidates = 0;
foreach $line (<>) {
	if ($len_lines == 0) {
		$len_lines = length $line;
	}
	for $line2 (@candidates) {
		my $differences = 0;
		my @common;
		for ($c = 0; $c < $len_lines && $differences < 2; $c += 1) {
			my $char1 = substr($line, $c, 1);
			my $char2 = substr($line2, $c, 1);
			if ($char1 ne $char2) {
				$differences += 1;
			}
			else {
				push @common, $char1;
			}
		}

		if ($differences == 1) {
			print join '', @common;
			exit;
		}
	}
	push @candidates, $line;
}
