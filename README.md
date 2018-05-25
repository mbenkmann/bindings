This project maintains
======================

 * Mostly auto-generated Go language bindings for various C libraries
 * Scripts to generate Go bindings for C libraries

Supported libraries
===================
Ready to Use
------------
 * Simple Direct Media Layer (SDL) core (i.e. SDL_* functions), http://www.libsdl.org

In Progress
-----------
 * SDL_image, https://www.libsdl.org/projects/SDL_image

Planned
-------
 * SDL_ttf, https://www.libsdl.org/projects/SDL_ttf
 * SDL_mixer, https://www.libsdl.org/projects/SDL_mixer
 * SDL2_gfx, http://www.ferzkopp.net/wordpress/2016/01/02/sdl_gfx-sdl2_gfx/

Examples
========
See what code using these bindings looks like:

https://github.com/mbenkmann/bindings/tree/master/examples

API Documentation
=================

https://godoc.org/github.com/mbenkmann/bindings/sdl

Using the ready-made bindings
=============================

Generating bindings yourself
============================

https://github.com/mbenkmann/bindings/tree/master/generators

Project Goals
=============

 * Provide developers who have used the supported libraries with other
   programming languages an easy way to use them with Go.
 * Offer a Go API that is similar to the C API, so that
     - developer knowledge transfers over from C to Go.
     - code snippets, examples and usage patterns can be translated between
       C and Go in a straightforward manner.
     - the original API reference documentation remains useable with the Go
       bindings.
     - tutorials written for the C API remain useable for Go programmers.
     - Go developers can participate in discussions with non-Go developers in
       the libraries' primary community (forums, mailing lists,...)

   However, API that permits out of bounds writes without use of package unsafe
   *must* be avoided; and API that permits out of bounds reads without package
   unsafe *should* be avoided.
 * API stability matching the original library, i.e. as long as the original
   library's API does not change, the Go binding API should not change.
 * Produce as much of the bindings as possible with generalized scripts, keep
   library-specific configuration to a minimum and use handwritten bindings
   only as a last resort. This way
     - bindings for new versions of the libraries can be generated instantly.
     - users of the Go bindings can generate new bindings themselves and don't
       have to wait for this project to catch up with new versions.
     - the bindings can be used with custom forks of the supported libraries.
     - developers can generate bindings for their own code they have developed
       around the supported libraries, as long as that code is reasonably
       similar in terms of types and conventions.
 * Do not introduce additional failure points, i.e. if a C code fragment using
   a supported library has no bugs, the equivalent Go code using the bindings
   should not have any bugs, either. In fact, the behaviour should be exactly
   the same.
 * Support Linux, Windows and macOS.

Explicit Non-Goals
==================

The following are things that could be legitimate project goals but that this
project has deliberately chosen not to pursue:

 * Following the conventions of native Go libraries: The Go bindings make some
   mild use of Go features and conventions, in particular multiple return values
   and function receivers (paired with the removal of the respective type name
   from the function name, e.g. SDL_WindowDestroy(win) => win.Destroy()),
   because these features make the API so much nicer to use.
   But in general every step towards making the bindings look more like a native
   Go library would compromise the project goals. Changes towards "more Go" are
   only made if they offer a huge benefit. E.g. Using error type returns in
   place of sdl.GetError() does not meet that threshold.
 * Being 100% complete. Some C functions do not work well with Go code, in
   particular everything including callbacks and varargs. Other functions have
   native Go alternatives that are preferable. If the generator scripts can
   create good bindings for these functions they will be left in. Otherwise
   they are blacklisted and omitted without replacement.
 * Replacing calls to the C library with native Go code. 
   While native Go code might be more efficient and/or have other benefits, its
   behaviour can never be guaranteed to be identical to that of the C library.
   When a wrapper for a C function does not make sense, it is better to omit it
   completely than to write a native Go function whose name and signature
   mislead users into believing that it is a wrapper.
 * Supporting different versions of a library with one version of bindings, i.e.
   making sure that a simple "go get" will always install a set of bindings
   that works with the version of the library installed on the system.
   The Go bindings are the counterpart of the C header files for the library.
   Old bindings with a new library version will work most of the time.
   New bindings with an old version will often fail, because the new bindings
   expect symbols that the old library does not have. The chosen solution is to
   tag generated bindings for different library versions in the repository to
   make it easy to switch to an appropriate set of bindings. Furthermore users
   can easily generate bindings that match their library version.

Project status
==============

*The project is being developed mainly on Ubuntu Linux. Your help in supporting
other platforms is appreciated.*

This project was initiated in May 2018. It is therefore very young and not
completely settled.

Because the bindings are mechanically
generated, most functions have only received little manual review. Especially
functions that take pointer arguments are an issue because unless the pointer
target is declared "const", it is not possible to automatically determine if it
is an input, output or both. The scripts rely on manually maintained lists of
defaults and exceptions for these cases. As this project gets used by more
people, these lists will get better. However, as a user of this library you
have to do your part in this effort.

*If you see something, say something!*

Filing an issue on Github is quick and painless and unlike calling the cops on
a "suspicious" cardboard box, an incorrect issue will not cause airports to
shut down or schools to be evacuated.

Examples of things to look out for:
 
 * Documentation that refers to NULL being a valid and useful value in a place
   where the Go bindings do not use a pointer. This is a particular issue with
   SDL_Rect which the Go bindings pass by value in most places.
 * Input or Input/Output parameters that the Go bindings incorrectly use as
   output parameters. Of particular note are pointers to a primitive type, such
   as int, which by default are treated as pure output parameters and returned
   by value in the Go bindings.
 * Documentation that mentions the transfer of ownership of memory blocks, e.g.
   that the caller has to free a pointer that is returned or that the caller
   must not free a string after passing it to the library.

In all such cases it doesn't hurt to take a quick look at the generated wrapper
function. Most issues are easy to spot, e.g. if the documentation mentions that
the caller has to free a returned string, but the wrapper function does not
contain the word "free", that's almost certainly an issue.
