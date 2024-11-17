# OpenEmu Controller Mappings Tool

## The Problem: No Per-Game Controller Mappings in OpenEmu

OpenEmu lacks support for per-game controller mappings. I found this to be an issue _immediately_. It is particularly glaring when playing games like **Perfect Dark** and **1080 Snowboarding** on N64. You have a go and come let me know whether you can play them on defaults without your brain overheating.

So, I made a plaster to cover this wound: a CLI tool to let you save, load, and list controller mappings. While it’s not the perfect solution, it’s functional and gets the job done.

Nothing clever is happening here. It literally copies files.

## Installation

1. Head over to the **[Releases page](https://github.com/mattcanty/openemu-controller-bindings/releases)** and download the latest version for your platform.
2. Extract the downloaded file.
3. Place the executable somewhere in your system's PATH (e.g., `/usr/local/bin` or `C:\Program Files\`).

## Usage

Edit your mappings using OpenEmu as normal. Then save the settings:

```bash
ocb save <NAME>
```

This will create a copy of the configuration with the name provided. You can then load that configuration back using load:

```bash
ocb load <NAME>
```

> [!IMPORTANT]  
> OpenEmu will not automatically start working with these mappings. You must restart OpenEmu.

```bash
ocb list
```

## References

- [OpenEmu/OpenEmu#206](https://github.com/OpenEmu/OpenEmu/issues/206)
- [OpenEmu/OpenEmu#1005](https://github.com/OpenEmu/OpenEmu/issues/1005)
- [OpenEmu/OpenEmu#1033](https://github.com/OpenEmu/OpenEmu/issues/1033)
- [OpenEmu/OpenEmu#1944](https://github.com/OpenEmu/OpenEmu/issues/1944)
- https://www.reddit.com/r/OpenEmu/comments/1cd2ate/different_controller_bindings_for_different_games/
- https://www.reddit.com/r/OpenEmu/comments/fxepjx/saving_control_schemes/
- https://www.reddit.com/r/OpenEmu/comments/d2uhiw/have_different_button_mapping_per_game_in/
