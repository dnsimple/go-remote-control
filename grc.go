package grc

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

const DEFAULT_SCRIPT_EXT = "sh"

var (
	scriptsDir = os.Getenv("SCRIPT_DIR")
	scriptExt  = os.Getenv("SCRIPT_EXT")
)

func RunStatus(w io.Writer) (err error) {
	return run("status", w)
}

func RunStart(w io.Writer) (err error) {
	return run("start", w)
}

func RunStop(w io.Writer) (err error) {
	return run("stop", w)
}

func RunUpdate(w io.Writer) (err error) {
	return run("update", w)
}

type flushWriter struct {
	f http.Flusher
	w io.Writer
}

func (fw *flushWriter) Write(p []byte) (n int, err error) {
	n, err = fw.w.Write(p)
	if fw.f != nil {
		fw.f.Flush()
	}
	return
}

func run(script string, w io.Writer) (err error) {
	fw := flushWriter{w: w}
	if f, ok := w.(http.Flusher); ok {
		fw.f = f
	}

	scriptPath := fmt.Sprintf("%s/%s.%s", scriptsDir, script, ext())
	log.Printf("Executing script: %s", scriptPath)
	cmd := exec.Command(scriptPath)
	cmd.Stdout = &fw
	cmd.Stderr = &fw
	err = cmd.Run()
	return
}

func ext() string {
	if scriptExt != "" {
		return scriptExt
	} else {
		return DEFAULT_SCRIPT_EXT
	}
}
