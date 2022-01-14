---
title: AppImage's in Ubuntu
layout: post
tags: linux
---
`AppImage`s are basically self contained scripts that can be run without installation.

We first need to make it executable via `chmod +x name.AppImage`. 

Then we can run it via `./name.AppImage`.

## Integrating with Ubuntu
It's useful to integrate the `AppImage`  into Ubuntu itself we can do this via the following steps:

1. Move the `AppImage` into a directory in your `PATH` (for me `~/bin`)
2. Create a new file in `~/.local/share/applications` with called `<name-of-software>.desktop`
3. Fill in the file with the following details:

```ini
[Desktop Entry]
Name=$name
Exec=/home/yourhome/bin/$name.AppImage
Terminal=false
Type=Application
Categories=Development
```
4. Done!



