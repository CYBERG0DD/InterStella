# General understanding

* We use cd ../.. to go two folders back instead of cd . to go back once
* cd - to jump straight into your previous working directory
* ">" To redirects code outputs to a file instead of terminal output careful usage cause it overwrites an existing file
* ">>" To append to an exsting the the outputs of your code
* 2> For redirecting errors if you want.

### IN SUMMARY

**stdin (fd 0)** - Stream a program reads input from. Defaults to keyboard, can be redirected from a file.

**stdout (fd 1)** - Stream a program's normal output goes to. Defaults to terminal screen.

**stderr (fd 2)** - Stream a program's error messages go to. Also defaults to terminal screen — separate from exit codes.

**Exit Code (Exit Status)** - A number (0-255) a program returns when it finishes. `0` = success, always. Non-zero = failure, meaning defined per-program.


**File Descriptor (fd)** - A number the OS assigns to any open input/output channel for a running program (a real file, the terminal, a pipe, a network connection). The program uses the number to reference it instead of re-specifying the whole resource each time.
→ Every command run in Git Bash gets fd 0 (stdin), fd 1 (stdout), fd 2 (stderr) auto-wired before it runs anything itself. Anything opened after that gets fd 3, 4, 5...
