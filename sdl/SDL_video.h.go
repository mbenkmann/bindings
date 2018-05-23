// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #cgo linux freebsd darwin pkg-config: sdl2
// #include <SDL.h>
// #include <SDL_shape.h>
import "C"
import "unsafe"

 // Header file for SDL video functions.

 // The structure that defines a display mode.
 // 
 // See also: SDL_GetNumDisplayModes()
 // 
 // See also: SDL_GetDisplayMode()
 // 
 // See also: SDL_GetDesktopDisplayMode()
 // 
 // See also: SDL_GetCurrentDisplayMode()
 // 
 // See also: SDL_GetClosestDisplayMode()
 // 
 // See also: SDL_SetWindowDisplayMode()
 // 
 // See also: SDL_GetWindowDisplayMode()
 // 
type DisplayMode struct {
     // pixel format
    Format uint32

     // width, in screen coordinates
    W int

     // height, in screen coordinates
    H int

     // refresh rate (or zero for unspecified)
    Refresh_rate int

     // driver-specific data, initialize to 0
    Driverdata uintptr
}

func fromC2DisplayMode(s C.SDL_DisplayMode) DisplayMode {
    return DisplayMode{uint32(s.format), int(s.w), int(s.h), int(s.refresh_rate), uintptr(s.driverdata)}
}

func toCFromDisplayMode(s DisplayMode) (d C.SDL_DisplayMode) {
    d.format = C.Uint32(s.Format)
    d.w = C.int(s.W)
    d.h = C.int(s.H)
    d.refresh_rate = C.int(s.Refresh_rate)
    d.driverdata = unsafe.Pointer(s.Driverdata)
    return
}

 // OpenGL support functions

 // Dynamically load an OpenGL library.
 // 
 // Returns: 0 on success, or -1 if the library couldn't be loaded.
 // 
 //   path
 //     The platform dependent OpenGL library name, or NULL to open the
 //     default OpenGL library.
 //   
 // This should be done after initializing the video driver, but before
 // creating any OpenGL windows. If no OpenGL library is loaded, the
 // default library will be loaded upon creation of the first OpenGL
 // window.
 // 
 // Note: If you do this, you need to retrieve all of the GL functions
 // used in your program from the dynamic library using
 // SDL_GL_GetProcAddress().
 // 
 // See also: SDL_GL_GetProcAddress()
 // 
 // See also: SDL_GL_UnloadLibrary()
 // 
func GL_LoadLibrary(path string) (retval int) {
    tmp_path := C.CString(path); defer C.free(unsafe.Pointer(tmp_path))
    retval = int(C.SDL_GL_LoadLibrary((*C.char)(tmp_path)))
    return
}

 // Get the address of an OpenGL function.
func GL_GetProcAddress(proc string) (retval uintptr) {
    tmp_proc := C.CString(proc); defer C.free(unsafe.Pointer(tmp_proc))
    retval = uintptr(C.SDL_GL_GetProcAddress((*C.char)(tmp_proc)))
    return
}

 // Unload the OpenGL library previously loaded by SDL_GL_LoadLibrary().
 // 
 // See also: SDL_GL_LoadLibrary()
 // 
func GL_UnloadLibrary() {
    C.SDL_GL_UnloadLibrary()
}

 // Return true if an OpenGL extension is supported for the current
 // context.
func GL_ExtensionSupported(extension string) (retval bool) {
    tmp_extension := C.CString(extension); defer C.free(unsafe.Pointer(tmp_extension))
    retval = C.SDL_TRUE==(C.SDL_GL_ExtensionSupported((*C.char)(tmp_extension)))
    return
}

 // Reset all previously set OpenGL context attributes to their default
 // values.
func GL_ResetAttributes() {
    C.SDL_GL_ResetAttributes()
}

 // Set an OpenGL window attribute before window creation.
func GL_SetAttribute(attr GLattr, value int) (retval int) {
    retval = int(C.SDL_GL_SetAttribute(C.SDL_GLattr(attr), C.int(value)))
    return
}

 // Get the actual value for an attribute from the current context.
func GL_GetAttribute(attr GLattr) (retval int, value int) {
    tmp_value := new(C.int)
    retval = int(C.SDL_GL_GetAttribute(C.SDL_GLattr(attr), (*C.int)(tmp_value)))
    value = deref_int_ptr(tmp_value)
    return
}

 // Create an OpenGL context for use with an OpenGL window, and make it
 // current.
 // 
 // See also: SDL_GL_DeleteContext()
 // 
func (window *Window) GL_CreateContext() (retval GLContext) {
    retval = GLContext(C.SDL_GL_CreateContext((*C.SDL_Window)(window)))
    return
}

 // Set up an OpenGL context for rendering into an OpenGL window.
 // 
 // Note: The context must have been created with a compatible window.
 // 
func (window *Window) GL_MakeCurrent(context GLContext) (retval int) {
    retval = int(C.SDL_GL_MakeCurrent((*C.SDL_Window)(window), C.SDL_GLContext(context)))
    return
}

 // Get the currently active OpenGL window.
func GL_GetCurrentWindow() (retval *Window) {
    retval = (*Window)(unsafe.Pointer(C.SDL_GL_GetCurrentWindow()))
    return
}

 // Get the currently active OpenGL context.
func GL_GetCurrentContext() (retval GLContext) {
    retval = GLContext(C.SDL_GL_GetCurrentContext())
    return
}

 // Get the size of a window's underlying drawable in pixels (for use with
 // glViewport).
 // 
 //   window
 //     Window from which the drawable size should be queried
 //   
 //   w
 //     Pointer to variable for storing the width in pixels, may be NULL
 //   
 //   h
 //     Pointer to variable for storing the height in pixels, may be NULL
 //   
 // This may differ from SDL_GetWindowSize() if we're rendering to a high-
 // DPI drawable, i.e. the window was created with
 // SDL_WINDOW_ALLOW_HIGHDPI on a platform with high-DPI support (Apple
 // calls this "Retina"), and not disabled by the
 // SDL_HINT_VIDEO_HIGHDPI_DISABLED hint.
 // 
 // See also: SDL_GetWindowSize()
 // 
 // See also: SDL_CreateWindow()
 // 
