# Install stoic

In order to not miss any updates you can subscribe to the release
notifications on [Github](https://github.com/skatkov/stoic) (at the top right:
“Watch”→“Custom”→“Releases”).

## MacOS
1. Download the latest version and unzip
   - [**Download for Intel**](https://github.com/skatkov/stoic/releases/latest/download/stoic-mac-intel.zip)
   - [**Download for M1 (ARM)**](https://github.com/skatkov/stoic/releases/latest/download/stoic-mac-arm.zip)
2. Right-click on the binary and select “Open“
   (due to [Gatekeeper](https://support.apple.com/en-us/HT202491))
3. Copy to path, e.g. `mv stoic /usr/local/bin/stoic` (might require `sudo`)

## Linux
1. [**Download**](https://github.com/skatkov/stoic/releases/latest/download/stoic-linux.zip)
   the latest version and unzip
2. Copy to path, e.g. `mv stoic /usr/local/bin/stoic` (might require `sudo`)

## Windows
Is not supported


# Build stoic from sources

Instead of downloading the binaries, you can also build stoic yourself.

As prerequisite, you need to have the [Go compiler](https://golang.org/doc/install).
Please check the [`go.mod`](go.mod) file to see what Go version stoic requires.

Fetch the sources:

```
git clone https://github.com/skatkov/stoic.git
cd stoic
```

In order to build the project, run:

```
go build stoic.go
```

This automatically resolves the dependencies and compiles the source code into an
executable for your platform.
