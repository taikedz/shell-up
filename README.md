# Shell-Up

An upgraded shell script writing experience.

## Motivation

Shell-Up is a macro-based transpiler for use with shell scripts, for working around some shortcomings of the shell scripting experience.

1. Code separation is essential for re-use
2. Resolving sourcing is error-prone as it is dependent on the top-script's current working directory.
    * Sourced scripts cannot reliably source other scripts
3. Function declaration syntax is limited. Shell-Up allows a macro for argument naming

It is desirable to transpile shell scripts typically because of thes limitations in code re-use, separation and being able to move-around single-scripts when deploying.

This is a rewrite of Bash Builder (see below) in Zig, which aims to achieve a few different goals:

* A no-dependencies executable for performing the transpilation
* A universal implementation that runs the same independently of execution platform
* Syntax macros built-in to the transpiler
* A more eloquent code base for maintainability
* Agnostic of system `*sh` type

## History

Previous incarnation was [bash-builder](https://gitlab.com/taikedz/bash-builder) which was implemented purely in bash "for dog-fooding".

Whilst it was a nice proof of concept, there were a number of burdens:

* token-processing and transformation is not the strongest suit for bash, and the syntax improvements had to be done as library content, not built-in
* due to the use of `echo` rather than `return` in shell scripting, and the lack of rich types, a lot of the resolution tracking and syntax is cumbersome and was difficult to maintain
* it was written specifically in bash, making most use-cases for plain `sh` scripting unachievable, especially on such systems as BSD where `bash` is usually not used
* it also made use of external programs from the system which meant it would not necessarily run the same on different OS-es
    * implementations (and options) of tar, gzip and some other key utilities varied from Ubuntu to Fedora to BSD to Alpine (and from version to version of each of these) making a truly universal implementation unlikely.
* it relied itself on system utilities being present, as all bash scripts do
    * not usually an issue when developing a bash script for itnernal use on platforms you control
    * again, the tool was meant to be universal, but fell short of this goal

These limitations also exist for its related project, [bash-libs](https://gitlab.com/taikedz/bash-libs), but on a less fundamental level. Libs for different target systems can be derived and maintained separately for each user's use-case, and (speculatively, off-topic) maybe even `bash-libs-linux` and `sh-libs-bsd` might be split out.