func (window *Window) GL_GetDrawableSize() (w int, h int) {
    tmp_w := new(C.int)
    tmp_h := new(C.int)
    C.SDL_GL_GetDrawableSize((*C.SDL_Window)(window), (*C.int)(tmp_w), (*C.int)(tmp_h))
    w = deref_int_ptr(tmp_w)
    h = deref_int_ptr(tmp_h)
    return
}

 // Set the swap interval for the current OpenGL context.
 // 
 // Returns: 0 on success, or -1 if setting the swap interval is not
 // supported.
 // 
 // See also: SDL_GL_GetSwapInterval()
 // 
 //   interval
 //     0 for immediate updates, 1 for updates synchronized with the vertical
 //     retrace. If the system supports it, you may specify -1 to allow late
 //     swaps to happen immediately instead of waiting for the next retrace.
 //   
func GL_SetSwapInterval(interval int) (retval int) {
    retval = int(C.SDL_GL_SetSwapInterval(C.int(interval)))
    return
}

 // Get the swap interval for the current OpenGL context.
 // 
 // Returns: 0 if there is no vertical retrace synchronization, 1 if the
 // buffer swap is synchronized with the vertical retrace, and -1 if late
 // swaps happen immediately instead of waiting for the next retrace. If
 // the system can't determine the swap interval, or there isn't a valid
 // current context, this will return 0 as a safe default.
 // 
 // See also: SDL_GL_SetSwapInterval()
 // 
func GL_GetSwapInterval() (retval int) {
    retval = int(C.SDL_GL_GetSwapInterval())
    return
}

 // Swap the OpenGL buffers for a window, if double-buffering is
 // supported.
func (window *Window) GL_Swap() {
    C.SDL_GL_SwapWindow((*C.SDL_Window)(window))
}

 // Delete an OpenGL context.
 // 
 // See also: SDL_GL_CreateContext()
 // 
func GL_DeleteContext(context GLContext) {
    C.SDL_GL_DeleteContext(C.SDL_GLContext(context))
}

const (
     // Used to indicate that you don't care what the window position is.
    WINDOWPOS_UNDEFINED_MASK = C.SDL_WINDOWPOS_UNDEFINED_MASK

    WINDOWPOS_UNDEFINED = C.SDL_WINDOWPOS_UNDEFINED

     // Used to indicate that the window position should be centered.
    WINDOWPOS_CENTERED_MASK = C.SDL_WINDOWPOS_CENTERED_MASK

    WINDOWPOS_CENTERED = C.SDL_WINDOWPOS_CENTERED
)

 // The flags on a window.
 // 
 // See also: SDL_GetWindowFlags()
 // 
type WindowFlags int
const (
     // fullscreen window
    WINDOW_FULLSCREEN WindowFlags = C.SDL_WINDOW_FULLSCREEN

     // window usable with OpenGL context
    WINDOW_OPENGL WindowFlags = C.SDL_WINDOW_OPENGL

     // window is visible
    WINDOW_SHOWN WindowFlags = C.SDL_WINDOW_SHOWN

     // window is not visible
    WINDOW_HIDDEN WindowFlags = C.SDL_WINDOW_HIDDEN

     // no window decoration
    WINDOW_BORDERLESS WindowFlags = C.SDL_WINDOW_BORDERLESS

     // window can be resized
    WINDOW_RESIZABLE WindowFlags = C.SDL_WINDOW_RESIZABLE

     // window is minimized
    WINDOW_MINIMIZED WindowFlags = C.SDL_WINDOW_MINIMIZED

     // window is maximized
    WINDOW_MAXIMIZED WindowFlags = C.SDL_WINDOW_MAXIMIZED

     // window has grabbed input focus
    WINDOW_INPUT_GRABBED WindowFlags = C.SDL_WINDOW_INPUT_GRABBED

     // window has input focus
    WINDOW_INPUT_FOCUS WindowFlags = C.SDL_WINDOW_INPUT_FOCUS

     // window has mouse focus
    WINDOW_MOUSE_FOCUS WindowFlags = C.SDL_WINDOW_MOUSE_FOCUS

    WINDOW_FULLSCREEN_DESKTOP WindowFlags = C.SDL_WINDOW_FULLSCREEN_DESKTOP

     // window not created by SDL
    WINDOW_FOREIGN WindowFlags = C.SDL_WINDOW_FOREIGN

     // window should be created in high-DPI mode if supported
    WINDOW_ALLOW_HIGHDPI WindowFlags = C.SDL_WINDOW_ALLOW_HIGHDPI

     // window has mouse captured (unrelated to INPUT_GRABBED)
    WINDOW_MOUSE_CAPTURE WindowFlags = C.SDL_WINDOW_MOUSE_CAPTURE
)

 // Event subtype for window events.
type WindowEventID int
const (
     // Never used
    WINDOWEVENT_NONE WindowEventID = C.SDL_WINDOWEVENT_NONE

     // Window has been shown
    WINDOWEVENT_SHOWN WindowEventID = C.SDL_WINDOWEVENT_SHOWN

     // Window has been hidden
    WINDOWEVENT_HIDDEN WindowEventID = C.SDL_WINDOWEVENT_HIDDEN

     // Window has been exposed and should be redrawn
    WINDOWEVENT_EXPOSED WindowEventID = C.SDL_WINDOWEVENT_EXPOSED

     // Window has been moved to data1, data2
    WINDOWEVENT_MOVED WindowEventID = C.SDL_WINDOWEVENT_MOVED

     // Window has been resized to data1xdata2
    WINDOWEVENT_RESIZED WindowEventID = C.SDL_WINDOWEVENT_RESIZED

     // The window size has changed, either as a result of an API call or
     // through the system or user changing the window size.
    WINDOWEVENT_SIZE_CHANGED WindowEventID = C.SDL_WINDOWEVENT_SIZE_CHANGED

     // Window has been minimized
    WINDOWEVENT_MINIMIZED WindowEventID = C.SDL_WINDOWEVENT_MINIMIZED

     // Window has been maximized
    WINDOWEVENT_MAXIMIZED WindowEventID = C.SDL_WINDOWEVENT_MAXIMIZED

     // Window has been restored to normal size and position
    WINDOWEVENT_RESTORED WindowEventID = C.SDL_WINDOWEVENT_RESTORED

     // Window has gained mouse focus
    WINDOWEVENT_ENTER WindowEventID = C.SDL_WINDOWEVENT_ENTER

     // Window has lost mouse focus
    WINDOWEVENT_LEAVE WindowEventID = C.SDL_WINDOWEVENT_LEAVE

     // Window has gained keyboard focus
    WINDOWEVENT_FOCUS_GAINED WindowEventID = C.SDL_WINDOWEVENT_FOCUS_GAINED

     // Window has lost keyboard focus
    WINDOWEVENT_FOCUS_LOST WindowEventID = C.SDL_WINDOWEVENT_FOCUS_LOST

     // The window manager requests that the window be closed
    WINDOWEVENT_CLOSE WindowEventID = C.SDL_WINDOWEVENT_CLOSE
)

 // OpenGL configuration attributes.
