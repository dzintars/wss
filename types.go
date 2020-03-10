
type Kind int

const (
  UI_LAUNCHER_SHOW = iota
  UI_LAUNCHER_HIDE
)

type Event struct {
  Type Kind
  User string
	Payload  interface{}
}


type App struct {
	// whatever your application state is
}

// Action is something that can operate on the application.
type Action interface {
	Run(app *App) error
}

type UiLauncherShowPayload struct {
	// ...
}

func (p *UiLauncherShowPayload) Run(app *App) error {
	// ...
}

type UiLauncherHidePayload struct {
	// ...
}

func (p *UiLauncherHidePayload) Run(app *App) error {
	// ...
}

var kindHandlers = map[Kind]func() Action{
	UI_LAUNCHER_SHOW:   func() Action { return &UiLauncherShowPayload{} },
	UI_LAUNCHER_HIDE: func() Action { return &UiLauncherHidePayload{} },
}

func main() {
	app := &App{
		// ...
	}

	// process an incoming message
	var raw json.RawMessage
	evt := Event{
		Msg: &raw,
	}
	if err := json.Unmarshal([]byte(input), &evt); err != nil {
		log.Fatal(err)
	}
	msg := kindHandlers[env.Type]()
	if err := json.Unmarshal(raw, msg); err != nil {
		log.Fatal(err)
	}
	if err := msg.Run(app); err != nil {
		// ...
	}
}
