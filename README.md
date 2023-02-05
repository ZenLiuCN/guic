# guic
golang gui console helper for windows.
## use
```go
package main
import(
  "os"
  "github.com/ZenLiuCN/guic"
)
func main() {
	if len(os.Args) > 1 && os.Args[1] == "-c" {
		guic.AttachConsole() //make console enable
	}else{
		//do things with GUI
	}
	//DO things
}
```
build with `go build -ldflags "-H windowsgui"`, you made it: a GUI application with console support.