type GLattr int
const (
    GL_RED_SIZE GLattr = C.SDL_GL_RED_SIZE

    GL_GREEN_SIZE GLattr = C.SDL_GL_GREEN_SIZE

    GL_BLUE_SIZE GLattr = C.SDL_GL_BLUE_SIZE

    GL_ALPHA_SIZE GLattr = C.SDL_GL_ALPHA_SIZE

    GL_BUFFER_SIZE GLattr = C.SDL_GL_BUFFER_SIZE

    GL_DOUBLEBUFFER GLattr = C.SDL_GL_DOUBLEBUFFER

    GL_DEPTH_SIZE GLattr = C.SDL_GL_DEPTH_SIZE

    GL_STENCIL_SIZE GLattr = C.SDL_GL_STENCIL_SIZE

    GL_ACCUM_RED_SIZE GLattr = C.SDL_GL_ACCUM_RED_SIZE

    GL_ACCUM_GREEN_SIZE GLattr = C.SDL_GL_ACCUM_GREEN_SIZE

    GL_ACCUM_BLUE_SIZE GLattr = C.SDL_GL_ACCUM_BLUE_SIZE

    GL_ACCUM_ALPHA_SIZE GLattr = C.SDL_GL_ACCUM_ALPHA_SIZE

    GL_STEREO GLattr = C.SDL_GL_STEREO

    GL_MULTISAMPLEBUFFERS GLattr = C.SDL_GL_MULTISAMPLEBUFFERS

    GL_MULTISAMPLESAMPLES GLattr = C.SDL_GL_MULTISAMPLESAMPLES

    GL_ACCELERATED_VISUAL GLattr = C.SDL_GL_ACCELERATED_VISUAL

    GL_RETAINED_BACKING GLattr = C.SDL_GL_RETAINED_BACKING

    GL_CONTEXT_MAJOR_VERSION GLattr = C.SDL_GL_CONTEXT_MAJOR_VERSION

    GL_CONTEXT_MINOR_VERSION GLattr = C.SDL_GL_CONTEXT_MINOR_VERSION

    GL_CONTEXT_EGL GLattr = C.SDL_GL_CONTEXT_EGL

    GL_CONTEXT_FLAGS GLattr = C.SDL_GL_CONTEXT_FLAGS

    GL_CONTEXT_PROFILE_MASK GLattr = C.SDL_GL_CONTEXT_PROFILE_MASK

    GL_SHARE_WITH_CURRENT_CONTEXT GLattr = C.SDL_GL_SHARE_WITH_CURRENT_CONTEXT

    GL_FRAMEBUFFER_SRGB_CAPABLE GLattr = C.SDL_GL_FRAMEBUFFER_SRGB_CAPABLE

    GL_CONTEXT_RELEASE_BEHAVIOR GLattr = C.SDL_GL_CONTEXT_RELEASE_BEHAVIOR
)

type GLprofile int
const (
    GL_CONTEXT_PROFILE_CORE GLprofile = C.SDL_GL_CONTEXT_PROFILE_CORE

    GL_CONTEXT_PROFILE_COMPATIBILITY GLprofile = C.SDL_GL_CONTEXT_PROFILE_COMPATIBILITY

    GL_CONTEXT_PROFILE_ES GLprofile = C.SDL_GL_CONTEXT_PROFILE_ES
)

type GLcontextFlag int
const (
    GL_CONTEXT_DEBUG_FLAG GLcontextFlag = C.SDL_GL_CONTEXT_DEBUG_FLAG

    GL_CONTEXT_FORWARD_COMPATIBLE_FLAG GLcontextFlag = C.SDL_GL_CONTEXT_FORWARD_COMPATIBLE_FLAG

    GL_CONTEXT_ROBUST_ACCESS_FLAG GLcontextFlag = C.SDL_GL_CONTEXT_ROBUST_ACCESS_FLAG

    GL_CONTEXT_RESET_ISOLATION_FLAG GLcontextFlag = C.SDL_GL_CONTEXT_RESET_ISOLATION_FLAG
)

type GLcontextReleaseFlag int
const (
    GL_CONTEXT_RELEASE_BEHAVIOR_NONE GLcontextReleaseFlag = C.SDL_GL_CONTEXT_RELEASE_BEHAVIOR_NONE

    GL_CONTEXT_RELEASE_BEHAVIOR_FLUSH GLcontextReleaseFlag = C.SDL_GL_CONTEXT_RELEASE_BEHAVIOR_FLUSH
)

 // Possible return values from the SDL_HitTest callback.
 // 
 // See also: SDL_HitTest
 // 
