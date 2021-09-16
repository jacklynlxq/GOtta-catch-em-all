# GOtta-catch-em-all
- [Package Visibility](#package-visibility)
- [Naming Convention](#naming-convention)
- [Godoc](#godoc)

### Go Cheatsheet
#### Package Visibility
1. Exported identifier (akin to public access modifier) begins with capital letter, eg: `Foobar()`.

2. Internal identifier (akin to private access modifier) begins with small letter, eg: `foobar()`.

Identifer means variable name, function name, constant, statement labels, package name, or types.

#### Naming Convention: 
`MixedCaps` or `mixedCaps` rather than underscores to write multiword names.

#### Godoc
Formatting rules by official Go doc:
- Subsequent lines of text are considered part of the same paragraph; you must leave a blank line to separate paragraphs.
- Pre-formatted text must be indented relative to the surrounding comment text (see gob’s [doc.go](https://golang.org/src/encoding/gob/doc.go) for an example).
- URLs will be converted to HTML links; no special markup is necessary.

[Official Godoc Guide](https://go.dev/blog/godoc) 



#### Reference
- [Official Go Documentation](https://golang.org/doc)
- [An Introduction to Programming in Go by Caleb Doxsey](https://www.golang-book.com/books/intro)
  - Note: This open source version was released in 2012, hence there might be some functions that have been deprecated but I think it's good enough to get yourself started on familiarising with Go's concept.