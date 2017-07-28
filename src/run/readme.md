run
===

[github.com/dwitp/golang/src/run](https://github.com/dwitp/golang/tree/master/src/run)

Skip 2 commands when try to run code:

```bash
go install
{program}
```

# Features

- single command to run
- error compile termination
```bash
user@desktop:~/workspace/src/github.com/dwitp/golang/src/run$ run
exit status 2: # github.com/dwitp/golang/src/run
./app.go:10: imported and not used: "math"
```

# Installation

go install

# Running

run
