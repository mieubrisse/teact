Teact
=====
Teact is a React-like framework built on top of Charm's [Bubbletea system](https://github.com/charmbracelet/bubbletea). Components built with Teact are reactive by default.

How to use it
-------------
You write components out of other components. Teact provides several 


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

TODO
====
- Make every component styleable, so we don't need styleboxes everywhere???
- Add some sort of inline/span thing
- Create a single "position" enum (so that we don't have different ones between flexbox and text, etc.)
- Make flexbox alignments purely "MainAxis" and "CrossAxis", so that when flipping the box things will be nicer
- Add Grid layout!!
