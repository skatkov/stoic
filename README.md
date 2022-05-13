# Stoic
stoic is a plain-text file format command line utility for daily journaling.

After you run 'stoic' command utility will:
- create a a file <year-month-day.txt> in a directory of your choosing.
- it will open up this file in editor (default: nano)

if file will be already present, it will just offer you a way to edit it.

You can create a file in ~/.stoic to alter utility parameters:

```
editor = "/usr/bin/nano"
directory = "~/MEGAsync/journal/
```

