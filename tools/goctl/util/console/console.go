package console

type (
	Console interface {
		Success(format string, a ...any)
		Info(format string, a ...any)
		Debug(format string, a ...any)
		Warning(format string, a ...any)
		Error(format string, a ...any)
		Fatalln(format string, a ...any)
		MarkDone()
		Must(err error)
	}

	colorConsole struct {
		enable bool
	}

	ideaConsole struct{}
)

func NewConsole(idea bool) Console { _ = "STUB: not implemented"; return *new(Console) }

func NewColorConsole(enable ...bool) Console { _ = "STUB: not implemented"; return *new(Console) }

func (c *colorConsole) Info(format string, a ...any) { _ = "STUB: not implemented"; return }

func (c *colorConsole) Debug(format string, a ...any) { _ = "STUB: not implemented"; return }

func (c *colorConsole) Success(format string, a ...any) { _ = "STUB: not implemented"; return }

func (c *colorConsole) Warning(format string, a ...any) { _ = "STUB: not implemented"; return }

func (c *colorConsole) Error(format string, a ...any) { _ = "STUB: not implemented"; return }

func (c *colorConsole) Fatalln(format string, a ...any) { _ = "STUB: not implemented"; return }

func (c *colorConsole) MarkDone() { _ = "STUB: not implemented"; return }

func (c *colorConsole) Must(err error) { _ = "STUB: not implemented"; return }

func NewIdeaConsole() Console { _ = "STUB: not implemented"; return *new(Console) }

func (i *ideaConsole) Info(format string, a ...any) { _ = "STUB: not implemented"; return }

func (i *ideaConsole) Debug(format string, a ...any) { _ = "STUB: not implemented"; return }

func (i *ideaConsole) Success(format string, a ...any) { _ = "STUB: not implemented"; return }

func (i *ideaConsole) Warning(format string, a ...any) { _ = "STUB: not implemented"; return }

func (i *ideaConsole) Error(format string, a ...any) { _ = "STUB: not implemented"; return }

func (i *ideaConsole) Fatalln(format string, a ...any) { _ = "STUB: not implemented"; return }

func (i *ideaConsole) MarkDone() { _ = "STUB: not implemented"; return }

func (i *ideaConsole) Must(err error) { _ = "STUB: not implemented"; return }

func println(msg any) { _ = "STUB: not implemented"; return }

var defaultConsole = &colorConsole{enable: true}

func Success(format string, a ...any) { _ = "STUB: not implemented"; return }

func Info(format string, a ...any) { _ = "STUB: not implemented"; return }

func Debug(format string, a ...any) { _ = "STUB: not implemented"; return }

func Warning(format string, a ...any) { _ = "STUB: not implemented"; return }

func Error(format string, a ...any) { _ = "STUB: not implemented"; return }

func Fatalln(format string, a ...any) { _ = "STUB: not implemented"; return }

func MarkDone() { _ = "STUB: not implemented"; return }

func Must(err error) { _ = "STUB: not implemented"; return }
