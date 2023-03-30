# configuration
Output ./docs/restore-ui.gif
Set Shell "bash"
Set FontSize 32
Set Width 1300
Set Height 600

# setup
Hide
Type "go install" Enter
Sleep 10s
Type "gotrash --help" Enter
Type "gotrash clear -f" Enter
Type "mkdir /workspace/gotrash" Enter
Type "cd /workspace/gotrash" Enter
Type "touch hello.go world.ts foo.js bar.rb" Enter
Ctrl+l
Show

# ---

Type "ls" Sleep 500ms Enter
Sleep 1s
Type "gotrash put hello.go world.ts foo.js bar.rb" Sleep 500ms Enter
Sleep 1s
Type "gotrash restore" Sleep 500ms Enter
Sleep 2s
Type "hello.go" Sleep 500ms Tab
Sleep 1s
Backspace 8
Sleep 1s
Type "foo.js" Sleep 500ms Tab
Sleep 1s
Backspace 6
Sleep 2s
Enter
Sleep 1s
Type "ls" Sleep 500ms Enter

Sleep 4s

# cleanup
Hide
Type "cd ~/" Enter
Type "\rm -rf /workspace/gotrash" Enter
Type "gotrash clear -f" Enter
