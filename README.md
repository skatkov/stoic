# Stoic
![stoic](./demo.png)

`stoic` is a command line app for daily journaling with plain-text files. It helps maintain day-to-day journaling habbit by allowing you to:
- Instantly edit current day entry
- Journal using any plain-text format (txt, md, [xit](https://xit.jotaen.net/) or anything else)
- Basic templates support
- Works on Linux and OSX
- Overview/filtering of existing entries

## Installation
Fast installation is possible through [brew](https://brew.sh/) on Linux or MacOS.

`brew install skatkov/tap/stoic`

If you want to download executable on your own, then there is another 📥 [Installation](INSTALL.md) instruction that comes with release.

## Usage
Run following commands in a terminal:

* `stoic` - opens todays journal entry in editor
* `stoic -list` (beta) - lists all existing entries and allows to pick one for edit.
* `stoic -edit "yesterday"` - will open or create file for previous days. Tries to understand natural dates ("2 days ago", "yesterday" and etc.)
* `stoic -quote` - shows random quote to encourage journaling habit
* `stoic -about` - shows information about cli app

## Configuration

- Editor could be changed by setting $EDITOR global variable. (default: nano)
- Directory for journal could be changed by setting $STOIC_DIR global variable. (default: ~/Journal/)
- Provide file template through $STOIC_TEMPLATE global variable.
- Provide new extension format through $STOIC_EXT global variables. (default: md)

```
export EDITOR="neovim"
export STOIC_DIR="~/MEGAsync/journal/"
export STOIC_TEMPLATE="~/MEGAsync/journal/template.md"
export STOIC_EXT="md"
```

## Motivation
There is a recurring theme through biographies of great people - they all had journaling as a hobby.

I've been battling my inner demons with different methods in the past with varying success. But journaling has helped me keep these demons permanently at bay. My sleep will be peaceful, if it follows after careful and honest self-examination in my journal. 

Existing software for journaling and note taking is too slow to load and filled with features I don'
t really need. Plain-text files stored in cloud storage and edited througn nano is more than enough. But some recurring manual work was still required - create new daily file and  modify according to my template. This command line utility completely removes that manual work for me. 

Epictetus, great Stoic philospher and slave, once told to his students that "philosophy is something one should write down day by day". Hence name of this tool is a reference to this great human and hat tip to practical philosophy called Stoicism.

## Obsidian integration
I wrote this tool our of frustration with existing note taking apps (everything has to be in the cloud these days), but there are other apps that followed similar design choices as `stoic` did.

I'm now using [Obsidian](https://obsidian.md/) as a fully featured writing app. Obsidian comes with some core plugins that you can enable, one of those is "Daily Notes" that does something similar to `stoic`. Notherless I still keep using `stoic` as a companion cli app to Obisidian, with light configuration changes they play perfectly together.

With 0.6 version of stoic a lot of settings are aligned with those Obsidian defaults to. In my case, only two settings are tweaked:
```
export STOIC_DIR="~/Obsidian/journal"
export STOIC_TEMPLATE="~/Obsidian/journal/template.md" 
```
This closely corresponds to settings I have in obsidian daily notes:
![obsidian daily notes config](./obsidian-daily-notes.png)


The only thing that stoic doesn't allow you to tweak is date format. So if you have that customized in obsidian - bad luck.

but PR's are welcome, everything is fixable.

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

## Contributions
This project is my little golang playground. It would be awesome to learn about any improvements that are appropriate for this codebase.

Everyone is welcome to contribute.

## TODO's
I've been brainstorming for possible improvements and here is a rough list of ideas with no particular order:

- Calendar view for existing records )show dates, mark with green dots if there is a record for that day) - `stoic -cal` (with `bubbletea` tui framework + `cal` utility.
- Use GoReleaser to automate all the routine (homebrew tap, building)
- Handle Errors better
- Use normal CLI framework ([cobra](https://github.com/spf13/cobra) seems popular)
- add support for Windows
- Add ability to store configuration in dotfiles (not just ENV variables), Including ability to add custom editor, not one defined in $EDITOR
- `stoic -stats` to receive statistics about journaling (how much % you didn''t journal, average journal lenght and etc)

## Feedback
Got some feedback or suggestions? Please open an issue or drop me a note!

* [Twitter](https://twitter.com/5katkov)
* [Personal Website](https://skatkov.com)
