# Rally_Mapper : A Rally Course Generator

![TrackGif](./demo.gif) [](./demo.gif)

***

## Rally: A skill and strategy based time trial race  

While I thoroughly enjoy riding both my **Fungineer FunWheel X7** Rally board and my **FutureMotion OneWheel pintX**
Rally board almost every day, I can't say I have ever enjoyed watching a competition involving them. Not only are there no events that can effectivly place these boards against one another, the few races and trick competitions I've seen have all failed to capture my attention (and bored me to tears...)

These are not mountain bikes, dirt bikes or off-road skateboards, Why are all the competitions on tracks built for that equipment?

## I want to change to game, literally....

> These Tracks are specifically designed to take advantage of the  
> capability of a Rally Board or EUC. Which are handsfree operated, partially self balancing,  
> personal electric vehicles or P.E.V.s

***

## 🚀 Quick Start

### 1. Install Rally_Mapper using the Go toolchain

```bash
# Install Rally_Mapper
go install github.com/Mossblac/Rally_Mapper@latest
```

***

## Getting Started

Choose the type of map  
Choose the size of the map  
Choose your obstacles  

Rally Mapper will generate a custom Track, Place Obstacles and Punches, and provide the calculated time reductions gained from each.  

## Usage

## A Rally Track is a **timed** course consisting of three components:

1. **A Course**  
2. **Obstacles**  
3. **and Punches**  

>**The Track** can be linear or a loop. changed with a toggle on the main menu.  
>**The Obstacles** are objects within the course that require the performance of certain riding skills in order to navigate successfully and even more skill to overcome at speed
>**Punches** are triggered checkpoints located throughout the Track that, when activated, grant a bonus time reduction to your run.

```bash
from the Track screen you can regenerate new maps using the same parameters, click save to store your favorites
and reload using the load menu on the homescreen 
```

*** 

**This is a work in progress** Here are some of the upcoming features and planned improvements:

> 1. ~~animation updates to place obstacles and punches after track is placed~~
> 2. delete saves from within the app
> 3. pinch zoom and rotate the map smoothly
> 4. run the track animation that takes you through the course at the average speed.
> 5. ability to upload and share Tracks
> 6. Login and Multiple user support
> 7. manual obstacle and punch placement
> 8. Rally Mapper Mobile
> 9. minor UI improvements.

***

This Application was made using **Fyne** UI and **Go** programming language.
Track Data is stored locally in JSON format allowing for the eventual implementation
of an online database.  

It is intended to be the first part of a larger application that will run  
Rallys based on the Track information created with the Rally Mapper.  
Automatically calculating times and logging multiple runs for each contestant,  
as well as creating sharable videos of their performance.  

