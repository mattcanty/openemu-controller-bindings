# OpenEmu Controller Mappings Tool

> [!IMPORTANT]
> 🕹️ I don't want anything here to detract away from what a fantastic job the team have done with [OpenEmu](https://openemu.org). I can't highly recommend it enough. Gone are the days where this was a challenging endeavour.
> [!INFO]
> I think it would be useful to have a simple way of ripping the controller mappings instead of manually writing out the game files. I'll do this one day soon.

## The Problem: No Per-Game Controller Mappings in OpenEmu

OpenEmu lacks support for per-game controller mappings. I found this to be an issue _immediately_. It is particularly glaring when playing games like **Perfect Dark** and **1080 Snowboarding** on N64. You have a go and come let me know whether you can play them on defaults without your brain overheating.

So, I made a plaster to cover this wound: a CLI tool to let you save, load, and list controller mappings. While it’s not the perfect solution, it’s functional and gets the job done.

Nothing clever is happening here. It literally copies files.

## Installation

1. Head over to the **[Releases page](https://github.com/mattcanty/openemu-controller-bindings/releases)** and download the latest version for your platform.
2. Extract the downloaded file.
3. Place the executable somewhere in your system's PATH (e.g., `/usr/local/bin` or `C:\Program Files\`).

## Usage

### Backup

```bash
ocb backup
```

### List

```bash
ocb list
```

### Load

```bash
ocb load <NAME>
```

> [!IMPORTANT]
> Exising controller bindings will be erased. Take a backup first if you need to.
> [!IMPORTANT]
> OpenEmu will not automatically start working with these mappings. You must restart OpenEmu.

## References

- [OpenEmu/OpenEmu#206](https://sgithub.com/OpenEmu/OpenEmu/issues/206)
- [OpenEmu/OpenEmu#1005](https://github.com/OpenEmu/OpenEmu/issues/1005)
- [OpenEmu/OpenEmu#1033](https://github.com/OpenEmu/OpenEmu/issues/1033)
- [OpenEmu/OpenEmu#1944](https://github.com/OpenEmu/OpenEmu/issues/1944)
- [r/OpenEmu:different_controller_bindings_for_different_games](https://www.reddit.com/r/OpenEmu/comments/1cd2ate/different_controller_bindings_for_different_games/)
- [r/OpenEmu:saving_control_schemes](https://www.reddit.com/r/OpenEmu/comments/fxepjx/saving_control_schemes/)
- [r/OpenEmu:have_different_button_mapping_per_game_in](https://www.reddit.com/r/OpenEmu/comments/d2uhiw/have_different_button_mapping_per_game_in/)
