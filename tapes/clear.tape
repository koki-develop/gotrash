# configuration
Output ./docs/clear.gif
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

Type "gotrash put hello.go" Sleep 500ms Enter
Sleep 1s
Type "gotrash list" Sleep 500ms Enter
Sleep 1s
Type "gotrash clear" Sleep 500ms Enter
Sleep 1s
Type "y" Sleep 500ms Enter
Sleep 1s
Type "gotrash list" Sleep 500ms Enter

Sleep 4s

# cleanup
Hide
Type "cd ~/" Enter
Type "\rm -rf /workspace/gotrash" Enter
Type "gotrash clear -f" Enter
