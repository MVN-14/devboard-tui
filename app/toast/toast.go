package toast

import (
	"strings"

	"github.com/MVN-14/devboard-tui/app/style"
)

type ToastType int

const (
	Success ToastType = iota
	Error
)

type Model struct {
	toast toast
}

type toast struct {
	msg       string
	toastType ToastType
}

func (m *Model) SetToast(msg string, t ToastType) {
	m.toast = toast{msg: msg, toastType: t}
}

func (m Model) Update() {
	if m.toast.msg == "" {
		return
	}
}

func (m Model) Render(v string) string {
	if m.toast.msg == "" {
		return v
	}

	s := style.ToastSuccess
	if m.toast.toastType == Error {
		s = style.ToastError
	}
	toast := s.Render(m.toast.msg)
	
	return overlay(v, toast)
}

func overlay(bg, fg string) string {
	bgLines := strings.Split(bg, "\n")
	fgLines := strings.Split(fg, "\n")
	
	for i, line := range fgLines {
		if i > len(bgLines) {
			bgLines = append(bgLines, line)
			continue
		}

		if len(bgLines[i]) > len(line) {
			newLine := line + bgLines[i][len(line):]
			bgLines[i] = newLine
		} else {
			bgLines[i] = line
		}
	}

	return strings.Join(bgLines, "\n")
}
