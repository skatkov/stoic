# Stoic
`stoic` is a command line app for daily journaling with plain-text files.

## Motivation
There is a recurring theme through biographies of great people - they all had journaling as a hobby.

I've been battling my inner demons with different methods in the past with varying success. But journaling has helped me keep these demons permanently at bay. My sleep will be peaceful, if it follows after careful and honest self-examination in my journal. 

Existing software for journaling and note taking is too slow to load and filled with features I don'
t really need. Plain-text files stored in cloud storage and edited througn nano is more than enough. But some recurring manual work was still required - create new daily file and  modify according to my template. This command line utility completely removes that manual work for me. 

Epictetus, great Stoic philospher and slave, once told to his students that "philosophy is something one should write down day by day". Hence name of this tool is a reference to this great human and hat tip to practical philosophy called Stoicism.

## Brief
After you run `stoic` command utility will:
- create or open file `<current_year>-<current_month>-<current_day>.txt` in a directory of your choosing (default: ~/Journal/).
- File will open in a editor you defined in $EDITOR global variable (default: nano)
- It's possible to define template for a new entry
- Plain-text files with any format (txt, md or anything else)

## Installation

📥 [Installation](INSTALL.md)

## Configration

- Editor could be changed by setting $EDITOR global variable.
- Directory for journal could be changed by setting $STOIC_DIR global variable.
- Provide file template through $STOIC_TEMPLATE global variable.
- Provide new extension format through $STOIC_EXT global variables

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
I've been brainstorming for possible ideas to implement further and here id what I came up with:

- Provide a calendar like view overlayed with existing entries - `stoic entries` 
- show/edit certain journal entires 
- Extract Context class (all the code that reads os.Getenv())
- Extract all entry related code into a Entry model
- Figure out testing
- Show app version  - `stoic version`
