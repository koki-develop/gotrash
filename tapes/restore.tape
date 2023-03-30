# configuration
Output ./docs/restore.gif
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
Type "mkdir /tmp/gotrash" Enter
Type "cd /tmp/gotrash" Enter
Type "touch hello.go world.ts foo.js bar.rb" Enter
Ctrl+l
Show

# ---

Type "ls" Sleep 500ms Enter
Sleep 1s
Type "gotrash put hello.go world.ts" Sleep 500ms Enter
Sleep 1s
Type "ls" Sleep 500ms Enter
Sleep 1s
Type "gotrash list" Sleep 500ms Enter
Sleep 2s
Type "gotrash restore 1" Sleep 500ms Enter
Sleep 1s
Type "ls" Sleep 500ms Enter

Sleep 4s

# cleanup
Hide
Type "cd ~/" Enter
Type "\rm -rf /tmp/gotrash" Enter
Type "gotrash clear -f" Enter
