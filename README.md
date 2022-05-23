<h1 align="center">
  <br>
  <a href="http://github.com/go-fn/fn"><img src="./docs/assets/function.png" alt="github.com/go-fn/fn" width="200px" /></a>
  <br>
  Fn
  <br>
</h1>

<p align="center">
  <a href="#introduction">Introduction</a> •
  <a href="#getting-started">Getting Started</a> •
  <a href="#contributing">Contributing</a> •
  <a href="#roadmap">Roadmap</a>
</p>

## Introduction

Fn is a general purpose automation tool that aims to be simpler and more flexible than similar tools such as: Make, Task, and Rake.

Fn aims to have a human-centered design, with emphasis on usability and aesthetics.

> Even the smartest among us can feel inept as we fail to figure out which light switch or oven burner to turn on, or whether to push, pull, or slide a door.
The fault lies not in ourselves, but in product design that ignores the needs of users and the principles of cognitive psychology. 
> 
> The problems range from ambiguous and hidden controls to arbitrary relationships between controls and functions, coupled with a lack of feedback or other assistance and unreasonable demands on memorization.

> The rules are simple: make things visible, exploit natural relationships that couple function and control, and make intelligent use of constraints. The goal: guide the user effortlessly to the right action on the right control at the right time.

` - The Design of Everyday Things`

An example to get you started:

> `fnfile.yml`
```yaml
version: '1'
fns:
  hello:
    do: echo "Hello, World!"
```

```shell
❯ fn hello
Hello, World
```

## Getting Started

```shell
go install github.com/go-fn/fn
```

## Features

### ✨ Functionality

|   Feature    | Explanation                                                                                                                                    |
|:------------:|:-----------------------------------------------------------------------------------------------------------------------------------------------|
|     `ui`     | An interactive UI that makes running `fn`'s easy and informative                                                                               |
| `templating` | Leverage Go Templates as an alternative to shell commands                                                                                      |
|  `include`   | Include content from local or remote `fnfile.yml` files.                                                                                       |
|  `outputs`   | Tired of reading walls of text? `Fn` supports a wide range of options for how output is presented, so you can find what you need, and move on. |
|   `watch`    | Integration with [viddy](https://github.com/sachaos/viddy) to watch a `fn`                                                                     |
|   `fwatch`   | Integration with [fsnotify](https://github.com/fsnotify/fsnotify) to provide "rerun" functionality when files change                           |
|  `src/gen`   | Fingerprint files between runs, useful to prevent unnecessary`fn` executions                                                                   |

### 🔑 Keywords

|    Keyword    | Does...                                                                     |
|:-------------:|:----------------------------------------------------------------------------|
|     `fn`      | A series of steps (aka function)                                            |
|    `step`     | A behavior                                                                  |
 |     `sh`      | A step that runs a `shell` command                                          |
|     `do`      | A step that runs other steps in sequence                                    |
|  `parallel`   | A step that runs other steps in parallel                                    |
|    `defer`    | Defer a `step` to be run when it's parent completes (success or failure)    |
|   `matrix`    | Dynamically defined steps                                                   |
|   `return`    | End early (such as when part of a `do`)                                     |
|     `var`     | A _lazily_ evaluated (then cached) value                                    |
|     `ctx`     | Ways to set or access information from different points of a `fn` or `step` |
|     `ns`      | Group `fn`'s by namespace                                                   |
| `serialgroup` | Control concurrency of `step`'s or `fn`'s through labelling                 |

## Contributing

TBD

## Roadmap

- [ ] Milestone 1 - Alpha
- [ ] Create Roadmap

## Attribution

<a href="https://www.flaticon.com/free-icons/function" title="function icons">Function icons created by Freepik - Flaticon</a>