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

Fn is a "function-oriented" general purpose automation tool that aims to be simpler and more flexible than similar tools such as Make, Task, and Rake.

Fn aims to have a human-centered design, with emphasis on usability and aesthetics.

> Even the smartest among us can feel inept as we fail to figure out which light switch or oven burner to turn on, or whether to push, pull, or slide a door.
The fault lies not in ourselves, but in product design that ignores the needs of users and the principles of cognitive psychology. 
> 
> The problems range from ambiguous and hidden controls to arbitrary relationships between controls and functions, coupled with a lack of feedback or other assistance and unreasonable demands on memorization.
>
> The rules are simple: make things visible, exploit natural relationships that couple function and control, and make intelligent use of constraints.
>
> The goal: guide the user effortlessly to the right action on the right control at the right time. 
>
> -- <cite>[The Design of Everyday Things](https://www.uxmatters.com/mt/archives/2021/03/book-review-the-design-of-everyday-things.php)</cite>

An example to get you started:

> `fnfile.yml`
```yaml
version: '0.1'
fns:
  hello: echo "Hello, World!"
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

|   Feature    | Explanation                                                                                                                                          |
|:------------:|:-----------------------------------------------------------------------------------------------------------------------------------------------------|
|     `ui`     | An interactive UI that makes running `fn` easy, fun, and informative                                                                                 |
|    `cli`     | POSIX compliant cli for use with automation                                                                                                          | 
| `templating` | Leverage Go Templates as an alternative to shell commands                                                                                            |
|  `include`   | Include content from local or remote `fnfile.yml` files.                                                                                             |
|  `outputs`   | Tired of reading walls of log output? `Fn` supports a wide range of options for how output is presented, so you can find what you need, and move on. |
|   `watch`    | Integration with [viddy](https://github.com/sachaos/viddy) to watch a `fn`                                                                           |
|   `fwatch`   | Integration with [fsnotify](https://github.com/fsnotify/fsnotify) to provide "rerun" functionality when files change                                 |
|  `src/gen`   | Fingerprint files between runs, useful to skip expensive steps                                                                                       |

### 🔑 Keywords

|    Keyword    | Description                                                                       |
|:-------------:|:----------------------------------------------------------------------------------|
|     `fn`      | A series of steps (aka function)                                                  |
|    `step`     | An abstract behavior                                                              |
 |     `sh`      | A `step` that runs a `shell` command                                              |
|     `do`      | A `step` that runs other steps in sequence                                        |
|  `parallel`   | A `step` that runs other steps in parallel                                        |
|    `defer`    | A `step` that is run when the surround `fn` completes (success or failure)        |
|   `matrix`    | Dynamically defined steps                                                         |
|   `return`    | A `step` that triggers the parent step to end early (such as when part of a `do`) |
|     `var`     | A _lazily_ evaluated (then optionally memoized/cached) value                      |
|     `ctx`     | A collection of values that are available at different points in an `fnfile`      |
|     `ns`      | A namespace for `fn`'s                                                            |
| `serialgroup` | Control concurrency of `step`'s or `fn`'s through labelling                       |

## Contributing

TBD

## Roadmap

| Status | Goal                                                   | Labels  |
|:------:|:-------------------------------------------------------|---------|
|   ✅    | Create RoadMap                                         |         |
|   🚧   | Implement Basic Step Types                             | `alpha` |
|   ✅    | *- Sh (shell out)*                                     | `alpha` |
|   ✅    | *- Do (serial steps)*                                  | `alpha` |
|   ✅    | *- Parallel (parallel steps)*                          | `alpha` |
|   ✅    | *- Defer (end of parent)*                              | `alpha` |
|   🚧   | *- Fn (call another function)*                         | `alpha` |
|   🚧   | Vars Support                                           | `alpha` |
|   🚧   | Env Support                                            | `alpha` |
|   🚧   | Ctx Support                                            | `alpha` |
|   🚧   | Templating Support                                     | `alpha` |
|   -    | CLI Args Support                                       | `alpha` |
|   -    | Linear Output (log-like)                               | `alpha` |
|   -    | Progress/Spinner Output                                | `alpha` |
|   -    | Robust Documentation (follows the code)                | `alpha` |
|   -    | BubbleTea based UI                                     | `beta`  |
|   -    | Viddy/Watch Integration                                | `beta`  |
|   -    | FSNotify Integration                                   | `beta`  |
|   -    | Implement Advanced Step Types                          | `beta`  |
|   -    | *- Matrix (dynamic steps based on combination matrix)* | `beta`  |
|   -    | *- Return (end early)*                                 | `beta`  |
|   -    | Other Advanced Features                                | `beta`  |
|   -    | *- Ns (namespaces/includes)*                           | `beta`  |
|   -    | *- SerialGroups (for use with parallel)*               | `beta`  |
|   -    | *- Inputs/Outputs (for functions)*                     | `beta`  |


## Attribution

<a href="https://www.flaticon.com/free-icons/function" title="function icons">Function icons created by Freepik - Flaticon</a>