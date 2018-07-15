# replacestrings
A command line tool to replace strings

This is a very simple program to replace multiple strings in a single file.
"Platform independent" in that it's written in plain/simple go, so theoretically
usable on any platform that supports go.

To install:

```
go get -u github.com/pdk/replacestrings
```

To use:

```
replacestrings -in /path/to/inputfile -out /path/to/output oldstring1 newstring1 oldstring2 newstring2 ...
```

The program operates by reading the input line by line and replacing all
`oldstringN` with `newstringN`. Note that if a later `newstring` matches a
previous `oldstring` those will get replaced, too.