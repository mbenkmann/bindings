Why generate the bindings yourself?
===================================
There are a variety of reasons why you may want to use the generator scripts
 to do the binding generation yourself. Some possible reasons include:

* If you have made changes to a supported library, you can generate bindings that properly   
  match your changes.
* If you have used a supported library in the past with a different programming language, 
  you will likely have created code around that library that you may want to re-use.
  Instead of porting this code to Go, you can generate bindings for it and use it directly.
  The only requirement is that the code has C .h header files. If the code is actual C
  code, you don't even need to convert it to a library. You can simply drop your .c files
  into the output directory of the generated bindings and the Go compiler will pull them in
  (see the documentation of the "`cgo`" package)
  
  If your code uses the same datatypes and follows the naming conventions of the supported
  library, chances are good that the generator scripts will produce good bindings for your
  own code right away with no or minimal tweaking.
* If you need bindings for a version of a supported library that is different from the 
  versions contained in the repository, you have to generate the bindings yourself.
  In particular this allows you to get bindings for a newly released library version
  on day 1, without having to wait for this project to react to the new release.
* Integrating the generator scripts and the generation of the bindings into your project's 
  build process insulates you from upstream changes and allows you to upgrade a library
  without upgrading to new bindings. While upgrading the bindings is usually not an issue
  because API stability is an explicit goal of this project, incompatible changes can never
  be ruled out completely. This could leave you in a situation where you have a deadline
  arriving and can't upgrade your library because the new bindings have an issue.
  Or even if there is no actual issue with the bindings, your project's testing process
  might require time consuming re-tests if the bindings are changed. If you integrate the
  binding generation into your process, you can avoid these issues.

Requirements
=============
* You need to have installed the development files for the libraries to generate bindings for
  and the ready-made bindings, as described in the main [README](../README.md).
* Python 3:

      apt-get install python3

* Beautiful Soup 4:

      apt-get install python3-bs4

* LXML:

      apt-get install python3-lxml

* Doxygen:

      apt-get install doxygen

'`master`' vs '`newest`'
========================
Even though the ready-made bindings in the '`master`' branch are kept at an older version,
this does not apply to the generator scripts. The generator scripts are supposed to be identical
in both branches. During periods of active development one or the other branch may temporarily
fall behind because development will only occur on one branch and will be synch'ed over at
regular intervals. There is no general rule about which branch has the most recent generator
scripts. You will have to look at the repository logs for the 2 branches to find out which one
has the most recent change to the generators/ directory.

Generating the bindings
=======================
1. Copy the `.h` header files to base the bindings on to the appropriate subdirectory of 
   `include/`
2. If necessary, edit `Makefile` to change the list of `.h` files to process.
3. Run the binding generation:

       make clean
       make


Customizing the generator scripts
=================================
The main code for the binding generation is in `generators/lib.py`. However for most customizations
you won't need to touch that code. A lot of the binding generation is controlled through maps and
lists in files like `generators/SDL.h.go`. Before you consider editing `lib.py`, see if you can't
achieve what you want through changing one of the files like `SDL.h.go` or creating your own.
If you do encounter a situation where you have to edit `lib.py` consider sharing your changes with
the main project by submitting them via github.
