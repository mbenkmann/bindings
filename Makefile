DOXYSDL:=doxygen/sdl
DOXYIMG:=doxygen/sdl_image
DOXYTTF:=doxygen/sdl_ttf
SDLOUT:=sdl
IMGOUT:=img
TTFOUT:=ttf
SDLHEADERS:=SDL.h SDL_error.h SDL_events.h SDL_joystick.h SDL_keyboard.h SDL_keycode.h SDL_scancode.h \
            SDL_rect.h SDL_touch.h SDL_gesture.h SDL_rwops.h SDL_video.h SDL_surface.h SDL_blendmode.h \
            SDL_pixels.h SDL_render.h SDL_filesystem.h SDL_timer.h SDL_version.h SDL_shape.h \
            SDL_power.h SDL_clipboard.h SDL_cpuinfo.h SDL_mouse.h SDL_gamecontroller.h SDL_messagebox.h \
            SDL_hints.h SDL_haptic.h SDL_audio.h
SDLGOOUT:=$(patsubst %.h,%.h.go,$(SDLHEADERS))

.PHONY: clean $(SDLHEADERS) SDL_image.h SDL_ttf.h all

all: $(DOXYSDL) $(DOXYIMG) $(DOXYTTF) $(SDLHEADERS) SDL_image.h SDL_ttf.h

# If there is a special case generators/headername.py we call that.
# Otherwise we use the generic SDL.py.
$(SDLHEADERS):
	generator=SDL.py ; \
	test -x generators/$(patsubst %h.,%.py,$@) && generator=$(patsubst %h.,%.py,$@) ; \
	generators/$$generator $@ $(DOXYSDL) >$(SDLOUT)/$@.go

SDL_image.h:
	generators/SDL_image.py $@ $(DOXYIMG) >$(IMGOUT)/$@.go

SDL_ttf.h:
	generators/SDL_ttf.py $@ $(DOXYTTF) >$(TTFOUT)/$@.go

$(DOXYSDL):
	doxygen doxygen/SDL.dox

$(DOXYIMG):
	doxygen doxygen/SDL_image.dox

$(DOXYTTF):
	doxygen doxygen/SDL_ttf.dox


clean:
	rm -rf $(DOXYSDL)
	rm -rf $(DOXYIMG)
	rm -rf $(DOXYTTF)
	rm -f bin/debug-temp
	cd $(SDLOUT) && rm -f $(SDLGOOUT)
	rm -f $(IMGOUT)/SDL_image.h.go
	rm -f $(TTFOUT)/SDL_ttf.h.go

distclean: clean
	find -name "*~" -exec rm {} \;