type HitTestResult int
const (
     // Region is normal. No special properties.
    HITTEST_NORMAL HitTestResult = C.SDL_HITTEST_NORMAL

     // Region can drag entire window.
    HITTEST_DRAGGABLE HitTestResult = C.SDL_HITTEST_DRAGGABLE

    HITTEST_RESIZE_TOPLEFT HitTestResult = C.SDL_HITTEST_RESIZE_TOPLEFT

    HITTEST_RESIZE_TOP HitTestResult = C.SDL_HITTEST_RESIZE_TOP

    HITTEST_RESIZE_TOPRIGHT HitTestResult = C.SDL_HITTEST_RESIZE_TOPRIGHT

    HITTEST_RESIZE_RIGHT HitTestResult = C.SDL_HITTEST_RESIZE_RIGHT

    HITTEST_RESIZE_BOTTOMRIGHT HitTestResult = C.SDL_HITTEST_RESIZE_BOTTOMRIGHT

    HITTEST_RESIZE_BOTTOM HitTestResult = C.SDL_HITTEST_RESIZE_BOTTOM

    HITTEST_RESIZE_BOTTOMLEFT HitTestResult = C.SDL_HITTEST_RESIZE_BOTTOMLEFT

    HITTEST_RESIZE_LEFT HitTestResult = C.SDL_HITTEST_RESIZE_LEFT
)

 // The type used to identify a window.
 // 
 // See also: SDL_CreateWindow()
 // 
 // See also: SDL_CreateWindowFrom()
 // 
 // See also: SDL_DestroyWindow()
 // 
 // See also: SDL_GetWindowData()
 // 
 // See also: SDL_GetWindowFlags()
 // 
 // See also: SDL_GetWindowGrab()
 // 
 // See also: SDL_GetWindowPosition()
 // 
 // See also: SDL_GetWindowSize()
 // 
 // See also: SDL_GetWindowTitle()
 // 
 // See also: SDL_HideWindow()
 // 
 // See also: SDL_MaximizeWindow()
 // 
 // See also: SDL_MinimizeWindow()
 // 
 // See also: SDL_RaiseWindow()
 // 
 // See also: SDL_RestoreWindow()
 // 
 // See also: SDL_SetWindowData()
 // 
 // See also: SDL_SetWindowFullscreen()
 // 
 // See also: SDL_SetWindowGrab()
 // 
 // See also: SDL_SetWindowIcon()
 // 
 // See also: SDL_SetWindowPosition()
 // 
 // See also: SDL_SetWindowSize()
 // 
 // See also: SDL_SetWindowBordered()
 // 
 // See also: SDL_SetWindowTitle()
 // 
 // See also: SDL_ShowWindow()
 // 
type Window C.SDL_Window

 // An opaque handle to an OpenGL context.
type GLContext C.SDL_GLContext

 // Callback used for hit-testing.
 // 
 // See also: SDL_SetWindowHitTest
 // 
type HitTest C.SDL_HitTest


 // Get the number of video drivers compiled into SDL.
 // 
 // See also: SDL_GetVideoDriver()
 // 
func GetNumVideoDrivers() (retval int) {
    retval = int(C.SDL_GetNumVideoDrivers())
    return
}

 // Get the name of a built in video driver.
 // 
 // Note: The video drivers are presented in the order in which they are
 // normally checked during initialization.
 // 
 // See also: SDL_GetNumVideoDrivers()
 // 
func GetVideoDriver(index int) (retval string) {
    retval = C.GoString(C.SDL_GetVideoDriver(C.int(index)))
    return
}

 // Initialize the video subsystem, optionally specifying a video driver.
 // 
 // Returns: 0 on success, -1 on error
 // 
 //   driver_name
 //     Initialize a specific driver by name, or NULL for the default video
 //     driver.
 //   
 // This function initializes the video subsystem; setting up a connection
 // to the window manager, etc, and determines the available display modes
 // and pixel formats, but does not initialize a window or graphics mode.
 // 
 // See also: SDL_VideoQuit()
 // 
func VideoInit(driver_name string) (retval int) {
    tmp_driver_name := C.CString(driver_name); defer C.free(unsafe.Pointer(tmp_driver_name))
    retval = int(C.SDL_VideoInit((*C.char)(tmp_driver_name)))
    return
}

 // Shuts down the video subsystem.
 // 
 // This function closes all windows, and restores the original video
 // mode.
 // 
 // See also: SDL_VideoInit()
 // 
func VideoQuit() {
    C.SDL_VideoQuit()
}

 // Returns the name of the currently initialized video driver.
 // 
 // Returns: The name of the current video driver or NULL if no driver has
 // been initialized
 // 
 // See also: SDL_GetNumVideoDrivers()
 // 
 // See also: SDL_GetVideoDriver()
 // 
func GetCurrentVideoDriver() (retval string) {
    retval = C.GoString(C.SDL_GetCurrentVideoDriver())
    return
}

 // Returns the number of available video displays.
 // 
 // See also: SDL_GetDisplayBounds()
 // 
func GetNumVideoDisplays() (retval int) {
    retval = int(C.SDL_GetNumVideoDisplays())
    return
}

 // Get the name of a display in UTF-8 encoding.
 // 
 // Returns: The name of a display, or NULL for an invalid display index.
 // 
 // See also: SDL_GetNumVideoDisplays()
 // 
func GetDisplayName(displayIndex int) (retval string) {
    retval = C.GoString(C.SDL_GetDisplayName(C.int(displayIndex)))
    return
}

 // Get the desktop area represented by a display, with the primary
 // display located at 0,0.
 // 
 // Returns: 0 on success, or -1 if the index is out of range.
 // 
 // See also: SDL_GetNumVideoDisplays()
 // 
func GetDisplayBounds(displayIndex int) (retval int, rect Rect) {
    tmp_rect := new(C.SDL_Rect)
    retval = int(C.SDL_GetDisplayBounds(C.int(displayIndex), (*C.SDL_Rect)(tmp_rect)))
    rect = fromC2Rect(*(tmp_rect))
    return
}

 // Get the dots/pixels-per-inch for a display.
 // 
 // Note: Diagonal, horizontal and vertical DPI can all be optionally
 // returned if the parameter is non-NULL.
 // 
 // Returns: 0 on success, or -1 if no DPI information is available or the
 // index is out of range.
 // 
 // See also: SDL_GetNumVideoDisplays()
 // 
func GetDisplayDPI(displayIndex int) (retval int, ddpi float32, hdpi float32, vdpi float32) {
    tmp_ddpi := new(C.float)
    tmp_hdpi := new(C.float)
    tmp_vdpi := new(C.float)
    retval = int(C.SDL_GetDisplayDPI(C.int(displayIndex), (*C.float)(tmp_ddpi), (*C.float)(tmp_hdpi), (*C.float)(tmp_vdpi)))
    ddpi = deref_float32_ptr(tmp_ddpi)
    hdpi = deref_float32_ptr(tmp_hdpi)
    vdpi = deref_float32_ptr(tmp_vdpi)
    return
}

 // Returns the number of available display modes.
 // 
 // See also: SDL_GetDisplayMode()
 // 
func GetNumDisplayModes(displayIndex int) (retval int) {
    retval = int(C.SDL_GetNumDisplayModes(C.int(displayIndex)))
    return
}

 // Fill in information about a specific display mode.
 // 
 // Note: The display modes are sorted in this priority:
 //   
 //   - bits per pixel -> more colors to fewer colors
 //   
 //   - width -> largest to smallest
 //   
 //   - height -> largest to smallest
 //   
 //   - refresh rate -> highest to lowest
 //   
 // 
 // See also: SDL_GetNumDisplayModes()
 // 
