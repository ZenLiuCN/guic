//go:build windows
//+build windows
package guic
import (
	"os"
	s "syscall"
)

// https://www.tillett.info/2013/05/13/how-to-create-a-windows-program-that-works-as-both-as-a-gui-and-console-application/
const attachParentProcess = ^uintptr(0)

func AttachConsole() {
	aOut := attachable(s.STD_OUTPUT_HANDLE)
	aIn := attachable(s.STD_INPUT_HANDLE)
	aErr := attachable(s.STD_ERROR_HANDLE)
	ac := false
	kernel32 := s.MustLoadDLL("kernel32.dll")
	if aErr && aOut {
		if r, _, _ := kernel32.MustFindProc("AttachConsole").Call(attachParentProcess); r != 0 {
			ac = true
		} else if r, _, _ = kernel32.MustFindProc("AllocConsole").Call(); r != 0 {
			ac = true
		}
	}
	if ac {
		if aOut {
			os.Stdout, _ = os.OpenFile("CONOUT$", os.O_WRONLY, 0666)

		}
		if aErr {
			os.Stderr, _ = os.OpenFile("CONOUT$", os.O_WRONLY, 0666)
		}
		if aIn {
			os.Stderr, _ = os.OpenFile("CONIN$", os.O_RDONLY, 0666)
		}
	}

}
func attachable(std int) bool {
	if hOut, _ := s.GetStdHandle(std); hOut != 0 {
		tOut, _ := s.GetFileType(hOut)
		if tOut == s.FILE_TYPE_DISK || tOut == s.FILE_TYPE_PIPE {
			return false
		}
	}
	return true
}
