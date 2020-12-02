#!/usr/bin/env raku
use v6;

my @lines = "input/day2.txt".IO.lines;
my $valid-pws = 0;
my $bvalid-pws = 0;
for @lines -> $line {
    $line ~~ /^ (\d+) \- (\d+) \s (.) \:\s (.*) $ /;
    #    say "match is $/";
    my ($lo, $hi, $letter, $pw) = $/.list;
    #    say "$lo -> $hi ($letter): $pw";
    my $count = $pw.comb($letter);
    $valid-pws++ if $lo <= $count and $count <= $hi;

    my $matches = 0;
    $matches++ if $pw.comb[$lo-1] eq $letter;
    $matches++ if $pw.comb[$hi-1] eq $letter;
    $bvalid-pws++ if $matches == 1;
}
say "{@lines.elems} lines $valid-pws valid $bvalid-pws bvalid";