func GetDisplayMode(displayIndex int, modeIndex int) (retval int, mode *DisplayMode) {
    tmp_mode := new(C.SDL_DisplayMode)
    retval = int(C.SDL_GetDisplayMode(C.int(displayIndex), C.int(modeIndex), (*C.SDL_DisplayMode)(tmp_mode)))
    tmp2_mode := fromC2DisplayMode(*(tmp_mode)); mode = &tmp2_mode
    return
}

 // Fill in information about the desktop display mode.
func GetDesktopDisplayMode(displayIndex int) (retval int, mode *DisplayMode) {
    tmp_mode := new(C.SDL_DisplayMode)
    retval = int(C.SDL_GetDesktopDisplayMode(C.int(displayIndex), (*C.SDL_DisplayMode)(tmp_mode)))
    tmp2_mode := fromC2DisplayMode(*(tmp_mode)); mode = &tmp2_mode
    return
}

 // Fill in information about the current display mode.
func GetCurrentDisplayMode(displayIndex int) (retval int, mode *DisplayMode) {
    tmp_mode := new(C.SDL_DisplayMode)
    retval = int(C.SDL_GetCurrentDisplayMode(C.int(displayIndex), (*C.SDL_DisplayMode)(tmp_mode)))
    tmp2_mode := fromC2DisplayMode(*(tmp_mode)); mode = &tmp2_mode
    return
}

 // Get the closest match to the requested display mode.
 // 
 // Returns: The passed in value closest, or NULL if no matching video
 // mode was available.
 // 
 //   displayIndex
 //     The index of display from which mode should be queried.
 //   
 //   mode
 //     The desired display mode
 //   
 //   closest
 //     A pointer to a display mode to be filled in with the closest match of
 //     the available display modes.
 //   
 // The available display modes are scanned, and closest is filled in with
 // the closest mode matching the requested mode and returned. The mode
 // format and refresh_rate default to the desktop mode if they are 0. The
 // modes are scanned with size being first priority, format being second
 // priority, and finally checking the refresh_rate. If all the available
 // modes are too small, then NULL is returned.
 // 
 // See also: SDL_GetNumDisplayModes()
 // 
 // See also: SDL_GetDisplayMode()
 // 
func GetClosestDisplayMode(displayIndex int, mode *DisplayMode) (retval *DisplayMode, closest *DisplayMode) {
    var tmp_mode *C.SDL_DisplayMode; if mode != nil { x := toCFromDisplayMode(*mode); tmp_mode = &x }
    tmp_closest := new(C.SDL_DisplayMode)
    tmp_retval  := fromC2DisplayMode(*(C.SDL_GetClosestDisplayMode(C.int(displayIndex), (*C.SDL_DisplayMode)(tmp_mode), (*C.SDL_DisplayMode)(tmp_closest))))
    retval  = &tmp_retval 
    tmp2_closest := fromC2DisplayMode(*(tmp_closest)); closest = &tmp2_closest
    return
}

 // Get the display index associated with a window.
 // 
 // Returns: the display index of the display containing the center of the
 // window, or -1 on error.
 // 
func (window *Window) GetDisplayIndex() (retval int) {
    retval = int(C.SDL_GetWindowDisplayIndex((*C.SDL_Window)(window)))
    return
}

 // Set the display mode used when a fullscreen window is visible.
 // 
 // By default the window's dimensions and the desktop format and refresh
 // rate are used.
 // 
 // Returns: 0 on success, or -1 if setting the display mode failed.
 // 
 // See also: SDL_GetWindowDisplayMode()
 // 
 // See also: SDL_SetWindowFullscreen()
 // 
 //   window
 //     The window for which the display mode should be set.
 //   
 //   mode
 //     The mode to use, or NULL for the default mode.
 //   
func (window *Window) SetDisplayMode(mode *DisplayMode) (retval int) {
    var tmp_mode *C.SDL_DisplayMode; if mode != nil { x := toCFromDisplayMode(*mode); tmp_mode = &x }
    retval = int(C.SDL_SetWindowDisplayMode((*C.SDL_Window)(window), (*C.SDL_DisplayMode)(tmp_mode)))
    return
}

 // Fill in information about the display mode used when a fullscreen
 // window is visible.
 // 
 // See also: SDL_SetWindowDisplayMode()
 // 
 // See also: SDL_SetWindowFullscreen()
 // 
func (window *Window) GetDisplayMode() (retval int, mode *DisplayMode) {
    tmp_mode := new(C.SDL_DisplayMode)
    retval = int(C.SDL_GetWindowDisplayMode((*C.SDL_Window)(window), (*C.SDL_DisplayMode)(tmp_mode)))
    tmp2_mode := fromC2DisplayMode(*(tmp_mode)); mode = &tmp2_mode
    return
}

 // Get the pixel format associated with the window.
func (window *Window) GetPixelFormat() (retval uint32) {
    retval = uint32(C.SDL_GetWindowPixelFormat((*C.SDL_Window)(window)))
    return
}

 // Create a window with the specified position, dimensions, and flags.
 // 
 // Returns: The id of the window created, or zero if window creation
 // failed.
 // 
 //   title
 //     The title of the window, in UTF-8 encoding.
 //   
 //   x
 //     The x position of the window, SDL_WINDOWPOS_CENTERED, or
 //     SDL_WINDOWPOS_UNDEFINED.
 //   
 //   y
 //     The y position of the window, SDL_WINDOWPOS_CENTERED, or
 //     SDL_WINDOWPOS_UNDEFINED.
 //   
 //   w
 //     The width of the window, in screen coordinates.
 //   
 //   h
 //     The height of the window, in screen coordinates.
 //   
 //   flags
 //     The flags for the window, a mask of any of the following:
 //     SDL_WINDOW_FULLSCREEN, SDL_WINDOW_OPENGL, SDL_WINDOW_HIDDEN,
 //     SDL_WINDOW_BORDERLESS, SDL_WINDOW_RESIZABLE, SDL_WINDOW_MAXIMIZED,
 //     SDL_WINDOW_MINIMIZED, SDL_WINDOW_INPUT_GRABBED,
 //     SDL_WINDOW_ALLOW_HIGHDPI.
 //   
 // If the window is created with the SDL_WINDOW_ALLOW_HIGHDPI flag, its
 // size in pixels may differ from its size in screen coordinates on
 // platforms with high-DPI support (e.g. iOS and Mac OS X). Use
 // SDL_GetWindowSize() to query the client area's size in screen
 // coordinates, and SDL_GL_GetDrawableSize() or
 // SDL_GetRendererOutputSize() to query the drawable size in pixels.
 // 
 // See also: SDL_DestroyWindow()
 // 
