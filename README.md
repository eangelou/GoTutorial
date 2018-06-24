# GoTutorial
A repository for various small go exercises

## Version struct
See main package. Use:

$ ./main -version 1.1.1
1.2.3
{"version":"1.2.3"}
{"version":"1.1.1"}
{"version":"0.0.0"}

$ ./main -major 1 -minor=1 -patch 1
1.2.3
{"version":"1.2.3"}
{"version":"1.1.1"}
{"version":"0.0.0"}

$ go build -ldflags "-X main.hardcoded=1.0.1" main.go 
$ ./main -version 1.1.1
1.2.3
{"version":"1.2.3"}
{"version":"1.1.1"}
{"version":"1.0.1"}

