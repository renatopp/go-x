# go-x

This is an attempt to extend the golang builtin package with more utilities while improving the consistence of its interface.

Packages:

- **`convx`** -- conversion utilities for primitive types. Replaces partially `strconv`.
- **`dsx`** -- data structures like `Heap`, `Queue`, `Stack`, etc. with python-like interface and indexing. Replaces `container/heap`.
- **`fmtx`** -- formatting utilities and terminal colors. Uses the formatting variation for print as default. Replaces `fmt`.
- **`fsx`** -- file system utilities. Replaces partially `os`, `io/ioutil`, `filepath` and more.
- **`httpx`** -- http utilities. Replaces partially `net/http`.