func CreateWindow(title string, x int, y int, w int, h int, flags WindowFlags) (retval *Window) {
    tmp_title := C.CString(title); defer C.free(unsafe.Pointer(tmp_title))
    retval = (*Window)(unsafe.Pointer(C.SDL_CreateWindow((*C.char)(tmp_title), C.int(x), C.int(y), C.int(w), C.int(h), C.Uint32(flags))))
    return
}

 // Create an SDL window from an existing native window.
 // 
 // Returns: The id of the window created, or zero if window creation
 // failed.
 // 
 // See also: SDL_DestroyWindow()
 // 
 //   data
 //     A pointer to driver-dependent window creation data
 //   
func CreateWindowFrom(data uintptr) (retval *Window) {
    retval = (*Window)(unsafe.Pointer(C.SDL_CreateWindowFrom(unsafe.Pointer(data))))
    return
}

 // Get the numeric ID of a window, for logging purposes.
func (window *Window) GetID() (retval uint32) {
    retval = uint32(C.SDL_GetWindowID((*C.SDL_Window)(window)))
    return
}

 // Get a window from a stored ID, or NULL if it doesn't exist.
func GetWindowFromID(id uint32) (retval *Window) {
    retval = (*Window)(unsafe.Pointer(C.SDL_GetWindowFromID(C.Uint32(id))))
    return
}

 // Get the window flags.
func (window *Window) GetFlags() (retval uint32) {
    retval = uint32(C.SDL_GetWindowFlags((*C.SDL_Window)(window)))
    return
}

 // Set the title of a window, in UTF-8 format.
 // 
 // See also: SDL_GetWindowTitle()
 // 
func (window *Window) SetTitle(title string) {
    tmp_title := C.CString(title); defer C.free(unsafe.Pointer(tmp_title))
    C.SDL_SetWindowTitle((*C.SDL_Window)(window), (*C.char)(tmp_title))
}

 // Get the title of a window, in UTF-8 format.
 // 
 // See also: SDL_SetWindowTitle()
 // 
func (window *Window) GetTitle() (retval string) {
    retval = C.GoString(C.SDL_GetWindowTitle((*C.SDL_Window)(window)))
    return
}

 // Set the icon for a window.
 // 
 //   window
 //     The window for which the icon should be set.
 //   
 //   icon
 //     The icon for the window.
 //   
func (window *Window) SetIcon(icon *Surface) {
    C.SDL_SetWindowIcon((*C.SDL_Window)(window), (*C.SDL_Surface)(icon))
}

 // Associate an arbitrary named pointer with a window.
 // 
 // Returns: The previous value associated with 'name'
 // 
 // Note: The name is case-sensitive.
 // 
 // See also: SDL_GetWindowData()
 // 
 //   window
 //     The window to associate with the pointer.
 //   
 //   name
 //     The name of the pointer.
 //   
 //   userdata
 //     The associated pointer.
 //   
func (window *Window) SetData(name string, userdata uintptr) (retval uintptr) {
    tmp_name := C.CString(name); defer C.free(unsafe.Pointer(tmp_name))
    retval = uintptr(C.SDL_SetWindowData((*C.SDL_Window)(window), (*C.char)(tmp_name), unsafe.Pointer(userdata)))
    return
}

 // Retrieve the data pointer associated with a window.
 // 
 // Returns: The value associated with 'name'
 // 
 // See also: SDL_SetWindowData()
 // 
 //   window
 //     The window to query.
 //   
 //   name
 //     The name of the pointer.
 //   
func (window *Window) GetData(name string) (retval uintptr) {
    tmp_name := C.CString(name); defer C.free(unsafe.Pointer(tmp_name))
    retval = uintptr(C.SDL_GetWindowData((*C.SDL_Window)(window), (*C.char)(tmp_name)))
    return
}

 // Set the position of a window.
 // 
 // Note: The window coordinate origin is the upper left of the display.
 // 
 // See also: SDL_GetWindowPosition()
 // 
 //   window
 //     The window to reposition.
 //   
 //   x
 //     The x coordinate of the window in screen coordinates, or
 //     SDL_WINDOWPOS_CENTERED or SDL_WINDOWPOS_UNDEFINED.
 //   
 //   y
 //     The y coordinate of the window in screen coordinates, or
 //     SDL_WINDOWPOS_CENTERED or SDL_WINDOWPOS_UNDEFINED.
 //   
func (window *Window) SetPosition(x int, y int) {
    C.SDL_SetWindowPosition((*C.SDL_Window)(window), C.int(x), C.int(y))
}

 // Get the position of a window.
 // 
 // See also: SDL_SetWindowPosition()
 // 
 //   window
 //     The window to query.
 //   
 //   x
 //     Pointer to variable for storing the x position, in screen coordinates.
 //     May be NULL.
 //   
 //   y
 //     Pointer to variable for storing the y position, in screen coordinates.
 //     May be NULL.
 //   
func (window *Window) GetPosition() (x int, y int) {
    tmp_x := new(C.int)
    tmp_y := new(C.int)
    C.SDL_GetWindowPosition((*C.SDL_Window)(window), (*C.int)(tmp_x), (*C.int)(tmp_y))
    x = deref_int_ptr(tmp_x)
    y = deref_int_ptr(tmp_y)
    return
}

 // Set the size of a window's client area.
 // 
 // Note: You can't change the size of a fullscreen window, it
 // automatically matches the size of the display mode.
 // 
 //   window
 //     The window to resize.
 //   
 //   w
 //     The width of the window, in screen coordinates. Must be >0.
 //   
 //   h
 //     The height of the window, in screen coordinates. Must be >0.
 //   
 // The window size in screen coordinates may differ from the size in
 // pixels, if the window was created with SDL_WINDOW_ALLOW_HIGHDPI on a
 // platform with high-dpi support (e.g. iOS or OS X). Use
 // SDL_GL_GetDrawableSize() or SDL_GetRendererOutputSize() to get the
 // real client area size in pixels.
 // 
 // See also: SDL_GetWindowSize()
 // 
