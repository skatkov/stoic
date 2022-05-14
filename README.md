# Stoic
`stoic` is a plain-text file format command line utility for daily journaling.

## Motivation
There is a recurring theme through biographies of great people - they all had journaling as a hobby.

I've been battling my inner demons with different methods in the past with varying success. But journaling has helped me keep these demons permanently at bay. My sleep became particularly peaceful, if it follows after careful and honest self-examination in my journal. 

Existing software for journaling and note taking is too slow to load and filled with features I don'
t really need. Plain-text files stored in cloud storage and edited througn nano is more than enough. But some recurring manual work was still required - create new daily file and  modify according to my template. This command line utility completely removes that manual work for me. 

Epictetus, great Stoic philospher and slave, once told to his students that "philosophy is something one should write down day by day". Hence name of this tool is a reference to this great human and hat tip to practical philosphy that changed my life.

## Brief
After you run 'stoic' command utility will:
- create or open file <current_year>-<current_month>-<current_day>.txt in a directory of your choosing (default: ~/Journal/).
- File will open in a editor you defined in $EDITOR global variable (but default is nano)
- It's possible to define template for a new entry (TODO)

## Configration

- Editor could be changed  by setting $EDITOR global variable.
- Directory for journal could be changed by setting $STOIC_DIR global variable.

## OS Support
This was primary written to be used on Linux, but should work OSX as well. 

Windows is not supported, I haven't touched it for ages.


## Contributions
This project is my little playground to learn golang. It would be awesome to learn about any improvements that are appropriate for this codebase.

Everyone is welcome to contribute.