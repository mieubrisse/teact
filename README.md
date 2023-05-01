TODO:
- Fix word wrapping
- Add Flexbox
- Add height

Best practices for custom components:
- Embed a `component.Component` in every custom component; use this as the "root" for rendering
- For any components you want to manipulate, store them as properties on the component
- `Component` does not receive focus. Instead, your `InteractiveComponent`s should, in their `Update`, control and pass to the children they want
