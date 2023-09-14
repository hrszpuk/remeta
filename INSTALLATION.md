# Installation guide for Remeta

## Linux/MacOS
We start with Linux/MacOS because people that use Linux are based and people who use MacOS have lots of money (sponsor me ;)).

Firstly clone the git repository onto your local machine (I would recommend doing this in the ~/Downloads folder).
``` 
git clone https://github.com/hrszpuk/remeta.git && cd remeta
```
Now that you've cloned the repository and changed to your current directory to remeta you need to build the source code (this requires Go/Golang to be installed).
```
go build -v . 
```
This should produce file called "remeta".
Once you've moved the "remeta" file to the location you would like to keep it on your machine, go do your home directory and open your `.bashrc` file or `.zshrc` file (either will do).
Go to the end of the file and enter the following:
``` 
export PATH="path/to/remeta:$PATH"
```
**IMPORTANT:** change the "path/to/remeta" of the line above to the location you installed the "remeta" file.
Save the changes and exit. Now, reload your shell (type `bash` or `zsh`).

:beers: Congratulations you've installed Remeta :beers:

Checkout the [user guide](./USER_GUIDE.md) for how to use Remeta, or simply type `remeta` into your terminal to get a basic usage.

## Windows
Firstly clone the git repository onto your local machine (I would recommend doing this in the ~/Downloads folder).
``` 
git clone https://github.com/hrszpuk/remeta.git
```
Now that you've cloned the repository, your need to open the "remeta" folder and build the project using go (this requires Go/Golang to be installed).
```
go build -v . 
```
This should produce file called "remeta.exe".

Place this executable somewhere safe, and add the path to "remeta.exe" to your PATH variable.

Checkout the [user guide](./USER_GUIDE.md) for how to use Remeta, or simply type `remeta` into your terminal to get a basic usage.