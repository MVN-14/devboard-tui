package toast

import (
	"strings"
	"time"

	"github.com/MVN-14/devboard-tui/app/style"
)

const duration = time.Second * 5

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
	start	  time.Time
}

func (m *Model) SetToast(msg string, t ToastType) {
	m.toast = toast{msg: msg, toastType: t, start: time.Now()}
}

func (m *Model) Update() {
	if m.toast.msg == "" {
		return
	}

	if time.Since(m.toast.start) > duration {
		m.toast = toast{}
	}
}

func (m Model) Render(v string, w int) string {
	if m.toast.msg == "" {
		return v
	}

	s := style.ToastSuccess.Width(w)
	if m.toast.toastType == Error {
		s = style.ToastError.Width(w)
	}
	toast := s.Render(m.toast.msg)

	return overlay(v, toast)
}

func overlay(bg, fg string) string {
	bgLines := strings.Split(bg, "\n")
	fgLines := strings.Split(fg, "\n")
	
	copy(bgLines, fgLines)

	return strings.Join(bgLines, "\n")
}
