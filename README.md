# OpenEmu Per-Game Controller Mappings CLI

## The Problem: No Per-Game Controller Mappings in OpenEmu

It’s baffling that OpenEmu, a tool so polished in many ways, lacks support for per-game controller mappings. This gap becomes glaring when playing games like **Perfect Dark** and **1080 Snowboarding**, which demand unique controller setups for an optimal experience. 

So, I made a plaster to cover this wound: a CLI tool to let you save, load, and manage controller mappings for specific games. While it’s not the perfect solution, it’s functional and gets the job done.

---

## Installation

1. Head over to the **[Releases page](https://github.com/your-repo/releases)** and download the latest version for your platform.
2. Extract the downloaded file.
3. Place the executable somewhere in your system's PATH (e.g., `/usr/local/bin` or `C:\Program Files\`).

---

## Usage

### Save a Mapping
Save your current controller mapping for a specific game:

```bash
oe-bindings save <NAME>
