# grep-tool

> This is a simple version of the Unix tool `grep`

*% man grep*
	The grep utility searches any given input files, selecting lines that
     match one or more patterns.  By default, a pattern matches an input line
     if the regular expression (RE) in the pattern matches the input line
     without its trailing newline.  An empty expression matches every line.
     Each input line that matches at least one of the patterns is written to
     the standard output.

`go run main.go "" test.txt | diff test.txt -`

`go run main.go J rockbands.txt`

`go run main.go -r "Nirvana" "*"`

`go run main.go -r Nirvana * | go run main.go -v Madonna`

`go run main.go "\d" test-subdir/BFS1985.txt`

`go run main.go "\w" symbols.txt`
