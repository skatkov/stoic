# Stoic
`stoic` is a command line app for daily journaling with plain-text files.

## Motivation
There is a recurring theme through biographies of great people - they all had journaling as a hobby.

I've been battling my inner demons with different methods in the past with varying success. But journaling has helped me keep these demons permanently at bay. My sleep will be peaceful, if it follows after careful and honest self-examination in my journal. 

Existing software for journaling and note taking is too slow to load and filled with features I don'
t really need. Plain-text files stored in cloud storage and edited througn nano is more than enough. But some recurring manual work was still required - create new daily file and  modify according to my template. This command line utility completely removes that manual work for me. 

Epictetus, great Stoic philospher and slave, once told to his students that "philosophy is something one should write down day by day". Hence name of this tool is a reference to this great human and hat tip to practical philosophy called Stoicism.

## Brief
Software helps maintain day-to-day journaling habbit by allowing you to:
- Instantly edit current day entry with one command in terminal
- Journal with any plain-text format
- Use templates for journal entries
- Works on Linux and OSX. 

## Installation

📥 [Installation](INSTALL.md)

## Usage
In a terminal of your chosing type following commands:

- `stoic`
- 
Will open todays journal entry in editor

- `stoic about`
- 
Prints out information about cli app


## Configration

- Editor could be changed by setting $EDITOR global variable. (default: nano)
- Directory for journal could be changed by setting $STOIC_DIR global variable. (default: ~/Journal/)
- Provide file template through $STOIC_TEMPLATE global variable.
- Provide new extension format through $STOIC_EXT global variables. (default: txt)

```
export EDITOR="neovim"
export STOIC_DIR="~/MEGAsync/journal/"
export STOIC_TEMPLATE="~/MEGAsync/journal/template.md"
export STOIC_EXT="md"
```

## Contributions
This project is my little playground to learn golang. It would be awesome to learn about any improvements that are appropriate for this codebase.

Everyone is welcome to contribute.

## TODO's
I've been brainstorming for possible implements and here is a rough list of ideas:

- Provide a calendar like view overlayed with existing entries - `stoic entries` 
- show/edit certain journal entires 
- Extract Context class (all the code that reads os.Getenv())
- Extract all entry related code into a Entry struct
- Figure out testing
- export to pdf
- add support for Windows
- expose all stoic config variables with `stoic -about` command


## Development
As prerequisite, you need to have the [Go compiler](https://golang.org/doc/install).
Please check the [`go.mod`](go.mod) file to see what Go version stoic requires.

Fetch the sources:

```
git clone https://github.com/skatkov/stoic.git
cd stoic
```

In order to build the project, run:

```
make build
```

This automatically resolves the dependencies and compiles the source code into an
executable for your platform.
