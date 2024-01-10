# Map Maker

> The game provides a very simple map maker that allows you to create your own maps.

## Creating a map

All maps are stored in the [`assets/maps`](../assets/maps) folder. To create a new map, simply create a new file in that folder.

## Format

Each map file should contain 2 sections: `metadata` and `map` divided by a line with this exact content: `---`.

The `metadata` section contains information about the map, such as its name, author, and description. It uses the same format as `.env` files (key-value pairs). However, right now this information is not used by the game.

The `map` section contains the actual map data. It must be composed of lines with the same length, and each character in the line must be one of the following:

- `_` (underscore): empty space
- `x`: ground
- `^`: spikes
- `s`: spring

## Example

This file create 12x5 map.

```
name=My Map
author=Me
description=This is my map
---
____________
____________
____xxxx___X
__xxxxxx___x
xxxxxxxxxxxx
```
