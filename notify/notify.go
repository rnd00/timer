// notify pkg
// this package is made to executing notify-send command
package notify

import "os/exec"

type message struct {
	Title string
	Text  string
}

type Message interface {
	Notify() error
	GetTitle() string
	GetText() string
}

func New(title, txt string) Message {
	return &message{
		Title: title,
		Text:  txt,
	}
}

func (m *message) Notify() error {
	cmd := exec.Command("notify-send", m.Title, m.Text)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func (m *message) GetTitle() string {
	return m.Title
}

func (m *message) GetText() string {
	return m.Text
}
