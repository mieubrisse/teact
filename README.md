Teact
=====
Teact is a React-like framework built on top of Charm's Bubbletea system.

TODO:
- Fix word wrapping
- Add Flexbox
- Add height

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
