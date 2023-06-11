# Rename Sourcetree

This is a tool to rename source code projects. It scratches a simple itch, namely renaming an Elixir/Phoenix web application.

For a project named "TestProject" this tool will appropriately rename files, directories, and file content while respecting two versions of the old/new name:
- the lowercase version of the string, e.g. `testproject`
- one specific MixEdCaSe version of the string -- for my purposes this is the camelcased Elixir module name for my project, e.g. TestProject

**WARNING** This is destructive! Always work on a copy of your project.

## Usage

```shell
# Copy your original project -- don't overwrite it in place!
cp -r /path/to/originalapp .

# Run the script
go run main.go --old=OriginalApp --new=SpinoffApp --targetdir=originalapp
```

## Testing/Development

There's a `skel/` directory with a test filesystem skeleton on it.

Run the `./reset.sh` script to automatically copy that to a `test` directory.

```shell
./reset.sh
```

Then you can test this tool by running

```
 go run main.go --old=old --new=new --targetdir=test
```

## Notes:

This is really just for me. Good luck!

## TODO
I could make this more flexible when it comes to mixed-case support, but I don't need to right now and this is such a tiny itch-scratcher that I doubt anyone else will use it. Modern IDEs probably have some function hidden in the menus that do this automatically. No idea, I use neovim now and have drunk all the Kool-Aid.


