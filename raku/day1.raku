use v6;

my $nums = "input/day1.txt".IO.lines.map(*.Int).Set;

my $target = 2020;
NUM: for $nums.keys -> $num {
    my $needed = $target - $num;
    if $nums{$needed} {
        say "$num + $needed = {$num+$needed} - product {$num*$needed}";
        last NUM;
    }
}

A: for $nums.keys -> $a {
    B: for $nums.keys -> $b {
        my $needed = $target - $a - $b;
        if $nums{$needed} {
            say "$a + $b + $needed = {$a+$b+$needed} - product {$a*$b*$needed}";
            last A;
        }
    }
}
