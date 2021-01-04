# Go-Mitizen
A Commitizen-like Git commit utility for written in Go. Used for standardizing commit messages.

## Compared to [cz-cli](https://github.com/commitizen/cz-cli)
- 👍 Starts significantly faster
- 👍 Doesn't require NPM
- 👍 Supports multiline bodies
- 👍 **TODO** Configurable within a repository without NPM or JavaScript
- 👎 Not as configurable
- 👎 Isn't compatible with the full commitizen toolset

## Installation
#### From Source
Download repository via Git or as a Zip
```
go build && ./go-mitizen --install
```

## Usage
```html
git cz [--install] [--log]
```
#### Flags
`--install`: Install the application to be runnable via `git cz`.

`--log`: Write log messages to the *logs.txt* file in the working directory for debugging.

## Configuration
**TODO**