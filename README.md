# OpenEmu Controller Mappings Tool

## The Problem: No Per-Game Controller Mappings in OpenEmu

OpenEmu lacks support for per-game controller mappings. This gap becomes glaring when playing games like **Perfect Dark** and **1080 Snowboarding** on N64. Just have a go and come let me know you can play them without your brain overheating.

So, I made a plaster to cover this wound: a CLI tool to let you save, load, and list controller mappings. While it’s not the perfect solution, it’s functional and gets the job done.

## Installation

1. Head over to the **[Releases page](https://github.com/mattcanty/openemu-controller-bindings/releases)** and download the latest version for your platform.
2. Extract the downloaded file.
3. Place the executable somewhere in your system's PATH (e.g., `/usr/local/bin` or `C:\Program Files\`).

## Usage

```bash
ocb save <NAME>
ocb load <NAME>
ocb list
```
