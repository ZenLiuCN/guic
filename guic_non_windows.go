//go:build !windows
//+build !windows

package guic

func AttachConsole(){
  panic("not supported")
}
