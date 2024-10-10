# Shell-Up Build

## Main pre-processor sequence

`shup build [-d OUT] FILE` will do the following:

* Take the main file name from argument FILE
* if OUT is specified, create the file and prepare it for writing
* process line-by-line, inspecting for file macros and syntax macros, and writing other lines to output
* when a file macro is found
    * derive the absolute and canonical file path for each file macro
    * if seen and macros is '#%include', skip
    * else if not seen or if macro is '#%insert'
        * register the absolute and canonical file path
        * recurse into this file
    * write the result to output
* when a syntax macro is found
    * perform the transformation and write to output

## File macros

There are two macros for pulling in external files ("file import"):

* `#%include TARGET` will include the target file contents if and only if the file has not been seen by any file macro
* `#%insert TARGET` will include the target file contents any time `#%insert` has been seen

Both macros push the file contents through the pre-processor.

### shup-file

Inclusion resolution is done from a home "shup-file" at `~/.config/shell-up/shup.paths`, or from a path specified in a local shup-file at `./shup.paths` (ignoring the general location unless its first line is `#%add`)

* this is a file of absolute directory paths, not necessarily canonical
* one path per line, empty lines and `#` comment lines accepted
* the locations are all directories
* the locations are searched first-to-last to resolve the include/insert target

## Syntax macros

Shell-Up natively supports some syntax macros for better QOL when writing shell scripts

* `$%function <name>(<arg names>) {`
    * supporting positional names as local variables
    * if name prefixed with `*` then the name is a `declare -n <name>` reference (only works in bash)
    * if name is prefixed with `!` it is made global
    * if token `?` is encountered, the subsequent positional names are optional, cannot be a declare name

* Multiline exclusion boundaries: `#--` and `#--#`
    * Excluded lines are not included in the finally assembled file at all.
    * Start marked by a line starting with optional whitespace then `#--` .
    * End marked by any line ending in `#--#` followed directly by newline or carriage return.
    * As this is a naive macro, this will also be interpreted from within heredoc strings.
    * Enders will be found if placed at the end of normal comment lines.

## Other macros

Some extra macros exist for library maintainer convenience

`#%warn Message` - message to be printed by the transpiler when seen during a file import. Typically will be used by 3rd party library maintainers to flag impending deprecations.

`###doc :<tags>` - denotes a documentation block, which will be removed during file import. It is ended by the `###/doc` line. A doc block can be looked for by `shell-doc` if implemented.

## Extra tools

`sh-doc` - allows printing of documentation string sections. `bash-doc [TAG]` will print only documentation items tagged with the `TAG` , or all documentation sections if none specified.

