# Messynotes CLI

This is a CLI to facilitate using Hugo as a personal notes/TODO system via the
[messynotes hugo theme](../).

## Name

I originally called this "barenotes", but then realized that totally collides
with the Bear Notes app.  Almost all synonyms for "simple" were taken, so I
went with "messy" because Messmore and whatever.

I'm mid-name change, so haven't adjusted everything.  So please bear with me.
(pun not intended, but enjoyed after typing)


## Installing/Building

I haven't created any releases yet.  I'm still test driving myself before I do
so.

You can clone the repo and run `make` to build the binary.

`make install` will install to `~/bin` by default and stick the shell
completion for zsh into `~/.zsh_functions`.  You can override those locations
by setting `BINDIR` and `SHELLCOMPDIR` but it's probably easier to just install
yourself if you don't want that setup.

I personally alias this to just `notes` and zsh handles the completion voodoo.
I didn't want to actually name the binary that because there are likely 100
`notes` commands out there.

## Usage

You probably want to have at least a `~/.messynotes.yaml` config file setting
the root directory of your hugo repo.  That makes it so you don't have to `cd`
to that every time.

Every top-level option can go in the config file.   You can also run
`messynotes showConfig` for a human-readable view of what it's actually using,
or `messynotes showConfig -y` to output the YAML that would be a full config
file.

Base Example:

```yaml
root: /home/mike/src/notes
```

This is the high-level usage.  All functionality requires the use of the
subcommands listed.

```text
The messynotes hugo theme is designed to be a minimalistic
system for maintaining personal notes and todo items.

This is a wrapper cli for providing simple access to maintain and use it
in this fashion.

This will try its best to choose the text editor of your choice.  The order
of precidence (first set wins): command line argument, config file,
$VISUAL, $EDITOR, editor command (for update-alternatives), and vi.  If
none are found, commands like 'new' will fail and you will need to
specify in the config or on the command line.

For the browser it will walk command line argument, config file, the 'open'
command, and then the 'xdg-open' command.  If none work, it will fail.

Usage:
  messynotes [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  edit        Edit an existing note
  git         Run git commands within the repo
  help        Help about any command
  init        Initialize new site/repository
  new         Create a new note
  open        Open a web browser to the hugo url (http://localhost:1313)
  restart     Restart hugo server
  showConfig  display configuration specified or implied
  start       Run the hugo server and open the browser to it.
  stop        Stop hugo server
  todo        Edit TODO file
  update      Update theme

Flags:
  -b, --browser string   Web browser binary/launcher
      --config string    config file (default $HOME/.messynotes.yaml)
  -e, --editor string    Text editor binary
  -g, --git string       Git binary (default "git")
  -h, --help             help for messynotes
  -H, --hugo string      Hugo binary (default "hugo")
  -r, --root string      Root of hugo repository (default ".")

Use "messynotes [command] --help" for more information about a command.
```

## Contributing

PRs will get the highest priority, either fixing bugs or extending features.
Issues that are bug reports will be next because they will bother me.  Feature
ask issues will be an "if I can get around to it."

I have a real job and a family and can't dedicate a lot of time to this.  It's
useful to me.  It may be to you.

## LICENSE

This is under the [MIT License](./LICENSE).
