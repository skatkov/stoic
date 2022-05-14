# Stoic
`stoic` is a plain-text file format command line utility for daily journaling.

## Motivation
There is a recurring theme through biographies of great people - they all had journaling as a hobby.

I've been battling my inner demons with different methods in the past with varying success. But journaling has helped me keep these demons permanently at bay. My sleep became particularly peaceful, if it follows after careful and honest self-examination in my journal. 


Non of current journaling or notes tools really stuck with me, but plain-text files edited with nano has been more than enough to stick with this habbit. Yeat, some recurring manual work has been required - create new daily file and  certain template.

Epictetus a great stoic philospher and slave, once told to his students that "philosophy was something they should write down day by day". Hence name of this tool is a reference to this great human and hat tip to philosphy that changed my life.

## Brief
After you run 'stoic' command utility will:
- create or open file <current_year>-<current_month>-<current_day>.txt in a directory of your choosing (default: ~/Journal/).
- File will be opened in a editr you defined in $EDITOR variable (or default: nano)
- It's possible to define template for a new entry (TODO)

## Configration

- Editor could be changed  by setting $EDITOR global variable.
- Directory for journal could be changed by setting $STOIC_DIR global variable.

## OS Support
This was primary written to be used on Linux, but should work OSX as well. 

Windows is not supported, I haven't touched it for ages.
