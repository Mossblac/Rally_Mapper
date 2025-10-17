# Rally_Mapper
**A Rally Course Generator**

![RallyMapper logo](./images/rally_mapper_logo_smaller.png)

## A Rally:

- A skill and strategy based time trial race.  

A rally course is a **timed** course consisting of three components:
1. A Course  
2. Obstacles *
3. and Punches *

`* Optional

**The Track** can be linear or a loop. changed with a toggle on the main menu.   
**The obstacles** are objects within the course that require the performance of certain riding skills in order to navigate successfully and even more skill to overcome at speed  
**Punches** are triggered checkpoints located throughout the Track that, when activated, present a bonus time reduction on your overall time.  

This Application was made using Fyne UI and Go programming language.
Track Data is stored locally in JSON format allowing for the eventual implementation
of an online database. 


It is intended to be the first part of a larger application that will run 
Rallys based on the map information created with the Rally Mapper. 
Automatically calculating times and logging multiple runs for each contestant,  
as well as creating sharable videos of their performance. 

-This is a work in progress- I hope to improve the UI and eventually add a custom editing feature at a later date. 

**required to build app**
1. Go Language: Fyne is built in Go, so you need a working Go installation. Version 1.17 or later is recommended.

2. C Compiler: Fyne utilizes CGo to interact with system graphics drivers (like GLFW for desktop platforms), which necessitates a C compiler.

3. OS specific :
- Linux: GCC is typically suitable. You might also need libgl1-mesa-dev and xorg-dev packages on Debian/Ubuntu-based systems.
- macOS: Xcode, which includes the necessary command-line tools, provides the C compiler.
- Windows: Options include MSYS2 with MingW-w64, TDM-GCC, or Cygwin.

4. System Graphics Driver: Fyne relies on your system's graphics driver to render the UI. This is typically handled by the C compiler and associated libraries. 

**Optional but Recommended:**
fyne.io/fyne/v2: The Fyne toolkit itself, which you install using go get fyne.io/fyne/v2@latest.
Fyne Setup tool: A utility to check your development environment and assist with troubleshooting.
fyne-cross: For cross-compiling your Fyne application to different platforms, particularly useful for mobile and other OS targets from a single development machine.
