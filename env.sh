export PATH=/home/john/src/go/bin:$PATH
export GOPATH=/home/john/dev/jbert/aoc2020
alias go2="go tool go2go"
alias go2test="cp src/dummy.go src/aoc/dummy.go; cp src/dummy.go src/aoc/dummy_test.go; go tool go2go test aoc"
