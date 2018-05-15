#!/usr/bin/python3

import lib

lib.sdl()
lib.out.append(lib.SDL_COMMON_FILE_HEADER)
import_idx = len(lib.out)
lib.out.append("")

lib.describe(lib.soup.doxygen.compounddef)

lib.structs()
lib.unions()

for section in lib.soup.doxygen.compounddef.find_all("sectiondef"):
    lib.out.append("")
    lib.describe(section)
    lib.define2const(section)
    lib.enum2const(section)
    lib.aliastypedefs(section)
    lib.simpletypedefs(section)
    lib.wrapfunctions(section)

lib.additional_boilerplate(import_idx)

print("\n".join(lib.out))
