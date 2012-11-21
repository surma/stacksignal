`stacksignal` registers a signal handler for SIGUSR1.
Whenever the importing program receives SIGUSR1, a full
stacktrace of all goroutines will be printed to `stderr`.

Since there are no functions to use, the import line should
be:

```Go
import _ "github.com/surma/stacksignal"
```
---
Version 1.1.1