func (window *Window) SetSize(w int, h int) {
    C.SDL_SetWindowSize((*C.SDL_Window)(window), C.int(w), C.int(h))
}

 // Get the size of a window's client area.
 // 
 //   window
 //     The window to query.
 //   
 //   w
 //     Pointer to variable for storing the width, in screen coordinates. May
 //     be NULL.
 //   
 //   h
 //     Pointer to variable for storing the height, in screen coordinates. May
 //     be NULL.
 //   
 // The window size in screen coordinates may differ from the size in
 // pixels, if the window was created with SDL_WINDOW_ALLOW_HIGHDPI on a
 // platform with high-dpi support (e.g. iOS or OS X). Use
 // SDL_GL_GetDrawableSize() or SDL_GetRendererOutputSize() to get the
 // real client area size in pixels.
 // 
 // See also: SDL_SetWindowSize()
 // 
func (window *Window) GetSize() (w int, h int) {
    tmp_w := new(C.int)
    tmp_h := new(C.int)
    C.SDL_GetWindowSize((*C.SDL_Window)(window), (*C.int)(tmp_w), (*C.int)(tmp_h))
    w = deref_int_ptr(tmp_w)
    h = deref_int_ptr(tmp_h)
    return
}

 // Set the minimum size of a window's client area.
 // 
 // Note: You can't change the minimum size of a fullscreen window, it
 // automatically matches the size of the display mode.
 // 
 // See also: SDL_GetWindowMinimumSize()
 // 
 // See also: SDL_SetWindowMaximumSize()
 // 
 //   window
 //     The window to set a new minimum size.
 //   
 //   min_w
 //     The minimum width of the window, must be >0
 //   
 //   min_h
 //     The minimum height of the window, must be >0
 //   
func (window *Window) SetMinimumSize(min_w int, min_h int) {
    C.SDL_SetWindowMinimumSize((*C.SDL_Window)(window), C.int(min_w), C.int(min_h))
}

 // Get the minimum size of a window's client area.
 // 
 // See also: SDL_GetWindowMaximumSize()
 // 
 // See also: SDL_SetWindowMinimumSize()
 // 
 //   window
 //     The window to query.
 //   
 //   w
 //     Pointer to variable for storing the minimum width, may be NULL
 //   
 //   h
 //     Pointer to variable for storing the minimum height, may be NULL
 //   
func (window *Window) GetMinimumSize() (w int, h int) {
    tmp_w := new(C.int)
    tmp_h := new(C.int)
    C.SDL_GetWindowMinimumSize((*C.SDL_Window)(window), (*C.int)(tmp_w), (*C.int)(tmp_h))
    w = deref_int_ptr(tmp_w)
    h = deref_int_ptr(tmp_h)
    return
}

 // Set the maximum size of a window's client area.
 // 
 // Note: You can't change the maximum size of a fullscreen window, it
 // automatically matches the size of the display mode.
 // 
 // See also: SDL_GetWindowMaximumSize()
 // 
 // See also: SDL_SetWindowMinimumSize()
 // 
 //   window
 //     The window to set a new maximum size.
 //   
 //   max_w
 //     The maximum width of the window, must be >0
 //   
 //   max_h
 //     The maximum height of the window, must be >0
 //   
func (window *Window) SetMaximumSize(max_w int, max_h int) {
    C.SDL_SetWindowMaximumSize((*C.SDL_Window)(window), C.int(max_w), C.int(max_h))
}

 // Get the maximum size of a window's client area.
 // 
 // See also: SDL_GetWindowMinimumSize()
 // 
 // See also: SDL_SetWindowMaximumSize()
 // 
 //   window
 //     The window to query.
 //   
 //   w
 //     Pointer to variable for storing the maximum width, may be NULL
 //   
 //   h
 //     Pointer to variable for storing the maximum height, may be NULL
 //   
func (window *Window) GetMaximumSize() (w int, h int) {
    tmp_w := new(C.int)
    tmp_h := new(C.int)
    C.SDL_GetWindowMaximumSize((*C.SDL_Window)(window), (*C.int)(tmp_w), (*C.int)(tmp_h))
    w = deref_int_ptr(tmp_w)
    h = deref_int_ptr(tmp_h)
    return
}

 // Set the border state of a window.
 // 
 // This will add or remove the window's SDL_WINDOW_BORDERLESS flag and
 // add or remove the border from the actual window. This is a no-op if
 // the window's border already matches the requested state.
 // 
 // Note: You can't change the border state of a fullscreen window.
 // 
 // See also: SDL_GetWindowFlags()
 // 
 //   window
 //     The window of which to change the border state.
 //   
 //   bordered
 //     SDL_FALSE to remove border, SDL_TRUE to add border.
 //   
func (window *Window) SetBordered(bordered bool) {
    C.SDL_SetWindowBordered((*C.SDL_Window)(window), bool2bool(bordered))
}

 // Show a window.
 // 
 // See also: SDL_HideWindow()
 // 
func (window *Window) Show() {
    C.SDL_ShowWindow((*C.SDL_Window)(window))
}

 // Hide a window.
 // 
 // See also: SDL_ShowWindow()
 // 
func (window *Window) Hide() {
    C.SDL_HideWindow((*C.SDL_Window)(window))
}

 // Raise a window above other windows and set the input focus.
func (window *Window) Raise() {
    C.SDL_RaiseWindow((*C.SDL_Window)(window))
}

 // Make a window as large as possible.
 // 
 // See also: SDL_RestoreWindow()
 // 
func (window *Window) Maximize() {
    C.SDL_MaximizeWindow((*C.SDL_Window)(window))
}

 // Minimize a window to an iconic representation.
 // 
 // See also: SDL_RestoreWindow()
 // 
func (window *Window) Minimize() {
    C.SDL_MinimizeWindow((*C.SDL_Window)(window))
}

 // Restore the size and position of a minimized or maximized window.
 // 
 // See also: SDL_MaximizeWindow()
 // 
 // See also: SDL_MinimizeWindow()
 // 
func (window *Window) Restore() {
    C.SDL_RestoreWindow((*C.SDL_Window)(window))
}

 // Set a window's fullscreen state.
 // 
 // Returns: 0 on success, or -1 if setting the display mode failed.
 // 
 // See also: SDL_SetWindowDisplayMode()
 // 
 // See also: SDL_GetWindowDisplayMode()
 // 
