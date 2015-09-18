# Gogam - A game engine written in Go

Built on top of the experimental mobile package

## Warning: Gogam is still at a very early stage, much of it is expected to change and therefore certain functionality may end up working differently in the future

### Current Features:

* Engine:
	* Simplified process to get up and running
* Sprites:
	* Animations
	* Built-in movement functions
		* Translate/MoveTo
		* Rotation (in place)
		* Scale
	
	
### Todos:
* General:
	* Consider reworking each part into its own folder for easier dependency management
	* Uncouple files
	* Look into web / console support options
	* Reconsider the public/private methods of the current engine and sprite implementations
* Engine:
	* Work in one resolution that scales automatically for different devices
	* Try to integrate app package into engine
		* Might require writing wrappers for events
	* Check for areas to improve performance with concurrency
* Sprites:
	* Simplify animation process
	* More options for movement
		* Choosing what to rotate around
		* Flip horizontal/vertical
	* Collision detection
* Buttons (not started):
	* onClick (could allow for passing a function here)
	* Default, Hover, onClick images
* Text (not start):
	* Choose font type, style, size
	* Move around like sprites
* Scenes (not started):
	* Easily manage what entities are within each scene
	* Allow for easier debugging
* Music (not started):
	* Need to see how it is currently implemented
	* Play music between scenes
	* Play sound effects on top of each other
* Game center integration (not started):
	* Need to see if there is anything in place currently
* Particles (not started):
	* Look into how this could be worked in
* Physics (not started)
* Ads (not started)