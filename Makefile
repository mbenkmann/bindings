DOXYXML:=output/xml
SDLOUT:=sdl
SDLHEADERS:=SDL.h SDL_error.h SDL_events.h SDL_joystick.h SDL_keyboard.h SDL_keycode.h
SDLSRC:=SDL2-*

.PHONY: doxygen clean $(SDLHEADERS) all

all: $(SDLHEADERS)

# If there is a special case generators/headername.h.py we call that.
# Otherwise we use the generic SDL.h.py.
$(SDLHEADERS):
	generator=SDL.h.py ; \
	test -x generators/$@.py && generator=$@.py ; \
	generators/$$generator $@ $(SDLSRC)/$(DOXYXML) >$(SDLOUT)/$@.go

doxygen:
	cd $(SDLSRC) && doxygen ../doxyfile


clean:
	rm -rf $(SDLSRC)/$(DOXYXML)
	rm -f bin/debug-temp

distclean: clean
	find -name "*~" -exec rm {} \;
	