# The user guide for Remeta

### :warning: **Important notice about generation quirks** :warning:
Remeta is a metaprogramming tool which takes a Go package and convert it to ReRect Go function calls.
This makes creating bindings much easier than the manual process of registering and implementing functions by hand.
However, Remeta is still in early development, this means some quirks/bugs may be generated along with the bindings.
Fortunately, we've put a lot of focus into generating readable bindings, so you should be able to fix any quirks in no time!

## Using Remeta to generate bindings
After you've installed Remeta (read the [installation guide](./INSTALLATION.md) if you haven't), you can start generating bindings by navigating to the ReRect source code.
If you don't have ReRect installed, you should clone and build it from source.
This will grant you access to the folder `ReRect/src/go_packages/` which is where all built-in Go packages for ReRects are stored.

Once you have opened the `go_packages` folder, download the Go package you want to bind (using `git clone`) into the `go_packages` folder.
Now, use remeta to generate the bindings:
``` 
remeta <path/to/package-folder> <output-name>
```
Remeta will read through the package and generate a single file with all the functions/structs/methods binded and registered.

After Remeta has finished, open the file called `load.go` and located the function named `Load`. 
You can see other packages that are loaded using the function, you should be able to add your own package by taking the output name and appending it to `Load` (i.e. if your output name is `test` then you should write `LoadTest()` into the `Load` function).

Now recompile ReRect (`go build .`) and your package bindings will be built into ReRect.