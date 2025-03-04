package internal

type SystemConfig struct {
	Namespace string
	Prefix    string
}

var SystemConfigs = map[string]SystemConfig{
	"Atari2600": {
		Namespace: "openemu.system.2600",
		Prefix:    "OE2600",
	},
	"SegaMegaDrive32x": {
		Namespace: "openemu.system.32x",
		Prefix:    "OESega32X",
	},
	"Atari5200": {
		Namespace: "openemu.system.5200",
		Prefix:    "OE5200",
	},
	"Atari7800": {
		Namespace: "openemu.system.7800",
		Prefix:    "OE7800",
	},
	"ColecoVision": {
		Namespace: "openemu.system.colecovision",
		Prefix:    "OEColecoVision",
	},
	"FamiconDiskSystem": {
		Namespace: "openemu.system.fds",
		Prefix:    "OEFDS",
	},
	"GameBoy": {
		Namespace: "openemu.system.gb",
		Prefix:    "OEGB",
	},
	"GameBoyAdvance": {
		Namespace: "openemu.system.gba",
		Prefix:    "OEGBA",
	},
	"GameCube": {
		Namespace: "openemu.system.gc",
		Prefix:    "OEGC",
	},
	"GameGear": {
		Namespace: "openemu.system.gg",
		Prefix:    "OEGG",
	},
	"Intellivision": {
		Namespace: "openemu.system.intellivision",
		Prefix:    "OEIntellivision",
	},
	"Lynx": {
		Namespace: "openemu.system.lynx",
		Prefix:    "OELynx",
	},
	"N64": {
		Namespace: "openemu.system.n64",
		Prefix:    "OEN64",
	},
	"NintendoDS": {
		Namespace: "openemu.system.nds",
		Prefix:    "OENDS",
	},
	"Nintendo": {
		Namespace: "openemu.system.nes",
		Prefix:    "OENES",
	},
	"NeoGeoPocket": {
		Namespace: "openemu.system.ngp",
		Prefix:    "OENGP",
	},
	"Odyssey2": {
		Namespace: "openemu.system.odyssey2",
		Prefix:    "OEOdyssey2",
	},
	"PCE": {
		Namespace: "openemu.system.pce",
		Prefix:    "OEPCE",
	},
	"PCECD": {
		Namespace: "openemu.system.pcecd",
		Prefix:    "OEPCECD",
	},
	"PCFX": {
		Namespace: "openemu.system.pcfx",
		Prefix:    "OEPCFX",
	},
	"PSP": {
		Namespace: "openemu.system.psp",
		Prefix:    "OEPSP",
	},
	"PlayStation": {
		Namespace: "openemu.system.psx",
		Prefix:    "OEPSX",
	},
	"SegaSaturn": {
		Namespace: "openemu.system.saturn",
		Prefix:    "OESaturn",
	},
	"SegaMegaCD": {
		Namespace: "openemu.system.scd",
		Prefix:    "OESegaCD",
	},
	"SG": {
		Namespace: "openemu.system.sg",
		Prefix:    "OEGenesis",
	},
	"SG1000": {
		Namespace: "openemu.system.sg1000",
		Prefix:    "OESG1000",
	},
	"SegaMasterSystem": {
		Namespace: "openemu.system.sms",
		Prefix:    "OESMS",
	},
	"SuperNintendo": {
		Namespace: "openemu.system.snes",
		Prefix:    "OESNES",
	},
	"VirtualBoy": {
		Namespace: "openemu.system.vb",
		Prefix:    "OEVB",
	},
	"Vectrex": {
		Namespace: "openemu.system.vectrex",
		Prefix:    "OEVectrex",
	},
	"WonderSwan": {
		Namespace: "openemu.system.ws",
		Prefix:    "OEWS",
	},
}
