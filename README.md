# Identigo

Identigo is a simple Go function for creating two-color identicons from a given Bytestring. It's based on Aaron Lasseigne's Clojure program (https://gist.github.com/AaronLasseigne/6255278) with some modifications.

# Usage

Identigo takes the following inputs:
1. bytestring - ([]Byte) Bytestream of input data
2. tilesPerSide - (Int) of number tile resolution per rectangle side
3. multiplier - (Int) To make a large icon the function multiplies the tilesPerSide with the multiplier and scales the rendering accordingly
4. Background Color - (string) Supported colors are given below
5. Foreground Color - (string) Supported colors are given below

It returns an image object of the size (tilesPerSide * multiplier)

## Supported Colors

 black
 navy
 blue
 green
 teal
 lime
 aqua
 maroon
 purple
 olive
 gray
 silver
 red
 fuchsia
 yellow
 white
