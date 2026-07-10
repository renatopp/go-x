package fmtx

import (
	"fmt"
	"io"
)

// Print is an alias to fmt.Printf.
func Print(msg string, v ...any) {
	fmt.Printf(msg, v...)
}

// Println is an alias to fmt.Printf with a newline at the end.
func Println(msg string, v ...any) {
	fmt.Printf(msg+"\n", v...)
}

// Sprintf is an alias to fmt.Sprintf.
func Sprint(msg string, v ...any) string {
	return fmt.Sprintf(msg, v...)
}

// Sprintln is an alias to fmt.Sprintf with a newline at the end.
func Sprintln(msg string, v ...any) string {
	return fmt.Sprintf(msg+"\n", v...)
}

// FPrint is an alias to fmt.Fprintf, using the given writer.
func FPrint(w io.Writer, msg string, v ...any) (n int, err error) {
	return fmt.Fprintf(w, msg, v...)
}

// FPrintln is an alias to fmt.Fprintf using the given writer with a newline at
// the end.
func FPrintln(w io.Writer, msg string, v ...any) (n int, err error) {
	return fmt.Fprintf(w, msg+"\n", v...)
}

// Fatal is an alias to panic with a formatted message.
func Fatal(msg string, v ...any) {
	panic(fmt.Sprintf(msg, v...))
}

// Error is an alias to fmt.Errorf.
func Error(msg string, v ...any) error {
	return fmt.Errorf(msg, v...)
}

// Scan is an alias to fmt.Scanf.
func Scan(format string, v ...any) (n int, err error) {
	return fmt.Scanf(format, v...)
}

// Scanln is an alias to fmt.Scanln.
func Scanln(v ...any) (n int, err error) {
	return fmt.Scanln(v...)
}

// Sscan is an alias to fmt.Sscanf.
func Sscan(str, format string, v ...any) (n int, err error) {
	return fmt.Sscanf(str, format, v...)
}

// Sscanln is an alias to fmt.Sscanln.
func Sscanln(str string, v ...any) (n int, err error) {
	return fmt.Sscanln(str, v...)
}

// Fscan is an alias to fmt.Fscanf, using the given reader.
func Fscan(r io.Reader, format string, v ...any) (n int, err error) {
	return fmt.Fscanf(r, format, v...)
}

// Fscanln is an alias to fmt.Fscanln, using the given reader.
func Fscanln(r io.Reader, v ...any) (n int, err error) {
	return fmt.Fscanln(r, v...)
}
