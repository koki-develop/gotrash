# configuration
Output ./docs/clear.gif
Set Shell "bash"
Set FontSize 32
Set Width 1300
Set Height 600

# setup
Hide
Type "go install"
Type "gotrash --help"
Type "gotrash clear -f" Enter
Type "mkdir /tmp/gotrash" Enter
Type "cd /tmp/gotrash" Enter
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
Type "\rm -rf /tmp/gotrash" Enter
Type "gotrash clear -f" Enter
