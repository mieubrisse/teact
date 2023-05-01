Teact 🍵
========
Teact is a React-like framework built on top of Charm's [Bubbletea system](https://github.com/charmbracelet/bubbletea) that will make your TUIs easier to build, and responsive to terminal size. It's like HTML + CSS + your browser's layout engine, all-in-one for the terminal.

Basic Teact
-----------
### Teact Apps
Every Teact app starts with a call to `teact.Run` in its `main.go`, to the component that will be the root of your application. For example, this runs a Hello World application ([source code here](https://github.com/mieubrisse/teact/blob/main/demos/hello_world/main.go)):

```go
func main() {
	myApp := app.New()
	if _, err := teact.Run(myApp, tea.WithAltScreen()); err != nil {
		fmt.Printf("An error occurred running the program:\n%v", err)
		os.Exit(1)
	}
}
```

Teact apps can be quit by default with `ctrl-c` or `ctrl-d` (and this can be changed).

### Teact Components
A Teact component is just an implementation of the `Component` interface, which gives sizing information to Teact's layout/rendering system. You don't need to do any size calculations though, because Teact ships with several common components (e.g. flexbox, text, input field) so you can just compose them together to form your components. For example, here's the Hello World app component ([source code here](https://github.com/mieubrisse/teact/blob/main/demos/hello_world/app/app.go)):

```go
// A custom component
type HelloWorldApp interface {
	components.Component
}

// Implementation of the custom component
type helloWorldAppImpl struct {
	// So long as we assign a component to this then our component will call down to it (via Go struct embedding)
	components.Component
}

func New() HelloWorldApp {
    // This is a tree, just like HTML, with leaf nodes indented the most
	root := flexbox.NewWithOpts(
		[]flexbox_item.FlexboxItem{
			flexbox_item.New(
				stylebox.New(
					text.New("Hello, world!"),
					stylebox.WithStyle(
						style.WithForeground(lipgloss.Color("#B6DCFE")),
					),
				),
			),
		},
		flexbox.WithVerticalAlignment(flexbox.AlignCenter),
		flexbox.WithHorizontalAlignment(flexbox.AlignCenter),
	)

	return &helloWorldAppImpl{
		Component: root,
	}
}
```

As long a component is assigned to the `Component` inner struct field, Teact will know how to render the object (via Go's struct embedding). In this way, you can simply build your custom components from preexisting components.

### Interactivity
Interactivity is accomplished by making a component implement the `InteractiveComponent` interface, which in turn uses the Bubbletea `Update` function. For example, this component keeps track of the number of keypresses it's seen and displays it ([source code here]():

```go
type KeypressCounter interface {
	components.InteractiveComponent
}

type keypressCounterImpl struct {
	components.Component

	keysPressed int
	output      text.Text
}

func New() KeypressCounter {
	output := text.New()
	result := &keypressCounterImpl{
		Component:   output,
		keysPressed: 0,
		output:      output,
	}
	result.updateOutputText()
	return result
}

func (k *keypressCounterImpl) Update(msg tea.Msg) tea.Cmd {
	if utilities.GetMaybeKeyMsgStr(msg) != "" {
		k.keysPressed += 1
		k.updateOutputText()
	}
	return nil
}

func (k keypressCounterImpl) SetFocus(isFocused bool) tea.Cmd {
	return nil
}

func (k keypressCounterImpl) IsFocused() bool {
	return true
}

func (b *keypressCounterImpl) updateOutputText() {
	b.output.SetContents(fmt.Sprintf("You've pressed %v keys", b.keysPressed))
}
```

You can see that 

Best Practices
--------------
Example non-interactive component:
```go
type MyCustomComponent interface {
    components.Component

    GetProperty() string
    SetProperty(val string) MyCustomComponent   // It's usually nice to make these fluent, but not required
}

type impl MyCustomComponent {
    property string
}

func (im *impl) GetProperty() string {
    return im.property
}

func (im *impl) SetProperty(val string) MyCustomComponent {
    im.property = val
    return im
}
```

- TODO add `XXXOpts` and create a `New(opts ...XXXOpts)` constructor, which allow for a very clean, hierarchical-nested declaration style:
  ```go
  flexbox_item.New(
      stylebox.New(
          text.New("Form", text.WithAlign(text.AlignCenter)),
          stylebox.WithStyle(style.NewStyle(
              style.WithBold(true),
              style.WithBorder(lipgloss.NormalBorder()),
          )),
      ),
      flexbox_item.WithHorizontalGrowthFactor(1),
  ),
  ```

- Write custom components with an interface that implements `Component` (or `InteractiveComponent` if your component is interactive), a private implementation, and a constructor. For example:
- Create an interface for your custom component. This will help you separate the external-facing functionality from the implementation, and keep your components clean.
- Use struct embedding to embed a `component.Component` inside your custom components, like so:
  ```go
  type myCustomComponentImpl struct {
    component.Component // <---- embedded struct
  }
  ```
  This embedded struct will serve as your root display element. Assign to it once in your
### 

- Write custom components that compose the default Teact components. Most of the time, your 
- Most of the 
- Embed a `component.Component` in every custom component; use this as the "root" for rendering
- For any components you want to manipulate, store them as properties on the component
- `Component` does not receive focus. Instead, your `InteractiveComponent`s should, in their `Update`, control and pass to the children they want
- TODO something about `XXXXOpt` for components that contain other components?





Teact provides several built-in components (e.g. flexbox, list). You'll write components that compose these components (which can be composed even further).

You'll write custom components that compose. Teact provides several 


- Parents should resize based on the sizes of their children
- 




It has several intents

TODO:
- Fix word wrapping
- Add Flexbox
- Add height

TODO explain the utility wrappers around `lipgloss.NewStyle()`

TODO note about `RunTeact`

TODO `TryUpdate`

TODO note about how putting a border around a thing is a good way to see what it's doing

Why not vanilla Bubbletea?
--------------------------


TODO
====
- Make every component styleable, so we don't need styleboxes everywhere???
- Add some sort of inline/span thing
- Create a single "position" enum (so that we don't have different ones between flexbox and text, etc.)
- Make flexbox alignments purely "MainAxis" and "CrossAxis", so that when flipping the box things will be nicer
- Add Grid layout!!

How Teact works
---------------
