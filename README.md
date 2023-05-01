Teact 🍵
========
Teact is a React-like abstraction built on top of Charm's [Bubbletea system](https://github.com/charmbracelet/bubbletea) that will make your TUIs easier to build, and responsive to terminal size. It's like HTML + CSS + your browser's layout engine, all-in-one for the terminal.

Basic Teact
-----------
### Teact Apps
Every Teact app starts with a call to `teact.Run` in its `main.go`, to the component that will be the root of your application. For example, this runs a Hello World application ([source code here](https://github.com/mieubrisse/teact/blob/main/demos/hello_world/main.go)):

```go
func main() {
	myApp := greeter.New()
	if _, err := teact.Run(myApp, tea.WithAltScreen()); err != nil {
		fmt.Printf("An error occurred running the program:\n%v", err)
		os.Exit(1)
	}
}
```

Teact apps can be quit by default with `ctrl-c` or `ctrl-d` (and this can be changed).

### Teact Components
A Teact component is just an implementation of the `Component` interface, and is analogous to an HTML element. It provides size & display information to Teact's layout/rendering system. 

However, 98% of the time you won't need to deal with any sizing because your custom components can be formed from [the default Teact components](https://github.com/mieubrisse/teact/tree/main/teact/components). You can think of the default Teact components like inbuilt HTML tags - `<p>`, `<div>`, `<li>`, etc. Your components should mostly be compositions of other components (just like in React).

For example, here's a `HelloWorldApp` component ([source code here](https://github.com/mieubrisse/teact/blob/main/demos/hello_world/greeter/greeter.go)) that's a composition of a flexbox containing styled text:

```go
// A custom component
type Greeter interface {
	components.Component
}

// Implementation of the custom component
type greeterImpl struct {
	// So long as we assign a component to this then our component will call down to it (via Go struct embedding)
	components.Component
}

func New() Greeter {
	// This is a tree, just like HTML, with leaf nodes indented the most
	root := flexbox.NewWithOpts(
		[]flexbox_item.FlexboxItem{
			flexbox_item.New(
				stylebox.New(
					text.New(text.WithContents("Hello, world!")),
					stylebox.WithStyle(
						style.WithForeground(lipgloss.Color("#B6DCFE")),
					),
				),
			),
		},
		flexbox.WithVerticalAlignment(flexbox.AlignCenter),
		flexbox.WithHorizontalAlignment(flexbox.AlignCenter),
	)

	return &greeterImpl{
		Component: root,
	}
}
```

Because the component has a `Component` struct embedded inside of it, `HelloWorldApp` fulfills the `Component` interface and Teact will know to use the embedded struct (which in this case is the flexbox) for rendering the `HelloWorldApp` component.

### Interactivity
Interactivity is accomplished by making a component implement the `InteractiveComponent` interface, which in turn uses the Bubbletea `Update` function. For example, this component keeps track of the number of keypresses it's seen and displays it ([source code here](https://github.com/mieubrisse/teact/blob/main/demos/keypress_counter/app/app.go):

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

You can see that each time the component receives a message, it checks if it's a keyboard message (since there are non-keyboard messages) and counts it.

### Utilities
Teact includes several utility functions ([under the `utilities` directory](https://github.com/mieubrisse/teact/tree/main/teact/utilities)) to make writing your components easier. Of note:

1. `NewStyle`, which allows building a `lipgloss.Style` object with the Go options pattern. For example:
    ```go
    NewStyle(
        WithBold(true),
        WithUnderline(true),
        WithBorder(lipgloss.Border()),
    )
    ```
1. `GetMaybeKeyMsgStr`, which is shorthand for testing if a `tea.Msg` is a `tea.KeyMsg`, and getting its string value if so.

### Testing
Teact includes [rudimentary component testing tools](https://github.com/mieubrisse/teact/tree/main/teact/component_test). These come in the form of assertions that are applied at various times in the component render loop (see below for more information on how this works). These are especially good if you're writing a new component from scratch (i.e. not embedding a `Component` in your impl `struct`).

Why not vanilla Bubbletea?
--------------------------
Bubbletea is a great foundation to build on, but it has several shortcomings that I hit when trying to build with it:

### No size communication from child to parent
In vanilla Bubbletea, parent components receive a simple `string` from child components via `tea.Model.View`. This means that the parent has no idea how to resize a given child's string - only the child knows how to render their `View` at the right size. 

The logical next step is a `Resize` method that cascades from child to parents, so that children are aware of the size they ought to be rendering at. However, this prevents a layout that responds to content: when a child grows of its own accord (say, it intercepted a keypress and added something to its width), a parent flexbox would need to resize the child's siblings. How does the child signify to the parent that it's wider and a recalculation needs to occur?

The way to do it in vanilla Bubbletea would be to have the child return a wider string. However, the parent might have preferences on how wide the child should be (e.g. to avoid overflowing the parent), so the parent might want to compress (perhaps by word-wrapping) the child text. But we know from earlier that a parent doesn't know how to resize a child's text - only a child knows that - so truly responsive layouts are impossible with vanilla Bubbletea.

Teact fixes this in the same way as your browser: doing a two-pass approach, where item preferred sizes are calculated first and then actual sizes are settled on using that information.

### By-value updating
`tea.Model.Update` returns a `tea.Model`. This means that a child Bubbletea component can either:

1. Implement `tea.Model`, but then its `Update` will return a `tea.Model` (thereby requiring the parent to cast it before storing the `Update` result)
2. Not actually implement `tea.Model` (which is what most of the components in [the Bubbles repo](https://github.com/charmbracelet/bubbles) do)

The by-value `Update` is also problematic when trying to create a generic component. For example, I was writing `FilterableList[T].Update`, with `T` being the element component that the list would contain. No matter how I tried, I couldn't get implementations of the `FilterableList[T]` interface to conform to the `Update(msg tea.Msg) T` function on the interface.

The by-value state transitioning is nice in theory (very Redux-y), but in practice I found it to be cumbersome so Teact only supports by-reference components.

### No first-class focus controls
The concept of "focusable component" is very useful and showed up in nearly all the example Bubbles, but it's not encoded in the BubbleTea framework in any way (all the example Bubbles recreate `Focus`, `Blur`, and `Focused` by hand).

### No flexbox
A resize of my terminal window should have each parent resizing their children (because the parent knows what size the children should be), but there was no out-of-the-box way for components to do this.

Best Practices
--------------
- 98% of the time, you should simply be assembling the [default Teact components](https://github.com/mieubrisse/teact/tree/main/teact/components) into a new component rather than writing the `View`, `GetContentMinMax`, etc. methods.
- Put each of your components in its own directory. This will help you stay organized.
- Give each component a public interface that implements either `Component` or `InteractiveComponent`. This will make it clear which type your component implements.
- Give each component a private implementation, built by a `New()` constructor. For example:

  In `my_component/my_component.go`
    ```go
    type Greeter interface {
        component.Component   // We know this isn't an interactive component
    }
    ```

  In `my_component/my_component_impl.go`
    ```go
    type greeterImpl struct {
        component.Component
    }

    func New() Greeter {
        root := text.New("Hello, world!")
        return &greeterImpl{
            Component: root,
        }
    }
    ```
- Embed a `Component` inside each private component implementation. This will transparently cause the `struct` to implement the `Component` interface, so that the rendering system will render however the embedded `Component` instance wants.
- To track subcomponents that your component needs to modify, store them as properties on your component (NOT replacing the embedded `Component`) and use them as needed. For example:
    ```go
    type greeterImpl struct {
        components.Component

        toUpdate text.Text
    }

    func New() Greeter {
        toUpdate := text.New(text.WithContents("Hello, World!"))
        root := flexbox.NewWithOpts(
            []flexbox_item.FlexboxItem{
                flexbox_item.New(
                    stylebox.New(
                        toUpdate, 
                        stylebox.WithStyle(
                            style.WithForeground(lipgloss.Color("#B6DCFE")),
                        ),
                    ),
                ),
            },
            flexbox.WithVerticalAlignment(flexbox.AlignCenter),
            flexbox.WithHorizontalAlignment(flexbox.AlignCenter),
        )

        return &greeterImpl{
            Component: root,
        }
    }

    func (impl *greeterImpl) UpdateGreeting(greeting string) Greeter {
        impl.toUpdate.SetContents(greeting)
    }
    ```
- When your component is configurable, use the [Go options pattern](https://michalzalecki.com/golang-options-pattern/) with a constructor like `New(opts ...MyComponentOpt)`. This will make it much easier to do the initial instantiation of your component, as all configuration for a component can be aligned visually. For comparison:
    ```go
    // If Teact components didn't have the Go optional pattern
    root := flexbox.New().SetChildren([]flexbox_item.FlexboxItem{
        flexbox_item.New().WithContent(
            stylebox.New(
                text.New().SetContent("Hello, world")
            ).SetStyle(someStyle)
        ).WithMaxWidth(flexbox_item.FixedSize(20)).WithVerticalGrowthFactor(1)
    }).SetHorizontalAlignment(flexbox.Center).SetVerticalAlignment(flexbox.Center)

    // With Go options pattern (notice how each component is an indentation level)
    root := flexbox.New(
        WithChildren(flexbox_item.FlexboxItem{
            flexbox_item.New(
                WithContent(
                    stylebox.New(
                        text.New(
                            WithContent("Hello, world")
                        )
                        WithStyle(someStyle),
                    )
                ),
                WithMaxWidth(flexbox_item.FixedSize(20)),
                WithVerticalGrowthFactor(1),
            ),
        }),
        WithHorizontalAlignment(flexbox.Center),
        WithVerticalAlignment(flexbox.Center),
    )
    ```
- Pass `tea.Msg` events solely from `InteractiveComponent` to `InteractiveComponent`. I.e., when an `InteractiveComponent` needs to pass a `tea.Msg` event downwards, have the parent's `Update` method pass the `tea.Msg` directly to the descendant that should receive it. Don't try to pass the event through a bunch of non-`InteractiveComponent`s (of which you will have many - `Flexbox`, `Stylebox`, `Text`, etc.).

<!--
Advanced Teact
--------------
- TODO embedding other components in
- TODO resizing based on viewport size
-->

How Teact Rendering Works
-------------------------
Teact rendering is a rudimentary version of what happens in your browser. Basically:

1. **X-Pass:** The minimum & desired widths & heights of each component in the graph is calculated (the equivalent of `min-content` and `max-content` in CSS), from bottom-to-top.
    - For those not familiar with CSS, components can have different sizes because word-wrapping can reduce the width of text (at a corresponding increase in height). The max width of a block of text is the length of its longest line without wrapping, and the min width is the length of the shortest word.
1. **Y-Pass:** Incorporating each component's desired width and the actual viewport width of your terminal, go top-to-bottom giving actual sizes to each component and calculating the components desired height given that width.
1. **Render:** Using all information, give an actual width to each component and render each component into a string to be displayed.

These three phases correspond to the three functions on the `Component` interface:

1. `GetContentMinMax()`
1. `SetWidthAndGetDesiredHeight(actualWidth)`
1. `View(actualWidth, actualHeight)`

Still TODO
----------
- Add Grid layout!!
- Make every component styleable, so we don't need styleboxes everywhere???
- Add some sort of inline/span thing
- Create a single "position" enum (so that we don't have different ones between flexbox and text, etc.)
- Make flexbox alignments purely "MainAxis" and "CrossAxis", so that when flipping the box things will be nicer