func (window *Window) SetFullscreen(flags uint32) (retval int) {
    retval = int(C.SDL_SetWindowFullscreen((*C.SDL_Window)(window), C.Uint32(flags)))
    return
}

 // Get the SDL surface associated with the window.
 // 
 // Returns: The window's framebuffer surface, or NULL on error.
 // 
 // A new surface will be created with the optimal format for the window,
 // if necessary. This surface will be freed when the window is destroyed.
 // 
 // Note: You may not combine this with 3D or the rendering API on this
 // window.
 // 
 // See also: SDL_UpdateWindowSurface()
 // 
 // See also: SDL_UpdateWindowSurfaceRects()
 // 
func (window *Window) GetSurface() (retval *Surface) {
    retval = (*Surface)(unsafe.Pointer(C.SDL_GetWindowSurface((*C.SDL_Window)(window))))
    return
}

 // Copy the window surface to the screen.
 // 
 // Returns: 0 on success, or -1 on error.
 // 
 // See also: SDL_GetWindowSurface()
 // 
 // See also: SDL_UpdateWindowSurfaceRects()
 // 
func (window *Window) UpdateSurface() (retval int) {
    retval = int(C.SDL_UpdateWindowSurface((*C.SDL_Window)(window)))
    return
}


 // Set a window's input grab mode.
 // 
 //   window
 //     The window for which the input grab mode should be set.
 //   
 //   grabbed
 //     This is SDL_TRUE to grab input, and SDL_FALSE to release input.
 //   
 // If the caller enables a grab while another window is currently
 // grabbed, the other window loses its grab in favor of the caller's
 // window.
 // 
 // See also: SDL_GetWindowGrab()
 // 
func (window *Window) SetGrab(grabbed bool) {
    C.SDL_SetWindowGrab((*C.SDL_Window)(window), bool2bool(grabbed))
}

 // Get a window's input grab mode.
 // 
 // Returns: This returns SDL_TRUE if input is grabbed, and SDL_FALSE
 // otherwise.
 // 
 // See also: SDL_SetWindowGrab()
 // 
func (window *Window) GetGrab() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_GetWindowGrab((*C.SDL_Window)(window)))
    return
}

 // Get the window that currently has an input grab enabled.
 // 
 // Returns: This returns the window if input is grabbed, and NULL
 // otherwise.
 // 
 // See also: SDL_SetWindowGrab()
 // 
func GetGrabbedWindow() (retval *Window) {
    retval = (*Window)(unsafe.Pointer(C.SDL_GetGrabbedWindow()))
    return
}

 // Set the brightness (gamma correction) for a window.
 // 
 // Returns: 0 on success, or -1 if setting the brightness isn't
 // supported.
 // 
 // See also: SDL_GetWindowBrightness()
 // 
 // See also: SDL_SetWindowGammaRamp()
 // 
func (window *Window) SetBrightness(brightness float32) (retval int) {
    retval = int(C.SDL_SetWindowBrightness((*C.SDL_Window)(window), C.float(brightness)))
    return
}

 // Get the brightness (gamma correction) for a window.
 // 
 // Returns: The last brightness value passed to SDL_SetWindowBrightness()
 // 
 // See also: SDL_SetWindowBrightness()
 // 
func (window *Window) GetBrightness() (retval float32) {
    retval = float32(C.SDL_GetWindowBrightness((*C.SDL_Window)(window)))
    return
}



 // Provide a callback that decides if a window region has special
 // properties.
 // 
 // Normally windows are dragged and resized by decorations provided by
 // the system window manager (a title bar, borders, etc), but for some
 // apps, it makes sense to drag them from somewhere else inside the
 // window itself; for example, one might have a borderless window that
 // wants to be draggable from any part, or simulate its own title bar,
 // etc.
 // 
 // This function lets the app provide a callback that designates pieces
 // of a given window as special. This callback is run during event
 // processing if we need to tell the OS to treat a region of the window
 // specially; the use of this callback is known as "hit testing."
 // 
 // Mouse input may not be delivered to your application if it is within a
 // special area; the OS will often apply that input to moving the window
 // or resizing the window and not deliver it to the application.
 // 
 // Specifying NULL for a callback disables hit-testing. Hit-testing is
 // disabled by default.
 // 
 // Platforms that don't support this functionality will return -1
 // unconditionally, even if you're attempting to disable hit-testing.
 // 
 // Your callback may fire at any time, and its firing does not indicate
 // any specific behavior (for example, on Windows, this certainly might
 // fire when the OS is deciding whether to drag your window, but it fires
 // for lots of other reasons, too, some unrelated to anything you
 // probably care about and when the mouse isn't actually at the location
 // it is testing). Since this can fire at any time, you should try to
 // keep your callback efficient, devoid of allocations, etc.
 // 
 // Returns: 0 on success, -1 on error (including unsupported).
 // 
 //   window
 //     The window to set hit-testing on.
 //   
 //   callback
 //     The callback to call when doing a hit-test.
 //   
 //   callback_data
 //     An app-defined void pointer passed to the callback.
 //   
func (window *Window) SetHitTest(callback HitTest, callback_data uintptr) (retval int) {
    retval = int(C.SDL_SetWindowHitTest((*C.SDL_Window)(window), C.SDL_HitTest(callback), unsafe.Pointer(callback_data)))
    return
}

 // Destroy a window.
func (window *Window) Destroy() {
    C.SDL_DestroyWindow((*C.SDL_Window)(window))
}

 // Returns whether the screensaver is currently enabled (default on).
 // 
 // See also: SDL_EnableScreenSaver()
 // 
 // See also: SDL_DisableScreenSaver()
 // 
func IsScreenSaverEnabled() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_IsScreenSaverEnabled())
    return
}

 // Allow the screen to be blanked by a screensaver.
 // 
 // See also: SDL_IsScreenSaverEnabled()
 // 
 // See also: SDL_DisableScreenSaver()
 // 
func EnableScreenSaver() {
    C.SDL_EnableScreenSaver()
}

 // Prevent the screen from being blanked by a screensaver.
 // 
 // See also: SDL_IsScreenSaverEnabled()
 // 
 // See also: SDL_EnableScreenSaver()
 // 
func DisableScreenSaver() {
    C.SDL_DisableScreenSaver()
}
