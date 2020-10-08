## github.com/eculver/go-play/pkg/confirmer

A simple confirmation library for blocking on user confirmation prompts.

### Installation

```
go get -u github.com/eculver/go-play/pkg/confirmer
```

### Usage

The `confirmer.Confirm` method provides a default configuration to cover the most common use-case:

```
import (
    "os"
    "github.com/eculver/go-play/pkg/confirmer"
)

if !confirmer.Confirm("Continue?", os.Stdin) {
    // not confirmed
    os.Exit(1)
}
// confirmed
```

This will prompt the user up to three times until a valid "accept" or "deny" input is given.
If after three times, no valid inputs have been given, it will exit and return `false`:

```
evan.culver@evan ➜  go-play (confirmer) go run cmd/confirmer/main.go
Continue? [Y/n]: not valid

Continue? [Y/n]: really?

Continue? [Y/n]: yes

not confirmed!
exit status 1
```
```
evan.culver@evan ➜  go-play (confirmer) go run cmd/confirmer/main.go
Continue? [Y/n]: Y
confirmed!
```
```
evan.culver@evan ➜  go-play (confirmer) go run cmd/confirmer/main.go
Continue? [Y/n]: N
Continue? [Y/n]: n
not confirmed!
exit status 1
```

A complete working example can be found in [`./cmd/confirmer/main.go`](../../cmd/confirmer/main.go) of this repository.
