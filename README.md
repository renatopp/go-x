# go-x

This is an attempt to extend the golang builtin package with more utilities while improving the consistence of its interface.

Packages:

- **`convx`** -- conversion utilities for primitive types. Replaces partially `strconv`.
- **`dsx`** -- data structures like `Heap`, `Queue`, `Stack`, etc. with python-like interface and indexing. Replaces `container/heap`.
- **`envx`** -- environment variable utilities, including loading as object. Replaces partially `os`.
- **`fmtx`** -- formatting utilities and terminal colors. Uses the formatting variation for print as default. Replaces `fmt`.
- **`fsx`** -- file system utilities. Replaces partially `os`, `io/ioutil`, `filepath` and more.
- **`httpx`** -- http utilities. Replaces partially `net/http`.
- **`iterx`** -- iterator utilities.
- **`jsonx`** -- json utilities. Replaces partially `encoding/json`.
- **`mapx`** -- map utilities. Replaces `maps`.
- **`mathx`** -- math utilities. Replaces `math`.
- **`randx`** -- random utilities. Replaces `math/rand`.
- **`runex`** -- runtime utilities. Replaces `runes`.
- **`slicex`** -- slice utilities. Replaces `slices`.
- **`strx`** -- string utilities. Adding lots of features, including tables and trees renderers. Replaces `strings` and partially `unicode`.
- **`testx`** -- testing utilities.
- **`timex`** -- time-based utilities, such as backoff and rate limiter functions.
- **`yamlx`** -- same as jsonx, but using `goccy/go-yaml`.
