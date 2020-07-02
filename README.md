# HTML Link Parser
##### Project from [gophercises](https://gophercises.com) ([html link parser repo](https://github.com/gophercises/link))

In this exercise goal is create a package that makes it easy to parse an HTML file and extract all of the links (`<a href="">...</a>` tags). For each extracted link should return a data structure that includes both the `href`, as well as the text inside the link. Any HTML inside of the link can be stripped out, along with any extra whitespace including newlines, back-to-back spaces, etc.

For links like this

```html
<a href="/dog">
  <span>Something in a span</span>
  Text not in a span
  <b>Bold text!</b>
</a>
```

We want to get output that looks roughly like:

```go
Link {
  Href: "/dog",
  Text: "Something in a span Text not in a span Bold text!",
}
```