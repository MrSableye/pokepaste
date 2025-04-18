package main

var pokemonData = map[string]map[string]interface{}{
	"Missingno": {
		"id":   uint(0),
		"form": uint(0),
		"type": "???",
	},
	"Bulbasaur": {
		"id":   uint(1),
		"form": uint(0),
		"type": "grass",
	},
	"Ivysaur": {
		"id":   uint(2),
		"form": uint(0),
		"type": "grass",
	},
	"Venusaur": {
		"id":   uint(3),
		"form": uint(0),
		"type": "grass",
	},
	"Venusaur-Mega": {
		"id":   uint(3),
		"form": uint(1),
		"type": "grass",
	},
	"Venusaur-Gmax": {
		"id":   uint(3),
		"form": uint(2),
		"type": "grass",
	},
	"Charmander": {
		"id":   uint(4),
		"form": uint(0),
		"type": "fire",
	},
	"Charmeleon": {
		"id":   uint(5),
		"form": uint(0),
		"type": "fire",
	},
	"Charizard": {
		"id":   uint(6),
		"form": uint(0),
		"type": "fire",
	},
	"Charizard-Mega-X": {
		"id":   uint(6),
		"form": uint(1),
		"type": "fire",
	},
	"Charizard-Mega-Y": {
		"id":   uint(6),
		"form": uint(2),
		"type": "fire",
	},
	"Charizard-Gmax": {
		"id":   uint(6),
		"form": uint(3),
		"type": "fire",
	},
	"Squirtle": {
		"id":   uint(7),
		"form": uint(0),
		"type": "water",
	},
	"Wartortle": {
		"id":   uint(8),
		"form": uint(0),
		"type": "water",
	},
	"Blastoise": {
		"id":   uint(9),
		"form": uint(0),
		"type": "water",
	},
	"Blastoise-Mega": {
		"id":   uint(9),
		"form": uint(1),
		"type": "water",
	},
	"Blastoise-Gmax": {
		"id":   uint(9),
		"form": uint(2),
		"type": "water",
	},
	"Caterpie": {
		"id":   uint(10),
		"form": uint(0),
		"type": "bug",
	},
	"Metapod": {
		"id":   uint(11),
		"form": uint(0),
		"type": "bug",
	},
	"Butterfree": {
		"id":   uint(12),
		"form": uint(0),
		"type": "bug",
	},
	"Butterfree-Gmax": {
		"id":   uint(12),
		"form": uint(1),
		"type": "bug",
	},
	"Weedle": {
		"id":   uint(13),
		"form": uint(0),
		"type": "bug",
	},
	"Kakuna": {
		"id":   uint(14),
		"form": uint(0),
		"type": "bug",
	},
	"Beedrill": {
		"id":   uint(15),
		"form": uint(0),
		"type": "bug",
	},
	"Beedrill-Mega": {
		"id":   uint(15),
		"form": uint(1),
		"type": "bug",
	},
	"Pidgey": {
		"id":   uint(16),
		"form": uint(0),
		"type": "normal",
	},
	"Pidgeotto": {
		"id":   uint(17),
		"form": uint(0),
		"type": "normal",
	},
	"Pidgeot": {
		"id":   uint(18),
		"form": uint(0),
		"type": "normal",
	},
	"Pidgeot-Mega": {
		"id":   uint(18),
		"form": uint(1),
		"type": "normal",
	},
	"Rattata": {
		"id":   uint(19),
		"form": uint(0),
		"type": "normal",
	},
	"Rattata-Alola": {
		"id":   uint(19),
		"form": uint(1),
		"type": "dark",
	},
	"Raticate": {
		"id":   uint(20),
		"form": uint(0),
		"type": "normal",
	},
	"Raticate-Alola": {
		"id":   uint(20),
		"form": uint(1),
		"type": "dark",
	},
	"Raticate-Alola-Totem": {
		"id":   uint(20),
		"form": uint(1),
		"type": "dark",
	},
	"Spearow": {
		"id":   uint(21),
		"form": uint(0),
		"type": "normal",
	},
	"Fearow": {
		"id":   uint(22),
		"form": uint(0),
		"type": "normal",
	},
	"Ekans": {
		"id":   uint(23),
		"form": uint(0),
		"type": "poison",
	},
	"Arbok": {
		"id":   uint(24),
		"form": uint(0),
		"type": "poison",
	},
	"Pikachu": {
		"id":   uint(25),
		"form": uint(0),
		"type": "electric",
	},
	"Pikachu-Gmax": {
		"id":   uint(25),
		"form": uint(1),
		"type": "electric",
	},
	"Raichu": {
		"id":   uint(26),
		"form": uint(0),
		"type": "electric",
	},
	"Raichu-Alola": {
		"id":   uint(26),
		"form": uint(1),
		"type": "electric",
	},
	"Sandshrew": {
		"id":   uint(27),
		"form": uint(0),
		"type": "ground",
	},
	"Sandshrew-Alola": {
		"id":   uint(27),
		"form": uint(1),
		"type": "ice",
	},
	"Sandslash": {
		"id":   uint(28),
		"form": uint(0),
		"type": "ground",
	},
	"Sandslash-Alola": {
		"id":   uint(28),
		"form": uint(1),
		"type": "ice",
	},
	"Nidoran-F": {
		"id":   uint(29),
		"form": uint(0),
		"type": "poison",
	},
	"Nidorina": {
		"id":   uint(30),
		"form": uint(0),
		"type": "poison",
	},
	"Nidoqueen": {
		"id":   uint(31),
		"form": uint(0),
		"type": "poison",
	},
	"Nidoran-M": {
		"id":   uint(32),
		"form": uint(0),
		"type": "poison",
	},
	"Nidorino": {
		"id":   uint(33),
		"form": uint(0),
		"type": "poison",
	},
	"Nidoking": {
		"id":   uint(34),
		"form": uint(0),
		"type": "poison",
	},
	"Clefairy": {
		"id":   uint(35),
		"form": uint(0),
		"type": "fairy",
	},
	"Clefable": {
		"id":   uint(36),
		"form": uint(0),
		"type": "fairy",
	},
	"Vulpix": {
		"id":   uint(37),
		"form": uint(0),
		"type": "fire",
	},
	"Vulpix-Alola": {
		"id":   uint(37),
		"form": uint(1),
		"type": "ice",
	},
	"Ninetales": {
		"id":   uint(38),
		"form": uint(0),
		"type": "fire",
	},
	"Ninetales-Alola": {
		"id":   uint(38),
		"form": uint(1),
		"type": "ice",
	},
	"Jigglypuff": {
		"id":   uint(39),
		"form": uint(0),
		"type": "normal",
	},
	"Wigglytuff": {
		"id":   uint(40),
		"form": uint(0),
		"type": "normal",
	},
	"Zubat": {
		"id":   uint(41),
		"form": uint(0),
		"type": "poison",
	},
	"Golbat": {
		"id":   uint(42),
		"form": uint(0),
		"type": "poison",
	},
	"Oddish": {
		"id":   uint(43),
		"form": uint(0),
		"type": "grass",
	},
	"Gloom": {
		"id":   uint(44),
		"form": uint(0),
		"type": "grass",
	},
	"Vileplume": {
		"id":   uint(45),
		"form": uint(0),
		"type": "grass",
	},
	"Paras": {
		"id":   uint(46),
		"form": uint(0),
		"type": "bug",
	},
	"Parasect": {
		"id":   uint(47),
		"form": uint(0),
		"type": "bug",
	},
	"Venonat": {
		"id":   uint(48),
		"form": uint(0),
		"type": "bug",
	},
	"Venomoth": {
		"id":   uint(49),
		"form": uint(0),
		"type": "bug",
	},
	"Diglett": {
		"id":   uint(50),
		"form": uint(0),
		"type": "ground",
	},
	"Diglett-Alola": {
		"id":   uint(50),
		"form": uint(1),
		"type": "ground",
	},
	"Dugtrio": {
		"id":   uint(51),
		"form": uint(0),
		"type": "ground",
	},
	"Dugtrio-Alola": {
		"id":   uint(51),
		"form": uint(1),
		"type": "ground",
	},
	"Meowth": {
		"id":   uint(52),
		"form": uint(0),
		"type": "normal",
	},
	"Meowth-Alola": {
		"id":   uint(52),
		"form": uint(1),
		"type": "dark",
	},
	"Meowth-Galar": {
		"id":   uint(52),
		"form": uint(2),
		"type": "steel",
	},
	"Meowth-Gmax": {
		"id":   uint(52),
		"form": uint(3),
		"type": "steel",
	},
	"Persian": {
		"id":   uint(53),
		"form": uint(0),
		"type": "normal",
	},
	"Persian-Alola": {
		"id":   uint(53),
		"form": uint(1),
		"type": "dark",
	},
	"Psyduck": {
		"id":   uint(54),
		"form": uint(0),
		"type": "water",
	},
	"Golduck": {
		"id":   uint(55),
		"form": uint(0),
		"type": "water",
	},
	"Mankey": {
		"id":   uint(56),
		"form": uint(0),
		"type": "fighting",
	},
	"Primeape": {
		"id":   uint(57),
		"form": uint(0),
		"type": "fighting",
	},
	"Growlithe": {
		"id":   uint(58),
		"form": uint(0),
		"type": "fire",
	},
	"Arcanine": {
		"id":   uint(59),
		"form": uint(0),
		"type": "fire",
	},
	"Poliwag": {
		"id":   uint(60),
		"form": uint(0),
		"type": "water",
	},
	"Poliwhirl": {
		"id":   uint(61),
		"form": uint(0),
		"type": "water",
	},
	"Poliwrath": {
		"id":   uint(62),
		"form": uint(0),
		"type": "water",
	},
	"Abra": {
		"id":   uint(63),
		"form": uint(0),
		"type": "psychic",
	},
	"Kadabra": {
		"id":   uint(64),
		"form": uint(0),
		"type": "psychic",
	},
	"Alakazam": {
		"id":   uint(65),
		"form": uint(0),
		"type": "psychic",
	},
	"Alakazam-Mega": {
		"id":   uint(65),
		"form": uint(1),
		"type": "psychic",
	},
	"Machop": {
		"id":   uint(66),
		"form": uint(0),
		"type": "fighting",
	},
	"Machoke": {
		"id":   uint(67),
		"form": uint(0),
		"type": "fighting",
	},
	"Machamp": {
		"id":   uint(68),
		"form": uint(0),
		"type": "fighting",
	},
	"Machamp-Gmax": {
		"id":   uint(68),
		"form": uint(1),
		"type": "fighting",
	},
	"Bellsprout": {
		"id":   uint(69),
		"form": uint(0),
		"type": "grass",
	},
	"Weepinbell": {
		"id":   uint(70),
		"form": uint(0),
		"type": "grass",
	},
	"Victreebel": {
		"id":   uint(71),
		"form": uint(0),
		"type": "grass",
	},
	"Tentacool": {
		"id":   uint(72),
		"form": uint(0),
		"type": "water",
	},
	"Tentacruel": {
		"id":   uint(73),
		"form": uint(0),
		"type": "water",
	},
	"Geodude": {
		"id":   uint(74),
		"form": uint(0),
		"type": "rock",
	},
	"Geodude-Alola": {
		"id":   uint(74),
		"form": uint(1),
		"type": "rock",
	},
	"Graveler": {
		"id":   uint(75),
		"form": uint(0),
		"type": "rock",
	},
	"Graveler-Alola": {
		"id":   uint(75),
		"form": uint(1),
		"type": "rock",
	},
	"Golem": {
		"id":   uint(76),
		"form": uint(0),
		"type": "rock",
	},
	"Golem-Alola": {
		"id":   uint(76),
		"form": uint(1),
		"type": "rock",
	},
	"Ponyta": {
		"id":   uint(77),
		"form": uint(0),
		"type": "fire",
	},
	"Ponyta-Galar": {
		"id":   uint(77),
		"form": uint(1),
		"type": "psychic",
	},
	"Rapidash": {
		"id":   uint(78),
		"form": uint(0),
		"type": "fire",
	},
	"Rapidash-Galar": {
		"id":   uint(78),
		"form": uint(1),
		"type": "psychic",
	},
	"Slowpoke": {
		"id":   uint(79),
		"form": uint(0),
		"type": "water",
	},
	"Slowpoke-Galar": {
		"id":   uint(79),
		"form": uint(1),
		"type": "psychic",
	},
	"Slowbro": {
		"id":   uint(80),
		"form": uint(0),
		"type": "water",
	},
	"Slowbro-Mega": {
		"id":   uint(80),
		"form": uint(1),
		"type": "water",
	},
	"Slowbro-Galar": {
		"id":   uint(80),
		"form": uint(2),
		"type": "poison",
	},
	"Magnemite": {
		"id":   uint(81),
		"form": uint(0),
		"type": "electric",
	},
	"Magneton": {
		"id":   uint(82),
		"form": uint(0),
		"type": "electric",
	},
	"Farfetch'd": {
		"id":   uint(83),
		"form": uint(0),
		"type": "normal",
	},
	"Farfetch&#x27;d": {
		"id":   uint(83),
		"form": uint(0),
		"type": "normal",
	},
	"Farfetch’d": {
		"id":   uint(83),
		"form": uint(0),
		"type": "normal",
	},
	"Farfetch'd-Galar": {
		"id":   uint(83),
		"form": uint(1),
		"type": "fighting",
	},
	"Farfetch&#x27;d-Galar": {
		"id":   uint(83),
		"form": uint(1),
		"type": "fighting",
	},
	"Farfetch’d-Galar": {
		"id":   uint(83),
		"form": uint(1),
		"type": "fighting",
	},
	"Doduo": {
		"id":   uint(84),
		"form": uint(0),
		"type": "normal",
	},
	"Dodrio": {
		"id":   uint(85),
		"form": uint(0),
		"type": "normal",
	},
	"Seel": {
		"id":   uint(86),
		"form": uint(0),
		"type": "water",
	},
	"Dewgong": {
		"id":   uint(87),
		"form": uint(0),
		"type": "water",
	},
	"Grimer": {
		"id":   uint(88),
		"form": uint(0),
		"type": "poison",
	},
	"Grimer-Alola": {
		"id":   uint(88),
		"form": uint(1),
		"type": "poison",
	},
	"Muk": {
		"id":   uint(89),
		"form": uint(0),
		"type": "poison",
	},
	"Muk-Alola": {
		"id":   uint(89),
		"form": uint(1),
		"type": "poison",
	},
	"Shellder": {
		"id":   uint(90),
		"form": uint(0),
		"type": "water",
	},
	"Cloyster": {
		"id":   uint(91),
		"form": uint(0),
		"type": "water",
	},
	"Gastly": {
		"id":   uint(92),
		"form": uint(0),
		"type": "ghost",
	},
	"Haunter": {
		"id":   uint(93),
		"form": uint(0),
		"type": "ghost",
	},
	"Gengar": {
		"id":   uint(94),
		"form": uint(0),
		"type": "ghost",
	},
	"Gengar-Mega": {
		"id":   uint(94),
		"form": uint(1),
		"type": "ghost",
	},
	"Gengar-Gmax": {
		"id":   uint(94),
		"form": uint(2),
		"type": "ghost",
	},
	"Onix": {
		"id":   uint(95),
		"form": uint(0),
		"type": "rock",
	},
	"Drowzee": {
		"id":   uint(96),
		"form": uint(0),
		"type": "psychic",
	},
	"Hypno": {
		"id":   uint(97),
		"form": uint(0),
		"type": "psychic",
	},
	"Krabby": {
		"id":   uint(98),
		"form": uint(0),
		"type": "water",
	},
	"Kingler": {
		"id":   uint(99),
		"form": uint(0),
		"type": "water",
	},
	"Kingler-Gmax": {
		"id":   uint(99),
		"form": uint(1),
		"type": "water",
	},
	"Voltorb": {
		"id":   uint(100),
		"form": uint(0),
		"type": "electric",
	},
	"Electrode": {
		"id":   uint(101),
		"form": uint(0),
		"type": "electric",
	},
	"Exeggcute": {
		"id":   uint(102),
		"form": uint(0),
		"type": "grass",
	},
	"Exeggutor": {
		"id":   uint(103),
		"form": uint(0),
		"type": "grass",
	},
	"Exeggutor-Alola": {
		"id":   uint(103),
		"form": uint(1),
		"type": "grass",
	},
	"Cubone": {
		"id":   uint(104),
		"form": uint(0),
		"type": "ground",
	},
	"Marowak": {
		"id":   uint(105),
		"form": uint(0),
		"type": "ground",
	},
	"Marowak-Alola": {
		"id":   uint(105),
		"form": uint(1),
		"type": "fire",
	},
	"Marowak-Alola-Totem": {
		"id":   uint(105),
		"form": uint(1),
		"type": "fire",
	},
	"Hitmonlee": {
		"id":   uint(106),
		"form": uint(0),
		"type": "fighting",
	},
	"Hitmonchan": {
		"id":   uint(107),
		"form": uint(0),
		"type": "fighting",
	},
	"Lickitung": {
		"id":   uint(108),
		"form": uint(0),
		"type": "normal",
	},
	"Koffing": {
		"id":   uint(109),
		"form": uint(0),
		"type": "poison",
	},
	"Weezing": {
		"id":   uint(110),
		"form": uint(0),
		"type": "poison",
	},
	"Weezing-Galar": {
		"id":   uint(110),
		"form": uint(1),
		"type": "poison",
	},
	"Rhyhorn": {
		"id":   uint(111),
		"form": uint(0),
		"type": "ground",
	},
	"Rhydon": {
		"id":   uint(112),
		"form": uint(0),
		"type": "ground",
	},
	"Chansey": {
		"id":   uint(113),
		"form": uint(0),
		"type": "normal",
	},
	"Tangela": {
		"id":   uint(114),
		"form": uint(0),
		"type": "grass",
	},
	"Kangaskhan": {
		"id":   uint(115),
		"form": uint(0),
		"type": "normal",
	},
	"Kangaskhan-Mega": {
		"id":   uint(115),
		"form": uint(1),
		"type": "normal",
	},
	"Horsea": {
		"id":   uint(116),
		"form": uint(0),
		"type": "water",
	},
	"Seadra": {
		"id":   uint(117),
		"form": uint(0),
		"type": "water",
	},
	"Goldeen": {
		"id":   uint(118),
		"form": uint(0),
		"type": "water",
	},
	"Seaking": {
		"id":   uint(119),
		"form": uint(0),
		"type": "water",
	},
	"Staryu": {
		"id":   uint(120),
		"form": uint(0),
		"type": "water",
	},
	"Starmie": {
		"id":   uint(121),
		"form": uint(0),
		"type": "water",
	},
	"Mr. Mime": {
		"id":   uint(122),
		"form": uint(0),
		"type": "psychic",
	},
	"Mr. Mime-Galar": {
		"id":   uint(122),
		"form": uint(1),
		"type": "ice",
	},
	"Scyther": {
		"id":   uint(123),
		"form": uint(0),
		"type": "bug",
	},
	"Jynx": {
		"id":   uint(124),
		"form": uint(0),
		"type": "ice",
	},
	"Electabuzz": {
		"id":   uint(125),
		"form": uint(0),
		"type": "electric",
	},
	"Magmar": {
		"id":   uint(126),
		"form": uint(0),
		"type": "fire",
	},
	"Pinsir": {
		"id":   uint(127),
		"form": uint(0),
		"type": "bug",
	},
	"Pinsir-Mega": {
		"id":   uint(127),
		"form": uint(1),
		"type": "bug",
	},
	"Tauros": {
		"id":   uint(128),
		"form": uint(0),
		"type": "normal",
	},
	"Magikarp": {
		"id":   uint(129),
		"form": uint(0),
		"type": "water",
	},
	"Gyarados": {
		"id":   uint(130),
		"form": uint(0),
		"type": "water",
	},
	"Gyarados-Mega": {
		"id":   uint(130),
		"form": uint(1),
		"type": "water",
	},
	"Lapras": {
		"id":   uint(131),
		"form": uint(0),
		"type": "water",
	},
	"Lapras-Gmax": {
		"id":   uint(131),
		"form": uint(1),
		"type": "water",
	},
	"Ditto": {
		"id":   uint(132),
		"form": uint(0),
		"type": "normal",
	},
	"Eevee": {
		"id":   uint(133),
		"form": uint(0),
		"type": "normal",
	},
	"Eevee-Gmax": {
		"id":   uint(133),
		"form": uint(1),
		"type": "normal",
	},
	"Vaporeon": {
		"id":   uint(134),
		"form": uint(0),
		"type": "water",
	},
	"Jolteon": {
		"id":   uint(135),
		"form": uint(0),
		"type": "electric",
	},
	"Flareon": {
		"id":   uint(136),
		"form": uint(0),
		"type": "fire",
	},
	"Porygon": {
		"id":   uint(137),
		"form": uint(0),
		"type": "normal",
	},
	"Omanyte": {
		"id":   uint(138),
		"form": uint(0),
		"type": "rock",
	},
	"Omastar": {
		"id":   uint(139),
		"form": uint(0),
		"type": "rock",
	},
	"Kabuto": {
		"id":   uint(140),
		"form": uint(0),
		"type": "rock",
	},
	"Kabutops": {
		"id":   uint(141),
		"form": uint(0),
		"type": "rock",
	},
	"Aerodactyl": {
		"id":   uint(142),
		"form": uint(0),
		"type": "rock",
	},
	"Aerodactyl-Mega": {
		"id":   uint(142),
		"form": uint(1),
		"type": "rock",
	},
	"Snorlax": {
		"id":   uint(143),
		"form": uint(0),
		"type": "normal",
	},
	"Snorlax-Gmax": {
		"id":   uint(143),
		"form": uint(1),
		"type": "normal",
	},
	"Articuno": {
		"id":   uint(144),
		"form": uint(0),
		"type": "ice",
	},
	"Articuno-Galar": {
		"id":   uint(144),
		"form": uint(1),
		"type": "psychic",
	},
	"Zapdos": {
		"id":   uint(145),
		"form": uint(0),
		"type": "electric",
	},
	"Zapdos-Galar": {
		"id":   uint(145),
		"form": uint(1),
		"type": "fighting",
	},
	"Moltres": {
		"id":   uint(146),
		"form": uint(0),
		"type": "fire",
	},
	"Moltres-Galar": {
		"id":   uint(146),
		"form": uint(1),
		"type": "fire",
	},
	"Dratini": {
		"id":   uint(147),
		"form": uint(0),
		"type": "dragon",
	},
	"Dragonair": {
		"id":   uint(148),
		"form": uint(0),
		"type": "dragon",
	},
	"Dragonite": {
		"id":   uint(149),
		"form": uint(0),
		"type": "dragon",
	},
	"Mewtwo": {
		"id":   uint(150),
		"form": uint(0),
		"type": "psychic",
	},
	"Mewtwo-Mega-X": {
		"id":   uint(150),
		"form": uint(1),
		"type": "psychic",
	},
	"Mewtwo-Mega-Y": {
		"id":   uint(150),
		"form": uint(2),
		"type": "psychic",
	},
	"Mew": {
		"id":   uint(151),
		"form": uint(0),
		"type": "psychic",
	},
	"Chikorita": {
		"id":   uint(152),
		"form": uint(0),
		"type": "grass",
	},
	"Bayleef": {
		"id":   uint(153),
		"form": uint(0),
		"type": "grass",
	},
	"Meganium": {
		"id":   uint(154),
		"form": uint(0),
		"type": "grass",
	},
	"Cyndaquil": {
		"id":   uint(155),
		"form": uint(0),
		"type": "fire",
	},
	"Quilava": {
		"id":   uint(156),
		"form": uint(0),
		"type": "fire",
	},
	"Typhlosion": {
		"id":   uint(157),
		"form": uint(0),
		"type": "fire",
	},
	"Totodile": {
		"id":   uint(158),
		"form": uint(0),
		"type": "water",
	},
	"Croconaw": {
		"id":   uint(159),
		"form": uint(0),
		"type": "water",
	},
	"Feraligatr": {
		"id":   uint(160),
		"form": uint(0),
		"type": "water",
	},
	"Sentret": {
		"id":   uint(161),
		"form": uint(0),
		"type": "normal",
	},
	"Furret": {
		"id":   uint(162),
		"form": uint(0),
		"type": "normal",
	},
	"Hoothoot": {
		"id":   uint(163),
		"form": uint(0),
		"type": "normal",
	},
	"Noctowl": {
		"id":   uint(164),
		"form": uint(0),
		"type": "normal",
	},
	"Ledyba": {
		"id":   uint(165),
		"form": uint(0),
		"type": "bug",
	},
	"Ledian": {
		"id":   uint(166),
		"form": uint(0),
		"type": "bug",
	},
	"Spinarak": {
		"id":   uint(167),
		"form": uint(0),
		"type": "bug",
	},
	"Ariados": {
		"id":   uint(168),
		"form": uint(0),
		"type": "bug",
	},
	"Crobat": {
		"id":   uint(169),
		"form": uint(0),
		"type": "poison",
	},
	"Chinchou": {
		"id":   uint(170),
		"form": uint(0),
		"type": "water",
	},
	"Lanturn": {
		"id":   uint(171),
		"form": uint(0),
		"type": "water",
	},
	"Pichu": {
		"id":   uint(172),
		"form": uint(0),
		"type": "electric",
	},
	"Cleffa": {
		"id":   uint(173),
		"form": uint(0),
		"type": "fairy",
	},
	"Igglybuff": {
		"id":   uint(174),
		"form": uint(0),
		"type": "normal",
	},
	"Togepi": {
		"id":   uint(175),
		"form": uint(0),
		"type": "fairy",
	},
	"Togetic": {
		"id":   uint(176),
		"form": uint(0),
		"type": "fairy",
	},
	"Natu": {
		"id":   uint(177),
		"form": uint(0),
		"type": "psychic",
	},
	"Xatu": {
		"id":   uint(178),
		"form": uint(0),
		"type": "psychic",
	},
	"Mareep": {
		"id":   uint(179),
		"form": uint(0),
		"type": "electric",
	},
	"Flaaffy": {
		"id":   uint(180),
		"form": uint(0),
		"type": "electric",
	},
	"Ampharos": {
		"id":   uint(181),
		"form": uint(0),
		"type": "electric",
	},
	"Ampharos-Mega": {
		"id":   uint(181),
		"form": uint(1),
		"type": "electric",
	},
	"Bellossom": {
		"id":   uint(182),
		"form": uint(0),
		"type": "grass",
	},
	"Marill": {
		"id":   uint(183),
		"form": uint(0),
		"type": "water",
	},
	"Azumarill": {
		"id":   uint(184),
		"form": uint(0),
		"type": "water",
	},
	"Sudowoodo": {
		"id":   uint(185),
		"form": uint(0),
		"type": "rock",
	},
	"Politoed": {
		"id":   uint(186),
		"form": uint(0),
		"type": "water",
	},
	"Hoppip": {
		"id":   uint(187),
		"form": uint(0),
		"type": "grass",
	},
	"Skiploom": {
		"id":   uint(188),
		"form": uint(0),
		"type": "grass",
	},
	"Jumpluff": {
		"id":   uint(189),
		"form": uint(0),
		"type": "grass",
	},
	"Aipom": {
		"id":   uint(190),
		"form": uint(0),
		"type": "normal",
	},
	"Sunkern": {
		"id":   uint(191),
		"form": uint(0),
		"type": "grass",
	},
	"Sunflora": {
		"id":   uint(192),
		"form": uint(0),
		"type": "grass",
	},
	"Yanma": {
		"id":   uint(193),
		"form": uint(0),
		"type": "bug",
	},
	"Wooper": {
		"id":   uint(194),
		"form": uint(0),
		"type": "water",
	},
	"Quagsire": {
		"id":   uint(195),
		"form": uint(0),
		"type": "water",
	},
	"Espeon": {
		"id":   uint(196),
		"form": uint(0),
		"type": "psychic",
	},
	"Umbreon": {
		"id":   uint(197),
		"form": uint(0),
		"type": "dark",
	},
	"Murkrow": {
		"id":   uint(198),
		"form": uint(0),
		"type": "dark",
	},
	"Slowking": {
		"id":   uint(199),
		"form": uint(0),
		"type": "water",
	},
	"Slowking-Galar": {
		"id":   uint(199),
		"form": uint(1),
		"type": "poison",
	},
	"Misdreavus": {
		"id":   uint(200),
		"form": uint(0),
		"type": "ghost",
	},
	"Unown": {
		"id":   uint(201),
		"form": uint(0),
		"type": "psychic",
	},
	"Wobbuffet": {
		"id":   uint(202),
		"form": uint(0),
		"type": "psychic",
	},
	"Girafarig": {
		"id":   uint(203),
		"form": uint(0),
		"type": "normal",
	},
	"Pineco": {
		"id":   uint(204),
		"form": uint(0),
		"type": "bug",
	},
	"Forretress": {
		"id":   uint(205),
		"form": uint(0),
		"type": "bug",
	},
	"Dunsparce": {
		"id":   uint(206),
		"form": uint(0),
		"type": "normal",
	},
	"Gligar": {
		"id":   uint(207),
		"form": uint(0),
		"type": "ground",
	},
	"Steelix": {
		"id":   uint(208),
		"form": uint(0),
		"type": "steel",
	},
	"Steelix-Mega": {
		"id":   uint(208),
		"form": uint(1),
		"type": "steel",
	},
	"Snubbull": {
		"id":   uint(209),
		"form": uint(0),
		"type": "fairy",
	},
	"Granbull": {
		"id":   uint(210),
		"form": uint(0),
		"type": "fairy",
	},
	"Qwilfish": {
		"id":   uint(211),
		"form": uint(0),
		"type": "water",
	},
	"Scizor": {
		"id":   uint(212),
		"form": uint(0),
		"type": "bug",
	},
	"Scizor-Mega": {
		"id":   uint(212),
		"form": uint(1),
		"type": "bug",
	},
	"Shuckle": {
		"id":   uint(213),
		"form": uint(0),
		"type": "bug",
	},
	"Heracross": {
		"id":   uint(214),
		"form": uint(0),
		"type": "bug",
	},
	"Heracross-Mega": {
		"id":   uint(214),
		"form": uint(1),
		"type": "bug",
	},
	"Sneasel": {
		"id":   uint(215),
		"form": uint(0),
		"type": "dark",
	},
	"Teddiursa": {
		"id":   uint(216),
		"form": uint(0),
		"type": "normal",
	},
	"Ursaring": {
		"id":   uint(217),
		"form": uint(0),
		"type": "normal",
	},
	"Slugma": {
		"id":   uint(218),
		"form": uint(0),
		"type": "fire",
	},
	"Magcargo": {
		"id":   uint(219),
		"form": uint(0),
		"type": "fire",
	},
	"Swinub": {
		"id":   uint(220),
		"form": uint(0),
		"type": "ice",
	},
	"Piloswine": {
		"id":   uint(221),
		"form": uint(0),
		"type": "ice",
	},
	"Corsola": {
		"id":   uint(222),
		"form": uint(0),
		"type": "water",
	},
	"Corsola-Galar": {
		"id":   uint(222),
		"form": uint(1),
		"type": "ghost",
	},
	"Remoraid": {
		"id":   uint(223),
		"form": uint(0),
		"type": "water",
	},
	"Octillery": {
		"id":   uint(224),
		"form": uint(0),
		"type": "water",
	},
	"Delibird": {
		"id":   uint(225),
		"form": uint(0),
		"type": "ice",
	},
	"Mantine": {
		"id":   uint(226),
		"form": uint(0),
		"type": "water",
	},
	"Skarmory": {
		"id":   uint(227),
		"form": uint(0),
		"type": "steel",
	},
	"Houndour": {
		"id":   uint(228),
		"form": uint(0),
		"type": "dark",
	},
	"Houndoom": {
		"id":   uint(229),
		"form": uint(0),
		"type": "dark",
	},
	"Houndoom-Mega": {
		"id":   uint(229),
		"form": uint(1),
		"type": "dark",
	},
	"Kingdra": {
		"id":   uint(230),
		"form": uint(0),
		"type": "water",
	},
	"Phanpy": {
		"id":   uint(231),
		"form": uint(0),
		"type": "ground",
	},
	"Donphan": {
		"id":   uint(232),
		"form": uint(0),
		"type": "ground",
	},
	"Porygon2": {
		"id":   uint(233),
		"form": uint(0),
		"type": "normal",
	},
	"Stantler": {
		"id":   uint(234),
		"form": uint(0),
		"type": "normal",
	},
	"Smeargle": {
		"id":   uint(235),
		"form": uint(0),
		"type": "normal",
	},
	"Tyrogue": {
		"id":   uint(236),
		"form": uint(0),
		"type": "fighting",
	},
	"Hitmontop": {
		"id":   uint(237),
		"form": uint(0),
		"type": "fighting",
	},
	"Smoochum": {
		"id":   uint(238),
		"form": uint(0),
		"type": "ice",
	},
	"Elekid": {
		"id":   uint(239),
		"form": uint(0),
		"type": "electric",
	},
	"Magby": {
		"id":   uint(240),
		"form": uint(0),
		"type": "fire",
	},
	"Miltank": {
		"id":   uint(241),
		"form": uint(0),
		"type": "normal",
	},
	"Blissey": {
		"id":   uint(242),
		"form": uint(0),
		"type": "normal",
	},
	"Raikou": {
		"id":   uint(243),
		"form": uint(0),
		"type": "electric",
	},
	"Entei": {
		"id":   uint(244),
		"form": uint(0),
		"type": "fire",
	},
	"Suicune": {
		"id":   uint(245),
		"form": uint(0),
		"type": "water",
	},
	"Larvitar": {
		"id":   uint(246),
		"form": uint(0),
		"type": "rock",
	},
	"Pupitar": {
		"id":   uint(247),
		"form": uint(0),
		"type": "rock",
	},
	"Tyranitar": {
		"id":   uint(248),
		"form": uint(0),
		"type": "rock",
	},
	"Tyranitar-Mega": {
		"id":   uint(248),
		"form": uint(1),
		"type": "rock",
	},
	"Lugia": {
		"id":   uint(249),
		"form": uint(0),
		"type": "psychic",
	},
	"Ho-Oh": {
		"id":   uint(250),
		"form": uint(0),
		"type": "fire",
	},
	"Celebi": {
		"id":   uint(251),
		"form": uint(0),
		"type": "psychic",
	},
	"Treecko": {
		"id":   uint(252),
		"form": uint(0),
		"type": "grass",
	},
	"Grovyle": {
		"id":   uint(253),
		"form": uint(0),
		"type": "grass",
	},
	"Sceptile": {
		"id":   uint(254),
		"form": uint(0),
		"type": "grass",
	},
	"Sceptile-Mega": {
		"id":   uint(254),
		"form": uint(1),
		"type": "grass",
	},
	"Torchic": {
		"id":   uint(255),
		"form": uint(0),
		"type": "fire",
	},
	"Combusken": {
		"id":   uint(256),
		"form": uint(0),
		"type": "fire",
	},
	"Blaziken": {
		"id":   uint(257),
		"form": uint(0),
		"type": "fire",
	},
	"Blaziken-Mega": {
		"id":   uint(257),
		"form": uint(1),
		"type": "fire",
	},
	"Mudkip": {
		"id":   uint(258),
		"form": uint(0),
		"type": "water",
	},
	"Marshtomp": {
		"id":   uint(259),
		"form": uint(0),
		"type": "water",
	},
	"Swampert": {
		"id":   uint(260),
		"form": uint(0),
		"type": "water",
	},
	"Swampert-Mega": {
		"id":   uint(260),
		"form": uint(1),
		"type": "water",
	},
	"Poochyena": {
		"id":   uint(261),
		"form": uint(0),
		"type": "dark",
	},
	"Mightyena": {
		"id":   uint(262),
		"form": uint(0),
		"type": "dark",
	},
	"Zigzagoon": {
		"id":   uint(263),
		"form": uint(0),
		"type": "normal",
	},
	"Zigzagoon-Galar": {
		"id":   uint(263),
		"form": uint(1),
		"type": "dark",
	},
	"Linoone": {
		"id":   uint(264),
		"form": uint(0),
		"type": "normal",
	},
	"Linoone-Galar": {
		"id":   uint(264),
		"form": uint(1),
		"type": "dark",
	},
	"Wurmple": {
		"id":   uint(265),
		"form": uint(0),
		"type": "bug",
	},
	"Silcoon": {
		"id":   uint(266),
		"form": uint(0),
		"type": "bug",
	},
	"Beautifly": {
		"id":   uint(267),
		"form": uint(0),
		"type": "bug",
	},
	"Cascoon": {
		"id":   uint(268),
		"form": uint(0),
		"type": "bug",
	},
	"Dustox": {
		"id":   uint(269),
		"form": uint(0),
		"type": "bug",
	},
	"Lotad": {
		"id":   uint(270),
		"form": uint(0),
		"type": "water",
	},
	"Lombre": {
		"id":   uint(271),
		"form": uint(0),
		"type": "water",
	},
	"Ludicolo": {
		"id":   uint(272),
		"form": uint(0),
		"type": "water",
	},
	"Seedot": {
		"id":   uint(273),
		"form": uint(0),
		"type": "grass",
	},
	"Nuzleaf": {
		"id":   uint(274),
		"form": uint(0),
		"type": "grass",
	},
	"Shiftry": {
		"id":   uint(275),
		"form": uint(0),
		"type": "grass",
	},
	"Taillow": {
		"id":   uint(276),
		"form": uint(0),
		"type": "normal",
	},
	"Swellow": {
		"id":   uint(277),
		"form": uint(0),
		"type": "normal",
	},
	"Wingull": {
		"id":   uint(278),
		"form": uint(0),
		"type": "water",
	},
	"Pelipper": {
		"id":   uint(279),
		"form": uint(0),
		"type": "water",
	},
	"Ralts": {
		"id":   uint(280),
		"form": uint(0),
		"type": "psychic",
	},
	"Kirlia": {
		"id":   uint(281),
		"form": uint(0),
		"type": "psychic",
	},
	"Gardevoir": {
		"id":   uint(282),
		"form": uint(0),
		"type": "psychic",
	},
	"Gardevoir-Mega": {
		"id":   uint(282),
		"form": uint(1),
		"type": "psychic",
	},
	"Surskit": {
		"id":   uint(283),
		"form": uint(0),
		"type": "bug",
	},
	"Masquerain": {
		"id":   uint(284),
		"form": uint(0),
		"type": "bug",
	},
	"Shroomish": {
		"id":   uint(285),
		"form": uint(0),
		"type": "grass",
	},
	"Breloom": {
		"id":   uint(286),
		"form": uint(0),
		"type": "grass",
	},
	"Slakoth": {
		"id":   uint(287),
		"form": uint(0),
		"type": "normal",
	},
	"Vigoroth": {
		"id":   uint(288),
		"form": uint(0),
		"type": "normal",
	},
	"Slaking": {
		"id":   uint(289),
		"form": uint(0),
		"type": "normal",
	},
	"Nincada": {
		"id":   uint(290),
		"form": uint(0),
		"type": "bug",
	},
	"Ninjask": {
		"id":   uint(291),
		"form": uint(0),
		"type": "bug",
	},
	"Shedinja": {
		"id":   uint(292),
		"form": uint(0),
		"type": "bug",
	},
	"Whismur": {
		"id":   uint(293),
		"form": uint(0),
		"type": "normal",
	},
	"Loudred": {
		"id":   uint(294),
		"form": uint(0),
		"type": "normal",
	},
	"Exploud": {
		"id":   uint(295),
		"form": uint(0),
		"type": "normal",
	},
	"Makuhita": {
		"id":   uint(296),
		"form": uint(0),
		"type": "fighting",
	},
	"Hariyama": {
		"id":   uint(297),
		"form": uint(0),
		"type": "fighting",
	},
	"Azurill": {
		"id":   uint(298),
		"form": uint(0),
		"type": "normal",
	},
	"Nosepass": {
		"id":   uint(299),
		"form": uint(0),
		"type": "rock",
	},
	"Skitty": {
		"id":   uint(300),
		"form": uint(0),
		"type": "normal",
	},
	"Delcatty": {
		"id":   uint(301),
		"form": uint(0),
		"type": "normal",
	},
	"Sableye": {
		"id":   uint(302),
		"form": uint(0),
		"type": "dark",
	},
	"Sableye-Mega": {
		"id":   uint(302),
		"form": uint(1),
		"type": "dark",
	},
	"Mawile": {
		"id":   uint(303),
		"form": uint(0),
		"type": "steel",
	},
	"Mawile-Mega": {
		"id":   uint(303),
		"form": uint(1),
		"type": "steel",
	},
	"Aron": {
		"id":   uint(304),
		"form": uint(0),
		"type": "steel",
	},
	"Lairon": {
		"id":   uint(305),
		"form": uint(0),
		"type": "steel",
	},
	"Aggron": {
		"id":   uint(306),
		"form": uint(0),
		"type": "steel",
	},
	"Aggron-Mega": {
		"id":   uint(306),
		"form": uint(1),
		"type": "steel",
	},
	"Meditite": {
		"id":   uint(307),
		"form": uint(0),
		"type": "fighting",
	},
	"Medicham": {
		"id":   uint(308),
		"form": uint(0),
		"type": "fighting",
	},
	"Medicham-Mega": {
		"id":   uint(308),
		"form": uint(1),
		"type": "fighting",
	},
	"Electrike": {
		"id":   uint(309),
		"form": uint(0),
		"type": "electric",
	},
	"Manectric": {
		"id":   uint(310),
		"form": uint(0),
		"type": "electric",
	},
	"Manectric-Mega": {
		"id":   uint(310),
		"form": uint(1),
		"type": "electric",
	},
	"Plusle": {
		"id":   uint(311),
		"form": uint(0),
		"type": "electric",
	},
	"Minun": {
		"id":   uint(312),
		"form": uint(0),
		"type": "electric",
	},
	"Volbeat": {
		"id":   uint(313),
		"form": uint(0),
		"type": "bug",
	},
	"Illumise": {
		"id":   uint(314),
		"form": uint(0),
		"type": "bug",
	},
	"Roselia": {
		"id":   uint(315),
		"form": uint(0),
		"type": "grass",
	},
	"Gulpin": {
		"id":   uint(316),
		"form": uint(0),
		"type": "poison",
	},
	"Swalot": {
		"id":   uint(317),
		"form": uint(0),
		"type": "poison",
	},
	"Carvanha": {
		"id":   uint(318),
		"form": uint(0),
		"type": "water",
	},
	"Sharpedo": {
		"id":   uint(319),
		"form": uint(0),
		"type": "water",
	},
	"Sharpedo-Mega": {
		"id":   uint(319),
		"form": uint(1),
		"type": "water",
	},
	"Wailmer": {
		"id":   uint(320),
		"form": uint(0),
		"type": "water",
	},
	"Wailord": {
		"id":   uint(321),
		"form": uint(0),
		"type": "water",
	},
	"Numel": {
		"id":   uint(322),
		"form": uint(0),
		"type": "fire",
	},
	"Camerupt": {
		"id":   uint(323),
		"form": uint(0),
		"type": "fire",
	},
	"Camerupt-Mega": {
		"id":   uint(323),
		"form": uint(1),
		"type": "fire",
	},
	"Torkoal": {
		"id":   uint(324),
		"form": uint(0),
		"type": "fire",
	},
	"Spoink": {
		"id":   uint(325),
		"form": uint(0),
		"type": "psychic",
	},
	"Grumpig": {
		"id":   uint(326),
		"form": uint(0),
		"type": "psychic",
	},
	"Spinda": {
		"id":   uint(327),
		"form": uint(0),
		"type": "normal",
	},
	"Trapinch": {
		"id":   uint(328),
		"form": uint(0),
		"type": "ground",
	},
	"Vibrava": {
		"id":   uint(329),
		"form": uint(0),
		"type": "ground",
	},
	"Flygon": {
		"id":   uint(330),
		"form": uint(0),
		"type": "ground",
	},
	"Cacnea": {
		"id":   uint(331),
		"form": uint(0),
		"type": "grass",
	},
	"Cacturne": {
		"id":   uint(332),
		"form": uint(0),
		"type": "grass",
	},
	"Swablu": {
		"id":   uint(333),
		"form": uint(0),
		"type": "normal",
	},
	"Altaria": {
		"id":   uint(334),
		"form": uint(0),
		"type": "dragon",
	},
	"Altaria-Mega": {
		"id":   uint(334),
		"form": uint(1),
		"type": "dragon",
	},
	"Zangoose": {
		"id":   uint(335),
		"form": uint(0),
		"type": "normal",
	},
	"Seviper": {
		"id":   uint(336),
		"form": uint(0),
		"type": "poison",
	},
	"Lunatone": {
		"id":   uint(337),
		"form": uint(0),
		"type": "rock",
	},
	"Solrock": {
		"id":   uint(338),
		"form": uint(0),
		"type": "rock",
	},
	"Barboach": {
		"id":   uint(339),
		"form": uint(0),
		"type": "water",
	},
	"Whiscash": {
		"id":   uint(340),
		"form": uint(0),
		"type": "water",
	},
	"Corphish": {
		"id":   uint(341),
		"form": uint(0),
		"type": "water",
	},
	"Crawdaunt": {
		"id":   uint(342),
		"form": uint(0),
		"type": "water",
	},
	"Baltoy": {
		"id":   uint(343),
		"form": uint(0),
		"type": "ground",
	},
	"Claydol": {
		"id":   uint(344),
		"form": uint(0),
		"type": "ground",
	},
	"Lileep": {
		"id":   uint(345),
		"form": uint(0),
		"type": "rock",
	},
	"Cradily": {
		"id":   uint(346),
		"form": uint(0),
		"type": "rock",
	},
	"Anorith": {
		"id":   uint(347),
		"form": uint(0),
		"type": "rock",
	},
	"Armaldo": {
		"id":   uint(348),
		"form": uint(0),
		"type": "rock",
	},
	"Feebas": {
		"id":   uint(349),
		"form": uint(0),
		"type": "water",
	},
	"Milotic": {
		"id":   uint(350),
		"form": uint(0),
		"type": "water",
	},
	"Castform": {
		"id":   uint(351),
		"form": uint(0),
		"type": "normal",
	},
	"Castform-Sunny": {
		"id":   uint(351),
		"form": uint(1),
		"type": "fire",
	},
	"Castform-Rainy": {
		"id":   uint(351),
		"form": uint(2),
		"type": "water",
	},
	"Castform-Snowy": {
		"id":   uint(351),
		"form": uint(3),
		"type": "ice",
	},
	"Kecleon": {
		"id":   uint(352),
		"form": uint(0),
		"type": "normal",
	},
	"Shuppet": {
		"id":   uint(353),
		"form": uint(0),
		"type": "ghost",
	},
	"Banette": {
		"id":   uint(354),
		"form": uint(0),
		"type": "ghost",
	},
	"Banette-Mega": {
		"id":   uint(354),
		"form": uint(1),
		"type": "ghost",
	},
	"Duskull": {
		"id":   uint(355),
		"form": uint(0),
		"type": "ghost",
	},
	"Dusclops": {
		"id":   uint(356),
		"form": uint(0),
		"type": "ghost",
	},
	"Tropius": {
		"id":   uint(357),
		"form": uint(0),
		"type": "grass",
	},
	"Chimecho": {
		"id":   uint(358),
		"form": uint(0),
		"type": "psychic",
	},
	"Absol": {
		"id":   uint(359),
		"form": uint(0),
		"type": "dark",
	},
	"Absol-Mega": {
		"id":   uint(359),
		"form": uint(1),
		"type": "dark",
	},
	"Wynaut": {
		"id":   uint(360),
		"form": uint(0),
		"type": "psychic",
	},
	"Snorunt": {
		"id":   uint(361),
		"form": uint(0),
		"type": "ice",
	},
	"Glalie": {
		"id":   uint(362),
		"form": uint(0),
		"type": "ice",
	},
	"Glalie-Mega": {
		"id":   uint(362),
		"form": uint(1),
		"type": "ice",
	},
	"Spheal": {
		"id":   uint(363),
		"form": uint(0),
		"type": "ice",
	},
	"Sealeo": {
		"id":   uint(364),
		"form": uint(0),
		"type": "ice",
	},
	"Walrein": {
		"id":   uint(365),
		"form": uint(0),
		"type": "ice",
	},
	"Clamperl": {
		"id":   uint(366),
		"form": uint(0),
		"type": "water",
	},
	"Huntail": {
		"id":   uint(367),
		"form": uint(0),
		"type": "water",
	},
	"Gorebyss": {
		"id":   uint(368),
		"form": uint(0),
		"type": "water",
	},
	"Relicanth": {
		"id":   uint(369),
		"form": uint(0),
		"type": "water",
	},
	"Luvdisc": {
		"id":   uint(370),
		"form": uint(0),
		"type": "water",
	},
	"Bagon": {
		"id":   uint(371),
		"form": uint(0),
		"type": "dragon",
	},
	"Shelgon": {
		"id":   uint(372),
		"form": uint(0),
		"type": "dragon",
	},
	"Salamence": {
		"id":   uint(373),
		"form": uint(0),
		"type": "dragon",
	},
	"Salamence-Mega": {
		"id":   uint(373),
		"form": uint(1),
		"type": "dragon",
	},
	"Beldum": {
		"id":   uint(374),
		"form": uint(0),
		"type": "steel",
	},
	"Metang": {
		"id":   uint(375),
		"form": uint(0),
		"type": "steel",
	},
	"Metagross": {
		"id":   uint(376),
		"form": uint(0),
		"type": "steel",
	},
	"Metagross-Mega": {
		"id":   uint(376),
		"form": uint(1),
		"type": "steel",
	},
	"Regirock": {
		"id":   uint(377),
		"form": uint(0),
		"type": "rock",
	},
	"Regice": {
		"id":   uint(378),
		"form": uint(0),
		"type": "ice",
	},
	"Registeel": {
		"id":   uint(379),
		"form": uint(0),
		"type": "steel",
	},
	"Latias": {
		"id":   uint(380),
		"form": uint(0),
		"type": "dragon",
	},
	"Latias-Mega": {
		"id":   uint(380),
		"form": uint(1),
		"type": "dragon",
	},
	"Latios": {
		"id":   uint(381),
		"form": uint(0),
		"type": "dragon",
	},
	"Latios-Mega": {
		"id":   uint(381),
		"form": uint(1),
		"type": "dragon",
	},
	"Kyogre": {
		"id":   uint(382),
		"form": uint(0),
		"type": "water",
	},
	"Kyogre-Primal": {
		"id":   uint(382),
		"form": uint(1),
		"type": "water",
	},
	"Groudon": {
		"id":   uint(383),
		"form": uint(0),
		"type": "ground",
	},
	"Groudon-Primal": {
		"id":   uint(383),
		"form": uint(1),
		"type": "ground",
	},
	"Rayquaza": {
		"id":   uint(384),
		"form": uint(0),
		"type": "dragon",
	},
	"Rayquaza-Mega": {
		"id":   uint(384),
		"form": uint(1),
		"type": "dragon",
	},
	"Jirachi": {
		"id":   uint(385),
		"form": uint(0),
		"type": "steel",
	},
	"Deoxys": {
		"id":   uint(386),
		"form": uint(0),
		"type": "psychic",
	},
	"Deoxys-Attack": {
		"id":   uint(386),
		"form": uint(1),
		"type": "psychic",
	},
	"Deoxys-Defense": {
		"id":   uint(386),
		"form": uint(2),
		"type": "psychic",
	},
	"Deoxys-Speed": {
		"id":   uint(386),
		"form": uint(3),
		"type": "psychic",
	},
	"Turtwig": {
		"id":   uint(387),
		"form": uint(0),
		"type": "grass",
	},
	"Grotle": {
		"id":   uint(388),
		"form": uint(0),
		"type": "grass",
	},
	"Torterra": {
		"id":   uint(389),
		"form": uint(0),
		"type": "grass",
	},
	"Chimchar": {
		"id":   uint(390),
		"form": uint(0),
		"type": "fire",
	},
	"Monferno": {
		"id":   uint(391),
		"form": uint(0),
		"type": "fire",
	},
	"Infernape": {
		"id":   uint(392),
		"form": uint(0),
		"type": "fire",
	},
	"Piplup": {
		"id":   uint(393),
		"form": uint(0),
		"type": "water",
	},
	"Prinplup": {
		"id":   uint(394),
		"form": uint(0),
		"type": "water",
	},
	"Empoleon": {
		"id":   uint(395),
		"form": uint(0),
		"type": "water",
	},
	"Starly": {
		"id":   uint(396),
		"form": uint(0),
		"type": "normal",
	},
	"Staravia": {
		"id":   uint(397),
		"form": uint(0),
		"type": "normal",
	},
	"Staraptor": {
		"id":   uint(398),
		"form": uint(0),
		"type": "normal",
	},
	"Bidoof": {
		"id":   uint(399),
		"form": uint(0),
		"type": "normal",
	},
	"Bibarel": {
		"id":   uint(400),
		"form": uint(0),
		"type": "normal",
	},
	"Kricketot": {
		"id":   uint(401),
		"form": uint(0),
		"type": "bug",
	},
	"Kricketune": {
		"id":   uint(402),
		"form": uint(0),
		"type": "bug",
	},
	"Shinx": {
		"id":   uint(403),
		"form": uint(0),
		"type": "electric",
	},
	"Luxio": {
		"id":   uint(404),
		"form": uint(0),
		"type": "electric",
	},
	"Luxray": {
		"id":   uint(405),
		"form": uint(0),
		"type": "electric",
	},
	"Budew": {
		"id":   uint(406),
		"form": uint(0),
		"type": "grass",
	},
	"Roserade": {
		"id":   uint(407),
		"form": uint(0),
		"type": "grass",
	},
	"Cranidos": {
		"id":   uint(408),
		"form": uint(0),
		"type": "rock",
	},
	"Rampardos": {
		"id":   uint(409),
		"form": uint(0),
		"type": "rock",
	},
	"Shieldon": {
		"id":   uint(410),
		"form": uint(0),
		"type": "rock",
	},
	"Bastiodon": {
		"id":   uint(411),
		"form": uint(0),
		"type": "rock",
	},
	"Burmy": {
		"id":   uint(412),
		"form": uint(0),
		"type": "bug",
	},
	"Burmy-Sandy": {
		"id":   uint(412),
		"form": uint(1),
		"type": "bug",
	},
	"Burmy-Trash": {
		"id":   uint(412),
		"form": uint(2),
		"type": "bug",
	},
	"Wormadam": {
		"id":   uint(413),
		"form": uint(0),
		"type": "bug",
	},
	"Wormadam-Sandy": {
		"id":   uint(413),
		"form": uint(1),
		"type": "bug",
	},
	"Wormadam-Trash": {
		"id":   uint(413),
		"form": uint(2),
		"type": "bug",
	},
	"Mothim": {
		"id":   uint(414),
		"form": uint(0),
		"type": "bug",
	},
	"Combee": {
		"id":   uint(415),
		"form": uint(0),
		"type": "bug",
	},
	"Vespiquen": {
		"id":   uint(416),
		"form": uint(0),
		"type": "bug",
	},
	"Pachirisu": {
		"id":   uint(417),
		"form": uint(0),
		"type": "electric",
	},
	"Buizel": {
		"id":   uint(418),
		"form": uint(0),
		"type": "water",
	},
	"Floatzel": {
		"id":   uint(419),
		"form": uint(0),
		"type": "water",
	},
	"Cherubi": {
		"id":   uint(420),
		"form": uint(0),
		"type": "grass",
	},
	"Cherrim": {
		"id":   uint(421),
		"form": uint(0),
		"type": "grass",
	},
	"Cherrim-Sunshine": {
		"id":   uint(421),
		"form": uint(1),
		"type": "grass",
	},
	"Shellos": {
		"id":   uint(422),
		"form": uint(0),
		"type": "water",
	},
	"Shellos-East": {
		"id":   uint(422),
		"form": uint(1),
		"type": "water",
	},
	"Gastrodon": {
		"id":   uint(423),
		"form": uint(0),
		"type": "water",
	},
	"Gastrodon-East": {
		"id":   uint(423),
		"form": uint(1),
		"type": "water",
	},
	"Ambipom": {
		"id":   uint(424),
		"form": uint(0),
		"type": "normal",
	},
	"Drifloon": {
		"id":   uint(425),
		"form": uint(0),
		"type": "ghost",
	},
	"Drifblim": {
		"id":   uint(426),
		"form": uint(0),
		"type": "ghost",
	},
	"Buneary": {
		"id":   uint(427),
		"form": uint(0),
		"type": "normal",
	},
	"Lopunny": {
		"id":   uint(428),
		"form": uint(0),
		"type": "normal",
	},
	"Lopunny-Mega": {
		"id":   uint(428),
		"form": uint(1),
		"type": "normal",
	},
	"Mismagius": {
		"id":   uint(429),
		"form": uint(0),
		"type": "ghost",
	},
	"Honchkrow": {
		"id":   uint(430),
		"form": uint(0),
		"type": "dark",
	},
	"Glameow": {
		"id":   uint(431),
		"form": uint(0),
		"type": "normal",
	},
	"Purugly": {
		"id":   uint(432),
		"form": uint(0),
		"type": "normal",
	},
	"Chingling": {
		"id":   uint(433),
		"form": uint(0),
		"type": "psychic",
	},
	"Stunky": {
		"id":   uint(434),
		"form": uint(0),
		"type": "poison",
	},
	"Skuntank": {
		"id":   uint(435),
		"form": uint(0),
		"type": "poison",
	},
	"Bronzor": {
		"id":   uint(436),
		"form": uint(0),
		"type": "steel",
	},
	"Bronzong": {
		"id":   uint(437),
		"form": uint(0),
		"type": "steel",
	},
	"Bonsly": {
		"id":   uint(438),
		"form": uint(0),
		"type": "rock",
	},
	"Mime Jr.": {
		"id":   uint(439),
		"form": uint(0),
		"type": "psychic",
	},
	"Happiny": {
		"id":   uint(440),
		"form": uint(0),
		"type": "normal",
	},
	"Chatot": {
		"id":   uint(441),
		"form": uint(0),
		"type": "normal",
	},
	"Spiritomb": {
		"id":   uint(442),
		"form": uint(0),
		"type": "ghost",
	},
	"Gible": {
		"id":   uint(443),
		"form": uint(0),
		"type": "dragon",
	},
	"Gabite": {
		"id":   uint(444),
		"form": uint(0),
		"type": "dragon",
	},
	"Garchomp": {
		"id":   uint(445),
		"form": uint(0),
		"type": "dragon",
	},
	"Garchomp-Mega": {
		"id":   uint(445),
		"form": uint(1),
		"type": "dragon",
	},
	"Munchlax": {
		"id":   uint(446),
		"form": uint(0),
		"type": "normal",
	},
	"Riolu": {
		"id":   uint(447),
		"form": uint(0),
		"type": "fighting",
	},
	"Lucario": {
		"id":   uint(448),
		"form": uint(0),
		"type": "fighting",
	},
	"Lucario-Mega": {
		"id":   uint(448),
		"form": uint(1),
		"type": "fighting",
	},
	"Hippopotas": {
		"id":   uint(449),
		"form": uint(0),
		"type": "ground",
	},
	"Hippowdon": {
		"id":   uint(450),
		"form": uint(0),
		"type": "ground",
	},
	"Skorupi": {
		"id":   uint(451),
		"form": uint(0),
		"type": "poison",
	},
	"Drapion": {
		"id":   uint(452),
		"form": uint(0),
		"type": "poison",
	},
	"Croagunk": {
		"id":   uint(453),
		"form": uint(0),
		"type": "poison",
	},
	"Toxicroak": {
		"id":   uint(454),
		"form": uint(0),
		"type": "poison",
	},
	"Carnivine": {
		"id":   uint(455),
		"form": uint(0),
		"type": "grass",
	},
	"Finneon": {
		"id":   uint(456),
		"form": uint(0),
		"type": "water",
	},
	"Lumineon": {
		"id":   uint(457),
		"form": uint(0),
		"type": "water",
	},
	"Mantyke": {
		"id":   uint(458),
		"form": uint(0),
		"type": "water",
	},
	"Snover": {
		"id":   uint(459),
		"form": uint(0),
		"type": "grass",
	},
	"Abomasnow": {
		"id":   uint(460),
		"form": uint(0),
		"type": "grass",
	},
	"Abomasnow-Mega": {
		"id":   uint(460),
		"form": uint(1),
		"type": "grass",
	},
	"Weavile": {
		"id":   uint(461),
		"form": uint(0),
		"type": "dark",
	},
	"Magnezone": {
		"id":   uint(462),
		"form": uint(0),
		"type": "electric",
	},
	"Lickilicky": {
		"id":   uint(463),
		"form": uint(0),
		"type": "normal",
	},
	"Rhyperior": {
		"id":   uint(464),
		"form": uint(0),
		"type": "ground",
	},
	"Tangrowth": {
		"id":   uint(465),
		"form": uint(0),
		"type": "grass",
	},
	"Electivire": {
		"id":   uint(466),
		"form": uint(0),
		"type": "electric",
	},
	"Magmortar": {
		"id":   uint(467),
		"form": uint(0),
		"type": "fire",
	},
	"Togekiss": {
		"id":   uint(468),
		"form": uint(0),
		"type": "fairy",
	},
	"Yanmega": {
		"id":   uint(469),
		"form": uint(0),
		"type": "bug",
	},
	"Leafeon": {
		"id":   uint(470),
		"form": uint(0),
		"type": "grass",
	},
	"Glaceon": {
		"id":   uint(471),
		"form": uint(0),
		"type": "ice",
	},
	"Gliscor": {
		"id":   uint(472),
		"form": uint(0),
		"type": "ground",
	},
	"Mamoswine": {
		"id":   uint(473),
		"form": uint(0),
		"type": "ice",
	},
	"Porygon-Z": {
		"id":   uint(474),
		"form": uint(0),
		"type": "normal",
	},
	"Gallade": {
		"id":   uint(475),
		"form": uint(0),
		"type": "psychic",
	},
	"Gallade-Mega": {
		"id":   uint(475),
		"form": uint(1),
		"type": "psychic",
	},
	"Probopass": {
		"id":   uint(476),
		"form": uint(0),
		"type": "rock",
	},
	"Dusknoir": {
		"id":   uint(477),
		"form": uint(0),
		"type": "ghost",
	},
	"Froslass": {
		"id":   uint(478),
		"form": uint(0),
		"type": "ice",
	},
	"Rotom": {
		"id":   uint(479),
		"form": uint(0),
		"type": "electric",
	},
	"Rotom-Heat": {
		"id":   uint(479),
		"form": uint(1),
		"type": "electric",
	},
	"Rotom-Wash": {
		"id":   uint(479),
		"form": uint(2),
		"type": "electric",
	},
	"Rotom-Frost": {
		"id":   uint(479),
		"form": uint(3),
		"type": "electric",
	},
	"Rotom-Fan": {
		"id":   uint(479),
		"form": uint(4),
		"type": "electric",
	},
	"Rotom-Mow": {
		"id":   uint(479),
		"form": uint(5),
		"type": "electric",
	},
	"Uxie": {
		"id":   uint(480),
		"form": uint(0),
		"type": "psychic",
	},
	"Mesprit": {
		"id":   uint(481),
		"form": uint(0),
		"type": "psychic",
	},
	"Azelf": {
		"id":   uint(482),
		"form": uint(0),
		"type": "psychic",
	},
	"Dialga": {
		"id":   uint(483),
		"form": uint(0),
		"type": "steel",
	},
	"Palkia": {
		"id":   uint(484),
		"form": uint(0),
		"type": "water",
	},
	"Heatran": {
		"id":   uint(485),
		"form": uint(0),
		"type": "fire",
	},
	"Regigigas": {
		"id":   uint(486),
		"form": uint(0),
		"type": "normal",
	},
	"Giratina": {
		"id":   uint(487),
		"form": uint(0),
		"type": "ghost",
	},
	"Giratina-Origin": {
		"id":   uint(487),
		"form": uint(1),
		"type": "ghost",
	},
	"Cresselia": {
		"id":   uint(488),
		"form": uint(0),
		"type": "psychic",
	},
	"Phione": {
		"id":   uint(489),
		"form": uint(0),
		"type": "water",
	},
	"Manaphy": {
		"id":   uint(490),
		"form": uint(0),
		"type": "water",
	},
	"Darkrai": {
		"id":   uint(491),
		"form": uint(0),
		"type": "dark",
	},
	"Shaymin": {
		"id":   uint(492),
		"form": uint(0),
		"type": "grass",
	},
	"Shaymin-Sky": {
		"id":   uint(492),
		"form": uint(1),
		"type": "grass",
	},
	"Arceus-Fairy": {
		"id":   uint(493),
		"form": uint(17),
		"type": "fairy",
	},
	"Arceus-Dark": {
		"id":   uint(493),
		"form": uint(16),
		"type": "dark",
	},
	"Arceus-Dragon": {
		"id":   uint(493),
		"form": uint(15),
		"type": "dragon",
	},
	"Arceus-Ice": {
		"id":   uint(493),
		"form": uint(14),
		"type": "ice",
	},
	"Arceus-Psychic": {
		"id":   uint(493),
		"form": uint(13),
		"type": "psychic",
	},
	"Arceus-Electric": {
		"id":   uint(493),
		"form": uint(12),
		"type": "electric",
	},
	"Arceus-Grass": {
		"id":   uint(493),
		"form": uint(11),
		"type": "grass",
	},
	"Arceus-Water": {
		"id":   uint(493),
		"form": uint(10),
		"type": "water",
	},
	"Arceus-Fire": {
		"id":   uint(493),
		"form": uint(9),
		"type": "fire",
	},
	"Arceus-Steel": {
		"id":   uint(493),
		"form": uint(8),
		"type": "steel",
	},
	"Arceus-Ghost": {
		"id":   uint(493),
		"form": uint(7),
		"type": "ghost",
	},
	"Arceus-Bug": {
		"id":   uint(493),
		"form": uint(6),
		"type": "bug",
	},
	"Arceus-Rock": {
		"id":   uint(493),
		"form": uint(5),
		"type": "rock",
	},
	"Arceus-Ground": {
		"id":   uint(493),
		"form": uint(4),
		"type": "ground",
	},
	"Arceus-Poison": {
		"id":   uint(493),
		"form": uint(3),
		"type": "poison",
	},
	"Arceus-Flying": {
		"id":   uint(493),
		"form": uint(2),
		"type": "flying",
	},
	"Arceus-Fighting": {
		"id":   uint(493),
		"form": uint(1),
		"type": "fighting",
	},
	"Arceus": {
		"id":   uint(493),
		"form": uint(0),
		"type": "normal",
	},
	"Victini": {
		"id":   uint(494),
		"form": uint(0),
		"type": "psychic",
	},
	"Snivy": {
		"id":   uint(495),
		"form": uint(0),
		"type": "grass",
	},
	"Servine": {
		"id":   uint(496),
		"form": uint(0),
		"type": "grass",
	},
	"Serperior": {
		"id":   uint(497),
		"form": uint(0),
		"type": "grass",
	},
	"Tepig": {
		"id":   uint(498),
		"form": uint(0),
		"type": "fire",
	},
	"Pignite": {
		"id":   uint(499),
		"form": uint(0),
		"type": "fire",
	},
	"Emboar": {
		"id":   uint(500),
		"form": uint(0),
		"type": "fire",
	},
	"Oshawott": {
		"id":   uint(501),
		"form": uint(0),
		"type": "water",
	},
	"Dewott": {
		"id":   uint(502),
		"form": uint(0),
		"type": "water",
	},
	"Samurott": {
		"id":   uint(503),
		"form": uint(0),
		"type": "water",
	},
	"Patrat": {
		"id":   uint(504),
		"form": uint(0),
		"type": "normal",
	},
	"Watchog": {
		"id":   uint(505),
		"form": uint(0),
		"type": "normal",
	},
	"Lillipup": {
		"id":   uint(506),
		"form": uint(0),
		"type": "normal",
	},
	"Herdier": {
		"id":   uint(507),
		"form": uint(0),
		"type": "normal",
	},
	"Stoutland": {
		"id":   uint(508),
		"form": uint(0),
		"type": "normal",
	},
	"Purrloin": {
		"id":   uint(509),
		"form": uint(0),
		"type": "dark",
	},
	"Liepard": {
		"id":   uint(510),
		"form": uint(0),
		"type": "dark",
	},
	"Pansage": {
		"id":   uint(511),
		"form": uint(0),
		"type": "grass",
	},
	"Simisage": {
		"id":   uint(512),
		"form": uint(0),
		"type": "grass",
	},
	"Pansear": {
		"id":   uint(513),
		"form": uint(0),
		"type": "fire",
	},
	"Simisear": {
		"id":   uint(514),
		"form": uint(0),
		"type": "fire",
	},
	"Panpour": {
		"id":   uint(515),
		"form": uint(0),
		"type": "water",
	},
	"Simipour": {
		"id":   uint(516),
		"form": uint(0),
		"type": "water",
	},
	"Munna": {
		"id":   uint(517),
		"form": uint(0),
		"type": "psychic",
	},
	"Musharna": {
		"id":   uint(518),
		"form": uint(0),
		"type": "psychic",
	},
	"Pidove": {
		"id":   uint(519),
		"form": uint(0),
		"type": "normal",
	},
	"Tranquill": {
		"id":   uint(520),
		"form": uint(0),
		"type": "normal",
	},
	"Unfezant": {
		"id":   uint(521),
		"form": uint(0),
		"type": "normal",
	},
	"Blitzle": {
		"id":   uint(522),
		"form": uint(0),
		"type": "electric",
	},
	"Zebstrika": {
		"id":   uint(523),
		"form": uint(0),
		"type": "electric",
	},
	"Roggenrola": {
		"id":   uint(524),
		"form": uint(0),
		"type": "rock",
	},
	"Boldore": {
		"id":   uint(525),
		"form": uint(0),
		"type": "rock",
	},
	"Gigalith": {
		"id":   uint(526),
		"form": uint(0),
		"type": "rock",
	},
	"Woobat": {
		"id":   uint(527),
		"form": uint(0),
		"type": "psychic",
	},
	"Swoobat": {
		"id":   uint(528),
		"form": uint(0),
		"type": "psychic",
	},
	"Drilbur": {
		"id":   uint(529),
		"form": uint(0),
		"type": "ground",
	},
	"Excadrill": {
		"id":   uint(530),
		"form": uint(0),
		"type": "ground",
	},
	"Audino": {
		"id":   uint(531),
		"form": uint(0),
		"type": "normal",
	},
	"Audino-Mega": {
		"id":   uint(531),
		"form": uint(1),
		"type": "normal",
	},
	"Timburr": {
		"id":   uint(532),
		"form": uint(0),
		"type": "fighting",
	},
	"Gurdurr": {
		"id":   uint(533),
		"form": uint(0),
		"type": "fighting",
	},
	"Conkeldurr": {
		"id":   uint(534),
		"form": uint(0),
		"type": "fighting",
	},
	"Tympole": {
		"id":   uint(535),
		"form": uint(0),
		"type": "water",
	},
	"Palpitoad": {
		"id":   uint(536),
		"form": uint(0),
		"type": "water",
	},
	"Seismitoad": {
		"id":   uint(537),
		"form": uint(0),
		"type": "water",
	},
	"Throh": {
		"id":   uint(538),
		"form": uint(0),
		"type": "fighting",
	},
	"Sawk": {
		"id":   uint(539),
		"form": uint(0),
		"type": "fighting",
	},
	"Sewaddle": {
		"id":   uint(540),
		"form": uint(0),
		"type": "bug",
	},
	"Swadloon": {
		"id":   uint(541),
		"form": uint(0),
		"type": "bug",
	},
	"Leavanny": {
		"id":   uint(542),
		"form": uint(0),
		"type": "bug",
	},
	"Venipede": {
		"id":   uint(543),
		"form": uint(0),
		"type": "bug",
	},
	"Whirlipede": {
		"id":   uint(544),
		"form": uint(0),
		"type": "bug",
	},
	"Scolipede": {
		"id":   uint(545),
		"form": uint(0),
		"type": "bug",
	},
	"Cottonee": {
		"id":   uint(546),
		"form": uint(0),
		"type": "grass",
	},
	"Whimsicott": {
		"id":   uint(547),
		"form": uint(0),
		"type": "grass",
	},
	"Petilil": {
		"id":   uint(548),
		"form": uint(0),
		"type": "grass",
	},
	"Lilligant": {
		"id":   uint(549),
		"form": uint(0),
		"type": "grass",
	},
	"Basculin": {
		"id":   uint(550),
		"form": uint(0),
		"type": "water",
	},
	"Basculin-Blue-Striped": {
		"id":   uint(550),
		"form": uint(1),
		"type": "water",
	},
	"Sandile": {
		"id":   uint(551),
		"form": uint(0),
		"type": "ground",
	},
	"Krokorok": {
		"id":   uint(552),
		"form": uint(0),
		"type": "ground",
	},
	"Krookodile": {
		"id":   uint(553),
		"form": uint(0),
		"type": "ground",
	},
	"Darumaka": {
		"id":   uint(554),
		"form": uint(0),
		"type": "fire",
	},
	"Darmanitan": {
		"id":   uint(555),
		"form": uint(0),
		"type": "fire",
	},
	"Darmanitan-Zen": {
		"id":   uint(555),
		"form": uint(1),
		"type": "fire",
	},
	"Darmanitan-Galar": {
		"id":   uint(555),
		"form": uint(2),
		"type": "ice",
	},
	"Darmanitan-Galar-Zen": {
		"id":   uint(555),
		"form": uint(3),
		"type": "ice",
	},
	"Maractus": {
		"id":   uint(556),
		"form": uint(0),
		"type": "grass",
	},
	"Dwebble": {
		"id":   uint(557),
		"form": uint(0),
		"type": "bug",
	},
	"Crustle": {
		"id":   uint(558),
		"form": uint(0),
		"type": "bug",
	},
	"Scraggy": {
		"id":   uint(559),
		"form": uint(0),
		"type": "dark",
	},
	"Scrafty": {
		"id":   uint(560),
		"form": uint(0),
		"type": "dark",
	},
	"Sigilyph": {
		"id":   uint(561),
		"form": uint(0),
		"type": "psychic",
	},
	"Yamask": {
		"id":   uint(562),
		"form": uint(0),
		"type": "ghost",
	},
	"Yamask-Galar": {
		"id":   uint(562),
		"form": uint(1),
		"type": "ground",
	},
	"Cofagrigus": {
		"id":   uint(563),
		"form": uint(0),
		"type": "ghost",
	},
	"Tirtouga": {
		"id":   uint(564),
		"form": uint(0),
		"type": "water",
	},
	"Carracosta": {
		"id":   uint(565),
		"form": uint(0),
		"type": "water",
	},
	"Archen": {
		"id":   uint(566),
		"form": uint(0),
		"type": "rock",
	},
	"Archeops": {
		"id":   uint(567),
		"form": uint(0),
		"type": "rock",
	},
	"Trubbish": {
		"id":   uint(568),
		"form": uint(0),
		"type": "poison",
	},
	"Garbodor": {
		"id":   uint(569),
		"form": uint(0),
		"type": "poison",
	},
	"Garbodor-Gmax": {
		"id":   uint(569),
		"form": uint(1),
		"type": "poison",
	},
	"Zorua": {
		"id":   uint(570),
		"form": uint(0),
		"type": "dark",
	},
	"Zoroark": {
		"id":   uint(571),
		"form": uint(0),
		"type": "dark",
	},
	"Minccino": {
		"id":   uint(572),
		"form": uint(0),
		"type": "normal",
	},
	"Cinccino": {
		"id":   uint(573),
		"form": uint(0),
		"type": "normal",
	},
	"Gothita": {
		"id":   uint(574),
		"form": uint(0),
		"type": "psychic",
	},
	"Gothorita": {
		"id":   uint(575),
		"form": uint(0),
		"type": "psychic",
	},
	"Gothitelle": {
		"id":   uint(576),
		"form": uint(0),
		"type": "psychic",
	},
	"Solosis": {
		"id":   uint(577),
		"form": uint(0),
		"type": "psychic",
	},
	"Duosion": {
		"id":   uint(578),
		"form": uint(0),
		"type": "psychic",
	},
	"Reuniclus": {
		"id":   uint(579),
		"form": uint(0),
		"type": "psychic",
	},
	"Ducklett": {
		"id":   uint(580),
		"form": uint(0),
		"type": "water",
	},
	"Swanna": {
		"id":   uint(581),
		"form": uint(0),
		"type": "water",
	},
	"Vanillite": {
		"id":   uint(582),
		"form": uint(0),
		"type": "ice",
	},
	"Vanillish": {
		"id":   uint(583),
		"form": uint(0),
		"type": "ice",
	},
	"Vanilluxe": {
		"id":   uint(584),
		"form": uint(0),
		"type": "ice",
	},
	"Deerling": {
		"id":   uint(585),
		"form": uint(0),
		"type": "normal",
	},
	"Deerling-Summer": {
		"id":   uint(585),
		"form": uint(1),
		"type": "normal",
	},
	"Deerling-Autumn": {
		"id":   uint(585),
		"form": uint(2),
		"type": "normal",
	},
	"Deerling-Winter": {
		"id":   uint(585),
		"form": uint(3),
		"type": "normal",
	},
	"Sawsbuck": {
		"id":   uint(586),
		"form": uint(0),
		"type": "normal",
	},
	"Sawsbuck-Summer": {
		"id":   uint(586),
		"form": uint(1),
		"type": "normal",
	},
	"Sawsbuck-Autumn": {
		"id":   uint(586),
		"form": uint(2),
		"type": "normal",
	},
	"Sawsbuck-Winter": {
		"id":   uint(586),
		"form": uint(3),
		"type": "normal",
	},
	"Emolga": {
		"id":   uint(587),
		"form": uint(0),
		"type": "electric",
	},
	"Karrablast": {
		"id":   uint(588),
		"form": uint(0),
		"type": "bug",
	},
	"Escavalier": {
		"id":   uint(589),
		"form": uint(0),
		"type": "bug",
	},
	"Foongus": {
		"id":   uint(590),
		"form": uint(0),
		"type": "grass",
	},
	"Amoonguss": {
		"id":   uint(591),
		"form": uint(0),
		"type": "grass",
	},
	"Frillish": {
		"id":   uint(592),
		"form": uint(0),
		"type": "water",
	},
	"Jellicent": {
		"id":   uint(593),
		"form": uint(0),
		"type": "water",
	},
	"Alomomola": {
		"id":   uint(594),
		"form": uint(0),
		"type": "water",
	},
	"Joltik": {
		"id":   uint(595),
		"form": uint(0),
		"type": "bug",
	},
	"Galvantula": {
		"id":   uint(596),
		"form": uint(0),
		"type": "bug",
	},
	"Ferroseed": {
		"id":   uint(597),
		"form": uint(0),
		"type": "grass",
	},
	"Ferrothorn": {
		"id":   uint(598),
		"form": uint(0),
		"type": "grass",
	},
	"Klink": {
		"id":   uint(599),
		"form": uint(0),
		"type": "steel",
	},
	"Klang": {
		"id":   uint(600),
		"form": uint(0),
		"type": "steel",
	},
	"Klinklang": {
		"id":   uint(601),
		"form": uint(0),
		"type": "steel",
	},
	"Tynamo": {
		"id":   uint(602),
		"form": uint(0),
		"type": "electric",
	},
	"Eelektrik": {
		"id":   uint(603),
		"form": uint(0),
		"type": "electric",
	},
	"Eelektross": {
		"id":   uint(604),
		"form": uint(0),
		"type": "electric",
	},
	"Elgyem": {
		"id":   uint(605),
		"form": uint(0),
		"type": "psychic",
	},
	"Beheeyem": {
		"id":   uint(606),
		"form": uint(0),
		"type": "psychic",
	},
	"Litwick": {
		"id":   uint(607),
		"form": uint(0),
		"type": "ghost",
	},
	"Lampent": {
		"id":   uint(608),
		"form": uint(0),
		"type": "ghost",
	},
	"Chandelure": {
		"id":   uint(609),
		"form": uint(0),
		"type": "ghost",
	},
	"Axew": {
		"id":   uint(610),
		"form": uint(0),
		"type": "dragon",
	},
	"Fraxure": {
		"id":   uint(611),
		"form": uint(0),
		"type": "dragon",
	},
	"Haxorus": {
		"id":   uint(612),
		"form": uint(0),
		"type": "dragon",
	},
	"Cubchoo": {
		"id":   uint(613),
		"form": uint(0),
		"type": "ice",
	},
	"Beartic": {
		"id":   uint(614),
		"form": uint(0),
		"type": "ice",
	},
	"Cryogonal": {
		"id":   uint(615),
		"form": uint(0),
		"type": "ice",
	},
	"Shelmet": {
		"id":   uint(616),
		"form": uint(0),
		"type": "bug",
	},
	"Accelgor": {
		"id":   uint(617),
		"form": uint(0),
		"type": "bug",
	},
	"Stunfisk": {
		"id":   uint(618),
		"form": uint(0),
		"type": "ground",
	},
	"Stunfisk-Galar": {
		"id":   uint(618),
		"form": uint(1),
		"type": "ground",
	},
	"Mienfoo": {
		"id":   uint(619),
		"form": uint(0),
		"type": "fighting",
	},
	"Mienshao": {
		"id":   uint(620),
		"form": uint(0),
		"type": "fighting",
	},
	"Druddigon": {
		"id":   uint(621),
		"form": uint(0),
		"type": "dragon",
	},
	"Golett": {
		"id":   uint(622),
		"form": uint(0),
		"type": "ground",
	},
	"Golurk": {
		"id":   uint(623),
		"form": uint(0),
		"type": "ground",
	},
	"Pawniard": {
		"id":   uint(624),
		"form": uint(0),
		"type": "dark",
	},
	"Bisharp": {
		"id":   uint(625),
		"form": uint(0),
		"type": "dark",
	},
	"Bouffalant": {
		"id":   uint(626),
		"form": uint(0),
		"type": "normal",
	},
	"Rufflet": {
		"id":   uint(627),
		"form": uint(0),
		"type": "normal",
	},
	"Braviary": {
		"id":   uint(628),
		"form": uint(0),
		"type": "normal",
	},
	"Vullaby": {
		"id":   uint(629),
		"form": uint(0),
		"type": "dark",
	},
	"Mandibuzz": {
		"id":   uint(630),
		"form": uint(0),
		"type": "dark",
	},
	"Heatmor": {
		"id":   uint(631),
		"form": uint(0),
		"type": "fire",
	},
	"Durant": {
		"id":   uint(632),
		"form": uint(0),
		"type": "bug",
	},
	"Deino": {
		"id":   uint(633),
		"form": uint(0),
		"type": "dark",
	},
	"Zweilous": {
		"id":   uint(634),
		"form": uint(0),
		"type": "dark",
	},
	"Hydreigon": {
		"id":   uint(635),
		"form": uint(0),
		"type": "dark",
	},
	"Larvesta": {
		"id":   uint(636),
		"form": uint(0),
		"type": "bug",
	},
	"Volcarona": {
		"id":   uint(637),
		"form": uint(0),
		"type": "bug",
	},
	"Cobalion": {
		"id":   uint(638),
		"form": uint(0),
		"type": "steel",
	},
	"Terrakion": {
		"id":   uint(639),
		"form": uint(0),
		"type": "rock",
	},
	"Virizion": {
		"id":   uint(640),
		"form": uint(0),
		"type": "grass",
	},
	"Tornadus": {
		"id":   uint(641),
		"form": uint(0),
		"type": "flying",
	},
	"Tornadus-Therian": {
		"id":   uint(641),
		"form": uint(1),
		"type": "flying",
	},
	"Thundurus": {
		"id":   uint(642),
		"form": uint(0),
		"type": "electric",
	},
	"Thundurus-Therian": {
		"id":   uint(642),
		"form": uint(1),
		"type": "electric",
	},
	"Reshiram": {
		"id":   uint(643),
		"form": uint(0),
		"type": "dragon",
	},
	"Zekrom": {
		"id":   uint(644),
		"form": uint(0),
		"type": "dragon",
	},
	"Landorus": {
		"id":   uint(645),
		"form": uint(0),
		"type": "ground",
	},
	"Landorus-Therian": {
		"id":   uint(645),
		"form": uint(1),
		"type": "ground",
	},
	"Kyurem": {
		"id":   uint(646),
		"form": uint(0),
		"type": "dragon",
	},
	"Kyurem-White": {
		"id":   uint(646),
		"form": uint(1),
		"type": "dragon",
	},
	"Kyurem-Black": {
		"id":   uint(646),
		"form": uint(2),
		"type": "dragon",
	},
	"Keldeo": {
		"id":   uint(647),
		"form": uint(0),
		"type": "water",
	},
	"Keldeo-Resolute": {
		"id":   uint(647),
		"form": uint(1),
		"type": "water",
	},
	"Meloetta": {
		"id":   uint(648),
		"form": uint(0),
		"type": "normal",
	},
	"Meloetta-Pirouette": {
		"id":   uint(648),
		"form": uint(1),
		"type": "normal",
	},
	"Genesect": {
		"id":   uint(649),
		"form": uint(0),
		"type": "bug",
	},
	"Chespin": {
		"id":   uint(650),
		"form": uint(0),
		"type": "grass",
	},
	"Quilladin": {
		"id":   uint(651),
		"form": uint(0),
		"type": "grass",
	},
	"Chesnaught": {
		"id":   uint(652),
		"form": uint(0),
		"type": "grass",
	},
	"Fennekin": {
		"id":   uint(653),
		"form": uint(0),
		"type": "fire",
	},
	"Braixen": {
		"id":   uint(654),
		"form": uint(0),
		"type": "fire",
	},
	"Delphox": {
		"id":   uint(655),
		"form": uint(0),
		"type": "fire",
	},
	"Froakie": {
		"id":   uint(656),
		"form": uint(0),
		"type": "water",
	},
	"Frogadier": {
		"id":   uint(657),
		"form": uint(0),
		"type": "water",
	},
	"Greninja-Ash": {
		"id":   uint(658),
		"form": uint(2),
		"type": "water",
	},
	"Greninja": {
		"id":   uint(658),
		"form": uint(0),
		"type": "water",
	},
	"Bunnelby": {
		"id":   uint(659),
		"form": uint(0),
		"type": "normal",
	},
	"Diggersby": {
		"id":   uint(660),
		"form": uint(0),
		"type": "normal",
	},
	"Fletchling": {
		"id":   uint(661),
		"form": uint(0),
		"type": "normal",
	},
	"Fletchinder": {
		"id":   uint(662),
		"form": uint(0),
		"type": "fire",
	},
	"Talonflame": {
		"id":   uint(663),
		"form": uint(0),
		"type": "fire",
	},
	"Scatterbug": {
		"id":   uint(664),
		"form": uint(0),
		"type": "bug",
	},
	"Spewpa": {
		"id":   uint(665),
		"form": uint(0),
		"type": "bug",
	},
	"Vivillon-Icysnow": {
		"id":   uint(666),
		"form": uint(0),
		"type": "bug",
	},
	"Vivillon-Polar": {
		"id":   uint(666),
		"form": uint(1),
		"type": "bug",
	},
	"Vivillon-Tundra": {
		"id":   uint(666),
		"form": uint(2),
		"type": "bug",
	},
	"Vivillon-Continental": {
		"id":   uint(666),
		"form": uint(3),
		"type": "bug",
	},
	"Vivillon-Garden": {
		"id":   uint(666),
		"form": uint(4),
		"type": "bug",
	},
	"Vivillon-Elegant": {
		"id":   uint(666),
		"form": uint(5),
		"type": "bug",
	},
	"Vivillon": {
		"id":   uint(666),
		"form": uint(6),
		"type": "bug",
	},
	"Vivillon-Modern": {
		"id":   uint(666),
		"form": uint(7),
		"type": "bug",
	},
	"Vivillon-Marine": {
		"id":   uint(666),
		"form": uint(8),
		"type": "bug",
	},
	"Vivillon-Archipelago": {
		"id":   uint(666),
		"form": uint(9),
		"type": "bug",
	},
	"Vivillon-Highplains": {
		"id":   uint(666),
		"form": uint(10),
		"type": "bug",
	},
	"Vivillon-Sandstorm": {
		"id":   uint(666),
		"form": uint(11),
		"type": "bug",
	},
	"Vivillon-River": {
		"id":   uint(666),
		"form": uint(12),
		"type": "bug",
	},
	"Vivillon-Monsoon": {
		"id":   uint(666),
		"form": uint(13),
		"type": "bug",
	},
	"Vivillon-Savanna": {
		"id":   uint(666),
		"form": uint(14),
		"type": "bug",
	},
	"Vivillon-Sun": {
		"id":   uint(666),
		"form": uint(15),
		"type": "bug",
	},
	"Vivillon-Ocean": {
		"id":   uint(666),
		"form": uint(16),
		"type": "bug",
	},
	"Vivillon-Jungle": {
		"id":   uint(666),
		"form": uint(17),
		"type": "bug",
	},
	"Vivillon-Fancy": {
		"id":   uint(666),
		"form": uint(18),
		"type": "bug",
	},
	"Vivillon-Pokeball": {
		"id":   uint(666),
		"form": uint(19),
		"type": "bug",
	},
	"Litleo": {
		"id":   uint(667),
		"form": uint(0),
		"type": "fire",
	},
	"Pyroar": {
		"id":   uint(668),
		"form": uint(0),
		"type": "fire",
	},
	"Flabebe": {
		"id":   uint(669),
		"form": uint(0),
		"type": "fairy",
	},
	"Flabebe-Yellow": {
		"id":   uint(669),
		"form": uint(1),
		"type": "fairy",
	},
	"Flabebe-Orange": {
		"id":   uint(669),
		"form": uint(2),
		"type": "fairy",
	},
	"Flabebe-Blue": {
		"id":   uint(669),
		"form": uint(3),
		"type": "fairy",
	},
	"Flabebe-White": {
		"id":   uint(669),
		"form": uint(4),
		"type": "fairy",
	},
	"Floette": {
		"id":   uint(670),
		"form": uint(0),
		"type": "fairy",
	},
	"Floette-Yellow": {
		"id":   uint(670),
		"form": uint(1),
		"type": "fairy",
	},
	"Floette-Orange": {
		"id":   uint(670),
		"form": uint(2),
		"type": "fairy",
	},
	"Floette-Blue": {
		"id":   uint(670),
		"form": uint(3),
		"type": "fairy",
	},
	"Floette-White": {
		"id":   uint(670),
		"form": uint(4),
		"type": "fairy",
	},
	"Floette-Eternal": {
		"id":   uint(670),
		"form": uint(5),
		"type": "fairy",
	},
	"Florges": {
		"id":   uint(671),
		"form": uint(0),
		"type": "fairy",
	},
	"Florges-Yellow": {
		"id":   uint(671),
		"form": uint(1),
		"type": "fairy",
	},
	"Florges-Orange": {
		"id":   uint(671),
		"form": uint(2),
		"type": "fairy",
	},
	"Florges-Blue": {
		"id":   uint(671),
		"form": uint(3),
		"type": "fairy",
	},
	"Florges-White": {
		"id":   uint(671),
		"form": uint(4),
		"type": "fairy",
	},
	"Skiddo": {
		"id":   uint(672),
		"form": uint(0),
		"type": "grass",
	},
	"Gogoat": {
		"id":   uint(673),
		"form": uint(0),
		"type": "grass",
	},
	"Pancham": {
		"id":   uint(674),
		"form": uint(0),
		"type": "fighting",
	},
	"Pangoro": {
		"id":   uint(675),
		"form": uint(0),
		"type": "fighting",
	},
	"Furfrou": {
		"id":   uint(676),
		"form": uint(0),
		"type": "normal",
	},
	"Espurr": {
		"id":   uint(677),
		"form": uint(0),
		"type": "psychic",
	},
	"Meowstic": {
		"id":   uint(678),
		"form": uint(0),
		"type": "psychic",
	},
	"Meowstic-F": {
		"id":   uint(678),
		"form": uint(1),
		"type": "psychic",
	},
	"Honedge": {
		"id":   uint(679),
		"form": uint(0),
		"type": "steel",
	},
	"Doublade": {
		"id":   uint(680),
		"form": uint(0),
		"type": "steel",
	},
	"Aegislash": {
		"id":   uint(681),
		"form": uint(0),
		"type": "steel",
	},
	"Aegislash-Blade": {
		"id":   uint(681),
		"form": uint(1),
		"type": "steel",
	},
	"Spritzee": {
		"id":   uint(682),
		"form": uint(0),
		"type": "fairy",
	},
	"Aromatisse": {
		"id":   uint(683),
		"form": uint(0),
		"type": "fairy",
	},
	"Swirlix": {
		"id":   uint(684),
		"form": uint(0),
		"type": "fairy",
	},
	"Slurpuff": {
		"id":   uint(685),
		"form": uint(0),
		"type": "fairy",
	},
	"Inkay": {
		"id":   uint(686),
		"form": uint(0),
		"type": "dark",
	},
	"Malamar": {
		"id":   uint(687),
		"form": uint(0),
		"type": "dark",
	},
	"Binacle": {
		"id":   uint(688),
		"form": uint(0),
		"type": "rock",
	},
	"Barbaracle": {
		"id":   uint(689),
		"form": uint(0),
		"type": "rock",
	},
	"Skrelp": {
		"id":   uint(690),
		"form": uint(0),
		"type": "poison",
	},
	"Dragalge": {
		"id":   uint(691),
		"form": uint(0),
		"type": "poison",
	},
	"Clauncher": {
		"id":   uint(692),
		"form": uint(0),
		"type": "water",
	},
	"Clawitzer": {
		"id":   uint(693),
		"form": uint(0),
		"type": "water",
	},
	"Helioptile": {
		"id":   uint(694),
		"form": uint(0),
		"type": "electric",
	},
	"Heliolisk": {
		"id":   uint(695),
		"form": uint(0),
		"type": "electric",
	},
	"Tyrunt": {
		"id":   uint(696),
		"form": uint(0),
		"type": "rock",
	},
	"Tyrantrum": {
		"id":   uint(697),
		"form": uint(0),
		"type": "rock",
	},
	"Amaura": {
		"id":   uint(698),
		"form": uint(0),
		"type": "rock",
	},
	"Aurorus": {
		"id":   uint(699),
		"form": uint(0),
		"type": "rock",
	},
	"Sylveon": {
		"id":   uint(700),
		"form": uint(0),
		"type": "fairy",
	},
	"Hawlucha": {
		"id":   uint(701),
		"form": uint(0),
		"type": "fighting",
	},
	"Dedenne": {
		"id":   uint(702),
		"form": uint(0),
		"type": "electric",
	},
	"Carbink": {
		"id":   uint(703),
		"form": uint(0),
		"type": "rock",
	},
	"Goomy": {
		"id":   uint(704),
		"form": uint(0),
		"type": "dragon",
	},
	"Sliggoo": {
		"id":   uint(705),
		"form": uint(0),
		"type": "dragon",
	},
	"Goodra": {
		"id":   uint(706),
		"form": uint(0),
		"type": "dragon",
	},
	"Klefki": {
		"id":   uint(707),
		"form": uint(0),
		"type": "steel",
	},
	"Phantump": {
		"id":   uint(708),
		"form": uint(0),
		"type": "ghost",
	},
	"Trevenant": {
		"id":   uint(709),
		"form": uint(0),
		"type": "ghost",
	},
	"Pumpkaboo": {
		"id":   uint(710),
		"form": uint(0),
		"type": "ghost",
	},
	"Pumpkaboo-Small": {
		"id":   uint(710),
		"form": uint(1),
		"type": "ghost",
	},
	"Pumpkaboo-Large": {
		"id":   uint(710),
		"form": uint(2),
		"type": "ghost",
	},
	"Pumpkaboo-Super": {
		"id":   uint(710),
		"form": uint(3),
		"type": "ghost",
	},
	"Gourgeist": {
		"id":   uint(711),
		"form": uint(0),
		"type": "ghost",
	},
	"Gourgeist-Small": {
		"id":   uint(711),
		"form": uint(1),
		"type": "ghost",
	},
	"Gourgeist-Large": {
		"id":   uint(711),
		"form": uint(2),
		"type": "ghost",
	},
	"Gourgeist-Super": {
		"id":   uint(711),
		"form": uint(3),
		"type": "ghost",
	},
	"Bergmite": {
		"id":   uint(712),
		"form": uint(0),
		"type": "ice",
	},
	"Avalugg": {
		"id":   uint(713),
		"form": uint(0),
		"type": "ice",
	},
	"Noibat": {
		"id":   uint(714),
		"form": uint(0),
		"type": "flying",
	},
	"Noivern": {
		"id":   uint(715),
		"form": uint(0),
		"type": "flying",
	},
	"Xerneas": {
		"id":   uint(716),
		"form": uint(0),
		"type": "fairy",
	},
	"Yveltal": {
		"id":   uint(717),
		"form": uint(0),
		"type": "dark",
	},
	"Zygarde": {
		"id":   uint(718),
		"form": uint(0),
		"type": "dragon",
	},
	"Zygarde-10%": {
		"id":   uint(718),
		"form": uint(1),
		"type": "dragon",
	},
	"Zygarde-Complete": {
		"id":   uint(718),
		"form": uint(4),
		"type": "dragon",
	},
	"Diancie": {
		"id":   uint(719),
		"form": uint(0),
		"type": "rock",
	},
	"Diancie-Mega": {
		"id":   uint(719),
		"form": uint(1),
		"type": "rock",
	},
	"Hoopa": {
		"id":   uint(720),
		"form": uint(0),
		"type": "psychic",
	},
	"Hoopa-Unbound": {
		"id":   uint(720),
		"form": uint(1),
		"type": "psychic",
	},
	"Volcanion": {
		"id":   uint(721),
		"form": uint(0),
		"type": "fire",
	},
	"Rowlet": {
		"id":   uint(722),
		"form": uint(0),
		"type": "grass",
	},
	"Dartrix": {
		"id":   uint(723),
		"form": uint(0),
		"type": "grass",
	},
	"Decidueye": {
		"id":   uint(724),
		"form": uint(0),
		"type": "grass",
	},
	"Litten": {
		"id":   uint(725),
		"form": uint(0),
		"type": "fire",
	},
	"Torracat": {
		"id":   uint(726),
		"form": uint(0),
		"type": "fire",
	},
	"Incineroar": {
		"id":   uint(727),
		"form": uint(0),
		"type": "fire",
	},
	"Popplio": {
		"id":   uint(728),
		"form": uint(0),
		"type": "water",
	},
	"Brionne": {
		"id":   uint(729),
		"form": uint(0),
		"type": "water",
	},
	"Primarina": {
		"id":   uint(730),
		"form": uint(0),
		"type": "water",
	},
	"Pikipek": {
		"id":   uint(731),
		"form": uint(0),
		"type": "normal",
	},
	"Trumbeak": {
		"id":   uint(732),
		"form": uint(0),
		"type": "normal",
	},
	"Toucannon": {
		"id":   uint(733),
		"form": uint(0),
		"type": "normal",
	},
	"Yungoos": {
		"id":   uint(734),
		"form": uint(0),
		"type": "normal",
	},
	"Gumshoos": {
		"id":   uint(735),
		"form": uint(0),
		"type": "normal",
	},
	"Gumshoos-Totem": {
		"id":   uint(735),
		"form": uint(0),
		"type": "normal",
	},
	"Grubbin": {
		"id":   uint(736),
		"form": uint(0),
		"type": "bug",
	},
	"Charjabug": {
		"id":   uint(737),
		"form": uint(0),
		"type": "bug",
	},
	"Vikavolt": {
		"id":   uint(738),
		"form": uint(0),
		"type": "bug",
	},
	"Vikavolt-Totem": {
		"id":   uint(738),
		"form": uint(0),
		"type": "bug",
	},
	"Crabrawler": {
		"id":   uint(739),
		"form": uint(0),
		"type": "fighting",
	},
	"Crabominable": {
		"id":   uint(740),
		"form": uint(0),
		"type": "fighting",
	},
	"Oricorio": {
		"id":   uint(741),
		"form": uint(0),
		"type": "fire",
	},
	"Oricorio-Pom-Pom": {
		"id":   uint(741),
		"form": uint(1),
		"type": "electric",
	},
	"Oricorio-Pa'u": {
		"id":   uint(741),
		"form": uint(2),
		"type": "psychic",
	},
	"Oricorio-Sensu": {
		"id":   uint(741),
		"form": uint(3),
		"type": "ghost",
	},
	"Cutiefly": {
		"id":   uint(742),
		"form": uint(0),
		"type": "bug",
	},
	"Ribombee": {
		"id":   uint(743),
		"form": uint(0),
		"type": "bug",
	},
	"Ribombee-Totem": {
		"id":   uint(743),
		"form": uint(0),
		"type": "bug",
	},
	"Rockruff": {
		"id":   uint(744),
		"form": uint(0),
		"type": "rock",
	},
	"Lycanroc": {
		"id":   uint(745),
		"form": uint(0),
		"type": "rock",
	},
	"Lycanroc-Midnight": {
		"id":   uint(745),
		"form": uint(1),
		"type": "rock",
	},
	"Lycanroc-Dusk": {
		"id":   uint(745),
		"form": uint(2),
		"type": "rock",
	},
	"Wishiwashi": {
		"id":   uint(746),
		"form": uint(0),
		"type": "water",
	},
	"Wishiwashi-School": {
		"id":   uint(746),
		"form": uint(1),
		"type": "water",
	},
	"Mareanie": {
		"id":   uint(747),
		"form": uint(0),
		"type": "poison",
	},
	"Toxapex": {
		"id":   uint(748),
		"form": uint(0),
		"type": "poison",
	},
	"Mudbray": {
		"id":   uint(749),
		"form": uint(0),
		"type": "ground",
	},
	"Mudsdale": {
		"id":   uint(750),
		"form": uint(0),
		"type": "ground",
	},
	"Dewpider": {
		"id":   uint(751),
		"form": uint(0),
		"type": "water",
	},
	"Araquanid": {
		"id":   uint(752),
		"form": uint(0),
		"type": "water",
	},
	"Araquanid-Totem": {
		"id":   uint(752),
		"form": uint(0),
		"type": "water",
	},
	"Fomantis": {
		"id":   uint(753),
		"form": uint(0),
		"type": "grass",
	},
	"Lurantis": {
		"id":   uint(754),
		"form": uint(0),
		"type": "grass",
	},
	"Lurantis-Totem": {
		"id":   uint(754),
		"form": uint(0),
		"type": "grass",
	},
	"Morelull": {
		"id":   uint(755),
		"form": uint(0),
		"type": "grass",
	},
	"Shiinotic": {
		"id":   uint(756),
		"form": uint(0),
		"type": "grass",
	},
	"Salandit": {
		"id":   uint(757),
		"form": uint(0),
		"type": "poison",
	},
	"Salazzle": {
		"id":   uint(758),
		"form": uint(0),
		"type": "poison",
	},
	"Salazzle-Totem": {
		"id":   uint(758),
		"form": uint(0),
		"type": "poison",
	},
	"Stufful": {
		"id":   uint(759),
		"form": uint(0),
		"type": "normal",
	},
	"Bewear": {
		"id":   uint(760),
		"form": uint(0),
		"type": "normal",
	},
	"Bounsweet": {
		"id":   uint(761),
		"form": uint(0),
		"type": "grass",
	},
	"Steenee": {
		"id":   uint(762),
		"form": uint(0),
		"type": "grass",
	},
	"Tsareena": {
		"id":   uint(763),
		"form": uint(0),
		"type": "grass",
	},
	"Comfey": {
		"id":   uint(764),
		"form": uint(0),
		"type": "fairy",
	},
	"Oranguru": {
		"id":   uint(765),
		"form": uint(0),
		"type": "normal",
	},
	"Passimian": {
		"id":   uint(766),
		"form": uint(0),
		"type": "fighting",
	},
	"Wimpod": {
		"id":   uint(767),
		"form": uint(0),
		"type": "bug",
	},
	"Golisopod": {
		"id":   uint(768),
		"form": uint(0),
		"type": "bug",
	},
	"Sandygast": {
		"id":   uint(769),
		"form": uint(0),
		"type": "ghost",
	},
	"Palossand": {
		"id":   uint(770),
		"form": uint(0),
		"type": "ghost",
	},
	"Pyukumuku": {
		"id":   uint(771),
		"form": uint(0),
		"type": "water",
	},
	"Type: Null": {
		"id":   uint(772),
		"form": uint(0),
		"type": "normal",
	},
	"Silvally-Fairy": {
		"id":   uint(773),
		"form": uint(17),
		"type": "fairy",
	},
	"Silvally-Dark": {
		"id":   uint(773),
		"form": uint(16),
		"type": "dark",
	},
	"Silvally-Dragon": {
		"id":   uint(773),
		"form": uint(15),
		"type": "dragon",
	},
	"Silvally-Ice": {
		"id":   uint(773),
		"form": uint(14),
		"type": "ice",
	},
	"Silvally-Psychic": {
		"id":   uint(773),
		"form": uint(13),
		"type": "psychic",
	},
	"Silvally-Electric": {
		"id":   uint(773),
		"form": uint(12),
		"type": "electric",
	},
	"Silvally-Grass": {
		"id":   uint(773),
		"form": uint(11),
		"type": "grass",
	},
	"Silvally-Water": {
		"id":   uint(773),
		"form": uint(10),
		"type": "water",
	},
	"Silvally-Fire": {
		"id":   uint(773),
		"form": uint(9),
		"type": "fire",
	},
	"Silvally-Steel": {
		"id":   uint(773),
		"form": uint(8),
		"type": "steel",
	},
	"Silvally-Ghost": {
		"id":   uint(773),
		"form": uint(7),
		"type": "ghost",
	},
	"Silvally-Bug": {
		"id":   uint(773),
		"form": uint(6),
		"type": "bug",
	},
	"Silvally-Rock": {
		"id":   uint(773),
		"form": uint(5),
		"type": "rock",
	},
	"Silvally-Ground": {
		"id":   uint(773),
		"form": uint(4),
		"type": "ground",
	},
	"Silvally-Poison": {
		"id":   uint(773),
		"form": uint(3),
		"type": "poison",
	},
	"Silvally-Flying": {
		"id":   uint(773),
		"form": uint(2),
		"type": "flying",
	},
	"Silvally-Fighting": {
		"id":   uint(773),
		"form": uint(1),
		"type": "fighting",
	},
	"Silvally": {
		"id":   uint(773),
		"form": uint(0),
		"type": "normal",
	},
	"Minior-Meteor": {
		"id":   uint(774),
		"form": uint(0),
		"type": "rock",
	},
	"Minior": {
		"id":   uint(774),
		"form": uint(7),
		"type": "rock",
	},
	"Minior-Orange": {
		"id":   uint(774),
		"form": uint(8),
		"type": "rock",
	},
	"Minior-Green": {
		"id":   uint(774),
		"form": uint(10),
		"type": "rock",
	},
	"Minior-Blue": {
		"id":   uint(774),
		"form": uint(11),
		"type": "rock",
	},
	"Minior-Indigo": {
		"id":   uint(774),
		"form": uint(12),
		"type": "rock",
	},
	"Minior-Violet": {
		"id":   uint(774),
		"form": uint(13),
		"type": "rock",
	},
	"Komala": {
		"id":   uint(775),
		"form": uint(0),
		"type": "normal",
	},
	"Turtonator": {
		"id":   uint(776),
		"form": uint(0),
		"type": "fire",
	},
	"Togedemaru": {
		"id":   uint(777),
		"form": uint(0),
		"type": "electric",
	},
	"Togedemaru-Totem": {
		"id":   uint(777),
		"form": uint(0),
		"type": "electric",
	},
	"Mimikyu": {
		"id":   uint(778),
		"form": uint(0),
		"type": "ghost",
	},
	"Mimikyu-Totem": {
		"id":   uint(778),
		"form": uint(0),
		"type": "ghost",
	},
	"Bruxish": {
		"id":   uint(779),
		"form": uint(0),
		"type": "water",
	},
	"Drampa": {
		"id":   uint(780),
		"form": uint(0),
		"type": "normal",
	},
	"Dhelmise": {
		"id":   uint(781),
		"form": uint(0),
		"type": "ghost",
	},
	"Jangmo-o": {
		"id":   uint(782),
		"form": uint(0),
		"type": "dragon",
	},
	"Hakamo-o": {
		"id":   uint(783),
		"form": uint(0),
		"type": "dragon",
	},
	"Kommo-o": {
		"id":   uint(784),
		"form": uint(0),
		"type": "dragon",
	},
	"Kommo-o-Totem": {
		"id":   uint(784),
		"form": uint(0),
		"type": "dragon",
	},
	"Tapu Koko": {
		"id":   uint(785),
		"form": uint(0),
		"type": "electric",
	},
	"Tapu Lele": {
		"id":   uint(786),
		"form": uint(0),
		"type": "psychic",
	},
	"Tapu Bulu": {
		"id":   uint(787),
		"form": uint(0),
		"type": "grass",
	},
	"Tapu Fini": {
		"id":   uint(788),
		"form": uint(0),
		"type": "water",
	},
	"Cosmog": {
		"id":   uint(789),
		"form": uint(0),
		"type": "psychic",
	},
	"Cosmoem": {
		"id":   uint(790),
		"form": uint(0),
		"type": "psychic",
	},
	"Solgaleo": {
		"id":   uint(791),
		"form": uint(0),
		"type": "psychic",
	},
	"Lunala": {
		"id":   uint(792),
		"form": uint(0),
		"type": "psychic",
	},
	"Nihilego": {
		"id":   uint(793),
		"form": uint(0),
		"type": "rock",
	},
	"Buzzwole": {
		"id":   uint(794),
		"form": uint(0),
		"type": "bug",
	},
	"Pheromosa": {
		"id":   uint(795),
		"form": uint(0),
		"type": "bug",
	},
	"Xurkitree": {
		"id":   uint(796),
		"form": uint(0),
		"type": "electric",
	},
	"Celesteela": {
		"id":   uint(797),
		"form": uint(0),
		"type": "steel",
	},
	"Kartana": {
		"id":   uint(798),
		"form": uint(0),
		"type": "grass",
	},
	"Guzzlord": {
		"id":   uint(799),
		"form": uint(0),
		"type": "dark",
	},
	"Necrozma": {
		"id":   uint(800),
		"form": uint(0),
		"type": "psychic",
	},
	"Necrozma-Dusk-Mane": {
		"id":   uint(800),
		"form": uint(1),
		"type": "psychic",
	},
	"Necrozma-Dawn-Wings": {
		"id":   uint(800),
		"form": uint(2),
		"type": "psychic",
	},
	"Necrozma-Ultra": {
		"id":   uint(800),
		"form": uint(3),
		"type": "psychic",
	},
	"Magearna": {
		"id":   uint(801),
		"form": uint(0),
		"type": "steel",
	},
	"Marshadow": {
		"id":   uint(802),
		"form": uint(0),
		"type": "ghost",
	},
	"Poipole": {
		"id":   uint(803),
		"form": uint(0),
		"type": "poison",
	},
	"Naganadel": {
		"id":   uint(804),
		"form": uint(0),
		"type": "poison",
	},
	"Stakataka": {
		"id":   uint(805),
		"form": uint(0),
		"type": "rock",
	},
	"Blacephalon": {
		"id":   uint(806),
		"form": uint(0),
		"type": "fire",
	},
	"Zeraora": {
		"id":   uint(807),
		"form": uint(0),
		"type": "electric",
	},
	"Meltan": {
		"id":   uint(808),
		"form": uint(0),
		"type": "steel",
	},
	"Melmetal": {
		"id":   uint(809),
		"form": uint(0),
		"type": "steel",
	},
	"Melmetal-Gmax": {
		"id":   uint(809),
		"form": uint(1),
		"type": "steel",
	},
	"Grookey": {
		"id":   uint(810),
		"form": uint(0),
		"type": "grass",
	},
	"Thwackey": {
		"id":   uint(811),
		"form": uint(0),
		"type": "grass",
	},
	"Rillaboom": {
		"id":   uint(812),
		"form": uint(0),
		"type": "grass",
	},
	"Rillaboom-Gmax": {
		"id":   uint(812),
		"form": uint(1),
		"type": "grass",
	},
	"Scorbunny": {
		"id":   uint(813),
		"form": uint(0),
		"type": "fire",
	},
	"Raboot": {
		"id":   uint(814),
		"form": uint(0),
		"type": "fire",
	},
	"Cinderace": {
		"id":   uint(815),
		"form": uint(0),
		"type": "fire",
	},
	"Cinderace-Gmax": {
		"id":   uint(815),
		"form": uint(1),
		"type": "fire",
	},
	"Sobble": {
		"id":   uint(816),
		"form": uint(0),
		"type": "water",
	},
	"Drizzile": {
		"id":   uint(817),
		"form": uint(0),
		"type": "water",
	},
	"Inteleon": {
		"id":   uint(818),
		"form": uint(0),
		"type": "water",
	},
	"Inteleon-Gmax": {
		"id":   uint(818),
		"form": uint(1),
		"type": "water",
	},
	"Skwovet": {
		"id":   uint(819),
		"form": uint(0),
		"type": "normal",
	},
	"Greedent": {
		"id":   uint(820),
		"form": uint(0),
		"type": "normal",
	},
	"Rookidee": {
		"id":   uint(821),
		"form": uint(0),
		"type": "flying",
	},
	"Corvisquire": {
		"id":   uint(822),
		"form": uint(0),
		"type": "flying",
	},
	"Corviknight": {
		"id":   uint(823),
		"form": uint(0),
		"type": "flying",
	},
	"Corviknight-Gmax": {
		"id":   uint(823),
		"form": uint(1),
		"type": "flying",
	},
	"Blipbug": {
		"id":   uint(824),
		"form": uint(0),
		"type": "bug",
	},
	"Dottler": {
		"id":   uint(825),
		"form": uint(0),
		"type": "bug",
	},
	"Orbeetle": {
		"id":   uint(826),
		"form": uint(0),
		"type": "bug",
	},
	"Orbeetle-Gmax": {
		"id":   uint(826),
		"form": uint(1),
		"type": "bug",
	},
	"Nickit": {
		"id":   uint(827),
		"form": uint(0),
		"type": "dark",
	},
	"Thievul": {
		"id":   uint(828),
		"form": uint(0),
		"type": "dark",
	},
	"Gossifleur": {
		"id":   uint(829),
		"form": uint(0),
		"type": "grass",
	},
	"Eldegoss": {
		"id":   uint(830),
		"form": uint(0),
		"type": "grass",
	},
	"Wooloo": {
		"id":   uint(831),
		"form": uint(0),
		"type": "normal",
	},
	"Dubwool": {
		"id":   uint(832),
		"form": uint(0),
		"type": "normal",
	},
	"Chewtle": {
		"id":   uint(833),
		"form": uint(0),
		"type": "water",
	},
	"Drednaw": {
		"id":   uint(834),
		"form": uint(0),
		"type": "water",
	},
	"Drednaw-Gmax": {
		"id":   uint(834),
		"form": uint(1),
		"type": "water",
	},
	"Yamper": {
		"id":   uint(835),
		"form": uint(0),
		"type": "electric",
	},
	"Boltund": {
		"id":   uint(836),
		"form": uint(0),
		"type": "electric",
	},
	"Rolycoly": {
		"id":   uint(837),
		"form": uint(0),
		"type": "rock",
	},
	"Carkol": {
		"id":   uint(838),
		"form": uint(0),
		"type": "rock",
	},
	"Coalossal": {
		"id":   uint(839),
		"form": uint(0),
		"type": "rock",
	},
	"Coalossal-Gmax": {
		"id":   uint(839),
		"form": uint(1),
		"type": "rock",
	},
	"Applin": {
		"id":   uint(840),
		"form": uint(0),
		"type": "grass",
	},
	"Flapple": {
		"id":   uint(841),
		"form": uint(0),
		"type": "grass",
	},
	"Flapple-Gmax": {
		"id":   uint(841),
		"form": uint(1),
		"type": "grass",
	},
	"Appletun": {
		"id":   uint(842),
		"form": uint(0),
		"type": "grass",
	},
	"Appletun-Gmax": {
		"id":   uint(842),
		"form": uint(1),
		"type": "grass",
	},
	"Silicobra": {
		"id":   uint(843),
		"form": uint(0),
		"type": "ground",
	},
	"Sandaconda": {
		"id":   uint(844),
		"form": uint(0),
		"type": "ground",
	},
	"Sandaconda-Gmax": {
		"id":   uint(844),
		"form": uint(1),
		"type": "ground",
	},
	"Cramorant": {
		"id":   uint(845),
		"form": uint(0),
		"type": "flying",
	},
	"Cramorant-Gulping": {
		"id":   uint(845),
		"form": uint(1),
		"type": "flying",
	},
	"Cramorant-Gorging": {
		"id":   uint(845),
		"form": uint(2),
		"type": "flying",
	},
	"Arrokuda": {
		"id":   uint(846),
		"form": uint(0),
		"type": "water",
	},
	"Barraskewda": {
		"id":   uint(847),
		"form": uint(0),
		"type": "water",
	},
	"Toxel": {
		"id":   uint(848),
		"form": uint(0),
		"type": "electric",
	},
	"Toxtricity": {
		"id":   uint(849),
		"form": uint(0),
		"type": "electric",
	},
	"Toxtricity-Low-Key": {
		"id":   uint(849),
		"form": uint(1),
		"type": "electric",
	},
	"Toxtricity-Gmax": {
		"id":   uint(844),
		"form": uint(2),
		"type": "electric",
	},
	"Sizzlipede": {
		"id":   uint(850),
		"form": uint(0),
		"type": "fire",
	},
	"Centiskorch": {
		"id":   uint(851),
		"form": uint(0),
		"type": "fire",
	},
	"Centiskorch-Gmax": {
		"id":   uint(851),
		"form": uint(1),
		"type": "fire",
	},
	"Clobbopus": {
		"id":   uint(852),
		"form": uint(0),
		"type": "fighting",
	},
	"Grapploct": {
		"id":   uint(853),
		"form": uint(0),
		"type": "fighting",
	},
	"Sinistea": {
		"id":   uint(854),
		"form": uint(0),
		"type": "ghost",
	},
	"Polteageist": {
		"id":   uint(855),
		"form": uint(0),
		"type": "ghost",
	},
	"Hatenna": {
		"id":   uint(856),
		"form": uint(0),
		"type": "psychic",
	},
	"Hattrem": {
		"id":   uint(857),
		"form": uint(0),
		"type": "psychic",
	},
	"Hatterene": {
		"id":   uint(858),
		"form": uint(0),
		"type": "psychic",
	},
	"Hatterene-Gmax": {
		"id":   uint(858),
		"form": uint(1),
		"type": "psychic",
	},
	"Impidimp": {
		"id":   uint(859),
		"form": uint(0),
		"type": "dark",
	},
	"Morgrem": {
		"id":   uint(860),
		"form": uint(0),
		"type": "dark",
	},
	"Grimmsnarl": {
		"id":   uint(861),
		"form": uint(0),
		"type": "dark",
	},
	"Grimmsnarl-Gmax": {
		"id":   uint(861),
		"form": uint(1),
		"type": "dark",
	},
	"Obstagoon": {
		"id":   uint(862),
		"form": uint(0),
		"type": "dark",
	},
	"Perrserker": {
		"id":   uint(863),
		"form": uint(0),
		"type": "steel",
	},
	"Cursola": {
		"id":   uint(864),
		"form": uint(0),
		"type": "ghost",
	},
	"Sirfetch'd": {
		"id":   uint(865),
		"form": uint(0),
		"type": "fighting",
	},
	"Mr. Rime": {
		"id":   uint(866),
		"form": uint(0),
		"type": "ice",
	},
	"Runerigus": {
		"id":   uint(867),
		"form": uint(0),
		"type": "ground",
	},
	"Milcery": {
		"id":   uint(868),
		"form": uint(0),
		"type": "fairy",
	},
	"Alcremie": {
		"id":   uint(869),
		"form": uint(0),
		"type": "fairy",
	},
	"Alcremie-Gmax": {
		"id":   uint(869),
		"form": uint(1),
		"type": "fairy",
	},
	"Falinks": {
		"id":   uint(870),
		"form": uint(0),
		"type": "fighting",
	},
	"Pincurchin": {
		"id":   uint(871),
		"form": uint(0),
		"type": "electric",
	},
	"Snom": {
		"id":   uint(872),
		"form": uint(0),
		"type": "ice",
	},
	"Frosmoth": {
		"id":   uint(873),
		"form": uint(0),
		"type": "ice",
	},
	"Stonjourner": {
		"id":   uint(874),
		"form": uint(0),
		"type": "rock",
	},
	"Eiscue": {
		"id":   uint(875),
		"form": uint(0),
		"type": "ice",
	},
	"Eiscue-Noice": {
		"id":   uint(875),
		"form": uint(1),
		"type": "ice",
	},
	"Indeedee": {
		"id":   uint(876),
		"form": uint(0),
		"type": "psychic",
	},
	"Indeedee-F": {
		"id":   uint(876),
		"form": uint(1),
		"type": "psychic",
	},
	"Morpeko": {
		"id":   uint(877),
		"form": uint(0),
		"type": "electric",
	},
	"Morpeko-Hangry": {
		"id":   uint(877),
		"form": uint(1),
		"type": "electric",
	},
	"Cufant": {
		"id":   uint(878),
		"form": uint(0),
		"type": "steel",
	},
	"Copperajah": {
		"id":   uint(879),
		"form": uint(0),
		"type": "steel",
	},
	"Copperajah-Gmax": {
		"id":   uint(879),
		"form": uint(1),
		"type": "steel",
	},
	"Dracozolt": {
		"id":   uint(880),
		"form": uint(0),
		"type": "electric",
	},
	"Arctozolt": {
		"id":   uint(881),
		"form": uint(0),
		"type": "electric",
	},
	"Dracovish": {
		"id":   uint(882),
		"form": uint(0),
		"type": "water",
	},
	"Arctovish": {
		"id":   uint(883),
		"form": uint(0),
		"type": "water",
	},
	"Duraludon": {
		"id":   uint(884),
		"form": uint(0),
		"type": "steel",
	},
	"Duraludon-Gmax": {
		"id":   uint(884),
		"form": uint(1),
		"type": "steel",
	},
	"Dreepy": {
		"id":   uint(885),
		"form": uint(0),
		"type": "dragon",
	},
	"Drakloak": {
		"id":   uint(886),
		"form": uint(0),
		"type": "dragon",
	},
	"Dragapult": {
		"id":   uint(887),
		"form": uint(0),
		"type": "dragon",
	},
	"Zacian": {
		"id":   uint(888),
		"form": uint(0),
		"type": "fairy",
	},
	"Zacian-Crowned": {
		"id":   uint(888),
		"form": uint(1),
		"type": "fairy",
	},
	"Zamazenta": {
		"id":   uint(889),
		"form": uint(0),
		"type": "fighting",
	},
	"Zamazenta-Crowned": {
		"id":   uint(889),
		"form": uint(1),
		"type": "fighting",
	},
	"Eternatus": {
		"id":   uint(890),
		"form": uint(0),
		"type": "poison",
	},
	"Eternatus-Eternamax": {
		"id":   uint(890),
		"form": uint(1),
		"type": "poison",
	},
	"Kubfu": {
		"id":   uint(891),
		"form": uint(0),
		"type": "fighting",
	},
	"Urshifu": {
		"id":   uint(892),
		"form": uint(0),
		"type": "fighting",
	},
	"Urshifu-Rapid-Strike": {
		"id":   uint(892),
		"form": uint(1),
		"type": "fighting",
	},
	"Urshifu-Gmax": {
		"id":   uint(892),
		"form": uint(2),
		"type": "fighting",
	},
	"Urshifu-Rapid-Strike-Gmax": {
		"id":   uint(892),
		"form": uint(3),
		"type": "fighting",
	},
	"Zarude": {
		"id":   uint(893),
		"form": uint(0),
		"type": "dark",
	},
	"Regieleki": {
		"id":   uint(894),
		"form": uint(0),
		"type": "electric",
	},
	"Regidrago": {
		"id":   uint(895),
		"form": uint(0),
		"type": "dragon",
	},
	"Glastrier": {
		"id":   uint(896),
		"form": uint(0),
		"type": "ice",
	},
	"Spectrier": {
		"id":   uint(897),
		"form": uint(0),
		"type": "ghost",
	},
	"Calyrex": {
		"id":   uint(898),
		"form": uint(0),
		"type": "psychic",
	},
	"Calyrex-Ice": {
		"id":   uint(898),
		"form": uint(1),
		"type": "psychic",
	},
	"Calyrex-Shadow": {
		"id":   uint(898),
		"form": uint(2),
		"type": "psychic",
	},
	"Syclant": {
		"id":   uint(10001),
		"form": uint(0),
		"type": "ice",
	},
	"Revenankh": {
		"id":   uint(10002),
		"form": uint(0),
		"type": "ghost",
	},
	"Pyroak": {
		"id":   uint(10003),
		"form": uint(0),
		"type": "fire",
	},
	"Fidgit": {
		"id":   uint(10004),
		"form": uint(0),
		"type": "poison",
	},
	"Stratagem": {
		"id":   uint(10005),
		"form": uint(0),
		"type": "rock",
	},
	"Arghonaut": {
		"id":   uint(10006),
		"form": uint(0),
		"type": "water",
	},
	"Kitsunoh": {
		"id":   uint(10007),
		"form": uint(0),
		"type": "steel",
	},
	"Cyclohm": {
		"id":   uint(10008),
		"form": uint(0),
		"type": "electric",
	},
	"Colossoil": {
		"id":   uint(10009),
		"form": uint(0),
		"type": "dark",
	},
	"Krilowatt": {
		"id":   uint(10010),
		"form": uint(0),
		"type": "electric",
	},
	"Voodoom": {
		"id":   uint(10011),
		"form": uint(0),
		"type": "fighting",
	},
	"Tomohawk": {
		"id":   uint(10012),
		"form": uint(0),
		"type": "flying",
	},
	"Necturna": {
		"id":   uint(10013),
		"form": uint(0),
		"type": "grass",
	},
	"Mollux": {
		"id":   uint(10014),
		"form": uint(0),
		"type": "fire",
	},
	"Aurumoth": {
		"id":   uint(10015),
		"form": uint(0),
		"type": "bug",
	},
	"Malaconda": {
		"id":   uint(10016),
		"form": uint(0),
		"type": "dark",
	},
	"Cawmodore": {
		"id":   uint(10017),
		"form": uint(0),
		"type": "steel",
	},
	"Volkraken": {
		"id":   uint(10018),
		"form": uint(0),
		"type": "water",
	},
	"Plasmanta": {
		"id":   uint(10019),
		"form": uint(0),
		"type": "electric",
	},
	"Naviathan": {
		"id":   uint(10020),
		"form": uint(0),
		"type": "water",
	},
	"Crucibelle-Mega": {
		"id":   uint(10021),
		"form": uint(1),
		"type": "rock",
	},
	"Crucibelle": {
		"id":   uint(10021),
		"form": uint(0),
		"type": "rock",
	},
	"Kerfluffle": {
		"id":   uint(10022),
		"form": uint(0),
		"type": "fairy",
	},
	"Pajantom": {
		"id":   uint(10023),
		"form": uint(0),
		"type": "dragon",
	},
	"Jumabo": {
		"id":   uint(10024),
		"form": uint(0),
		"type": "grass",
	},
	"Caribolt": {
		"id":   uint(10025),
		"form": uint(0),
		"type": "grass",
	},
	"Smokomodo": {
		"id":   uint(10026),
		"form": uint(0),
		"type": "fire",
	},
	"Snaelstrom": {
		"id":   uint(10027),
		"form": uint(0),
		"type": "water",
	},
	"Equilibra": {
		"id":   uint(10028),
		"form": uint(0),
		"type": "steel",
	},
	"Astrolotl": {
		"id":   uint(10029),
		"form": uint(0),
		"type": "fire",
	},
	"Miasmaw": {
		"id":   uint(10030),
		"form": uint(0),
		"type": "bug",
	},
	"Maplage": {
		"id":   uint(42001),
		"form": uint(0),
		"type": "grass",
	},
	"Hazelnaut": {
		"id":   uint(42002),
		"form": uint(0),
		"type": "grass",
	},
	"Ginocchio": {
		"id":   uint(42003),
		"form": uint(0),
		"type": "grass",
	},
	"Magmata": {
		"id":   uint(42004),
		"form": uint(0),
		"type": "fire",
	},
	"Delava": {
		"id":   uint(42005),
		"form": uint(0),
		"type": "fire",
	},
	"Salezerker": {
		"id":   uint(42006),
		"form": uint(0),
		"type": "fire",
	},
	"Frudge": {
		"id":   uint(42007),
		"form": uint(0),
		"type": "water",
	},
	"Aquaos": {
		"id":   uint(42008),
		"form": uint(0),
		"type": "water",
	},
	"Lemirethun": {
		"id":   uint(42009),
		"form": uint(0),
		"type": "water",
	},
	"Jermin": {
		"id":   uint(42010),
		"form": uint(0),
		"type": "normal",
	},
	"Jermin-Swarm": {
		"id":   uint(42010),
		"form": uint(1),
		"type": "normal",
	},
	"Skuba": {
		"id":   uint(42011),
		"form": uint(0),
		"type": "water",
	},
	"Skuba-Anti": {
		"id":   uint(42011),
		"form": uint(1),
		"type": "water",
	},
	"Spiranha": {
		"id":   uint(42012),
		"form": uint(0),
		"type": "water",
	},
	"Selaghast": {
		"id":   uint(42013),
		"form": uint(0),
		"type": "water",
	},
	"Karkutlass": {
		"id":   uint(42014),
		"form": uint(0),
		"type": "water",
	},
	"Bandicute": {
		"id":   uint(42015),
		"form": uint(0),
		"type": "ground",
	},
	"Naughtycoot": {
		"id":   uint(42016),
		"form": uint(0),
		"type": "ground",
	},
	"Frijolero": {
		"id":   uint(42017),
		"form": uint(0),
		"type": "grass",
	},
	"Picantero": {
		"id":   uint(42018),
		"form": uint(0),
		"type": "grass",
	},
	"Sherifuego": {
		"id":   uint(42019),
		"form": uint(0),
		"type": "ghost",
	},
	"Magghost": {
		"id":   uint(42020),
		"form": uint(0),
		"type": "bug",
	},
	"Coccurn": {
		"id":   uint(42021),
		"form": uint(0),
		"type": "bug",
	},
	"Phamothom": {
		"id":   uint(42022),
		"form": uint(0),
		"type": "bug",
	},
	"Larvitty": {
		"id":   uint(42023),
		"form": uint(0),
		"type": "bug",
	},
	"Grumpoon": {
		"id":   uint(42024),
		"form": uint(0),
		"type": "bug",
	},
	"Konekoth": {
		"id":   uint(42025),
		"form": uint(0),
		"type": "bug",
	},
	"Hyduck": {
		"id":   uint(42026),
		"form": uint(0),
		"type": "psychic",
	},
	"Bluduck": {
		"id":   uint(42027),
		"form": uint(0),
		"type": "psychic",
	},
	"Platylics": {
		"id":   uint(42028),
		"form": uint(0),
		"type": "psychic",
	},
	"Hamstatic": {
		"id":   uint(42029),
		"form": uint(0),
		"type": "normal",
	},
	"Ampstar": {
		"id":   uint(42030),
		"form": uint(0),
		"type": "normal",
	},
	"Disbeary": {
		"id":   uint(42031),
		"form": uint(0),
		"type": "normal",
	},
	"Disbeary-Ebil": {
		"id":   uint(42031),
		"form": uint(1),
		"type": "normal",
	},
	"Fungnet": {
		"id":   uint(42032),
		"form": uint(0),
		"type": "grass",
	},
	"Floribel": {
		"id":   uint(42033),
		"form": uint(0),
		"type": "grass",
	},
	"Mogeria": {
		"id":   uint(42034),
		"form": uint(0),
		"type": "grass",
	},
	"Flarhea": {
		"id":   uint(42035),
		"form": uint(0),
		"type": "fire",
	},
	"Pedestone": {
		"id":   uint(42036),
		"form": uint(0),
		"type": "rock",
	},
	"Pillaia": {
		"id":   uint(42037),
		"form": uint(0),
		"type": "rock",
	},
	"Erochre": {
		"id":   uint(42038),
		"form": uint(0),
		"type": "rock",
	},
	"Canamaple": {
		"id":   uint(42039),
		"form": uint(0),
		"type": "grass",
	},
	"Frozweed": {
		"id":   uint(42040),
		"form": uint(0),
		"type": "grass",
	},
	"Arachnote": {
		"id":   uint(42041),
		"form": uint(0),
		"type": "bug",
	},
	"Turturret": {
		"id":   uint(42042),
		"form": uint(0),
		"type": "ground",
	},
	"Spycrab": {
		"id":   uint(42043),
		"form": uint(0),
		"type": "water",
	},
	"Weatherbane": {
		"id":   uint(42044),
		"form": uint(0),
		"type": "steel",
	},
	"Cyclonian": {
		"id":   uint(42045),
		"form": uint(0),
		"type": "flying",
	},
	"Simionach": {
		"id":   uint(42046),
		"form": uint(0),
		"type": "fighting",
	},
	"Simionach-Zen": {
		"id":   uint(42046),
		"form": uint(1),
		"type": "fighting",
	},
	"Wattitude": {
		"id":   uint(42047),
		"form": uint(0),
		"type": "electric",
	},
	"Korokami": {
		"id":   uint(42048),
		"form": uint(0),
		"type": "dark",
	},
	"Sabsute": {
		"id":   uint(42049),
		"form": uint(0),
		"type": "normal",
	},
	"Dollghost": {
		"id":   uint(42050),
		"form": uint(0),
		"type": "normal",
	},
	"Tarquail": {
		"id":   uint(42051),
		"form": uint(0),
		"type": "poison",
	},
	"Eccosmic": {
		"id":   uint(42052),
		"form": uint(0),
		"type": "water",
	},
	"Eccosmic-Magellanic": {
		"id":   uint(42052),
		"form": uint(1),
		"type": "water",
	},
	"Eccosmic-Sol": {
		"id":   uint(42052),
		"form": uint(2),
		"type": "water",
	},
	"Eccosmic-Starburst": {
		"id":   uint(42052),
		"form": uint(3),
		"type": "water",
	},
	"Eccosmic-Quasar": {
		"id":   uint(42052),
		"form": uint(4),
		"type": "water",
	},
	"Eccosmic-Andromeda": {
		"id":   uint(42052),
		"form": uint(5),
		"type": "water",
	},
	"Eccosmic-Nebula": {
		"id":   uint(42052),
		"form": uint(6),
		"type": "water",
	},
	"Eccosmic-Milky Way": {
		"id":   uint(42052),
		"form": uint(7),
		"type": "water",
	},
	"Bearedaze": {
		"id":   uint(42053),
		"form": uint(0),
		"type": "normal",
	},
	"Kumataro": {
		"id":   uint(42054),
		"form": uint(0),
		"type": "normal",
	},
	"Jarumite": {
		"id":   uint(42055),
		"form": uint(0),
		"type": "dragon",
	},
	"Serpotine": {
		"id":   uint(42056),
		"form": uint(0),
		"type": "dragon",
	},
	"Tsemani": {
		"id":   uint(42057),
		"form": uint(0),
		"type": "dragon",
	},
	"Tadpoison": {
		"id":   uint(42058),
		"form": uint(0),
		"type": "poison",
	},
	"Bretkhelm": {
		"id":   uint(42059),
		"form": uint(0),
		"type": "poison",
	},
	"Berfgoyle": {
		"id":   uint(42060),
		"form": uint(0),
		"type": "poison",
	},
	"Flubunny": {
		"id":   uint(42061),
		"form": uint(0),
		"type": "fairy",
	},
	"Lepooze": {
		"id":   uint(42062),
		"form": uint(0),
		"type": "fairy",
	},
	"Lunabbit": {
		"id":   uint(42063),
		"form": uint(0),
		"type": "fairy",
	},
	"Helmata": {
		"id":   uint(42064),
		"form": uint(0),
		"type": "ground",
	},
	"Punkuyu": {
		"id":   uint(42065),
		"form": uint(0),
		"type": "ground",
	},
	"Chamelee": {
		"id":   uint(42066),
		"form": uint(0),
		"type": "ground",
	},
	"Toastort": {
		"id":   uint(42067),
		"form": uint(0),
		"type": "steel",
	},
	"Horntuba": {
		"id":   uint(42068),
		"form": uint(0),
		"type": "steel",
	},
	"Cristanium": {
		"id":   uint(42069),
		"form": uint(0),
		"type": "steel",
	},
	"Cheesetah": {
		"id":   uint(42070),
		"form": uint(0),
		"type": "electric",
	},
	"Whipsicuffs": {
		"id":   uint(42071),
		"form": uint(0),
		"type": "bug",
	},
	"Tarantagon": {
		"id":   uint(42072),
		"form": uint(0),
		"type": "bug",
	},
	"Pegui": {
		"id":   uint(42073),
		"form": uint(0),
		"type": "ice",
	},
	"Pechit": {
		"id":   uint(42074),
		"form": uint(0),
		"type": "ice",
	},
	"Copolar": {
		"id":   uint(42075),
		"form": uint(0),
		"type": "ice",
	},
	"Yuukiino": {
		"id":   uint(42076),
		"form": uint(0),
		"type": "ice",
	},
	"Aquadiiva": {
		"id":   uint(42077),
		"form": uint(0),
		"type": "water",
	},
	"Monkite": {
		"id":   uint(42078),
		"form": uint(0),
		"type": "flying",
	},
	"Konglide": {
		"id":   uint(42079),
		"form": uint(0),
		"type": "flying",
	},
	"Yokite": {
		"id":   uint(42080),
		"form": uint(0),
		"type": "flying",
	},
	"Borkosmos": {
		"id":   uint(42081),
		"form": uint(0),
		"type": "flying",
	},
	"Skallaxy": {
		"id":   uint(42082),
		"form": uint(0),
		"type": "flying",
	},
	"Icehugger": {
		"id":   uint(42083),
		"form": uint(0),
		"type": "ice",
	},
	"Xeninter": {
		"id":   uint(42084),
		"form": uint(0),
		"type": "ice",
	},
	"Bubbasaur": {
		"id":   uint(42085),
		"form": uint(0),
		"type": "fairy",
	},
	"Bungaloon": {
		"id":   uint(42086),
		"form": uint(0),
		"type": "fairy",
	},
	"Envile": {
		"id":   uint(42087),
		"form": uint(0),
		"type": "dark",
	},
	"Kaustikrok": {
		"id":   uint(42088),
		"form": uint(0),
		"type": "dark",
	},
	"3dawg": {
		"id":   uint(42089),
		"form": uint(0),
		"type": "normal",
	},
	"Drattel": {
		"id":   uint(42090),
		"form": uint(0),
		"type": "dragon",
	},
	"Zillichina": {
		"id":   uint(42091),
		"form": uint(0),
		"type": "dragon",
	},
	"Venireal": {
		"id":   uint(42092),
		"form": uint(0),
		"type": "psychic",
	},
	"Apocamise": {
		"id":   uint(42093),
		"form": uint(0),
		"type": "psychic",
	},
	"Voliath": {
		"id":   uint(42094),
		"form": uint(0),
		"type": "dragon",
	},
	"Mafiadon": {
		"id":   uint(42095),
		"form": uint(0),
		"type": "dark",
	},
	"Wycern": {
		"id":   uint(42096),
		"form": uint(0),
		"type": "dragon",
	},
	"Bitteragon": {
		"id":   uint(42097),
		"form": uint(0),
		"type": "dragon",
	},
	"Trickling": {
		"id":   uint(42098),
		"form": uint(0),
		"type": "dark",
	},
	"Draklown": {
		"id":   uint(42099),
		"form": uint(0),
		"type": "dark",
	},
	"Moyalith": {
		"id":   uint(42100),
		"form": uint(0),
		"type": "rock",
	},
	"Dinomight": {
		"id":   uint(42101),
		"form": uint(0),
		"type": "rock",
	},
	"Cocken": {
		"id":   uint(42102),
		"form": uint(0),
		"type": "",
	},
	"Cocken-Mega": {
		"id":   uint(42102),
		"form": uint(1),
		"type": "",
	},
	"Primiteve": {
		"id":   uint(42103),
		"form": uint(0),
		"type": "rock",
	},
	"Tokoyaki": {
		"id":   uint(42104),
		"form": uint(0),
		"type": "water",
	},
	"Onsenpura": {
		"id":   uint(42105),
		"form": uint(0),
		"type": "rock",
	},
	"Molterra": {
		"id":   uint(42106),
		"form": uint(0),
		"type": "fire",
	},
	"Dubsnake": {
		"id":   uint(42107),
		"form": uint(0),
		"type": "ice",
	},
	"Hydroil": {
		"id":   uint(42108),
		"form": uint(0),
		"type": "water",
	},
	"Equuan": {
		"id":   uint(42109),
		"form": uint(0),
		"type": "normal",
	},
	"Impfection": {
		"id":   uint(42110),
		"form": uint(0),
		"type": "poison",
	},
	"Junkgeist": {
		"id":   uint(42111),
		"form": uint(0),
		"type": "poison",
	},
	"Krissy": {
		"id":   uint(42112),
		"form": uint(0),
		"type": "electric",
	},
	"Krackodemon": {
		"id":   uint(42113),
		"form": uint(0),
		"type": "electric",
	},
	"Strollge": {
		"id":   uint(42114),
		"form": uint(0),
		"type": "normal",
	},
	"Mamini": {
		"id":   uint(42115),
		"form": uint(0),
		"type": "ice",
	},
	"Mammount": {
		"id":   uint(42116),
		"form": uint(0),
		"type": "ice",
	},
	"Vriskeleton": {
		"id":   uint(42117),
		"form": uint(0),
		"type": "dark",
	},
	"Bulbfrog": {
		"id":   uint(42118),
		"form": uint(0),
		"type": "electric",
	},
	"Ballboa": {
		"id":   uint(42119),
		"form": uint(0),
		"type": "rock",
	},
	"Semmush": {
		"id":   uint(42120),
		"form": uint(0),
		"type": "water",
	},
	"Husbin": {
		"id":   uint(42121),
		"form": uint(0),
		"type": "fighting",
	},
	"Smogars": {
		"id":   uint(42122),
		"form": uint(0),
		"type": "poison",
	},
	"Gargarramer": {
		"id":   uint(42123),
		"form": uint(0),
		"type": "rock",
	},
	"Gargarramer-Awoken": {
		"id":   uint(42123),
		"form": uint(1),
		"type": "rock",
	},
	"Noxilium": {
		"id":   uint(42124),
		"form": uint(0),
		"type": "ghost",
	},
	"Fusjahl": {
		"id":   uint(42125),
		"form": uint(0),
		"type": "",
	},
	"Indignifly": {
		"id":   uint(42126),
		"form": uint(0),
		"type": "bug",
	},
	"Regishort": {
		"id":   uint(42127),
		"form": uint(0),
		"type": "electric",
	},
	"Regicide": {
		"id":   uint(42128),
		"form": uint(0),
		"type": "dragon",
	},
	"Regigigone": {
		"id":   uint(42129),
		"form": uint(0),
		"type": "normal",
	},
	"Glaciun": {
		"id":   uint(42130),
		"form": uint(0),
		"type": "ice",
	},
	"Galvadeux": {
		"id":   uint(42131),
		"form": uint(0),
		"type": "electric",
	},
	"Pyrotrois": {
		"id":   uint(42132),
		"form": uint(0),
		"type": "fire",
	},
	"Ranruu": {
		"id":   uint(42133),
		"form": uint(0),
		"type": "dark",
	},
	"Loituma": {
		"id":   uint(42134),
		"form": uint(0),
		"type": "grass",
	},
	"Ignifatu": {
		"id":   uint(42135),
		"form": uint(0),
		"type": "fire",
	},
	"Spenjbab": {
		"id":   uint(42136),
		"form": uint(0),
		"type": "grass",
	},
	"Gigapuddi": {
		"id":   uint(42137),
		"form": uint(0),
		"type": "fairy",
	},
	"Empidae": {
		"id":   uint(42138),
		"form": uint(0),
		"type": "normal",
	},
	"Phantorney": {
		"id":   uint(42139),
		"form": uint(0),
		"type": "ghost",
	},
	"Cirnumiru": {
		"id":   uint(42140),
		"form": uint(0),
		"type": "ice",
	},
	"Seraphill": {
		"id":   uint(42141),
		"form": uint(0),
		"type": "electric",
	},
	"Shiribiko": {
		"id":   uint(42142),
		"form": uint(0),
		"type": "psychic",
	},
	"Nyanonite": {
		"id":   uint(42143),
		"form": uint(0),
		"type": "fairy",
	},
	"Ankhira": {
		"id":   uint(42144),
		"form": uint(0),
		"type": "ground",
	},
	"Ankhpu": {
		"id":   uint(42145),
		"form": uint(0),
		"type": "ground",
	},
	"Ankhurah": {
		"id":   uint(42146),
		"form": uint(0),
		"type": "ground",
	},
	"Squithee": {
		"id":   uint(42147),
		"form": uint(0),
		"type": "steel",
	},
	"Goblazer": {
		"id":   uint(42148),
		"form": uint(0),
		"type": "steel",
	},
	"Warghork": {
		"id":   uint(42149),
		"form": uint(0),
		"type": "steel",
	},
	"Wawho": {
		"id":   uint(42150),
		"form": uint(0),
		"type": "",
	},
	"Sablini": {
		"id":   uint(42151),
		"form": uint(0),
		"type": "steel",
	},
	"Doomsday": {
		"id":   uint(42152),
		"form": uint(0),
		"type": "ghost",
	},
	"Peasol": {
		"id":   uint(42161),
		"form": uint(0),
		"type": "normal",
	},
	"Vaponobi": {
		"id":   uint(42162),
		"form": uint(0),
		"type": "normal",
	},
	"Haarshogun": {
		"id":   uint(42163),
		"form": uint(0),
		"type": "normal",
	},
	"Vinebomb": {
		"id":   uint(42164),
		"form": uint(0),
		"type": "grass",
	},
	"Fucker": {
		"id":   uint(42165),
		"form": uint(0),
		"type": "fighting",
	},
	"Fucker-Konata": {
		"id":   uint(42165),
		"form": uint(1),
		"type": "fighting",
	},
	"Emojinn": {
		"id":   uint(42166),
		"form": uint(0),
		"type": "psychic",
	},
	"Sporestool": {
		"id":   uint(42167),
		"form": uint(0),
		"type": "fairy",
	},
	"Peashroom": {
		"id":   uint(42168),
		"form": uint(0),
		"type": "fairy",
	},
	"Grasshole": {
		"id":   uint(69001),
		"form": uint(0),
		"type": "grass",
	},
	"Analgae": {
		"id":   uint(69002),
		"form": uint(0),
		"type": "grass",
	},
	"Rectreem": {
		"id":   uint(69003),
		"form": uint(0),
		"type": "grass",
	},
	"Arabomb": {
		"id":   uint(69004),
		"form": uint(0),
		"type": "fire",
	},
	"Iguallah": {
		"id":   uint(69005),
		"form": uint(0),
		"type": "fire",
	},
	"Lizakbar": {
		"id":   uint(69006),
		"form": uint(0),
		"type": "fire",
	},
	"Ejacasm": {
		"id":   uint(69007),
		"form": uint(0),
		"type": "water",
	},
	"Hosajack": {
		"id":   uint(69008),
		"form": uint(0),
		"type": "water",
	},
	"Condoom": {
		"id":   uint(69009),
		"form": uint(0),
		"type": "water",
	},
	"Squirrap": {
		"id":   uint(69010),
		"form": uint(0),
		"type": "normal",
	},
	"Gampster": {
		"id":   uint(69011),
		"form": uint(0),
		"type": "normal",
	},
	"Kengeon": {
		"id":   uint(69012),
		"form": uint(0),
		"type": "normal",
	},
	"Frieden": {
		"id":   uint(69013),
		"form": uint(0),
		"type": "normal",
	},
	"Corooster": {
		"id":   uint(69014),
		"form": uint(0),
		"type": "normal",
	},
	"Curicrawl": {
		"id":   uint(69015),
		"form": uint(0),
		"type": "bug",
	},
	"Gutsicoon": {
		"id":   uint(69016),
		"form": uint(0),
		"type": "bug",
	},
	"Gutsifly": {
		"id":   uint(69017),
		"form": uint(0),
		"type": "bug",
	},
	"Larvades": {
		"id":   uint(69018),
		"form": uint(0),
		"type": "bug",
	},
	"Proboskito": {
		"id":   uint(69019),
		"form": uint(0),
		"type": "bug",
	},
	"Clovour": {
		"id":   uint(69020),
		"form": uint(0),
		"type": "grass",
	},
	"Chanolour": {
		"id":   uint(69021),
		"form": uint(0),
		"type": "grass",
	},
	"Nauseon": {
		"id":   uint(69022),
		"form": uint(0),
		"type": "poison",
	},
	"Hazmate": {
		"id":   uint(69023),
		"form": uint(0),
		"type": "poison",
	},
	"Ebolable": {
		"id":   uint(69024),
		"form": uint(0),
		"type": "poison",
	},
	"Pikotton": {
		"id":   uint(69025),
		"form": uint(0),
		"type": "electric",
	},
	"Birdhouse": {
		"id":   uint(69026),
		"form": uint(0),
		"type": "normal",
	},
	"Lanshil": {
		"id":   uint(69027),
		"form": uint(0),
		"type": "rock",
	},
	"Middril": {
		"id":   uint(69028),
		"form": uint(0),
		"type": "rock",
	},
	"Shurismash": {
		"id":   uint(69029),
		"form": uint(0),
		"type": "rock",
	},
	"Maymay": {
		"id":   uint(69030),
		"form": uint(0),
		"type": "normal",
	},
	"Memenace": {
		"id":   uint(69031),
		"form": uint(0),
		"type": "normal",
	},
	"Vandash": {
		"id":   uint(69032),
		"form": uint(0),
		"type": "dark",
	},
	"Piguson": {
		"id":   uint(69033),
		"form": uint(0),
		"type": "fighting",
	},
	"Armando": {
		"id":   uint(69034),
		"form": uint(0),
		"type": "fighting",
	},
	"Muscledude": {
		"id":   uint(69035),
		"form": uint(0),
		"type": "fighting",
	},
	"Tripecs": {
		"id":   uint(69036),
		"form": uint(0),
		"type": "fighting",
	},
	"Semdrop": {
		"id":   uint(69037),
		"form": uint(0),
		"type": "normal",
	},
	"Semrust": {
		"id":   uint(69038),
		"form": uint(0),
		"type": "normal",
	},
	"Furnazi": {
		"id":   uint(69039),
		"form": uint(0),
		"type": "fire",
	},
	"Finasoven": {
		"id":   uint(69040),
		"form": uint(0),
		"type": "fire",
	},
	"Troubait": {
		"id":   uint(69041),
		"form": uint(0),
		"type": "water",
	},
	"Tulure": {
		"id":   uint(69042),
		"form": uint(0),
		"type": "water",
	},
	"Monstrap": {
		"id":   uint(69043),
		"form": uint(0),
		"type": "water",
	},
	"Pixila": {
		"id":   uint(69044),
		"form": uint(0),
		"type": "bug",
	},
	"Fairileon": {
		"id":   uint(69045),
		"form": uint(0),
		"type": "bug",
	},
	"Wedgard": {
		"id":   uint(69046),
		"form": uint(0),
		"type": "ground",
	},
	"Knokedge": {
		"id":   uint(69047),
		"form": uint(0),
		"type": "ground",
	},
	"Ribbizap": {
		"id":   uint(69048),
		"form": uint(0),
		"type": "electric",
	},
	"Elephas": {
		"id":   uint(69049),
		"form": uint(0),
		"type": "psychic",
	},
	"Seamapan": {
		"id":   uint(69050),
		"form": uint(0),
		"type": "water",
	},
	"Caroline": {
		"id":   uint(69051),
		"form": uint(0),
		"type": "ground",
	},
	"Chompest": {
		"id":   uint(69052),
		"form": uint(0),
		"type": "rock",
	},
	"Masdawg": {
		"id":   uint(69053),
		"form": uint(0),
		"type": "normal",
	},
	"Pasdawg": {
		"id":   uint(69054),
		"form": uint(0),
		"type": "normal",
	},
	"Spanke": {
		"id":   uint(69055),
		"form": uint(0),
		"type": "grass",
	},
	"Anaconduke": {
		"id":   uint(69056),
		"form": uint(0),
		"type": "grass",
	},
	"Fishnism": {
		"id":   uint(69057),
		"form": uint(0),
		"type": "water",
	},
	"Sjwhale": {
		"id":   uint(69058),
		"form": uint(0),
		"type": "water",
	},
	"Chezetta": {
		"id":   uint(69059),
		"form": uint(0),
		"type": "fairy",
	},
	"Mozzamazel": {
		"id":   uint(69060),
		"form": uint(0),
		"type": "fairy",
	},
	"Krokling": {
		"id":   uint(69061),
		"form": uint(0),
		"type": "ground",
	},
	"Krokizon": {
		"id":   uint(69062),
		"form": uint(0),
		"type": "ground",
	},
	"Conchilla": {
		"id":   uint(69063),
		"form": uint(0),
		"type": "normal",
	},
	"Haremit": {
		"id":   uint(69064),
		"form": uint(0),
		"type": "normal",
	},
	"Coolcube": {
		"id":   uint(69065),
		"form": uint(0),
		"type": "ice",
	},
	"Tankube": {
		"id":   uint(69066),
		"form": uint(0),
		"type": "ice",
	},
	"Fabkube": {
		"id":   uint(69067),
		"form": uint(0),
		"type": "ice",
	},
	"Smogaroben": {
		"id":   uint(69068),
		"form": uint(0),
		"type": "poison",
	},
	"Smoxilon": {
		"id":   uint(69069),
		"form": uint(0),
		"type": "poison",
	},
	"Machmona": {
		"id":   uint(69070),
		"form": uint(0),
		"type": "fighting",
	},
	"Bacub": {
		"id":   uint(69071),
		"form": uint(0),
		"type": "normal",
	},
	"Urswine": {
		"id":   uint(69072),
		"form": uint(0),
		"type": "normal",
	},
	"Moostatic": {
		"id":   uint(69073),
		"form": uint(0),
		"type": "electric",
	},
	"Mooshock": {
		"id":   uint(69074),
		"form": uint(0),
		"type": "electric",
	},
	"Pretzely": {
		"id":   uint(69075),
		"form": uint(0),
		"type": "rock",
	},
	"Deemdow": {
		"id":   uint(69076),
		"form": uint(0),
		"type": "ghost",
	},
	"Dowster": {
		"id":   uint(69077),
		"form": uint(0),
		"type": "ghost",
	},
	"Cheerly": {
		"id":   uint(69078),
		"form": uint(0),
		"type": "fairy",
	},
	"Cheerific": {
		"id":   uint(69079),
		"form": uint(0),
		"type": "fairy",
	},
	"Mennopaws": {
		"id":   uint(69080),
		"form": uint(0),
		"type": "fairy",
	},
	"Caracold": {
		"id":   uint(69081),
		"form": uint(0),
		"type": "ice",
	},
	"Glacialynx": {
		"id":   uint(69082),
		"form": uint(0),
		"type": "ice",
	},
	"Pengas": {
		"id":   uint(69083),
		"form": uint(0),
		"type": "ice",
	},
	"Praeteur": {
		"id":   uint(69084),
		"form": uint(0),
		"type": "water",
	},
	"Praestish": {
		"id":   uint(69085),
		"form": uint(0),
		"type": "water",
	},
	"Flowre": {
		"id":   uint(69086),
		"form": uint(0),
		"type": "poison",
	},
	"Florious": {
		"id":   uint(69087),
		"form": uint(0),
		"type": "poison",
	},
	"Juarecito": {
		"id":   uint(69088),
		"form": uint(0),
		"type": "ground",
	},
	"Ponchito": {
		"id":   uint(69089),
		"form": uint(0),
		"type": "ground",
	},
	"Somboludo": {
		"id":   uint(69090),
		"form": uint(0),
		"type": "ground",
	},
	"Puppacti": {
		"id":   uint(69091),
		"form": uint(0),
		"type": "grass",
	},
	"Marionettl": {
		"id":   uint(69092),
		"form": uint(0),
		"type": "grass",
	},
	"Tittai": {
		"id":   uint(69093),
		"form": uint(0),
		"type": "water",
	},
	"Octai": {
		"id":   uint(69094),
		"form": uint(0),
		"type": "water",
	},
	"Charagon": {
		"id":   uint(69095),
		"form": uint(0),
		"type": "steel",
	},
	"Hyletrack": {
		"id":   uint(69096),
		"form": uint(0),
		"type": "steel",
	},
	"Euphoreal": {
		"id":   uint(69097),
		"form": uint(0),
		"type": "dark",
	},
	"Armowite": {
		"id":   uint(69098),
		"form": uint(0),
		"type": "dark",
	},
	"Praunch": {
		"id":   uint(69099),
		"form": uint(0),
		"type": "water",
	},
	"Karawn": {
		"id":   uint(69100),
		"form": uint(0),
		"type": "water",
	},
	"Mabster": {
		"id":   uint(69101),
		"form": uint(0),
		"type": "water",
	},
	"Tricient": {
		"id":   uint(69102),
		"form": uint(0),
		"type": "rock",
	},
	"Tricillion": {
		"id":   uint(69103),
		"form": uint(0),
		"type": "rock",
	},
	"Poizookie": {
		"id":   uint(69104),
		"form": uint(0),
		"type": "rock",
	},
	"Toxiraptor": {
		"id":   uint(69105),
		"form": uint(0),
		"type": "rock",
	},
	"Sabreck": {
		"id":   uint(69106),
		"form": uint(0),
		"type": "rock",
	},
	"Huntabre": {
		"id":   uint(69107),
		"form": uint(0),
		"type": "rock",
	},
	"Pyralink": {
		"id":   uint(69108),
		"form": uint(0),
		"type": "dragon",
	},
	"Singlets": {
		"id":   uint(69109),
		"form": uint(0),
		"type": "psychic",
	},
	"Dubus": {
		"id":   uint(69110),
		"form": uint(0),
		"type": "psychic",
	},
	"Tripsius": {
		"id":   uint(69111),
		"form": uint(0),
		"type": "psychic",
	},
	"Miasmer": {
		"id":   uint(69112),
		"form": uint(0),
		"type": "poison",
	},
	"Miasmortor": {
		"id":   uint(69113),
		"form": uint(0),
		"type": "poison",
	},
	"Cacademon": {
		"id":   uint(69114),
		"form": uint(0),
		"type": "ground",
	},
	"Hanginy": {
		"id":   uint(69115),
		"form": uint(0),
		"type": "ghost",
	},
	"Chancer": {
		"id":   uint(69116),
		"form": uint(0),
		"type": "ghost",
	},
	"Unjoy": {
		"id":   uint(69117),
		"form": uint(0),
		"type": "ghost",
	},
	"Frosowl": {
		"id":   uint(69118),
		"form": uint(0),
		"type": "ice",
	},
	"Whizzard": {
		"id":   uint(69119),
		"form": uint(0),
		"type": "ice",
	},
	"Drapillar": {
		"id":   uint(69120),
		"form": uint(0),
		"type": "bug",
	},
	"Caparagon": {
		"id":   uint(69121),
		"form": uint(0),
		"type": "bug",
	},
	"Mantrake": {
		"id":   uint(69122),
		"form": uint(0),
		"type": "bug",
	},
	"Ogrelord": {
		"id":   uint(69123),
		"form": uint(0),
		"type": "dark",
	},
	"Chantruth": {
		"id":   uint(69124),
		"form": uint(0),
		"type": "dark",
	},
	"Inbitween": {
		"id":   uint(69125),
		"form": uint(0),
		"type": "fairy",
	},
	"Geigh": {
		"id":   uint(69126),
		"form": uint(0),
		"type": "fairy",
	},
	"Jarape": {
		"id":   uint(69127),
		"form": uint(0),
		"type": "electric",
	},
	"Spilefree": {
		"id":   uint(69128),
		"form": uint(0),
		"type": "electric",
	},
	"Bongecko": {
		"id":   uint(69129),
		"form": uint(0),
		"type": "grass",
	},
	"Mariguana": {
		"id":   uint(69130),
		"form": uint(0),
		"type": "grass",
	},
	"Marleyzard": {
		"id":   uint(69131),
		"form": uint(0),
		"type": "grass",
	},
	"Blobbos": {
		"id":   uint(69132),
		"form": uint(0),
		"type": "ice",
	},
	"Ayylamo": {
		"id":   uint(69133),
		"form": uint(0),
		"type": "psychic",
	},
	"Smellox": {
		"id":   uint(69134),
		"form": uint(0),
		"type": "poison",
	},
	"Chasumo": {
		"id":   uint(69135),
		"form": uint(0),
		"type": "fighting",
	},
	"Oilslam": {
		"id":   uint(69136),
		"form": uint(0),
		"type": "poison",
	},
	"Isissin": {
		"id":   uint(69137),
		"form": uint(0),
		"type": "poison",
	},
	"Stuffowl": {
		"id":   uint(69138),
		"form": uint(0),
		"type": "normal",
	},
	"Durkey": {
		"id":   uint(69139),
		"form": uint(0),
		"type": "normal",
	},
	"Upchucken": {
		"id":   uint(69140),
		"form": uint(0),
		"type": "normal",
	},
	"Dragun": {
		"id":   uint(69141),
		"form": uint(0),
		"type": "dragon",
	},
	"Triggedon": {
		"id":   uint(69142),
		"form": uint(0),
		"type": "dragon",
	},
	"Cannonance": {
		"id":   uint(69143),
		"form": uint(0),
		"type": "dragon",
	},
	"Evacycle": {
		"id":   uint(69144),
		"form": uint(0),
		"type": "steel",
	},
	"Foryu": {
		"id":   uint(69145),
		"form": uint(0),
		"type": "steel",
	},
	"Scytill": {
		"id":   uint(69146),
		"form": uint(0),
		"type": "steel",
	},
	"Baddon": {
		"id":   uint(69147),
		"form": uint(0),
		"type": "dragon",
	},
	"Endranther": {
		"id":   uint(69148),
		"form": uint(0),
		"type": "poison",
	},
	"Notridley": {
		"id":   uint(69149),
		"form": uint(0),
		"type": "dragon",
	},
	"Demiwaifu": {
		"id":   uint(69150),
		"form": uint(0),
		"type": "fairy",
	},
	"Clovenix": {
		"id":   uint(69151),
		"form": uint(0),
		"type": "grass",
	},
	"Reptyke": {
		"id":   uint(69152),
		"form": uint(0),
		"type": "fighting",
	},
	"Rasclobber": {
		"id":   uint(69153),
		"form": uint(0),
		"type": "fighting",
	},
	"Batterex": {
		"id":   uint(69154),
		"form": uint(0),
		"type": "fighting",
	},
	"Faeriock": {
		"id":   uint(69155),
		"form": uint(0),
		"type": "rock",
	},
	"Sylvuin": {
		"id":   uint(69156),
		"form": uint(0),
		"type": "rock",
	},
	"Oburonyxo": {
		"id":   uint(69157),
		"form": uint(0),
		"type": "rock",
	},
	"Cloucat": {
		"id":   uint(69158),
		"form": uint(0),
		"type": "flying",
	},
	"Felimbus": {
		"id":   uint(69159),
		"form": uint(0),
		"type": "flying",
	},
	"Deathorus": {
		"id":   uint(69160),
		"form": uint(0),
		"type": "flying",
	},
	"Knogoat": {
		"id":   uint(69161),
		"form": uint(0),
		"type": "normal",
	},
	"Knogrinyu": {
		"id":   uint(69162),
		"form": uint(0),
		"type": "normal",
	},
	"Glasshot": {
		"id":   uint(69163),
		"form": uint(0),
		"type": "ice",
	},
	"Glassannon": {
		"id":   uint(69164),
		"form": uint(0),
		"type": "ice",
	},
	"Fluffyeti": {
		"id":   uint(69165),
		"form": uint(0),
		"type": "ice",
	},
	"Yedoom": {
		"id":   uint(69166),
		"form": uint(0),
		"type": "ice",
	},
	"Mirrostine": {
		"id":   uint(69167),
		"form": uint(0),
		"type": "steel",
	},
	"Mudpants": {
		"id":   uint(69168),
		"form": uint(0),
		"type": "ground",
	},
	"Mudslacks": {
		"id":   uint(69169),
		"form": uint(0),
		"type": "ground",
	},
	"Bukitlee": {
		"id":   uint(69170),
		"form": uint(0),
		"type": "steel",
	},
	"Lossolith": {
		"id":   uint(69171),
		"form": uint(0),
		"type": "steel",
	},
	"Mimimie": {
		"id":   uint(69172),
		"form": uint(0),
		"type": "water",
	},
	"Sprucifix": {
		"id":   uint(69173),
		"form": uint(0),
		"type": "grass",
	},
	"Hulkan": {
		"id":   uint(69174),
		"form": uint(0),
		"type": "fighting",
	},
	"Nonite": {
		"id":   uint(69175),
		"form": uint(0),
		"type": "fairy",
	},
	"Nonegative": {
		"id":   uint(69176),
		"form": uint(0),
		"type": "fairy",
	},
	"Nomaestro": {
		"id":   uint(69177),
		"form": uint(0),
		"type": "fairy",
	},
	"Dunkypea": {
		"id":   uint(69178),
		"form": uint(0),
		"type": "poison",
	},
	"Burdmen": {
		"id":   uint(69179),
		"form": uint(0),
		"type": "poison",
	},
	"Nutzboltz": {
		"id":   uint(69180),
		"form": uint(0),
		"type": "poison",
	},
	"Kuklux": {
		"id":   uint(69181),
		"form": uint(0),
		"type": "fire",
	},
	"Kuklan": {
		"id":   uint(69182),
		"form": uint(0),
		"type": "fire",
	},
	"Flameboyan": {
		"id":   uint(69183),
		"form": uint(0),
		"type": "fire",
	},
	"Piicee": {
		"id":   uint(69184),
		"form": uint(0),
		"type": "psychic",
	},
	"Sonnanos": {
		"id":   uint(69185),
		"form": uint(0),
		"type": "steel",
	},
	"Plebbles": {
		"id":   uint(69186),
		"form": uint(0),
		"type": "rock",
	},
	"Sapleaf": {
		"id":   uint(69187),
		"form": uint(0),
		"type": "grass",
	},
	"Ignut": {
		"id":   uint(69188),
		"form": uint(0),
		"type": "fire",
	},
	"Phantash": {
		"id":   uint(69189),
		"form": uint(0),
		"type": "ghost",
	},
	"Polossus": {
		"id":   uint(69190),
		"form": uint(0),
		"type": "poison",
	},
	"Hitmonana": {
		"id":   uint(69191),
		"form": uint(0),
		"type": "grass",
	},
	"Ballacle": {
		"id":   uint(69192),
		"form": uint(0),
		"type": "rock",
	},
	"Barbarkley": {
		"id":   uint(69193),
		"form": uint(0),
		"type": "rock",
	},
	"Kekroach": {
		"id":   uint(69194),
		"form": uint(0),
		"type": "dark",
	},
	"Rekroach": {
		"id":   uint(69195),
		"form": uint(0),
		"type": "dark",
	},
	"Ricosheep": {
		"id":   uint(69196),
		"form": uint(0),
		"type": "normal",
	},
	"Ricosuave": {
		"id":   uint(69197),
		"form": uint(0),
		"type": "normal",
	},
	"Spaghefant": {
		"id":   uint(69198),
		"form": uint(0),
		"type": "ground",
	},
	"Crystaquil": {
		"id":   uint(69199),
		"form": uint(0),
		"type": "ice",
	},
	"Crystaking": {
		"id":   uint(69200),
		"form": uint(0),
		"type": "ice",
	},
	"Fontaba": {
		"id":   uint(69201),
		"form": uint(0),
		"type": "psychic",
	},
	"Floconut": {
		"id":   uint(69202),
		"form": uint(0),
		"type": "grass",
	},
	"Sappalm": {
		"id":   uint(69203),
		"form": uint(0),
		"type": "grass",
	},
	"Roynapalm": {
		"id":   uint(69204),
		"form": uint(0),
		"type": "grass",
	},
	"Manateet": {
		"id":   uint(69205),
		"form": uint(0),
		"type": "water",
	},
	"Manatorque": {
		"id":   uint(69206),
		"form": uint(0),
		"type": "water",
	},
	"Manatank": {
		"id":   uint(69207),
		"form": uint(0),
		"type": "water",
	},
	"Anonymouse": {
		"id":   uint(69208),
		"form": uint(0),
		"type": "dark",
	},
	"Motherfuck": {
		"id":   uint(69209),
		"form": uint(0),
		"type": "ground",
	},
	"Ninoop": {
		"id":   uint(69210),
		"form": uint(0),
		"type": "normal",
	},
	"Draconius": {
		"id":   uint(69211),
		"form": uint(0),
		"type": "normal",
	},
	"Ticktac": {
		"id":   uint(69212),
		"form": uint(0),
		"type": "bug",
	},
	"Eareticle": {
		"id":   uint(69213),
		"form": uint(0),
		"type": "bug",
	},
	"Scythear": {
		"id":   uint(69214),
		"form": uint(0),
		"type": "bug",
	},
	"Cozload": {
		"id":   uint(69215),
		"form": uint(0),
		"type": "electric",
	},
	"Snugware": {
		"id":   uint(69216),
		"form": uint(0),
		"type": "electric",
	},
	"Fleespecs": {
		"id":   uint(69217),
		"form": uint(0),
		"type": "electric",
	},
	"Honrade": {
		"id":   uint(69218),
		"form": uint(0),
		"type": "fairy",
	},
	"Embortion": {
		"id":   uint(69219),
		"form": uint(0),
		"type": "dark",
	},
	"Premantom": {
		"id":   uint(69220),
		"form": uint(0),
		"type": "dark",
	},
	"Galavik": {
		"id":   uint(69221),
		"form": uint(0),
		"type": "steel",
	},
	"Galavire": {
		"id":   uint(69222),
		"form": uint(0),
		"type": "steel",
	},
	"Galaviste": {
		"id":   uint(69223),
		"form": uint(0),
		"type": "steel",
	},
	"Pollefin": {
		"id":   uint(69224),
		"form": uint(0),
		"type": "water",
	},
	"Florigrace": {
		"id":   uint(69225),
		"form": uint(0),
		"type": "water",
	},
	"Lasslee": {
		"id":   uint(69226),
		"form": uint(0),
		"type": "psychic",
	},
	"Wifemin": {
		"id":   uint(69227),
		"form": uint(0),
		"type": "psychic",
	},
	"Uwotto": {
		"id":   uint(69228),
		"form": uint(0),
		"type": "psychic",
	},
	"Venowatt": {
		"id":   uint(69229),
		"form": uint(0),
		"type": "electric",
	},
	"Vaultevour": {
		"id":   uint(69230),
		"form": uint(0),
		"type": "steel",
	},
	"Semdemen": {
		"id":   uint(69231),
		"form": uint(0),
		"type": "water",
	},
	"Dragking": {
		"id":   uint(69232),
		"form": uint(0),
		"type": "fairy",
	},
	"Kuuroba": {
		"id":   uint(69233),
		"form": uint(0),
		"type": "grass",
	},
	"Baitmaster": {
		"id":   uint(69234),
		"form": uint(0),
		"type": "water",
	},
	"Spookster": {
		"id":   uint(69235),
		"form": uint(0),
		"type": "ghost",
	},
	"Spookeer": {
		"id":   uint(69236),
		"form": uint(0),
		"type": "ghost",
	},
	"Spookzilla": {
		"id":   uint(69237),
		"form": uint(0),
		"type": "ghost",
	},
	"Derpato": {
		"id":   uint(69238),
		"form": uint(0),
		"type": "ground",
	},
	"Retater": {
		"id":   uint(69239),
		"form": uint(0),
		"type": "ground",
	},
	"Potarded": {
		"id":   uint(69240),
		"form": uint(0),
		"type": "ground",
	},
	"Sadfish": {
		"id":   uint(69241),
		"form": uint(0),
		"type": "water",
	},
	"Deloris": {
		"id":   uint(69242),
		"form": uint(0),
		"type": "grass",
	},
	"Oreon": {
		"id":   uint(69243),
		"form": uint(0),
		"type": "fairy",
	},
	"Farfigtron": {
		"id":   uint(69244),
		"form": uint(0),
		"type": "fire",
	},
	"Regirode": {
		"id":   uint(69245),
		"form": uint(0),
		"type": "rock",
	},
	"Regimelt": {
		"id":   uint(69246),
		"form": uint(0),
		"type": "ice",
	},
	"Regirust": {
		"id":   uint(69247),
		"form": uint(0),
		"type": "steel",
	},
	"Jewipede": {
		"id":   uint(69248),
		"form": uint(0),
		"type": "bug",
	},
	"Vivaiger": {
		"id":   uint(69249),
		"form": uint(0),
		"type": "electric",
	},
	"Heliofug": {
		"id":   uint(69250),
		"form": uint(0),
		"type": "dragon",
	},
	"Adesign": {
		"id":   uint(69251),
		"form": uint(0),
		"type": "???",
	},
	"Sadrog": {
		"id":   uint(69252),
		"form": uint(0),
		"type": "grass",
	},
	"Smugrok": {
		"id":   uint(69253),
		"form": uint(0),
		"type": "grass",
	},
	"Peperee": {
		"id":   uint(69254),
		"form": uint(0),
		"type": "grass",
	},
	"Hodtog": {
		"id":   uint(69255),
		"form": uint(0),
		"type": "fire",
	},
	"Wienebark": {
		"id":   uint(69256),
		"form": uint(0),
		"type": "fire",
	},
	"Saudoge": {
		"id":   uint(69257),
		"form": uint(0),
		"type": "fire",
	},
	"Gnarrk": {
		"id":   uint(69258),
		"form": uint(0),
		"type": "water",
	},
	"Corsharrk": {
		"id":   uint(69259),
		"form": uint(0),
		"type": "water",
	},
	"Davyjaws": {
		"id":   uint(69260),
		"form": uint(0),
		"type": "water",
	},
	"Humpunny": {
		"id":   uint(69261),
		"form": uint(0),
		"type": "normal",
	},
	"Bunnorgy": {
		"id":   uint(69262),
		"form": uint(0),
		"type": "normal",
	},
	"Emelgy": {
		"id":   uint(69263),
		"form": uint(0),
		"type": "normal",
	},
	"Fedorawk": {
		"id":   uint(69264),
		"form": uint(0),
		"type": "normal",
	},
	"FaZeagle": {
		"id":   uint(69265),
		"form": uint(0),
		"type": "fire",
	},
	"Catikillar": {
		"id":   uint(69266),
		"form": uint(0),
		"type": "bug",
	},
	"Tikoon": {
		"id":   uint(69267),
		"form": uint(0),
		"type": "bug",
	},
	"Oogabuga": {
		"id":   uint(69268),
		"form": uint(0),
		"type": "bug",
	},
	"Ebining": {
		"id":   uint(69269),
		"form": uint(0),
		"type": "dark",
	},
	"Emplyin": {
		"id":   uint(69270),
		"form": uint(0),
		"type": "dark",
	},
	"Upbote": {
		"id":   uint(69271),
		"form": uint(0),
		"type": "psychic",
	},
	"Upbeddit": {
		"id":   uint(69272),
		"form": uint(0),
		"type": "psychic",
	},
	"Flipbird": {
		"id":   uint(69273),
		"form": uint(0),
		"type": "normal",
	},
	"Wheygle": {
		"id":   uint(69274),
		"form": uint(0),
		"type": "fighting",
	},
	"Gangnome": {
		"id":   uint(69275),
		"form": uint(0),
		"type": "fairy",
	},
	"Pyongnome": {
		"id":   uint(69276),
		"form": uint(0),
		"type": "steel",
	},
	"Bruhkid": {
		"id":   uint(69277),
		"form": uint(0),
		"type": "ground",
	},
	"Ubruh": {
		"id":   uint(69278),
		"form": uint(0),
		"type": "ground",
	},
	"Walruse": {
		"id":   uint(69279),
		"form": uint(0),
		"type": "water",
	},
	"Walruskie": {
		"id":   uint(69280),
		"form": uint(0),
		"type": "water",
	},
	"Diobat": {
		"id":   uint(69281),
		"form": uint(0),
		"type": "dark",
	},
	"Warudio": {
		"id":   uint(69282),
		"form": uint(0),
		"type": "dark",
	},
	"Docee": {
		"id":   uint(69283),
		"form": uint(0),
		"type": "steel",
	},
	"Donutsteel": {
		"id":   uint(69284),
		"form": uint(0),
		"type": "steel",
	},
	"Cakupple": {
		"id":   uint(69285),
		"form": uint(0),
		"type": "grass",
	},
	"Fondupple": {
		"id":   uint(69286),
		"form": uint(0),
		"type": "grass",
	},
	"Baloofang": {
		"id":   uint(69287),
		"form": uint(0),
		"type": "ghost",
	},
	"Socckat": {
		"id":   uint(69288),
		"form": uint(0),
		"type": "normal",
	},
	"Egylamp": {
		"id":   uint(69289),
		"form": uint(0),
		"type": "psychic",
	},
	"Ruselamp": {
		"id":   uint(69290),
		"form": uint(0),
		"type": "psychic",
	},
	"Mehndior": {
		"id":   uint(69291),
		"form": uint(0),
		"type": "fire",
	},
	"Tunakking": {
		"id":   uint(69292),
		"form": uint(0),
		"type": "fire",
	},
	"Missletoe": {
		"id":   uint(69293),
		"form": uint(0),
		"type": "ice",
	},
	"Hohohoming": {
		"id":   uint(69294),
		"form": uint(0),
		"type": "ice",
	},
	"Flaa": {
		"id":   uint(69295),
		"form": uint(0),
		"type": "bug",
	},
	"Hopault": {
		"id":   uint(69296),
		"form": uint(0),
		"type": "bug",
	},
	"Araketsu": {
		"id":   uint(69297),
		"form": uint(0),
		"type": "bug",
	},
	"Senketula": {
		"id":   uint(69298),
		"form": uint(0),
		"type": "bug",
	},
	"Falcaptain": {
		"id":   uint(69299),
		"form": uint(0),
		"type": "fighting",
	},
	"Cirpent": {
		"id":   uint(69300),
		"form": uint(0),
		"type": "poison",
	},
	"Traumobra": {
		"id":   uint(69301),
		"form": uint(0),
		"type": "poison",
	},
	"Lizascoop": {
		"id":   uint(69302),
		"form": uint(0),
		"type": "ground",
	},
	"Reptrill": {
		"id":   uint(69303),
		"form": uint(0),
		"type": "ground",
	},
	"Fefeion": {
		"id":   uint(69304),
		"form": uint(0),
		"type": "bug",
	},
	"Feferun": {
		"id":   uint(69305),
		"form": uint(0),
		"type": "bug",
	},
	"Jerkle": {
		"id":   uint(69306),
		"form": uint(0),
		"type": "rock",
	},
	"Faptime": {
		"id":   uint(69307),
		"form": uint(0),
		"type": "rock",
	},
	"Shiriman": {
		"id":   uint(69308),
		"form": uint(0),
		"type": "psychic",
	},
	"Fresrye": {
		"id":   uint(69309),
		"form": uint(0),
		"type": "grass",
	},
	"Dedwheat": {
		"id":   uint(69310),
		"form": uint(0),
		"type": "grass",
	},
	"Nutjobber": {
		"id":   uint(69311),
		"form": uint(0),
		"type": "ghost",
	},
	"Deisnutz": {
		"id":   uint(69312),
		"form": uint(0),
		"type": "ghost",
	},
	"Spargle": {
		"id":   uint(69313),
		"form": uint(0),
		"type": "fairy",
	},
	"Spedo": {
		"id":   uint(69314),
		"form": uint(0),
		"type": "fairy",
	},
	"Housant": {
		"id":   uint(69315),
		"form": uint(0),
		"type": "bug",
	},
	"Apartmant": {
		"id":   uint(69316),
		"form": uint(0),
		"type": "bug",
	},
	"Fuermiga": {
		"id":   uint(69317),
		"form": uint(0),
		"type": "bug",
	},
	"Maripyro": {
		"id":   uint(69318),
		"form": uint(0),
		"type": "bug",
	},
	"Confirmabi": {
		"id":   uint(69319),
		"form": uint(0),
		"type": "grass",
	},
	"Bulbapedo": {
		"id":   uint(69320),
		"form": uint(0),
		"type": "grass",
	},
	"Wastenaut": {
		"id":   uint(69321),
		"form": uint(0),
		"type": "steel",
	},
	"Wantnaut": {
		"id":   uint(69322),
		"form": uint(0),
		"type": "steel",
	},
	"Cansumor": {
		"id":   uint(69323),
		"form": uint(0),
		"type": "poison",
	},
	"Snuffant": {
		"id":   uint(69324),
		"form": uint(0),
		"type": "ice",
	},
	"Cokemmoth": {
		"id":   uint(69325),
		"form": uint(0),
		"type": "ice",
	},
	"Grimdak": {
		"id":   uint(69326),
		"form": uint(0),
		"type": "dark",
	},
	"Shroofle": {
		"id":   uint(69327),
		"form": uint(0),
		"type": "dark",
	},
	"Betacluck": {
		"id":   uint(69328),
		"form": uint(0),
		"type": "electric",
	},
	"Thundacock": {
		"id":   uint(69329),
		"form": uint(0),
		"type": "electric",
	},
	"Twinfowl": {
		"id":   uint(69330),
		"form": uint(0),
		"type": "steel",
	},
	"Avianjoin": {
		"id":   uint(69331),
		"form": uint(0),
		"type": "steel",
	},
	"Illuminowl": {
		"id":   uint(69332),
		"form": uint(0),
		"type": "psychic",
	},
	"Illumatrix": {
		"id":   uint(69333),
		"form": uint(0),
		"type": "psychic",
	},
	"Typobop": {
		"id":   uint(69334),
		"form": uint(0),
		"type": "rock",
	},
	"Yeerex": {
		"id":   uint(69335),
		"form": uint(0),
		"type": "rock",
	},
	"Randomix": {
		"id":   uint(69336),
		"form": uint(0),
		"type": "rock",
	},
	"Cthullord": {
		"id":   uint(69337),
		"form": uint(0),
		"type": "rock",
	},
	"Quiboom": {
		"id":   uint(69338),
		"form": uint(0),
		"type": "water",
	},
	"Gynuke": {
		"id":   uint(69339),
		"form": uint(0),
		"type": "water",
	},
	"Rainglock": {
		"id":   uint(69340),
		"form": uint(0),
		"type": "water",
	},
	"Beatmarine": {
		"id":   uint(69341),
		"form": uint(0),
		"type": "water",
	},
	"Substarr": {
		"id":   uint(69342),
		"form": uint(0),
		"type": "water",
	},
	"Hofucno": {
		"id":   uint(69343),
		"form": uint(0),
		"type": "water",
	},
	"Hornigiri": {
		"id":   uint(69344),
		"form": uint(0),
		"type": "ice",
	},
	"Hornicier": {
		"id":   uint(69345),
		"form": uint(0),
		"type": "ice",
	},
	"Hornititan": {
		"id":   uint(69346),
		"form": uint(0),
		"type": "ice",
	},
	"Preasu": {
		"id":   uint(69347),
		"form": uint(0),
		"type": "electric",
	},
	"Undastand": {
		"id":   uint(69348),
		"form": uint(0),
		"type": "electric",
	},
	"Warney": {
		"id":   uint(69349),
		"form": uint(0),
		"type": "normal",
	},
	"Banageddon": {
		"id":   uint(69350),
		"form": uint(0),
		"type": "normal",
	},
	"Acufront": {
		"id":   uint(69351),
		"form": uint(0),
		"type": "normal",
	},
	"Militant": {
		"id":   uint(69352),
		"form": uint(0),
		"type": "bug",
	},
	"Shiggydig": {
		"id":   uint(69353),
		"form": uint(0),
		"type": "normal",
	},
	"Signot": {
		"id":   uint(69354),
		"form": uint(0),
		"type": "ground",
	},
	"Reagain": {
		"id":   uint(69355),
		"form": uint(0),
		"type": "ground",
	},
	"Ormite": {
		"id":   uint(69356),
		"form": uint(0),
		"type": "rock",
	},
	"Viristal": {
		"id":   uint(69357),
		"form": uint(0),
		"type": "rock",
	},
	"Rolango": {
		"id":   uint(69358),
		"form": uint(0),
		"type": "dragon",
	},
	"Dreameme": {
		"id":   uint(69359),
		"form": uint(0),
		"type": "dragon",
	},
	"Pearchie": {
		"id":   uint(69360),
		"form": uint(0),
		"type": "grass",
	},
	"Tarditank": {
		"id":   uint(69361),
		"form": uint(0),
		"type": "bug",
	},
	"Iplora": {
		"id":   uint(69362),
		"form": uint(0),
		"type": "bug",
	},
	"Chromox": {
		"id":   uint(69363),
		"form": uint(0),
		"type": "fire",
	},
	"Pitayen": {
		"id":   uint(69364),
		"form": uint(0),
		"type": "grass",
	},
	"Dragaya": {
		"id":   uint(69365),
		"form": uint(0),
		"type": "grass",
	},
	"Frutagon": {
		"id":   uint(69366),
		"form": uint(0),
		"type": "grass",
	},
	"Biteki": {
		"id":   uint(69367),
		"form": uint(0),
		"type": "ice",
	},
	"Sesquatch": {
		"id":   uint(69368),
		"form": uint(0),
		"type": "normal",
	},
	"Fireshi": {
		"id":   uint(69369),
		"form": uint(0),
		"type": "fire",
	},
	"Fireshitwi": {
		"id":   uint(69370),
		"form": uint(0),
		"type": "fire",
	},
	"Fireshitre": {
		"id":   uint(69371),
		"form": uint(0),
		"type": "fire",
	},
	"Slugbud": {
		"id":   uint(69372),
		"form": uint(0),
		"type": "water",
	},
	"Slughug": {
		"id":   uint(69373),
		"form": uint(0),
		"type": "water",
	},
	"Slugfugg": {
		"id":   uint(69374),
		"form": uint(0),
		"type": "water",
	},
	"Skeletroll": {
		"id":   uint(69375),
		"form": uint(0),
		"type": "ground",
	},
	"Spookscare": {
		"id":   uint(69376),
		"form": uint(0),
		"type": "ground",
	},
	"Goryannus": {
		"id":   uint(69377),
		"form": uint(0),
		"type": "ground",
	},
	"Lankong": {
		"id":   uint(69378),
		"form": uint(0),
		"type": "normal",
	},
	"Ballankey": {
		"id":   uint(69379),
		"form": uint(0),
		"type": "normal",
	},
	"Funnedong": {
		"id":   uint(69380),
		"form": uint(0),
		"type": "normal",
	},
	"Narwhiz": {
		"id":   uint(69381),
		"form": uint(0),
		"type": "water",
	},
	"Niterpent": {
		"id":   uint(69382),
		"form": uint(0),
		"type": "electric",
	},
	"Griffawork": {
		"id":   uint(69383),
		"form": uint(0),
		"type": "psychic",
	},
	"Boarnograf": {
		"id":   uint(69384),
		"form": uint(0),
		"type": "dark",
	},
	"Admoot": {
		"id":   uint(69385),
		"form": uint(0),
		"type": "fairy",
	},
	"Tentaquil": {
		"id":   uint(69386),
		"form": uint(0),
		"type": "bug",
	},
	"Mega Blobbos": {
		"id":   uint(69387),
		"form": uint(0),
		"type": "???",
	},
	"Woot": {
		"id":   uint(666000),
		"form": uint(0),
		"type": "wood",
	},
	"Woodie": {
		"id":   uint(666001),
		"form": uint(0),
		"type": "wood",
	},
	"Arbrood": {
		"id":   uint(666002),
		"form": uint(0),
		"type": "wood",
	},
	"Maggie": {
		"id":   uint(666003),
		"form": uint(0),
		"type": "magma",
	},
	"Magman": {
		"id":   uint(666004),
		"form": uint(0),
		"type": "magma",
	},
	"Lavagun": {
		"id":   uint(666005),
		"form": uint(0),
		"type": "magma",
	},
	"Steamin": {
		"id":   uint(666006),
		"form": uint(0),
		"type": "steam",
	},
	"Steamer": {
		"id":   uint(666007),
		"form": uint(0),
		"type": "steam",
	},
	"Steamboatle": {
		"id":   uint(666008),
		"form": uint(0),
		"type": "steam",
	},
	"Pagie": {
		"id":   uint(666009),
		"form": uint(0),
		"type": "normal",
	},
	"Swanaper": {
		"id":   uint(666010),
		"form": uint(0),
		"type": "paper",
	},
	"Smoball": {
		"id":   uint(666011),
		"form": uint(0),
		"type": "ice",
	},
	"Smoman": {
		"id":   uint(666012),
		"form": uint(0),
		"type": "ice",
	},
	"Snofistor": {
		"id":   uint(666013),
		"form": uint(0),
		"type": "ice",
	},
	"Frostear": {
		"id":   uint(666014),
		"form": uint(0),
		"type": "ice",
	},
	"Elastim": {
		"id":   uint(666015),
		"form": uint(0),
		"type": "bug",
	},
	"Bouncect": {
		"id":   uint(666016),
		"form": uint(0),
		"type": "bug",
	},
	"Rubberfly": {
		"id":   uint(666017),
		"form": uint(0),
		"type": "bug",
	},
	"Rubbero": {
		"id":   uint(666018),
		"form": uint(0),
		"type": "normal",
	},
	"Cosmock": {
		"id":   uint(666019),
		"form": uint(0),
		"type": "rock",
	},
	"Solaioss": {
		"id":   uint(666020),
		"form": uint(0),
		"type": "rock",
	},
	"Lunighton": {
		"id":   uint(666021),
		"form": uint(0),
		"type": "rock",
	},
	"Chipit": {
		"id":   uint(666022),
		"form": uint(0),
		"type": "food",
	},
	"Chokolit": {
		"id":   uint(666023),
		"form": uint(0),
		"type": "food",
	},
	"Undlouis": {
		"id":   uint(666024),
		"form": uint(0),
		"type": "zombie",
	},
	"Hundead": {
		"id":   uint(666025),
		"form": uint(0),
		"type": "zombie",
	},
	"Zomboom": {
		"id":   uint(666026),
		"form": uint(0),
		"type": "zombie",
	},
	"Smobill": {
		"id":   uint(666027),
		"form": uint(0),
		"type": "zombie",
	},
	"Uncrow": {
		"id":   uint(666028),
		"form": uint(0),
		"type": "zombie",
	},
	"Haicaw": {
		"id":   uint(666029),
		"form": uint(0),
		"type": "zombie",
	},
	"Rubadubb": {
		"id":   uint(666030),
		"form": uint(0),
		"type": "rubber",
	},
	"Splissplash": {
		"id":   uint(666031),
		"form": uint(0),
		"type": "rubber",
	},
	"Bottleo": {
		"id":   uint(666032),
		"form": uint(0),
		"type": "plastic",
	},
	"Bottleodrake": {
		"id":   uint(666033),
		"form": uint(0),
		"type": "plastic",
	},
	"Chillzie": {
		"id":   uint(666034),
		"form": uint(0),
		"type": "zombie",
	},
	"Spotni": {
		"id":   uint(666035),
		"form": uint(0),
		"type": "tech",
	},
	"Sateli": {
		"id":   uint(666036),
		"form": uint(0),
		"type": "tech",
	},
	"Shattarate": {
		"id":   uint(666037),
		"form": uint(0),
		"type": "glass",
	},
	"Floatube": {
		"id":   uint(666038),
		"form": uint(0),
		"type": "rubber",
	},
	"Drubber": {
		"id":   uint(666039),
		"form": uint(0),
		"type": "rubber",
	},
	"Battie": {
		"id":   uint(666040),
		"form": uint(0),
		"type": "wood",
	},
	"Capsileau": {
		"id":   uint(666041),
		"form": uint(0),
		"type": "wood",
	},
	"Tuterpill": {
		"id":   uint(666042),
		"form": uint(0),
		"type": "bug",
	},
	"Caskoon": {
		"id":   uint(666043),
		"form": uint(0),
		"type": "bug",
	},
	"Tutterfly": {
		"id":   uint(666044),
		"form": uint(0),
		"type": "bug",
	},
	"Tiroar": {
		"id":   uint(666045),
		"form": uint(0),
		"type": "rubber",
	},
	"Wheelus": {
		"id":   uint(666046),
		"form": uint(0),
		"type": "rubber",
	},
	"Viriv": {
		"id":   uint(666047),
		"form": uint(0),
		"type": "cyber",
	},
	"Sphinxiant": {
		"id":   uint(666048),
		"form": uint(0),
		"type": "ground",
	},
	"Volcime": {
		"id":   uint(666049),
		"form": uint(0),
		"type": "rock",
	},
	"Heruptin": {
		"id":   uint(666050),
		"form": uint(0),
		"type": "rock",
	},
	"Bairy": {
		"id":   uint(666051),
		"form": uint(0),
		"type": "fairy",
	},
	"Chairy": {
		"id":   uint(666052),
		"form": uint(0),
		"type": "fairy",
	},
	"Goom": {
		"id":   uint(666053),
		"form": uint(0),
		"type": "rubber",
	},
	"Chewim": {
		"id":   uint(666054),
		"form": uint(0),
		"type": "rubber",
	},
	"Gummi": {
		"id":   uint(666055),
		"form": uint(0),
		"type": "rubber",
	},
	"Cuplet": {
		"id":   uint(666056),
		"form": uint(0),
		"type": "food",
	},
	"Beaucake": {
		"id":   uint(666057),
		"form": uint(0),
		"type": "food",
	},
	"Gatteri": {
		"id":   uint(666058),
		"form": uint(0),
		"type": "food",
	},
	"Palmetric": {
		"id":   uint(666059),
		"form": uint(0),
		"type": "grass",
	},
	"Tropilightning": {
		"id":   uint(666060),
		"form": uint(0),
		"type": "grass",
	},
	"Seatie": {
		"id":   uint(666061),
		"form": uint(0),
		"type": "wood",
	},
	"Tenablerus": {
		"id":   uint(666062),
		"form": uint(0),
		"type": "wood",
	},
	"Bargear": {
		"id":   uint(666063),
		"form": uint(0),
		"type": "fairy",
	},
	"Carpla": {
		"id":   uint(666064),
		"form": uint(0),
		"type": "bug",
	},
	"Beetarp": {
		"id":   uint(666065),
		"form": uint(0),
		"type": "bug",
	},
	"Dusite": {
		"id":   uint(666066),
		"form": uint(0),
		"type": "wind",
	},
	"Warchon": {
		"id":   uint(666067),
		"form": uint(0),
		"type": "wind",
	},
	"Teskare": {
		"id":   uint(666068),
		"form": uint(0),
		"type": "zombie",
	},
	"Clavies": {
		"id":   uint(666069),
		"form": uint(0),
		"type": "cyber",
	},
	"Tweeter": {
		"id":   uint(666070),
		"form": uint(0),
		"type": "cyber",
	},
	"Myosmic": {
		"id":   uint(666071),
		"form": uint(0),
		"type": "cyber",
	},
	"Feborius": {
		"id":   uint(666072),
		"form": uint(0),
		"type": "cyber",
	},
	"Hompuff": {
		"id":   uint(666073),
		"form": uint(0),
		"type": "wind",
	},
	"Cummunculus": {
		"id":   uint(666074),
		"form": uint(0),
		"type": "wind",
	},
	"Espionot": {
		"id":   uint(666075),
		"form": uint(0),
		"type": "tech",
	},
	"Hoverot": {
		"id":   uint(666076),
		"form": uint(0),
		"type": "tech",
	},
	"Patima": {
		"id":   uint(666077),
		"form": uint(0),
		"type": "wood",
	},
	"Saplom": {
		"id":   uint(666078),
		"form": uint(0),
		"type": "wood",
	},
	"Fanin": {
		"id":   uint(666079),
		"form": uint(0),
		"type": "tech",
	},
	"Turbind": {
		"id":   uint(666080),
		"form": uint(0),
		"type": "tech",
	},
	"Budsee": {
		"id":   uint(666081),
		"form": uint(0),
		"type": "cyber",
	},
	"Sonotech": {
		"id":   uint(666082),
		"form": uint(0),
		"type": "cyber",
	},
	"Creemo": {
		"id":   uint(666083),
		"form": uint(0),
		"type": "dark",
	},
	"Creeposs": {
		"id":   uint(666084),
		"form": uint(0),
		"type": "dark",
	},
	"Nukreep": {
		"id":   uint(666085),
		"form": uint(0),
		"type": "dark",
	},
	"Recyclat": {
		"id":   uint(666086),
		"form": uint(0),
		"type": "plastic",
	},
	"Drummin": {
		"id":   uint(666087),
		"form": uint(0),
		"type": "ghost",
	},
	"Drumgheist": {
		"id":   uint(666088),
		"form": uint(0),
		"type": "ghost",
	},
	"Grinix": {
		"id":   uint(666089),
		"form": uint(0),
		"type": "fear",
	},
	"Horrorux": {
		"id":   uint(666090),
		"form": uint(0),
		"type": "fear",
	},
	"Nixox": {
		"id":   uint(666091),
		"form": uint(0),
		"type": "normal",
	},
	"Namino": {
		"id":   uint(666092),
		"form": uint(0),
		"type": "normal",
	},
	"Radien": {
		"id":   uint(666093),
		"form": uint(0),
		"type": "nuclear",
	},
	"Fissiom": {
		"id":   uint(666094),
		"form": uint(0),
		"type": "nuclear",
	},
	"Mousee": {
		"id":   uint(666095),
		"form": uint(0),
		"type": "tech",
	},
	"Mopockit": {
		"id":   uint(666096),
		"form": uint(0),
		"type": "food",
	},
	"Hotpack": {
		"id":   uint(666097),
		"form": uint(0),
		"type": "food",
	},
	"Magmocket": {
		"id":   uint(666098),
		"form": uint(0),
		"type": "food",
	},
	"Storactus": {
		"id":   uint(666099),
		"form": uint(0),
		"type": "grass",
	},
	"Miracact": {
		"id":   uint(666100),
		"form": uint(0),
		"type": "grass",
	},
	"Boxy": {
		"id":   uint(666101),
		"form": uint(0),
		"type": "wood",
	},
	"Chestox": {
		"id":   uint(666102),
		"form": uint(0),
		"type": "wood",
	},
	"Gargie": {
		"id":   uint(666103),
		"form": uint(0),
		"type": "rock",
	},
	"Dargouille": {
		"id":   uint(666104),
		"form": uint(0),
		"type": "rock",
	},
	"Pareyemid": {
		"id":   uint(666105),
		"form": uint(0),
		"type": "rock",
	},
	"Chaosemple": {
		"id":   uint(666106),
		"form": uint(0),
		"type": "rock",
	},
	"Orderymid": {
		"id":   uint(666107),
		"form": uint(0),
		"type": "rock",
	},
	"Kawaidesha": {
		"id":   uint(666108),
		"form": uint(0),
		"type": "fairy",
	},
	"Machetit": {
		"id":   uint(666109),
		"form": uint(0),
		"type": "steel",
	},
	"Buzzjaw": {
		"id":   uint(666110),
		"form": uint(0),
		"type": "steel",
	},
	"Razorbash": {
		"id":   uint(666111),
		"form": uint(0),
		"type": "steel",
	},
	"Fabreel": {
		"id":   uint(666112),
		"form": uint(0),
		"type": "fabric",
	},
	"Facarveel": {
		"id":   uint(666113),
		"form": uint(0),
		"type": "fabric",
	},
	"Amorphitto": {
		"id":   uint(666114),
		"form": uint(0),
		"type": "normal",
	},
	"Orscine": {
		"id":   uint(666115),
		"form": uint(0),
		"type": "water",
	},
	"Carolden": {
		"id":   uint(666116),
		"form": uint(0),
		"type": "water",
	},
	"Zomby": {
		"id":   uint(666117),
		"form": uint(0),
		"type": "zombie",
	},
	"Zombeast": {
		"id":   uint(666118),
		"form": uint(0),
		"type": "zombie",
	},
	"Wiondeath": {
		"id":   uint(666119),
		"form": uint(0),
		"type": "zombie",
	},
	"Artsy": {
		"id":   uint(666120),
		"form": uint(0),
		"type": "paint",
	},
	"Canvast": {
		"id":   uint(666121),
		"form": uint(0),
		"type": "paint",
	},
	"Eleoler": {
		"id":   uint(666122),
		"form": uint(0),
		"type": "steam",
	},
	"Behemist": {
		"id":   uint(666123),
		"form": uint(0),
		"type": "steam",
	},
	"Tideno": {
		"id":   uint(666124),
		"form": uint(0),
		"type": "water",
	},
	"Flowjaw": {
		"id":   uint(666125),
		"form": uint(0),
		"type": "water",
	},
	"Warshin": {
		"id":   uint(666126),
		"form": uint(0),
		"type": "water",
	},
	"Arod": {
		"id":   uint(666127),
		"form": uint(0),
		"type": "flying",
	},
	"Atmosro": {
		"id":   uint(666128),
		"form": uint(0),
		"type": "flying",
	},
	"Cosmix": {
		"id":   uint(666129),
		"form": uint(0),
		"type": "cosmic",
	},
	"Univax": {
		"id":   uint(666130),
		"form": uint(0),
		"type": "cosmic",
	},
	"Prisy": {
		"id":   uint(666131),
		"form": uint(0),
		"type": "glass",
	},
	"Prismolis": {
		"id":   uint(666132),
		"form": uint(0),
		"type": "glass",
	},
	"Anol": {
		"id":   uint(666133),
		"form": uint(0),
		"type": "divine",
	},
	"Angeist": {
		"id":   uint(666134),
		"form": uint(0),
		"type": "divine",
	},
	"Eyat": {
		"id":   uint(666135),
		"form": uint(0),
		"type": "dark",
	},
	"Coraut": {
		"id":   uint(666136),
		"form": uint(0),
		"type": "chaos",
	},
	"Oregi": {
		"id":   uint(666137),
		"form": uint(0),
		"type": "wood",
	},
	"Canyou": {
		"id":   uint(666138),
		"form": uint(0),
		"type": "wood",
	},
	"Gueriest": {
		"id":   uint(666139),
		"form": uint(0),
		"type": "normal",
	},
	"Flufamb": {
		"id":   uint(666140),
		"form": uint(0),
		"type": "normal",
	},
	"Lamcoud": {
		"id":   uint(666141),
		"form": uint(0),
		"type": "normal",
	},
	"Gorale": {
		"id":   uint(666142),
		"form": uint(0),
		"type": "normal",
	},
	"Plastolphin": {
		"id":   uint(666143),
		"form": uint(0),
		"type": "water",
	},
	"Dolphottle": {
		"id":   uint(666144),
		"form": uint(0),
		"type": "water",
	},
	"Barab": {
		"id":   uint(666145),
		"form": uint(0),
		"type": "bug",
	},
	"Radicrab": {
		"id":   uint(666146),
		"form": uint(0),
		"type": "bug",
	},
	"Baifre": {
		"id":   uint(666147),
		"form": uint(0),
		"type": "fire",
	},
	"Baumine": {
		"id":   uint(666148),
		"form": uint(0),
		"type": "fire",
	},
	"Microette": {
		"id":   uint(666149),
		"form": uint(0),
		"type": "sound",
	},
	"Speakerro": {
		"id":   uint(666150),
		"form": uint(0),
		"type": "tech",
	},
	"Boomblix": {
		"id":   uint(666151),
		"form": uint(0),
		"type": "tech",
	},
	"Otyab": {
		"id":   uint(666152),
		"form": uint(0),
		"type": "poison",
	},
	"Otyash": {
		"id":   uint(666153),
		"form": uint(0),
		"type": "poison",
	},
	"Ijusgetsbigger": {
		"id":   uint(666154),
		"form": uint(0),
		"type": "glass",
	},
	"Ijusgotbigger": {
		"id":   uint(666155),
		"form": uint(0),
		"type": "glass",
	},
	"Ijusgotsmal": {
		"id":   uint(666156),
		"form": uint(0),
		"type": "glass",
	},
	"Sharpladon": {
		"id":   uint(666157),
		"form": uint(0),
		"type": "water",
	},
	"Ancienkrab": {
		"id":   uint(666158),
		"form": uint(0),
		"type": "water",
	},
	"Fedoraz": {
		"id":   uint(666159),
		"form": uint(0),
		"type": "dark",
	},
	"Euphoreye": {
		"id":   uint(666160),
		"form": uint(0),
		"type": "dark",
	},
	"Mladee": {
		"id":   uint(666161),
		"form": uint(0),
		"type": "dark",
	},
	"Grasses": {
		"id":   uint(666162),
		"form": uint(0),
		"type": "grass",
	},
	"Grasstacles": {
		"id":   uint(666163),
		"form": uint(0),
		"type": "grass",
	},
	"Anguar": {
		"id":   uint(666164),
		"form": uint(0),
		"type": "divine",
	},
	"Archango": {
		"id":   uint(666165),
		"form": uint(0),
		"type": "divine",
	},
	"Devio": {
		"id":   uint(666166),
		"form": uint(0),
		"type": "chaos",
	},
	"Ledevilain": {
		"id":   uint(666167),
		"form": uint(0),
		"type": "chaos",
	},
	"Sola": {
		"id":   uint(666168),
		"form": uint(0),
		"type": "fire",
	},
	"Novais": {
		"id":   uint(666169),
		"form": uint(0),
		"type": "fire",
	},
	"Darkole": {
		"id":   uint(666170),
		"form": uint(0),
		"type": "dark",
	},
	"Bulbight": {
		"id":   uint(666171),
		"form": uint(0),
		"type": "glass",
	},
	"Lightulb": {
		"id":   uint(666172),
		"form": uint(0),
		"type": "glass",
	},
	"Plasbul": {
		"id":   uint(666173),
		"form": uint(0),
		"type": "glass",
	},
	"Solareno": {
		"id":   uint(666174),
		"form": uint(0),
		"type": "light",
	},
	"Electreno": {
		"id":   uint(666175),
		"form": uint(0),
		"type": "electric",
	},
	"Nuclissheno": {
		"id":   uint(666176),
		"form": uint(0),
		"type": "nuclear",
	},
	"Lun": {
		"id":   uint(666177),
		"form": uint(0),
		"type": "dragon",
	},
	"Ryerpent": {
		"id":   uint(666178),
		"form": uint(0),
		"type": "dragon",
	},
	"Ryulung": {
		"id":   uint(666179),
		"form": uint(0),
		"type": "dragon",
	},
	"Winecho": {
		"id":   uint(666180),
		"form": uint(0),
		"type": "psychic",
	},
	"Datagon": {
		"id":   uint(666181),
		"form": uint(0),
		"type": "cyber",
	},
	"Ultirus": {
		"id":   uint(666182),
		"form": uint(0),
		"type": "virus",
	},
	"Techolute": {
		"id":   uint(666183),
		"form": uint(0),
		"type": "tech",
	},
	"Zenature": {
		"id":   uint(666184),
		"form": uint(0),
		"type": "grass",
	},
	"Frankenspook": {
		"id":   uint(666185),
		"form": uint(0),
		"type": "zombie",
	},
	"Johnspleena": {
		"id":   uint(666186),
		"form": uint(0),
		"type": "fighting",
	},
	"Gigasvyre": {
		"id":   uint(666187),
		"form": uint(0),
		"type": "cyber",
	},
	"Iceis": {
		"id":   uint(666188),
		"form": uint(0),
		"type": "ice",
	},
	"Blizzlam": {
		"id":   uint(666189),
		"form": uint(0),
		"type": "ice",
	},
	"Acies": {
		"id":   uint(666190),
		"form": uint(0),
		"type": "dragon",
	},
	"Eldrake": {
		"id":   uint(666191),
		"form": uint(0),
		"type": "dragon",
	},
	"Eldragon": {
		"id":   uint(666192),
		"form": uint(0),
		"type": "dragon",
	},
	"Cosmeak": {
		"id":   uint(666193),
		"form": uint(0),
		"type": "rock",
	},
	"Comistorin": {
		"id":   uint(666194),
		"form": uint(0),
		"type": "rock",
	},
	"Coolor": {
		"id":   uint(666195),
		"form": uint(0),
		"type": "plastic",
	},
	"Fridigor": {
		"id":   uint(666196),
		"form": uint(0),
		"type": "tech",
	},
	"Spongee": {
		"id":   uint(666197),
		"form": uint(0),
		"type": "water",
	},
	"Absorsponge": {
		"id":   uint(666198),
		"form": uint(0),
		"type": "water",
	},
	"Melonboy": {
		"id":   uint(666199),
		"form": uint(0),
		"type": "food",
	},
	"Mellown": {
		"id":   uint(666200),
		"form": uint(0),
		"type": "food",
	},
	"Melonvile": {
		"id":   uint(666201),
		"form": uint(0),
		"type": "food",
	},
	"Earring": {
		"id":   uint(666202),
		"form": uint(0),
		"type": "fighting",
	},
	"Toneout": {
		"id":   uint(666203),
		"form": uint(0),
		"type": "fighting",
	},
	"Xeraphox": {
		"id":   uint(666204),
		"form": uint(0),
		"type": "fighting",
	},
	"Bushin": {
		"id":   uint(666205),
		"form": uint(0),
		"type": "steel",
	},
	"Bushidon": {
		"id":   uint(666206),
		"form": uint(0),
		"type": "steel",
	},
	"Doit": {
		"id":   uint(666207),
		"form": uint(0),
		"type": "sound",
	},
	"Juzt": {
		"id":   uint(666208),
		"form": uint(0),
		"type": "sound",
	},
	"Doita": {
		"id":   uint(666209),
		"form": uint(0),
		"type": "sound",
	},
	"Spooksheet": {
		"id":   uint(666210),
		"form": uint(0),
		"type": "ghost",
	},
	"Strangspook": {
		"id":   uint(666211),
		"form": uint(0),
		"type": "ghost",
	},
	"Lynchost": {
		"id":   uint(666212),
		"form": uint(0),
		"type": "ghost",
	},
	"Boxo": {
		"id":   uint(666213),
		"form": uint(0),
		"type": "paper",
	},
	"Robox": {
		"id":   uint(666214),
		"form": uint(0),
		"type": "paper",
	},
	"Boxtress": {
		"id":   uint(666215),
		"form": uint(0),
		"type": "paper",
	},
	"Eldrute": {
		"id":   uint(666216),
		"form": uint(0),
		"type": "chaos",
	},
	"Eldriz": {
		"id":   uint(666217),
		"form": uint(0),
		"type": "chaos",
	},
	"Eldcraftian": {
		"id":   uint(666218),
		"form": uint(0),
		"type": "chaos",
	},
	"Brigo": {
		"id":   uint(666219),
		"form": uint(0),
		"type": "plastic",
	},
	"Phantoy": {
		"id":   uint(666220),
		"form": uint(0),
		"type": "plastic",
	},
	"Synthite": {
		"id":   uint(666221),
		"form": uint(0),
		"type": "plastic",
	},
	"Belcalf": {
		"id":   uint(666222),
		"form": uint(0),
		"type": "normal",
	},
	"Medijaa": {
		"id":   uint(666223),
		"form": uint(0),
		"type": "fighting",
	},
	"Audinette": {
		"id":   uint(666224),
		"form": uint(0),
		"type": "normal",
	},
	"Nanos": {
		"id":   uint(666225),
		"form": uint(0),
		"type": "tech",
	},
	"Propod": {
		"id":   uint(666226),
		"form": uint(0),
		"type": "tech",
	},
	"Propad": {
		"id":   uint(666227),
		"form": uint(0),
		"type": "tech",
	},
	"Eyeluminid": {
		"id":   uint(666228),
		"form": uint(0),
		"type": "rock",
	},
	"Mexigust": {
		"id":   uint(666229),
		"form": uint(0),
		"type": "ground",
	},
	"Muchoshot": {
		"id":   uint(666230),
		"form": uint(0),
		"type": "ground",
	},
	"Poryrus": {
		"id":   uint(666231),
		"form": uint(0),
		"type": "normal",
	},
	"Poryrusz": {
		"id":   uint(666232),
		"form": uint(0),
		"type": "normal",
	},
	"Playne": {
		"id":   uint(666233),
		"form": uint(0),
		"type": "normal",
	},
	"Jetcraft": {
		"id":   uint(666234),
		"form": uint(0),
		"type": "tech",
	},
	"Eldsparce": {
		"id":   uint(666235),
		"form": uint(0),
		"type": "normal",
	},
	"Dunesdrill": {
		"id":   uint(666236),
		"form": uint(0),
		"type": "normal",
	},
	"Dunsparth": {
		"id":   uint(666237),
		"form": uint(0),
		"type": "normal",
	},
	"Vespiguard": {
		"id":   uint(666238),
		"form": uint(0),
		"type": "bug",
	},
	"Lapsy": {
		"id":   uint(666239),
		"form": uint(0),
		"type": "water",
	},
	"Neonee": {
		"id":   uint(666240),
		"form": uint(0),
		"type": "light",
	},
	"Neonite": {
		"id":   uint(666241),
		"form": uint(0),
		"type": "light",
	},
	"Neonazi": {
		"id":   uint(666242),
		"form": uint(0),
		"type": "light",
	},
	"Encyclodia": {
		"id":   uint(666243),
		"form": uint(0),
		"type": "paper",
	},
	"Bookly": {
		"id":   uint(666244),
		"form": uint(0),
		"type": "paper",
	},
	"Willisp": {
		"id":   uint(666245),
		"form": uint(0),
		"type": "ghost",
	},
	"Willost": {
		"id":   uint(666246),
		"form": uint(0),
		"type": "ghost",
	},
	"Zilly": {
		"id":   uint(666247),
		"form": uint(0),
		"type": "dragon",
	},
	"Godzillus": {
		"id":   uint(666248),
		"form": uint(0),
		"type": "dragon",
	},
	"Raydiat": {
		"id":   uint(666249),
		"form": uint(0),
		"type": "water",
	},
	"Gammaray": {
		"id":   uint(666250),
		"form": uint(0),
		"type": "water",
	},
	"Banor": {
		"id":   uint(666251),
		"form": uint(0),
		"type": "fighting",
	},
	"Hulklear": {
		"id":   uint(666252),
		"form": uint(0),
		"type": "nuclear",
	},
	"Arcaneon": {
		"id":   uint(666253),
		"form": uint(0),
		"type": "magic",
	},
	"Spaceon": {
		"id":   uint(666254),
		"form": uint(0),
		"type": "cosmic",
	},
	"Zyklas": {
		"id":   uint(666255),
		"form": uint(0),
		"type": "poison",
	},
	"Gachamber": {
		"id":   uint(666256),
		"form": uint(0),
		"type": "poison",
	},
	"Elding": {
		"id":   uint(666257),
		"form": uint(0),
		"type": "ice",
	},
	"Eldone": {
		"id":   uint(666258),
		"form": uint(0),
		"type": "ice",
	},
	"Attot": {
		"id":   uint(666259),
		"form": uint(0),
		"type": "electric",
	},
	"Reactiox": {
		"id":   uint(666260),
		"form": uint(0),
		"type": "electric",
	},
	"Nukomb": {
		"id":   uint(666261),
		"form": uint(0),
		"type": "steel",
	},
	"Reqimissile": {
		"id":   uint(666262),
		"form": uint(0),
		"type": "steel",
	},
	"Goatkun": {
		"id":   uint(666263),
		"form": uint(0),
		"type": "chaos",
	},
	"Goatemon": {
		"id":   uint(666264),
		"form": uint(0),
		"type": "chaos",
	},
	"Sushish": {
		"id":   uint(666265),
		"form": uint(0),
		"type": "water",
	},
	"Sashumish": {
		"id":   uint(666266),
		"form": uint(0),
		"type": "water",
	},
	"Solareon": {
		"id":   uint(666267),
		"form": uint(0),
		"type": "light",
	},
	"Snowdoge": {
		"id":   uint(666268),
		"form": uint(0),
		"type": "poison",
	},
	"Cocaiturd": {
		"id":   uint(666269),
		"form": uint(0),
		"type": "poison",
	},
	"Bongorilla": {
		"id":   uint(666270),
		"form": uint(0),
		"type": "grass",
	},
	"Dongorilla": {
		"id":   uint(666271),
		"form": uint(0),
		"type": "grass",
	},
	"Altarule": {
		"id":   uint(666272),
		"form": uint(0),
		"type": "dragon",
	},
	"Camerano": {
		"id":   uint(666273),
		"form": uint(0),
		"type": "fire",
	},
	"Sharpoat": {
		"id":   uint(666274),
		"form": uint(0),
		"type": "water",
	},
	"Houndhell": {
		"id":   uint(666275),
		"form": uint(0),
		"type": "dark",
	},
	"Pinswing": {
		"id":   uint(666276),
		"form": uint(0),
		"type": "bug",
	},
	"Aerodactold": {
		"id":   uint(666277),
		"form": uint(0),
		"type": "rock",
	},
	"Herculecross": {
		"id":   uint(666278),
		"form": uint(0),
		"type": "bug",
	},
	"Diebbuk": {
		"id":   uint(666279),
		"form": uint(0),
		"type": "ghost",
	},
	"Diebbyuk": {
		"id":   uint(666280),
		"form": uint(0),
		"type": "ghost",
	},
	"Pixall": {
		"id":   uint(666281),
		"form": uint(0),
		"type": "cyber",
	},
	"Pixension": {
		"id":   uint(666282),
		"form": uint(0),
		"type": "cyber",
	},
	"Combug": {
		"id":   uint(666283),
		"form": uint(0),
		"type": "bug",
	},
	"Buglitch": {
		"id":   uint(666284),
		"form": uint(0),
		"type": "bug",
	},
	"Smirtsnek": {
		"id":   uint(666285),
		"form": uint(0),
		"type": "tech",
	},
	"Smirtrsnekek": {
		"id":   uint(666286),
		"form": uint(0),
		"type": "tech",
	},
	"Narwhar": {
		"id":   uint(666287),
		"form": uint(0),
		"type": "water",
	},
	"Introll": {
		"id":   uint(666288),
		"form": uint(0),
		"type": "cyber",
	},
	"Trollternet": {
		"id":   uint(666289),
		"form": uint(0),
		"type": "cyber",
	},
	"Brico": {
		"id":   uint(666290),
		"form": uint(0),
		"type": "rock",
	},
	"Bricken": {
		"id":   uint(666291),
		"form": uint(0),
		"type": "rock",
	},
	"Brickor": {
		"id":   uint(666292),
		"form": uint(0),
		"type": "rock",
	},
	"Trojony": {
		"id":   uint(666293),
		"form": uint(0),
		"type": "wood",
	},
	"Trojorse": {
		"id":   uint(666294),
		"form": uint(0),
		"type": "wood",
	},
	"All Is Well": {
		"id":   uint(666295),
		"form": uint(0),
		"type": "virus",
	},
	"There Is Only Hell": {
		"id":   uint(666296),
		"form": uint(0),
		"type": "virus",
	},
	"Helpme": {
		"id":   uint(666297),
		"form": uint(0),
		"type": "virus",
	},
	"Sixsnek": {
		"id":   uint(666298),
		"form": uint(0),
		"type": "magic",
	},
	"Snatan": {
		"id":   uint(666299),
		"form": uint(0),
		"type": "magic",
	},
	"Humanoid": {
		"id":   uint(666300),
		"form": uint(0),
		"type": "normal",
	},
	"Nonhumanoid": {
		"id":   uint(666301),
		"form": uint(0),
		"type": "fear",
	},
	"Black Belt": {
		"id":   uint(666302),
		"form": uint(0),
		"type": "fighting",
	},
	"Omegnown": {
		"id":   uint(666303),
		"form": uint(0),
		"type": "psychic",
	},
	"Maghat": {
		"id":   uint(666304),
		"form": uint(0),
		"type": "fabric",
	},
	"Wizcap": {
		"id":   uint(666305),
		"form": uint(0),
		"type": "fabric",
	},
	"Stando": {
		"id":   uint(666306),
		"form": uint(0),
		"type": "steel",
	},
	"Stando-TheWorld": {
		"id":   uint(666306),
		"form": uint(1),
		"type": "time",
	},
	"Stando-HierophantGreen": {
		"id":   uint(666306),
		"form": uint(2),
		"type": "ghost",
	},
	"Stando-HermitPurple": {
		"id":   uint(666306),
		"form": uint(3),
		"type": "grass",
	},
	"Stando-MagiciansRed": {
		"id":   uint(666306),
		"form": uint(4),
		"type": "magic",
	},
	"Stando-SilverChariot": {
		"id":   uint(666306),
		"form": uint(5),
		"type": "steel",
	},
	"Stando-TheFool": {
		"id":   uint(666306),
		"form": uint(6),
		"type": "tech",
	},
	"Stando-Cream": {
		"id":   uint(666306),
		"form": uint(7),
		"type": "ghost",
	},
	"Stando-DarkBlueMoon": {
		"id":   uint(666306),
		"form": uint(8),
		"type": "water",
	},
	"Stando-TowerOfGray": {
		"id":   uint(666306),
		"form": uint(9),
		"type": "bug",
	},
	"Stando-EbonyDevil": {
		"id":   uint(666306),
		"form": uint(10),
		"type": "ghost",
	},
	"Stando-YellowTemperance": {
		"id":   uint(666306),
		"form": uint(11),
		"type": "normal",
	},
	"Stando-Strength": {
		"id":   uint(666306),
		"form": uint(12),
		"type": "steel",
	},
	"Stando-WheelofFortune": {
		"id":   uint(666306),
		"form": uint(13),
		"type": "tech",
	},
	"Stando-HangedMan": {
		"id":   uint(666306),
		"form": uint(14),
		"type": "zombie",
	},
	"Stando-Emperor": {
		"id":   uint(666306),
		"form": uint(15),
		"type": "steel",
	},
	"Stando-Emperess": {
		"id":   uint(666306),
		"form": uint(16),
		"type": "virus",
	},
	"Stando-Justice": {
		"id":   uint(666306),
		"form": uint(17),
		"type": "ghost",
	},
	"Stando-Lovers": {
		"id":   uint(666306),
		"form": uint(18),
		"type": "bug",
	},
	"Stando-TheSun": {
		"id":   uint(666306),
		"form": uint(19),
		"type": "fire",
	},
	"Stando-DeathThirteen": {
		"id":   uint(666306),
		"form": uint(20),
		"type": "psychic",
	},
	"Stando-Judgment": {
		"id":   uint(666306),
		"form": uint(21),
		"type": "ground",
	},
	"Stando-HighPriestess": {
		"id":   uint(666306),
		"form": uint(22),
		"type": "steel",
	},
	"Stando-Geb": {
		"id":   uint(666306),
		"form": uint(23),
		"type": "water",
	},
	"Stando-Khnum": {
		"id":   uint(666306),
		"form": uint(24),
		"type": "ghost",
	},
	"Stando-Thoth": {
		"id":   uint(666306),
		"form": uint(25),
		"type": "paper",
	},
	"Stando-Anubis": {
		"id":   uint(666306),
		"form": uint(26),
		"type": "ghost",
	},
	"Stando-Bastet": {
		"id":   uint(666306),
		"form": uint(27),
		"type": "steel",
	},
	"Stando-Sethan": {
		"id":   uint(666306),
		"form": uint(28),
		"type": "dark",
	},
	"Stando-Osiris": {
		"id":   uint(666306),
		"form": uint(29),
		"type": "ghost",
	},
	"Stando-Horus": {
		"id":   uint(666306),
		"form": uint(30),
		"type": "ice",
	},
	"Stando-Atum": {
		"id":   uint(666306),
		"form": uint(31),
		"type": "ghost",
	},
	"Stando-TenoreSax": {
		"id":   uint(666306),
		"form": uint(32),
		"type": "ghost",
	},
	"Stando-CrazyDiamond": {
		"id":   uint(666306),
		"form": uint(33),
		"type": "heart",
	},
	"Stando-TheHand": {
		"id":   uint(666306),
		"form": uint(34),
		"type": "fighting",
	},
	"Stando-EchoesAct1": {
		"id":   uint(666306),
		"form": uint(35),
		"type": "sound",
	},
	"Stando-EchoesAct2": {
		"id":   uint(666306),
		"form": uint(36),
		"type": "sound",
	},
	"Stando-EchoesAct3": {
		"id":   uint(666306),
		"form": uint(37),
		"type": "sound",
	},
	"Stando-HeavensDoor": {
		"id":   uint(666306),
		"form": uint(38),
		"type": "paper",
	},
	"Stando-KillerQueen": {
		"id":   uint(666306),
		"form": uint(39),
		"type": "ghost",
	},
	"Stando-AquaNecklace": {
		"id":   uint(666306),
		"form": uint(40),
		"type": "water",
	},
	"Stando-BadCompany": {
		"id":   uint(666306),
		"form": uint(41),
		"type": "steel",
	},
	"Stando-RedHotChilliPepper": {
		"id":   uint(666306),
		"form": uint(42),
		"type": "electric",
	},
	"Stando-TheLock": {
		"id":   uint(666306),
		"form": uint(43),
		"type": "steel",
	},
	"Stando-Surface": {
		"id":   uint(666306),
		"form": uint(44),
		"type": "wood",
	},
	"Stando-LoveDeluxe": {
		"id":   uint(666306),
		"form": uint(45),
		"type": "fabric",
	},
	"Stando-PearlJam": {
		"id":   uint(666306),
		"form": uint(46),
		"type": "food",
	},
	"Stando-Ratt": {
		"id":   uint(666306),
		"form": uint(47),
		"type": "tech",
	},
	"Stando-Harvest": {
		"id":   uint(666306),
		"form": uint(48),
		"type": "bug",
	},
	"Stando-Cinderella": {
		"id":   uint(666306),
		"form": uint(49),
		"type": "heart",
	},
	"Stando-BoyIIMan": {
		"id":   uint(666306),
		"form": uint(50),
		"type": "rock",
	},
	"Stando-HighwayStar": {
		"id":   uint(666306),
		"form": uint(51),
		"type": "rubber",
	},
	"Stando-StrayCat": {
		"id":   uint(666306),
		"form": uint(52),
		"type": "grass",
	},
	"Stando-SuperFly": {
		"id":   uint(666306),
		"form": uint(53),
		"type": "steel",
	},
	"Stando-Enigma": {
		"id":   uint(666306),
		"form": uint(54),
		"type": "paper",
	},
	"Stando-CheapTrick": {
		"id":   uint(666306),
		"form": uint(55),
		"type": "dark",
	},
	"Stando-AtomHeartFather": {
		"id":   uint(666306),
		"form": uint(56),
		"type": "ghost",
	},
	"Stando-ActhungBaby": {
		"id":   uint(666306),
		"form": uint(57),
		"type": "normal",
	},
	"Stando-EarthWindandFire": {
		"id":   uint(666306),
		"form": uint(58),
		"type": "qmarks",
	},
	"Stando-GoldExperience": {
		"id":   uint(666306),
		"form": uint(59),
		"type": "light",
	},
	"Stando-StickyFingers": {
		"id":   uint(666306),
		"form": uint(60),
		"type": "fabric",
	},
	"Stando-MoodyBlues": {
		"id":   uint(666306),
		"form": uint(61),
		"type": "sound",
	},
	"Stando-SexPistols": {
		"id":   uint(666306),
		"form": uint(62),
		"type": "steel",
	},
	"Stando-Aerosmith": {
		"id":   uint(666306),
		"form": uint(63),
		"type": "tech",
	},
	"Stando-PurpleHaze": {
		"id":   uint(666306),
		"form": uint(64),
		"type": "virus",
	},
	"Stando-SpiceGirl": {
		"id":   uint(666306),
		"form": uint(65),
		"type": "rubber",
	},
	"Stando-KingCrimson": {
		"id":   uint(666306),
		"form": uint(66),
		"type": "chaos",
	},
	"Stando-BlackSabbath": {
		"id":   uint(666306),
		"form": uint(67),
		"type": "dark",
	},
	"Stando-SoftMachine": {
		"id":   uint(666306),
		"form": uint(68),
		"type": "steel",
	},
	"Stando-Kraftwork": {
		"id":   uint(666306),
		"form": uint(69),
		"type": "cosmic",
	},
	"Stando-LittleFeet": {
		"id":   uint(666306),
		"form": uint(70),
		"type": "normal",
	},
	"Stando-ManintheMirror": {
		"id":   uint(666306),
		"form": uint(71),
		"type": "glass",
	},
	"Stando-MrPresident": {
		"id":   uint(666306),
		"form": uint(72),
		"type": "ground",
	},
	"Stando-BeachBoy": {
		"id":   uint(666306),
		"form": uint(73),
		"type": "water",
	},
	"Stando-GratefulDead": {
		"id":   uint(666306),
		"form": uint(74),
		"type": "zombie",
	},
	"Stando-BabyFace": {
		"id":   uint(666306),
		"form": uint(75),
		"type": "cyber",
	},
	"Stando-WhiteAlbum": {
		"id":   uint(666306),
		"form": uint(76),
		"type": "ice",
	},
	"Stando-Clash": {
		"id":   uint(666306),
		"form": uint(77),
		"type": "water",
	},
	"Stando-TalkingHead": {
		"id":   uint(666306),
		"form": uint(78),
		"type": "dark",
	},
	"Stando-NotoriousB.I.G": {
		"id":   uint(666306),
		"form": uint(79),
		"type": "virus",
	},
	"Stando-Metallica": {
		"id":   uint(666306),
		"form": uint(80),
		"type": "steel",
	},
	"Stando-GreenDay": {
		"id":   uint(666306),
		"form": uint(81),
		"type": "grass",
	},
	"Stando-Oasis": {
		"id":   uint(666306),
		"form": uint(82),
		"type": "ground",
	},
	"Stando-RollingStones": {
		"id":   uint(666306),
		"form": uint(83),
		"type": "rock",
	},
	"Stando-StoneFree": {
		"id":   uint(666306),
		"form": uint(84),
		"type": "fabric",
	},
	"Stando-WeatherReport": {
		"id":   uint(666306),
		"form": uint(85),
		"type": "wind",
	},
	"Stando-Whitesnake": {
		"id":   uint(666306),
		"form": uint(86),
		"type": "ghost",
	},
	"Stando-ScaryMonsters": {
		"id":   uint(666306),
		"form": uint(87),
		"type": "dragon",
	},
	"Stando-HighwayToHell": {
		"id":   uint(666306),
		"form": uint(88),
		"type": "dark",
	},
	"Stando-Survivor": {
		"id":   uint(666306),
		"form": uint(89),
		"type": "electric",
	},
	"Stando-SkyHigh": {
		"id":   uint(666306),
		"form": uint(90),
		"type": "wind",
	},
	"Stando-SoftandWet": {
		"id":   uint(666306),
		"form": uint(91),
		"type": "water",
	},
	"Stando-CMoon": {
		"id":   uint(666306),
		"form": uint(92),
		"type": "cosmic",
	},
	"Stando-TuskAct1": {
		"id":   uint(666306),
		"form": uint(93),
		"type": "normal",
	},
	"Stando-Ballbreaker": {
		"id":   uint(666306),
		"form": uint(94),
		"type": "steel",
	},
	"Stando-D4C": {
		"id":   uint(666306),
		"form": uint(95),
		"type": "cosmic",
	},
	"Stando-InASilentWay": {
		"id":   uint(666306),
		"form": uint(96),
		"type": "sound",
	},
	"Stando-HeyYa!": {
		"id":   uint(666306),
		"form": uint(97),
		"type": "ghost",
	},
	"Stando-SilverChariotRequiem": {
		"id":   uint(666306),
		"form": uint(98),
		"type": "ghost",
	},
	"Stando-BlueNight": {
		"id":   uint(666306),
		"form": uint(99),
		"type": "water",
	},
	"Stando-HighHighPriestess": {
		"id":   uint(666306),
		"form": uint(100),
		"type": "crystal",
	},
	"Monkid": {
		"id":   uint(666307),
		"form": uint(0),
		"type": "normal",
	},
	"Undyrate": {
		"id":   uint(666308),
		"form": uint(0),
		"type": "water",
	},
	"Skeleton": {
		"id":   uint(666309),
		"form": uint(0),
		"type": "dark",
	},
	"Asgoat": {
		"id":   uint(666310),
		"form": uint(0),
		"type": "steel",
	},
	"Mettabot": {
		"id":   uint(666311),
		"form": uint(0),
		"type": "tech",
	},
	"Nitegant": {
		"id":   uint(666312),
		"form": uint(0),
		"type": "fear",
	},
	"Huntorror": {
		"id":   uint(666313),
		"form": uint(0),
		"type": "fear",
	},
	"Cutehulhu": {
		"id":   uint(666314),
		"form": uint(0),
		"type": "water",
	},
	"Cthuspawn": {
		"id":   uint(666315),
		"form": uint(0),
		"type": "water",
	},
	"Cultulzu": {
		"id":   uint(666316),
		"form": uint(0),
		"type": "water",
	},
	"Steamos": {
		"id":   uint(666317),
		"form": uint(0),
		"type": "steam",
	},
	"Steamist": {
		"id":   uint(666318),
		"form": uint(0),
		"type": "steam",
	},
	"Black Door": {
		"id":   uint(666319),
		"form": uint(0),
		"type": "grass",
	},
	"F00": {
		"id":   uint(666320),
		"form": uint(0),
		"type": "tech",
	},
	"F00-Zen": {
		"id":   uint(666320),
		"form": uint(1),
		"type": "tech",
	},
	"White Hand": {
		"id":   uint(666321),
		"form": uint(0),
		"type": "ghost",
	},
	"Black": {
		"id":   uint(666322),
		"form": uint(0),
		"type": "ghost",
	},
	"Macat": {
		"id":   uint(666323),
		"form": uint(0),
		"type": "zombie",
	},
	"Jackorse": {
		"id":   uint(666324),
		"form": uint(0),
		"type": "zombie",
	},
	"White Door": {
		"id":   uint(666325),
		"form": uint(0),
		"type": "fire",
	},
	"Nukadow": {
		"id":   uint(666326),
		"form": uint(0),
		"type": "ghost",
	},
	"Dragurve": {
		"id":   uint(666327),
		"form": uint(0),
		"type": "dragon",
	},
	"Kingellow": {
		"id":   uint(666328),
		"form": uint(0),
		"type": "fear",
	},
	"Hasterror": {
		"id":   uint(666329),
		"form": uint(0),
		"type": "fear",
	},
	"Byaklight": {
		"id":   uint(666330),
		"form": uint(0),
		"type": "fear",
	},
	"Nyarly": {
		"id":   uint(666331),
		"form": uint(0),
		"type": "chaos",
	},
	"Virurm": {
		"id":   uint(666332),
		"form": uint(0),
		"type": "virus",
	},
	"Malworm": {
		"id":   uint(666333),
		"form": uint(0),
		"type": "virus",
	},
	"Majin": {
		"id":   uint(666334),
		"form": uint(0),
		"type": "dark",
	},
	"Red Mist": {
		"id":   uint(666335),
		"form": uint(0),
		"type": "steam",
	},
	"Snaschiel": {
		"id":   uint(666336),
		"form": uint(0),
		"type": "divine",
	},
	"Sangiel": {
		"id":   uint(666337),
		"form": uint(0),
		"type": "divine",
	},
	"Krisiel": {
		"id":   uint(666338),
		"form": uint(0),
		"type": "crystal",
	},
	"Finpip": {
		"id":   uint(666339),
		"form": uint(0),
		"type": "water",
	},
	"Finpelio": {
		"id":   uint(666340),
		"form": uint(0),
		"type": "water",
	},
	"Pinnister": {
		"id":   uint(666341),
		"form": uint(0),
		"type": "water",
	},
	"Roothoot": {
		"id":   uint(666342),
		"form": uint(0),
		"type": "grass",
	},
	"Owlhood": {
		"id":   uint(666343),
		"form": uint(0),
		"type": "grass",
	},
	"Rootspook": {
		"id":   uint(666344),
		"form": uint(0),
		"type": "grass",
	},
	"Emcat": {
		"id":   uint(666345),
		"form": uint(0),
		"type": "fire",
	},
	"Purshot": {
		"id":   uint(666346),
		"form": uint(0),
		"type": "fire",
	},
	"Panthannon": {
		"id":   uint(666347),
		"form": uint(0),
		"type": "fire",
	},
	"Guttsu": {
		"id":   uint(666348),
		"form": uint(0),
		"type": "fighting",
	},
	"Gatsu": {
		"id":   uint(666349),
		"form": uint(0),
		"type": "fighting",
	},
	"Berserguts": {
		"id":   uint(666350),
		"form": uint(0),
		"type": "fighting",
	},
	"Gublin": {
		"id":   uint(666351),
		"form": uint(0),
		"type": "steel",
	},
	"Goblord": {
		"id":   uint(666352),
		"form": uint(0),
		"type": "steel",
	},
	"Gaghocean": {
		"id":   uint(666353),
		"form": uint(0),
		"type": "divine",
	},
	"Indhead": {
		"id":   uint(666354),
		"form": uint(0),
		"type": "zombie",
	},
	"Israfel": {
		"id":   uint(666355),
		"form": uint(0),
		"type": "divine",
	},
	"Sandolcano": {
		"id":   uint(666356),
		"form": uint(0),
		"type": "divine",
	},
	"Mataracid": {
		"id":   uint(666357),
		"form": uint(0),
		"type": "poison",
	},
	"Wimpangel": {
		"id":   uint(666358),
		"form": uint(0),
		"type": "rock",
	},
	"Weepangel": {
		"id":   uint(666359),
		"form": uint(0),
		"type": "rock",
	},
	"Agonangel": {
		"id":   uint(666360),
		"form": uint(0),
		"type": "fear",
	},
	"Seratime": {
		"id":   uint(666361),
		"form": uint(0),
		"type": "divine",
	},
	"Lunound": {
		"id":   uint(666362),
		"form": uint(0),
		"type": "dark",
	},
	"Lycanmoon": {
		"id":   uint(666363),
		"form": uint(0),
		"type": "dark",
	},
	"Acroagunk": {
		"id":   uint(666364),
		"form": uint(0),
		"type": "poison",
	},
	"Atoxicroak": {
		"id":   uint(666365),
		"form": uint(0),
		"type": "poison",
	},
	"Aqwilfish": {
		"id":   uint(666366),
		"form": uint(0),
		"type": "steel",
	},
	"Shelldor": {
		"id":   uint(666367),
		"form": uint(0),
		"type": "water",
	},
	"Dedfishe": {
		"id":   uint(666368),
		"form": uint(0),
		"type": "water",
	},
	"Lightsy": {
		"id":   uint(666369),
		"form": uint(0),
		"type": "light",
	},
	"Syli": {
		"id":   uint(666370),
		"form": uint(0),
		"type": "light",
	},
	"Mawcar": {
		"id":   uint(666371),
		"form": uint(0),
		"type": "cosmic",
	},
	"Acasmoc": {
		"id":   uint(666372),
		"form": uint(0),
		"type": "cosmic",
	},
	"Snowpil": {
		"id":   uint(666373),
		"form": uint(0),
		"type": "bug",
	},
	"Froskoon": {
		"id":   uint(666374),
		"form": uint(0),
		"type": "bug",
	},
	"Rhemorice": {
		"id":   uint(666375),
		"form": uint(0),
		"type": "bug",
	},
	"Twickotweet": {
		"id":   uint(666376),
		"form": uint(0),
		"type": "fear",
	},
	"Phantmask": {
		"id":   uint(666377),
		"form": uint(0),
		"type": "fear",
	},
	"Vampask": {
		"id":   uint(666378),
		"form": uint(0),
		"type": "fear",
	},
	"Witchise": {
		"id":   uint(666379),
		"form": uint(0),
		"type": "fear",
	},
	"Demask": {
		"id":   uint(666380),
		"form": uint(0),
		"type": "fear",
	},
	"Buried Alive": {
		"id":   uint(666381),
		"form": uint(0),
		"type": "ghost",
	},
	"Ningice": {
		"id":   uint(666382),
		"form": uint(0),
		"type": "ice",
	},
	"Ningenice": {
		"id":   uint(666383),
		"form": uint(0),
		"type": "ice",
	},
	"Fireall": {
		"id":   uint(666384),
		"form": uint(0),
		"type": "cyber",
	},
	"Terreon": {
		"id":   uint(666385),
		"form": uint(0),
		"type": "ground",
	},
	"Seapi": {
		"id":   uint(666386),
		"form": uint(0),
		"type": "water",
	},
	"Seahog": {
		"id":   uint(666387),
		"form": uint(0),
		"type": "water",
	},
	"Eldri": {
		"id":   uint(666388),
		"form": uint(0),
		"type": "chaos",
	},
	"Eldreye": {
		"id":   uint(666389),
		"form": uint(0),
		"type": "chaos",
	},
	"Dlycanroc": {
		"id":   uint(666390),
		"form": uint(0),
		"type": "rock",
	},
	"Nlycanroc": {
		"id":   uint(666391),
		"form": uint(0),
		"type": "rock",
	},
	"Peacon": {
		"id":   uint(666392),
		"form": uint(0),
		"type": "light",
	},
	"Pearoyal": {
		"id":   uint(666393),
		"form": uint(0),
		"type": "light",
	},
	"Cragite": {
		"id":   uint(666394),
		"form": uint(0),
		"type": "bug",
	},
	"Canyonite": {
		"id":   uint(666395),
		"form": uint(0),
		"type": "bug",
	},
	"Spirus": {
		"id":   uint(666396),
		"form": uint(0),
		"type": "virus",
	},
	"Splague": {
		"id":   uint(666397),
		"form": uint(0),
		"type": "virus",
	},
	"Lunafake": {
		"id":   uint(666398),
		"form": uint(0),
		"type": "tech",
	},
	"Fonighton": {
		"id":   uint(666399),
		"form": uint(0),
		"type": "tech",
	},
	"Moonse": {
		"id":   uint(666400),
		"form": uint(0),
		"type": "food",
	},
	"Lunarat": {
		"id":   uint(666401),
		"form": uint(0),
		"type": "food",
	},
	"Mooneast": {
		"id":   uint(666402),
		"form": uint(0),
		"type": "psychic",
	},
	"Amagikarp": {
		"id":   uint(666403),
		"form": uint(0),
		"type": "water",
	},
	"Agyarados": {
		"id":   uint(666404),
		"form": uint(0),
		"type": "water",
	},
	"Adratini": {
		"id":   uint(666405),
		"form": uint(0),
		"type": "dragon",
	},
	"Adragonair": {
		"id":   uint(666406),
		"form": uint(0),
		"type": "dragon",
	},
	"Adragonite": {
		"id":   uint(666407),
		"form": uint(0),
		"type": "dragon",
	},
	"Poisapp": {
		"id":   uint(666408),
		"form": uint(0),
		"type": "food",
	},
	"Toxapple": {
		"id":   uint(666409),
		"form": uint(0),
		"type": "food",
	},
	"Burgite": {
		"id":   uint(666410),
		"form": uint(0),
		"type": "ice",
	},
	"Icendless": {
		"id":   uint(666411),
		"form": uint(0),
		"type": "ice",
	},
	"Asandshrew": {
		"id":   uint(666412),
		"form": uint(0),
		"type": "ice",
	},
	"Asandslash": {
		"id":   uint(666413),
		"form": uint(0),
		"type": "ice",
	},
	"Arattata": {
		"id":   uint(666414),
		"form": uint(0),
		"type": "normal",
	},
	"Araticate": {
		"id":   uint(666415),
		"form": uint(0),
		"type": "normal",
	},
	"Avulpix": {
		"id":   uint(666416),
		"form": uint(0),
		"type": "ice",
	},
	"Aninetales": {
		"id":   uint(666417),
		"form": uint(0),
		"type": "ice",
	},
	"Ameowth": {
		"id":   uint(666418),
		"form": uint(0),
		"type": "dark",
	},
	"Apersian": {
		"id":   uint(666419),
		"form": uint(0),
		"type": "dark",
	},
	"Aexeggutor": {
		"id":   uint(666420),
		"form": uint(0),
		"type": "grass",
	},
	"Agrimer": {
		"id":   uint(666421),
		"form": uint(0),
		"type": "poison",
	},
	"Amuk": {
		"id":   uint(666422),
		"form": uint(0),
		"type": "poison",
	},
	"Amarowak": {
		"id":   uint(666423),
		"form": uint(0),
		"type": "bone",
	},
	"Aparas": {
		"id":   uint(666424),
		"form": uint(0),
		"type": "bug",
	},
	"Aparasect": {
		"id":   uint(666425),
		"form": uint(0),
		"type": "bug",
	},
	"Akoffing": {
		"id":   uint(666426),
		"form": uint(0),
		"type": "fire",
	},
	"Aweezing": {
		"id":   uint(666427),
		"form": uint(0),
		"type": "fire",
	},
	"Bmagikarp": {
		"id":   uint(666428),
		"form": uint(0),
		"type": "water",
	},
	"Bgyarados": {
		"id":   uint(666429),
		"form": uint(0),
		"type": "water",
	},
	"Adiglett": {
		"id":   uint(666430),
		"form": uint(0),
		"type": "ground",
	},
	"Adugtrio": {
		"id":   uint(666431),
		"form": uint(0),
		"type": "ground",
	},
	"Ageodude": {
		"id":   uint(666432),
		"form": uint(0),
		"type": "rock",
	},
	"Agraveler": {
		"id":   uint(666433),
		"form": uint(0),
		"type": "rock",
	},
	"Agolem": {
		"id":   uint(666434),
		"form": uint(0),
		"type": "rock",
	},
	"Araichu": {
		"id":   uint(666435),
		"form": uint(0),
		"type": "electric",
	},
	"Apoochyena": {
		"id":   uint(666436),
		"form": uint(0),
		"type": "dark",
	},
	"Amightyena": {
		"id":   uint(666437),
		"form": uint(0),
		"type": "dark",
	},
	"Aelectrode": {
		"id":   uint(666438),
		"form": uint(0),
		"type": "water",
	},
	"Plagfume": {
		"id":   uint(666439),
		"form": uint(0),
		"type": "fairy",
	},
	"Wendigi": {
		"id":   uint(666440),
		"form": uint(0),
		"type": "ice",
	},
	"Wendigore": {
		"id":   uint(666441),
		"form": uint(0),
		"type": "ice",
	},
	"Transformed": {
		"id":   uint(666442),
		"form": uint(0),
		"type": "ice",
	},
	"Transformedf": {
		"id":   uint(666443),
		"form": uint(0),
		"type": "ice",
	},
	"Sahareph": {
		"id":   uint(666444),
		"form": uint(0),
		"type": "cosmic",
	},
	"Delireul": {
		"id":   uint(666445),
		"form": uint(0),
		"type": "cyber",
	},
	"Leliaw": {
		"id":   uint(666446),
		"form": uint(0),
		"type": "ghost",
	},
	"Bardius": {
		"id":   uint(666447),
		"form": uint(0),
		"type": "divine",
	},
	"Zereurior": {
		"id":   uint(666448),
		"form": uint(0),
		"type": "divine",
	},
	"Araelight": {
		"id":   uint(666449),
		"form": uint(0),
		"type": "wind",
	},
	"Armisy": {
		"id":   uint(666450),
		"form": uint(0),
		"type": "divine",
	},
	"Onigi": {
		"id":   uint(666451),
		"form": uint(0),
		"type": "food",
	},
	"Onigirice": {
		"id":   uint(666452),
		"form": uint(0),
		"type": "food",
	},
	"Stoneon": {
		"id":   uint(666453),
		"form": uint(0),
		"type": "rock",
	},
	"Sloggoth": {
		"id":   uint(666454),
		"form": uint(0),
		"type": "cosmic",
	},
	"Sloggeroth": {
		"id":   uint(666455),
		"form": uint(0),
		"type": "cosmic",
	},
	"Baxbak": {
		"id":   uint(666456),
		"form": uint(0),
		"type": "zombie",
	},
	"Gowen": {
		"id":   uint(666457),
		"form": uint(0),
		"type": "divine",
	},
	"Laprageddon": {
		"id":   uint(666458),
		"form": uint(0),
		"type": "dragon",
	},
	"Pmagikarp": {
		"id":   uint(666459),
		"form": uint(0),
		"type": "water",
	},
	"Pgyarados": {
		"id":   uint(666460),
		"form": uint(0),
		"type": "water",
	},
	"Bartle": {
		"id":   uint(666461),
		"form": uint(0),
		"type": "wood",
	},
	"Torteel": {
		"id":   uint(666462),
		"form": uint(0),
		"type": "wood",
	},
	"Knood": {
		"id":   uint(666463),
		"form": uint(0),
		"type": "wood",
	},
	"Mageno": {
		"id":   uint(666464),
		"form": uint(0),
		"type": "magma",
	},
	"Magamen": {
		"id":   uint(666465),
		"form": uint(0),
		"type": "magma",
	},
	"Mageyser": {
		"id":   uint(666466),
		"form": uint(0),
		"type": "magma",
	},
	"Mistice": {
		"id":   uint(666467),
		"form": uint(0),
		"type": "steam",
	},
	"Phantice": {
		"id":   uint(666468),
		"form": uint(0),
		"type": "steam",
	},
	"Serpentice": {
		"id":   uint(666469),
		"form": uint(0),
		"type": "steam",
	},
	"Slinka": {
		"id":   uint(666470),
		"form": uint(0),
		"type": "normal",
	},
	"Slinpis": {
		"id":   uint(666471),
		"form": uint(0),
		"type": "normal",
	},
	"Cornird": {
		"id":   uint(666472),
		"form": uint(0),
		"type": "normal",
	},
	"Cornosera": {
		"id":   uint(666473),
		"form": uint(0),
		"type": "normal",
	},
	"Regacaw": {
		"id":   uint(666474),
		"form": uint(0),
		"type": "normal",
	},
	"Rockon": {
		"id":   uint(666475),
		"form": uint(0),
		"type": "rock",
	},
	"Mossolem": {
		"id":   uint(666476),
		"form": uint(0),
		"type": "rock",
	},
	"Frostolem": {
		"id":   uint(666477),
		"form": uint(0),
		"type": "rock",
	},
	"Paintip": {
		"id":   uint(666478),
		"form": uint(0),
		"type": "paint",
	},
	"Blat": {
		"id":   uint(666479),
		"form": uint(0),
		"type": "blood",
	},
	"Batood": {
		"id":   uint(666480),
		"form": uint(0),
		"type": "blood",
	},
	"Coneet": {
		"id":   uint(666481),
		"form": uint(0),
		"type": "bug",
	},
	"Concoona": {
		"id":   uint(666482),
		"form": uint(0),
		"type": "bug",
	},
	"Stingcano": {
		"id":   uint(666483),
		"form": uint(0),
		"type": "bug",
	},
	"Nothyp": {
		"id":   uint(666484),
		"form": uint(0),
		"type": "normal",
	},
	"Serpentina": {
		"id":   uint(666485),
		"form": uint(0),
		"type": "normal",
	},
	"Conix": {
		"id":   uint(666486),
		"form": uint(0),
		"type": "crystal",
	},
	"Csteelix": {
		"id":   uint(666487),
		"form": uint(0),
		"type": "steel",
	},
	"Metaleon": {
		"id":   uint(666488),
		"form": uint(0),
		"type": "steel",
	},
	"Zephyreon": {
		"id":   uint(666489),
		"form": uint(0),
		"type": "flying",
	},
	"Champeon": {
		"id":   uint(666490),
		"form": uint(0),
		"type": "fighting",
	},
	"Phanteon": {
		"id":   uint(666491),
		"form": uint(0),
		"type": "ghost",
	},
	"Asableye": {
		"id":   uint(666492),
		"form": uint(0),
		"type": "cosmic",
	},
	"Searam": {
		"id":   uint(666493),
		"form": uint(0),
		"type": "fairy",
	},
	"Baitish": {
		"id":   uint(666494),
		"form": uint(0),
		"type": "water",
	},
	"Fishor": {
		"id":   uint(666495),
		"form": uint(0),
		"type": "water",
	},
	"Left Arm": {
		"id":   uint(666496),
		"form": uint(0),
		"type": "ghost",
	},
	"Right Arm": {
		"id":   uint(666497),
		"form": uint(0),
		"type": "divine",
	},
	"Right Leg": {
		"id":   uint(666498),
		"form": uint(0),
		"type": "chaos",
	},
	"Left Leg": {
		"id":   uint(666499),
		"form": uint(0),
		"type": "cosmic",
	},
	"Moon Core": {
		"id":   uint(666500),
		"form": uint(0),
		"type": "blood",
	},
	"Lunar Head": {
		"id":   uint(666501),
		"form": uint(0),
		"type": "divine",
	},
	"Nymphad": {
		"id":   uint(666502),
		"form": uint(0),
		"type": "fairy",
	},
	"Dryasin": {
		"id":   uint(666503),
		"form": uint(0),
		"type": "fairy",
	},
	"Naiqua": {
		"id":   uint(666504),
		"form": uint(0),
		"type": "fairy",
	},
	"Aulpriland": {
		"id":   uint(666505),
		"form": uint(0),
		"type": "fairy",
	},
	"Sylphind": {
		"id":   uint(666506),
		"form": uint(0),
		"type": "fairy",
	},
	"Lampeath": {
		"id":   uint(666507),
		"form": uint(0),
		"type": "fairy",
	},
	"Lunar Guardian": {
		"id":   uint(666508),
		"form": uint(0),
		"type": "divine",
	},
	"Stalig": {
		"id":   uint(666509),
		"form": uint(0),
		"type": "rock",
	},
	"Stalagfright": {
		"id":   uint(666510),
		"form": uint(0),
		"type": "rock",
	},
	"Drafly": {
		"id":   uint(666511),
		"form": uint(0),
		"type": "bug",
	},
	"Drakefly": {
		"id":   uint(666512),
		"form": uint(0),
		"type": "bug",
	},
	"Anguille": {
		"id":   uint(666513),
		"form": uint(0),
		"type": "water",
	},
	"Saphireau": {
		"id":   uint(666514),
		"form": uint(0),
		"type": "water",
	},
	"Foggot": {
		"id":   uint(666515),
		"form": uint(0),
		"type": "poison",
	},
	"Foggeam": {
		"id":   uint(666516),
		"form": uint(0),
		"type": "poison",
	},
	"Nout": {
		"id":   uint(666517),
		"form": uint(0),
		"type": "sound",
	},
	"Donout": {
		"id":   uint(666518),
		"form": uint(0),
		"type": "sound",
	},
	"Sonout": {
		"id":   uint(666519),
		"form": uint(0),
		"type": "sound",
	},
	"Lonout": {
		"id":   uint(666520),
		"form": uint(0),
		"type": "sound",
	},
	"Treblout": {
		"id":   uint(666521),
		"form": uint(0),
		"type": "sound",
	},
	"Twirlano": {
		"id":   uint(666522),
		"form": uint(0),
		"type": "grass",
	},
	"Seedin": {
		"id":   uint(666523),
		"form": uint(0),
		"type": "grass",
	},
	"Diog": {
		"id":   uint(666524),
		"form": uint(0),
		"type": "ground",
	},
	"Demonig": {
		"id":   uint(666525),
		"form": uint(0),
		"type": "ground",
	},
	"Wigglearth": {
		"id":   uint(666526),
		"form": uint(0),
		"type": "bug",
	},
	"Dooworm": {
		"id":   uint(666527),
		"form": uint(0),
		"type": "bug",
	},
	"Perfectzygarde": {
		"id":   uint(666528),
		"form": uint(0),
		"type": "dragon",
	},
	"Dogzygarde": {
		"id":   uint(666529),
		"form": uint(0),
		"type": "dragon",
	},
	"Cellzygarde": {
		"id":   uint(666530),
		"form": uint(0),
		"type": "dragon",
	},
	"Corezygarde": {
		"id":   uint(666531),
		"form": uint(0),
		"type": "dragon",
	},
	"Chypper": {
		"id":   uint(666532),
		"form": uint(0),
		"type": "cyber",
	},
	"Enchip": {
		"id":   uint(666533),
		"form": uint(0),
		"type": "cyber",
	},
	"Systerus": {
		"id":   uint(666534),
		"form": uint(0),
		"type": "cyber",
	},
	"Demoncake": {
		"id":   uint(666535),
		"form": uint(0),
		"type": "food",
	},
	"Devilcak": {
		"id":   uint(666536),
		"form": uint(0),
		"type": "food",
	},
	"Ambroel": {
		"id":   uint(666537),
		"form": uint(0),
		"type": "food",
	},
	"Ambrosine": {
		"id":   uint(666538),
		"form": uint(0),
		"type": "food",
	},
	"Pawno": {
		"id":   uint(666539),
		"form": uint(0),
		"type": "normal",
	},
	"Rookower": {
		"id":   uint(666540),
		"form": uint(0),
		"type": "rock",
	},
	"Chevak": {
		"id":   uint(666541),
		"form": uint(0),
		"type": "ground",
	},
	"Bishess": {
		"id":   uint(666542),
		"form": uint(0),
		"type": "magic",
	},
	"Queeness": {
		"id":   uint(666543),
		"form": uint(0),
		"type": "crystal",
	},
	"Kingo": {
		"id":   uint(666544),
		"form": uint(0),
		"type": "normal",
	},
	"Bubblet": {
		"id":   uint(666545),
		"form": uint(0),
		"type": "water",
	},
	"Bubloform": {
		"id":   uint(666546),
		"form": uint(0),
		"type": "water",
	},
	"Cobrubble": {
		"id":   uint(666547),
		"form": uint(0),
		"type": "water",
	},
	"Drabon": {
		"id":   uint(666548),
		"form": uint(0),
		"type": "water",
	},
	"Leobble": {
		"id":   uint(666549),
		"form": uint(0),
		"type": "water",
	},
	"Burble": {
		"id":   uint(666550),
		"form": uint(0),
		"type": "water",
	},
	"Turble": {
		"id":   uint(666551),
		"form": uint(0),
		"type": "water",
	},
	"Elebubble": {
		"id":   uint(666552),
		"form": uint(0),
		"type": "water",
	},
	"Wock": {
		"id":   uint(666553),
		"form": uint(0),
		"type": "rock",
	},
	"Stowin": {
		"id":   uint(666554),
		"form": uint(0),
		"type": "rock",
	},
	"Winolith": {
		"id":   uint(666555),
		"form": uint(0),
		"type": "rock",
	},
	"Tomatea": {
		"id":   uint(666556),
		"form": uint(0),
		"type": "grass",
	},
	"Chillyip": {
		"id":   uint(666557),
		"form": uint(0),
		"type": "food",
	},
	"Pepperella": {
		"id":   uint(666558),
		"form": uint(0),
		"type": "grass",
	},
	"Bause": {
		"id":   uint(666559),
		"form": uint(0),
		"type": "food",
	},
	"Cloudlet": {
		"id":   uint(666560),
		"form": uint(0),
		"type": "wind",
	},
	"Rainoud": {
		"id":   uint(666561),
		"form": uint(0),
		"type": "wind",
	},
	"Sunnound": {
		"id":   uint(666562),
		"form": uint(0),
		"type": "wind",
	},
	"Frwind": {
		"id":   uint(666563),
		"form": uint(0),
		"type": "wind",
	},
	"Clawnd": {
		"id":   uint(666564),
		"form": uint(0),
		"type": "wind",
	},
	"Thundoud": {
		"id":   uint(666565),
		"form": uint(0),
		"type": "wind",
	},
	"Clwister": {
		"id":   uint(666566),
		"form": uint(0),
		"type": "wind",
	},
	"Tridler": {
		"id":   uint(666567),
		"form": uint(0),
		"type": "tech",
	},
	"Trirayd": {
		"id":   uint(666568),
		"form": uint(0),
		"type": "tech",
	},
	"Trinvasio": {
		"id":   uint(666569),
		"form": uint(0),
		"type": "tech",
	},
	"Lycanmars": {
		"id":   uint(666570),
		"form": uint(0),
		"type": "fighting",
	},
	"Ayyliem": {
		"id":   uint(666571),
		"form": uint(0),
		"type": "cosmic",
	},
	"Ayymao": {
		"id":   uint(666572),
		"form": uint(0),
		"type": "cosmic",
	},
	"Cybirb": {
		"id":   uint(666573),
		"form": uint(0),
		"type": "cyber",
	},
	"Cypecker": {
		"id":   uint(666574),
		"form": uint(0),
		"type": "cyber",
	},
	"Ufowoah": {
		"id":   uint(666575),
		"form": uint(0),
		"type": "grass",
	},
	"Ufolien": {
		"id":   uint(666576),
		"form": uint(0),
		"type": "grass",
	},
	"Ufoear": {
		"id":   uint(666577),
		"form": uint(0),
		"type": "cosmic",
	},
	"Searice": {
		"id":   uint(666578),
		"form": uint(0),
		"type": "fire",
	},
	"Regidark": {
		"id":   uint(666579),
		"form": uint(0),
		"type": "dark",
	},
	"Drakosm": {
		"id":   uint(666580),
		"form": uint(0),
		"type": "dragon",
	},
	"Pagolem": {
		"id":   uint(666581),
		"form": uint(0),
		"type": "paper",
	},
	"Cristo": {
		"id":   uint(666582),
		"form": uint(0),
		"type": "crystal",
	},
	"Flashard": {
		"id":   uint(666583),
		"form": uint(0),
		"type": "crystal",
	},
	"Prishard": {
		"id":   uint(666584),
		"form": uint(0),
		"type": "crystal",
	},
	"Phanshard": {
		"id":   uint(666585),
		"form": uint(0),
		"type": "crystal",
	},
	"Storshard": {
		"id":   uint(666586),
		"form": uint(0),
		"type": "crystal",
	},
	"Baall": {
		"id":   uint(666587),
		"form": uint(0),
		"type": "dark",
	},
	"Batt": {
		"id":   uint(666588),
		"form": uint(0),
		"type": "wood",
	},
	"Tadpal": {
		"id":   uint(666589),
		"form": uint(0),
		"type": "water",
	},
	"Frogpal": {
		"id":   uint(666590),
		"form": uint(0),
		"type": "water",
	},
	"Marir": {
		"id":   uint(666591),
		"form": uint(0),
		"type": "glass",
	},
	"Castleet": {
		"id":   uint(666592),
		"form": uint(0),
		"type": "glass",
	},
	"Cornel": {
		"id":   uint(666593),
		"form": uint(0),
		"type": "food",
	},
	"Grancorn": {
		"id":   uint(666594),
		"form": uint(0),
		"type": "food",
	},
	"Shalight": {
		"id":   uint(666595),
		"form": uint(0),
		"type": "light",
	},
	"Lumix": {
		"id":   uint(666596),
		"form": uint(0),
		"type": "light",
	},
	"Shadrix": {
		"id":   uint(666597),
		"form": uint(0),
		"type": "dark",
	},
	"Lettucee": {
		"id":   uint(666598),
		"form": uint(0),
		"type": "grass",
	},
	"Saladman": {
		"id":   uint(666599),
		"form": uint(0),
		"type": "grass",
	},
	"Sliwa": {
		"id":   uint(666600),
		"form": uint(0),
		"type": "ice",
	},
	"Cawice": {
		"id":   uint(666601),
		"form": uint(0),
		"type": "ice",
	},
	"Hassassin": {
		"id":   uint(666602),
		"form": uint(0),
		"type": "dark",
	},
	"Berserkalot": {
		"id":   uint(666603),
		"form": uint(0),
		"type": "fairy",
	},
	"Saibah": {
		"id":   uint(666604),
		"form": uint(0),
		"type": "dragon",
	},
	"Rideusa": {
		"id":   uint(666605),
		"form": uint(0),
		"type": "blood",
	},
	"Girgamesh": {
		"id":   uint(666606),
		"form": uint(0),
		"type": "divine",
	},
	"Castedea": {
		"id":   uint(666607),
		"form": uint(0),
		"type": "magic",
	},
	"Lanculain": {
		"id":   uint(666608),
		"form": uint(0),
		"type": "fighting",
	},
	"Bersercules": {
		"id":   uint(666609),
		"form": uint(0),
		"type": "fighting",
	},
	"Ballim": {
		"id":   uint(666610),
		"form": uint(0),
		"type": "rubber",
	},
	"Helift": {
		"id":   uint(666611),
		"form": uint(0),
		"type": "rubber",
	},
	"Shrul": {
		"id":   uint(666612),
		"form": uint(0),
		"type": "grass",
	},
	"Smoov": {
		"id":   uint(666613),
		"form": uint(0),
		"type": "grass",
	},
	"Sporeel": {
		"id":   uint(666614),
		"form": uint(0),
		"type": "grass",
	},
	"Bonnacalf": {
		"id":   uint(666615),
		"form": uint(0),
		"type": "poison",
	},
	"Bonaflatuen": {
		"id":   uint(666616),
		"form": uint(0),
		"type": "poison",
	},
	"Viria": {
		"id":   uint(666617),
		"form": uint(0),
		"type": "virus",
	},
	"Verius": {
		"id":   uint(666618),
		"form": uint(0),
		"type": "virus",
	},
	"Viroison": {
		"id":   uint(666619),
		"form": uint(0),
		"type": "virus",
	},
	"Feaverus": {
		"id":   uint(666620),
		"form": uint(0),
		"type": "virus",
	},
	"Cuberus": {
		"id":   uint(666621),
		"form": uint(0),
		"type": "virus",
	},
	"Shockerus": {
		"id":   uint(666622),
		"form": uint(0),
		"type": "virus",
	},
	"Snoozerus": {
		"id":   uint(666623),
		"form": uint(0),
		"type": "virus",
	},
	"Bonfen": {
		"id":   uint(666624),
		"form": uint(0),
		"type": "wood",
	},
	"Fitree": {
		"id":   uint(666625),
		"form": uint(0),
		"type": "wood",
	},
	"Campyre": {
		"id":   uint(666626),
		"form": uint(0),
		"type": "wood",
	},
	"Cobret": {
		"id":   uint(666627),
		"form": uint(0),
		"type": "fabric",
	},
	"Serfric": {
		"id":   uint(666628),
		"form": uint(0),
		"type": "fabric",
	},
	"Spiortal": {
		"id":   uint(666629),
		"form": uint(0),
		"type": "bug",
	},
	"Spiortex": {
		"id":   uint(666630),
		"form": uint(0),
		"type": "bug",
	},
	"Grifawk": {
		"id":   uint(666631),
		"form": uint(0),
		"type": "light",
	},
	"Grifallen": {
		"id":   uint(666632),
		"form": uint(0),
		"type": "normal",
	},
	"Griffisen": {
		"id":   uint(666633),
		"form": uint(0),
		"type": "divine",
	},
	"Ulycanroc": {
		"id":   uint(666634),
		"form": uint(0),
		"type": "rock",
	},
	"Acaterpie": {
		"id":   uint(666635),
		"form": uint(0),
		"type": "bug",
	},
	"Ametapod": {
		"id":   uint(666636),
		"form": uint(0),
		"type": "bug",
	},
	"Abutterfree": {
		"id":   uint(666637),
		"form": uint(0),
		"type": "bug",
	},
	"Amantyke": {
		"id":   uint(666638),
		"form": uint(0),
		"type": "electric",
	},
	"Amantine": {
		"id":   uint(666639),
		"form": uint(0),
		"type": "electric",
	},
	"Aoddish": {
		"id":   uint(666640),
		"form": uint(0),
		"type": "grass",
	},
	"Agloom": {
		"id":   uint(666641),
		"form": uint(0),
		"type": "grass",
	},
	"Avileplume": {
		"id":   uint(666642),
		"form": uint(0),
		"type": "grass",
	},
	"Agrowlithe": {
		"id":   uint(666643),
		"form": uint(0),
		"type": "water",
	},
	"Aarcanine": {
		"id":   uint(666644),
		"form": uint(0),
		"type": "water",
	},
	"Atorkoal": {
		"id":   uint(666645),
		"form": uint(0),
		"type": "steam",
	},
	"Atogepi": {
		"id":   uint(666646),
		"form": uint(0),
		"type": "dragon",
	},
	"Atogetic": {
		"id":   uint(666647),
		"form": uint(0),
		"type": "dragon",
	},
	"Atogekiss": {
		"id":   uint(666648),
		"form": uint(0),
		"type": "dragon",
	},
	"Rathy": {
		"id":   uint(666649),
		"form": uint(0),
		"type": "fire",
	},
	"Rathor": {
		"id":   uint(666650),
		"form": uint(0),
		"type": "fire",
	},
	"Rathlos": {
		"id":   uint(666651),
		"form": uint(0),
		"type": "fire",
	},
	"Rathiana": {
		"id":   uint(666652),
		"form": uint(0),
		"type": "fire",
	},
	"Jhemoran": {
		"id":   uint(666653),
		"form": uint(0),
		"type": "dragon",
	},
	"Reimaiden": {
		"id":   uint(666654),
		"form": uint(0),
		"type": "normal",
	},
	"Mariswitch": {
		"id":   uint(666655),
		"form": uint(0),
		"type": "light",
	},
	"Sakuyaid": {
		"id":   uint(666656),
		"form": uint(0),
		"type": "steel",
	},
	"Youmurai": {
		"id":   uint(666657),
		"form": uint(0),
		"type": "ghost",
	},
	"Cirnai": {
		"id":   uint(666658),
		"form": uint(0),
		"type": "fairy",
	},
	"Godzillante": {
		"id":   uint(666659),
		"form": uint(0),
		"type": "normal",
	},
	"Gorillaimo": {
		"id":   uint(666660),
		"form": uint(0),
		"type": "normal",
	},
	"Goopars": {
		"id":   uint(666661),
		"form": uint(0),
		"type": "rubber",
	},
	"Tiwisone": {
		"id":   uint(666662),
		"form": uint(0),
		"type": "rock",
	},
	"Bouldars": {
		"id":   uint(666663),
		"form": uint(0),
		"type": "steel",
	},
	"Cacemon": {
		"id":   uint(666664),
		"form": uint(0),
		"type": "chaos",
	},
	"Cubaoop": {
		"id":   uint(666665),
		"form": uint(0),
		"type": "plastic",
	},
	"Patchician": {
		"id":   uint(666666),
		"form": uint(0),
		"type": "magic",
	},
	"Stelloch": {
		"id":   uint(666667),
		"form": uint(0),
		"type": "ghost",
	},
	"Koisheep": {
		"id":   uint(666668),
		"form": uint(0),
		"type": "normal",
	},
	"Mt": {
		"id":   uint(666669),
		"form": uint(0),
		"type": "steel",
	},
	"Mt2": {
		"id":   uint(666670),
		"form": uint(0),
		"type": "steel",
	},
	"Giant Woman": {
		"id":   uint(666671),
		"form": uint(0),
		"type": "normal",
	},
	"Transport": {
		"id":   uint(666672),
		"form": uint(0),
		"type": "tech",
	},
	"Cufo": {
		"id":   uint(666673),
		"form": uint(0),
		"type": "flying",
	},
	"Cufo2": {
		"id":   uint(666674),
		"form": uint(0),
		"type": "psychic",
	},
	"Lidra": {
		"id":   uint(666675),
		"form": uint(0),
		"type": "light",
	},
	"Dardra": {
		"id":   uint(666676),
		"form": uint(0),
		"type": "dark",
	},
	"Shadight": {
		"id":   uint(666677),
		"form": uint(0),
		"type": "light",
	},
	"Jaggu": {
		"id":   uint(666678),
		"form": uint(0),
		"type": "water",
	},
	"Crocky": {
		"id":   uint(666679),
		"form": uint(0),
		"type": "dragon",
	},
	"Deer": {
		"id":   uint(666680),
		"form": uint(0),
		"type": "normal",
	},
	"Kokana": {
		"id":   uint(666681),
		"form": uint(0),
		"type": "bug",
	},
	"Kasanagi": {
		"id":   uint(666682),
		"form": uint(0),
		"type": "bug",
	},
	"Kasanova": {
		"id":   uint(666683),
		"form": uint(0),
		"type": "bug",
	},
	"Aslugma": {
		"id":   uint(666684),
		"form": uint(0),
		"type": "fire",
	},
	"Amagcargo": {
		"id":   uint(666685),
		"form": uint(0),
		"type": "fire",
	},
	"Amerimutt": {
		"id":   uint(666686),
		"form": uint(0),
		"type": "normal",
	},
	"Bremoraid": {
		"id":   uint(666687),
		"form": uint(0),
		"type": "water",
	},
	"Boctillery": {
		"id":   uint(666688),
		"form": uint(0),
		"type": "water",
	},
	"Blossomole": {
		"id":   uint(666689),
		"form": uint(0),
		"type": "grass",
	},
	"Animon": {
		"id":   uint(666690),
		"form": uint(0),
		"type": "normal",
	},
	"Balistagon": {
		"id":   uint(666691),
		"form": uint(0),
		"type": "dragon",
	},
	"Supercannon": {
		"id":   uint(666692),
		"form": uint(0),
		"type": "dragon",
	},
	"Memee": {
		"id":   uint(666693),
		"form": uint(0),
		"type": "meme",
	},
	"Memeo": {
		"id":   uint(666694),
		"form": uint(0),
		"type": "meme",
	},
	"Memaster": {
		"id":   uint(666695),
		"form": uint(0),
		"type": "meme",
	},
	"Smilepup": {
		"id":   uint(666696),
		"form": uint(0),
		"type": "meme",
	},
	"Smilehound": {
		"id":   uint(666697),
		"form": uint(0),
		"type": "meme",
	},
	"Flambear": {
		"id":   uint(666698),
		"form": uint(0),
		"type": "fire",
	},
	"Volbear": {
		"id":   uint(666699),
		"form": uint(0),
		"type": "fire",
	},
	"Dynabear": {
		"id":   uint(666700),
		"form": uint(0),
		"type": "fire",
	},
	"Cruz": {
		"id":   uint(666701),
		"form": uint(0),
		"type": "water",
	},
	"Aqua": {
		"id":   uint(666702),
		"form": uint(0),
		"type": "water",
	},
	"Aquaria": {
		"id":   uint(666703),
		"form": uint(0),
		"type": "water",
	},
	"Hoohoo": {
		"id":   uint(666704),
		"form": uint(0),
		"type": "dark",
	},
	"Leafeo": {
		"id":   uint(666705),
		"form": uint(0),
		"type": "grass",
	},
	"Trifox": {
		"id":   uint(666706),
		"form": uint(0),
		"type": "fire",
	},
	"Tangel": {
		"id":   uint(666707),
		"form": uint(0),
		"type": "grass",
	},
	"Gelanla": {
		"id":   uint(666708),
		"form": uint(0),
		"type": "grass",
	},
	"Rayleep": {
		"id":   uint(666709),
		"form": uint(0),
		"type": "water",
	},
	"Numpuff": {
		"id":   uint(666710),
		"form": uint(0),
		"type": "water",
	},
	"Apichu": {
		"id":   uint(666711),
		"form": uint(0),
		"type": "electric",
	},
	"Acleffa": {
		"id":   uint(666712),
		"form": uint(0),
		"type": "fairy",
	},
	"Aigglybuff": {
		"id":   uint(666713),
		"form": uint(0),
		"type": "fairy",
	},
	"Golppy": {
		"id":   uint(666714),
		"form": uint(0),
		"type": "water",
	},
	"Sunmola": {
		"id":   uint(666715),
		"form": uint(0),
		"type": "water",
	},
	"Anchorage": {
		"id":   uint(666716),
		"form": uint(0),
		"type": "water",
	},
	"Grotess": {
		"id":   uint(666717),
		"form": uint(0),
		"type": "water",
	},
	"Para": {
		"id":   uint(666718),
		"form": uint(0),
		"type": "bug",
	},
	"Spidette": {
		"id":   uint(666719),
		"form": uint(0),
		"type": "bug",
	},
	"Tuhedd": {
		"id":   uint(666720),
		"form": uint(0),
		"type": "bug",
	},
	"Chiks": {
		"id":   uint(666721),
		"form": uint(0),
		"type": "normal",
	},
	"Twinz": {
		"id":   uint(666722),
		"form": uint(0),
		"type": "normal",
	},
	"Kirinriki": {
		"id":   uint(666723),
		"form": uint(0),
		"type": "normal",
	},
	"Meowsy": {
		"id":   uint(666724),
		"form": uint(0),
		"type": "normal",
	},
	"Rinring": {
		"id":   uint(666725),
		"form": uint(0),
		"type": "dark",
	},
	"Bellboyant": {
		"id":   uint(666726),
		"form": uint(0),
		"type": "dark",
	},
	"WackTentaquil": {
		"id":   uint(666727),
		"form": uint(0),
		"type": "water",
	},
	"Tripstar": {
		"id":   uint(666728),
		"form": uint(0),
		"type": "bug",
	},
	"Minicorn": {
		"id":   uint(666729),
		"form": uint(0),
		"type": "fire",
	},
	"Aumbreon": {
		"id":   uint(666730),
		"form": uint(0),
		"type": "poison",
	},
	"Turbann": {
		"id":   uint(666731),
		"form": uint(0),
		"type": "water",
	},
	"Grimey": {
		"id":   uint(666732),
		"form": uint(0),
		"type": "poison",
	},
	"Gohng": {
		"id":   uint(666733),
		"form": uint(0),
		"type": "fighting",
	},
	"Topmonhit": {
		"id":   uint(666734),
		"form": uint(0),
		"type": "fighting",
	},
	"Puddi": {
		"id":   uint(666735),
		"form": uint(0),
		"type": "fire",
	},
	"Lefmew": {
		"id":   uint(666736),
		"form": uint(0),
		"type": "grass",
	},
	"Dandemew": {
		"id":   uint(666737),
		"form": uint(0),
		"type": "grass",
	},
	"Cottomew": {
		"id":   uint(666738),
		"form": uint(0),
		"type": "grass",
	},
	"Ballerine": {
		"id":   uint(666739),
		"form": uint(0),
		"type": "psychic",
	},
	"Lipp": {
		"id":   uint(666740),
		"form": uint(0),
		"type": "ice",
	},
	"Elababee": {
		"id":   uint(666741),
		"form": uint(0),
		"type": "electric",
	},
	"Bmagby": {
		"id":   uint(666742),
		"form": uint(0),
		"type": "fire",
	},
	"Bbellossom": {
		"id":   uint(666743),
		"form": uint(0),
		"type": "grass",
	},
	"Belmitt": {
		"id":   uint(666744),
		"form": uint(0),
		"type": "grass",
	},
	"Bomseel": {
		"id":   uint(666745),
		"form": uint(0),
		"type": "water",
	},
	"Ghift": {
		"id":   uint(666746),
		"form": uint(0),
		"type": "ice",
	},
	"Tigrette": {
		"id":   uint(666747),
		"form": uint(0),
		"type": "electric",
	},
	"Electiger": {
		"id":   uint(666748),
		"form": uint(0),
		"type": "electric",
	},
	"Madame": {
		"id":   uint(666749),
		"form": uint(0),
		"type": "normal",
	},
	"Kurstraw": {
		"id":   uint(666750),
		"form": uint(0),
		"type": "ghost",
	},
	"Pangshi": {
		"id":   uint(666751),
		"form": uint(0),
		"type": "ghost",
	},
	"Bmurkrow": {
		"id":   uint(666752),
		"form": uint(0),
		"type": "dark",
	},
	"Happi": {
		"id":   uint(666753),
		"form": uint(0),
		"type": "normal",
	},
	"Scizors": {
		"id":   uint(666754),
		"form": uint(0),
		"type": "bug",
	},
	"Plux": {
		"id":   uint(666755),
		"form": uint(0),
		"type": "bug",
	},
	"Wolfman": {
		"id":   uint(666756),
		"form": uint(0),
		"type": "ice",
	},
	"Warwolf": {
		"id":   uint(666757),
		"form": uint(0),
		"type": "ice",
	},
	"Borygon2": {
		"id":   uint(666758),
		"form": uint(0),
		"type": "normal",
	},
	"Likk": {
		"id":   uint(666759),
		"form": uint(0),
		"type": "normal",
	},
	"Bkingdra": {
		"id":   uint(666760),
		"form": uint(0),
		"type": "water",
	},
	"Rai": {
		"id":   uint(666761),
		"form": uint(0),
		"type": "electric",
	},
	"En": {
		"id":   uint(666762),
		"form": uint(0),
		"type": "fire",
	},
	"Sui": {
		"id":   uint(666763),
		"form": uint(0),
		"type": "water",
	},
	"Bsneasel": {
		"id":   uint(666764),
		"form": uint(0),
		"type": "dark",
	},
	"Tael": {
		"id":   uint(666765),
		"form": uint(0),
		"type": "normal",
	},
	"Pikablu": {
		"id":   uint(666766),
		"form": uint(0),
		"type": "water",
	},
	"Bhooh": {
		"id":   uint(666767),
		"form": uint(0),
		"type": "flying",
	},
	"Udongabbit": {
		"id":   uint(666768),
		"form": uint(0),
		"type": "psychic",
	},
	"Lanceffish": {
		"id":   uint(666769),
		"form": uint(0),
		"type": "water",
	},
	"Sunnee": {
		"id":   uint(666770),
		"form": uint(0),
		"type": "grass",
	},
	"Shrinae": {
		"id":   uint(666771),
		"form": uint(0),
		"type": "wind",
	},
	"Bpoliwag": {
		"id":   uint(666772),
		"form": uint(0),
		"type": "water",
	},
	"Nyosuka": {
		"id":   uint(666773),
		"form": uint(0),
		"type": "water",
	},
	"Bdragonair": {
		"id":   uint(666774),
		"form": uint(0),
		"type": "dragon",
	},
	"Augumon": {
		"id":   uint(666775),
		"form": uint(0),
		"type": "fire",
	},
	"Greymon": {
		"id":   uint(666776),
		"form": uint(0),
		"type": "fire",
	},
	"Devull": {
		"id":   uint(666777),
		"form": uint(0),
		"type": "fire",
	},
	"Bhoundoom": {
		"id":   uint(666778),
		"form": uint(0),
		"type": "fire",
	},
	"Bespeon": {
		"id":   uint(666779),
		"form": uint(0),
		"type": "psychic",
	},
	"Urchall": {
		"id":   uint(666780),
		"form": uint(0),
		"type": "plastic",
	},
	"Spurchin": {
		"id":   uint(666781),
		"form": uint(0),
		"type": "plastic",
	},
	"Luigi": {
		"id":   uint(666782),
		"form": uint(0),
		"type": "normal",
	},
	"Yoshi": {
		"id":   uint(666783),
		"form": uint(0),
		"type": "normal",
	},
	"Ahole": {
		"id":   uint(666784),
		"form": uint(0),
		"type": "ground",
	},
	"Crachole": {
		"id":   uint(666785),
		"form": uint(0),
		"type": "ground",
	},
	"Cracholiath": {
		"id":   uint(666786),
		"form": uint(0),
		"type": "ground",
	},
	"Staresh": {
		"id":   uint(666787),
		"form": uint(0),
		"type": "crystal",
	},
	"Stargem": {
		"id":   uint(666788),
		"form": uint(0),
		"type": "crystal",
	},
	"Eggit": {
		"id":   uint(666789),
		"form": uint(0),
		"type": "food",
	},
	"Eggotto": {
		"id":   uint(666790),
		"form": uint(0),
		"type": "food",
	},
	"Friedack": {
		"id":   uint(666791),
		"form": uint(0),
		"type": "food",
	},
	"Dustunny": {
		"id":   uint(666792),
		"form": uint(0),
		"type": "fabric",
	},
	"Narlord": {
		"id":   uint(666793),
		"form": uint(0),
		"type": "water",
	},
	"Cbulbasaur": {
		"id":   uint(666794),
		"form": uint(0),
		"type": "qmarks",
	},
	"Civysaur": {
		"id":   uint(666795),
		"form": uint(0),
		"type": "qmarks",
	},
	"Cvenusaur": {
		"id":   uint(666796),
		"form": uint(0),
		"type": "qmarks",
	},
	"Ccharmander": {
		"id":   uint(666797),
		"form": uint(0),
		"type": "qmarks",
	},
	"Ccharmeleon": {
		"id":   uint(666798),
		"form": uint(0),
		"type": "qmarks",
	},
	"Ccharizard": {
		"id":   uint(666799),
		"form": uint(0),
		"type": "qmarks",
	},
	"Csquirtle": {
		"id":   uint(666800),
		"form": uint(0),
		"type": "qmarks",
	},
	"Cwartortle": {
		"id":   uint(666801),
		"form": uint(0),
		"type": "qmarks",
	},
	"Cblastoise": {
		"id":   uint(666802),
		"form": uint(0),
		"type": "qmarks",
	},
	"Asolosis": {
		"id":   uint(666803),
		"form": uint(0),
		"type": "water",
	},
	"Aduosion": {
		"id":   uint(666804),
		"form": uint(0),
		"type": "water",
	},
	"Areuniclus": {
		"id":   uint(666805),
		"form": uint(0),
		"type": "water",
	},
	"Mars Shade": {
		"id":   uint(666806),
		"form": uint(0),
		"type": "ghost",
	},
	"Futovant": {
		"id":   uint(666807),
		"form": uint(0),
		"type": "wind",
	},
	"Suikaoni": {
		"id":   uint(666808),
		"form": uint(0),
		"type": "fighting",
	},
	"Ayawind": {
		"id":   uint(666809),
		"form": uint(0),
		"type": "wind",
	},
	"Adrifloon": {
		"id":   uint(666810),
		"form": uint(0),
		"type": "fire",
	},
	"Adrifblim": {
		"id":   uint(666811),
		"form": uint(0),
		"type": "fire",
	},
	"Adarumaka": {
		"id":   uint(666812),
		"form": uint(0),
		"type": "ice",
	},
	"Adarmanitan": {
		"id":   uint(666813),
		"form": uint(0),
		"type": "ice",
	},
	"Adarmanitan-Zen": {
		"id":   uint(666813),
		"form": uint(1),
		"type": "ice",
	},
	"Kogasella": {
		"id":   uint(666814),
		"form": uint(0),
		"type": "water",
	},
	"Sukumini": {
		"id":   uint(666815),
		"form": uint(0),
		"type": "steel",
	},
	"Clownpairy": {
		"id":   uint(666816),
		"form": uint(0),
		"type": "fairy",
	},
	"Tenshearth": {
		"id":   uint(666817),
		"form": uint(0),
		"type": "ground",
	},
	"Ligoop": {
		"id":   uint(666818),
		"form": uint(0),
		"type": "light",
	},
	"Aghost": {
		"id":   uint(666819),
		"form": uint(0),
		"type": "ghost",
	},
	"Amissingno": {
		"id":   uint(666820),
		"form": uint(0),
		"type": "qmarks",
	},
	"Bmissingno": {
		"id":   uint(666821),
		"form": uint(0),
		"type": "qmarks",
	},
	"Ymissingno": {
		"id":   uint(666822),
		"form": uint(0),
		"type": "qmarks",
	},
	"Tissoe": {
		"id":   uint(666823),
		"form": uint(0),
		"type": "paper",
	},
	"Kleenot": {
		"id":   uint(666824),
		"form": uint(0),
		"type": "paper",
	},
	"Breadamo": {
		"id":   uint(666825),
		"form": uint(0),
		"type": "food",
	},
	"Toasted": {
		"id":   uint(666826),
		"form": uint(0),
		"type": "food",
	},
	"Toater": {
		"id":   uint(666827),
		"form": uint(0),
		"type": "tech",
	},
	"Toastaem": {
		"id":   uint(666828),
		"form": uint(0),
		"type": "tech",
	},
	"Gyarth": {
		"id":   uint(666829),
		"form": uint(0),
		"type": "ground",
	},
	"Bspearow": {
		"id":   uint(666830),
		"form": uint(0),
		"type": "normal",
	},
	"Bclefairy": {
		"id":   uint(666831),
		"form": uint(0),
		"type": "normal",
	},
	"Bvenusaur": {
		"id":   uint(666832),
		"form": uint(0),
		"type": "grass",
	},
	"Bseel": {
		"id":   uint(666833),
		"form": uint(0),
		"type": "water",
	},
	"Arrowa": {
		"id":   uint(666834),
		"form": uint(0),
		"type": "qmarks",
	},
	"Ac0": {
		"id":   uint(666835),
		"form": uint(0),
		"type": "water",
	},
	"Fuidic1": {
		"id":   uint(666836),
		"form": uint(0),
		"type": "ground",
	},
	"Mechamew2": {
		"id":   uint(666837),
		"form": uint(0),
		"type": "tech",
	},
	"Frozone": {
		"id":   uint(666838),
		"form": uint(0),
		"type": "ghost",
	},
	"Hellraiser": {
		"id":   uint(666839),
		"form": uint(0),
		"type": "ghost",
	},
	"Icepick": {
		"id":   uint(666840),
		"form": uint(0),
		"type": "bug",
	},
	"Flamefetchd": {
		"id":   uint(666841),
		"form": uint(0),
		"type": "fire",
	},
	"Vshuckle": {
		"id":   uint(666842),
		"form": uint(0),
		"type": "bug",
	},
	"Darpharos": {
		"id":   uint(666843),
		"form": uint(0),
		"type": "electric",
	},
	"Slowloss": {
		"id":   uint(666844),
		"form": uint(0),
		"type": "water",
	},
	"Fracture": {
		"id":   uint(666845),
		"form": uint(0),
		"type": "bone",
	},
	"Pikared": {
		"id":   uint(666846),
		"form": uint(0),
		"type": "electric",
	},
	"Feliflame": {
		"id":   uint(666847),
		"form": uint(0),
		"type": "normal",
	},
	"Raticlaw": {
		"id":   uint(666848),
		"form": uint(0),
		"type": "normal",
	},
	"Clegnyana": {
		"id":   uint(666849),
		"form": uint(0),
		"type": "divine",
	},
	"Pikabud": {
		"id":   uint(666850),
		"form": uint(0),
		"type": "electric",
	},
	"Alenkar": {
		"id":   uint(666851),
		"form": uint(0),
		"type": "psychic",
	},
	"Dreamaster": {
		"id":   uint(666852),
		"form": uint(0),
		"type": "psychic",
	},
	"Fletchice": {
		"id":   uint(666853),
		"form": uint(0),
		"type": "normal",
	},
	"Cluckatrice": {
		"id":   uint(666854),
		"form": uint(0),
		"type": "normal",
	},
	"Roosteratrice": {
		"id":   uint(666855),
		"form": uint(0),
		"type": "dragon",
	},
	"Birttle": {
		"id":   uint(666856),
		"form": uint(0),
		"type": "glass",
	},
	"Amorvian": {
		"id":   uint(666857),
		"form": uint(0),
		"type": "glass",
	},
	"Koap": {
		"id":   uint(666858),
		"form": uint(0),
		"type": "dark",
	},
	"Hatos": {
		"id":   uint(666859),
		"form": uint(0),
		"type": "dark",
	},
	"Unhichaos": {
		"id":   uint(666860),
		"form": uint(0),
		"type": "dark",
	},
	"Suckert": {
		"id":   uint(666861),
		"form": uint(0),
		"type": "food",
	},
	"Swirlia": {
		"id":   uint(666862),
		"form": uint(0),
		"type": "food",
	},
	"Lolichap": {
		"id":   uint(666863),
		"form": uint(0),
		"type": "food",
	},
	"Flwood": {
		"id":   uint(666864),
		"form": uint(0),
		"type": "wood",
	},
	"Woopeck": {
		"id":   uint(666865),
		"form": uint(0),
		"type": "wood",
	},
	"Extinguil": {
		"id":   uint(666866),
		"form": uint(0),
		"type": "steel",
	},
	"Putout": {
		"id":   uint(666867),
		"form": uint(0),
		"type": "steel",
	},
	"Soobu": {
		"id":   uint(666868),
		"form": uint(0),
		"type": "electric",
	},
	"Sebushus": {
		"id":   uint(666869),
		"form": uint(0),
		"type": "electric",
	},
	"Magu": {
		"id":   uint(666870),
		"form": uint(0),
		"type": "magma",
	},
	"Magbru": {
		"id":   uint(666871),
		"form": uint(0),
		"type": "magma",
	},
	"Zombry": {
		"id":   uint(666872),
		"form": uint(0),
		"type": "zombie",
	},
	"Charund": {
		"id":   uint(666873),
		"form": uint(0),
		"type": "zombie",
	},
	"Spiombie": {
		"id":   uint(666874),
		"form": uint(0),
		"type": "zombie",
	},
	"Jombey": {
		"id":   uint(666875),
		"form": uint(0),
		"type": "zombie",
	},
	"Sunisk": {
		"id":   uint(666876),
		"form": uint(0),
		"type": "water",
	},
	"Solfish": {
		"id":   uint(666877),
		"form": uint(0),
		"type": "water",
	},
	"Sewan": {
		"id":   uint(666878),
		"form": uint(0),
		"type": "tech",
	},
	"Sewantor": {
		"id":   uint(666879),
		"form": uint(0),
		"type": "tech",
	},
	"Pyrork": {
		"id":   uint(666880),
		"form": uint(0),
		"type": "fire",
	},
	"Artifire": {
		"id":   uint(666881),
		"form": uint(0),
		"type": "fire",
	},
	"Artiwyrk": {
		"id":   uint(666882),
		"form": uint(0),
		"type": "fire",
	},
	"Hatakoro": {
		"id":   uint(666883),
		"form": uint(0),
		"type": "normal",
	},
	"Slinkala": {
		"id":   uint(666884),
		"form": uint(0),
		"type": "normal",
	},
	"Mimaost": {
		"id":   uint(666885),
		"form": uint(0),
		"type": "ghost",
	},
	"Ahchah": {
		"id":   uint(666886),
		"form": uint(0),
		"type": "steel",
	},
	"Kojirassin": {
		"id":   uint(666887),
		"form": uint(0),
		"type": "normal",
	},
	"Rogo": {
		"id":   uint(666888),
		"form": uint(0),
		"type": "rock",
	},
	"Rogyst": {
		"id":   uint(666889),
		"form": uint(0),
		"type": "rock",
	},
	"Spino": {
		"id":   uint(666890),
		"form": uint(0),
		"type": "plastic",
	},
	"Spinga": {
		"id":   uint(666891),
		"form": uint(0),
		"type": "plastic",
	},
	"Pikachutwo": {
		"id":   uint(666892),
		"form": uint(0),
		"type": "qmarks",
	},
	"Shinkiang": {
		"id":   uint(666893),
		"form": uint(0),
		"type": "psychic",
	},
	"Soundagon": {
		"id":   uint(666894),
		"form": uint(0),
		"type": "dragon",
	},
	"Sonardrake": {
		"id":   uint(666895),
		"form": uint(0),
		"type": "dragon",
	},
	"Carplosion": {
		"id":   uint(666896),
		"form": uint(0),
		"type": "fabric",
	},
	"Carpetbomb": {
		"id":   uint(666897),
		"form": uint(0),
		"type": "fabric",
	},
	"Finish": {
		"id":   uint(666898),
		"form": uint(0),
		"type": "water",
	},
	"Bludgreigon": {
		"id":   uint(666899),
		"form": uint(0),
		"type": "water",
	},
	"Salember": {
		"id":   uint(666900),
		"form": uint(0),
		"type": "fire",
	},
	"Axofire": {
		"id":   uint(666901),
		"form": uint(0),
		"type": "dragon",
	},
	"Searamander": {
		"id":   uint(666902),
		"form": uint(0),
		"type": "dragon",
	},
	"Squinky": {
		"id":   uint(666903),
		"form": uint(0),
		"type": "water",
	},
	"Inkuid": {
		"id":   uint(666904),
		"form": uint(0),
		"type": "water",
	},
	"Lemmong": {
		"id":   uint(666905),
		"form": uint(0),
		"type": "food",
	},
	"Lemmongade": {
		"id":   uint(666906),
		"form": uint(0),
		"type": "food",
	},
	"Voodem": {
		"id":   uint(666907),
		"form": uint(0),
		"type": "magic",
	},
	"Calculus": {
		"id":   uint(666908),
		"form": uint(0),
		"type": "tech",
	},
	"Calculusus": {
		"id":   uint(666909),
		"form": uint(0),
		"type": "tech",
	},
	"Maskett": {
		"id":   uint(666910),
		"form": uint(0),
		"type": "steel",
	},
	"Visagelle": {
		"id":   uint(666911),
		"form": uint(0),
		"type": "steel",
	},
	"Tardida": {
		"id":   uint(666912),
		"form": uint(0),
		"type": "bug",
	},
	"Blimy": {
		"id":   uint(666913),
		"form": uint(0),
		"type": "steel",
	},
	"Hidenblim": {
		"id":   uint(666914),
		"form": uint(0),
		"type": "steel",
	},
	"Eikishiki": {
		"id":   uint(666915),
		"form": uint(0),
		"type": "divine",
	},
	"Trarsile": {
		"id":   uint(666916),
		"form": uint(0),
		"type": "poison",
	},
	"Rogodo": {
		"id":   uint(666917),
		"form": uint(0),
		"type": "divine",
	},
	"Marsarrior": {
		"id":   uint(666918),
		"form": uint(0),
		"type": "divine",
	},
	"Venusess": {
		"id":   uint(666919),
		"form": uint(0),
		"type": "divine",
	},
	"Gapkari": {
		"id":   uint(666920),
		"form": uint(0),
		"type": "psychic",
	},
	"Gyaradeth": {
		"id":   uint(666921),
		"form": uint(0),
		"type": "water",
	},
	"Babirys": {
		"id":   uint(666922),
		"form": uint(0),
		"type": "steel",
	},
	"Suriv": {
		"id":   uint(666923),
		"form": uint(0),
		"type": "steel",
	},
	"Shoison": {
		"id":   uint(666924),
		"form": uint(0),
		"type": "electric",
	},
	"Battacid": {
		"id":   uint(666925),
		"form": uint(0),
		"type": "electric",
	},
	"Aliceoll": {
		"id":   uint(666926),
		"form": uint(0),
		"type": "steel",
	},
	"Yuukawer": {
		"id":   uint(666927),
		"form": uint(0),
		"type": "grass",
	},
	"Hongling": {
		"id":   uint(666928),
		"form": uint(0),
		"type": "fighting",
	},
	"Chengami": {
		"id":   uint(666929),
		"form": uint(0),
		"type": "wind",
	},
	"Arachoe": {
		"id":   uint(666930),
		"form": uint(0),
		"type": "tech",
	},
	"Backscorpion": {
		"id":   uint(666931),
		"form": uint(0),
		"type": "tech",
	},
	"Blolocute": {
		"id":   uint(666932),
		"form": uint(0),
		"type": "electric",
	},
	"Bloculo": {
		"id":   uint(666933),
		"form": uint(0),
		"type": "electric",
	},
	"Blolossod": {
		"id":   uint(666934),
		"form": uint(0),
		"type": "electric",
	},
	"Jawetech": {
		"id":   uint(666935),
		"form": uint(0),
		"type": "water",
	},
	"Humowl": {
		"id":   uint(666936),
		"form": uint(0),
		"type": "fear",
	},
	"Mothmowl": {
		"id":   uint(666937),
		"form": uint(0),
		"type": "fear",
	},
	"Windoh": {
		"id":   uint(666938),
		"form": uint(0),
		"type": "glass",
	},
	"Winblow": {
		"id":   uint(666939),
		"form": uint(0),
		"type": "glass",
	},
	"Winpane": {
		"id":   uint(666940),
		"form": uint(0),
		"type": "glass",
	},
	"Peekindow": {
		"id":   uint(666941),
		"form": uint(0),
		"type": "glass",
	},
	"Bunn": {
		"id":   uint(666942),
		"form": uint(0),
		"type": "food",
	},
	"Draguette": {
		"id":   uint(666943),
		"form": uint(0),
		"type": "food",
	},
	"Subwaydon": {
		"id":   uint(666944),
		"form": uint(0),
		"type": "food",
	},
	"Predorb": {
		"id":   uint(666945),
		"form": uint(0),
		"type": "crystal",
	},
	"Seerorb": {
		"id":   uint(666946),
		"form": uint(0),
		"type": "crystal",
	},
	"Prophorb": {
		"id":   uint(666947),
		"form": uint(0),
		"type": "crystal",
	},
	"Tapistle": {
		"id":   uint(666948),
		"form": uint(0),
		"type": "plastic",
	},
	"Blopire": {
		"id":   uint(666949),
		"form": uint(0),
		"type": "plastic",
	},
	"Larvick": {
		"id":   uint(666950),
		"form": uint(0),
		"type": "bug",
	},
	"Tickood": {
		"id":   uint(666951),
		"form": uint(0),
		"type": "bug",
	},
	"Eirinurse": {
		"id":   uint(666952),
		"form": uint(0),
		"type": "cosmic",
	},
	"Smoodie": {
		"id":   uint(666953),
		"form": uint(0),
		"type": "ice",
	},
	"Coocamodie": {
		"id":   uint(666954),
		"form": uint(0),
		"type": "ice",
	},
	"Rockandy": {
		"id":   uint(666955),
		"form": uint(0),
		"type": "crystal",
	},
	"Crystandy": {
		"id":   uint(666956),
		"form": uint(0),
		"type": "crystal",
	},
	"Cracandy": {
		"id":   uint(666957),
		"form": uint(0),
		"type": "crystal",
	},
	"Seeare": {
		"id":   uint(666958),
		"form": uint(0),
		"type": "wood",
	},
	"Feasume": {
		"id":   uint(666959),
		"form": uint(0),
		"type": "wood",
	},
	"Ranox": {
		"id":   uint(666960),
		"form": uint(0),
		"type": "fire",
	},
	"Remilire": {
		"id":   uint(666961),
		"form": uint(0),
		"type": "dark",
	},
	"Marpet": {
		"id":   uint(666962),
		"form": uint(0),
		"type": "fabric",
	},
	"Flarpet": {
		"id":   uint(666963),
		"form": uint(0),
		"type": "fabric",
	},
	"Yuyukhost": {
		"id":   uint(666964),
		"form": uint(0),
		"type": "ghost",
	},
	"Ikugae": {
		"id":   uint(666965),
		"form": uint(0),
		"type": "electric",
	},
	"Yumemerry": {
		"id":   uint(666966),
		"form": uint(0),
		"type": "magic",
	},
	"Marsybie": {
		"id":   uint(666967),
		"form": uint(0),
		"type": "zombie",
	},
	"Honarsal": {
		"id":   uint(666968),
		"form": uint(0),
		"type": "meme",
	},
	"Torturlens": {
		"id":   uint(666969),
		"form": uint(0),
		"type": "tech",
	},
	"Maidcoresh": {
		"id":   uint(666970),
		"form": uint(0),
		"type": "tech",
	},
	"Maidcoresh-Blade": {
		"id":   uint(666970),
		"form": uint(1),
		"type": "tech",
	},
	"Flandryre": {
		"id":   uint(666971),
		"form": uint(0),
		"type": "blood",
	},
	"Sarieangel": {
		"id":   uint(666972),
		"form": uint(0),
		"type": "chaos",
	},
	"Sarieangel-Zen": {
		"id":   uint(666972),
		"form": uint(1),
		"type": "chaos",
	},
	"Starphire": {
		"id":   uint(666973),
		"form": uint(0),
		"type": "fairy",
	},
	"Artania": {
		"id":   uint(666974),
		"form": uint(0),
		"type": "divine",
	},
	"Apollolar": {
		"id":   uint(666975),
		"form": uint(0),
		"type": "divine",
	},
	"Marsortured": {
		"id":   uint(666976),
		"form": uint(0),
		"type": "cosmic",
	},
	"Shouger": {
		"id":   uint(666977),
		"form": uint(0),
		"type": "light",
	},
	"Demiurge": {
		"id":   uint(666978),
		"form": uint(0),
		"type": "divine",
	},
	"Komacheaper": {
		"id":   uint(666979),
		"form": uint(0),
		"type": "water",
	},
	"Doveart": {
		"id":   uint(666980),
		"form": uint(0),
		"type": "heart",
	},
	"Snakeart": {
		"id":   uint(666981),
		"form": uint(0),
		"type": "heart",
	},
	"Dnaby": {
		"id":   uint(666982),
		"form": uint(0),
		"type": "blood",
	},
	"Dnaelix": {
		"id":   uint(666983),
		"form": uint(0),
		"type": "blood",
	},
	"Dnanium": {
		"id":   uint(666984),
		"form": uint(0),
		"type": "blood",
	},
	"Supaninteno": {
		"id":   uint(666985),
		"form": uint(0),
		"type": "tech",
	},
	"Cubegamo": {
		"id":   uint(666986),
		"form": uint(0),
		"type": "tech",
	},
	"Wiilly": {
		"id":   uint(666987),
		"form": uint(0),
		"type": "tech",
	},
	"Threediesh": {
		"id":   uint(666988),
		"form": uint(0),
		"type": "tech",
	},
	"Crawdrake": {
		"id":   uint(666989),
		"form": uint(0),
		"type": "dragon",
	},
	"Lightip": {
		"id":   uint(666990),
		"form": uint(0),
		"type": "steel",
	},
	"Flameup": {
		"id":   uint(666991),
		"form": uint(0),
		"type": "steel",
	},
	"Seijaverse": {
		"id":   uint(666992),
		"form": uint(0),
		"type": "poison",
	},
	"Shingyokorb": {
		"id":   uint(666993),
		"form": uint(0),
		"type": "normal",
	},
	"Saibaltah": {
		"id":   uint(666994),
		"form": uint(0),
		"type": "dragon",
	},
	"Gillesaster": {
		"id":   uint(666995),
		"form": uint(0),
		"type": "magic",
	},
	"Iskanrider": {
		"id":   uint(666996),
		"form": uint(0),
		"type": "fighting",
	},
	"Diamurdancer": {
		"id":   uint(666997),
		"form": uint(0),
		"type": "heart",
	},
	"Angrymanjew": {
		"id":   uint(666998),
		"form": uint(0),
		"type": "qmarks",
	},
	"Angrymanjew-Zen": {
		"id":   uint(666998),
		"form": uint(1),
		"type": "qmarks",
	},
	"Merquick": {
		"id":   uint(666999),
		"form": uint(0),
		"type": "divine",
	},
	"Nerober": {
		"id":   uint(667000),
		"form": uint(0),
		"type": "light",
	},
	"Hundssassin": {
		"id":   uint(667001),
		"form": uint(0),
		"type": "dark",
	},
	"Himekoto": {
		"id":   uint(667002),
		"form": uint(0),
		"type": "normal",
	},
	"Archlanta": {
		"id":   uint(667003),
		"form": uint(0),
		"type": "divine",
	},
	"Sudorokku": {
		"id":   uint(667004),
		"form": uint(0),
		"type": "grass",
	},
	"Traino": {
		"id":   uint(667005),
		"form": uint(0),
		"type": "steel",
	},
	"Strain": {
		"id":   uint(667006),
		"form": uint(0),
		"type": "steel",
	},
	"Electrain": {
		"id":   uint(667007),
		"form": uint(0),
		"type": "steel",
	},
	"Shinkantrain": {
		"id":   uint(667008),
		"form": uint(0),
		"type": "steel",
	},
	"Konngarrior": {
		"id":   uint(667009),
		"form": uint(0),
		"type": "steel",
	},
	"Starwi": {
		"id":   uint(667010),
		"form": uint(0),
		"type": "water",
	},
	"Chiyailor": {
		"id":   uint(667011),
		"form": uint(0),
		"type": "water",
	},
	"Abomizzard": {
		"id":   uint(667012),
		"form": uint(0),
		"type": "grass",
	},
	"Mmawile": {
		"id":   uint(667013),
		"form": uint(0),
		"type": "steel",
	},
	"Aabsol": {
		"id":   uint(667014),
		"form": uint(0),
		"type": "dark",
	},
	"Bbanette": {
		"id":   uint(667015),
		"form": uint(0),
		"type": "ghost",
	},
	"Rumiarkness": {
		"id":   uint(667016),
		"form": uint(0),
		"type": "dark",
	},
	"Suwakrog": {
		"id":   uint(667017),
		"form": uint(0),
		"type": "ground",
	},
	"Kanakake": {
		"id":   uint(667018),
		"form": uint(0),
		"type": "wind",
	},
	"Daiyairy": {
		"id":   uint(667019),
		"form": uint(0),
		"type": "fairy",
	},
	"Koakuma": {
		"id":   uint(667020),
		"form": uint(0),
		"type": "chaos",
	},
	"Yorihimoon": {
		"id":   uint(667021),
		"form": uint(0),
		"type": "steel",
	},
	"Toyohimoon": {
		"id":   uint(667022),
		"form": uint(0),
		"type": "water",
	},
	"Elisemon": {
		"id":   uint(667023),
		"form": uint(0),
		"type": "dark",
	},
	"Yuugenmagan": {
		"id":   uint(667024),
		"form": uint(0),
		"type": "electric",
	},
	"Sworesh": {
		"id":   uint(667025),
		"form": uint(0),
		"type": "fighting",
	},
	"Marlencin": {
		"id":   uint(667026),
		"form": uint(0),
		"type": "fighting",
	},
	"Grenir": {
		"id":   uint(667027),
		"form": uint(0),
		"type": "magic",
	},
	"Marvelamp": {
		"id":   uint(667028),
		"form": uint(0),
		"type": "magic",
	},
	"Tusk": {
		"id":   uint(667029),
		"form": uint(0),
		"type": "fighting",
	},
	"Tusk Act 3": {
		"id":   uint(667030),
		"form": uint(0),
		"type": "fighting",
	},
	"Mystiarrow": {
		"id":   uint(667031),
		"form": uint(0),
		"type": "dark",
	},
	"Bowser": {
		"id":   uint(667032),
		"form": uint(0),
		"type": "fire",
	},
	"Risatsuki": {
		"id":   uint(667033),
		"form": uint(0),
		"type": "fire",
	},
	"Lagiaunder": {
		"id":   uint(667034),
		"form": uint(0),
		"type": "water",
	},
	"Sunnilk": {
		"id":   uint(667035),
		"form": uint(0),
		"type": "fairy",
	},
	"Lunild": {
		"id":   uint(667036),
		"form": uint(0),
		"type": "fairy",
	},
	"Bootporygon": {
		"id":   uint(667037),
		"form": uint(0),
		"type": "normal",
	},
	"Dlcorygon": {
		"id":   uint(667038),
		"form": uint(0),
		"type": "normal",
	},
	"Tewiabbit": {
		"id":   uint(667039),
		"form": uint(0),
		"type": "normal",
	},
	"Ranmoose": {
		"id":   uint(667040),
		"form": uint(0),
		"type": "normal",
	},
	"Kyotsufall": {
		"id":   uint(667041),
		"form": uint(0),
		"type": "water",
	},
	"Reikuto": {
		"id":   uint(667042),
		"form": uint(0),
		"type": "normal",
	},
	"Junkure": {
		"id":   uint(667043),
		"form": uint(0),
		"type": "divine",
	},
	"Pendrasp": {
		"id":   uint(667044),
		"form": uint(0),
		"type": "flying",
	},
	"Hillipod": {
		"id":   uint(667045),
		"form": uint(0),
		"type": "flying",
	},
	"Axelasp": {
		"id":   uint(667046),
		"form": uint(0),
		"type": "flying",
	},
	"Medicholy": {
		"id":   uint(667047),
		"form": uint(0),
		"type": "poison",
	},
	"Wrigglight": {
		"id":   uint(667048),
		"form": uint(0),
		"type": "bug",
	},
	"Kaguyainbow": {
		"id":   uint(667049),
		"form": uint(0),
		"type": "cosmic",
	},
	"Mokanou": {
		"id":   uint(667050),
		"form": uint(0),
		"type": "fire",
	},
	"Keinistory": {
		"id":   uint(667051),
		"form": uint(0),
		"type": "steel",
	},
	"Keinistory-Zen": {
		"id":   uint(667051),
		"form": uint(1),
		"type": "steel",
	},
	"Swastitan": {
		"id":   uint(667052),
		"form": uint(0),
		"type": "qmarks",
	},
	"Hatatoto": {
		"id":   uint(667053),
		"form": uint(0),
		"type": "psychic",
	},
	"Lilight": {
		"id":   uint(667054),
		"form": uint(0),
		"type": "normal",
	},
	"Lilack": {
		"id":   uint(667055),
		"form": uint(0),
		"type": "dark",
	},
	"Lettyice": {
		"id":   uint(667056),
		"form": uint(0),
		"type": "ice",
	},
	"Lunasound": {
		"id":   uint(667057),
		"form": uint(0),
		"type": "ghost",
	},
	"Lyricound": {
		"id":   uint(667058),
		"form": uint(0),
		"type": "ghost",
	},
	"Merlinound": {
		"id":   uint(667059),
		"form": uint(0),
		"type": "ghost",
	},
	"Ryutoli": {
		"id":   uint(667060),
		"form": uint(0),
		"type": "light",
	},
	"Tigrampage": {
		"id":   uint(667061),
		"form": uint(0),
		"type": "normal",
	},
	"Mikorince": {
		"id":   uint(667062),
		"form": uint(0),
		"type": "divine",
	},
	"Byakuraint": {
		"id":   uint(667063),
		"form": uint(0),
		"type": "fighting",
	},
	"Nitorech": {
		"id":   uint(667064),
		"form": uint(0),
		"type": "water",
	},
	"Kikurisk": {
		"id":   uint(667065),
		"form": uint(0),
		"type": "rock",
	},
	"Khezear": {
		"id":   uint(667066),
		"form": uint(0),
		"type": "electric",
	},
	"Veloprey": {
		"id":   uint(667067),
		"form": uint(0),
		"type": "normal",
	},
	"Velodrome": {
		"id":   uint(667068),
		"form": uint(0),
		"type": "normal",
	},
	"Diablex": {
		"id":   uint(667069),
		"form": uint(0),
		"type": "ground",
	},
	"Shachich": {
		"id":   uint(667070),
		"form": uint(0),
		"type": "dragon",
	},
	"Utsuclear": {
		"id":   uint(667071),
		"form": uint(0),
		"type": "nuclear",
	},
	"Communarve": {
		"id":   uint(667072),
		"form": uint(0),
		"type": "steel",
	},
	"Ellythe": {
		"id":   uint(667073),
		"form": uint(0),
		"type": "steel",
	},
	"Mugetseam": {
		"id":   uint(667074),
		"form": uint(0),
		"type": "psychic",
	},
	"Gengetseam": {
		"id":   uint(667075),
		"form": uint(0),
		"type": "psychic",
	},
	"Dalo": {
		"id":   uint(667076),
		"form": uint(0),
		"type": "fire",
	},
	"Balremon": {
		"id":   uint(667077),
		"form": uint(0),
		"type": "fire",
	},
	"Bewdragon": {
		"id":   uint(667078),
		"form": uint(0),
		"type": "dragon",
	},
	"Plesiocean": {
		"id":   uint(667079),
		"form": uint(0),
		"type": "water",
	},
	"Spacird": {
		"id":   uint(667080),
		"form": uint(0),
		"type": "cosmic",
	},
	"Tosporygon": {
		"id":   uint(667081),
		"form": uint(0),
		"type": "divine",
	},
	"Alactrode": {
		"id":   uint(667082),
		"form": uint(0),
		"type": "electric",
	},
	"Gamble": {
		"id":   uint(667083),
		"form": uint(0),
		"type": "dark",
	},
	"Pinsire": {
		"id":   uint(667084),
		"form": uint(0),
		"type": "bug",
	},
	"Cliffracer": {
		"id":   uint(667085),
		"form": uint(0),
		"type": "flying",
	},
	"Graviore": {
		"id":   uint(667086),
		"form": uint(0),
		"type": "steel",
	},
	"Nazrat": {
		"id":   uint(667087),
		"form": uint(0),
		"type": "normal",
	},
	"Ichiran": {
		"id":   uint(667088),
		"form": uint(0),
		"type": "fighting",
	},
	"Doreamy": {
		"id":   uint(667089),
		"form": uint(0),
		"type": "psychic",
	},
	"Bowsette": {
		"id":   uint(667090),
		"form": uint(0),
		"type": "dragon",
	},
	"Lagannech": {
		"id":   uint(667091),
		"form": uint(0),
		"type": "tech",
	},
	"Gurrenlagannech": {
		"id":   uint(667092),
		"form": uint(0),
		"type": "tech",
	},
	"Dry Bowser": {
		"id":   uint(667093),
		"form": uint(0),
		"type": "bone",
	},
	"Dark Bowser": {
		"id":   uint(667094),
		"form": uint(0),
		"type": "fire",
	},
	"Retro Bowser": {
		"id":   uint(667095),
		"form": uint(0),
		"type": "fire",
	},
	"Megadragonbowser": {
		"id":   uint(667096),
		"form": uint(0),
		"type": "fire",
	},
	"Metal Bowser": {
		"id":   uint(667097),
		"form": uint(0),
		"type": "fire",
	},
	"Meowser": {
		"id":   uint(667098),
		"form": uint(0),
		"type": "fire",
	},
	"Parsenvy": {
		"id":   uint(667099),
		"form": uint(0),
		"type": "heart",
	},
	"Kisumket": {
		"id":   uint(667100),
		"form": uint(0),
		"type": "fire",
	},
	"Fishember": {
		"id":   uint(667101),
		"form": uint(0),
		"type": "water",
	},
	"Searish": {
		"id":   uint(667102),
		"form": uint(0),
		"type": "water",
	},
	"Mamizouki": {
		"id":   uint(667103),
		"form": uint(0),
		"type": "normal",
	},
	"Satorind": {
		"id":   uint(667104),
		"form": uint(0),
		"type": "heart",
	},
	"Desound": {
		"id":   uint(667105),
		"form": uint(0),
		"type": "normal",
	},
	"Descreech": {
		"id":   uint(667106),
		"form": uint(0),
		"type": "chaos",
	},
	"Yamamease": {
		"id":   uint(667107),
		"form": uint(0),
		"type": "bug",
	},
	"Gasheleton": {
		"id":   uint(667108),
		"form": uint(0),
		"type": "ghost",
	},
	"Chillizar": {
		"id":   uint(667109),
		"form": uint(0),
		"type": "fire",
	},
	"Frozark": {
		"id":   uint(667110),
		"form": uint(0),
		"type": "fire",
	},
	"Frizard": {
		"id":   uint(667111),
		"form": uint(0),
		"type": "fire",
	},
	"Jackask": {
		"id":   uint(667112),
		"form": uint(0),
		"type": "fear",
	},
	"Catace": {
		"id":   uint(667113),
		"form": uint(0),
		"type": "chaos",
	},
	"Catface": {
		"id":   uint(667114),
		"form": uint(0),
		"type": "chaos",
	},
	"Phobeon": {
		"id":   uint(667115),
		"form": uint(0),
		"type": "fear",
	},
	"Skullask": {
		"id":   uint(667116),
		"form": uint(0),
		"type": "fear",
	},
	"Shizautumn": {
		"id":   uint(667117),
		"form": uint(0),
		"type": "grass",
	},
	"Minoravest": {
		"id":   uint(667118),
		"form": uint(0),
		"type": "grass",
	},
	"Gettarobo": {
		"id":   uint(667119),
		"form": uint(0),
		"type": "tech",
	},
	"Jestevil": {
		"id":   uint(667120),
		"form": uint(0),
		"type": "chaos",
	},
	"Shincityo": {
		"id":   uint(667121),
		"form": uint(0),
		"type": "water",
	},
	"Msableye": {
		"id":   uint(667122),
		"form": uint(0),
		"type": "dark",
	},
	"Msteelix": {
		"id":   uint(667123),
		"form": uint(0),
		"type": "steel",
	},
	"Mcsteelix": {
		"id":   uint(667124),
		"form": uint(0),
		"type": "steel",
	},
	"Amimejr": {
		"id":   uint(667125),
		"form": uint(0),
		"type": "psychic",
	},
	"Amrmime": {
		"id":   uint(667126),
		"form": uint(0),
		"type": "psychic",
	},
	"Apurrloin": {
		"id":   uint(667127),
		"form": uint(0),
		"type": "normal",
	},
	"Aliepard": {
		"id":   uint(667128),
		"form": uint(0),
		"type": "normal",
	},
	"Acacnea": {
		"id":   uint(667129),
		"form": uint(0),
		"type": "normal",
	},
	"Acacturne": {
		"id":   uint(667130),
		"form": uint(0),
		"type": "normal",
	},
	"Snowtorb": {
		"id":   uint(667131),
		"form": uint(0),
		"type": "ice",
	},
	"Snowtrode": {
		"id":   uint(667132),
		"form": uint(0),
		"type": "ice",
	},
	"Atreecko": {
		"id":   uint(667133),
		"form": uint(0),
		"type": "fairy",
	},
	"Agrovyle": {
		"id":   uint(667134),
		"form": uint(0),
		"type": "fairy",
	},
	"Asceptile": {
		"id":   uint(667135),
		"form": uint(0),
		"type": "fairy",
	},
	"Atorchic": {
		"id":   uint(667136),
		"form": uint(0),
		"type": "rock",
	},
	"Acombusken": {
		"id":   uint(667137),
		"form": uint(0),
		"type": "rock",
	},
	"Ablaziken": {
		"id":   uint(667138),
		"form": uint(0),
		"type": "rock",
	},
	"Amudkip": {
		"id":   uint(667139),
		"form": uint(0),
		"type": "fighting",
	},
	"Amarshtomp": {
		"id":   uint(667140),
		"form": uint(0),
		"type": "fighting",
	},
	"Aswampert": {
		"id":   uint(667141),
		"form": uint(0),
		"type": "fighting",
	},
	"Aves": {
		"id":   uint(667142),
		"form": uint(0),
		"type": "normal",
	},
	"Laprince": {
		"id":   uint(667143),
		"form": uint(0),
		"type": "water",
	},
	"Firefree": {
		"id":   uint(667144),
		"form": uint(0),
		"type": "bug",
	},
	"Pegazeus": {
		"id":   uint(667145),
		"form": uint(0),
		"type": "electric",
	},
	"Shockra": {
		"id":   uint(667146),
		"form": uint(0),
		"type": "electric",
	},
	"Barkanite": {
		"id":   uint(667147),
		"form": uint(0),
		"type": "electric",
	},
	"Momiawoo": {
		"id":   uint(667148),
		"form": uint(0),
		"type": "steel",
	},
	"Kyecho": {
		"id":   uint(667149),
		"form": uint(0),
		"type": "sound",
	},
	"Replicore": {
		"id":   uint(667150),
		"form": uint(0),
		"type": "meme",
	},
	"Thwompo": {
		"id":   uint(667151),
		"form": uint(0),
		"type": "rock",
	},
	"Hbuneary": {
		"id":   uint(667152),
		"form": uint(0),
		"type": "normal",
	},
	"Hlopunny": {
		"id":   uint(667153),
		"form": uint(0),
		"type": "normal",
	},
	"Yuugihol": {
		"id":   uint(667154),
		"form": uint(0),
		"type": "fighting",
	},
	"Bigsquatch": {
		"id":   uint(667155),
		"form": uint(0),
		"type": "fear",
	},
	"Nessieous": {
		"id":   uint(667156),
		"form": uint(0),
		"type": "fear",
	},
	"Kart Bowser": {
		"id":   uint(667157),
		"form": uint(0),
		"type": "fire",
	},
	"Holeecow": {
		"id":   uint(667158),
		"form": uint(0),
		"type": "normal",
	},
	"Acidoud": {
		"id":   uint(667159),
		"form": uint(0),
		"type": "wind",
	},
	"Clounight": {
		"id":   uint(667160),
		"form": uint(0),
		"type": "wind",
	},
	"Coyot": {
		"id":   uint(667161),
		"form": uint(0),
		"type": "ground",
	},
	"Coyotears": {
		"id":   uint(667162),
		"form": uint(0),
		"type": "ground",
	},
	"Levelcap": {
		"id":   uint(667163),
		"form": uint(0),
		"type": "fabric",
	},
	"Luckat": {
		"id":   uint(667164),
		"form": uint(0),
		"type": "magic",
	},
	"Christree": {
		"id":   uint(667165),
		"form": uint(0),
		"type": "grass",
	},
	"Mario": {
		"id":   uint(667166),
		"form": uint(0),
		"type": "normal",
	},
	"Marthlord": {
		"id":   uint(667167),
		"form": uint(0),
		"type": "steel",
	},
	"Royboy": {
		"id":   uint(667168),
		"form": uint(0),
		"type": "steel",
	},
	"Lucinord": {
		"id":   uint(667169),
		"form": uint(0),
		"type": "steel",
	},
	"Balloom": {
		"id":   uint(667170),
		"form": uint(0),
		"type": "rubber",
	},
	"Wyrmados": {
		"id":   uint(667171),
		"form": uint(0),
		"type": "water",
	},
	"Corrupted Sol": {
		"id":   uint(667172),
		"form": uint(0),
		"type": "divine",
	},
	"Isthis": {
		"id":   uint(667173),
		"form": uint(0),
		"type": "chaos",
	},
	"Fbanette": {
		"id":   uint(667174),
		"form": uint(0),
		"type": "ghost",
	},
	"Indianelephant": {
		"id":   uint(667175),
		"form": uint(0),
		"type": "normal",
	},
	"Solaress": {
		"id":   uint(667176),
		"form": uint(0),
		"type": "fire",
	},
	"Sanzucrow": {
		"id":   uint(667177),
		"form": uint(0),
		"type": "flying",
	},
	"Suncrab": {
		"id":   uint(667178),
		"form": uint(0),
		"type": "magma",
	},
	"Ohmega": {
		"id":   uint(667179),
		"form": uint(0),
		"type": "electric",
	},
	"Ittybatty": {
		"id":   uint(667180),
		"form": uint(0),
		"type": "poison",
	},
	"Gorochu": {
		"id":   uint(667181),
		"form": uint(0),
		"type": "electric",
	},
	"Mordrayal": {
		"id":   uint(667182),
		"form": uint(0),
		"type": "blood",
	},
	"Berserfran": {
		"id":   uint(667183),
		"form": uint(0),
		"type": "electric",
	},
	"Tamamon": {
		"id":   uint(667184),
		"form": uint(0),
		"type": "magic",
	},
	"Karnancer": {
		"id":   uint(667185),
		"form": uint(0),
		"type": "fire",
	},
	"Chargra": {
		"id":   uint(667186),
		"form": uint(0),
		"type": "grass",
	},
	"Chargradon": {
		"id":   uint(667187),
		"form": uint(0),
		"type": "grass",
	},
	"Chargroar": {
		"id":   uint(667188),
		"form": uint(0),
		"type": "grass",
	},
	"Astolftrap": {
		"id":   uint(667189),
		"form": uint(0),
		"type": "fairy",
	},
	"Twotow": {
		"id":   uint(667190),
		"form": uint(0),
		"type": "steel",
	},
	"Ashgaaaaaa": {
		"id":   uint(667191),
		"form": uint(0),
		"type": "fairy",
	},
	"Infinistake": {
		"id":   uint(667192),
		"form": uint(0),
		"type": "dragon",
	},
	"Elias": {
		"id":   uint(667193),
		"form": uint(0),
		"type": "divine",
	},
	"Jeannearc": {
		"id":   uint(667194),
		"form": uint(0),
		"type": "divine",
	},
	"Jalter": {
		"id":   uint(667195),
		"form": uint(0),
		"type": "dragon",
	},
	"Cactoos": {
		"id":   uint(667196),
		"form": uint(0),
		"type": "grass",
	},
	"Golpsyduck": {
		"id":   uint(667197),
		"form": uint(0),
		"type": "water",
	},
	"Matriawak": {
		"id":   uint(667198),
		"form": uint(0),
		"type": "bone",
	},
	"Buu": {
		"id":   uint(667199),
		"form": uint(0),
		"type": "ice",
	},
	"Gyao": {
		"id":   uint(667200),
		"form": uint(0),
		"type": "normal",
	},
	"Gyaoon": {
		"id":   uint(667201),
		"form": uint(0),
		"type": "normal",
	},
	"Eleko": {
		"id":   uint(667202),
		"form": uint(0),
		"type": "ground",
	},
	"Squi": {
		"id":   uint(667203),
		"form": uint(0),
		"type": "water",
	},
	"Squink": {
		"id":   uint(667204),
		"form": uint(0),
		"type": "water",
	},
	"Cheepa": {
		"id":   uint(667205),
		"form": uint(0),
		"type": "water",
	},
	"Beeta": {
		"id":   uint(667206),
		"form": uint(0),
		"type": "water",
	},
	"Wartoise": {
		"id":   uint(667207),
		"form": uint(0),
		"type": "water",
	},
	"Blastyke": {
		"id":   uint(667208),
		"form": uint(0),
		"type": "water",
	},
	"Sumorog": {
		"id":   uint(667209),
		"form": uint(0),
		"type": "water",
	},
	"Kungfrog": {
		"id":   uint(667210),
		"form": uint(0),
		"type": "water",
	},
	"Wadgon": {
		"id":   uint(667211),
		"form": uint(0),
		"type": "water",
	},
	"Waterake": {
		"id":   uint(667212),
		"form": uint(0),
		"type": "water",
	},
	"Wateragon": {
		"id":   uint(667213),
		"form": uint(0),
		"type": "water",
	},
	"Vladpire": {
		"id":   uint(667214),
		"form": uint(0),
		"type": "blood",
	},
	"Chirontaur": {
		"id":   uint(667215),
		"form": uint(0),
		"type": "divine",
	},
	"Le Loiber": {
		"id":   uint(667216),
		"form": uint(0),
		"type": "dragon",
	},
	"Bersertacus": {
		"id":   uint(667217),
		"form": uint(0),
		"type": "fighting",
	},
	"Jackripper": {
		"id":   uint(667218),
		"form": uint(0),
		"type": "dark",
	},
	"Avicecaster": {
		"id":   uint(667219),
		"form": uint(0),
		"type": "magic",
	},
	"Achillesder": {
		"id":   uint(667220),
		"form": uint(0),
		"type": "fighting",
	},
	"Shakespearster": {
		"id":   uint(667221),
		"form": uint(0),
		"type": "normal",
	},
	"Semiramissin": {
		"id":   uint(667222),
		"form": uint(0),
		"type": "magic",
	},
	"Siegfriedaber": {
		"id":   uint(667223),
		"form": uint(0),
		"type": "dragon",
	},
	"Amakuruler": {
		"id":   uint(667224),
		"form": uint(0),
		"type": "divine",
	},
	"Gillesaber": {
		"id":   uint(667225),
		"form": uint(0),
		"type": "steel",
	},
	"Arasharcher": {
		"id":   uint(667226),
		"form": uint(0),
		"type": "normal",
	},
	"Georgiorider": {
		"id":   uint(667227),
		"form": uint(0),
		"type": "divine",
	},
	"Elizidol": {
		"id":   uint(667228),
		"form": uint(0),
		"type": "dragon",
	},
	"Bunyanserk": {
		"id":   uint(667229),
		"form": uint(0),
		"type": "wood",
	},
	"Euryarcher": {
		"id":   uint(667230),
		"form": uint(0),
		"type": "divine",
	},
	"Sthenassin": {
		"id":   uint(667231),
		"form": uint(0),
		"type": "divine",
	},
	"Drakrider": {
		"id":   uint(667232),
		"form": uint(0),
		"type": "water",
	},
	"Tamamocat": {
		"id":   uint(667233),
		"form": uint(0),
		"type": "food",
	},
	"Enkiduvine": {
		"id":   uint(667234),
		"form": uint(0),
		"type": "divine",
	},
	"Hanscaster": {
		"id":   uint(667235),
		"form": uint(0),
		"type": "paper",
	},
	"Robinarcher": {
		"id":   uint(667236),
		"form": uint(0),
		"type": "grass",
	},
	"Cryongi": {
		"id":   uint(667237),
		"form": uint(0),
		"type": "chaos",
	},
	"Medbrider": {
		"id":   uint(667238),
		"form": uint(0),
		"type": "fairy",
	},
	"Hasserenity": {
		"id":   uint(667239),
		"form": uint(0),
		"type": "dark",
	},
	"Tikisque": {
		"id":   uint(667240),
		"form": uint(0),
		"type": "grass",
	},
	"Nurserhyme": {
		"id":   uint(667241),
		"form": uint(0),
		"type": "fairy",
	},
	"Annesque": {
		"id":   uint(667242),
		"form": uint(0),
		"type": "steel",
	},
	"Comfstrictor": {
		"id":   uint(667243),
		"form": uint(0),
		"type": "fabric",
	},
	"Lancartoria": {
		"id":   uint(667244),
		"form": uint(0),
		"type": "dragon",
	},
	"Lu Buserker": {
		"id":   uint(667245),
		"form": uint(0),
		"type": "fighting",
	},
	"Kiarayoin": {
		"id":   uint(667246),
		"form": uint(0),
		"type": "normal",
	},
	"Tristanarcher": {
		"id":   uint(667247),
		"form": uint(0),
		"type": "sound",
	},
	"Lancartorialter": {
		"id":   uint(667248),
		"form": uint(0),
		"type": "dragon",
	},
	"Merlinaster": {
		"id":   uint(667249),
		"form": uint(0),
		"type": "magic",
	},
	"Ozyridias": {
		"id":   uint(667250),
		"form": uint(0),
		"type": "divine",
	},
	"Lishuwassin": {
		"id":   uint(667251),
		"form": uint(0),
		"type": "fighting",
	},
	"Beoserker": {
		"id":   uint(667252),
		"form": uint(0),
		"type": "fighting",
	},
	"Gawaiknight": {
		"id":   uint(667253),
		"form": uint(0),
		"type": "light",
	},
	"Sularcl": {
		"id":   uint(667254),
		"form": uint(0),
		"type": "rock",
	},
	"Gooseling": {
		"id":   uint(667255),
		"form": uint(0),
		"type": "normal",
	},
	"Gooseloose": {
		"id":   uint(667256),
		"form": uint(0),
		"type": "normal",
	},
	"Gooseon": {
		"id":   uint(667257),
		"form": uint(0),
		"type": "normal",
	},
	"Brynhildcer": {
		"id":   uint(667258),
		"form": uint(0),
		"type": "heart",
	},
	"Quachildust": {
		"id":   uint(667259),
		"form": uint(0),
		"type": "zombie",
	},
	"Kinghassassin": {
		"id":   uint(667260),
		"form": uint(0),
		"type": "dark",
	},
	"Ballooneon": {
		"id":   uint(667261),
		"form": uint(0),
		"type": "rubber",
	},
	"Elechair": {
		"id":   uint(667262),
		"form": uint(0),
		"type": "wood",
	},
	"Gadgeteon": {
		"id":   uint(667263),
		"form": uint(0),
		"type": "tech",
	},
	"Chaneller": {
		"id":   uint(667264),
		"form": uint(0),
		"type": "ghost",
	},
	"Mlucario": {
		"id":   uint(667265),
		"form": uint(0),
		"type": "fighting",
	},
	"Mmanectric": {
		"id":   uint(667266),
		"form": uint(0),
		"type": "electric",
	},
	"Mslowbro": {
		"id":   uint(667267),
		"form": uint(0),
		"type": "water",
	},
	"Mgyarados": {
		"id":   uint(667268),
		"form": uint(0),
		"type": "water",
	},
	"Mscizor": {
		"id":   uint(667269),
		"form": uint(0),
		"type": "bug",
	},
	"Mglalie": {
		"id":   uint(667270),
		"form": uint(0),
		"type": "ice",
	},
	"Mlopunny": {
		"id":   uint(667271),
		"form": uint(0),
		"type": "normal",
	},
	"Snapchost": {
		"id":   uint(667272),
		"form": uint(0),
		"type": "cyber",
	},
	"Virtualeon": {
		"id":   uint(667273),
		"form": uint(0),
		"type": "cyber",
	},
	"Alteraber": {
		"id":   uint(667274),
		"form": uint(0),
		"type": "divine",
	},
	"Mariantoinette": {
		"id":   uint(667275),
		"form": uint(0),
		"type": "glass",
	},
	"Mhlopunny": {
		"id":   uint(667276),
		"form": uint(0),
		"type": "fighting",
	},
	"Rodac": {
		"id":   uint(667277),
		"form": uint(0),
		"type": "normal",
	},
	"Rodactyl": {
		"id":   uint(667278),
		"form": uint(0),
		"type": "nuclear",
	},
	"Cut Cut": {
		"id":   uint(667279),
		"form": uint(0),
		"type": "steel",
	},
	"Eyehaed": {
		"id":   uint(667280),
		"form": uint(0),
		"type": "fear",
	},
	"Maniquesque": {
		"id":   uint(667281),
		"form": uint(0),
		"type": "plastic",
	},
	"Drownskel": {
		"id":   uint(667282),
		"form": uint(0),
		"type": "water",
	},
	"Rotree": {
		"id":   uint(667283),
		"form": uint(0),
		"type": "wood",
	},
	"Mestwi": {
		"id":   uint(667284),
		"form": uint(0),
		"type": "dark",
	},
	"Mestwi-Zen": {
		"id":   uint(667284),
		"form": uint(1),
		"type": "dark",
	},
	"Kawainnocent": {
		"id":   uint(667285),
		"form": uint(0),
		"type": "qmarks",
	},
	"Kawainnocent-Zen": {
		"id":   uint(667285),
		"form": uint(1),
		"type": "qmarks",
	},
	"Elephyxiation": {
		"id":   uint(667286),
		"form": uint(0),
		"type": "poison",
	},
	"Tristilt": {
		"id":   uint(667287),
		"form": uint(0),
		"type": "bug",
	},
	"Faircore": {
		"id":   uint(667288),
		"form": uint(0),
		"type": "fairy",
	},
	"Slugeongoo": {
		"id":   uint(667289),
		"form": uint(0),
		"type": "greasy",
	},
	"Cuchucaster": {
		"id":   uint(667290),
		"form": uint(0),
		"type": "fairy",
	},
	"Currenceagle": {
		"id":   uint(667291),
		"form": uint(0),
		"type": "steel",
	},
	"Maggotfly": {
		"id":   uint(667292),
		"form": uint(0),
		"type": "bug",
	},
	"Masigil": {
		"id":   uint(667293),
		"form": uint(0),
		"type": "chaos",
	},
	"Anatomyan": {
		"id":   uint(667294),
		"form": uint(0),
		"type": "blood",
	},
	"Lilskars": {
		"id":   uint(667295),
		"form": uint(0),
		"type": "ghost",
	},
	"Meatcko": {
		"id":   uint(667296),
		"form": uint(0),
		"type": "chaos",
	},
	"Whaledriff": {
		"id":   uint(667297),
		"form": uint(0),
		"type": "water",
	},
	"Crustolor": {
		"id":   uint(667298),
		"form": uint(0),
		"type": "qmarks",
	},
	"Muschum": {
		"id":   uint(667299),
		"form": uint(0),
		"type": "fighting",
	},
	"Drawbie": {
		"id":   uint(667300),
		"form": uint(0),
		"type": "paper",
	},
	"Tortuagel": {
		"id":   uint(667301),
		"form": uint(0),
		"type": "divine",
	},
	"Cakeross": {
		"id":   uint(667302),
		"form": uint(0),
		"type": "food",
	},
	"Rockodile": {
		"id":   uint(667303),
		"form": uint(0),
		"type": "rock",
	},
	"Crocwistul": {
		"id":   uint(667304),
		"form": uint(0),
		"type": "rock",
	},
	"Orinrin": {
		"id":   uint(667305),
		"form": uint(0),
		"type": "fire",
	},
	"Beeabee": {
		"id":   uint(667306),
		"form": uint(0),
		"type": "bug",
	},
	"Mistroo": {
		"id":   uint(667307),
		"form": uint(0),
		"type": "ghost",
	},
	"Statikman": {
		"id":   uint(667308),
		"form": uint(0),
		"type": "tech",
	},
	"Praticate": {
		"id":   uint(667309),
		"form": uint(0),
		"type": "normal",
	},
	"Snowlax": {
		"id":   uint(667310),
		"form": uint(0),
		"type": "ice",
	},
	"Cradisk": {
		"id":   uint(667311),
		"form": uint(0),
		"type": "rock",
	},
	"Yogsoth": {
		"id":   uint(667312),
		"form": uint(0),
		"type": "cosmic",
	},
	"Chanseychanse": {
		"id":   uint(667313),
		"form": uint(0),
		"type": "qmarks",
	},
	"C Encyclodia": {
		"id":   uint(667314),
		"form": uint(0),
		"type": "paper",
	},
	"Nameless Demon": {
		"id":   uint(667315),
		"form": uint(0),
		"type": "divine",
	},
	"Darkyo": {
		"id":   uint(667316),
		"form": uint(0),
		"type": "wood",
	},
	"Darkyoun": {
		"id":   uint(667317),
		"form": uint(0),
		"type": "wood",
	},
	"Shubniggurother": {
		"id":   uint(667318),
		"form": uint(0),
		"type": "wood",
	},
	"Ithaqwind": {
		"id":   uint(667319),
		"form": uint(0),
		"type": "ice",
	},
	"Venustoise": {
		"id":   uint(667320),
		"form": uint(0),
		"type": "water",
	},
	"Frigo": {
		"id":   uint(667321),
		"form": uint(0),
		"type": "ice",
	},
	"Hiker": {
		"id":   uint(667322),
		"form": uint(0),
		"type": "normal",
	},
	"Glassbol": {
		"id":   uint(667323),
		"form": uint(0),
		"type": "glass",
	},
	"Aquariuol": {
		"id":   uint(667324),
		"form": uint(0),
		"type": "glass",
	},
	"Terraol": {
		"id":   uint(667325),
		"form": uint(0),
		"type": "glass",
	},
	"Voltbyor": {
		"id":   uint(667326),
		"form": uint(0),
		"type": "electric",
	},
	"Smariados": {
		"id":   uint(667327),
		"form": uint(0),
		"type": "bug",
	},
	"Tuberm": {
		"id":   uint(667328),
		"form": uint(0),
		"type": "water",
	},
	"Tuberf": {
		"id":   uint(667329),
		"form": uint(0),
		"type": "water",
	},
	"Swimmerm": {
		"id":   uint(667330),
		"form": uint(0),
		"type": "water",
	},
	"Swimmerf": {
		"id":   uint(667331),
		"form": uint(0),
		"type": "water",
	},
	"Painter": {
		"id":   uint(667332),
		"form": uint(0),
		"type": "paint",
	},
	"Bird Keeper": {
		"id":   uint(667333),
		"form": uint(0),
		"type": "normal",
	},
	"Aroma Lady": {
		"id":   uint(667334),
		"form": uint(0),
		"type": "grass",
	},
	"Beauty": {
		"id":   uint(667335),
		"form": uint(0),
		"type": "heart",
	},
	"Elsloggeroth": {
		"id":   uint(667336),
		"form": uint(0),
		"type": "cosmic",
	},
	"Husloggeroth": {
		"id":   uint(667337),
		"form": uint(0),
		"type": "cosmic",
	},
	"Datatrainer": {
		"id":   uint(667338),
		"form": uint(0),
		"type": "cyber",
	},
	"Youngster": {
		"id":   uint(667339),
		"form": uint(0),
		"type": "normal",
	},
	"Lass": {
		"id":   uint(667340),
		"form": uint(0),
		"type": "normal",
	},
	"Picknicker": {
		"id":   uint(667341),
		"form": uint(0),
		"type": "normal",
	},
	"Sailor": {
		"id":   uint(667342),
		"form": uint(0),
		"type": "water",
	},
	"Bugcatcher": {
		"id":   uint(667343),
		"form": uint(0),
		"type": "normal",
	},
	"Fisherman": {
		"id":   uint(667344),
		"form": uint(0),
		"type": "water",
	},
	"Crush Girl": {
		"id":   uint(667345),
		"form": uint(0),
		"type": "fighting",
	},
	"Enderite": {
		"id":   uint(667346),
		"form": uint(0),
		"type": "cosmic",
	},
	"Endmeran": {
		"id":   uint(667347),
		"form": uint(0),
		"type": "cosmic",
	},
	"Endragon": {
		"id":   uint(667348),
		"form": uint(0),
		"type": "cosmic",
	},
	"Endershulk": {
		"id":   uint(667349),
		"form": uint(0),
		"type": "cosmic",
	},
	"Maverock": {
		"id":   uint(667350),
		"form": uint(0),
		"type": "tech",
	},
	"Chillrock": {
		"id":   uint(667351),
		"form": uint(0),
		"type": "tech",
	},
	"Overock": {
		"id":   uint(667352),
		"form": uint(0),
		"type": "tech",
	},
	"Blastrock": {
		"id":   uint(667353),
		"form": uint(0),
		"type": "tech",
	},
	"Pfpoof": {
		"id":   uint(667354),
		"form": uint(0),
		"type": "magic",
	},
	"Mecha Zaydolf": {
		"id":   uint(667355),
		"form": uint(0),
		"type": "tech",
	},
	"Scrunrafe": {
		"id":   uint(667356),
		"form": uint(0),
		"type": "normal",
	},
	"Loganne": {
		"id":   uint(667357),
		"form": uint(0),
		"type": "wood",
	},
	"Alphanne": {
		"id":   uint(667358),
		"form": uint(0),
		"type": "divine",
	},
	"Voidanne": {
		"id":   uint(667359),
		"form": uint(0),
		"type": "cosmic",
	},
	"Fatherat": {
		"id":   uint(667360),
		"form": uint(0),
		"type": "blood",
	},
	"Coneanne": {
		"id":   uint(667361),
		"form": uint(0),
		"type": "psychic",
	},
	"Emptyanne": {
		"id":   uint(667362),
		"form": uint(0),
		"type": "qmarks",
	},
	"Emptyanne-Zen": {
		"id":   uint(667362),
		"form": uint(1),
		"type": "qmarks",
	},
	"Anneoo": {
		"id":   uint(667363),
		"form": uint(0),
		"type": "qmarks",
	},
	"Wcanvast": {
		"id":   uint(667364),
		"form": uint(0),
		"type": "paint",
	},
	"Godluna": {
		"id":   uint(667365),
		"form": uint(0),
		"type": "divine",
	},
	"Godsol": {
		"id":   uint(667366),
		"form": uint(0),
		"type": "divine",
	},
	"Godmercury": {
		"id":   uint(667367),
		"form": uint(0),
		"type": "divine",
	},
	"Godvenus": {
		"id":   uint(667368),
		"form": uint(0),
		"type": "divine",
	},
	"Godmars": {
		"id":   uint(667369),
		"form": uint(0),
		"type": "divine",
	},
	"Godjupiter": {
		"id":   uint(667370),
		"form": uint(0),
		"type": "divine",
	},
	"Godsaturn": {
		"id":   uint(667371),
		"form": uint(0),
		"type": "divine",
	},
	"Godneptune": {
		"id":   uint(667372),
		"form": uint(0),
		"type": "divine",
	},
	"Goduranus": {
		"id":   uint(667373),
		"form": uint(0),
		"type": "cosmic",
	},
	"Godpluto": {
		"id":   uint(667374),
		"form": uint(0),
		"type": "cosmic",
	},
	"Starmish": {
		"id":   uint(667375),
		"form": uint(0),
		"type": "water",
	},
	"Ghostly": {
		"id":   uint(667376),
		"form": uint(0),
		"type": "ghost",
	},
	"Scytheper": {
		"id":   uint(667377),
		"form": uint(0),
		"type": "bug",
	},
	"Asmoochum": {
		"id":   uint(667378),
		"form": uint(0),
		"type": "fire",
	},
	"Ajynx": {
		"id":   uint(667379),
		"form": uint(0),
		"type": "fire",
	},
	"Ivictini": {
		"id":   uint(667380),
		"form": uint(0),
		"type": "paper",
	},
	"Ishtarin": {
		"id":   uint(667381),
		"form": uint(0),
		"type": "divine",
	},
	"Hearhild": {
		"id":   uint(667382),
		"form": uint(0),
		"type": "normal",
	},
	"Heartlenine": {
		"id":   uint(667383),
		"form": uint(0),
		"type": "dark",
	},
	"Nobotile": {
		"id":   uint(667384),
		"form": uint(0),
		"type": "zombie",
	},
	"Unlion": {
		"id":   uint(667385),
		"form": uint(0),
		"type": "ghost",
	},
	"Dreagamorph": {
		"id":   uint(667386),
		"form": uint(0),
		"type": "fairy",
	},
	"Royickish": {
		"id":   uint(667387),
		"form": uint(0),
		"type": "light",
	},
	"Ezigzagoon": {
		"id":   uint(667388),
		"form": uint(0),
		"type": "electric",
	},
	"Elinoone": {
		"id":   uint(667389),
		"form": uint(0),
		"type": "electric",
	},
	"Pengui": {
		"id":   uint(667390),
		"form": uint(0),
		"type": "ice",
	},
	"Pengemperor": {
		"id":   uint(667391),
		"form": uint(0),
		"type": "ice",
	},
	"Albinguin": {
		"id":   uint(667392),
		"form": uint(0),
		"type": "ice",
	},
	"Gzigzagoon": {
		"id":   uint(667393),
		"form": uint(0),
		"type": "dark",
	},
	"Glinoone": {
		"id":   uint(667394),
		"form": uint(0),
		"type": "dark",
	},
	"Gweezing": {
		"id":   uint(667395),
		"form": uint(0),
		"type": "poison",
	},
	"Mistaibah X": {
		"id":   uint(667396),
		"form": uint(0),
		"type": "dragon",
	},
	"Nanachu": {
		"id":   uint(667397),
		"form": uint(0),
		"type": "electric",
	},
	"Scpee173": {
		"id":   uint(667398),
		"form": uint(0),
		"type": "rock",
	},
	"Lilsaibah": {
		"id":   uint(667399),
		"form": uint(0),
		"type": "dragon",
	},
	"Summarthoria": {
		"id":   uint(667400),
		"form": uint(0),
		"type": "dragon",
	},
	"Babbaster": {
		"id":   uint(667401),
		"form": uint(0),
		"type": "tech",
	},
	"Paperbowser": {
		"id":   uint(667402),
		"form": uint(0),
		"type": "fire",
	},
	"Skylanderbowser": {
		"id":   uint(667403),
		"form": uint(0),
		"type": "plastic",
	},
	"Bowletta": {
		"id":   uint(667404),
		"form": uint(0),
		"type": "fire",
	},
	"Aseedot": {
		"id":   uint(667405),
		"form": uint(0),
		"type": "wood",
	},
	"Anuzleaf": {
		"id":   uint(667406),
		"form": uint(0),
		"type": "wood",
	},
	"Ashiftry": {
		"id":   uint(667407),
		"form": uint(0),
		"type": "wood",
	},
	"Condenseon": {
		"id":   uint(667408),
		"form": uint(0),
		"type": "steam",
	},
	"Volcaneon": {
		"id":   uint(667409),
		"form": uint(0),
		"type": "magma",
	},
	"Timberreon": {
		"id":   uint(667410),
		"form": uint(0),
		"type": "wood",
	},
	"Sphealedge": {
		"id":   uint(667411),
		"form": uint(0),
		"type": "dark",
	},
	"Sealeoedge": {
		"id":   uint(667412),
		"form": uint(0),
		"type": "dark",
	},
	"Walreinedge": {
		"id":   uint(667413),
		"form": uint(0),
		"type": "dark",
	},
	"Shanggai": {
		"id":   uint(667414),
		"form": uint(0),
		"type": "bug",
	},
	"Xadazoth": {
		"id":   uint(667415),
		"form": uint(0),
		"type": "qmarks",
	},
	"Gcharmander": {
		"id":   uint(667416),
		"form": uint(0),
		"type": "ground",
	},
	"Gcharmeleon": {
		"id":   uint(667417),
		"form": uint(0),
		"type": "ground",
	},
	"Gcharizard": {
		"id":   uint(667418),
		"form": uint(0),
		"type": "ground",
	},
	"Xiclotlee": {
		"id":   uint(667419),
		"form": uint(0),
		"type": "steel",
	},
	"Plutoneon": {
		"id":   uint(667420),
		"form": uint(0),
		"type": "nuclear",
	},
	"Unownking": {
		"id":   uint(667421),
		"form": uint(0),
		"type": "psychic",
	},
	"Yithimeless": {
		"id":   uint(667422),
		"form": uint(0),
		"type": "psychic",
	},
	"Flolyp": {
		"id":   uint(667423),
		"form": uint(0),
		"type": "flying",
	},
	"Flyongolyp": {
		"id":   uint(667424),
		"form": uint(0),
		"type": "flying",
	},
	"Lusfairy": {
		"id":   uint(667425),
		"form": uint(0),
		"type": "heart",
	},
	"Mylph": {
		"id":   uint(667426),
		"form": uint(0),
		"type": "heart",
	},
	"Migus": {
		"id":   uint(667427),
		"form": uint(0),
		"type": "cosmic",
	},
	"Migoggoth": {
		"id":   uint(667428),
		"form": uint(0),
		"type": "cosmic",
	},
	"Stillnut": {
		"id":   uint(667429),
		"form": uint(0),
		"type": "food",
	},
	"Mr Stillnut": {
		"id":   uint(667430),
		"form": uint(0),
		"type": "food",
	},
	"Lord Stillnut": {
		"id":   uint(667431),
		"form": uint(0),
		"type": "food",
	},
	"Gentleman": {
		"id":   uint(667432),
		"form": uint(0),
		"type": "normal",
	},
	"Lady": {
		"id":   uint(667433),
		"form": uint(0),
		"type": "heart",
	},
	"Breeder": {
		"id":   uint(667434),
		"form": uint(0),
		"type": "normal",
	},
	"Twins": {
		"id":   uint(667435),
		"form": uint(0),
		"type": "normal",
	},
	"Burglar": {
		"id":   uint(667436),
		"form": uint(0),
		"type": "dark",
	},
	"Lustubust": {
		"id":   uint(667437),
		"form": uint(0),
		"type": "heart",
	},
	"Toriegoat": {
		"id":   uint(667438),
		"form": uint(0),
		"type": "magic",
	},
	"Sgeodude": {
		"id":   uint(667439),
		"form": uint(0),
		"type": "steam",
	},
	"Sgraveler": {
		"id":   uint(667440),
		"form": uint(0),
		"type": "steam",
	},
	"Sgolem": {
		"id":   uint(667441),
		"form": uint(0),
		"type": "steam",
	},
	"Steamea": {
		"id":   uint(667442),
		"form": uint(0),
		"type": "steam",
	},
	"Stweams": {
		"id":   uint(667443),
		"form": uint(0),
		"type": "steam",
	},
	"Yamcha": {
		"id":   uint(667444),
		"form": uint(0),
		"type": "fighting",
	},
	"Kagerawoo": {
		"id":   uint(667445),
		"form": uint(0),
		"type": "dark",
	},
	"Ndratini": {
		"id":   uint(667446),
		"form": uint(0),
		"type": "nuclear",
	},
	"Ndragonair": {
		"id":   uint(667447),
		"form": uint(0),
		"type": "nuclear",
	},
	"Ndragonite": {
		"id":   uint(667448),
		"form": uint(0),
		"type": "nuclear",
	},
	"Windows98": {
		"id":   uint(667449),
		"form": uint(0),
		"type": "cyber",
	},
	"Rotteon": {
		"id":   uint(667450),
		"form": uint(0),
		"type": "zombie",
	},
	"Himster": {
		"id":   uint(667451),
		"form": uint(0),
		"type": "normal",
	},
	"Blasphlute": {
		"id":   uint(667452),
		"form": uint(0),
		"type": "chaos",
	},
	"Smolstron": {
		"id":   uint(667453),
		"form": uint(0),
		"type": "ghost",
	},
	"Cthylltulla": {
		"id":   uint(667454),
		"form": uint(0),
		"type": "water",
	},
	"Snakeny": {
		"id":   uint(667455),
		"form": uint(0),
		"type": "poison",
	},
	"Serpeople": {
		"id":   uint(667456),
		"form": uint(0),
		"type": "poison",
	},
	"Yigsnake": {
		"id":   uint(667457),
		"form": uint(0),
		"type": "poison",
	},
	"The Guy": {
		"id":   uint(667458),
		"form": uint(0),
		"type": "qmarks",
	},
	"Hydeep": {
		"id":   uint(667459),
		"form": uint(0),
		"type": "normal",
	},
	"Deepone": {
		"id":   uint(667460),
		"form": uint(0),
		"type": "water",
	},
	"Dagonther": {
		"id":   uint(667461),
		"form": uint(0),
		"type": "water",
	},
	"Mothydra": {
		"id":   uint(667462),
		"form": uint(0),
		"type": "water",
	},
	"Killnut": {
		"id":   uint(667463),
		"form": uint(0),
		"type": "food",
	},
	"Mr Killnut": {
		"id":   uint(667464),
		"form": uint(0),
		"type": "food",
	},
	"Lord Killnut": {
		"id":   uint(667465),
		"form": uint(0),
		"type": "food",
	},
	"Geno": {
		"id":   uint(667466),
		"form": uint(0),
		"type": "wood",
	},
	"Deltrune": {
		"id":   uint(667467),
		"form": uint(0),
		"type": "light",
	},
	"Sweetergeist": {
		"id":   uint(667468),
		"form": uint(0),
		"type": "ghost",
	},
	"Regighost": {
		"id":   uint(667469),
		"form": uint(0),
		"type": "ghost",
	},
	"Regiwood": {
		"id":   uint(667470),
		"form": uint(0),
		"type": "wood",
	},
	"Regifight": {
		"id":   uint(667471),
		"form": uint(0),
		"type": "fighting",
	},
	"Bokrugod": {
		"id":   uint(667472),
		"form": uint(0),
		"type": "water",
	},
	"Ghatanovisage": {
		"id":   uint(667473),
		"form": uint(0),
		"type": "fear",
	},
	"Rhantegice": {
		"id":   uint(667474),
		"form": uint(0),
		"type": "bug",
	},
	"Tulzsflame": {
		"id":   uint(667475),
		"form": uint(0),
		"type": "fire",
	},
	"Dimenshambler": {
		"id":   uint(667476),
		"form": uint(0),
		"type": "cosmic",
	},
	"Timedalos": {
		"id":   uint(667477),
		"form": uint(0),
		"type": "time",
	},
	"Chaugnaphant": {
		"id":   uint(667478),
		"form": uint(0),
		"type": "rock",
	},
	"Azathoth": {
		"id":   uint(667479),
		"form": uint(0),
		"type": "chaos",
	},
	"Firepire": {
		"id":   uint(667480),
		"form": uint(0),
		"type": "fire",
	},
	"Cthugyre": {
		"id":   uint(667481),
		"form": uint(0),
		"type": "fire",
	},
	"Colorspace": {
		"id":   uint(667482),
		"form": uint(0),
		"type": "cosmic",
	},
	"Necroli": {
		"id":   uint(667483),
		"form": uint(0),
		"type": "zombie",
	},
	"Necrolili": {
		"id":   uint(667484),
		"form": uint(0),
		"type": "zombie",
	},
	"Necrolilith": {
		"id":   uint(667485),
		"form": uint(0),
		"type": "zombie",
	},
	"Himw": {
		"id":   uint(667486),
		"form": uint(0),
		"type": "water",
	},
	"Gothot": {
		"id":   uint(667487),
		"form": uint(0),
		"type": "fairy",
	},
	"Gooei": {
		"id":   uint(667488),
		"form": uint(0),
		"type": "water",
	},
	"Gooarm": {
		"id":   uint(667489),
		"form": uint(0),
		"type": "rubber",
	},
	"Gooswole": {
		"id":   uint(667490),
		"form": uint(0),
		"type": "rubber",
	},
	"Ubbosathlime": {
		"id":   uint(667491),
		"form": uint(0),
		"type": "qmarks",
	},
	"Himg": {
		"id":   uint(667492),
		"form": uint(0),
		"type": "grass",
	},
	"Chanaroy": {
		"id":   uint(667493),
		"form": uint(0),
		"type": "normal",
	},
	"Chanarang": {
		"id":   uint(667494),
		"form": uint(0),
		"type": "normal",
	},
	"Chanoken": {
		"id":   uint(667495),
		"form": uint(0),
		"type": "normal",
	},
	"Dbounsweet": {
		"id":   uint(667496),
		"form": uint(0),
		"type": "dark",
	},
	"Dsteenee": {
		"id":   uint(667497),
		"form": uint(0),
		"type": "dark",
	},
	"Dtsareena": {
		"id":   uint(667498),
		"form": uint(0),
		"type": "dark",
	},
	"Earbuzz": {
		"id":   uint(667499),
		"form": uint(0),
		"type": "bug",
	},
	"Phornet": {
		"id":   uint(667500),
		"form": uint(0),
		"type": "bug",
	},
	"Madegg": {
		"id":   uint(667501),
		"form": uint(0),
		"type": "dark",
	},
	"Sbellsprout": {
		"id":   uint(667502),
		"form": uint(0),
		"type": "grass",
	},
	"Mpikachu": {
		"id":   uint(667503),
		"form": uint(0),
		"type": "electric",
	},
	"Melectrode": {
		"id":   uint(667504),
		"form": uint(0),
		"type": "electric",
	},
	"Mjynx": {
		"id":   uint(667505),
		"form": uint(0),
		"type": "ice",
	},
	"Majynx": {
		"id":   uint(667506),
		"form": uint(0),
		"type": "rock",
	},
	"Mariados": {
		"id":   uint(667507),
		"form": uint(0),
		"type": "bug",
	},
	"Mdelcatty": {
		"id":   uint(667508),
		"form": uint(0),
		"type": "normal",
	},
	"Msunflora": {
		"id":   uint(667509),
		"form": uint(0),
		"type": "grass",
	},
	"Mcorsola": {
		"id":   uint(667510),
		"form": uint(0),
		"type": "water",
	},
	"Deli Delibird": {
		"id":   uint(667511),
		"form": uint(0),
		"type": "ice",
	},
	"Moctillery": {
		"id":   uint(667512),
		"form": uint(0),
		"type": "water",
	},
	"Mdonphan": {
		"id":   uint(667513),
		"form": uint(0),
		"type": "ground",
	},
	"Mstantler": {
		"id":   uint(667514),
		"form": uint(0),
		"type": "normal",
	},
	"Mtorkoal": {
		"id":   uint(667515),
		"form": uint(0),
		"type": "fire",
	},
	"Mkecleon": {
		"id":   uint(667516),
		"form": uint(0),
		"type": "normal",
	},
	"Mhuntail": {
		"id":   uint(667517),
		"form": uint(0),
		"type": "water",
	},
	"Mgorebyss": {
		"id":   uint(667518),
		"form": uint(0),
		"type": "water",
	},
	"Mchatot": {
		"id":   uint(667519),
		"form": uint(0),
		"type": "sound",
	},
	"Mcarnivine": {
		"id":   uint(667520),
		"form": uint(0),
		"type": "grass",
	},
	"Mlumineon": {
		"id":   uint(667521),
		"form": uint(0),
		"type": "water",
	},
	"Mswoobat": {
		"id":   uint(667522),
		"form": uint(0),
		"type": "psychic",
	},
	"Mheliolisk": {
		"id":   uint(667523),
		"form": uint(0),
		"type": "electric",
	},
	"Nfarfetchd": {
		"id":   uint(667524),
		"form": uint(0),
		"type": "nuclear",
	},
	"Mplusle": {
		"id":   uint(667525),
		"form": uint(0),
		"type": "electric",
	},
	"Mminun": {
		"id":   uint(667526),
		"form": uint(0),
		"type": "electric",
	},
	"Voltcube": {
		"id":   uint(667527),
		"form": uint(0),
		"type": "electric",
	},
	"Daolangle": {
		"id":   uint(667528),
		"form": uint(0),
		"type": "steel",
	},
	"Mwhiscash": {
		"id":   uint(667529),
		"form": uint(0),
		"type": "water",
	},
	"Mrelicanth": {
		"id":   uint(667530),
		"form": uint(0),
		"type": "water",
	},
	"Momastar": {
		"id":   uint(667531),
		"form": uint(0),
		"type": "rock",
	},
	"Mkabutops": {
		"id":   uint(667532),
		"form": uint(0),
		"type": "rock",
	},
	"Mgarbodor": {
		"id":   uint(667533),
		"form": uint(0),
		"type": "poison",
	},
	"Starpire": {
		"id":   uint(667534),
		"form": uint(0),
		"type": "blood",
	},
	"Protecdoll": {
		"id":   uint(667535),
		"form": uint(0),
		"type": "steel",
	},
	"Proteknight": {
		"id":   uint(667536),
		"form": uint(0),
		"type": "steel",
	},
	"Skittiglock": {
		"id":   uint(667537),
		"form": uint(0),
		"type": "normal",
	},
	"Skittigang": {
		"id":   uint(667538),
		"form": uint(0),
		"type": "normal",
	},
	"Mluvdiscx": {
		"id":   uint(667539),
		"form": uint(0),
		"type": "water",
	},
	"Mluvdiscy": {
		"id":   uint(667540),
		"form": uint(0),
		"type": "water",
	},
	"Karmris": {
		"id":   uint(667541),
		"form": uint(0),
		"type": "light",
	},
	"Swimnecrolilith": {
		"id":   uint(667542),
		"form": uint(0),
		"type": "water",
	},
	"Gloob": {
		"id":   uint(667543),
		"form": uint(0),
		"type": "dark",
	},
	"Dreamulus": {
		"id":   uint(667544),
		"form": uint(0),
		"type": "fairy",
	},
	"Duodreamulus": {
		"id":   uint(667545),
		"form": uint(0),
		"type": "fairy",
	},
	"Dreamungus": {
		"id":   uint(667546),
		"form": uint(0),
		"type": "fairy",
	},
	"Dreamare": {
		"id":   uint(667547),
		"form": uint(0),
		"type": "fear",
	},
	"Drasleep": {
		"id":   uint(667548),
		"form": uint(0),
		"type": "dragon",
	},
	"Drabite": {
		"id":   uint(667549),
		"form": uint(0),
		"type": "dragon",
	},
	"Biteragon": {
		"id":   uint(667550),
		"form": uint(0),
		"type": "dragon",
	},
	"Herculeagon": {
		"id":   uint(667551),
		"form": uint(0),
		"type": "dragon",
	},
	"Mmagcargo": {
		"id":   uint(667552),
		"form": uint(0),
		"type": "fire",
	},
	"Drowsleep": {
		"id":   uint(667553),
		"form": uint(0),
		"type": "fairy",
	},
	"Sleepevil": {
		"id":   uint(667554),
		"form": uint(0),
		"type": "fairy",
	},
	"Dreamarx": {
		"id":   uint(667555),
		"form": uint(0),
		"type": "fairy",
	},
	"Pizzaur": {
		"id":   uint(667556),
		"form": uint(0),
		"type": "food",
	},
	"Pizzaurex": {
		"id":   uint(667557),
		"form": uint(0),
		"type": "food",
	},
	"Drewi": {
		"id":   uint(667558),
		"form": uint(0),
		"type": "psychic",
	},
	"Dreamwi": {
		"id":   uint(667559),
		"form": uint(0),
		"type": "psychic",
	},
	"Psycasso": {
		"id":   uint(667560),
		"form": uint(0),
		"type": "psychic",
	},
	"Scathatcher": {
		"id":   uint(667561),
		"form": uint(0),
		"type": "fairy",
	},
	"Mninetales": {
		"id":   uint(667562),
		"form": uint(0),
		"type": "fire",
	},
	"Mrapidash": {
		"id":   uint(667563),
		"form": uint(0),
		"type": "fire",
	},
	"Formspawn": {
		"id":   uint(667564),
		"form": uint(0),
		"type": "greasy",
	},
	"Mheatmor": {
		"id":   uint(667565),
		"form": uint(0),
		"type": "fire",
	},
	"Himf": {
		"id":   uint(667566),
		"form": uint(0),
		"type": "fire",
	},
	"Xmrapidash": {
		"id":   uint(667567),
		"form": uint(0),
		"type": "fire",
	},
	"Tsoadggua": {
		"id":   uint(667568),
		"form": uint(0),
		"type": "ground",
	},
	"Atlachnacha": {
		"id":   uint(667569),
		"form": uint(0),
		"type": "bug",
	},
	"Nukrylic": {
		"id":   uint(667570),
		"form": uint(0),
		"type": "nuclear",
	},
	"Marcanine": {
		"id":   uint(667571),
		"form": uint(0),
		"type": "fire",
	},
	"Spooklypuff": {
		"id":   uint(667572),
		"form": uint(0),
		"type": "ghost",
	},
	"Stogenie": {
		"id":   uint(667573),
		"form": uint(0),
		"type": "poison",
	},
	"Ebolette": {
		"id":   uint(667574),
		"form": uint(0),
		"type": "virus",
	},
	"Ebolass": {
		"id":   uint(667575),
		"form": uint(0),
		"type": "virus",
	},
	"Tortaco": {
		"id":   uint(667576),
		"form": uint(0),
		"type": "food",
	},
	"Abfilth": {
		"id":   uint(667577),
		"form": uint(0),
		"type": "poison",
	},
	"Itsy": {
		"id":   uint(667578),
		"form": uint(0),
		"type": "poison",
	},
	"Bitsy": {
		"id":   uint(667579),
		"form": uint(0),
		"type": "tech",
	},
	"Sleamy": {
		"id":   uint(667580),
		"form": uint(0),
		"type": "water",
	},
	"Dementad": {
		"id":   uint(667581),
		"form": uint(0),
		"type": "water",
	},
	"Lombression": {
		"id":   uint(667582),
		"form": uint(0),
		"type": "water",
	},
	"Lunaticolo": {
		"id":   uint(667583),
		"form": uint(0),
		"type": "water",
	},
	"Glaakine": {
		"id":   uint(667584),
		"form": uint(0),
		"type": "water",
	},
	"Kelpight": {
		"id":   uint(667585),
		"form": uint(0),
		"type": "grass",
	},
	"Ghrothmet": {
		"id":   uint(667586),
		"form": uint(0),
		"type": "cosmic",
	},
	"Pijammin": {
		"id":   uint(667587),
		"form": uint(0),
		"type": "fabric",
	},
	"Tendream": {
		"id":   uint(667588),
		"form": uint(0),
		"type": "fabric",
	},
	"Amigoggoth": {
		"id":   uint(667589),
		"form": uint(0),
		"type": "cosmic",
	},
	"Nebrownie": {
		"id":   uint(667590),
		"form": uint(0),
		"type": "food",
	},
	"Galactcake": {
		"id":   uint(667591),
		"form": uint(0),
		"type": "food",
	},
	"Ygolonacol": {
		"id":   uint(667592),
		"form": uint(0),
		"type": "dark",
	},
	"Broodhort": {
		"id":   uint(667593),
		"form": uint(0),
		"type": "dark",
	},
	"Eihortli": {
		"id":   uint(667594),
		"form": uint(0),
		"type": "fear",
	},
	"Cthoniurm": {
		"id":   uint(667595),
		"form": uint(0),
		"type": "ground",
	},
	"Shuddermell": {
		"id":   uint(667596),
		"form": uint(0),
		"type": "ground",
	},
	"Himstm": {
		"id":   uint(667597),
		"form": uint(0),
		"type": "steam",
	},
	"Protogil": {
		"id":   uint(667598),
		"form": uint(0),
		"type": "divine",
	},
	"Steamercom": {
		"id":   uint(667599),
		"form": uint(0),
		"type": "cyber",
	},
	"Steamnie": {
		"id":   uint(667600),
		"form": uint(0),
		"type": "steam",
	},
	"Starvipup": {
		"id":   uint(667601),
		"form": uint(0),
		"type": "dark",
	},
	"Hungier": {
		"id":   uint(667602),
		"form": uint(0),
		"type": "dark",
	},
	"Hollowland": {
		"id":   uint(667603),
		"form": uint(0),
		"type": "dark",
	},
	"Loyaland": {
		"id":   uint(667604),
		"form": uint(0),
		"type": "dark",
	},
	"Nugger": {
		"id":   uint(667605),
		"form": uint(0),
		"type": "ice",
	},
	"Yebus": {
		"id":   uint(667606),
		"form": uint(0),
		"type": "fire",
	},
	"Pumpkapie": {
		"id":   uint(667607),
		"form": uint(0),
		"type": "ghost",
	},
	"Soulatte": {
		"id":   uint(667608),
		"form": uint(0),
		"type": "ghost",
	},
	"Emojicon": {
		"id":   uint(667609),
		"form": uint(0),
		"type": "meme",
	},
	"Cursedcon": {
		"id":   uint(667610),
		"form": uint(0),
		"type": "meme",
	},
	"Crycon": {
		"id":   uint(667611),
		"form": uint(0),
		"type": "meme",
	},
	"Xokcon": {
		"id":   uint(667612),
		"form": uint(0),
		"type": "meme",
	},
	"Noocon": {
		"id":   uint(667613),
		"form": uint(0),
		"type": "meme",
	},
	"Qwailmer": {
		"id":   uint(667614),
		"form": uint(0),
		"type": "qmarks",
	},
	"Qwailord": {
		"id":   uint(667615),
		"form": uint(0),
		"type": "qmarks",
	},
	"Lagoonask": {
		"id":   uint(667616),
		"form": uint(0),
		"type": "fear",
	},
	"Msteamboatle": {
		"id":   uint(667617),
		"form": uint(0),
		"type": "steam",
	},
	"Mlavagun": {
		"id":   uint(667618),
		"form": uint(0),
		"type": "magma",
	},
	"Marbrood": {
		"id":   uint(667619),
		"form": uint(0),
		"type": "wood",
	},
	"Opticus": {
		"id":   uint(667620),
		"form": uint(0),
		"type": "fear",
	},
	"Opticusowl": {
		"id":   uint(667621),
		"form": uint(0),
		"type": "fear",
	},
	"Opticusufo": {
		"id":   uint(667622),
		"form": uint(0),
		"type": "fear",
	},
	"Optulzu": {
		"id":   uint(667623),
		"form": uint(0),
		"type": "fear",
	},
	"Wololo": {
		"id":   uint(667624),
		"form": uint(0),
		"type": "normal",
	},
	"Gmeowth": {
		"id":   uint(667625),
		"form": uint(0),
		"type": "normal",
	},
	"Gameowth": {
		"id":   uint(667626),
		"form": uint(0),
		"type": "dark",
	},
	"Gpikachu": {
		"id":   uint(667627),
		"form": uint(0),
		"type": "electric",
	},
	"Geevee": {
		"id":   uint(667628),
		"form": uint(0),
		"type": "normal",
	},
	"Galcremie": {
		"id":   uint(667629),
		"form": uint(0),
		"type": "fairy",
	},
	"Gdrednaw": {
		"id":   uint(667630),
		"form": uint(0),
		"type": "water",
	},
	"Gcorviknight": {
		"id":   uint(667631),
		"form": uint(0),
		"type": "flying",
	},
	"Mspinda": {
		"id":   uint(667632),
		"form": uint(0),
		"type": "normal",
	},
	"Mpyroar": {
		"id":   uint(667633),
		"form": uint(0),
		"type": "fire",
	},
	"Ayymask": {
		"id":   uint(667634),
		"form": uint(0),
		"type": "fear",
	},
	"Wphantump": {
		"id":   uint(667635),
		"form": uint(0),
		"type": "ghost",
	},
	"Wtrevenant": {
		"id":   uint(667636),
		"form": uint(0),
		"type": "ghost",
	},
	"Zombmask": {
		"id":   uint(667637),
		"form": uint(0),
		"type": "fear",
	},
	"Battora": {
		"id":   uint(667638),
		"form": uint(0),
		"type": "normal",
	},
	"Bbattora": {
		"id":   uint(667639),
		"form": uint(0),
		"type": "dark",
	},
	"Venomothra": {
		"id":   uint(667640),
		"form": uint(0),
		"type": "bug",
	},
	"Beatoriche": {
		"id":   uint(667641),
		"form": uint(0),
		"type": "magic",
	},
	"Torappu": {
		"id":   uint(667642),
		"form": uint(0),
		"type": "fear",
	},
	"Fugury": {
		"id":   uint(667643),
		"form": uint(0),
		"type": "water",
	},
	"Navabrine": {
		"id":   uint(667644),
		"form": uint(0),
		"type": "steel",
	},
	"Koilrig": {
		"id":   uint(667645),
		"form": uint(0),
		"type": "steel",
	},
	"Wobbugeist": {
		"id":   uint(667646),
		"form": uint(0),
		"type": "fear",
	},
	"Yeerby": {
		"id":   uint(667647),
		"form": uint(0),
		"type": "normal",
	},
	"Psyelk": {
		"id":   uint(667648),
		"form": uint(0),
		"type": "normal",
	},
	"Suprago": {
		"id":   uint(667649),
		"form": uint(0),
		"type": "normal",
	},
	"Moosid": {
		"id":   uint(667650),
		"form": uint(0),
		"type": "normal",
	},
	"Swirlpup": {
		"id":   uint(667651),
		"form": uint(0),
		"type": "fairy",
	},
	"Frostpup": {
		"id":   uint(667652),
		"form": uint(0),
		"type": "fairy",
	},
	"Giramare": {
		"id":   uint(667653),
		"form": uint(0),
		"type": "normal",
	},
	"Giradream": {
		"id":   uint(667654),
		"form": uint(0),
		"type": "normal",
	},
	"Himmagma": {
		"id":   uint(667655),
		"form": uint(0),
		"type": "magma",
	},
	"Himwood": {
		"id":   uint(667656),
		"form": uint(0),
		"type": "wood",
	},
	"Trpsychic": {
		"id":   uint(667657),
		"form": uint(0),
		"type": "normal",
	},
	"Rocker": {
		"id":   uint(667658),
		"form": uint(0),
		"type": "normal",
	},
	"Clownask": {
		"id":   uint(667659),
		"form": uint(0),
		"type": "fear",
	},
	"Autisma": {
		"id":   uint(667660),
		"form": uint(0),
		"type": "virus",
	},
	"Nucleshark": {
		"id":   uint(667661),
		"form": uint(0),
		"type": "steel",
	},
	"Mummyask": {
		"id":   uint(667662),
		"form": uint(0),
		"type": "fear",
	},
	"Juggler": {
		"id":   uint(667663),
		"form": uint(0),
		"type": "normal",
	},
	"Withrew": {
		"id":   uint(667664),
		"form": uint(0),
		"type": "chaos",
	},
	"Withrew-Zen": {
		"id":   uint(667664),
		"form": uint(1),
		"type": "chaos",
	},
	"Malevia": {
		"id":   uint(667665),
		"form": uint(0),
		"type": "chaos",
	},
	"Rosemalum": {
		"id":   uint(667666),
		"form": uint(0),
		"type": "chaos",
	},
	"Subroclock": {
		"id":   uint(667667),
		"form": uint(0),
		"type": "ground",
	},
	"Egguish": {
		"id":   uint(667668),
		"form": uint(0),
		"type": "normal",
	},
	"Friendfly": {
		"id":   uint(667669),
		"form": uint(0),
		"type": "psychic",
	},
	"Deleggme": {
		"id":   uint(667670),
		"form": uint(0),
		"type": "blood",
	},
	"Nymbiosego": {
		"id":   uint(667671),
		"form": uint(0),
		"type": "rock",
	},
	"Mother Beast": {
		"id":   uint(667672),
		"form": uint(0),
		"type": "chaos",
	},
	"Slugia": {
		"id":   uint(667673),
		"form": uint(0),
		"type": "shadow",
	},
	"Smewtwo": {
		"id":   uint(667674),
		"form": uint(0),
		"type": "psychic",
	},
	"Stabra": {
		"id":   uint(667675),
		"form": uint(0),
		"type": "fighting",
	},
	"Kapierce": {
		"id":   uint(667676),
		"form": uint(0),
		"type": "fighting",
	},
	"Alakatana": {
		"id":   uint(667677),
		"form": uint(0),
		"type": "fighting",
	},
	"Servglaak": {
		"id":   uint(667678),
		"form": uint(0),
		"type": "zombie",
	},
	"Slimee": {
		"id":   uint(667679),
		"form": uint(0),
		"type": "normal",
	},
	"Kingslimee": {
		"id":   uint(667680),
		"form": uint(0),
		"type": "normal",
	},
	"Darlime": {
		"id":   uint(667681),
		"form": uint(0),
		"type": "dark",
	},
	"Kingdarlime": {
		"id":   uint(667682),
		"form": uint(0),
		"type": "dark",
	},
	"Taurosteak": {
		"id":   uint(667683),
		"form": uint(0),
		"type": "food",
	},
	"Milsteak": {
		"id":   uint(667684),
		"form": uint(0),
		"type": "ghost",
	},
	"Gponyta": {
		"id":   uint(667685),
		"form": uint(0),
		"type": "fairy",
	},
	"Grapidash": {
		"id":   uint(667686),
		"form": uint(0),
		"type": "fairy",
	},
	"Gstunfisk": {
		"id":   uint(667687),
		"form": uint(0),
		"type": "ground",
	},
	"Rookiedie": {
		"id":   uint(667688),
		"form": uint(0),
		"type": "flying",
	},
	"Nicket": {
		"id":   uint(667689),
		"form": uint(0),
		"type": "dark",
	},
	"Carkoal": {
		"id":   uint(667690),
		"form": uint(0),
		"type": "rock",
	},
	"Galmeowth": {
		"id":   uint(667691),
		"form": uint(0),
		"type": "steel",
	},
	"Gcorsola": {
		"id":   uint(667692),
		"form": uint(0),
		"type": "ghost",
	},
	"Hattena": {
		"id":   uint(667693),
		"form": uint(0),
		"type": "psychic",
	},
	"Gyamask": {
		"id":   uint(667694),
		"form": uint(0),
		"type": "ghost",
	},
	"Tbreloom": {
		"id":   uint(667695),
		"form": uint(0),
		"type": "grass",
	},
	"Astrobunny": {
		"id":   uint(667696),
		"form": uint(0),
		"type": "fire",
	},
	"Wsnover": {
		"id":   uint(667697),
		"form": uint(0),
		"type": "wood",
	},
	"Wabomasnow": {
		"id":   uint(667698),
		"form": uint(0),
		"type": "wood",
	},
	"Gmrmime": {
		"id":   uint(667699),
		"form": uint(0),
		"type": "psychic",
	},
	"Zazamenta": {
		"id":   uint(667700),
		"form": uint(0),
		"type": "fighting",
	},
	"Lhooh": {
		"id":   uint(667701),
		"form": uint(0),
		"type": "light",
	},
	"Scpee682": {
		"id":   uint(667702),
		"form": uint(0),
		"type": "dragon",
	},
	"Locoalmotive": {
		"id":   uint(667703),
		"form": uint(0),
		"type": "rock",
	},
	"Rotomchu": {
		"id":   uint(667704),
		"form": uint(0),
		"type": "electric",
	},
	"Tdyamask": {
		"id":   uint(667705),
		"form": uint(0),
		"type": "ghost",
	},
	"Tdcofagrigus": {
		"id":   uint(667706),
		"form": uint(0),
		"type": "ghost",
	},
	"Tgyamask": {
		"id":   uint(667707),
		"form": uint(0),
		"type": "ghost",
	},
	"Tgcofagrigus": {
		"id":   uint(667708),
		"form": uint(0),
		"type": "ghost",
	},
	"Tzyamask": {
		"id":   uint(667709),
		"form": uint(0),
		"type": "ghost",
	},
	"Tzcofagrigus": {
		"id":   uint(667710),
		"form": uint(0),
		"type": "ghost",
	},
	"Tmlyamask": {
		"id":   uint(667711),
		"form": uint(0),
		"type": "ghost",
	},
	"Tikcofagrigus": {
		"id":   uint(667712),
		"form": uint(0),
		"type": "ghost",
	},
	"Kasensage": {
		"id":   uint(667713),
		"form": uint(0),
		"type": "fighting",
	},
	"Tduskull": {
		"id":   uint(667714),
		"form": uint(0),
		"type": "ghost",
	},
	"Tdusclops": {
		"id":   uint(667715),
		"form": uint(0),
		"type": "ghost",
	},
	"Duscelle": {
		"id":   uint(667716),
		"form": uint(0),
		"type": "ghost",
	},
	"Tspoink": {
		"id":   uint(667717),
		"form": uint(0),
		"type": "dark",
	},
	"Tgrumpig": {
		"id":   uint(667718),
		"form": uint(0),
		"type": "dark",
	},
	"Tbmspoink": {
		"id":   uint(667719),
		"form": uint(0),
		"type": "dark",
	},
	"Tmimikyu": {
		"id":   uint(667720),
		"form": uint(0),
		"type": "fairy",
	},
	"Tmrmime": {
		"id":   uint(667721),
		"form": uint(0),
		"type": "ghost",
	},
	"Tmagcargo": {
		"id":   uint(667722),
		"form": uint(0),
		"type": "fire",
	},
	"Tspiritomb": {
		"id":   uint(667723),
		"form": uint(0),
		"type": "ghost",
	},
	"Pbvenonat": {
		"id":   uint(667724),
		"form": uint(0),
		"type": "bug",
	},
	"Lktoxtricity": {
		"id":   uint(667725),
		"form": uint(0),
		"type": "poison",
	},
	"Gfarfetchd": {
		"id":   uint(667726),
		"form": uint(0),
		"type": "fighting",
	},
	"Neapolitaneiscue": {
		"id":   uint(667727),
		"form": uint(0),
		"type": "ice",
	},
	"Neapolitaneiscue-Noice": {
		"id":   uint(667727),
		"form": uint(1),
		"type": "fire",
	},
	"Susully": {
		"id":   uint(667728),
		"form": uint(0),
		"type": "light",
	},
	"Gdarumaka": {
		"id":   uint(667729),
		"form": uint(0),
		"type": "ice",
	},
	"Gdarmanitan": {
		"id":   uint(667730),
		"form": uint(0),
		"type": "ice",
	},
	"Scpee999": {
		"id":   uint(667731),
		"form": uint(0),
		"type": "normal",
	},
	"Dampelectrode": {
		"id":   uint(667732),
		"form": uint(0),
		"type": "electric",
	},
	"Shadsquirtle": {
		"id":   uint(667733),
		"form": uint(0),
		"type": "water",
	},
	"Whitegengar": {
		"id":   uint(667734),
		"form": uint(0),
		"type": "ghost",
	},
	"Mpelipper": {
		"id":   uint(667735),
		"form": uint(0),
		"type": "water",
	},
	"Mpersian": {
		"id":   uint(667736),
		"form": uint(0),
		"type": "normal",
	},
	"Mapersian": {
		"id":   uint(667737),
		"form": uint(0),
		"type": "dark",
	},
	"Marbok": {
		"id":   uint(667738),
		"form": uint(0),
		"type": "poison",
	},
	"Clowdpoke": {
		"id":   uint(667739),
		"form": uint(0),
		"type": "water",
	},
	"Torcano": {
		"id":   uint(667740),
		"form": uint(0),
		"type": "fire",
	},
	"Grumboar": {
		"id":   uint(667741),
		"form": uint(0),
		"type": "psychic",
	},
	"Baidy": {
		"id":   uint(667742),
		"form": uint(0),
		"type": "normal",
	},
	"Melodyno": {
		"id":   uint(667743),
		"form": uint(0),
		"type": "normal",
	},
	"Wbaidy": {
		"id":   uint(667744),
		"form": uint(0),
		"type": "fairy",
	},
	"Waudino": {
		"id":   uint(667745),
		"form": uint(0),
		"type": "fairy",
	},
	"Wmelodyno": {
		"id":   uint(667746),
		"form": uint(0),
		"type": "fairy",
	},
	"Wfurret": {
		"id":   uint(667747),
		"form": uint(0),
		"type": "rock",
	},
	"Wfurret-Zen": {
		"id":   uint(667747),
		"form": uint(1),
		"type": "fairy",
	},
	"Taker": {
		"id":   uint(667748),
		"form": uint(0),
		"type": "dark",
	},
	"Cakebowser": {
		"id":   uint(667749),
		"form": uint(0),
		"type": "food",
	},
	"Giradrake": {
		"id":   uint(667750),
		"form": uint(0),
		"type": "normal",
	},
	"Seafarig": {
		"id":   uint(667751),
		"form": uint(0),
		"type": "water",
	},
	"Pyooraffe": {
		"id":   uint(667752),
		"form": uint(0),
		"type": "fairy",
	},
	"Darkmare": {
		"id":   uint(667753),
		"form": uint(0),
		"type": "dark",
	},
	"Stunfish": {
		"id":   uint(667754),
		"form": uint(0),
		"type": "ground",
	},
	"Shockeel": {
		"id":   uint(667755),
		"form": uint(0),
		"type": "ground",
	},
	"Mantastorm": {
		"id":   uint(667756),
		"form": uint(0),
		"type": "ground",
	},
	"Stuntrap": {
		"id":   uint(667757),
		"form": uint(0),
		"type": "ground",
	},
	"Cgeodude": {
		"id":   uint(667758),
		"form": uint(0),
		"type": "crystal",
	},
	"Cgraveler": {
		"id":   uint(667759),
		"form": uint(0),
		"type": "crystal",
	},
	"Cgolem": {
		"id":   uint(667760),
		"form": uint(0),
		"type": "crystal",
	},
	"Turkking": {
		"id":   uint(667761),
		"form": uint(0),
		"type": "normal",
	},
	"Wchansey": {
		"id":   uint(667762),
		"form": uint(0),
		"type": "normal",
	},
	"Wblissey": {
		"id":   uint(667763),
		"form": uint(0),
		"type": "normal",
	},
	"Hhgregg": {
		"id":   uint(667764),
		"form": uint(0),
		"type": "paper",
	},
	"Mcradily": {
		"id":   uint(667765),
		"form": uint(0),
		"type": "rock",
	},
	"Marmaldo": {
		"id":   uint(667766),
		"form": uint(0),
		"type": "rock",
	},
	"Drbowser": {
		"id":   uint(667767),
		"form": uint(0),
		"type": "fire",
	},
	"Himrock": {
		"id":   uint(667768),
		"form": uint(0),
		"type": "rock",
	},
	"Turlucky": {
		"id":   uint(667769),
		"form": uint(0),
		"type": "food",
	},
	"Depremeleon": {
		"id":   uint(667770),
		"form": uint(0),
		"type": "water",
	},
	"Suiceleon": {
		"id":   uint(667771),
		"form": uint(0),
		"type": "water",
	},
	"Zimbaga": {
		"id":   uint(667772),
		"form": uint(0),
		"type": "grass",
	},
	"Somalorilla": {
		"id":   uint(667773),
		"form": uint(0),
		"type": "grass",
	},
	"Cindemiga": {
		"id":   uint(667774),
		"form": uint(0),
		"type": "fire",
	},
	"Heakazuking": {
		"id":   uint(667775),
		"form": uint(0),
		"type": "fire",
	},
	"Koromon": {
		"id":   uint(667776),
		"form": uint(0),
		"type": "normal",
	},
	"Metalgreymon": {
		"id":   uint(667777),
		"form": uint(0),
		"type": "steel",
	},
	"Metalgreymonv": {
		"id":   uint(667778),
		"form": uint(0),
		"type": "steel",
	},
	"Pancosmic": {
		"id":   uint(667779),
		"form": uint(0),
		"type": "food",
	},
	"Pancalamity": {
		"id":   uint(667780),
		"form": uint(0),
		"type": "food",
	},
	"Pancallous": {
		"id":   uint(667781),
		"form": uint(0),
		"type": "food",
	},
	"Aloby": {
		"id":   uint(667782),
		"form": uint(0),
		"type": "water",
	},
	"Seraphmola": {
		"id":   uint(667783),
		"form": uint(0),
		"type": "water",
	},
	"Molashade": {
		"id":   uint(667784),
		"form": uint(0),
		"type": "water",
	},
	"Wbasculin": {
		"id":   uint(667785),
		"form": uint(0),
		"type": "ghost",
	},
	"Anglereist": {
		"id":   uint(667786),
		"form": uint(0),
		"type": "ghost",
	},
	"Fabdeer": {
		"id":   uint(667787),
		"form": uint(0),
		"type": "normal",
	},
	"Stormdler": {
		"id":   uint(667788),
		"form": uint(0),
		"type": "normal",
	},
	"Noseer": {
		"id":   uint(667789),
		"form": uint(0),
		"type": "normal",
	},
	"Smvoodoll": {
		"id":   uint(667790),
		"form": uint(0),
		"type": "normal",
	},
	"Himpaper": {
		"id":   uint(667791),
		"form": uint(0),
		"type": "paper",
	},
	"Himpaper-Open": {
		"id":   uint(667791),
		"form": uint(1),
		"type": "paper",
	},
	"Wargreymon": {
		"id":   uint(667792),
		"form": uint(0),
		"type": "dragon",
	},
	"Wargreymonv": {
		"id":   uint(667793),
		"form": uint(0),
		"type": "dragon",
	},
	"Victorygreymon": {
		"id":   uint(667794),
		"form": uint(0),
		"type": "dragon",
	},
	"Skullgreymon": {
		"id":   uint(667795),
		"form": uint(0),
		"type": "zombie",
	},
	"Phagus": {
		"id":   uint(667796),
		"form": uint(0),
		"type": "virus",
	},
	"Nexaphago": {
		"id":   uint(667797),
		"form": uint(0),
		"type": "virus",
	},
	"Himmagic": {
		"id":   uint(667798),
		"form": uint(0),
		"type": "magic",
	},
	"Standoge": {
		"id":   uint(667799),
		"form": uint(0),
		"type": "light",
	},
	"Standoger": {
		"id":   uint(667800),
		"form": uint(0),
		"type": "light",
	},
	"H1E1": {
		"id":   uint(667801),
		"form": uint(0),
		"type": "virus",
	},
	"E Boli": {
		"id":   uint(667802),
		"form": uint(0),
		"type": "virus",
	},
	"Charzar": {
		"id":   uint(667803),
		"form": uint(0),
		"type": "fire",
	},
	"Devilby": {
		"id":   uint(667804),
		"form": uint(0),
		"type": "dark",
	},
	"Devilmar": {
		"id":   uint(667805),
		"form": uint(0),
		"type": "dark",
	},
	"Hellmortar": {
		"id":   uint(667806),
		"form": uint(0),
		"type": "dark",
	},
	"Corneen": {
		"id":   uint(667807),
		"form": uint(0),
		"type": "food",
	},
	"Maze": {
		"id":   uint(667808),
		"form": uint(0),
		"type": "food",
	},
	"Ricardo Milos": {
		"id":   uint(667809),
		"form": uint(0),
		"type": "heart",
	},
	"Colosshale": {
		"id":   uint(667810),
		"form": uint(0),
		"type": "water",
	},
	"Cascavian": {
		"id":   uint(667811),
		"form": uint(0),
		"type": "water",
	},
	"Vitalimar": {
		"id":   uint(667812),
		"form": uint(0),
		"type": "fighting",
	},
	"Skulloton": {
		"id":   uint(667813),
		"form": uint(0),
		"type": "bone",
	},
	"Mmismagius": {
		"id":   uint(667814),
		"form": uint(0),
		"type": "ghost",
	},
	"Toyagumon": {
		"id":   uint(667815),
		"form": uint(0),
		"type": "plastic",
	},
	"Stoyagumon": {
		"id":   uint(667816),
		"form": uint(0),
		"type": "plastic",
	},
	"Hitleragumon": {
		"id":   uint(667817),
		"form": uint(0),
		"type": "fire",
	},
	"Biker": {
		"id":   uint(667818),
		"form": uint(0),
		"type": "normal",
	},
	"Cueball": {
		"id":   uint(667819),
		"form": uint(0),
		"type": "dark",
	},
	"Pokemaniac": {
		"id":   uint(667820),
		"form": uint(0),
		"type": "normal",
	},
	"Engineer": {
		"id":   uint(667821),
		"form": uint(0),
		"type": "tech",
	},
	"Gambler": {
		"id":   uint(667822),
		"form": uint(0),
		"type": "normal",
	},
	"Cool Couple": {
		"id":   uint(667823),
		"form": uint(0),
		"type": "normal",
	},
	"Ruin Maniac": {
		"id":   uint(667824),
		"form": uint(0),
		"type": "normal",
	},
	"Pkmn Ranger": {
		"id":   uint(667825),
		"form": uint(0),
		"type": "grass",
	},
	"Tamer": {
		"id":   uint(667826),
		"form": uint(0),
		"type": "normal",
	},
	"Super Nerd": {
		"id":   uint(667827),
		"form": uint(0),
		"type": "normal",
	},
	"Scientist": {
		"id":   uint(667828),
		"form": uint(0),
		"type": "tech",
	},
	"Cool Trainer": {
		"id":   uint(667829),
		"form": uint(0),
		"type": "normal",
	},
	"Santaltah": {
		"id":   uint(667830),
		"form": uint(0),
		"type": "ice",
	},
	"Biyomon": {
		"id":   uint(667831),
		"form": uint(0),
		"type": "cyber",
	},
	"Birdramon": {
		"id":   uint(667832),
		"form": uint(0),
		"type": "fire",
	},
	"Garudamon": {
		"id":   uint(667833),
		"form": uint(0),
		"type": "fire",
	},
	"Hououmon": {
		"id":   uint(667834),
		"form": uint(0),
		"type": "fire",
	},
	"Varodurumon": {
		"id":   uint(667835),
		"form": uint(0),
		"type": "light",
	},
	"Subaluga": {
		"id":   uint(667836),
		"form": uint(0),
		"type": "steel",
	},
	"Trapgeist": {
		"id":   uint(667837),
		"form": uint(0),
		"type": "steel",
	},
	"Troch": {
		"id":   uint(667838),
		"form": uint(0),
		"type": "fire",
	},
	"Floller": {
		"id":   uint(667839),
		"form": uint(0),
		"type": "fire",
	},
	"Tsunomon": {
		"id":   uint(667840),
		"form": uint(0),
		"type": "normal",
	},
	"Gabumon": {
		"id":   uint(667841),
		"form": uint(0),
		"type": "normal",
	},
	"Garurumon": {
		"id":   uint(667842),
		"form": uint(0),
		"type": "ice",
	},
	"Weregarurumon": {
		"id":   uint(667843),
		"form": uint(0),
		"type": "dark",
	},
	"Metalgarurumon": {
		"id":   uint(667844),
		"form": uint(0),
		"type": "cyber",
	},
	"Zeedgarurumon": {
		"id":   uint(667845),
		"form": uint(0),
		"type": "cyber",
	},
	"Motimon": {
		"id":   uint(667846),
		"form": uint(0),
		"type": "normal",
	},
	"Tentomon": {
		"id":   uint(667847),
		"form": uint(0),
		"type": "bug",
	},
	"Kabuterimon": {
		"id":   uint(667848),
		"form": uint(0),
		"type": "bug",
	},
	"Rmegakabuterimon": {
		"id":   uint(667849),
		"form": uint(0),
		"type": "bug",
	},
	"Herculeskabuterimon": {
		"id":   uint(667850),
		"form": uint(0),
		"type": "bug",
	},
	"Tyrantkabuterimon": {
		"id":   uint(667851),
		"form": uint(0),
		"type": "bug",
	},
	"Bmegakabuterimon": {
		"id":   uint(667852),
		"form": uint(0),
		"type": "bug",
	},
	"Napseel": {
		"id":   uint(667853),
		"form": uint(0),
		"type": "water",
	},
	"Frostiseel": {
		"id":   uint(667854),
		"form": uint(0),
		"type": "water",
	},
	"Shockaseel": {
		"id":   uint(667855),
		"form": uint(0),
		"type": "water",
	},
	"Retulzu": {
		"id":   uint(667856),
		"form": uint(0),
		"type": "fear",
	},
	"Retulzu-Zen": {
		"id":   uint(667856),
		"form": uint(1),
		"type": "fear",
	},
	"Spatulzu": {
		"id":   uint(667857),
		"form": uint(0),
		"type": "fear",
	},
	"Spatulzu-Zen": {
		"id":   uint(667857),
		"form": uint(1),
		"type": "fear",
	},
	"Palmon": {
		"id":   uint(667858),
		"form": uint(0),
		"type": "grass",
	},
	"Togemon": {
		"id":   uint(667859),
		"form": uint(0),
		"type": "grass",
	},
	"Lilymon": {
		"id":   uint(667860),
		"form": uint(0),
		"type": "grass",
	},
	"Rosemon": {
		"id":   uint(667861),
		"form": uint(0),
		"type": "grass",
	},
	"Babamon": {
		"id":   uint(667862),
		"form": uint(0),
		"type": "grass",
	},
	"Ponchomon": {
		"id":   uint(667863),
		"form": uint(0),
		"type": "grass",
	},
	"Kongdom": {
		"id":   uint(667864),
		"form": uint(0),
		"type": "grass",
	},
	"Scorknight": {
		"id":   uint(667865),
		"form": uint(0),
		"type": "fire",
	},
	"Drakobble": {
		"id":   uint(667866),
		"form": uint(0),
		"type": "water",
	},
	"Lurquad": {
		"id":   uint(667867),
		"form": uint(0),
		"type": "normal",
	},
	"Disfunbot": {
		"id":   uint(667868),
		"form": uint(0),
		"type": "cyber",
	},
	"Sheridryoid": {
		"id":   uint(667869),
		"form": uint(0),
		"type": "tech",
	},
	"Kizzeep": {
		"id":   uint(667870),
		"form": uint(0),
		"type": "heart",
	},
	"Xiaconisect": {
		"id":   uint(667871),
		"form": uint(0),
		"type": "bug",
	},
	"Gdiglett": {
		"id":   uint(667872),
		"form": uint(0),
		"type": "ghost",
	},
	"Gdugtrio": {
		"id":   uint(667873),
		"form": uint(0),
		"type": "ghost",
	},
	"Dkeldeo": {
		"id":   uint(667874),
		"form": uint(0),
		"type": "fire",
	},
	"Himice": {
		"id":   uint(667875),
		"form": uint(0),
		"type": "ice",
	},
	"Necroseel": {
		"id":   uint(667876),
		"form": uint(0),
		"type": "water",
	},
	"Gloseel": {
		"id":   uint(667877),
		"form": uint(0),
		"type": "water",
	},
	"Nursebola": {
		"id":   uint(667878),
		"form": uint(0),
		"type": "virus",
	},
	"Marsech": {
		"id":   uint(667879),
		"form": uint(0),
		"type": "cosmic",
	},
	"Himnuclear": {
		"id":   uint(667880),
		"form": uint(0),
		"type": "nuclear",
	},
	"Amiibobowser": {
		"id":   uint(667881),
		"form": uint(0),
		"type": "plastic",
	},
	"Cnosepass": {
		"id":   uint(667882),
		"form": uint(0),
		"type": "crystal",
	},
	"Cprobopass": {
		"id":   uint(667883),
		"form": uint(0),
		"type": "crystal",
	},
	"Ccarbink": {
		"id":   uint(667884),
		"form": uint(0),
		"type": "crystal",
	},
	"Qwilptain": {
		"id":   uint(667885),
		"form": uint(0),
		"type": "water",
	},
	"Penguinmon": {
		"id":   uint(667886),
		"form": uint(0),
		"type": "water",
	},
	"Dolphmon": {
		"id":   uint(667887),
		"form": uint(0),
		"type": "water",
	},
	"Whamon": {
		"id":   uint(667888),
		"form": uint(0),
		"type": "water",
	},
	"Neptunemon": {
		"id":   uint(667889),
		"form": uint(0),
		"type": "water",
	},
	"Marineangemon": {
		"id":   uint(667890),
		"form": uint(0),
		"type": "water",
	},
	"Rprocker": {
		"id":   uint(667891),
		"form": uint(0),
		"type": "tech",
	},
	"Rprockerz": {
		"id":   uint(667892),
		"form": uint(0),
		"type": "tech",
	},
	"Rprazor": {
		"id":   uint(667893),
		"form": uint(0),
		"type": "tech",
	},
	"Rprazorz": {
		"id":   uint(667894),
		"form": uint(0),
		"type": "tech",
	},
	"Rpboba": {
		"id":   uint(667895),
		"form": uint(0),
		"type": "tech",
	},
	"Rpmochi": {
		"id":   uint(667896),
		"form": uint(0),
		"type": "wood",
	},
	"Rpsunny": {
		"id":   uint(667897),
		"form": uint(0),
		"type": "tech",
	},
	"Rpsun02": {
		"id":   uint(667898),
		"form": uint(0),
		"type": "tech",
	},
	"Rpsuncust": {
		"id":   uint(667899),
		"form": uint(0),
		"type": "tech",
	},
	"Rpmedbot": {
		"id":   uint(667900),
		"form": uint(0),
		"type": "tech",
	},
	"Rpmeddy": {
		"id":   uint(667901),
		"form": uint(0),
		"type": "tech",
	},
	"Rpmeddyz": {
		"id":   uint(667902),
		"form": uint(0),
		"type": "tech",
	},
	"Rpdocbot": {
		"id":   uint(667903),
		"form": uint(0),
		"type": "tech",
	},
	"Wtrubbish": {
		"id":   uint(667904),
		"form": uint(0),
		"type": "food",
	},
	"Wgarbodor": {
		"id":   uint(667905),
		"form": uint(0),
		"type": "food",
	},
	"Cottandy": {
		"id":   uint(667906),
		"form": uint(0),
		"type": "food",
	},
	"Wgulpin": {
		"id":   uint(667907),
		"form": uint(0),
		"type": "ice",
	},
	"Wswalot": {
		"id":   uint(667908),
		"form": uint(0),
		"type": "ice",
	},
	"Swallice": {
		"id":   uint(667909),
		"form": uint(0),
		"type": "ice",
	},
	"Wapplin": {
		"id":   uint(667910),
		"form": uint(0),
		"type": "food",
	},
	"Wflapple": {
		"id":   uint(667911),
		"form": uint(0),
		"type": "food",
	},
	"Wappletun": {
		"id":   uint(667912),
		"form": uint(0),
		"type": "food",
	},
	"Charizardo": {
		"id":   uint(667913),
		"form": uint(0),
		"type": "fire",
	},
	"Dinofur": {
		"id":   uint(667914),
		"form": uint(0),
		"type": "fabric",
	},
	"Furrygon": {
		"id":   uint(667915),
		"form": uint(0),
		"type": "fabric",
	},
	"Furror": {
		"id":   uint(667916),
		"form": uint(0),
		"type": "fabric",
	},
	"Dreadark": {
		"id":   uint(667917),
		"form": uint(0),
		"type": "fear",
	},
	"Stareha": {
		"id":   uint(667918),
		"form": uint(0),
		"type": "water",
	},
	"Whattena": {
		"id":   uint(667919),
		"form": uint(0),
		"type": "heart",
	},
	"Whattrem": {
		"id":   uint(667920),
		"form": uint(0),
		"type": "heart",
	},
	"Whatterene": {
		"id":   uint(667921),
		"form": uint(0),
		"type": "heart",
	},
	"Wgothita": {
		"id":   uint(667922),
		"form": uint(0),
		"type": "water",
	},
	"Sailorita": {
		"id":   uint(667923),
		"form": uint(0),
		"type": "water",
	},
	"Admirelle": {
		"id":   uint(667924),
		"form": uint(0),
		"type": "water",
	},
	"Paintgon": {
		"id":   uint(667925),
		"form": uint(0),
		"type": "paint",
	},
	"Acrylomodo": {
		"id":   uint(667926),
		"form": uint(0),
		"type": "paint",
	},
	"Rattastic": {
		"id":   uint(667927),
		"form": uint(0),
		"type": "psychic",
	},
	"Edgucate": {
		"id":   uint(667928),
		"form": uint(0),
		"type": "psychic",
	},
	"Passta": {
		"id":   uint(667929),
		"form": uint(0),
		"type": "fire",
	},
	"Sspaghetti": {
		"id":   uint(667930),
		"form": uint(0),
		"type": "fire",
	},
	"Barkbark": {
		"id":   uint(667931),
		"form": uint(0),
		"type": "grass",
	},
	"Timbark": {
		"id":   uint(667932),
		"form": uint(0),
		"type": "grass",
	},
	"Floatskull": {
		"id":   uint(667933),
		"form": uint(0),
		"type": "ghost",
	},
	"Bloodetan": {
		"id":   uint(667934),
		"form": uint(0),
		"type": "ghost",
	},
	"Bloodoom": {
		"id":   uint(667935),
		"form": uint(0),
		"type": "ghost",
	},
	"Pistoff": {
		"id":   uint(667936),
		"form": uint(0),
		"type": "steel",
	},
	"Porkopter": {
		"id":   uint(667937),
		"form": uint(0),
		"type": "dragon",
	},
	"Bacopter": {
		"id":   uint(667938),
		"form": uint(0),
		"type": "dragon",
	},
	"Boombowl": {
		"id":   uint(667939),
		"form": uint(0),
		"type": "steel",
	},
	"Eggy": {
		"id":   uint(667940),
		"form": uint(0),
		"type": "normal",
	},
	"Eggception": {
		"id":   uint(667941),
		"form": uint(0),
		"type": "normal",
	},
	"Panslash": {
		"id":   uint(667942),
		"form": uint(0),
		"type": "dark",
	},
	"Simislash": {
		"id":   uint(667943),
		"form": uint(0),
		"type": "dark",
	},
	"Loccoon": {
		"id":   uint(667944),
		"form": uint(0),
		"type": "bug",
	},
	"Buzzerk": {
		"id":   uint(667945),
		"form": uint(0),
		"type": "bug",
	},
	"Pillost": {
		"id":   uint(667946),
		"form": uint(0),
		"type": "ghost",
	},
	"Shrieet": {
		"id":   uint(667947),
		"form": uint(0),
		"type": "ghost",
	},
	"Duvetious": {
		"id":   uint(667948),
		"form": uint(0),
		"type": "ghost",
	},
	"Bidoom": {
		"id":   uint(667949),
		"form": uint(0),
		"type": "normal",
	},
	"Kv 1": {
		"id":   uint(667950),
		"form": uint(0),
		"type": "steel",
	},
	"Rpramjet": {
		"id":   uint(667951),
		"form": uint(0),
		"type": "tech",
	},
	"Himfighting": {
		"id":   uint(667952),
		"form": uint(0),
		"type": "fighting",
	},
	"Rpmolbot": {
		"id":   uint(667953),
		"form": uint(0),
		"type": "ground",
	},
	"Rpdigger": {
		"id":   uint(667954),
		"form": uint(0),
		"type": "ground",
	},
	"Rpcookey": {
		"id":   uint(667955),
		"form": uint(0),
		"type": "tech",
	},
	"Tankdoge": {
		"id":   uint(667956),
		"form": uint(0),
		"type": "normal",
	},
	"Wnosepass": {
		"id":   uint(667957),
		"form": uint(0),
		"type": "wood",
	},
	"Wprobopass": {
		"id":   uint(667958),
		"form": uint(0),
		"type": "wood",
	},
	"Wexeggcute": {
		"id":   uint(667959),
		"form": uint(0),
		"type": "rock",
	},
	"Wexeggutor": {
		"id":   uint(667960),
		"form": uint(0),
		"type": "rock",
	},
	"Wkoffing": {
		"id":   uint(667961),
		"form": uint(0),
		"type": "cosmic",
	},
	"Wweezing": {
		"id":   uint(667962),
		"form": uint(0),
		"type": "cosmic",
	},
	"Rpboomer": {
		"id":   uint(667963),
		"form": uint(0),
		"type": "tech",
	},
	"Rpairraid": {
		"id":   uint(667964),
		"form": uint(0),
		"type": "tech",
	},
	"Rpboomb1": {
		"id":   uint(667965),
		"form": uint(0),
		"type": "tech",
	},
	"Rpgidget": {
		"id":   uint(667966),
		"form": uint(0),
		"type": "tech",
	},
	"Rpbetty": {
		"id":   uint(667967),
		"form": uint(0),
		"type": "tech",
	},
	"Rpbettyz": {
		"id":   uint(667968),
		"form": uint(0),
		"type": "tech",
	},
	"Rpgranny": {
		"id":   uint(667969),
		"form": uint(0),
		"type": "tech",
	},
	"Ssurshifu": {
		"id":   uint(667970),
		"form": uint(0),
		"type": "fighting",
	},
	"Rsurshifu": {
		"id":   uint(667971),
		"form": uint(0),
		"type": "fighting",
	},
	"Gslowpoke": {
		"id":   uint(667972),
		"form": uint(0),
		"type": "psychic",
	},
	"Himcyber": {
		"id":   uint(667973),
		"form": uint(0),
		"type": "cyber",
	},
	"Garticuno": {
		"id":   uint(667974),
		"form": uint(0),
		"type": "ice",
	},
	"Gzapdos": {
		"id":   uint(667975),
		"form": uint(0),
		"type": "electric",
	},
	"Gmoltres": {
		"id":   uint(667976),
		"form": uint(0),
		"type": "fire",
	},
	"Regielectric": {
		"id":   uint(667977),
		"form": uint(0),
		"type": "electric",
	},
	"Regidragon": {
		"id":   uint(667978),
		"form": uint(0),
		"type": "dragon",
	},
	"Tsunonasu": {
		"id":   uint(667979),
		"form": uint(0),
		"type": "rock",
	},
	"Gilerth": {
		"id":   uint(667980),
		"form": uint(0),
		"type": "rock",
	},
	"Gilgierth": {
		"id":   uint(667981),
		"form": uint(0),
		"type": "rock",
	},
	"Gigagigerth": {
		"id":   uint(667982),
		"form": uint(0),
		"type": "rock",
	},
	"Elsnow": {
		"id":   uint(667983),
		"form": uint(0),
		"type": "heart",
	},
	"Noofr": {
		"id":   uint(667984),
		"form": uint(0),
		"type": "chaos",
	},
	"Zodiacunown": {
		"id":   uint(667985),
		"form": uint(0),
		"type": "cosmic",
	},
	"Fireunown": {
		"id":   uint(667986),
		"form": uint(0),
		"type": "fire",
	},
	"Earthunown": {
		"id":   uint(667987),
		"form": uint(0),
		"type": "ground",
	},
	"Airunown": {
		"id":   uint(667988),
		"form": uint(0),
		"type": "wind",
	},
	"Waterunown": {
		"id":   uint(667989),
		"form": uint(0),
		"type": "water",
	},
	"Mfunown": {
		"id":   uint(667990),
		"form": uint(0),
		"type": "heart",
	},
	"Cspiritomb": {
		"id":   uint(667991),
		"form": uint(0),
		"type": "ghost",
	},
	"Rbsnorlax": {
		"id":   uint(667992),
		"form": uint(0),
		"type": "normal",
	},
	"Bdrifloon": {
		"id":   uint(667993),
		"form": uint(0),
		"type": "water",
	},
	"Bdrifblim": {
		"id":   uint(667994),
		"form": uint(0),
		"type": "water",
	},
	"Ivenonat": {
		"id":   uint(667995),
		"form": uint(0),
		"type": "bug",
	},
	"Ivenomoth": {
		"id":   uint(667996),
		"form": uint(0),
		"type": "bug",
	},
	"Ozodiacunown": {
		"id":   uint(667997),
		"form": uint(0),
		"type": "cosmic",
	},
	"Adrowzee": {
		"id":   uint(667998),
		"form": uint(0),
		"type": "ghost",
	},
	"Ahypno": {
		"id":   uint(667999),
		"form": uint(0),
		"type": "ghost",
	},
	"Maideedee": {
		"id":   uint(668000),
		"form": uint(0),
		"type": "psychic",
	},
	"Bylethemble": {
		"id":   uint(668001),
		"form": uint(0),
		"type": "magic",
	},
	"Curry": {
		"id":   uint(668002),
		"form": uint(0),
		"type": "food",
	},
	"Mcurry": {
		"id":   uint(668003),
		"form": uint(0),
		"type": "food",
	},
	"Bcurry": {
		"id":   uint(668004),
		"form": uint(0),
		"type": "food",
	},
	"Barkhol": {
		"id":   uint(668005),
		"form": uint(0),
		"type": "wood",
	},
	"Octopine": {
		"id":   uint(668006),
		"form": uint(0),
		"type": "wood",
	},
	"Yummydew": {
		"id":   uint(668007),
		"form": uint(0),
		"type": "dragon",
	},
	"Sucklaxis": {
		"id":   uint(668008),
		"form": uint(0),
		"type": "zombie",
	},
	"Honbrosia": {
		"id":   uint(668009),
		"form": uint(0),
		"type": "dragon",
	},
	"Pawnembra": {
		"id":   uint(668010),
		"form": uint(0),
		"type": "magic",
	},
	"Dusklaw": {
		"id":   uint(668011),
		"form": uint(0),
		"type": "magic",
	},
	"Femaledot": {
		"id":   uint(668012),
		"form": uint(0),
		"type": "normal",
	},
	"Froglide": {
		"id":   uint(668013),
		"form": uint(0),
		"type": "flying",
	},
	"Amphibifloat": {
		"id":   uint(668014),
		"form": uint(0),
		"type": "flying",
	},
	"Mfearow": {
		"id":   uint(668015),
		"form": uint(0),
		"type": "fear",
	},
	"Gomamon": {
		"id":   uint(668016),
		"form": uint(0),
		"type": "water",
	},
	"Ikkakumon": {
		"id":   uint(668017),
		"form": uint(0),
		"type": "water",
	},
	"Zudomon": {
		"id":   uint(668018),
		"form": uint(0),
		"type": "water",
	},
	"Vikemon": {
		"id":   uint(668019),
		"form": uint(0),
		"type": "ice",
	},
	"Himwind": {
		"id":   uint(668020),
		"form": uint(0),
		"type": "wind",
	},
	"Wdgolem": {
		"id":   uint(668021),
		"form": uint(0),
		"type": "rock",
	},
	"Wbarboach": {
		"id":   uint(668022),
		"form": uint(0),
		"type": "steam",
	},
	"Wwhiscash": {
		"id":   uint(668023),
		"form": uint(0),
		"type": "steam",
	},
	"Tokomon": {
		"id":   uint(668024),
		"form": uint(0),
		"type": "normal",
	},
	"Patamon": {
		"id":   uint(668025),
		"form": uint(0),
		"type": "normal",
	},
	"Angemon": {
		"id":   uint(668026),
		"form": uint(0),
		"type": "divine",
	},
	"Magnaangemon": {
		"id":   uint(668027),
		"form": uint(0),
		"type": "divine",
	},
	"Seraphimon": {
		"id":   uint(668028),
		"form": uint(0),
		"type": "divine",
	},
	"Salamon": {
		"id":   uint(668029),
		"form": uint(0),
		"type": "fairy",
	},
	"Tailmon": {
		"id":   uint(668030),
		"form": uint(0),
		"type": "fairy",
	},
	"Angewomon": {
		"id":   uint(668031),
		"form": uint(0),
		"type": "divine",
	},
	"Holydramon": {
		"id":   uint(668032),
		"form": uint(0),
		"type": "divine",
	},
	"Ophanimon": {
		"id":   uint(668033),
		"form": uint(0),
		"type": "divine",
	},
	"Ccleffa": {
		"id":   uint(668034),
		"form": uint(0),
		"type": "rock",
	},
	"Cclefairy": {
		"id":   uint(668035),
		"form": uint(0),
		"type": "rock",
	},
	"Cclefairy-Zen": {
		"id":   uint(668035),
		"form": uint(1),
		"type": "rock",
	},
	"Cclefable": {
		"id":   uint(668036),
		"form": uint(0),
		"type": "rock",
	},
	"Cclefable-Zen": {
		"id":   uint(668036),
		"form": uint(1),
		"type": "rock",
	},
	"Elecmon": {
		"id":   uint(668037),
		"form": uint(0),
		"type": "electric",
	},
	"Leomon": {
		"id":   uint(668038),
		"form": uint(0),
		"type": "fighting",
	},
	"Astaurios": {
		"id":   uint(668039),
		"form": uint(0),
		"type": "fighting",
	},
	"Grapleomon": {
		"id":   uint(668040),
		"form": uint(0),
		"type": "fighting",
	},
	"Saberleomon": {
		"id":   uint(668041),
		"form": uint(0),
		"type": "fighting",
	},
	"Doratgon": {
		"id":   uint(668042),
		"form": uint(0),
		"type": "normal",
	},
	"Kingdorah": {
		"id":   uint(668043),
		"form": uint(0),
		"type": "dragon",
	},
	"Candicegong": {
		"id":   uint(668044),
		"form": uint(0),
		"type": "normal",
	},
	"Gardebuck": {
		"id":   uint(668045),
		"form": uint(0),
		"type": "normal",
	},
	"Mistydos": {
		"id":   uint(668046),
		"form": uint(0),
		"type": "normal",
	},
	"Janinedos": {
		"id":   uint(668047),
		"form": uint(0),
		"type": "normal",
	},
	"Hapbray": {
		"id":   uint(668048),
		"form": uint(0),
		"type": "normal",
	},
	"Lucyper": {
		"id":   uint(668049),
		"form": uint(0),
		"type": "normal",
	},
	"Dgenesect": {
		"id":   uint(668050),
		"form": uint(0),
		"type": "bug",
	},
	"Pagumon": {
		"id":   uint(668051),
		"form": uint(0),
		"type": "dark",
	},
	"Goblimon": {
		"id":   uint(668052),
		"form": uint(0),
		"type": "dark",
	},
	"Ogremon": {
		"id":   uint(668053),
		"form": uint(0),
		"type": "dark",
	},
	"Fugamon": {
		"id":   uint(668054),
		"form": uint(0),
		"type": "dark",
	},
	"Hyogamon": {
		"id":   uint(668055),
		"form": uint(0),
		"type": "dark",
	},
	"Digitamamon": {
		"id":   uint(668056),
		"form": uint(0),
		"type": "cyber",
	},
	"Titamon": {
		"id":   uint(668057),
		"form": uint(0),
		"type": "chaos",
	},
	"Fuggirl": {
		"id":   uint(668058),
		"form": uint(0),
		"type": "dragon",
	},
	"Devitamamon": {
		"id":   uint(668059),
		"form": uint(0),
		"type": "chaos",
	},
	"Fdigitamamon": {
		"id":   uint(668060),
		"form": uint(0),
		"type": "magic",
	},
	"Minervamon": {
		"id":   uint(668061),
		"form": uint(0),
		"type": "divine",
	},
	"Relemon": {
		"id":   uint(668062),
		"form": uint(0),
		"type": "normal",
	},
	"Viximon": {
		"id":   uint(668063),
		"form": uint(0),
		"type": "normal",
	},
	"Renamon": {
		"id":   uint(668064),
		"form": uint(0),
		"type": "magic",
	},
	"Kyubimon": {
		"id":   uint(668065),
		"form": uint(0),
		"type": "magic",
	},
	"Taomon": {
		"id":   uint(668066),
		"form": uint(0),
		"type": "magic",
	},
	"Sakuyamon": {
		"id":   uint(668067),
		"form": uint(0),
		"type": "magic",
	},
	"Pageon": {
		"id":   uint(668068),
		"form": uint(0),
		"type": "paper",
	},
	"Wpachirisu": {
		"id":   uint(668069),
		"form": uint(0),
		"type": "fire",
	},
	"Mkachirisu": {
		"id":   uint(668070),
		"form": uint(0),
		"type": "fire",
	},
	"Fkachirisu": {
		"id":   uint(668071),
		"form": uint(0),
		"type": "fire",
	},
	"Mmienshao": {
		"id":   uint(668072),
		"form": uint(0),
		"type": "fighting",
	},
	"Himvirus": {
		"id":   uint(668073),
		"form": uint(0),
		"type": "virus",
	},
	"Samcurry": {
		"id":   uint(668074),
		"form": uint(0),
		"type": "food",
	},
	"Sabcurry": {
		"id":   uint(668075),
		"form": uint(0),
		"type": "food",
	},
	"Gkingler": {
		"id":   uint(668076),
		"form": uint(0),
		"type": "water",
	},
	"Dfpsyduck": {
		"id":   uint(668077),
		"form": uint(0),
		"type": "dark",
	},
	"Dfgolduck": {
		"id":   uint(668078),
		"form": uint(0),
		"type": "dark",
	},
	"Solcrab": {
		"id":   uint(668079),
		"form": uint(0),
		"type": "magma",
	},
	"Gordrarm": {
		"id":   uint(668080),
		"form": uint(0),
		"type": "dragon",
	},
	"Garagonarm": {
		"id":   uint(668081),
		"form": uint(0),
		"type": "dragon",
	},
	"Acolittle": {
		"id":   uint(668082),
		"form": uint(0),
		"type": "normal",
	},
	"Acthulist": {
		"id":   uint(668083),
		"form": uint(0),
		"type": "fear",
	},
	"Acolthep": {
		"id":   uint(668084),
		"form": uint(0),
		"type": "fear",
	},
	"Acosoth": {
		"id":   uint(668085),
		"form": uint(0),
		"type": "fear",
	},
	"Acultgoat": {
		"id":   uint(668086),
		"form": uint(0),
		"type": "fear",
	},
	"Aculthqua": {
		"id":   uint(668087),
		"form": uint(0),
		"type": "fear",
	},
	"Aculggua": {
		"id":   uint(668088),
		"form": uint(0),
		"type": "fear",
	},
	"Aculyig": {
		"id":   uint(668089),
		"form": uint(0),
		"type": "fear",
	},
	"Aculyellow": {
		"id":   uint(668090),
		"form": uint(0),
		"type": "fear",
	},
	"Acultchotcho": {
		"id":   uint(668091),
		"form": uint(0),
		"type": "fear",
	},
	"Aculzoth": {
		"id":   uint(668092),
		"form": uint(0),
		"type": "fear",
	},
	"Regicrawl": {
		"id":   uint(668093),
		"form": uint(0),
		"type": "normal",
	},
	"Aculthuga": {
		"id":   uint(668094),
		"form": uint(0),
		"type": "fear",
	},
	"Voormisan": {
		"id":   uint(668095),
		"form": uint(0),
		"type": "greasy",
	},
	"Teaplasm": {
		"id":   uint(668096),
		"form": uint(0),
		"type": "ghost",
	},
	"Maglamp": {
		"id":   uint(668097),
		"form": uint(0),
		"type": "magma",
	},
	"Shatterene": {
		"id":   uint(668098),
		"form": uint(0),
		"type": "water",
	},
	"Whoredevoir": {
		"id":   uint(668099),
		"form": uint(0),
		"type": "psychic",
	},
	"Steamingo": {
		"id":   uint(668100),
		"form": uint(0),
		"type": "steam",
	},
	"Wyrmaiden": {
		"id":   uint(668101),
		"form": uint(0),
		"type": "heart",
	},
	"Himtech": {
		"id":   uint(668102),
		"form": uint(0),
		"type": "tech",
	},
	"Tfkochia": {
		"id":   uint(668103),
		"form": uint(0),
		"type": "flying",
	},
	"Tfraygirth": {
		"id":   uint(668104),
		"form": uint(0),
		"type": "flying",
	},
	"Tfraygoten": {
		"id":   uint(668105),
		"form": uint(0),
		"type": "flying",
	},
	"Dratyke": {
		"id":   uint(668106),
		"form": uint(0),
		"type": "normal",
	},
	"Mdrampa": {
		"id":   uint(668107),
		"form": uint(0),
		"type": "time",
	},
	"Mbruxish": {
		"id":   uint(668108),
		"form": uint(0),
		"type": "water",
	},
	"Rfalinks": {
		"id":   uint(668109),
		"form": uint(0),
		"type": "steel",
	},
	"Magiroll": {
		"id":   uint(668110),
		"form": uint(0),
		"type": "water",
	},
	"Ebichu": {
		"id":   uint(668111),
		"form": uint(0),
		"type": "normal",
	},
	"Shrawk": {
		"id":   uint(668112),
		"form": uint(0),
		"type": "fighting",
	},
	"Amariantoinette": {
		"id":   uint(668113),
		"form": uint(0),
		"type": "glass",
	},
	"Himghost": {
		"id":   uint(668114),
		"form": uint(0),
		"type": "ghost",
	},
	"Selizidol": {
		"id":   uint(668115),
		"form": uint(0),
		"type": "dragon",
	},
	"Hinaspin": {
		"id":   uint(668116),
		"form": uint(0),
		"type": "poison",
	},
	"Dsnorunt": {
		"id":   uint(668117),
		"form": uint(0),
		"type": "dark",
	},
	"Dglalie": {
		"id":   uint(668118),
		"form": uint(0),
		"type": "dark",
	},
	"Dfroslass": {
		"id":   uint(668119),
		"form": uint(0),
		"type": "dark",
	},
	"Dpsyduck": {
		"id":   uint(668120),
		"form": uint(0),
		"type": "water",
	},
	"Dgolduck": {
		"id":   uint(668121),
		"form": uint(0),
		"type": "water",
	},
	"Lobotite": {
		"id":   uint(668122),
		"form": uint(0),
		"type": "virus",
	},
	"Lobotitan": {
		"id":   uint(668123),
		"form": uint(0),
		"type": "virus",
	},
	"Lobautomous": {
		"id":   uint(668124),
		"form": uint(0),
		"type": "virus",
	},
	"Steamling": {
		"id":   uint(668125),
		"form": uint(0),
		"type": "steam",
	},
	"Burniel": {
		"id":   uint(668126),
		"form": uint(0),
		"type": "fire",
	},
	"Zophlare": {
		"id":   uint(668127),
		"form": uint(0),
		"type": "fire",
	},
	"Ywwhisper": {
		"id":   uint(668128),
		"form": uint(0),
		"type": "ghost",
	},
	"Spongechu": {
		"id":   uint(668129),
		"form": uint(0),
		"type": "electric",
	},
	"Dralts": {
		"id":   uint(668130),
		"form": uint(0),
		"type": "dark",
	},
	"Dkirlia": {
		"id":   uint(668131),
		"form": uint(0),
		"type": "dark",
	},
	"Dgardevoir": {
		"id":   uint(668132),
		"form": uint(0),
		"type": "dark",
	},
	"Dgallade": {
		"id":   uint(668133),
		"form": uint(0),
		"type": "dark",
	},
	"Blatsoup": {
		"id":   uint(668134),
		"form": uint(0),
		"type": "blood",
	},
	"Coronablat": {
		"id":   uint(668135),
		"form": uint(0),
		"type": "virus",
	},
	"Muraship": {
		"id":   uint(668136),
		"form": uint(0),
		"type": "water",
	},
	"Tokikok": {
		"id":   uint(668137),
		"form": uint(0),
		"type": "flying",
	},
	"Kedama": {
		"id":   uint(668138),
		"form": uint(0),
		"type": "flying",
	},
	"Nuealien": {
		"id":   uint(668139),
		"form": uint(0),
		"type": "cosmic",
	},
	"Akyuu": {
		"id":   uint(668140),
		"form": uint(0),
		"type": "normal",
	},
	"Himcosmic": {
		"id":   uint(668141),
		"form": uint(0),
		"type": "cosmic",
	},
	"Whoppip": {
		"id":   uint(668142),
		"form": uint(0),
		"type": "glass",
	},
	"Wskiploom": {
		"id":   uint(668143),
		"form": uint(0),
		"type": "glass",
	},
	"Wjumpluff": {
		"id":   uint(668144),
		"form": uint(0),
		"type": "glass",
	},
	"Himground": {
		"id":   uint(668145),
		"form": uint(0),
		"type": "ground",
	},
	"Himfood": {
		"id":   uint(668146),
		"form": uint(0),
		"type": "food",
	},
	"Himzombie": {
		"id":   uint(668147),
		"form": uint(0),
		"type": "zombie",
	},
	"Aukomala": {
		"id":   uint(668148),
		"form": uint(0),
		"type": "normal",
	},
	"Zombiefairy": {
		"id":   uint(668149),
		"form": uint(0),
		"type": "fairy",
	},
	"Tenmonarch": {
		"id":   uint(668150),
		"form": uint(0),
		"type": "flying",
	},
	"Sendaimom": {
		"id":   uint(668151),
		"form": uint(0),
		"type": "fighting",
	},
	"Yoshianshi": {
		"id":   uint(668152),
		"form": uint(0),
		"type": "zombie",
	},
	"Seiganyan": {
		"id":   uint(668153),
		"form": uint(0),
		"type": "ghost",
	},
	"Shanghaioll": {
		"id":   uint(668154),
		"form": uint(0),
		"type": "psychic",
	},
	"Houraioll": {
		"id":   uint(668155),
		"form": uint(0),
		"type": "steel",
	},
	"Goliatholl": {
		"id":   uint(668156),
		"form": uint(0),
		"type": "psychic",
	},
	"Tojiko": {
		"id":   uint(668157),
		"form": uint(0),
		"type": "ghost",
	},
	"Orangeuard": {
		"id":   uint(668158),
		"form": uint(0),
		"type": "normal",
	},
	"Alterlanta": {
		"id":   uint(668159),
		"form": uint(0),
		"type": "chaos",
	},
	"Nursingale": {
		"id":   uint(668160),
		"form": uint(0),
		"type": "normal",
	},
	"Charlesanson": {
		"id":   uint(668161),
		"form": uint(0),
		"type": "dark",
	},
	"Genjiturtle": {
		"id":   uint(668162),
		"form": uint(0),
		"type": "ground",
	},
	"Greatcatfish": {
		"id":   uint(668163),
		"form": uint(0),
		"type": "wind",
	},
	"Thtori": {
		"id":   uint(668164),
		"form": uint(0),
		"type": "flying",
	},
	"Youkirai": {
		"id":   uint(668165),
		"form": uint(0),
		"type": "steel",
	},
	"Mimi Chan": {
		"id":   uint(668166),
		"form": uint(0),
		"type": "tech",
	},
	"Thrukoto": {
		"id":   uint(668167),
		"form": uint(0),
		"type": "tech",
	},
	"Chernophoot": {
		"id":   uint(668168),
		"form": uint(0),
		"type": "steam",
	},
	"Meirablade": {
		"id":   uint(668169),
		"form": uint(0),
		"type": "steel",
	},
	"Crikatank": {
		"id":   uint(668170),
		"form": uint(0),
		"type": "steel",
	},
	"Rikatank": {
		"id":   uint(668171),
		"form": uint(0),
		"type": "steel",
	},
	"Himdragon": {
		"id":   uint(668172),
		"form": uint(0),
		"type": "dragon",
	},
	"Ellenhu": {
		"id":   uint(668173),
		"form": uint(0),
		"type": "heart",
	},
	"Kanabarel": {
		"id":   uint(668174),
		"form": uint(0),
		"type": "ghost",
	},
	"Rikaksearch": {
		"id":   uint(668175),
		"form": uint(0),
		"type": "psychic",
	},
	"Kurumire": {
		"id":   uint(668176),
		"form": uint(0),
		"type": "dark",
	},
	"Saraguard": {
		"id":   uint(668177),
		"form": uint(0),
		"type": "fighting",
	},
	"Louisehu": {
		"id":   uint(668178),
		"form": uint(0),
		"type": "heart",
	},
	"Yukifire": {
		"id":   uint(668179),
		"form": uint(0),
		"type": "fire",
	},
	"Maiice": {
		"id":   uint(668180),
		"form": uint(0),
		"type": "ice",
	},
	"Yumekaid": {
		"id":   uint(668181),
		"form": uint(0),
		"type": "steel",
	},
	"Orinrincat": {
		"id":   uint(668182),
		"form": uint(0),
		"type": "fire",
	},
	"Suikamini": {
		"id":   uint(668183),
		"form": uint(0),
		"type": "fire",
	},
	"Rinnosudude": {
		"id":   uint(668184),
		"form": uint(0),
		"type": "steel",
	},
	"Tensokuhu": {
		"id":   uint(668185),
		"form": uint(0),
		"type": "steel",
	},
	"Ayakashtree": {
		"id":   uint(668186),
		"form": uint(0),
		"type": "ghost",
	},
	"Fmagicstones": {
		"id":   uint(668187),
		"form": uint(0),
		"type": "crystal",
	},
	"Twohu": {
		"id":   uint(668188),
		"form": uint(0),
		"type": "normal",
	},
	"Cmaribel": {
		"id":   uint(668189),
		"form": uint(0),
		"type": "normal",
	},
	"Thmaribel": {
		"id":   uint(668190),
		"form": uint(0),
		"type": "normal",
	},
	"Crenko": {
		"id":   uint(668191),
		"form": uint(0),
		"type": "normal",
	},
	"Threnko": {
		"id":   uint(668192),
		"form": uint(0),
		"type": "normal",
	},
	"Wakamaid": {
		"id":   uint(668193),
		"form": uint(0),
		"type": "water",
	},
	"Sekibankead": {
		"id":   uint(668194),
		"form": uint(0),
		"type": "dark",
	},
	"Fionnancer": {
		"id":   uint(668195),
		"form": uint(0),
		"type": "fighting",
	},
	"Sakuraber": {
		"id":   uint(668196),
		"form": uint(0),
		"type": "steel",
	},
	"Nerobaltah": {
		"id":   uint(668197),
		"form": uint(0),
		"type": "light",
	},
	"Boudichariot": {
		"id":   uint(668198),
		"form": uint(0),
		"type": "fighting",
	},
	"Nerosummer": {
		"id":   uint(668199),
		"form": uint(0),
		"type": "water",
	},
	"Caligulunatic": {
		"id":   uint(668200),
		"form": uint(0),
		"type": "cosmic",
	},
	"Cwarshin": {
		"id":   uint(668201),
		"form": uint(0),
		"type": "water",
	},
	"Orionartemis": {
		"id":   uint(668202),
		"form": uint(0),
		"type": "divine",
	},
	"Jungrilla": {
		"id":   uint(668203),
		"form": uint(0),
		"type": "grass",
	},
	"Eevee Girl": {
		"id":   uint(668204),
		"form": uint(0),
		"type": "normal",
	},
	"Doomseye": {
		"id":   uint(668205),
		"form": uint(0),
		"type": "chaos",
	},
	"Floweyr": {
		"id":   uint(668206),
		"form": uint(0),
		"type": "normal",
	},
	"Asrielgot": {
		"id":   uint(668207),
		"form": uint(0),
		"type": "normal",
	},
	"Chrislumbus": {
		"id":   uint(668208),
		"form": uint(0),
		"type": "water",
	},
	"Floweyrst": {
		"id":   uint(668209),
		"form": uint(0),
		"type": "normal",
	},
	"Dwelder": {
		"id":   uint(668210),
		"form": uint(0),
		"type": "steel",
	},
	"Weldigger": {
		"id":   uint(668211),
		"form": uint(0),
		"type": "steel",
	},
	"Himchaos": {
		"id":   uint(668212),
		"form": uint(0),
		"type": "chaos",
	},
	"Himdivine": {
		"id":   uint(668213),
		"form": uint(0),
		"type": "divine",
	},
	"Ngrimer": {
		"id":   uint(668214),
		"form": uint(0),
		"type": "poison",
	},
	"Nmuk": {
		"id":   uint(668215),
		"form": uint(0),
		"type": "poison",
	},
	"Catnobanana": {
		"id":   uint(668216),
		"form": uint(0),
		"type": "qmarks",
	},
	"Domineko": {
		"id":   uint(668217),
		"form": uint(0),
		"type": "heart",
	},
	"Doomsayw": {
		"id":   uint(668218),
		"form": uint(0),
		"type": "ghost",
	},
	"Sdoomsday": {
		"id":   uint(668219),
		"form": uint(0),
		"type": "ghost",
	},
	"Milunatone": {
		"id":   uint(668220),
		"form": uint(0),
		"type": "rock",
	},
	"Milunatone-Zen": {
		"id":   uint(668220),
		"form": uint(1),
		"type": "light",
	},
	"Misolrock": {
		"id":   uint(668221),
		"form": uint(0),
		"type": "rock",
	},
	"Misolrock-Zen": {
		"id":   uint(668221),
		"form": uint(1),
		"type": "light",
	},
	"Weepinbarf": {
		"id":   uint(668222),
		"form": uint(0),
		"type": "grass",
	},
	"Tmagnemite": {
		"id":   uint(668223),
		"form": uint(0),
		"type": "electric",
	},
	"Tmagneton": {
		"id":   uint(668224),
		"form": uint(0),
		"type": "electric",
	},
	"Bjigglypuff": {
		"id":   uint(668225),
		"form": uint(0),
		"type": "rubber",
	},
	"Bwigglytuff": {
		"id":   uint(668226),
		"form": uint(0),
		"type": "rubber",
	},
	"Truehim": {
		"id":   uint(668227),
		"form": uint(0),
		"type": "qmarks",
	},
	"Himless": {
		"id":   uint(668228),
		"form": uint(0),
		"type": "qmarks",
	},
	"Himwall": {
		"id":   uint(668229),
		"form": uint(0),
		"type": "qmarks",
	},
	"Himanne": {
		"id":   uint(668230),
		"form": uint(0),
		"type": "qmarks",
	},
	"Himzaydolf": {
		"id":   uint(668231),
		"form": uint(0),
		"type": "qmarks",
	},
	"Himvoid": {
		"id":   uint(668232),
		"form": uint(0),
		"type": "void",
	},
	"Fhorsea": {
		"id":   uint(668233),
		"form": uint(0),
		"type": "water",
	},
	"Fseadra": {
		"id":   uint(668234),
		"form": uint(0),
		"type": "water",
	},
	"Fkingdra": {
		"id":   uint(668235),
		"form": uint(0),
		"type": "water",
	},
	"Corviknighteen": {
		"id":   uint(668236),
		"form": uint(0),
		"type": "flying",
	},
	"Beholdie": {
		"id":   uint(668237),
		"form": uint(0),
		"type": "magic",
	},
	"Beholdeye": {
		"id":   uint(668238),
		"form": uint(0),
		"type": "magic",
	},
	"Teyerant": {
		"id":   uint(668239),
		"form": uint(0),
		"type": "magic",
	},
	"Gassporeye": {
		"id":   uint(668240),
		"form": uint(0),
		"type": "grass",
	},
	"Himplastic": {
		"id":   uint(668241),
		"form": uint(0),
		"type": "plastic",
	},
	"Solahog": {
		"id":   uint(668242),
		"form": uint(0),
		"type": "grass",
	},
	"Cheslios": {
		"id":   uint(668243),
		"form": uint(0),
		"type": "grass",
	},
	"Neofen": {
		"id":   uint(668244),
		"form": uint(0),
		"type": "fire",
	},
	"Omephox": {
		"id":   uint(668245),
		"form": uint(0),
		"type": "fire",
	},
	"Simprog": {
		"id":   uint(668246),
		"form": uint(0),
		"type": "water",
	},
	"Neoille": {
		"id":   uint(668247),
		"form": uint(0),
		"type": "water",
	},
	"Pumpird": {
		"id":   uint(668248),
		"form": uint(0),
		"type": "tech",
	},
	"Pumpjackird": {
		"id":   uint(668249),
		"form": uint(0),
		"type": "tech",
	},
	"Gigimon": {
		"id":   uint(668250),
		"form": uint(0),
		"type": "dragon",
	},
	"Guilmon": {
		"id":   uint(668251),
		"form": uint(0),
		"type": "dragon",
	},
	"Growlmon": {
		"id":   uint(668252),
		"form": uint(0),
		"type": "dragon",
	},
	"Wargrowlmon": {
		"id":   uint(668253),
		"form": uint(0),
		"type": "dragon",
	},
	"Gallantmon": {
		"id":   uint(668254),
		"form": uint(0),
		"type": "fighting",
	},
	"Megidramon": {
		"id":   uint(668255),
		"form": uint(0),
		"type": "dragon",
	},
	"Larvipede": {
		"id":   uint(668256),
		"form": uint(0),
		"type": "bug",
	},
	"Sandrawler": {
		"id":   uint(668257),
		"form": uint(0),
		"type": "bug",
	},
	"Gummymon": {
		"id":   uint(668258),
		"form": uint(0),
		"type": "normal",
	},
	"Terriermon": {
		"id":   uint(668259),
		"form": uint(0),
		"type": "normal",
	},
	"Gargomon": {
		"id":   uint(668260),
		"form": uint(0),
		"type": "tech",
	},
	"Rapidmon": {
		"id":   uint(668261),
		"form": uint(0),
		"type": "tech",
	},
	"Megagargomon": {
		"id":   uint(668262),
		"form": uint(0),
		"type": "tech",
	},
	"Numemon": {
		"id":   uint(668263),
		"form": uint(0),
		"type": "poison",
	},
	"Sukamon": {
		"id":   uint(668264),
		"form": uint(0),
		"type": "poison",
	},
	"Platinumsukamon": {
		"id":   uint(668265),
		"form": uint(0),
		"type": "poison",
	},
	"Karatsukinumemon": {
		"id":   uint(668266),
		"form": uint(0),
		"type": "poison",
	},
	"Monzaemon": {
		"id":   uint(668267),
		"form": uint(0),
		"type": "fabric",
	},
	"Etemon": {
		"id":   uint(668268),
		"form": uint(0),
		"type": "dark",
	},
	"Metaletemon": {
		"id":   uint(668269),
		"form": uint(0),
		"type": "steel",
	},
	"Kingetemon": {
		"id":   uint(668270),
		"form": uint(0),
		"type": "light",
	},
	"Chuumon": {
		"id":   uint(668271),
		"form": uint(0),
		"type": "normal",
	},
	"Snowagumon": {
		"id":   uint(668272),
		"form": uint(0),
		"type": "ice",
	},
	"Blackagumon": {
		"id":   uint(668273),
		"form": uint(0),
		"type": "fire",
	},
	"Vgreymon": {
		"id":   uint(668274),
		"form": uint(0),
		"type": "fire",
	},
	"Snowgoblimon": {
		"id":   uint(668275),
		"form": uint(0),
		"type": "dark",
	},
	"Shamanmon": {
		"id":   uint(668276),
		"form": uint(0),
		"type": "fairy",
	},
	"Dotaugumon": {
		"id":   uint(668277),
		"form": uint(0),
		"type": "fire",
	},
	"Geremon": {
		"id":   uint(668278),
		"form": uint(0),
		"type": "poison",
	},
	"Otamamon": {
		"id":   uint(668279),
		"form": uint(0),
		"type": "water",
	},
	"Gekomon": {
		"id":   uint(668280),
		"form": uint(0),
		"type": "water",
	},
	"Shogungekomon": {
		"id":   uint(668281),
		"form": uint(0),
		"type": "water",
	},
	"Pukumon": {
		"id":   uint(668282),
		"form": uint(0),
		"type": "water",
	},
	"Nanimon": {
		"id":   uint(668283),
		"form": uint(0),
		"type": "poison",
	},
	"Bombernanimon": {
		"id":   uint(668284),
		"form": uint(0),
		"type": "steel",
	},
	"Snuggice": {
		"id":   uint(668285),
		"form": uint(0),
		"type": "ice",
	},
	"Sweateryeti": {
		"id":   uint(668286),
		"form": uint(0),
		"type": "ice",
	},
	"Frequenceon": {
		"id":   uint(668287),
		"form": uint(0),
		"type": "sound",
	},
	"Primplup": {
		"id":   uint(668288),
		"form": uint(0),
		"type": "water",
	},
	"Primpendlup": {
		"id":   uint(668289),
		"form": uint(0),
		"type": "water",
	},
	"Manaphake": {
		"id":   uint(668290),
		"form": uint(0),
		"type": "water",
	},
	"Magmortard": {
		"id":   uint(668291),
		"form": uint(0),
		"type": "fire",
	},
	"Grotlefake": {
		"id":   uint(668292),
		"form": uint(0),
		"type": "grass",
	},
	"Faketerra": {
		"id":   uint(668293),
		"form": uint(0),
		"type": "grass",
	},
	"Riodog": {
		"id":   uint(668294),
		"form": uint(0),
		"type": "fighting",
	},
	"Betalapras": {
		"id":   uint(668295),
		"form": uint(0),
		"type": "water",
	},
	"Saunarilla": {
		"id":   uint(668296),
		"form": uint(0),
		"type": "grass",
	},
	"Helialope": {
		"id":   uint(668297),
		"form": uint(0),
		"type": "fire",
	},
	"Societrex": {
		"id":   uint(668298),
		"form": uint(0),
		"type": "water",
	},
	"Noroiko": {
		"id":   uint(668299),
		"form": uint(0),
		"type": "ghost",
	},
	"Gabsol": {
		"id":   uint(668300),
		"form": uint(0),
		"type": "dark",
	},
	"Benbenee": {
		"id":   uint(668301),
		"form": uint(0),
		"type": "sound",
	},
	"Yatsusic": {
		"id":   uint(668302),
		"form": uint(0),
		"type": "sound",
	},
	"Raikodrum": {
		"id":   uint(668303),
		"form": uint(0),
		"type": "sound",
	},
	"Tornadeon": {
		"id":   uint(668304),
		"form": uint(0),
		"type": "wind",
	},
	"Ghost Maid": {
		"id":   uint(668305),
		"form": uint(0),
		"type": "normal",
	},
	"Thoothoot": {
		"id":   uint(668306),
		"form": uint(0),
		"type": "time",
	},
	"Tnoctowl": {
		"id":   uint(668307),
		"form": uint(0),
		"type": "time",
	},
	"Mprobopass": {
		"id":   uint(668308),
		"form": uint(0),
		"type": "rock",
	},
	"Adventcirnai": {
		"id":   uint(668309),
		"form": uint(0),
		"type": "ice",
	},
	"Advent Reisen": {
		"id":   uint(668310),
		"form": uint(0),
		"type": "tech",
	},
	"Advent Meiling": {
		"id":   uint(668311),
		"form": uint(0),
		"type": "fighting",
	},
	"Advent Marisa": {
		"id":   uint(668312),
		"form": uint(0),
		"type": "electric",
	},
	"Advent Alice": {
		"id":   uint(668313),
		"form": uint(0),
		"type": "magic",
	},
	"Advent Tewi": {
		"id":   uint(668314),
		"form": uint(0),
		"type": "normal",
	},
	"Advent Letty": {
		"id":   uint(668315),
		"form": uint(0),
		"type": "ice",
	},
	"Advent Mokanou": {
		"id":   uint(668316),
		"form": uint(0),
		"type": "fire",
	},
	"Advent Ran": {
		"id":   uint(668317),
		"form": uint(0),
		"type": "ground",
	},
	"Advent Chen": {
		"id":   uint(668318),
		"form": uint(0),
		"type": "ground",
	},
	"Advent Yukari": {
		"id":   uint(668319),
		"form": uint(0),
		"type": "psychic",
	},
	"Advent Mystia": {
		"id":   uint(668320),
		"form": uint(0),
		"type": "flying",
	},
	"Advent Youmu": {
		"id":   uint(668321),
		"form": uint(0),
		"type": "steel",
	},
	"Advent Yuyuko": {
		"id":   uint(668322),
		"form": uint(0),
		"type": "ghost",
	},
	"Mercary": {
		"id":   uint(668323),
		"form": uint(0),
		"type": "steel",
	},
	"Merqetal": {
		"id":   uint(668324),
		"form": uint(0),
		"type": "steel",
	},
	"Tunacury": {
		"id":   uint(668325),
		"form": uint(0),
		"type": "steel",
	},
	"Davistar": {
		"id":   uint(668326),
		"form": uint(0),
		"type": "cosmic",
	},
	"Woolshep": {
		"id":   uint(668327),
		"form": uint(0),
		"type": "normal",
	},
	"Matzio": {
		"id":   uint(668328),
		"form": uint(0),
		"type": "food",
	},
	"Matzahz": {
		"id":   uint(668329),
		"form": uint(0),
		"type": "food",
	},
	"Aeongel": {
		"id":   uint(668330),
		"form": uint(0),
		"type": "divine",
	},
	"Archongel": {
		"id":   uint(668331),
		"form": uint(0),
		"type": "divine",
	},
	"Sakiriff": {
		"id":   uint(668332),
		"form": uint(0),
		"type": "fighting",
	},
	"Deputymon": {
		"id":   uint(668333),
		"form": uint(0),
		"type": "tech",
	},
	"Billykid": {
		"id":   uint(668334),
		"form": uint(0),
		"type": "normal",
	},
	"Ringabbit": {
		"id":   uint(668335),
		"form": uint(0),
		"type": "food",
	},
	"Dreidelspi": {
		"id":   uint(668336),
		"form": uint(0),
		"type": "wood",
	},
	"Jgolett": {
		"id":   uint(668337),
		"form": uint(0),
		"type": "ground",
	},
	"Jgolurk": {
		"id":   uint(668338),
		"form": uint(0),
		"type": "ground",
	},
	"Wrufflet": {
		"id":   uint(668339),
		"form": uint(0),
		"type": "fighting",
	},
	"Wbraviary": {
		"id":   uint(668340),
		"form": uint(0),
		"type": "fighting",
	},
	"Hanukkora": {
		"id":   uint(668341),
		"form": uint(0),
		"type": "fire",
	},
	"Kaiju Ralts": {
		"id":   uint(668342),
		"form": uint(0),
		"type": "psychic",
	},
	"Kaiju Cacnea": {
		"id":   uint(668343),
		"form": uint(0),
		"type": "grass",
	},
	"Kaiju Seedot": {
		"id":   uint(668344),
		"form": uint(0),
		"type": "grass",
	},
	"Kaiju Machop": {
		"id":   uint(668345),
		"form": uint(0),
		"type": "fighting",
	},
	"Kaiju Venusaur": {
		"id":   uint(668346),
		"form": uint(0),
		"type": "grass",
	},
	"Kaiju Cutiefly": {
		"id":   uint(668347),
		"form": uint(0),
		"type": "bug",
	},
	"Kaiju Fearow": {
		"id":   uint(668348),
		"form": uint(0),
		"type": "normal",
	},
	"Kaiju Primarina": {
		"id":   uint(668349),
		"form": uint(0),
		"type": "water",
	},
	"Kaiju Latias": {
		"id":   uint(668350),
		"form": uint(0),
		"type": "dragon",
	},
	"Kaiju Natu": {
		"id":   uint(668351),
		"form": uint(0),
		"type": "psychic",
	},
	"Kaiju Flaaffy": {
		"id":   uint(668352),
		"form": uint(0),
		"type": "electric",
	},
	"Kaiju Girafarig": {
		"id":   uint(668353),
		"form": uint(0),
		"type": "normal",
	},
	"Kaiju Krabby": {
		"id":   uint(668354),
		"form": uint(0),
		"type": "water",
	},
	"Kaiju Minccino": {
		"id":   uint(668355),
		"form": uint(0),
		"type": "normal",
	},
	"Kaiju Cinccino": {
		"id":   uint(668356),
		"form": uint(0),
		"type": "normal",
	},
	"Kaiju Ditto": {
		"id":   uint(668357),
		"form": uint(0),
		"type": "normal",
	},
	"Kaiju Stoutland": {
		"id":   uint(668358),
		"form": uint(0),
		"type": "normal",
	},
	"Kaiju Quilladin": {
		"id":   uint(668359),
		"form": uint(0),
		"type": "grass",
	},
	"Kaiju Croagunk": {
		"id":   uint(668360),
		"form": uint(0),
		"type": "poison",
	},
	"Kaiju Umbreon": {
		"id":   uint(668361),
		"form": uint(0),
		"type": "dark",
	},
	"Kaiju Blacephalon": {
		"id":   uint(668362),
		"form": uint(0),
		"type": "fire",
	},
	"Kaiju Weavile": {
		"id":   uint(668363),
		"form": uint(0),
		"type": "dark",
	},
	"Kaiju Magnezone": {
		"id":   uint(668364),
		"form": uint(0),
		"type": "electric",
	},
	"Kaiju Klinklang": {
		"id":   uint(668365),
		"form": uint(0),
		"type": "steel",
	},
	"Kaiju Lileep": {
		"id":   uint(668366),
		"form": uint(0),
		"type": "rock",
	},
	"Corida": {
		"id":   uint(668367),
		"form": uint(0),
		"type": "fighting",
	},
	"Kungfloo": {
		"id":   uint(668368),
		"form": uint(0),
		"type": "fighting",
	},
	"Skirona": {
		"id":   uint(668369),
		"form": uint(0),
		"type": "rubber",
	},
	"Woodman": {
		"id":   uint(668370),
		"form": uint(0),
		"type": "chaos",
	},
	"Timbito": {
		"id":   uint(668371),
		"form": uint(0),
		"type": "food",
	},
	"Doughnutter": {
		"id":   uint(668372),
		"form": uint(0),
		"type": "food",
	},
	"Mavalugg": {
		"id":   uint(668373),
		"form": uint(0),
		"type": "ice",
	},
	"Himpoison": {
		"id":   uint(668374),
		"form": uint(0),
		"type": "poison",
	},
	"Mervamon": {
		"id":   uint(668375),
		"form": uint(0),
		"type": "divine",
	},
	"Mikemon": {
		"id":   uint(668376),
		"form": uint(0),
		"type": "normal",
	},
	"Bastemon": {
		"id":   uint(668377),
		"form": uint(0),
		"type": "heart",
	},
	"Lilithmon": {
		"id":   uint(668378),
		"form": uint(0),
		"type": "heart",
	},
	"Artificeon": {
		"id":   uint(668379),
		"form": uint(0),
		"type": "plastic",
	},
	"Transpareon": {
		"id":   uint(668380),
		"form": uint(0),
		"type": "glass",
	},
	"Carmillapire": {
		"id":   uint(668381),
		"form": uint(0),
		"type": "blood",
	},
	"Mozarter": {
		"id":   uint(668382),
		"form": uint(0),
		"type": "sound",
	},
	"Mashielder": {
		"id":   uint(668383),
		"form": uint(0),
		"type": "steel",
	},
	"Chevaleon": {
		"id":   uint(668384),
		"form": uint(0),
		"type": "fairy",
	},
	"Benkeilancer": {
		"id":   uint(668385),
		"form": uint(0),
		"type": "fighting",
	},
	"Kiyodere": {
		"id":   uint(668386),
		"form": uint(0),
		"type": "dragon",
	},
	"Mumcurry": {
		"id":   uint(668387),
		"form": uint(0),
		"type": "food",
	},
	"Mumbcurry": {
		"id":   uint(668388),
		"form": uint(0),
		"type": "food",
	},
	"Pacmcurry": {
		"id":   uint(668389),
		"form": uint(0),
		"type": "food",
	},
	"Pacbcurry": {
		"id":   uint(668390),
		"form": uint(0),
		"type": "food",
	},
	"Appmcurry": {
		"id":   uint(668391),
		"form": uint(0),
		"type": "food",
	},
	"Appbcurry": {
		"id":   uint(668392),
		"form": uint(0),
		"type": "food",
	},
	"Paracelcast": {
		"id":   uint(668393),
		"form": uint(0),
		"type": "magic",
	},
	"Martharider": {
		"id":   uint(668394),
		"form": uint(0),
		"type": "divine",
	},
	"Phantopera": {
		"id":   uint(668395),
		"form": uint(0),
		"type": "fear",
	},
	"Skelossil": {
		"id":   uint(668396),
		"form": uint(0),
		"type": "bone",
	},
	"Livepunch": {
		"id":   uint(668397),
		"form": uint(0),
		"type": "normal",
	},
	"Firepuncher": {
		"id":   uint(668398),
		"form": uint(0),
		"type": "normal",
	},
	"Kinopunch": {
		"id":   uint(668399),
		"form": uint(0),
		"type": "normal",
	},
	"Fukoujo": {
		"id":   uint(668400),
		"form": uint(0),
		"type": "blood",
	},
	"Mammarina": {
		"id":   uint(668401),
		"form": uint(0),
		"type": "heart",
	},
	"Kokopelli": {
		"id":   uint(668402),
		"form": uint(0),
		"type": "grass",
	},
	"Ballbu": {
		"id":   uint(668403),
		"form": uint(0),
		"type": "ice",
	},
	"Snowbun": {
		"id":   uint(668404),
		"form": uint(0),
		"type": "ice",
	},
	"Snoware": {
		"id":   uint(668405),
		"form": uint(0),
		"type": "ice",
	},
	"Snowugg": {
		"id":   uint(668406),
		"form": uint(0),
		"type": "ice",
	},
	"Larv": {
		"id":   uint(668407),
		"form": uint(0),
		"type": "rock",
	},
	"Zatu": {
		"id":   uint(668408),
		"form": uint(0),
		"type": "psychic",
	},
	"Sunsprout": {
		"id":   uint(668409),
		"form": uint(0),
		"type": "grass",
	},
	"Bwooper": {
		"id":   uint(668410),
		"form": uint(0),
		"type": "water",
	},
	"Feiscue": {
		"id":   uint(668411),
		"form": uint(0),
		"type": "fire",
	},
	"Feiscue-Noice": {
		"id":   uint(668411),
		"form": uint(1),
		"type": "fire",
	},
	"Wobbo": {
		"id":   uint(668412),
		"form": uint(0),
		"type": "ghost",
	},
	"Dunspa": {
		"id":   uint(668413),
		"form": uint(0),
		"type": "normal",
	},
	"Dunspit": {
		"id":   uint(668414),
		"form": uint(0),
		"type": "normal",
	},
	"Dunsnek": {
		"id":   uint(668415),
		"form": uint(0),
		"type": "normal",
	},
	"Azumarillegg": {
		"id":   uint(668416),
		"form": uint(0),
		"type": "water",
	},
	"Bchinchou": {
		"id":   uint(668417),
		"form": uint(0),
		"type": "water",
	},
	"Blanturn": {
		"id":   uint(668418),
		"form": uint(0),
		"type": "water",
	},
	"Clefmontop": {
		"id":   uint(668419),
		"form": uint(0),
		"type": "fighting",
	},
	"Bshuckle": {
		"id":   uint(668420),
		"form": uint(0),
		"type": "bug",
	},
	"Beforretress": {
		"id":   uint(668421),
		"form": uint(0),
		"type": "bug",
	},
	"Bcyndaquil": {
		"id":   uint(668422),
		"form": uint(0),
		"type": "steel",
	},
	"Acyndaquil": {
		"id":   uint(668423),
		"form": uint(0),
		"type": "steel",
	},
	"Floxo": {
		"id":   uint(668424),
		"form": uint(0),
		"type": "flying",
	},
	"Fennefire": {
		"id":   uint(668425),
		"form": uint(0),
		"type": "fire",
	},
	"Leosolg": {
		"id":   uint(668426),
		"form": uint(0),
		"type": "ice",
	},
	"Snekindian": {
		"id":   uint(668427),
		"form": uint(0),
		"type": "grass",
	},
	"Keelisk": {
		"id":   uint(668428),
		"form": uint(0),
		"type": "dragon",
	},
	"Anvikin": {
		"id":   uint(668429),
		"form": uint(0),
		"type": "water",
	},
	"Molaee": {
		"id":   uint(668430),
		"form": uint(0),
		"type": "water",
	},
	"Molaugg": {
		"id":   uint(668431),
		"form": uint(0),
		"type": "water",
	},
	"Gostee": {
		"id":   uint(668432),
		"form": uint(0),
		"type": "ghost",
	},
	"Puppee": {
		"id":   uint(668433),
		"form": uint(0),
		"type": "normal",
	},
	"Hippsycho": {
		"id":   uint(668434),
		"form": uint(0),
		"type": "psychic",
	},
	"Boarstag": {
		"id":   uint(668435),
		"form": uint(0),
		"type": "ground",
	},
	"Elephian": {
		"id":   uint(668436),
		"form": uint(0),
		"type": "steel",
	},
	"Koalma": {
		"id":   uint(668437),
		"form": uint(0),
		"type": "normal",
	},
	"Kachimatch": {
		"id":   uint(668438),
		"form": uint(0),
		"type": "normal",
	},
	"Kiteana": {
		"id":   uint(668439),
		"form": uint(0),
		"type": "flying",
	},
	"Bladeagon": {
		"id":   uint(668440),
		"form": uint(0),
		"type": "dragon",
	},
	"Squidrill": {
		"id":   uint(668441),
		"form": uint(0),
		"type": "water",
	},
	"Quailnote": {
		"id":   uint(668442),
		"form": uint(0),
		"type": "normal",
	},
	"Choruseph": {
		"id":   uint(668443),
		"form": uint(0),
		"type": "sound",
	},
	"Dootbirb": {
		"id":   uint(668444),
		"form": uint(0),
		"type": "sound",
	},
	"Okaabirb": {
		"id":   uint(668445),
		"form": uint(0),
		"type": "normal",
	},
	"Crowscare": {
		"id":   uint(668446),
		"form": uint(0),
		"type": "grass",
	},
	"Keewee": {
		"id":   uint(668447),
		"form": uint(0),
		"type": "grass",
	},
	"Sneewii": {
		"id":   uint(668448),
		"form": uint(0),
		"type": "fairy",
	},
	"Flyfish": {
		"id":   uint(668449),
		"form": uint(0),
		"type": "water",
	},
	"WackLunabbit": {
		"id":   uint(668450),
		"form": uint(0),
		"type": "fairy",
	},
	"Qwilsmug": {
		"id":   uint(668451),
		"form": uint(0),
		"type": "water",
	},
	"Froggo": {
		"id":   uint(668452),
		"form": uint(0),
		"type": "water",
	},
	"Echocho": {
		"id":   uint(668453),
		"form": uint(0),
		"type": "normal",
	},
	"Bashkull": {
		"id":   uint(668454),
		"form": uint(0),
		"type": "rock",
	},
	"Virualower": {
		"id":   uint(668455),
		"form": uint(0),
		"type": "grass",
	},
	"Cocoobag": {
		"id":   uint(668456),
		"form": uint(0),
		"type": "bug",
	},
	"Mothca": {
		"id":   uint(668457),
		"form": uint(0),
		"type": "bug",
	},
	"Misfly": {
		"id":   uint(668458),
		"form": uint(0),
		"type": "bug",
	},
	"Beetul": {
		"id":   uint(668459),
		"form": uint(0),
		"type": "bug",
	},
	"Stingscorp": {
		"id":   uint(668460),
		"form": uint(0),
		"type": "bug",
	},
	"Bhoothoot": {
		"id":   uint(668461),
		"form": uint(0),
		"type": "dark",
	},
	"Bsmantine": {
		"id":   uint(668462),
		"form": uint(0),
		"type": "water",
	},
	"Bcmantine": {
		"id":   uint(668463),
		"form": uint(0),
		"type": "water",
	},
	"Bslugma": {
		"id":   uint(668464),
		"form": uint(0),
		"type": "fire",
	},
	"Bpichu": {
		"id":   uint(668465),
		"form": uint(0),
		"type": "electric",
	},
	"Bblissey": {
		"id":   uint(668466),
		"form": uint(0),
		"type": "normal",
	},
	"Blarvitar": {
		"id":   uint(668467),
		"form": uint(0),
		"type": "rock",
	},
	"Btyranitar": {
		"id":   uint(668468),
		"form": uint(0),
		"type": "rock",
	},
	"Wobslime": {
		"id":   uint(668469),
		"form": uint(0),
		"type": "psychic",
	},
	"Bpolitoed": {
		"id":   uint(668470),
		"form": uint(0),
		"type": "water",
	},
	"Bgligar": {
		"id":   uint(668471),
		"form": uint(0),
		"type": "ground",
	},
	"Caenislancer": {
		"id":   uint(668472),
		"form": uint(0),
		"type": "fighting",
	},
	"Summervoir": {
		"id":   uint(668473),
		"form": uint(0),
		"type": "psychic",
	},
	"Kaiju Scrafty": {
		"id":   uint(668474),
		"form": uint(0),
		"type": "dark",
	},
	"Kaiju Porygon": {
		"id":   uint(668475),
		"form": uint(0),
		"type": "normal",
	},
	"Kaiju Kabuto": {
		"id":   uint(668476),
		"form": uint(0),
		"type": "rock",
	},
	"Kaiju Happiny": {
		"id":   uint(668477),
		"form": uint(0),
		"type": "normal",
	},
	"Kaiju Pyukumuku": {
		"id":   uint(668478),
		"form": uint(0),
		"type": "water",
	},
	"Kaiju Hypno": {
		"id":   uint(668479),
		"form": uint(0),
		"type": "psychic",
	},
	"Kaiju Gligar": {
		"id":   uint(668480),
		"form": uint(0),
		"type": "ground",
	},
	"Kaiju Dunsparce": {
		"id":   uint(668481),
		"form": uint(0),
		"type": "normal",
	},
	"Kaiju Ursaring": {
		"id":   uint(668482),
		"form": uint(0),
		"type": "normal",
	},
	"Kaiju Eleko": {
		"id":   uint(668483),
		"form": uint(0),
		"type": "ground",
	},
	"Kaiju Goomy": {
		"id":   uint(668484),
		"form": uint(0),
		"type": "dragon",
	},
	"Kaiju Stufful": {
		"id":   uint(668485),
		"form": uint(0),
		"type": "normal",
	},
	"Kaiju Gumshoos": {
		"id":   uint(668486),
		"form": uint(0),
		"type": "normal",
	},
	"Kaiju Toucannon": {
		"id":   uint(668487),
		"form": uint(0),
		"type": "normal",
	},
	"Kaiju Kingler": {
		"id":   uint(668488),
		"form": uint(0),
		"type": "water",
	},
	"Kaiju Machoke": {
		"id":   uint(668489),
		"form": uint(0),
		"type": "fighting",
	},
	"Kaiju Shelmet": {
		"id":   uint(668490),
		"form": uint(0),
		"type": "bug",
	},
	"Kaiju Skarmory": {
		"id":   uint(668491),
		"form": uint(0),
		"type": "steel",
	},
	"Kaiju Galvantula": {
		"id":   uint(668492),
		"form": uint(0),
		"type": "bug",
	},
	"Kaiju Maractus": {
		"id":   uint(668493),
		"form": uint(0),
		"type": "grass",
	},
	"Kaiju Nosepass": {
		"id":   uint(668494),
		"form": uint(0),
		"type": "rock",
	},
	"Kaiju Anorith": {
		"id":   uint(668495),
		"form": uint(0),
		"type": "rock",
	},
	"Kaiju Heliolisk": {
		"id":   uint(668496),
		"form": uint(0),
		"type": "electric",
	},
	"Kaiju Gastly": {
		"id":   uint(668497),
		"form": uint(0),
		"type": "ghost",
	},
	"Kaiju Shroomish": {
		"id":   uint(668498),
		"form": uint(0),
		"type": "grass",
	},
	"Bsentret": {
		"id":   uint(668499),
		"form": uint(0),
		"type": "normal",
	},
	"Bfurret": {
		"id":   uint(668500),
		"form": uint(0),
		"type": "normal",
	},
	"Btotodile": {
		"id":   uint(668501),
		"form": uint(0),
		"type": "water",
	},
	"Bforretress": {
		"id":   uint(668502),
		"form": uint(0),
		"type": "bug",
	},
	"Baipom": {
		"id":   uint(668503),
		"form": uint(0),
		"type": "normal",
	},
	"Qigglybuff": {
		"id":   uint(668504),
		"form": uint(0),
		"type": "qmarks",
	},
	"Bjumpluff": {
		"id":   uint(668505),
		"form": uint(0),
		"type": "grass",
	},
	"Bkingdratwo": {
		"id":   uint(668506),
		"form": uint(0),
		"type": "water",
	},
	"Boctillerytwo": {
		"id":   uint(668507),
		"form": uint(0),
		"type": "water",
	},
	"Bteddiursa": {
		"id":   uint(668508),
		"form": uint(0),
		"type": "normal",
	},
	"Bsneaseltwo": {
		"id":   uint(668509),
		"form": uint(0),
		"type": "ice",
	},
	"Bsneaselthree": {
		"id":   uint(668510),
		"form": uint(0),
		"type": "dark",
	},
	"Bphanpy": {
		"id":   uint(668511),
		"form": uint(0),
		"type": "ground",
	},
	"Bsunkern": {
		"id":   uint(668512),
		"form": uint(0),
		"type": "grass",
	},
	"Unfinkingdra": {
		"id":   uint(668513),
		"form": uint(0),
		"type": "qmarks",
	},
	"Bariados": {
		"id":   uint(668514),
		"form": uint(0),
		"type": "bug",
	},
	"Halfblissey": {
		"id":   uint(668515),
		"form": uint(0),
		"type": "normal",
	},
	"Poopgloom": {
		"id":   uint(668516),
		"form": uint(0),
		"type": "grass",
	},
	"Bcelebi": {
		"id":   uint(668517),
		"form": uint(0),
		"type": "psychic",
	},
	"Camainoe": {
		"id":   uint(668518),
		"form": uint(0),
		"type": "water",
	},
	"Beamcurry": {
		"id":   uint(668519),
		"form": uint(0),
		"type": "food",
	},
	"Beabcurry": {
		"id":   uint(668520),
		"form": uint(0),
		"type": "food",
	},
	"Burmcurry": {
		"id":   uint(668521),
		"form": uint(0),
		"type": "food",
	},
	"Burbcurry": {
		"id":   uint(668522),
		"form": uint(0),
		"type": "food",
	},
	"Tropmcurry": {
		"id":   uint(668523),
		"form": uint(0),
		"type": "food",
	},
	"Tropbcurry": {
		"id":   uint(668524),
		"form": uint(0),
		"type": "food",
	},
	"Gaotora": {
		"id":   uint(668525),
		"form": uint(0),
		"type": "electric",
	},
	"Arjunarcher": {
		"id":   uint(668526),
		"form": uint(0),
		"type": "divine",
	},
	"Lotamiesq": {
		"id":   uint(668527),
		"form": uint(0),
		"type": "grass",
	},
	"Titsdevoir": {
		"id":   uint(668528),
		"form": uint(0),
		"type": "dark",
	},
	"Caesaber": {
		"id":   uint(668529),
		"form": uint(0),
		"type": "fighting",
	},
	"Wdeino": {
		"id":   uint(668530),
		"form": uint(0),
		"type": "steel",
	},
	"Wzweilous": {
		"id":   uint(668531),
		"form": uint(0),
		"type": "steel",
	},
	"Whydreigon": {
		"id":   uint(668532),
		"form": uint(0),
		"type": "steel",
	},
	"Ebisurock": {
		"id":   uint(668533),
		"form": uint(0),
		"type": "ghost",
	},
	"Urumicow": {
		"id":   uint(668534),
		"form": uint(0),
		"type": "normal",
	},
	"Cocktaka": {
		"id":   uint(668535),
		"form": uint(0),
		"type": "divine",
	},
	"Yachieou": {
		"id":   uint(668536),
		"form": uint(0),
		"type": "dragon",
	},
	"Mayumior": {
		"id":   uint(668537),
		"form": uint(0),
		"type": "ground",
	},
	"Keikidol": {
		"id":   uint(668538),
		"form": uint(0),
		"type": "divine",
	},
	"Eternalarva": {
		"id":   uint(668539),
		"form": uint(0),
		"type": "fairy",
	},
	"Nemunag": {
		"id":   uint(668540),
		"form": uint(0),
		"type": "steel",
	},
	"Doggaunn": {
		"id":   uint(668541),
		"form": uint(0),
		"type": "rock",
	},
	"Narumistomp": {
		"id":   uint(668542),
		"form": uint(0),
		"type": "rock",
	},
	"Maidancer": {
		"id":   uint(668543),
		"form": uint(0),
		"type": "sound",
	},
	"Satanodancer": {
		"id":   uint(668544),
		"form": uint(0),
		"type": "sound",
	},
	"Okinaseason": {
		"id":   uint(668545),
		"form": uint(0),
		"type": "divine",
	},
	"Jooney": {
		"id":   uint(668546),
		"form": uint(0),
		"type": "fighting",
	},
	"Shionhobo": {
		"id":   uint(668547),
		"form": uint(0),
		"type": "ghost",
	},
	"Sagumeverse": {
		"id":   uint(668548),
		"form": uint(0),
		"type": "flying",
	},
	"Seirabbit": {
		"id":   uint(668549),
		"form": uint(0),
		"type": "normal",
	},
	"Hecatiaology": {
		"id":   uint(668550),
		"form": uint(0),
		"type": "divine",
	},
	"Kasenoni": {
		"id":   uint(668551),
		"form": uint(0),
		"type": "fighting",
	},
	"Praticlaw": {
		"id":   uint(668552),
		"form": uint(0),
		"type": "virus",
	},
	"Numemon X": {
		"id":   uint(668553),
		"form": uint(0),
		"type": "poison",
	},
	"Sky Regigigas": {
		"id":   uint(668554),
		"form": uint(0),
		"type": "flying",
	},
	"Bellscrop": {
		"id":   uint(668555),
		"form": uint(0),
		"type": "grass",
	},
	"Atlasbell": {
		"id":   uint(668556),
		"form": uint(0),
		"type": "grass",
	},
	"Hoppot": {
		"id":   uint(668557),
		"form": uint(0),
		"type": "grass",
	},
	"Hoprake": {
		"id":   uint(668558),
		"form": uint(0),
		"type": "grass",
	},
	"Potish": {
		"id":   uint(668559),
		"form": uint(0),
		"type": "grass",
	},
	"Trickish": {
		"id":   uint(668560),
		"form": uint(0),
		"type": "grass",
	},
	"Roseling": {
		"id":   uint(668561),
		"form": uint(0),
		"type": "grass",
	},
	"Witherade": {
		"id":   uint(668562),
		"form": uint(0),
		"type": "grass",
	},
	"Dmareep": {
		"id":   uint(668563),
		"form": uint(0),
		"type": "electric",
	},
	"Dflaaffy": {
		"id":   uint(668564),
		"form": uint(0),
		"type": "electric",
	},
	"Dampharos": {
		"id":   uint(668565),
		"form": uint(0),
		"type": "electric",
	},
	"Gswablu": {
		"id":   uint(668566),
		"form": uint(0),
		"type": "grass",
	},
	"Galtaria": {
		"id":   uint(668567),
		"form": uint(0),
		"type": "dragon",
	},
	"Brycenman": {
		"id":   uint(668568),
		"form": uint(0),
		"type": "dark",
	},
	"Brycencop": {
		"id":   uint(668569),
		"form": uint(0),
		"type": "tech",
	},
	"Brycenjet": {
		"id":   uint(668570),
		"form": uint(0),
		"type": "dark",
	},
	"Invadaifu": {
		"id":   uint(668571),
		"form": uint(0),
		"type": "cosmic",
	},
	"Pwt Golem": {
		"id":   uint(668572),
		"form": uint(0),
		"type": "wood",
	},
	"Gobogeast": {
		"id":   uint(668573),
		"form": uint(0),
		"type": "cosmic",
	},
	"Twinztwo": {
		"id":   uint(668574),
		"form": uint(0),
		"type": "normal",
	},
	"Mummeem": {
		"id":   uint(668575),
		"form": uint(0),
		"type": "zombie",
	},
	"Pharoammy": {
		"id":   uint(668576),
		"form": uint(0),
		"type": "zombie",
	},
	"Momummy": {
		"id":   uint(668577),
		"form": uint(0),
		"type": "zombie",
	},
	"Dootnoot": {
		"id":   uint(668578),
		"form": uint(0),
		"type": "sound",
	},
	"Frymcurry": {
		"id":   uint(668579),
		"form": uint(0),
		"type": "food",
	},
	"Frybcurry": {
		"id":   uint(668580),
		"form": uint(0),
		"type": "food",
	},
	"Crmmcurry": {
		"id":   uint(668581),
		"form": uint(0),
		"type": "food",
	},
	"Crmbcurry": {
		"id":   uint(668582),
		"form": uint(0),
		"type": "food",
	},
	"Rammcurry": {
		"id":   uint(668583),
		"form": uint(0),
		"type": "food",
	},
	"Rambcurry": {
		"id":   uint(668584),
		"form": uint(0),
		"type": "food",
	},
	"Seamcurry": {
		"id":   uint(668585),
		"form": uint(0),
		"type": "food",
	},
	"Seabcurry": {
		"id":   uint(668586),
		"form": uint(0),
		"type": "food",
	},
	"Salmcurry": {
		"id":   uint(668587),
		"form": uint(0),
		"type": "food",
	},
	"Salbcurry": {
		"id":   uint(668588),
		"form": uint(0),
		"type": "food",
	},
	"Slomcurry": {
		"id":   uint(668589),
		"form": uint(0),
		"type": "food",
	},
	"Slobcurry": {
		"id":   uint(668590),
		"form": uint(0),
		"type": "food",
	},
	"Bonmcurry": {
		"id":   uint(668591),
		"form": uint(0),
		"type": "food",
	},
	"Bonbcurry": {
		"id":   uint(668592),
		"form": uint(0),
		"type": "food",
	},
	"Eggmcurry": {
		"id":   uint(668593),
		"form": uint(0),
		"type": "food",
	},
	"Eggbcurry": {
		"id":   uint(668594),
		"form": uint(0),
		"type": "food",
	},
	"Potmcurry": {
		"id":   uint(668595),
		"form": uint(0),
		"type": "food",
	},
	"Potbcurry": {
		"id":   uint(668596),
		"form": uint(0),
		"type": "food",
	},
	"Herbmcurry": {
		"id":   uint(668597),
		"form": uint(0),
		"type": "food",
	},
	"Herbbcurry": {
		"id":   uint(668598),
		"form": uint(0),
		"type": "food",
	},
	"Leekmcurry": {
		"id":   uint(668599),
		"form": uint(0),
		"type": "food",
	},
	"Leekbcurry": {
		"id":   uint(668600),
		"form": uint(0),
		"type": "food",
	},
	"Cocomcurry": {
		"id":   uint(668601),
		"form": uint(0),
		"type": "food",
	},
	"Cocobcurry": {
		"id":   uint(668602),
		"form": uint(0),
		"type": "food",
	},
	"Breadmcurry": {
		"id":   uint(668603),
		"form": uint(0),
		"type": "food",
	},
	"Breadbcurry": {
		"id":   uint(668604),
		"form": uint(0),
		"type": "food",
	},
	"Pastamcurry": {
		"id":   uint(668605),
		"form": uint(0),
		"type": "food",
	},
	"Pastabcurry": {
		"id":   uint(668606),
		"form": uint(0),
		"type": "food",
	},
	"Cheesemcurry": {
		"id":   uint(668607),
		"form": uint(0),
		"type": "food",
	},
	"Cheesebcurry": {
		"id":   uint(668608),
		"form": uint(0),
		"type": "food",
	},
	"Juicymcurry": {
		"id":   uint(668609),
		"form": uint(0),
		"type": "food",
	},
	"Juicybcurry": {
		"id":   uint(668610),
		"form": uint(0),
		"type": "food",
	},
	"Richmcurry": {
		"id":   uint(668611),
		"form": uint(0),
		"type": "food",
	},
	"Richbcurry": {
		"id":   uint(668612),
		"form": uint(0),
		"type": "food",
	},
	"Roven": {
		"id":   uint(668613),
		"form": uint(0),
		"type": "rock",
	},
	"Infernace": {
		"id":   uint(668614),
		"form": uint(0),
		"type": "steel",
	},
	"Jekhyde": {
		"id":   uint(668615),
		"form": uint(0),
		"type": "normal",
	},
	"Jekhyde-Zen": {
		"id":   uint(668615),
		"form": uint(1),
		"type": "chaos",
	},
	"Hasshinji": {
		"id":   uint(668616),
		"form": uint(0),
		"type": "dark",
	},
	"Falsessan": {
		"id":   uint(668617),
		"form": uint(0),
		"type": "qmarks",
	},
	"Maryanncer": {
		"id":   uint(668618),
		"form": uint(0),
		"type": "rock",
	},
	"Cherbim": {
		"id":   uint(668619),
		"form": uint(0),
		"type": "grass",
	},
	"Jasonber": {
		"id":   uint(668620),
		"form": uint(0),
		"type": "normal",
	},
	"Hbagon": {
		"id":   uint(668621),
		"form": uint(0),
		"type": "dragon",
	},
	"Hshelgon": {
		"id":   uint(668622),
		"form": uint(0),
		"type": "dragon",
	},
	"Hsalamence": {
		"id":   uint(668623),
		"form": uint(0),
		"type": "dragon",
	},
	"Pyukuluku": {
		"id":   uint(668624),
		"form": uint(0),
		"type": "water",
	},
	"Sstonjourner": {
		"id":   uint(668625),
		"form": uint(0),
		"type": "rock",
	},
	"Dmeowth": {
		"id":   uint(668626),
		"form": uint(0),
		"type": "dragon",
	},
	"Dpersian": {
		"id":   uint(668627),
		"form": uint(0),
		"type": "dragon",
	},
	"Galpersian": {
		"id":   uint(668628),
		"form": uint(0),
		"type": "steel",
	},
	"Kperrserker": {
		"id":   uint(668629),
		"form": uint(0),
		"type": "normal",
	},
	"Aperrserker": {
		"id":   uint(668630),
		"form": uint(0),
		"type": "dark",
	},
	"Sufdonyut": {
		"id":   uint(668631),
		"form": uint(0),
		"type": "food",
	},
	"Huneehi": {
		"id":   uint(668632),
		"form": uint(0),
		"type": "bug",
	},
	"Huneyler": {
		"id":   uint(668633),
		"form": uint(0),
		"type": "bug",
	},
	"Jcursola": {
		"id":   uint(668634),
		"form": uint(0),
		"type": "water",
	},
	"Felekid": {
		"id":   uint(668635),
		"form": uint(0),
		"type": "fire",
	},
	"Felectabuzz": {
		"id":   uint(668636),
		"form": uint(0),
		"type": "fire",
	},
	"Felectivire": {
		"id":   uint(668637),
		"form": uint(0),
		"type": "fire",
	},
	"Emagby": {
		"id":   uint(668638),
		"form": uint(0),
		"type": "electric",
	},
	"Emagmar": {
		"id":   uint(668639),
		"form": uint(0),
		"type": "electric",
	},
	"Emagmortar": {
		"id":   uint(668640),
		"form": uint(0),
		"type": "electric",
	},
	"Ibouffalant": {
		"id":   uint(668641),
		"form": uint(0),
		"type": "normal",
	},
	"Bou": {
		"id":   uint(668642),
		"form": uint(0),
		"type": "virus",
	},
	"Boukka": {
		"id":   uint(668643),
		"form": uint(0),
		"type": "virus",
	},
	"Bouquetke": {
		"id":   uint(668644),
		"form": uint(0),
		"type": "virus",
	},
	"Clair": {
		"id":   uint(668645),
		"form": uint(0),
		"type": "dragon",
	},
	"Venusauri": {
		"id":   uint(668646),
		"form": uint(0),
		"type": "grass",
	},
	"Midger": {
		"id":   uint(668647),
		"form": uint(0),
		"type": "steel",
	},
	"Fingle": {
		"id":   uint(668648),
		"form": uint(0),
		"type": "steel",
	},
	"Faakyu": {
		"id":   uint(668649),
		"form": uint(0),
		"type": "steel",
	},
	"Blobby": {
		"id":   uint(668650),
		"form": uint(0),
		"type": "water",
	},
	"Blobuff": {
		"id":   uint(668651),
		"form": uint(0),
		"type": "water",
	},
	"Shaunaph": {
		"id":   uint(668652),
		"form": uint(0),
		"type": "steam",
	},
	"Himfear": {
		"id":   uint(668653),
		"form": uint(0),
		"type": "fear",
	},
	"Coalgate": {
		"id":   uint(668654),
		"form": uint(0),
		"type": "rock",
	},
	"Yanenvy": {
		"id":   uint(668655),
		"form": uint(0),
		"type": "fighting",
	},
	"Dshinx": {
		"id":   uint(668656),
		"form": uint(0),
		"type": "electric",
	},
	"Liberinx": {
		"id":   uint(668657),
		"form": uint(0),
		"type": "electric",
	},
	"Pslurpuff": {
		"id":   uint(668658),
		"form": uint(0),
		"type": "fairy",
	},
	"Urunerigus": {
		"id":   uint(668659),
		"form": uint(0),
		"type": "ghost",
	},
	"Dmunna": {
		"id":   uint(668660),
		"form": uint(0),
		"type": "dark",
	},
	"Dmusharna": {
		"id":   uint(668661),
		"form": uint(0),
		"type": "dark",
	},
	"Muspecta": {
		"id":   uint(668662),
		"form": uint(0),
		"type": "dark",
	},
	"Tdarumaka": {
		"id":   uint(668663),
		"form": uint(0),
		"type": "grass",
	},
	"Tdarmanitan": {
		"id":   uint(668664),
		"form": uint(0),
		"type": "grass",
	},
	"Tdarmanitan-Zen": {
		"id":   uint(668664),
		"form": uint(1),
		"type": "grass",
	},
	"Fagyamask": {
		"id":   uint(668665),
		"form": uint(0),
		"type": "fairy",
	},
	"Yangelic": {
		"id":   uint(668666),
		"form": uint(0),
		"type": "fairy",
	},
	"Pmrmime": {
		"id":   uint(668667),
		"form": uint(0),
		"type": "dark",
	},
	"Dvenonat": {
		"id":   uint(668668),
		"form": uint(0),
		"type": "bug",
	},
	"Manomoth": {
		"id":   uint(668669),
		"form": uint(0),
		"type": "bug",
	},
	"Rsmeargle": {
		"id":   uint(668670),
		"form": uint(0),
		"type": "normal",
	},
	"Psqwilfish": {
		"id":   uint(668671),
		"form": uint(0),
		"type": "psychic",
	},
	"Gtropius": {
		"id":   uint(668672),
		"form": uint(0),
		"type": "grass",
	},
	"Draerodactyl": {
		"id":   uint(668673),
		"form": uint(0),
		"type": "dragon",
	},
	"Saoranguru": {
		"id":   uint(668674),
		"form": uint(0),
		"type": "normal",
	},
	"Fsmoochum": {
		"id":   uint(668675),
		"form": uint(0),
		"type": "grass",
	},
	"Fjynx": {
		"id":   uint(668676),
		"form": uint(0),
		"type": "grass",
	},
	"Jybaba": {
		"id":   uint(668677),
		"form": uint(0),
		"type": "grass",
	},
	"Hubulbasaur": {
		"id":   uint(668678),
		"form": uint(0),
		"type": "grass",
	},
	"Huivysaur": {
		"id":   uint(668679),
		"form": uint(0),
		"type": "grass",
	},
	"Huvenusaur": {
		"id":   uint(668680),
		"form": uint(0),
		"type": "grass",
	},
	"Hucharmander": {
		"id":   uint(668681),
		"form": uint(0),
		"type": "fire",
	},
	"Hucharmeleon": {
		"id":   uint(668682),
		"form": uint(0),
		"type": "fire",
	},
	"Hucharizard": {
		"id":   uint(668683),
		"form": uint(0),
		"type": "fire",
	},
	"Husquirtle": {
		"id":   uint(668684),
		"form": uint(0),
		"type": "water",
	},
	"Huwartortle": {
		"id":   uint(668685),
		"form": uint(0),
		"type": "water",
	},
	"Hublastoise": {
		"id":   uint(668686),
		"form": uint(0),
		"type": "water",
	},
	"Hucaterpie": {
		"id":   uint(668687),
		"form": uint(0),
		"type": "bug",
	},
	"Humetapod": {
		"id":   uint(668688),
		"form": uint(0),
		"type": "bug",
	},
	"Hubutterfree": {
		"id":   uint(668689),
		"form": uint(0),
		"type": "bug",
	},
	"Huweedle": {
		"id":   uint(668690),
		"form": uint(0),
		"type": "bug",
	},
	"Hukakuna": {
		"id":   uint(668691),
		"form": uint(0),
		"type": "bug",
	},
	"Hubeedrill": {
		"id":   uint(668692),
		"form": uint(0),
		"type": "bug",
	},
	"Hupidgey": {
		"id":   uint(668693),
		"form": uint(0),
		"type": "normal",
	},
	"Hupidgeotto": {
		"id":   uint(668694),
		"form": uint(0),
		"type": "normal",
	},
	"Hupidgeot": {
		"id":   uint(668695),
		"form": uint(0),
		"type": "normal",
	},
	"Hurattata": {
		"id":   uint(668696),
		"form": uint(0),
		"type": "normal",
	},
	"Huraticate": {
		"id":   uint(668697),
		"form": uint(0),
		"type": "normal",
	},
	"Huspearow": {
		"id":   uint(668698),
		"form": uint(0),
		"type": "normal",
	},
	"Hufearow": {
		"id":   uint(668699),
		"form": uint(0),
		"type": "normal",
	},
	"Huekans": {
		"id":   uint(668700),
		"form": uint(0),
		"type": "poison",
	},
	"Huarbok": {
		"id":   uint(668701),
		"form": uint(0),
		"type": "poison",
	},
	"Hupikachu": {
		"id":   uint(668702),
		"form": uint(0),
		"type": "electric",
	},
	"Huraichu": {
		"id":   uint(668703),
		"form": uint(0),
		"type": "electric",
	},
	"Husandshrew": {
		"id":   uint(668704),
		"form": uint(0),
		"type": "ground",
	},
	"Husandslash": {
		"id":   uint(668705),
		"form": uint(0),
		"type": "ground",
	},
	"Hunidoranfe": {
		"id":   uint(668706),
		"form": uint(0),
		"type": "poison",
	},
	"Hunidorina": {
		"id":   uint(668707),
		"form": uint(0),
		"type": "poison",
	},
	"Hunidoqueen": {
		"id":   uint(668708),
		"form": uint(0),
		"type": "poison",
	},
	"Hunidoranma": {
		"id":   uint(668709),
		"form": uint(0),
		"type": "poison",
	},
	"Hunidorino": {
		"id":   uint(668710),
		"form": uint(0),
		"type": "poison",
	},
	"Hunidoking": {
		"id":   uint(668711),
		"form": uint(0),
		"type": "poison",
	},
	"Huclefairy": {
		"id":   uint(668712),
		"form": uint(0),
		"type": "fairy",
	},
	"Huclefable": {
		"id":   uint(668713),
		"form": uint(0),
		"type": "fairy",
	},
	"Huvulpix": {
		"id":   uint(668714),
		"form": uint(0),
		"type": "fire",
	},
	"Huninetales": {
		"id":   uint(668715),
		"form": uint(0),
		"type": "fire",
	},
	"Hujigglypuff": {
		"id":   uint(668716),
		"form": uint(0),
		"type": "sound",
	},
	"Huwigglytuff": {
		"id":   uint(668717),
		"form": uint(0),
		"type": "sound",
	},
	"Huzubat": {
		"id":   uint(668718),
		"form": uint(0),
		"type": "poison",
	},
	"Hugolbat": {
		"id":   uint(668719),
		"form": uint(0),
		"type": "poison",
	},
	"Huoddish": {
		"id":   uint(668720),
		"form": uint(0),
		"type": "grass",
	},
	"Hugloom": {
		"id":   uint(668721),
		"form": uint(0),
		"type": "grass",
	},
	"Huvileplume": {
		"id":   uint(668722),
		"form": uint(0),
		"type": "grass",
	},
	"Huparas": {
		"id":   uint(668723),
		"form": uint(0),
		"type": "bug",
	},
	"Huparasect": {
		"id":   uint(668724),
		"form": uint(0),
		"type": "bug",
	},
	"Huvenonat": {
		"id":   uint(668725),
		"form": uint(0),
		"type": "bug",
	},
	"Huvenomoth": {
		"id":   uint(668726),
		"form": uint(0),
		"type": "bug",
	},
	"Hudiglett": {
		"id":   uint(668727),
		"form": uint(0),
		"type": "ground",
	},
	"Hudugtrio": {
		"id":   uint(668728),
		"form": uint(0),
		"type": "ground",
	},
	"Humeowth": {
		"id":   uint(668729),
		"form": uint(0),
		"type": "normal",
	},
	"Hupersian": {
		"id":   uint(668730),
		"form": uint(0),
		"type": "normal",
	},
	"Hupsyduck": {
		"id":   uint(668731),
		"form": uint(0),
		"type": "water",
	},
	"Hugolduck": {
		"id":   uint(668732),
		"form": uint(0),
		"type": "water",
	},
	"Humankey": {
		"id":   uint(668733),
		"form": uint(0),
		"type": "fighting",
	},
	"Huprimeape": {
		"id":   uint(668734),
		"form": uint(0),
		"type": "fighting",
	},
	"Hugrowlithe": {
		"id":   uint(668735),
		"form": uint(0),
		"type": "fire",
	},
	"Huarcanine": {
		"id":   uint(668736),
		"form": uint(0),
		"type": "fire",
	},
	"Hupoliwag": {
		"id":   uint(668737),
		"form": uint(0),
		"type": "water",
	},
	"Hupoliwhirl": {
		"id":   uint(668738),
		"form": uint(0),
		"type": "water",
	},
	"Hupoliwrath": {
		"id":   uint(668739),
		"form": uint(0),
		"type": "water",
	},
	"Huabra": {
		"id":   uint(668740),
		"form": uint(0),
		"type": "psychic",
	},
	"Hukadabra": {
		"id":   uint(668741),
		"form": uint(0),
		"type": "psychic",
	},
	"Hualakazam": {
		"id":   uint(668742),
		"form": uint(0),
		"type": "psychic",
	},
	"Humachop": {
		"id":   uint(668743),
		"form": uint(0),
		"type": "fighting",
	},
	"Humachoke": {
		"id":   uint(668744),
		"form": uint(0),
		"type": "fighting",
	},
	"Humachamp": {
		"id":   uint(668745),
		"form": uint(0),
		"type": "fighting",
	},
	"Hubellsprout": {
		"id":   uint(668746),
		"form": uint(0),
		"type": "grass",
	},
	"Huweepinbell": {
		"id":   uint(668747),
		"form": uint(0),
		"type": "grass",
	},
	"Huvictreebel": {
		"id":   uint(668748),
		"form": uint(0),
		"type": "grass",
	},
	"Hutentacool": {
		"id":   uint(668749),
		"form": uint(0),
		"type": "water",
	},
	"Hutentacruel": {
		"id":   uint(668750),
		"form": uint(0),
		"type": "water",
	},
	"Hugeodude": {
		"id":   uint(668751),
		"form": uint(0),
		"type": "rock",
	},
	"Hugraveler": {
		"id":   uint(668752),
		"form": uint(0),
		"type": "rock",
	},
	"Hugolem": {
		"id":   uint(668753),
		"form": uint(0),
		"type": "rock",
	},
	"Huponyta": {
		"id":   uint(668754),
		"form": uint(0),
		"type": "fire",
	},
	"Hurapidash": {
		"id":   uint(668755),
		"form": uint(0),
		"type": "fire",
	},
	"Huslowpoke": {
		"id":   uint(668756),
		"form": uint(0),
		"type": "water",
	},
	"Huslowbro": {
		"id":   uint(668757),
		"form": uint(0),
		"type": "water",
	},
	"Humagnemite": {
		"id":   uint(668758),
		"form": uint(0),
		"type": "electric",
	},
	"Humagneton": {
		"id":   uint(668759),
		"form": uint(0),
		"type": "electric",
	},
	"Hufarfetchd": {
		"id":   uint(668760),
		"form": uint(0),
		"type": "normal",
	},
	"Hudoduo": {
		"id":   uint(668761),
		"form": uint(0),
		"type": "normal",
	},
	"Hudodrio": {
		"id":   uint(668762),
		"form": uint(0),
		"type": "normal",
	},
	"Huseel": {
		"id":   uint(668763),
		"form": uint(0),
		"type": "water",
	},
	"Hudewgong": {
		"id":   uint(668764),
		"form": uint(0),
		"type": "water",
	},
	"Hugrimer": {
		"id":   uint(668765),
		"form": uint(0),
		"type": "poison",
	},
	"Humuk": {
		"id":   uint(668766),
		"form": uint(0),
		"type": "poison",
	},
	"Hushellder": {
		"id":   uint(668767),
		"form": uint(0),
		"type": "water",
	},
	"Hucloyster": {
		"id":   uint(668768),
		"form": uint(0),
		"type": "water",
	},
	"Hugastly": {
		"id":   uint(668769),
		"form": uint(0),
		"type": "ghost",
	},
	"Huhaunter": {
		"id":   uint(668770),
		"form": uint(0),
		"type": "ghost",
	},
	"Hugengar": {
		"id":   uint(668771),
		"form": uint(0),
		"type": "ghost",
	},
	"Huonix": {
		"id":   uint(668772),
		"form": uint(0),
		"type": "rock",
	},
	"Hudrowzee": {
		"id":   uint(668773),
		"form": uint(0),
		"type": "psychic",
	},
	"Huhypno": {
		"id":   uint(668774),
		"form": uint(0),
		"type": "psychic",
	},
	"Hukrabby": {
		"id":   uint(668775),
		"form": uint(0),
		"type": "water",
	},
	"Hukingler": {
		"id":   uint(668776),
		"form": uint(0),
		"type": "water",
	},
	"Huvoltorb": {
		"id":   uint(668777),
		"form": uint(0),
		"type": "electric",
	},
	"Huelectrode": {
		"id":   uint(668778),
		"form": uint(0),
		"type": "electric",
	},
	"Huexeggcute": {
		"id":   uint(668779),
		"form": uint(0),
		"type": "grass",
	},
	"Huexeggutor": {
		"id":   uint(668780),
		"form": uint(0),
		"type": "grass",
	},
	"Hucubone": {
		"id":   uint(668781),
		"form": uint(0),
		"type": "bone",
	},
	"Humarowak": {
		"id":   uint(668782),
		"form": uint(0),
		"type": "bone",
	},
	"Huhitmonlee": {
		"id":   uint(668783),
		"form": uint(0),
		"type": "fighting",
	},
	"Huhitmonchan": {
		"id":   uint(668784),
		"form": uint(0),
		"type": "fighting",
	},
	"Hulickitung": {
		"id":   uint(668785),
		"form": uint(0),
		"type": "normal",
	},
	"Hukoffing": {
		"id":   uint(668786),
		"form": uint(0),
		"type": "poison",
	},
	"Huweezing": {
		"id":   uint(668787),
		"form": uint(0),
		"type": "poison",
	},
	"Hurhyhorn": {
		"id":   uint(668788),
		"form": uint(0),
		"type": "bone",
	},
	"Hurhydon": {
		"id":   uint(668789),
		"form": uint(0),
		"type": "bone",
	},
	"Huchansey": {
		"id":   uint(668790),
		"form": uint(0),
		"type": "normal",
	},
	"Hutangela": {
		"id":   uint(668791),
		"form": uint(0),
		"type": "grass",
	},
	"Hukangaskhan": {
		"id":   uint(668792),
		"form": uint(0),
		"type": "normal",
	},
	"Huhorsea": {
		"id":   uint(668793),
		"form": uint(0),
		"type": "water",
	},
	"Huseadra": {
		"id":   uint(668794),
		"form": uint(0),
		"type": "water",
	},
	"Hugoldeen": {
		"id":   uint(668795),
		"form": uint(0),
		"type": "water",
	},
	"Huseaking": {
		"id":   uint(668796),
		"form": uint(0),
		"type": "water",
	},
	"Hustaryu": {
		"id":   uint(668797),
		"form": uint(0),
		"type": "water",
	},
	"Hustarmie": {
		"id":   uint(668798),
		"form": uint(0),
		"type": "water",
	},
	"Humrmime": {
		"id":   uint(668799),
		"form": uint(0),
		"type": "psychic",
	},
	"Huscyther": {
		"id":   uint(668800),
		"form": uint(0),
		"type": "bug",
	},
	"Hujynx": {
		"id":   uint(668801),
		"form": uint(0),
		"type": "ice",
	},
	"Huelectabuzz": {
		"id":   uint(668802),
		"form": uint(0),
		"type": "electric",
	},
	"Humagmar": {
		"id":   uint(668803),
		"form": uint(0),
		"type": "fire",
	},
	"Hupinsir": {
		"id":   uint(668804),
		"form": uint(0),
		"type": "bug",
	},
	"Hutauros": {
		"id":   uint(668805),
		"form": uint(0),
		"type": "normal",
	},
	"Humagikarp": {
		"id":   uint(668806),
		"form": uint(0),
		"type": "water",
	},
	"Hugyarados": {
		"id":   uint(668807),
		"form": uint(0),
		"type": "water",
	},
	"Hulapras": {
		"id":   uint(668808),
		"form": uint(0),
		"type": "water",
	},
	"Huditto": {
		"id":   uint(668809),
		"form": uint(0),
		"type": "normal",
	},
	"Hueevee": {
		"id":   uint(668810),
		"form": uint(0),
		"type": "normal",
	},
	"Huvaporeon": {
		"id":   uint(668811),
		"form": uint(0),
		"type": "water",
	},
	"Hujolteon": {
		"id":   uint(668812),
		"form": uint(0),
		"type": "electric",
	},
	"Huflareon": {
		"id":   uint(668813),
		"form": uint(0),
		"type": "fire",
	},
	"Huporygon": {
		"id":   uint(668814),
		"form": uint(0),
		"type": "normal",
	},
	"Huomanyte": {
		"id":   uint(668815),
		"form": uint(0),
		"type": "rock",
	},
	"Huomastar": {
		"id":   uint(668816),
		"form": uint(0),
		"type": "rock",
	},
	"Hukabuto": {
		"id":   uint(668817),
		"form": uint(0),
		"type": "rock",
	},
	"Hukabutops": {
		"id":   uint(668818),
		"form": uint(0),
		"type": "rock",
	},
	"Huaerodactyl": {
		"id":   uint(668819),
		"form": uint(0),
		"type": "rock",
	},
	"Husnorlax": {
		"id":   uint(668820),
		"form": uint(0),
		"type": "normal",
	},
	"Huarticuno": {
		"id":   uint(668821),
		"form": uint(0),
		"type": "ice",
	},
	"Huzapdos": {
		"id":   uint(668822),
		"form": uint(0),
		"type": "electric",
	},
	"Humoltres": {
		"id":   uint(668823),
		"form": uint(0),
		"type": "fire",
	},
	"Hudratini": {
		"id":   uint(668824),
		"form": uint(0),
		"type": "dragon",
	},
	"Hudragonair": {
		"id":   uint(668825),
		"form": uint(0),
		"type": "dragon",
	},
	"Hudragonite": {
		"id":   uint(668826),
		"form": uint(0),
		"type": "dragon",
	},
	"Humewtwo": {
		"id":   uint(668827),
		"form": uint(0),
		"type": "psychic",
	},
	"Humew": {
		"id":   uint(668828),
		"form": uint(0),
		"type": "psychic",
	},
	"Drgnchamber": {
		"id":   uint(668829),
		"form": uint(0),
		"type": "dragon",
	},
	"Zbagon": {
		"id":   uint(668830),
		"form": uint(0),
		"type": "dragon",
	},
	"Zshelgon": {
		"id":   uint(668831),
		"form": uint(0),
		"type": "dragon",
	},
	"Zsalamence": {
		"id":   uint(668832),
		"form": uint(0),
		"type": "dragon",
	},
	"Snkgrimreaper": {
		"id":   uint(668833),
		"form": uint(0),
		"type": "dark",
	},
	"Ironreaper": {
		"id":   uint(668834),
		"form": uint(0),
		"type": "zombie",
	},
	"Houndsour": {
		"id":   uint(668835),
		"form": uint(0),
		"type": "dark",
	},
	"Headdoom": {
		"id":   uint(668836),
		"form": uint(0),
		"type": "dark",
	},
	"Skinskin": {
		"id":   uint(668837),
		"form": uint(0),
		"type": "ice",
	},
	"Snkeyeeye": {
		"id":   uint(668838),
		"form": uint(0),
		"type": "ice",
	},
	"Elecharge": {
		"id":   uint(668839),
		"form": uint(0),
		"type": "electric",
	},
	"Electabugs": {
		"id":   uint(668840),
		"form": uint(0),
		"type": "electric",
	},
	"Miaby": {
		"id":   uint(668841),
		"form": uint(0),
		"type": "fire",
	},
	"Miasmar": {
		"id":   uint(668842),
		"form": uint(0),
		"type": "fire",
	},
	"Irivatar": {
		"id":   uint(668843),
		"form": uint(0),
		"type": "steel",
	},
	"Metalitar": {
		"id":   uint(668844),
		"form": uint(0),
		"type": "steel",
	},
	"Zyranizila": {
		"id":   uint(668845),
		"form": uint(0),
		"type": "steel",
	},
	"Scarlugia": {
		"id":   uint(668846),
		"form": uint(0),
		"type": "psychic",
	},
	"Vigourlan": {
		"id":   uint(668847),
		"form": uint(0),
		"type": "normal",
	},
	"Rotting": {
		"id":   uint(668848),
		"form": uint(0),
		"type": "blood",
	},
	"Snowkitty": {
		"id":   uint(668849),
		"form": uint(0),
		"type": "ice",
	},
	"Frozcatty": {
		"id":   uint(668850),
		"form": uint(0),
		"type": "ice",
	},
	"Zangol": {
		"id":   uint(668851),
		"form": uint(0),
		"type": "normal",
	},
	"Sevicious": {
		"id":   uint(668852),
		"form": uint(0),
		"type": "poison",
	},
	"Galactimon": {
		"id":   uint(668853),
		"form": uint(0),
		"type": "rock",
	},
	"Yeutdoom": {
		"id":   uint(668854),
		"form": uint(0),
		"type": "dark",
	},
	"Wobdoomer": {
		"id":   uint(668855),
		"form": uint(0),
		"type": "dark",
	},
	"Evidisc": {
		"id":   uint(668856),
		"form": uint(0),
		"type": "dark",
	},
	"Demohate": {
		"id":   uint(668857),
		"form": uint(0),
		"type": "dark",
	},
	"Rustum": {
		"id":   uint(668858),
		"form": uint(0),
		"type": "steel",
	},
	"Rustang": {
		"id":   uint(668859),
		"form": uint(0),
		"type": "steel",
	},
	"Rustegross": {
		"id":   uint(668860),
		"form": uint(0),
		"type": "steel",
	},
	"Stitcher": {
		"id":   uint(668861),
		"form": uint(0),
		"type": "dragon",
	},
	"Slimex": {
		"id":   uint(668862),
		"form": uint(0),
		"type": "fire",
	},
	"Blackuto": {
		"id":   uint(668863),
		"form": uint(0),
		"type": "zombie",
	},
	"Darkutops": {
		"id":   uint(668864),
		"form": uint(0),
		"type": "zombie",
	},
	"Gorelax": {
		"id":   uint(668865),
		"form": uint(0),
		"type": "normal",
	},
	"Mutagon": {
		"id":   uint(668866),
		"form": uint(0),
		"type": "psychic",
	},
	"Mewby": {
		"id":   uint(668867),
		"form": uint(0),
		"type": "psychic",
	},
	"Thorita": {
		"id":   uint(668868),
		"form": uint(0),
		"type": "poison",
	},
	"Thornleef": {
		"id":   uint(668869),
		"form": uint(0),
		"type": "poison",
	},
	"Thornium": {
		"id":   uint(668870),
		"form": uint(0),
		"type": "poison",
	},
	"Mawnium": {
		"id":   uint(668871),
		"form": uint(0),
		"type": "poison",
	},
	"Cleffgar": {
		"id":   uint(668872),
		"form": uint(0),
		"type": "rock",
	},
	"Clegoyle": {
		"id":   uint(668873),
		"form": uint(0),
		"type": "rock",
	},
	"Garfable": {
		"id":   uint(668874),
		"form": uint(0),
		"type": "rock",
	},
	"Snkzombified": {
		"id":   uint(668875),
		"form": uint(0),
		"type": "water",
	},
	"Snktoxeon": {
		"id":   uint(668876),
		"form": uint(0),
		"type": "nuclear",
	},
	"Snkdemoneon": {
		"id":   uint(668877),
		"form": uint(0),
		"type": "chaos",
	},
	"Dirtkrow": {
		"id":   uint(668878),
		"form": uint(0),
		"type": "magic",
	},
	"Dunschum": {
		"id":   uint(668879),
		"form": uint(0),
		"type": "normal",
	},
	"Qwilshark": {
		"id":   uint(668880),
		"form": uint(0),
		"type": "water",
	},
	"Telson": {
		"id":   uint(668881),
		"form": uint(0),
		"type": "water",
	},
	"Telshark": {
		"id":   uint(668882),
		"form": uint(0),
		"type": "water",
	},
	"Sevydra": {
		"id":   uint(668883),
		"form": uint(0),
		"type": "poison",
	},
	"Polihag": {
		"id":   uint(668884),
		"form": uint(0),
		"type": "water",
	},
	"Poliworm": {
		"id":   uint(668885),
		"form": uint(0),
		"type": "water",
	},
	"Poliwraith": {
		"id":   uint(668886),
		"form": uint(0),
		"type": "zombie",
	},
	"Moulder": {
		"id":   uint(668887),
		"form": uint(0),
		"type": "poison",
	},
	"Yuck": {
		"id":   uint(668888),
		"form": uint(0),
		"type": "poison",
	},
	"Roteen": {
		"id":   uint(668889),
		"form": uint(0),
		"type": "water",
	},
	"Seatyrant": {
		"id":   uint(668890),
		"form": uint(0),
		"type": "water",
	},
	"Zombicater": {
		"id":   uint(668891),
		"form": uint(0),
		"type": "zombie",
	},
	"Deaderant": {
		"id":   uint(668892),
		"form": uint(0),
		"type": "zombie",
	},
	"Brimshrew": {
		"id":   uint(668893),
		"form": uint(0),
		"type": "ground",
	},
	"Toxislash": {
		"id":   uint(668894),
		"form": uint(0),
		"type": "fire",
	},
	"Nidoblood": {
		"id":   uint(668895),
		"form": uint(0),
		"type": "poison",
	},
	"Rubirina": {
		"id":   uint(668896),
		"form": uint(0),
		"type": "poison",
	},
	"Rosaqueen": {
		"id":   uint(668897),
		"form": uint(0),
		"type": "poison",
	},
	"Metalran": {
		"id":   uint(668898),
		"form": uint(0),
		"type": "poison",
	},
	"Cyborino": {
		"id":   uint(668899),
		"form": uint(0),
		"type": "poison",
	},
	"Mechaking": {
		"id":   uint(668900),
		"form": uint(0),
		"type": "poison",
	},
	"Gluemadio": {
		"id":   uint(668901),
		"form": uint(0),
		"type": "ghost",
	},
	"Earthmadio": {
		"id":   uint(668902),
		"form": uint(0),
		"type": "ghost",
	},
	"Bonemadio": {
		"id":   uint(668903),
		"form": uint(0),
		"type": "bone",
	},
	"Crazymadio": {
		"id":   uint(668904),
		"form": uint(0),
		"type": "ghost",
	},
	"Seamadio": {
		"id":   uint(668905),
		"form": uint(0),
		"type": "ghost",
	},
	"Burstmadio": {
		"id":   uint(668906),
		"form": uint(0),
		"type": "ghost",
	},
	"Blastmadio": {
		"id":   uint(668907),
		"form": uint(0),
		"type": "ghost",
	},
	"Treemadio": {
		"id":   uint(668908),
		"form": uint(0),
		"type": "ghost",
	},
	"Aeromadio": {
		"id":   uint(668909),
		"form": uint(0),
		"type": "ghost",
	},
	"Normadio": {
		"id":   uint(668910),
		"form": uint(0),
		"type": "ghost",
	},
	"WackKingmadio": {
		"id":   uint(668911),
		"form": uint(0),
		"type": "ghost",
	},
	"Calfby": {
		"id":   uint(668912),
		"form": uint(0),
		"type": "normal",
	},
	"Ramshaker": {
		"id":   uint(668913),
		"form": uint(0),
		"type": "normal",
	},
	"A Monster": {
		"id":   uint(668914),
		"form": uint(0),
		"type": "zombie",
	},
	"Snkantispiral": {
		"id":   uint(668915),
		"form": uint(0),
		"type": "normal",
	},
	"Abyssparce": {
		"id":   uint(668916),
		"form": uint(0),
		"type": "ghost",
	},
	"Snkdragoone": {
		"id":   uint(668917),
		"form": uint(0),
		"type": "normal",
	},
	"Roclobster": {
		"id":   uint(668918),
		"form": uint(0),
		"type": "rock",
	},
	"Vala": {
		"id":   uint(668919),
		"form": uint(0),
		"type": "normal",
	},
	"Snkkamina": {
		"id":   uint(668920),
		"form": uint(0),
		"type": "fighting",
	},
	"Diamandix": {
		"id":   uint(668921),
		"form": uint(0),
		"type": "steel",
	},
	"Cinderco": {
		"id":   uint(668922),
		"form": uint(0),
		"type": "fire",
	},
	"Radiorange": {
		"id":   uint(668923),
		"form": uint(0),
		"type": "nuclear",
	},
	"Chocwork": {
		"id":   uint(668924),
		"form": uint(0),
		"type": "tech",
	},
	"Mysteryegg": {
		"id":   uint(668925),
		"form": uint(0),
		"type": "normal",
	},
	"Glitchlett": {
		"id":   uint(668926),
		"form": uint(0),
		"type": "ground",
	},
	"Glitchtrio": {
		"id":   uint(668927),
		"form": uint(0),
		"type": "ground",
	},
	"Glitchtet": {
		"id":   uint(668928),
		"form": uint(0),
		"type": "ground",
	},
	"Secretegg": {
		"id":   uint(668929),
		"form": uint(0),
		"type": "normal",
	},
	"Hyperegg": {
		"id":   uint(668930),
		"form": uint(0),
		"type": "normal",
	},
	"Kajilianth": {
		"id":   uint(668931),
		"form": uint(0),
		"type": "poison",
	},
	"Skylax": {
		"id":   uint(668932),
		"form": uint(0),
		"type": "normal",
	},
	"Hombone": {
		"id":   uint(668933),
		"form": uint(0),
		"type": "rock",
	},
	"Hombeast": {
		"id":   uint(668934),
		"form": uint(0),
		"type": "rock",
	},
	"Kenchira": {
		"id":   uint(668935),
		"form": uint(0),
		"type": "normal",
	},
	"Kenchukuo": {
		"id":   uint(668936),
		"form": uint(0),
		"type": "fighting",
	},
	"X32763": {
		"id":   uint(668937),
		"form": uint(0),
		"type": "steel",
	},
	"Furi X": {
		"id":   uint(668938),
		"form": uint(0),
		"type": "water",
	},
	"Furi Z": {
		"id":   uint(668939),
		"form": uint(0),
		"type": "fire",
	},
	"Furi Q": {
		"id":   uint(668940),
		"form": uint(0),
		"type": "grass",
	},
	"Turmur": {
		"id":   uint(668941),
		"form": uint(0),
		"type": "virus",
	},
	"Senex": {
		"id":   uint(668942),
		"form": uint(0),
		"type": "fighting",
	},
	"Luca Zamon": {
		"id":   uint(668943),
		"form": uint(0),
		"type": "dragon",
	},
	"Tea Barqan": {
		"id":   uint(668944),
		"form": uint(0),
		"type": "dragon",
	},
	"Shaderu": {
		"id":   uint(668945),
		"form": uint(0),
		"type": "ghost",
	},
	"Faceleech": {
		"id":   uint(668946),
		"form": uint(0),
		"type": "qmarks",
	},
	"Boilbasaur": {
		"id":   uint(668947),
		"form": uint(0),
		"type": "grass",
	},
	"Shrivlsaur": {
		"id":   uint(668948),
		"form": uint(0),
		"type": "grass",
	},
	"Vivosaur": {
		"id":   uint(668949),
		"form": uint(0),
		"type": "grass",
	},
	"Rotmander": {
		"id":   uint(668950),
		"form": uint(0),
		"type": "fire",
	},
	"Charmeworm": {
		"id":   uint(668951),
		"form": uint(0),
		"type": "fire",
	},
	"Snkdragon": {
		"id":   uint(668952),
		"form": uint(0),
		"type": "dragon",
	},
	"Lichizard": {
		"id":   uint(668953),
		"form": uint(0),
		"type": "fire",
	},
	"Oozle": {
		"id":   uint(668954),
		"form": uint(0),
		"type": "water",
	},
	"Entrailtle": {
		"id":   uint(668955),
		"form": uint(0),
		"type": "water",
	},
	"Bleedtoise": {
		"id":   uint(668956),
		"form": uint(0),
		"type": "water",
	},
	"Decarill": {
		"id":   uint(668957),
		"form": uint(0),
		"type": "zombie",
	},
	"Graveill": {
		"id":   uint(668958),
		"form": uint(0),
		"type": "zombie",
	},
	"Azombarill": {
		"id":   uint(668959),
		"form": uint(0),
		"type": "zombie",
	},
	"Psylet": {
		"id":   uint(668960),
		"form": uint(0),
		"type": "fire",
	},
	"Psypig": {
		"id":   uint(668961),
		"form": uint(0),
		"type": "fire",
	},
	"Zurmex": {
		"id":   uint(668962),
		"form": uint(0),
		"type": "bug",
	},
	"Sxsalamence": {
		"id":   uint(668963),
		"form": uint(0),
		"type": "dragon",
	},
	"Seimenkongou": {
		"id":   uint(668964),
		"form": uint(0),
		"type": "light",
	},
	"Metalslimee": {
		"id":   uint(668965),
		"form": uint(0),
		"type": "steel",
	},
	"Metalkingslimee": {
		"id":   uint(668966),
		"form": uint(0),
		"type": "steel",
	},
	"Burnasaur": {
		"id":   uint(668967),
		"form": uint(0),
		"type": "fire",
	},
	"Embersaur": {
		"id":   uint(668968),
		"form": uint(0),
		"type": "fire",
	},
	"Vulcasaur": {
		"id":   uint(668969),
		"form": uint(0),
		"type": "fire",
	},
	"Seamander": {
		"id":   uint(668970),
		"form": uint(0),
		"type": "water",
	},
	"Seameleon": {
		"id":   uint(668971),
		"form": uint(0),
		"type": "water",
	},
	"Sealizard": {
		"id":   uint(668972),
		"form": uint(0),
		"type": "water",
	},
	"Fertle": {
		"id":   uint(668973),
		"form": uint(0),
		"type": "grass",
	},
	"Rootortle": {
		"id":   uint(668974),
		"form": uint(0),
		"type": "grass",
	},
	"Grasstoise": {
		"id":   uint(668975),
		"form": uint(0),
		"type": "grass",
	},
	"Caterpebb": {
		"id":   uint(668976),
		"form": uint(0),
		"type": "rock",
	},
	"Minepod": {
		"id":   uint(668977),
		"form": uint(0),
		"type": "rock",
	},
	"Boulderfree": {
		"id":   uint(668978),
		"form": uint(0),
		"type": "rock",
	},
	"Geodle": {
		"id":   uint(668979),
		"form": uint(0),
		"type": "rock",
	},
	"Prikuna": {
		"id":   uint(668980),
		"form": uint(0),
		"type": "rock",
	},
	"Bristrill": {
		"id":   uint(668981),
		"form": uint(0),
		"type": "rock",
	},
	"Fridgey": {
		"id":   uint(668982),
		"form": uint(0),
		"type": "ice",
	},
	"Fridgeotto": {
		"id":   uint(668983),
		"form": uint(0),
		"type": "ice",
	},
	"Fridgeot": {
		"id":   uint(668984),
		"form": uint(0),
		"type": "ice",
	},
	"Rottata": {
		"id":   uint(668985),
		"form": uint(0),
		"type": "zombie",
	},
	"Roticate": {
		"id":   uint(668986),
		"form": uint(0),
		"type": "zombie",
	},
	"Ghastata": {
		"id":   uint(668987),
		"form": uint(0),
		"type": "normal",
	},
	"Ghasticate": {
		"id":   uint(668988),
		"form": uint(0),
		"type": "normal",
	},
	"Spearolt": {
		"id":   uint(668989),
		"form": uint(0),
		"type": "electric",
	},
	"Fearolt": {
		"id":   uint(668990),
		"form": uint(0),
		"type": "electric",
	},
	"Mechans": {
		"id":   uint(668991),
		"form": uint(0),
		"type": "tech",
	},
	"Mecharbok": {
		"id":   uint(668992),
		"form": uint(0),
		"type": "tech",
	},
	"Pikacoal": {
		"id":   uint(668993),
		"form": uint(0),
		"type": "ground",
	},
	"Raicoal": {
		"id":   uint(668994),
		"form": uint(0),
		"type": "ground",
	},
	"Bushrew": {
		"id":   uint(668995),
		"form": uint(0),
		"type": "grass",
	},
	"Bushlash": {
		"id":   uint(668996),
		"form": uint(0),
		"type": "grass",
	},
	"Boxiranfe": {
		"id":   uint(668997),
		"form": uint(0),
		"type": "paper",
	},
	"Boxirina": {
		"id":   uint(668998),
		"form": uint(0),
		"type": "paper",
	},
	"Boxiqueen": {
		"id":   uint(668999),
		"form": uint(0),
		"type": "poison",
	},
	"Boxiranma": {
		"id":   uint(669000),
		"form": uint(0),
		"type": "paper",
	},
	"Boxirino": {
		"id":   uint(669001),
		"form": uint(0),
		"type": "paper",
	},
	"Boxiking": {
		"id":   uint(669002),
		"form": uint(0),
		"type": "paper",
	},
	"Clefair": {
		"id":   uint(669003),
		"form": uint(0),
		"type": "flying",
	},
	"Inflable": {
		"id":   uint(669004),
		"form": uint(0),
		"type": "flying",
	},
	"Vulstyx": {
		"id":   uint(669005),
		"form": uint(0),
		"type": "dark",
	},
	"Sinistales": {
		"id":   uint(669006),
		"form": uint(0),
		"type": "fire",
	},
	"Jigglisect": {
		"id":   uint(669007),
		"form": uint(0),
		"type": "bug",
	},
	"Arachtuff": {
		"id":   uint(669008),
		"form": uint(0),
		"type": "bug",
	},
	"Sinjbat": {
		"id":   uint(669009),
		"form": uint(0),
		"type": "fire",
	},
	"Volbat": {
		"id":   uint(669010),
		"form": uint(0),
		"type": "fire",
	},
	"Odprong": {
		"id":   uint(669011),
		"form": uint(0),
		"type": "steel",
	},
	"Turboon": {
		"id":   uint(669012),
		"form": uint(0),
		"type": "steel",
	},
	"Panelplume": {
		"id":   uint(669013),
		"form": uint(0),
		"type": "steel",
	},
	"Faaras": {
		"id":   uint(669014),
		"form": uint(0),
		"type": "psychic",
	},
	"Omnisect": {
		"id":   uint(669015),
		"form": uint(0),
		"type": "bug",
	},
	"Fernonat": {
		"id":   uint(669016),
		"form": uint(0),
		"type": "grass",
	},
	"Fernomoth": {
		"id":   uint(669017),
		"form": uint(0),
		"type": "bug",
	},
	"Nimbett": {
		"id":   uint(669018),
		"form": uint(0),
		"type": "flying",
	},
	"Stratrio": {
		"id":   uint(669019),
		"form": uint(0),
		"type": "flying",
	},
	"Meowxie": {
		"id":   uint(669020),
		"form": uint(0),
		"type": "fairy",
	},
	"Magisian": {
		"id":   uint(669021),
		"form": uint(0),
		"type": "fairy",
	},
	"Sluduck": {
		"id":   uint(669022),
		"form": uint(0),
		"type": "poison",
	},
	"Taintduck": {
		"id":   uint(669023),
		"form": uint(0),
		"type": "poison",
	},
	"Mankember": {
		"id":   uint(669024),
		"form": uint(0),
		"type": "fire",
	},
	"Flameape": {
		"id":   uint(669025),
		"form": uint(0),
		"type": "fire",
	},
	"Iolithe": {
		"id":   uint(669026),
		"form": uint(0),
		"type": "electric",
	},
	"Ionnine": {
		"id":   uint(669027),
		"form": uint(0),
		"type": "electric",
	},
	"Mandiwag": {
		"id":   uint(669028),
		"form": uint(0),
		"type": "bug",
	},
	"Mandiwhirl": {
		"id":   uint(669029),
		"form": uint(0),
		"type": "bug",
	},
	"Mandiwrath": {
		"id":   uint(669030),
		"form": uint(0),
		"type": "bug",
	},
	"Abrark": {
		"id":   uint(669031),
		"form": uint(0),
		"type": "dark",
	},
	"Cutabra": {
		"id":   uint(669032),
		"form": uint(0),
		"type": "dark",
	},
	"Shinigazam": {
		"id":   uint(669033),
		"form": uint(0),
		"type": "dark",
	},
	"Zenchop": {
		"id":   uint(669034),
		"form": uint(0),
		"type": "psychic",
	},
	"Zenchoke": {
		"id":   uint(669035),
		"form": uint(0),
		"type": "psychic",
	},
	"Zenchamp": {
		"id":   uint(669036),
		"form": uint(0),
		"type": "psychic",
	},
	"Bellpowder": {
		"id":   uint(669037),
		"form": uint(0),
		"type": "fire",
	},
	"Cannonbell": {
		"id":   uint(669038),
		"form": uint(0),
		"type": "fire",
	},
	"Chernobell": {
		"id":   uint(669039),
		"form": uint(0),
		"type": "fire",
	},
	"Tentapunch": {
		"id":   uint(669040),
		"form": uint(0),
		"type": "fighting",
	},
	"Tentabrawl": {
		"id":   uint(669041),
		"form": uint(0),
		"type": "fighting",
	},
	"Aquadude": {
		"id":   uint(669042),
		"form": uint(0),
		"type": "water",
	},
	"Coraller": {
		"id":   uint(669043),
		"form": uint(0),
		"type": "water",
	},
	"Golacier": {
		"id":   uint(669044),
		"form": uint(0),
		"type": "water",
	},
	"Ponytamp": {
		"id":   uint(669045),
		"form": uint(0),
		"type": "water",
	},
	"Rapidouse": {
		"id":   uint(669046),
		"form": uint(0),
		"type": "water",
	},
	"Stalagnemite": {
		"id":   uint(669047),
		"form": uint(0),
		"type": "ice",
	},
	"Stalagneton": {
		"id":   uint(669048),
		"form": uint(0),
		"type": "ice",
	},
	"Farganc": {
		"id":   uint(669049),
		"form": uint(0),
		"type": "grass",
	},
	"Slowslab": {
		"id":   uint(669050),
		"form": uint(0),
		"type": "rock",
	},
	"Slowsidian": {
		"id":   uint(669051),
		"form": uint(0),
		"type": "rock",
	},
	"Drillduo": {
		"id":   uint(669052),
		"form": uint(0),
		"type": "ground",
	},
	"Drilldrio": {
		"id":   uint(669053),
		"form": uint(0),
		"type": "ground",
	},
	"Seewage": {
		"id":   uint(669054),
		"form": uint(0),
		"type": "poison",
	},
	"Pollugong": {
		"id":   uint(669055),
		"form": uint(0),
		"type": "poison",
	},
	"Grimag": {
		"id":   uint(669056),
		"form": uint(0),
		"type": "magma",
	},
	"Mukma": {
		"id":   uint(669057),
		"form": uint(0),
		"type": "magma",
	},
	"Celesder": {
		"id":   uint(669058),
		"form": uint(0),
		"type": "cosmic",
	},
	"Cloystar": {
		"id":   uint(669059),
		"form": uint(0),
		"type": "cosmic",
	},
	"Goomly": {
		"id":   uint(669060),
		"form": uint(0),
		"type": "normal",
	},
	"Paragoomer": {
		"id":   uint(669061),
		"form": uint(0),
		"type": "normal",
	},
	"Goomgar": {
		"id":   uint(669062),
		"form": uint(0),
		"type": "normal",
	},
	"Eelix": {
		"id":   uint(669063),
		"form": uint(0),
		"type": "water",
	},
	"Ectozee": {
		"id":   uint(669064),
		"form": uint(0),
		"type": "ghost",
	},
	"Hypnoplasm": {
		"id":   uint(669065),
		"form": uint(0),
		"type": "ghost",
	},
	"Krabbune": {
		"id":   uint(669066),
		"form": uint(0),
		"type": "ground",
	},
	"Kraggler": {
		"id":   uint(669067),
		"form": uint(0),
		"type": "water",
	},
	"Frientorb": {
		"id":   uint(669068),
		"form": uint(0),
		"type": "fairy",
	},
	"Luxutrode": {
		"id":   uint(669069),
		"form": uint(0),
		"type": "fairy",
	},
	"Swarmute": {
		"id":   uint(669070),
		"form": uint(0),
		"type": "bug",
	},
	"Exeggunest": {
		"id":   uint(669071),
		"form": uint(0),
		"type": "bug",
	},
	"Cublade": {
		"id":   uint(669072),
		"form": uint(0),
		"type": "dragon",
	},
	"Marowarior": {
		"id":   uint(669073),
		"form": uint(0),
		"type": "dragon",
	},
	"Crustenga": {
		"id":   uint(669074),
		"form": uint(0),
		"type": "rock",
	},
	"Crustone": {
		"id":   uint(669075),
		"form": uint(0),
		"type": "rock",
	},
	"Lickoral": {
		"id":   uint(669076),
		"form": uint(0),
		"type": "rock",
	},
	"Sparkking": {
		"id":   uint(669077),
		"form": uint(0),
		"type": "dark",
	},
	"Shockking": {
		"id":   uint(669078),
		"form": uint(0),
		"type": "dark",
	},
	"Drachorn": {
		"id":   uint(669079),
		"form": uint(0),
		"type": "dragon",
	},
	"Dracdon": {
		"id":   uint(669080),
		"form": uint(0),
		"type": "dragon",
	},
	"Basiney": {
		"id":   uint(669081),
		"form": uint(0),
		"type": "water",
	},
	"Lazzmela": {
		"id":   uint(669082),
		"form": uint(0),
		"type": "electric",
	},
	"Kankaghan": {
		"id":   uint(669083),
		"form": uint(0),
		"type": "fighting",
	},
	"Horsteed": {
		"id":   uint(669084),
		"form": uint(0),
		"type": "normal",
	},
	"Steedra": {
		"id":   uint(669085),
		"form": uint(0),
		"type": "normal",
	},
	"Golbee": {
		"id":   uint(669086),
		"form": uint(0),
		"type": "bug",
	},
	"Beeking": {
		"id":   uint(669087),
		"form": uint(0),
		"type": "bug",
	},
	"Starite": {
		"id":   uint(669088),
		"form": uint(0),
		"type": "fairy",
	},
	"Starau": {
		"id":   uint(669089),
		"form": uint(0),
		"type": "fairy",
	},
	"Mr Melee": {
		"id":   uint(669090),
		"form": uint(0),
		"type": "fighting",
	},
	"Sickler": {
		"id":   uint(669091),
		"form": uint(0),
		"type": "dark",
	},
	"Sirynx": {
		"id":   uint(669092),
		"form": uint(0),
		"type": "water",
	},
	"Necrablast": {
		"id":   uint(669093),
		"form": uint(0),
		"type": "ghost",
	},
	"Plasgar": {
		"id":   uint(669094),
		"form": uint(0),
		"type": "electric",
	},
	"Mincir": {
		"id":   uint(669095),
		"form": uint(0),
		"type": "bug",
	},
	"Taurice": {
		"id":   uint(669096),
		"form": uint(0),
		"type": "ice",
	},
	"Chilikarp": {
		"id":   uint(669097),
		"form": uint(0),
		"type": "fire",
	},
	"Flarados": {
		"id":   uint(669098),
		"form": uint(0),
		"type": "fire",
	},
	"Gunkpras": {
		"id":   uint(669099),
		"form": uint(0),
		"type": "poison",
	},
	"Clayto": {
		"id":   uint(669100),
		"form": uint(0),
		"type": "ground",
	},
	"Threevee": {
		"id":   uint(669101),
		"form": uint(0),
		"type": "normal",
	},
	"Drilleon": {
		"id":   uint(669102),
		"form": uint(0),
		"type": "ground",
	},
	"Insecteon": {
		"id":   uint(669103),
		"form": uint(0),
		"type": "bug",
	},
	"Caireon": {
		"id":   uint(669104),
		"form": uint(0),
		"type": "rock",
	},
	"Viragon": {
		"id":   uint(669105),
		"form": uint(0),
		"type": "poison",
	},
	"Sandanice": {
		"id":   uint(669106),
		"form": uint(0),
		"type": "ice",
	},
	"Sandasnow": {
		"id":   uint(669107),
		"form": uint(0),
		"type": "ice",
	},
	"Leviuto": {
		"id":   uint(669108),
		"form": uint(0),
		"type": "water",
	},
	"Leviatops": {
		"id":   uint(669109),
		"form": uint(0),
		"type": "water",
	},
	"Sykodactyl": {
		"id":   uint(669110),
		"form": uint(0),
		"type": "psychic",
	},
	"Snorsessed": {
		"id":   uint(669111),
		"form": uint(0),
		"type": "ghost",
	},
	"Cosmicuno": {
		"id":   uint(669112),
		"form": uint(0),
		"type": "cosmic",
	},
	"Mortaldos": {
		"id":   uint(669113),
		"form": uint(0),
		"type": "dark",
	},
	"Spectres": {
		"id":   uint(669114),
		"form": uint(0),
		"type": "ghost",
	},
	"Pixini": {
		"id":   uint(669115),
		"form": uint(0),
		"type": "fairy",
	},
	"Pixinair": {
		"id":   uint(669116),
		"form": uint(0),
		"type": "fairy",
	},
	"Pixinite": {
		"id":   uint(669117),
		"form": uint(0),
		"type": "fairy",
	},
	"Altermew": {
		"id":   uint(669118),
		"form": uint(0),
		"type": "normal",
	},
	"Mewzero": {
		"id":   uint(669119),
		"form": uint(0),
		"type": "psychic",
	},
	"Gslowbro": {
		"id":   uint(669120),
		"form": uint(0),
		"type": "poison",
	},
	"Sirganc": {
		"id":   uint(669121),
		"form": uint(0),
		"type": "fighting",
	},
	"Regishark": {
		"id":   uint(669122),
		"form": uint(0),
		"type": "water",
	},
	"Regimagma": {
		"id":   uint(669123),
		"form": uint(0),
		"type": "magma",
	},
	"Regiweather": {
		"id":   uint(669124),
		"form": uint(0),
		"type": "wack",
	},
	"Asrieldrem": {
		"id":   uint(669125),
		"form": uint(0),
		"type": "divine",
	},
	"Explobat": {
		"id":   uint(669126),
		"form": uint(0),
		"type": "fire",
	},
	"Amon Ra": {
		"id":   uint(669127),
		"form": uint(0),
		"type": "divine",
	},
	"Kaiju Sandrawler": {
		"id":   uint(669128),
		"form": uint(0),
		"type": "bug",
	},
	"Flamoozer": {
		"id":   uint(669129),
		"form": uint(0),
		"type": "fire",
	},
	"Phallusaur": {
		"id":   uint(669130),
		"form": uint(0),
		"type": "rock",
	},
	"Arctosaur": {
		"id":   uint(669131),
		"form": uint(0),
		"type": "ice",
	},
	"Assdeer": {
		"id":   uint(669132),
		"form": uint(0),
		"type": "grass",
	},
	"Cspoink": {
		"id":   uint(669133),
		"form": uint(0),
		"type": "cosmic",
	},
	"Cgrumpig": {
		"id":   uint(669134),
		"form": uint(0),
		"type": "cosmic",
	},
	"Ms Pain": {
		"id":   uint(669135),
		"form": uint(0),
		"type": "paint",
	},
	"Monster Pain": {
		"id":   uint(669136),
		"form": uint(0),
		"type": "paint",
	},
	"Diemic": {
		"id":   uint(669137),
		"form": uint(0),
		"type": "dark",
	},
	"Scrubber": {
		"id":   uint(669138),
		"form": uint(0),
		"type": "water",
	},
	"Clamppelo": {
		"id":   uint(669139),
		"form": uint(0),
		"type": "plastic",
	},
	"Rushedsalamence": {
		"id":   uint(669140),
		"form": uint(0),
		"type": "dragon",
	},
	"Vanilltank": {
		"id":   uint(669141),
		"form": uint(0),
		"type": "ice",
	},
	"Kingsdale": {
		"id":   uint(669142),
		"form": uint(0),
		"type": "normal",
	},
	"Eelectrix": {
		"id":   uint(669143),
		"form": uint(0),
		"type": "water",
	},
	"Blissea": {
		"id":   uint(669144),
		"form": uint(0),
		"type": "water",
	},
	"Slowspike": {
		"id":   uint(669145),
		"form": uint(0),
		"type": "rock",
	},
	"Ujamer": {
		"id":   uint(669146),
		"form": uint(0),
		"type": "bug",
	},
	"Ujadalord": {
		"id":   uint(669147),
		"form": uint(0),
		"type": "bug",
	},
	"Holloach": {
		"id":   uint(669148),
		"form": uint(0),
		"type": "grass",
	},
	"Illexash": {
		"id":   uint(669149),
		"form": uint(0),
		"type": "grass",
	},
	"Mandirit": {
		"id":   uint(669150),
		"form": uint(0),
		"type": "ghost",
	},
	"Scizorch": {
		"id":   uint(669151),
		"form": uint(0),
		"type": "dark",
	},
	"Swooloo": {
		"id":   uint(669152),
		"form": uint(0),
		"type": "steel",
	},
	"Sdubwool": {
		"id":   uint(669153),
		"form": uint(0),
		"type": "steel",
	},
	"Necratomic": {
		"id":   uint(669154),
		"form": uint(0),
		"type": "ghost",
	},
	"Plasmission": {
		"id":   uint(669155),
		"form": uint(0),
		"type": "nuclear",
	},
	"Dracperior": {
		"id":   uint(669156),
		"form": uint(0),
		"type": "dragon",
	},
	"Lazzplug": {
		"id":   uint(669157),
		"form": uint(0),
		"type": "electric",
	},
	"Reeficoral": {
		"id":   uint(669158),
		"form": uint(0),
		"type": "rock",
	},
	"Stalagcola": {
		"id":   uint(669159),
		"form": uint(0),
		"type": "ice",
	},
	"Trapunk": {
		"id":   uint(669160),
		"form": uint(0),
		"type": "dark",
	},
	"Motorava": {
		"id":   uint(669161),
		"form": uint(0),
		"type": "dark",
	},
	"Gangon": {
		"id":   uint(669162),
		"form": uint(0),
		"type": "dark",
	},
	"Plusheon": {
		"id":   uint(669163),
		"form": uint(0),
		"type": "fabric",
	},
	"Eobstagoon": {
		"id":   uint(669164),
		"form": uint(0),
		"type": "electric",
	},
	"Kamisama": {
		"id":   uint(669165),
		"form": uint(0),
		"type": "qmarks",
	},
	"Pkmnwack": {
		"id":   uint(669166),
		"form": uint(0),
		"type": "wack",
	},
	"Barceus": {
		"id":   uint(669167),
		"form": uint(0),
		"type": "qmarks",
	},
	"Bgiratina": {
		"id":   uint(669168),
		"form": uint(0),
		"type": "light",
	},
	"Bdarkrai": {
		"id":   uint(669169),
		"form": uint(0),
		"type": "dark",
	},
	"Brotom": {
		"id":   uint(669170),
		"form": uint(0),
		"type": "electric",
	},
	"Bheatran": {
		"id":   uint(669171),
		"form": uint(0),
		"type": "fire",
	},
	"Mikuloid": {
		"id":   uint(669172),
		"form": uint(0),
		"type": "cyber",
	},
	"Vokagamine": {
		"id":   uint(669173),
		"form": uint(0),
		"type": "cyber",
	},
	"Crewmongus": {
		"id":   uint(669174),
		"form": uint(0),
		"type": "normal",
	},
	"Chupacryptid": {
		"id":   uint(669175),
		"form": uint(0),
		"type": "fear",
	},
	"Prinsiel": {
		"id":   uint(669176),
		"form": uint(0),
		"type": "dark",
	},
	"Scpee1128": {
		"id":   uint(669177),
		"form": uint(0),
		"type": "water",
	},
	"Birbley": {
		"id":   uint(669178),
		"form": uint(0),
		"type": "normal",
	},
	"Ridteen": {
		"id":   uint(669179),
		"form": uint(0),
		"type": "dragon",
	},
	"Ridleyuge": {
		"id":   uint(669180),
		"form": uint(0),
		"type": "dragon",
	},
	"Mbboo": {
		"id":   uint(669181),
		"form": uint(0),
		"type": "ghost",
	},
	"King Boo": {
		"id":   uint(669182),
		"form": uint(0),
		"type": "ghost",
	},
	"Facefugger": {
		"id":   uint(669183),
		"form": uint(0),
		"type": "poison",
	},
	"Chestburste": {
		"id":   uint(669184),
		"form": uint(0),
		"type": "poison",
	},
	"Xenoalien": {
		"id":   uint(669185),
		"form": uint(0),
		"type": "poison",
	},
	"Predalien": {
		"id":   uint(669186),
		"form": uint(0),
		"type": "cosmic",
	},
	"Mokelembe": {
		"id":   uint(669187),
		"form": uint(0),
		"type": "fear",
	},
	"Garfielf": {
		"id":   uint(669188),
		"form": uint(0),
		"type": "food",
	},
	"Gorefielf": {
		"id":   uint(669189),
		"form": uint(0),
		"type": "food",
	},
	"Scpee3166": {
		"id":   uint(669190),
		"form": uint(0),
		"type": "food",
	},
	"Scpee939": {
		"id":   uint(669191),
		"form": uint(0),
		"type": "fear",
	},
	"Higurika": {
		"id":   uint(669192),
		"form": uint(0),
		"type": "normal",
	},
	"Shionzaki": {
		"id":   uint(669193),
		"form": uint(0),
		"type": "normal",
	},
	"Hanyuu": {
		"id":   uint(669194),
		"form": uint(0),
		"type": "divine",
	},
	"Mionzaki": {
		"id":   uint(669195),
		"form": uint(0),
		"type": "normal",
	},
	"Higusatoko": {
		"id":   uint(669196),
		"form": uint(0),
		"type": "normal",
	},
	"Higurena": {
		"id":   uint(669197),
		"form": uint(0),
		"type": "steel",
	},
	"Higukeiichi": {
		"id":   uint(669198),
		"form": uint(0),
		"type": "normal",
	},
	"Tomitake": {
		"id":   uint(669199),
		"form": uint(0),
		"type": "normal",
	},
	"Takano": {
		"id":   uint(669200),
		"form": uint(0),
		"type": "dark",
	},
	"Oishigurashi": {
		"id":   uint(669201),
		"form": uint(0),
		"type": "normal",
	},
	"Akasaka": {
		"id":   uint(669202),
		"form": uint(0),
		"type": "fighting",
	},
	"Higuirie": {
		"id":   uint(669203),
		"form": uint(0),
		"type": "normal",
	},
	"Higurachie": {
		"id":   uint(669204),
		"form": uint(0),
		"type": "normal",
	},
	"Higukasai": {
		"id":   uint(669205),
		"form": uint(0),
		"type": "normal",
	},
	"Akanezaki": {
		"id":   uint(669206),
		"form": uint(0),
		"type": "normal",
	},
	"Natsumiyoshi": {
		"id":   uint(669207),
		"form": uint(0),
		"type": "normal",
	},
	"Scpee106": {
		"id":   uint(669208),
		"form": uint(0),
		"type": "dark",
	},
	"Scpee096": {
		"id":   uint(669209),
		"form": uint(0),
		"type": "fear",
	},
	"Scpee096-Zen": {
		"id":   uint(669209),
		"form": uint(1),
		"type": "fear",
	},
	"Scpee049": {
		"id":   uint(669210),
		"form": uint(0),
		"type": "virus",
	},
	"Eeeee": {
		"id":   uint(669211),
		"form": uint(0),
		"type": "qmarks",
	},
	"Scpee504": {
		"id":   uint(669212),
		"form": uint(0),
		"type": "grass",
	},
	"Glaakultist": {
		"id":   uint(669213),
		"form": uint(0),
		"type": "fear",
	},
	"Jersevil": {
		"id":   uint(669214),
		"form": uint(0),
		"type": "fear",
	},
	"Bgrotle": {
		"id":   uint(669215),
		"form": uint(0),
		"type": "grass",
	},
	"Btorterra": {
		"id":   uint(669216),
		"form": uint(0),
		"type": "grass",
	},
	"Bprinplup": {
		"id":   uint(669217),
		"form": uint(0),
		"type": "water",
	},
	"Bempoleon": {
		"id":   uint(669218),
		"form": uint(0),
		"type": "water",
	},
	"Bshinx": {
		"id":   uint(669219),
		"form": uint(0),
		"type": "electric",
	},
	"Bachichi": {
		"id":   uint(669220),
		"form": uint(0),
		"type": "electric",
	},
	"Bachibacchi": {
		"id":   uint(669221),
		"form": uint(0),
		"type": "electric",
	},
	"Bcranidos": {
		"id":   uint(669222),
		"form": uint(0),
		"type": "rock",
	},
	"Brampardos": {
		"id":   uint(669223),
		"form": uint(0),
		"type": "rock",
	},
	"Bcombee": {
		"id":   uint(669224),
		"form": uint(0),
		"type": "bug",
	},
	"Bvespiquen": {
		"id":   uint(669225),
		"form": uint(0),
		"type": "bug",
	},
	"Bbronzong": {
		"id":   uint(669226),
		"form": uint(0),
		"type": "steel",
	},
	"Bgible": {
		"id":   uint(669227),
		"form": uint(0),
		"type": "dragon",
	},
	"Bgabite": {
		"id":   uint(669228),
		"form": uint(0),
		"type": "dragon",
	},
	"Bgarchomp": {
		"id":   uint(669229),
		"form": uint(0),
		"type": "dragon",
	},
	"Bhippopotas": {
		"id":   uint(669230),
		"form": uint(0),
		"type": "ground",
	},
	"Bhippowdon": {
		"id":   uint(669231),
		"form": uint(0),
		"type": "ground",
	},
	"Bcarnivine": {
		"id":   uint(669232),
		"form": uint(0),
		"type": "grass",
	},
	"Bfinneon": {
		"id":   uint(669233),
		"form": uint(0),
		"type": "water",
	},
	"Blumineon": {
		"id":   uint(669234),
		"form": uint(0),
		"type": "water",
	},
	"Bsnover": {
		"id":   uint(669235),
		"form": uint(0),
		"type": "grass",
	},
	"Babomasnow": {
		"id":   uint(669236),
		"form": uint(0),
		"type": "grass",
	},
	"Btogekiss": {
		"id":   uint(669237),
		"form": uint(0),
		"type": "fairy",
	},
	"Byanmega": {
		"id":   uint(669238),
		"form": uint(0),
		"type": "bug",
	},
	"Bprobopass": {
		"id":   uint(669239),
		"form": uint(0),
		"type": "rock",
	},
	"Bregigigas": {
		"id":   uint(669240),
		"form": uint(0),
		"type": "normal",
	},
	"Bshaymin": {
		"id":   uint(669241),
		"form": uint(0),
		"type": "grass",
	},
	"Sadakoring": {
		"id":   uint(669242),
		"form": uint(0),
		"type": "ghost",
	},
	"Scpee079": {
		"id":   uint(669243),
		"form": uint(0),
		"type": "cyber",
	},
	"Pyraheadmid": {
		"id":   uint(669244),
		"form": uint(0),
		"type": "fear",
	},
	"Bubbnurse": {
		"id":   uint(669245),
		"form": uint(0),
		"type": "fear",
	},
	"Pumpkinopunch": {
		"id":   uint(669246),
		"form": uint(0),
		"type": "ghost",
	},
	"Wolfmansk": {
		"id":   uint(669247),
		"form": uint(0),
		"type": "fear",
	},
	"Invisimask": {
		"id":   uint(669248),
		"form": uint(0),
		"type": "fear",
	},
	"Slendmeran": {
		"id":   uint(669249),
		"form": uint(0),
		"type": "fear",
	},
	"Hachishaku": {
		"id":   uint(669250),
		"form": uint(0),
		"type": "ghost",
	},
	"Trumporn": {
		"id":   uint(669251),
		"form": uint(0),
		"type": "fire",
	},
	"Trollel": {
		"id":   uint(669252),
		"form": uint(0),
		"type": "dark",
	},
	"Elaboruse": {
		"id":   uint(669253),
		"form": uint(0),
		"type": "dark",
	},
	"Wetloon": {
		"id":   uint(669254),
		"form": uint(0),
		"type": "water",
	},
	"Aqualoon": {
		"id":   uint(669255),
		"form": uint(0),
		"type": "water",
	},
	"Jubiloon": {
		"id":   uint(669256),
		"form": uint(0),
		"type": "water",
	},
	"Loleye": {
		"id":   uint(669257),
		"form": uint(0),
		"type": "heart",
	},
	"Oppeye": {
		"id":   uint(669258),
		"form": uint(0),
		"type": "heart",
	},
	"Thicclops": {
		"id":   uint(669259),
		"form": uint(0),
		"type": "heart",
	},
	"Huranium": {
		"id":   uint(669260),
		"form": uint(0),
		"type": "normal",
	},
	"Alcohog": {
		"id":   uint(669261),
		"form": uint(0),
		"type": "poison",
	},
	"Swineyard": {
		"id":   uint(669262),
		"form": uint(0),
		"type": "poison",
	},
	"Hillarge": {
		"id":   uint(669263),
		"form": uint(0),
		"type": "ground",
	},
	"Bigliff": {
		"id":   uint(669264),
		"form": uint(0),
		"type": "ground",
	},
	"Gigantain": {
		"id":   uint(669265),
		"form": uint(0),
		"type": "ground",
	},
	"Blanda": {
		"id":   uint(669266),
		"form": uint(0),
		"type": "normal",
	},
	"Frospidey": {
		"id":   uint(669267),
		"form": uint(0),
		"type": "bug",
	},
	"Froscorpio": {
		"id":   uint(669268),
		"form": uint(0),
		"type": "bug",
	},
	"Moustab": {
		"id":   uint(669269),
		"form": uint(0),
		"type": "steel",
	},
	"Daggerat": {
		"id":   uint(669270),
		"form": uint(0),
		"type": "steel",
	},
	"Sworodent": {
		"id":   uint(669271),
		"form": uint(0),
		"type": "steel",
	},
	"Spovum": {
		"id":   uint(669272),
		"form": uint(0),
		"type": "psychic",
	},
	"Boebii": {
		"id":   uint(669273),
		"form": uint(0),
		"type": "fairy",
	},
	"Bowaii": {
		"id":   uint(669274),
		"form": uint(0),
		"type": "fairy",
	},
	"Bowtesque": {
		"id":   uint(669275),
		"form": uint(0),
		"type": "fairy",
	},
	"Ploop": {
		"id":   uint(669276),
		"form": uint(0),
		"type": "poison",
	},
	"Dungus": {
		"id":   uint(669277),
		"form": uint(0),
		"type": "poison",
	},
	"Poolossus": {
		"id":   uint(669278),
		"form": uint(0),
		"type": "poison",
	},
	"Celesphere": {
		"id":   uint(669279),
		"form": uint(0),
		"type": "rock",
	},
	"Gullbage": {
		"id":   uint(669280),
		"form": uint(0),
		"type": "poison",
	},
	"Albatrash": {
		"id":   uint(669281),
		"form": uint(0),
		"type": "poison",
	},
	"Gruppy": {
		"id":   uint(669282),
		"form": uint(0),
		"type": "bug",
	},
	"Barkoon": {
		"id":   uint(669283),
		"form": uint(0),
		"type": "bug",
	},
	"Stingnarl": {
		"id":   uint(669284),
		"form": uint(0),
		"type": "bug",
	},
	"Pinnint": {
		"id":   uint(669285),
		"form": uint(0),
		"type": "normal",
	},
	"Badmini": {
		"id":   uint(669286),
		"form": uint(0),
		"type": "normal",
	},
	"Pinsochist": {
		"id":   uint(669287),
		"form": uint(0),
		"type": "ghost",
	},
	"Aphrenus": {
		"id":   uint(669288),
		"form": uint(0),
		"type": "dragon",
	},
	"Pengwarn": {
		"id":   uint(669289),
		"form": uint(0),
		"type": "ice",
	},
	"Banguin": {
		"id":   uint(669290),
		"form": uint(0),
		"type": "ice",
	},
	"Pantom": {
		"id":   uint(669291),
		"form": uint(0),
		"type": "ghost",
	},
	"Punchortz": {
		"id":   uint(669292),
		"form": uint(0),
		"type": "ghost",
	},
	"Lingereist": {
		"id":   uint(669293),
		"form": uint(0),
		"type": "ghost",
	},
	"Sparroof": {
		"id":   uint(669294),
		"form": uint(0),
		"type": "normal",
	},
	"Spastle": {
		"id":   uint(669295),
		"form": uint(0),
		"type": "rock",
	},
	"Spalcazar": {
		"id":   uint(669296),
		"form": uint(0),
		"type": "rock",
	},
	"Slugar": {
		"id":   uint(669297),
		"form": uint(0),
		"type": "fairy",
	},
	"Escarbun": {
		"id":   uint(669298),
		"form": uint(0),
		"type": "fairy",
	},
	"Pizzauro": {
		"id":   uint(669299),
		"form": uint(0),
		"type": "food",
	},
	"Mozzarex": {
		"id":   uint(669300),
		"form": uint(0),
		"type": "food",
	},
	"Pepperoari": {
		"id":   uint(669301),
		"form": uint(0),
		"type": "food",
	},
	"Nyuvel": {
		"id":   uint(669302),
		"form": uint(0),
		"type": "cyber",
	},
	"Nyuvel2": {
		"id":   uint(669303),
		"form": uint(0),
		"type": "cyber",
	},
	"Nyuvel0": {
		"id":   uint(669304),
		"form": uint(0),
		"type": "cyber",
	},
	"Swelteroid": {
		"id":   uint(669305),
		"form": uint(0),
		"type": "rock",
	},
	"Winteroid": {
		"id":   uint(669306),
		"form": uint(0),
		"type": "rock",
	},
	"Yellynth": {
		"id":   uint(669307),
		"form": uint(0),
		"type": "sound",
	},
	"Magiynth": {
		"id":   uint(669308),
		"form": uint(0),
		"type": "sound",
	},
	"Orchesynth": {
		"id":   uint(669309),
		"form": uint(0),
		"type": "sound",
	},
	"Twiggerd": {
		"id":   uint(669310),
		"form": uint(0),
		"type": "wood",
	},
	"Trunkannon": {
		"id":   uint(669311),
		"form": uint(0),
		"type": "wood",
	},
	"Pictoji": {
		"id":   uint(669312),
		"form": uint(0),
		"type": "rock",
	},
	"Schnozard": {
		"id":   uint(669313),
		"form": uint(0),
		"type": "dragon",
	},
	"Schnormous": {
		"id":   uint(669314),
		"form": uint(0),
		"type": "dragon",
	},
	"Squbble": {
		"id":   uint(669315),
		"form": uint(0),
		"type": "water",
	},
	"Kraclean": {
		"id":   uint(669316),
		"form": uint(0),
		"type": "water",
	},
	"Cartcap": {
		"id":   uint(669317),
		"form": uint(0),
		"type": "ice",
	},
	"Molstub": {
		"id":   uint(669318),
		"form": uint(0),
		"type": "ground",
	},
	"Molstache": {
		"id":   uint(669319),
		"form": uint(0),
		"type": "ground",
	},
	"Molbeardo": {
		"id":   uint(669320),
		"form": uint(0),
		"type": "ground",
	},
	"Cyunicoji": {
		"id":   uint(669321),
		"form": uint(0),
		"type": "psychic",
	},
	"Beanoy": {
		"id":   uint(669322),
		"form": uint(0),
		"type": "grass",
	},
	"Oniogre": {
		"id":   uint(669323),
		"form": uint(0),
		"type": "grass",
	},
	"Cuffy": {
		"id":   uint(669324),
		"form": uint(0),
		"type": "steel",
	},
	"Whippain": {
		"id":   uint(669325),
		"form": uint(0),
		"type": "steel",
	},
	"Collatrix": {
		"id":   uint(669326),
		"form": uint(0),
		"type": "steel",
	},
	"Hairjog": {
		"id":   uint(669327),
		"form": uint(0),
		"type": "normal",
	},
	"Kerajog": {
		"id":   uint(669328),
		"form": uint(0),
		"type": "normal",
	},
	"Senteddi": {
		"id":   uint(669329),
		"form": uint(0),
		"type": "fighting",
	},
	"Urscort": {
		"id":   uint(669330),
		"form": uint(0),
		"type": "fighting",
	},
	"Tardough": {
		"id":   uint(669331),
		"form": uint(0),
		"type": "food",
	},
	"Breadiot": {
		"id":   uint(669332),
		"form": uint(0),
		"type": "food",
	},
	"Pastiot": {
		"id":   uint(669333),
		"form": uint(0),
		"type": "food",
	},
	"Cookiot": {
		"id":   uint(669334),
		"form": uint(0),
		"type": "food",
	},
	"Oinxploda": {
		"id":   uint(669335),
		"form": uint(0),
		"type": "dark",
	},
	"Salsik": {
		"id":   uint(669336),
		"form": uint(0),
		"type": "water",
	},
	"Sikrak": {
		"id":   uint(669337),
		"form": uint(0),
		"type": "water",
	},
	"Sikteria": {
		"id":   uint(669338),
		"form": uint(0),
		"type": "water",
	},
	"Autistrain": {
		"id":   uint(669339),
		"form": uint(0),
		"type": "steel",
	},
	"Teevz": {
		"id":   uint(669340),
		"form": uint(0),
		"type": "grass",
	},
	"Teepod": {
		"id":   uint(669341),
		"form": uint(0),
		"type": "grass",
	},
	"Cuppatee": {
		"id":   uint(669342),
		"form": uint(0),
		"type": "grass",
	},
	"Oymayt": {
		"id":   uint(669343),
		"form": uint(0),
		"type": "dark",
	},
	"Yuwotmayt": {
		"id":   uint(669344),
		"form": uint(0),
		"type": "dark",
	},
	"Pinana": {
		"id":   uint(669345),
		"form": uint(0),
		"type": "grass",
	},
	"Peelngas": {
		"id":   uint(669346),
		"form": uint(0),
		"type": "grass",
	},
	"Splug": {
		"id":   uint(669347),
		"form": uint(0),
		"type": "electric",
	},
	"Adaplug": {
		"id":   uint(669348),
		"form": uint(0),
		"type": "electric",
	},
	"Cabilisk": {
		"id":   uint(669349),
		"form": uint(0),
		"type": "electric",
	},
	"Flopple": {
		"id":   uint(669350),
		"form": uint(0),
		"type": "grass",
	},
	"Buktopple": {
		"id":   uint(669351),
		"form": uint(0),
		"type": "grass",
	},
	"Corottle": {
		"id":   uint(669352),
		"form": uint(0),
		"type": "grass",
	},
	"Fleap": {
		"id":   uint(669353),
		"form": uint(0),
		"type": "bug",
	},
	"Flitchi": {
		"id":   uint(669354),
		"form": uint(0),
		"type": "bug",
	},
	"Ostreoch": {
		"id":   uint(669355),
		"form": uint(0),
		"type": "ghost",
	},
	"Shlario": {
		"id":   uint(669356),
		"form": uint(0),
		"type": "normal",
	},
	"Razario": {
		"id":   uint(669357),
		"form": uint(0),
		"type": "normal",
	},
	"Snotto": {
		"id":   uint(669358),
		"form": uint(0),
		"type": "poison",
	},
	"Boglob": {
		"id":   uint(669359),
		"form": uint(0),
		"type": "poison",
	},
	"Fluinter": {
		"id":   uint(669360),
		"form": uint(0),
		"type": "poison",
	},
	"Ovilock": {
		"id":   uint(669361),
		"form": uint(0),
		"type": "steel",
	},
	"Alkeydin": {
		"id":   uint(669362),
		"form": uint(0),
		"type": "steel",
	},
	"Clownim": {
		"id":   uint(669363),
		"form": uint(0),
		"type": "psychic",
	},
	"Tentaclown": {
		"id":   uint(669364),
		"form": uint(0),
		"type": "psychic",
	},
	"Zizizii": {
		"id":   uint(669365),
		"form": uint(0),
		"type": "normal",
	},
	"Tinna": {
		"id":   uint(669366),
		"form": uint(0),
		"type": "ghost",
	},
	"Skeletuna": {
		"id":   uint(669367),
		"form": uint(0),
		"type": "ghost",
	},
	"Capseal": {
		"id":   uint(669368),
		"form": uint(0),
		"type": "water",
	},
	"Nurseal": {
		"id":   uint(669369),
		"form": uint(0),
		"type": "water",
	},
	"Manursee": {
		"id":   uint(669370),
		"form": uint(0),
		"type": "water",
	},
	"Shizzlip": {
		"id":   uint(669371),
		"form": uint(0),
		"type": "dark",
	},
	"Marmin": {
		"id":   uint(669372),
		"form": uint(0),
		"type": "bug",
	},
	"Marmight": {
		"id":   uint(669373),
		"form": uint(0),
		"type": "bug",
	},
	"Chairor": {
		"id":   uint(669374),
		"form": uint(0),
		"type": "dragon",
	},
	"Throgon": {
		"id":   uint(669375),
		"form": uint(0),
		"type": "dragon",
	},
	"Hologg": {
		"id":   uint(669376),
		"form": uint(0),
		"type": "electric",
	},
	"Holectric": {
		"id":   uint(669377),
		"form": uint(0),
		"type": "electric",
	},
	"Zelenvus": {
		"id":   uint(669378),
		"form": uint(0),
		"type": "dragon",
	},
	"Clovote": {
		"id":   uint(669379),
		"form": uint(0),
		"type": "bug",
	},
	"Clagma": {
		"id":   uint(669380),
		"form": uint(0),
		"type": "magma",
	},
	"Lavalve": {
		"id":   uint(669381),
		"form": uint(0),
		"type": "magma",
	},
	"Volclamno": {
		"id":   uint(669382),
		"form": uint(0),
		"type": "magma",
	},
	"Trichic": {
		"id":   uint(669383),
		"form": uint(0),
		"type": "psychic",
	},
	"Squaychic": {
		"id":   uint(669384),
		"form": uint(0),
		"type": "psychic",
	},
	"Pentragon": {
		"id":   uint(669385),
		"form": uint(0),
		"type": "psychic",
	},
	"Sodrade": {
		"id":   uint(669386),
		"form": uint(0),
		"type": "water",
	},
	"Tabbee": {
		"id":   uint(669387),
		"form": uint(0),
		"type": "bug",
	},
	"Tabbumble": {
		"id":   uint(669388),
		"form": uint(0),
		"type": "bug",
	},
	"Eaglass": {
		"id":   uint(669389),
		"form": uint(0),
		"type": "ice",
	},
	"Threaglass": {
		"id":   uint(669390),
		"form": uint(0),
		"type": "ice",
	},
	"Gouwheel": {
		"id":   uint(669391),
		"form": uint(0),
		"type": "steel",
	},
	"Hauntaxi": {
		"id":   uint(669392),
		"form": uint(0),
		"type": "steel",
	},
	"Bussacre": {
		"id":   uint(669393),
		"form": uint(0),
		"type": "steel",
	},
	"Chubbirb": {
		"id":   uint(669394),
		"form": uint(0),
		"type": "normal",
	},
	"Birbutter": {
		"id":   uint(669395),
		"form": uint(0),
		"type": "normal",
	},
	"Goleg": {
		"id":   uint(669396),
		"form": uint(0),
		"type": "rock",
	},
	"Figoleg": {
		"id":   uint(669397),
		"form": uint(0),
		"type": "rock",
	},
	"Fagworm": {
		"id":   uint(669398),
		"form": uint(0),
		"type": "poison",
	},
	"Cigraworm": {
		"id":   uint(669399),
		"form": uint(0),
		"type": "poison",
	},
	"Powdillar": {
		"id":   uint(669400),
		"form": uint(0),
		"type": "bug",
	},
	"Grencoon": {
		"id":   uint(669401),
		"form": uint(0),
		"type": "bug",
	},
	"Exploth": {
		"id":   uint(669402),
		"form": uint(0),
		"type": "bug",
	},
	"Nimbud": {
		"id":   uint(669403),
		"form": uint(0),
		"type": "grass",
	},
	"Bonskai": {
		"id":   uint(669404),
		"form": uint(0),
		"type": "wood",
	},
	"Tyrannomon": {
		"id":   uint(669405),
		"form": uint(0),
		"type": "fire",
	},
	"Mastertyrannomon": {
		"id":   uint(669406),
		"form": uint(0),
		"type": "dark",
	},
	"Samudramon": {
		"id":   uint(669407),
		"form": uint(0),
		"type": "fighting",
	},
	"Gslowking": {
		"id":   uint(669408),
		"form": uint(0),
		"type": "poison",
	},
	"Rlimshaikworm": {
		"id":   uint(669409),
		"form": uint(0),
		"type": "ice",
	},
	"Clefairicon": {
		"id":   uint(669410),
		"form": uint(0),
		"type": "fairy",
	},
	"Cheezipake": {
		"id":   uint(669411),
		"form": uint(0),
		"type": "food",
	},
	"Slasherpie": {
		"id":   uint(669412),
		"form": uint(0),
		"type": "bug",
	},
	"Metalblud": {
		"id":   uint(669413),
		"form": uint(0),
		"type": "bug",
	},
	"Scorpiofree": {
		"id":   uint(669414),
		"form": uint(0),
		"type": "fear",
	},
	"Starruz": {
		"id":   uint(669415),
		"form": uint(0),
		"type": "cosmic",
	},
	"Intestinela": {
		"id":   uint(669416),
		"form": uint(0),
		"type": "blood",
	},
	"Entrailgrowth": {
		"id":   uint(669417),
		"form": uint(0),
		"type": "blood",
	},
	"Kickkecleon": {
		"id":   uint(669418),
		"form": uint(0),
		"type": "normal",
	},
	"Kerkecleon": {
		"id":   uint(669419),
		"form": uint(0),
		"type": "meme",
	},
	"Eeeveelatte": {
		"id":   uint(669420),
		"form": uint(0),
		"type": "food",
	},
	"Litwickcocoa": {
		"id":   uint(669421),
		"form": uint(0),
		"type": "food",
	},
	"Liliganttea": {
		"id":   uint(669422),
		"form": uint(0),
		"type": "food",
	},
	"Bunearyfrappe": {
		"id":   uint(669423),
		"form": uint(0),
		"type": "food",
	},
	"Pachirisufloat": {
		"id":   uint(669424),
		"form": uint(0),
		"type": "food",
	},
	"Gratinpumpkaboo": {
		"id":   uint(669425),
		"form": uint(0),
		"type": "food",
	},
	"Thiccremie": {
		"id":   uint(669426),
		"form": uint(0),
		"type": "heart",
	},
	"Characide": {
		"id":   uint(669427),
		"form": uint(0),
		"type": "fighting",
	},
	"Bsgrava": {
		"id":   uint(669428),
		"form": uint(0),
		"type": "steel",
	},
	"Pikachucurry": {
		"id":   uint(669429),
		"form": uint(0),
		"type": "food",
	},
	"Rowlet Pizza": {
		"id":   uint(669430),
		"form": uint(0),
		"type": "grass",
	},
	"Dugtriosandwich": {
		"id":   uint(669431),
		"form": uint(0),
		"type": "ground",
	},
	"Oricoriopopcorn": {
		"id":   uint(669432),
		"form": uint(0),
		"type": "fire",
	},
	"Yamperpasta": {
		"id":   uint(669433),
		"form": uint(0),
		"type": "electric",
	},
	"Vulpixsundae": {
		"id":   uint(669434),
		"form": uint(0),
		"type": "food",
	},
	"Combeewaffles": {
		"id":   uint(669435),
		"form": uint(0),
		"type": "bug",
	},
	"Alcremiecupcakes": {
		"id":   uint(669436),
		"form": uint(0),
		"type": "fairy",
	},
	"Eeveepancakes": {
		"id":   uint(669437),
		"form": uint(0),
		"type": "food",
	},
	"Swablushavedice": {
		"id":   uint(669438),
		"form": uint(0),
		"type": "ice",
	},
	"Gossifleurcombo": {
		"id":   uint(669439),
		"form": uint(0),
		"type": "grass",
	},
	"Miltank Lait": {
		"id":   uint(669440),
		"form": uint(0),
		"type": "food",
	},
	"Eiscuecaprese": {
		"id":   uint(669441),
		"form": uint(0),
		"type": "food",
	},
	"Ribombeecake": {
		"id":   uint(669442),
		"form": uint(0),
		"type": "food",
	},
	"Brionnesoda": {
		"id":   uint(669443),
		"form": uint(0),
		"type": "food",
	},
	"Pikachufruitflan": {
		"id":   uint(669444),
		"form": uint(0),
		"type": "electric",
	},
	"Snorlaxlocomoco": {
		"id":   uint(669445),
		"form": uint(0),
		"type": "normal",
	},
	"Teddiursaicedcoffee": {
		"id":   uint(669446),
		"form": uint(0),
		"type": "food",
	},
	"Torchicomelet": {
		"id":   uint(669447),
		"form": uint(0),
		"type": "fire",
	},
	"Scorbunnysandwich": {
		"id":   uint(669448),
		"form": uint(0),
		"type": "fire",
	},
	"Chikoritabruschetta": {
		"id":   uint(669449),
		"form": uint(0),
		"type": "grass",
	},
	"Mimikyufeast": {
		"id":   uint(669450),
		"form": uint(0),
		"type": "ghost",
	},
	"Eternalfloette": {
		"id":   uint(669451),
		"form": uint(0),
		"type": "fairy",
	},
	"Bsshell": {
		"id":   uint(669452),
		"form": uint(0),
		"type": "normal",
	},
	"Shrowser": {
		"id":   uint(669453),
		"form": uint(0),
		"type": "fire",
	},
	"Papyrousle": {
		"id":   uint(669454),
		"form": uint(0),
		"type": "light",
	},
	"Hammerygon": {
		"id":   uint(669455),
		"form": uint(0),
		"type": "normal",
	},
	"Swordygon": {
		"id":   uint(669456),
		"form": uint(0),
		"type": "normal",
	},
	"Shieldygon": {
		"id":   uint(669457),
		"form": uint(0),
		"type": "normal",
	},
	"Cosplay Pikachu": {
		"id":   uint(669458),
		"form": uint(0),
		"type": "electric",
	},
	"Bgroudon": {
		"id":   uint(669459),
		"form": uint(0),
		"type": "fire",
	},
	"Btreecko": {
		"id":   uint(669460),
		"form": uint(0),
		"type": "grass",
	},
	"Btorchic": {
		"id":   uint(669461),
		"form": uint(0),
		"type": "fire",
	},
	"Latiaziken": {
		"id":   uint(669462),
		"form": uint(0),
		"type": "fire",
	},
	"Bgastrodon": {
		"id":   uint(669463),
		"form": uint(0),
		"type": "water",
	},
	"Hidelra": {
		"id":   uint(669464),
		"form": uint(0),
		"type": "normal",
	},
	"Huldire": {
		"id":   uint(669465),
		"form": uint(0),
		"type": "fear",
	},
	"Trolluhyd": {
		"id":   uint(669466),
		"form": uint(0),
		"type": "rock",
	},
	"Consumetsou": {
		"id":   uint(669467),
		"form": uint(0),
		"type": "fear",
	},
	"Bigmallow": {
		"id":   uint(669468),
		"form": uint(0),
		"type": "heart",
	},
	"Ballough": {
		"id":   uint(669469),
		"form": uint(0),
		"type": "food",
	},
	"Hunballun": {
		"id":   uint(669470),
		"form": uint(0),
		"type": "food",
	},
	"Sneaviwrath": {
		"id":   uint(669471),
		"form": uint(0),
		"type": "shadow",
	},
	"Dmagikarp": {
		"id":   uint(669472),
		"form": uint(0),
		"type": "water",
	},
	"Dgyarados": {
		"id":   uint(669473),
		"form": uint(0),
		"type": "water",
	},
	"Gmimejr": {
		"id":   uint(669474),
		"form": uint(0),
		"type": "psychic",
	},
	"Krampird": {
		"id":   uint(669475),
		"form": uint(0),
		"type": "ghost",
	},
	"Santajynx": {
		"id":   uint(669476),
		"form": uint(0),
		"type": "ice",
	},
	"Hyonix": {
		"id":   uint(669477),
		"form": uint(0),
		"type": "steel",
	},
	"Malicix": {
		"id":   uint(669478),
		"form": uint(0),
		"type": "steel",
	},
	"Miasva": {
		"id":   uint(669479),
		"form": uint(0),
		"type": "bug",
	},
	"Fgroudon": {
		"id":   uint(669480),
		"form": uint(0),
		"type": "qmarks",
	},
	"Fkyogre": {
		"id":   uint(669481),
		"form": uint(0),
		"type": "qmarks",
	},
	"Cmsnivy": {
		"id":   uint(669482),
		"form": uint(0),
		"type": "grass",
	},
	"Cmservine": {
		"id":   uint(669483),
		"form": uint(0),
		"type": "grass",
	},
	"Cmserperior": {
		"id":   uint(669484),
		"form": uint(0),
		"type": "grass",
	},
	"Splusle": {
		"id":   uint(669485),
		"form": uint(0),
		"type": "ice",
	},
	"Sminun": {
		"id":   uint(669486),
		"form": uint(0),
		"type": "ice",
	},
	"Cmstantler": {
		"id":   uint(669487),
		"form": uint(0),
		"type": "normal",
	},
	"Armewtwo": {
		"id":   uint(669488),
		"form": uint(0),
		"type": "psychic",
	},
	"Taswellow": {
		"id":   uint(669489),
		"form": uint(0),
		"type": "normal",
	},
	"Kangaskhid": {
		"id":   uint(669490),
		"form": uint(0),
		"type": "normal",
	},
	"Gkoffing": {
		"id":   uint(669491),
		"form": uint(0),
		"type": "poison",
	},
	"Zgardevoir": {
		"id":   uint(669492),
		"form": uint(0),
		"type": "blood",
	},
	"Seqola": {
		"id":   uint(669493),
		"form": uint(0),
		"type": "grass",
	},
	"Sequoisola": {
		"id":   uint(669494),
		"form": uint(0),
		"type": "grass",
	},
	"Zotsune": {
		"id":   uint(669495),
		"form": uint(0),
		"type": "psychic",
	},
	"Zoroterasu": {
		"id":   uint(669496),
		"form": uint(0),
		"type": "psychic",
	},
	"Bristik": {
		"id":   uint(669497),
		"form": uint(0),
		"type": "bug",
	},
	"Airantula": {
		"id":   uint(669498),
		"form": uint(0),
		"type": "bug",
	},
	"Hsnorunt": {
		"id":   uint(669499),
		"form": uint(0),
		"type": "fear",
	},
	"Hglalie": {
		"id":   uint(669500),
		"form": uint(0),
		"type": "fear",
	},
	"Hfroslass": {
		"id":   uint(669501),
		"form": uint(0),
		"type": "fear",
	},
	"Merchum": {
		"id":   uint(669502),
		"form": uint(0),
		"type": "water",
	},
	"Mermynx": {
		"id":   uint(669503),
		"form": uint(0),
		"type": "water",
	},
	"Hex Maniac": {
		"id":   uint(669504),
		"form": uint(0),
		"type": "ghost",
	},
	"Waifuractus": {
		"id":   uint(669505),
		"form": uint(0),
		"type": "grass",
	},
	"Princess Peach": {
		"id":   uint(669506),
		"form": uint(0),
		"type": "normal",
	},
	"Princess Daisy": {
		"id":   uint(669507),
		"form": uint(0),
		"type": "fighting",
	},
	"Donkey Kong": {
		"id":   uint(669508),
		"form": uint(0),
		"type": "normal",
	},
	"Lozlink": {
		"id":   uint(669509),
		"form": uint(0),
		"type": "fairy",
	},
	"Samus": {
		"id":   uint(669510),
		"form": uint(0),
		"type": "cosmic",
	},
	"Dark Samus": {
		"id":   uint(669511),
		"form": uint(0),
		"type": "cosmic",
	},
	"Kirby": {
		"id":   uint(669512),
		"form": uint(0),
		"type": "normal",
	},
	"Starfox": {
		"id":   uint(669513),
		"form": uint(0),
		"type": "normal",
	},
	"Ness": {
		"id":   uint(669514),
		"form": uint(0),
		"type": "psychic",
	},
	"Captain Falcon": {
		"id":   uint(669515),
		"form": uint(0),
		"type": "fire",
	},
	"Starfalco": {
		"id":   uint(669516),
		"form": uint(0),
		"type": "flying",
	},
	"Dmcdante": {
		"id":   uint(669517),
		"form": uint(0),
		"type": "steel",
	},
	"Ice Climbers": {
		"id":   uint(669518),
		"form": uint(0),
		"type": "ice",
	},
	"Lozzelda": {
		"id":   uint(669519),
		"form": uint(0),
		"type": "fairy",
	},
	"Lozsheik": {
		"id":   uint(669520),
		"form": uint(0),
		"type": "fairy",
	},
	"Starwolf": {
		"id":   uint(669521),
		"form": uint(0),
		"type": "dark",
	},
	"Wario": {
		"id":   uint(669522),
		"form": uint(0),
		"type": "poison",
	},
	"Ikelord": {
		"id":   uint(669523),
		"form": uint(0),
		"type": "steel",
	},
	"Chromlord": {
		"id":   uint(669524),
		"form": uint(0),
		"type": "steel",
	},
	"Game And Watch": {
		"id":   uint(669525),
		"form": uint(0),
		"type": "cyber",
	},
	"Meta Knight": {
		"id":   uint(669526),
		"form": uint(0),
		"type": "dark",
	},
	"King Dedede": {
		"id":   uint(669527),
		"form": uint(0),
		"type": "greasy",
	},
	"Dr Mario": {
		"id":   uint(669528),
		"form": uint(0),
		"type": "normal",
	},
	"Young Link": {
		"id":   uint(669529),
		"form": uint(0),
		"type": "fairy",
	},
	"Ganondorf": {
		"id":   uint(669530),
		"form": uint(0),
		"type": "dark",
	},
	"Kipit": {
		"id":   uint(669531),
		"form": uint(0),
		"type": "divine",
	},
	"Dark Pit": {
		"id":   uint(669532),
		"form": uint(0),
		"type": "shadow",
	},
	"Palutena": {
		"id":   uint(669533),
		"form": uint(0),
		"type": "divine",
	},
	"Sonichedgehog": {
		"id":   uint(669534),
		"form": uint(0),
		"type": "normal",
	},
	"Diddy Kong": {
		"id":   uint(669535),
		"form": uint(0),
		"type": "normal",
	},
	"Eblucas": {
		"id":   uint(669536),
		"form": uint(0),
		"type": "psychic",
	},
	"Solid Snake": {
		"id":   uint(669537),
		"form": uint(0),
		"type": "dark",
	},
	"Olimar": {
		"id":   uint(669538),
		"form": uint(0),
		"type": "grass",
	},
	"Rob": {
		"id":   uint(669539),
		"form": uint(0),
		"type": "tech",
	},
	"Toon Link": {
		"id":   uint(669540),
		"form": uint(0),
		"type": "fairy",
	},
	"Acvillager": {
		"id":   uint(669541),
		"form": uint(0),
		"type": "normal",
	},
	"Wii Fit Trainer": {
		"id":   uint(669542),
		"form": uint(0),
		"type": "normal",
	},
	"Rosalina": {
		"id":   uint(669543),
		"form": uint(0),
		"type": "cosmic",
	},
	"Little Mac": {
		"id":   uint(669544),
		"form": uint(0),
		"type": "fighting",
	},
	"Bowser Jr": {
		"id":   uint(669545),
		"form": uint(0),
		"type": "fire",
	},
	"Duck Hunt Dog": {
		"id":   uint(669546),
		"form": uint(0),
		"type": "flying",
	},
	"Pacman": {
		"id":   uint(669547),
		"form": uint(0),
		"type": "normal",
	},
	"Mii Brawler": {
		"id":   uint(669548),
		"form": uint(0),
		"type": "normal",
	},
	"Mii Swordfighter": {
		"id":   uint(669549),
		"form": uint(0),
		"type": "normal",
	},
	"Mii Gunner": {
		"id":   uint(669550),
		"form": uint(0),
		"type": "normal",
	},
	"Ferobin": {
		"id":   uint(669551),
		"form": uint(0),
		"type": "magic",
	},
	"Sfryu": {
		"id":   uint(669552),
		"form": uint(0),
		"type": "fighting",
	},
	"Sfken": {
		"id":   uint(669553),
		"form": uint(0),
		"type": "fighting",
	},
	"Chun Li": {
		"id":   uint(669554),
		"form": uint(0),
		"type": "fighting",
	},
	"Sfcammy": {
		"id":   uint(669555),
		"form": uint(0),
		"type": "fighting",
	},
	"Dsfelicia": {
		"id":   uint(669556),
		"form": uint(0),
		"type": "normal",
	},
	"Acisabelle": {
		"id":   uint(669557),
		"form": uint(0),
		"type": "normal",
	},
	"Dqhero": {
		"id":   uint(669558),
		"form": uint(0),
		"type": "steel",
	},
	"Fecorrin": {
		"id":   uint(669559),
		"form": uint(0),
		"type": "dragon",
	},
	"Xbshulk": {
		"id":   uint(669560),
		"form": uint(0),
		"type": "psychic",
	},
	"Sinkling": {
		"id":   uint(669561),
		"form": uint(0),
		"type": "paint",
	},
	"Bayonetta": {
		"id":   uint(669562),
		"form": uint(0),
		"type": "dark",
	},
	"Ffcloud": {
		"id":   uint(669563),
		"form": uint(0),
		"type": "steel",
	},
	"Mega Man": {
		"id":   uint(669564),
		"form": uint(0),
		"type": "tech",
	},
	"Simon Belmont": {
		"id":   uint(669565),
		"form": uint(0),
		"type": "fighting",
	},
	"Richter Belmont": {
		"id":   uint(669566),
		"form": uint(0),
		"type": "fighting",
	},
	"King K Rool": {
		"id":   uint(669567),
		"form": uint(0),
		"type": "ground",
	},
	"Pirahnaplant": {
		"id":   uint(669568),
		"form": uint(0),
		"type": "grass",
	},
	"Personajoker": {
		"id":   uint(669569),
		"form": uint(0),
		"type": "dark",
	},
	"Personajoker-Zen": {
		"id":   uint(669569),
		"form": uint(1),
		"type": "dark",
	},
	"Snkterry": {
		"id":   uint(669570),
		"form": uint(0),
		"type": "fighting",
	},
	"Banjo Kazooie": {
		"id":   uint(669571),
		"form": uint(0),
		"type": "ground",
	},
	"Min Min": {
		"id":   uint(669572),
		"form": uint(0),
		"type": "food",
	},
	"Minecraftsteve": {
		"id":   uint(669573),
		"form": uint(0),
		"type": "normal",
	},
	"Sephiroth": {
		"id":   uint(669574),
		"form": uint(0),
		"type": "steel",
	},
	"Sephiroth-Zen": {
		"id":   uint(669574),
		"form": uint(1),
		"type": "divine",
	},
	"Safersephiroth": {
		"id":   uint(669575),
		"form": uint(0),
		"type": "divine",
	},
	"Herobrine": {
		"id":   uint(669576),
		"form": uint(0),
		"type": "normal",
	},
	"Mkscorpion": {
		"id":   uint(669577),
		"form": uint(0),
		"type": "zombie",
	},
	"Mksubzero": {
		"id":   uint(669578),
		"form": uint(0),
		"type": "dark",
	},
	"Mileena": {
		"id":   uint(669579),
		"form": uint(0),
		"type": "dark",
	},
	"Mkskarlet": {
		"id":   uint(669580),
		"form": uint(0),
		"type": "blood",
	},
	"Edash": {
		"id":   uint(669581),
		"form": uint(0),
		"type": "fighting",
	},
	"Krampusird": {
		"id":   uint(669582),
		"form": uint(0),
		"type": "ghost",
	},
	"Santa Cirno": {
		"id":   uint(669583),
		"form": uint(0),
		"type": "fairy",
	},
	"Santa Kisume": {
		"id":   uint(669584),
		"form": uint(0),
		"type": "fire",
	},
	"Santamystia": {
		"id":   uint(669585),
		"form": uint(0),
		"type": "ice",
	},
	"Santaparsee": {
		"id":   uint(669586),
		"form": uint(0),
		"type": "heart",
	},
	"Santarumia": {
		"id":   uint(669587),
		"form": uint(0),
		"type": "dark",
	},
	"Santawriggle": {
		"id":   uint(669588),
		"form": uint(0),
		"type": "bug",
	},
	"Beach Alice": {
		"id":   uint(669589),
		"form": uint(0),
		"type": "water",
	},
	"Isami": {
		"id":   uint(669590),
		"form": uint(0),
		"type": "water",
	},
	"Skyrimzard": {
		"id":   uint(669591),
		"form": uint(0),
		"type": "fire",
	},
	"Lozmidna": {
		"id":   uint(669592),
		"form": uint(0),
		"type": "dark",
	},
	"Poppip": {
		"id":   uint(669593),
		"form": uint(0),
		"type": "ice",
	},
	"Skipop": {
		"id":   uint(669594),
		"form": uint(0),
		"type": "ice",
	},
	"Yumsicle": {
		"id":   uint(669595),
		"form": uint(0),
		"type": "grass",
	},
	"Yarrbull": {
		"id":   uint(669596),
		"form": uint(0),
		"type": "fighting",
	},
	"Swashbull": {
		"id":   uint(669597),
		"form": uint(0),
		"type": "fighting",
	},
	"Goldeneye007": {
		"id":   uint(669598),
		"form": uint(0),
		"type": "normal",
	},
	"Santa Joon": {
		"id":   uint(669599),
		"form": uint(0),
		"type": "fighting",
	},
	"Santa Shion": {
		"id":   uint(669600),
		"form": uint(0),
		"type": "ghost",
	},
	"Meduka": {
		"id":   uint(669601),
		"form": uint(0),
		"type": "magic",
	},
	"Mmsayaka": {
		"id":   uint(669602),
		"form": uint(0),
		"type": "magic",
	},
	"Mmmami": {
		"id":   uint(669603),
		"form": uint(0),
		"type": "magic",
	},
	"Mmkyoko": {
		"id":   uint(669604),
		"form": uint(0),
		"type": "magic",
	},
	"Moemura": {
		"id":   uint(669605),
		"form": uint(0),
		"type": "magic",
	},
	"Mmhomura": {
		"id":   uint(669606),
		"form": uint(0),
		"type": "magic",
	},
	"Kage No Mushi": {
		"id":   uint(669607),
		"form": uint(0),
		"type": "bug",
	},
	"Layla": {
		"id":   uint(669608),
		"form": uint(0),
		"type": "ghost",
	},
	"Beachutsuho": {
		"id":   uint(669609),
		"form": uint(0),
		"type": "nuclear",
	},
	"Beachshikieiki": {
		"id":   uint(669610),
		"form": uint(0),
		"type": "divine",
	},
	"Barbeon": {
		"id":   uint(669611),
		"form": uint(0),
		"type": "poison",
	},
	"Senseion": {
		"id":   uint(669612),
		"form": uint(0),
		"type": "fighting",
	},
	"Doomimp": {
		"id":   uint(669613),
		"form": uint(0),
		"type": "chaos",
	},
	"Doompinky": {
		"id":   uint(669614),
		"form": uint(0),
		"type": "chaos",
	},
	"Dlostsoul": {
		"id":   uint(669615),
		"form": uint(0),
		"type": "ghost",
	},
	"Betalostsoul": {
		"id":   uint(669616),
		"form": uint(0),
		"type": "ghost",
	},
	"Mancubus": {
		"id":   uint(669617),
		"form": uint(0),
		"type": "chaos",
	},
	"Doomrevenant": {
		"id":   uint(669618),
		"form": uint(0),
		"type": "chaos",
	},
	"Doomhellknight": {
		"id":   uint(669619),
		"form": uint(0),
		"type": "chaos",
	},
	"Baron Of Hell": {
		"id":   uint(669620),
		"form": uint(0),
		"type": "chaos",
	},
	"Darchvile": {
		"id":   uint(669621),
		"form": uint(0),
		"type": "chaos",
	},
	"Pain Elemental": {
		"id":   uint(669622),
		"form": uint(0),
		"type": "chaos",
	},
	"Dcyberdemon": {
		"id":   uint(669623),
		"form": uint(0),
		"type": "chaos",
	},
	"Arachnotron": {
		"id":   uint(669624),
		"form": uint(0),
		"type": "chaos",
	},
	"Dspiderdemon": {
		"id":   uint(669625),
		"form": uint(0),
		"type": "chaos",
	},
	"Sfkrystal": {
		"id":   uint(669626),
		"form": uint(0),
		"type": "normal",
	},
	"Sitryvern": {
		"id":   uint(669627),
		"form": uint(0),
		"type": "dragon",
	},
	"Sitragon": {
		"id":   uint(669628),
		"form": uint(0),
		"type": "dragon",
	},
	"Dracelium": {
		"id":   uint(669629),
		"form": uint(0),
		"type": "dragon",
	},
	"Mtoad": {
		"id":   uint(669630),
		"form": uint(0),
		"type": "normal",
	},
	"Toadette": {
		"id":   uint(669631),
		"form": uint(0),
		"type": "normal",
	},
	"Cmtorterra": {
		"id":   uint(669632),
		"form": uint(0),
		"type": "grass",
	},
	"Tailsprower": {
		"id":   uint(669633),
		"form": uint(0),
		"type": "normal",
	},
	"Knucklesechidna": {
		"id":   uint(669634),
		"form": uint(0),
		"type": "ground",
	},
	"Shadowhedgehog": {
		"id":   uint(669635),
		"form": uint(0),
		"type": "dark",
	},
	"Dr Eggman": {
		"id":   uint(669636),
		"form": uint(0),
		"type": "normal",
	},
	"Amyrose": {
		"id":   uint(669637),
		"form": uint(0),
		"type": "normal",
	},
	"Rougebat": {
		"id":   uint(669638),
		"form": uint(0),
		"type": "dark",
	},
	"Chaossonic": {
		"id":   uint(669639),
		"form": uint(0),
		"type": "water",
	},
	"Creamrabbit": {
		"id":   uint(669640),
		"form": uint(0),
		"type": "normal",
	},
	"E102": {
		"id":   uint(669641),
		"form": uint(0),
		"type": "tech",
	},
	"Emerl": {
		"id":   uint(669642),
		"form": uint(0),
		"type": "tech",
	},
	"Dmagician Girl": {
		"id":   uint(669643),
		"form": uint(0),
		"type": "heart",
	},
	"Dark Magician": {
		"id":   uint(669644),
		"form": uint(0),
		"type": "dark",
	},
	"Redeyes Blkdragon": {
		"id":   uint(669645),
		"form": uint(0),
		"type": "dragon",
	},
	"Sawcubus": {
		"id":   uint(669646),
		"form": uint(0),
		"type": "chaos",
	},
	"Dwatcher": {
		"id":   uint(669647),
		"form": uint(0),
		"type": "chaos",
	},
	"Doomguy": {
		"id":   uint(669648),
		"form": uint(0),
		"type": "fighting",
	},
	"Hufennekin": {
		"id":   uint(669649),
		"form": uint(0),
		"type": "fire",
	},
	"Hubraixen": {
		"id":   uint(669650),
		"form": uint(0),
		"type": "fire",
	},
	"Hudelphox": {
		"id":   uint(669651),
		"form": uint(0),
		"type": "fire",
	},
	"Big The Cat": {
		"id":   uint(669652),
		"form": uint(0),
		"type": "normal",
	},
	"Borosu": {
		"id":   uint(669653),
		"form": uint(0),
		"type": "cosmic",
	},
	"Borosu-Zen": {
		"id":   uint(669653),
		"form": uint(1),
		"type": "cosmic",
	},
	"Dawnpoke": {
		"id":   uint(669654),
		"form": uint(0),
		"type": "water",
	},
	"Birthdayraticate": {
		"id":   uint(669655),
		"form": uint(0),
		"type": "normal",
	},
	"Mumenrider": {
		"id":   uint(669656),
		"form": uint(0),
		"type": "normal",
	},
	"Crabblante": {
		"id":   uint(669657),
		"form": uint(0),
		"type": "water",
	},
	"Pripriprisoner": {
		"id":   uint(669658),
		"form": uint(0),
		"type": "normal",
	},
	"Metabat": {
		"id":   uint(669659),
		"form": uint(0),
		"type": "normal",
	},
	"Manvaccine": {
		"id":   uint(669660),
		"form": uint(0),
		"type": "psychic",
	},
	"Genosu": {
		"id":   uint(669661),
		"form": uint(0),
		"type": "tech",
	},
	"Tanktopmastah": {
		"id":   uint(669662),
		"form": uint(0),
		"type": "fighting",
	},
	"Hammerheaddo": {
		"id":   uint(669663),
		"form": uint(0),
		"type": "normal",
	},
	"Watchdoggoman": {
		"id":   uint(669664),
		"form": uint(0),
		"type": "normal",
	},
	"Childoemperor": {
		"id":   uint(669665),
		"form": uint(0),
		"type": "normal",
	},
	"Shy Guy": {
		"id":   uint(669666),
		"form": uint(0),
		"type": "normal",
	},
	"Shy Fly Guy": {
		"id":   uint(669667),
		"form": uint(0),
		"type": "normal",
	},
	"Shyguyghost": {
		"id":   uint(669668),
		"form": uint(0),
		"type": "normal",
	},
	"Shy Gal": {
		"id":   uint(669669),
		"form": uint(0),
		"type": "normal",
	},
	"Koopa Troopa": {
		"id":   uint(669670),
		"form": uint(0),
		"type": "normal",
	},
	"Koopa Paratroopa": {
		"id":   uint(669671),
		"form": uint(0),
		"type": "normal",
	},
	"Bobomb": {
		"id":   uint(669672),
		"form": uint(0),
		"type": "normal",
	},
	"Mawvenus": {
		"id":   uint(669673),
		"form": uint(0),
		"type": "grass",
	},
	"Dry Bones": {
		"id":   uint(669674),
		"form": uint(0),
		"type": "bone",
	},
	"Opmmetalknight": {
		"id":   uint(669675),
		"form": uint(0),
		"type": "tech",
	},
	"Sonikku": {
		"id":   uint(669676),
		"form": uint(0),
		"type": "normal",
	},
	"Skull Kid": {
		"id":   uint(669677),
		"form": uint(0),
		"type": "dark",
	},
	"Gogengar": {
		"id":   uint(669678),
		"form": uint(0),
		"type": "ghost",
	},
	"Fedorakirlia": {
		"id":   uint(669679),
		"form": uint(0),
		"type": "psychic",
	},
	"Excite Bike": {
		"id":   uint(669680),
		"form": uint(0),
		"type": "cyber",
	},
	"Tomba": {
		"id":   uint(669681),
		"form": uint(0),
		"type": "rock",
	},
	"Ben Drowned": {
		"id":   uint(669682),
		"form": uint(0),
		"type": "ghost",
	},
	"Bandana Dee": {
		"id":   uint(669683),
		"form": uint(0),
		"type": "normal",
	},
	"Gsisaac": {
		"id":   uint(669684),
		"form": uint(0),
		"type": "ground",
	},
	"Waluigi": {
		"id":   uint(669685),
		"form": uint(0),
		"type": "poison",
	},
	"Rayman": {
		"id":   uint(669686),
		"form": uint(0),
		"type": "normal",
	},
	"Smshsandbag": {
		"id":   uint(669687),
		"form": uint(0),
		"type": "normal",
	},
	"Chibi Robo": {
		"id":   uint(669688),
		"form": uint(0),
		"type": "tech",
	},
	"Bomberman": {
		"id":   uint(669689),
		"form": uint(0),
		"type": "normal",
	},
	"Ffblackmage": {
		"id":   uint(669690),
		"form": uint(0),
		"type": "magic",
	},
	"Khsora": {
		"id":   uint(669691),
		"form": uint(0),
		"type": "light",
	},
	"Lloydirving": {
		"id":   uint(669692),
		"form": uint(0),
		"type": "steel",
	},
	"Mach Rider": {
		"id":   uint(669693),
		"form": uint(0),
		"type": "tech",
	},
	"Urban Champion": {
		"id":   uint(669694),
		"form": uint(0),
		"type": "fighting",
	},
	"Dnryuk": {
		"id":   uint(669695),
		"form": uint(0),
		"type": "ghost",
	},
	"Pdplip": {
		"id":   uint(669696),
		"form": uint(0),
		"type": "fairy",
	},
	"Michael Jackson": {
		"id":   uint(669697),
		"form": uint(0),
		"type": "fairy",
	},
	"Mksheeva": {
		"id":   uint(669698),
		"form": uint(0),
		"type": "fighting",
	},
	"Mkgoro": {
		"id":   uint(669699),
		"form": uint(0),
		"type": "fighting",
	},
	"Morshu": {
		"id":   uint(669700),
		"form": uint(0),
		"type": "normal",
	},
	"Hosmoke": {
		"id":   uint(669701),
		"form": uint(0),
		"type": "fire",
	},
	"Takamaru": {
		"id":   uint(669702),
		"form": uint(0),
		"type": "normal",
	},
	"Ayumi Tachibana": {
		"id":   uint(669703),
		"form": uint(0),
		"type": "normal",
	},
	"Nintendog": {
		"id":   uint(669704),
		"form": uint(0),
		"type": "normal",
	},
	"Loztetra": {
		"id":   uint(669705),
		"form": uint(0),
		"type": "water",
	},
	"Dixie Kong": {
		"id":   uint(669706),
		"form": uint(0),
		"type": "normal",
	},
	"Balloonfighter": {
		"id":   uint(669707),
		"form": uint(0),
		"type": "rubber",
	},
	"Johnny Cage": {
		"id":   uint(669708),
		"form": uint(0),
		"type": "fighting",
	},
	"Sonya Blade": {
		"id":   uint(669709),
		"form": uint(0),
		"type": "fighting",
	},
	"Shang Tsung": {
		"id":   uint(669710),
		"form": uint(0),
		"type": "magic",
	},
	"Mkraiden": {
		"id":   uint(669711),
		"form": uint(0),
		"type": "electric",
	},
	"Mkreptile": {
		"id":   uint(669712),
		"form": uint(0),
		"type": "poison",
	},
	"Mkkano": {
		"id":   uint(669713),
		"form": uint(0),
		"type": "dark",
	},
	"Liu Kang": {
		"id":   uint(669714),
		"form": uint(0),
		"type": "fighting",
	},
	"Baraka": {
		"id":   uint(669715),
		"form": uint(0),
		"type": "dark",
	},
	"Sindel": {
		"id":   uint(669716),
		"form": uint(0),
		"type": "dark",
	},
	"Fftifa": {
		"id":   uint(669717),
		"form": uint(0),
		"type": "normal",
	},
	"Ffbarret": {
		"id":   uint(669718),
		"form": uint(0),
		"type": "normal",
	},
	"Bkgruntilda": {
		"id":   uint(669719),
		"form": uint(0),
		"type": "magic",
	},
	"Metroid": {
		"id":   uint(669720),
		"form": uint(0),
		"type": "virus",
	},
	"Ffaerith": {
		"id":   uint(669721),
		"form": uint(0),
		"type": "normal",
	},
	"Red Xiii": {
		"id":   uint(669722),
		"form": uint(0),
		"type": "fire",
	},
	"Ffyuffie": {
		"id":   uint(669723),
		"form": uint(0),
		"type": "dark",
	},
	"Cait Sith": {
		"id":   uint(669724),
		"form": uint(0),
		"type": "normal",
	},
	"Ffzack": {
		"id":   uint(669725),
		"form": uint(0),
		"type": "steel",
	},
	"Cluclububbles": {
		"id":   uint(669726),
		"form": uint(0),
		"type": "rubber",
	},
	"Heihachi": {
		"id":   uint(669727),
		"form": uint(0),
		"type": "fighting",
	},
	"Chorus Kids": {
		"id":   uint(669728),
		"form": uint(0),
		"type": "sound",
	},
	"Ctcrono": {
		"id":   uint(669729),
		"form": uint(0),
		"type": "steel",
	},
	"Soma Cruz": {
		"id":   uint(669730),
		"form": uint(0),
		"type": "magic",
	},
	"Alucard": {
		"id":   uint(669731),
		"form": uint(0),
		"type": "dark",
	},
	"Cvdracula": {
		"id":   uint(669732),
		"form": uint(0),
		"type": "blood",
	},
	"Scivy": {
		"id":   uint(669733),
		"form": uint(0),
		"type": "dark",
	},
	"Master Chief": {
		"id":   uint(669734),
		"form": uint(0),
		"type": "fighting",
	},
	"Arbiter": {
		"id":   uint(669735),
		"form": uint(0),
		"type": "cosmic",
	},
	"Halogrunt": {
		"id":   uint(669736),
		"form": uint(0),
		"type": "cosmic",
	},
	"Haloelite": {
		"id":   uint(669737),
		"form": uint(0),
		"type": "cosmic",
	},
	"Halobrute": {
		"id":   uint(669738),
		"form": uint(0),
		"type": "cosmic",
	},
	"Halojackal": {
		"id":   uint(669739),
		"form": uint(0),
		"type": "cosmic",
	},
	"Haloskirmisher": {
		"id":   uint(669740),
		"form": uint(0),
		"type": "cosmic",
	},
	"Halohunter": {
		"id":   uint(669741),
		"form": uint(0),
		"type": "cosmic",
	},
	"Halodrone": {
		"id":   uint(669742),
		"form": uint(0),
		"type": "cosmic",
	},
	"Detna": {
		"id":   uint(669743),
		"form": uint(0),
		"type": "dark",
	},
	"Kosmos": {
		"id":   uint(669744),
		"form": uint(0),
		"type": "tech",
	},
	"Amaterasu": {
		"id":   uint(669745),
		"form": uint(0),
		"type": "fire",
	},
	"Pheonix Wright": {
		"id":   uint(669746),
		"form": uint(0),
		"type": "normal",
	},
	"Garcynthia": {
		"id":   uint(669747),
		"form": uint(0),
		"type": "dragon",
	},
	"Elesatrika": {
		"id":   uint(669748),
		"form": uint(0),
		"type": "electric",
	},
	"Maytias": {
		"id":   uint(669749),
		"form": uint(0),
		"type": "dragon",
	},
	"Hildaroark": {
		"id":   uint(669750),
		"form": uint(0),
		"type": "dark",
	},
	"Beach Patchi": {
		"id":   uint(669751),
		"form": uint(0),
		"type": "magic",
	},
	"Beachsakuya": {
		"id":   uint(669752),
		"form": uint(0),
		"type": "time",
	},
	"Beach Reimu": {
		"id":   uint(669753),
		"form": uint(0),
		"type": "normal",
	},
	"Umigeorge": {
		"id":   uint(669754),
		"form": uint(0),
		"type": "normal",
	},
	"Umijessica": {
		"id":   uint(669755),
		"form": uint(0),
		"type": "normal",
	},
	"Rainbow Mika": {
		"id":   uint(669756),
		"form": uint(0),
		"type": "fighting",
	},
	"Umikrauss": {
		"id":   uint(669757),
		"form": uint(0),
		"type": "normal",
	},
	"Uminatsuhi": {
		"id":   uint(669758),
		"form": uint(0),
		"type": "normal",
	},
	"Umihideyoshi": {
		"id":   uint(669759),
		"form": uint(0),
		"type": "normal",
	},
	"Umieva": {
		"id":   uint(669760),
		"form": uint(0),
		"type": "normal",
	},
	"Umirudolf": {
		"id":   uint(669761),
		"form": uint(0),
		"type": "normal",
	},
	"Umikyrie": {
		"id":   uint(669762),
		"form": uint(0),
		"type": "normal",
	},
	"Umirosa": {
		"id":   uint(669763),
		"form": uint(0),
		"type": "normal",
	},
	"Umimaria": {
		"id":   uint(669764),
		"form": uint(0),
		"type": "normal",
	},
	"Umigenji": {
		"id":   uint(669765),
		"form": uint(0),
		"type": "normal",
	},
	"Umikumasawa": {
		"id":   uint(669766),
		"form": uint(0),
		"type": "normal",
	},
	"Fury Bowser": {
		"id":   uint(669767),
		"form": uint(0),
		"type": "dragon",
	},
	"Bowser Statue": {
		"id":   uint(669768),
		"form": uint(0),
		"type": "fire",
	},
	"Mechabowserstatue": {
		"id":   uint(669769),
		"form": uint(0),
		"type": "tech",
	},
	"Bowser Koobala": {
		"id":   uint(669770),
		"form": uint(0),
		"type": "fire",
	},
	"Umishannon": {
		"id":   uint(669771),
		"form": uint(0),
		"type": "normal",
	},
	"Umikanon": {
		"id":   uint(669772),
		"form": uint(0),
		"type": "normal",
	},
	"Birdo": {
		"id":   uint(669773),
		"form": uint(0),
		"type": "normal",
	},
	"Goombella": {
		"id":   uint(669774),
		"form": uint(0),
		"type": "normal",
	},
	"Metal Mario": {
		"id":   uint(669775),
		"form": uint(0),
		"type": "steel",
	},
	"Gooigi": {
		"id":   uint(669776),
		"form": uint(0),
		"type": "poison",
	},
	"Darkskylanderbowser": {
		"id":   uint(669777),
		"form": uint(0),
		"type": "plastic",
	},
	"Akira Yuki": {
		"id":   uint(669778),
		"form": uint(0),
		"type": "fighting",
	},
	"Andross": {
		"id":   uint(669779),
		"form": uint(0),
		"type": "tech",
	},
	"Arcade Bunny": {
		"id":   uint(669780),
		"form": uint(0),
		"type": "normal",
	},
	"Wwashley": {
		"id":   uint(669781),
		"form": uint(0),
		"type": "magic",
	},
	"Barbarabat": {
		"id":   uint(669782),
		"form": uint(0),
		"type": "sound",
	},
	"Feblackknight": {
		"id":   uint(669783),
		"form": uint(0),
		"type": "steel",
	},
	"Chainchomp": {
		"id":   uint(669784),
		"form": uint(0),
		"type": "steel",
	},
	"Chef Kawasaki": {
		"id":   uint(669785),
		"form": uint(0),
		"type": "normal",
	},
	"Lenneth": {
		"id":   uint(669786),
		"form": uint(0),
		"type": "divine",
	},
	"Phantom Ganon": {
		"id":   uint(669787),
		"form": uint(0),
		"type": "dark",
	},
	"King Bobomb": {
		"id":   uint(669788),
		"form": uint(0),
		"type": "steel",
	},
	"Gold Mario": {
		"id":   uint(669789),
		"form": uint(0),
		"type": "steel",
	},
	"Ice Mario": {
		"id":   uint(669790),
		"form": uint(0),
		"type": "ice",
	},
	"Dsmorrigan": {
		"id":   uint(669791),
		"form": uint(0),
		"type": "heart",
	},
	"Hammer Bro": {
		"id":   uint(669792),
		"form": uint(0),
		"type": "fighting",
	},
	"Dwdevil": {
		"id":   uint(669793),
		"form": uint(0),
		"type": "dark",
	},
	"Dillon": {
		"id":   uint(669794),
		"form": uint(0),
		"type": "ground",
	},
	"Lozganon": {
		"id":   uint(669795),
		"form": uint(0),
		"type": "dark",
	},
	"Dark Link": {
		"id":   uint(669796),
		"form": uint(0),
		"type": "fairy",
	},
	"Dr Kawashima": {
		"id":   uint(669797),
		"form": uint(0),
		"type": "normal",
	},
	"Ray Mk Iii": {
		"id":   uint(669798),
		"form": uint(0),
		"type": "tech",
	},
	"Nier2B": {
		"id":   uint(669799),
		"form": uint(0),
		"type": "tech",
	},
	"Nier9S": {
		"id":   uint(669800),
		"form": uint(0),
		"type": "tech",
	},
	"Niera2": {
		"id":   uint(669801),
		"form": uint(0),
		"type": "tech",
	},
	"Dr Wright": {
		"id":   uint(669802),
		"form": uint(0),
		"type": "normal",
	},
	"Elec Man": {
		"id":   uint(669803),
		"form": uint(0),
		"type": "electric",
	},
	"Color Tvgame 15": {
		"id":   uint(669804),
		"form": uint(0),
		"type": "cyber",
	},
	"Blazecat": {
		"id":   uint(669805),
		"form": uint(0),
		"type": "fire",
	},
	"Bullet Bill": {
		"id":   uint(669806),
		"form": uint(0),
		"type": "steel",
	},
	"Kung Fu Man": {
		"id":   uint(669807),
		"form": uint(0),
		"type": "fighting",
	},
	"Kung Fu Girl": {
		"id":   uint(669808),
		"form": uint(0),
		"type": "fighting",
	},
	"M Bison": {
		"id":   uint(669809),
		"form": uint(0),
		"type": "dark",
	},
	"Dhalsim": {
		"id":   uint(669810),
		"form": uint(0),
		"type": "fighting",
	},
	"Sfguile": {
		"id":   uint(669811),
		"form": uint(0),
		"type": "fighting",
	},
	"E Honda": {
		"id":   uint(669812),
		"form": uint(0),
		"type": "fighting",
	},
	"Zangief": {
		"id":   uint(669813),
		"form": uint(0),
		"type": "fighting",
	},
	"Blanka": {
		"id":   uint(669814),
		"form": uint(0),
		"type": "electric",
	},
	"C Viper": {
		"id":   uint(669815),
		"form": uint(0),
		"type": "fighting",
	},
	"Crwutemate": {
		"id":   uint(669816),
		"form": uint(0),
		"type": "cosmic",
	},
	"Coommate": {
		"id":   uint(669817),
		"form": uint(0),
		"type": "cosmic",
	},
	"Penguira": {
		"id":   uint(669818),
		"form": uint(0),
		"type": "ice",
	},
	"Lugiel": {
		"id":   uint(669819),
		"form": uint(0),
		"type": "psychic",
	},
	"Dqdragonlord": {
		"id":   uint(669820),
		"form": uint(0),
		"type": "dark",
	},
	"Dqdragonlord-Zen": {
		"id":   uint(669820),
		"form": uint(1),
		"type": "dark",
	},
	"Dqmalroth": {
		"id":   uint(669821),
		"form": uint(0),
		"type": "dark",
	},
	"Dqzoma": {
		"id":   uint(669822),
		"form": uint(0),
		"type": "dark",
	},
	"Bubble Slime": {
		"id":   uint(669823),
		"form": uint(0),
		"type": "poison",
	},
	"Heal Slimee": {
		"id":   uint(669824),
		"form": uint(0),
		"type": "normal",
	},
	"Liquidmetalslime": {
		"id":   uint(669825),
		"form": uint(0),
		"type": "poison",
	},
	"Cure Slimee": {
		"id":   uint(669826),
		"form": uint(0),
		"type": "normal",
	},
	"Seaslimee": {
		"id":   uint(669827),
		"form": uint(0),
		"type": "water",
	},
	"Wackqualitydoge": {
		"id":   uint(669828),
		"form": uint(0),
		"type": "normal",
	},
	"Wackmuscleman": {
		"id":   uint(669829),
		"form": uint(0),
		"type": "normal",
	},
	"Wackchad": {
		"id":   uint(669830),
		"form": uint(0),
		"type": "normal",
	},
	"Wackvirgin": {
		"id":   uint(669831),
		"form": uint(0),
		"type": "normal",
	},
	"Sprite007": {
		"id":   uint(669832),
		"form": uint(0),
		"type": "normal",
	},
	"Youkawaii": {
		"id":   uint(669833),
		"form": uint(0),
		"type": "heart",
	},
	"Xbpyra": {
		"id":   uint(669834),
		"form": uint(0),
		"type": "fire",
	},
	"Xbrex": {
		"id":   uint(669835),
		"form": uint(0),
		"type": "normal",
	},
	"Xbmythra": {
		"id":   uint(669836),
		"form": uint(0),
		"type": "light",
	},
	"Master Hand": {
		"id":   uint(669837),
		"form": uint(0),
		"type": "normal",
	},
	"Fgonobunaga": {
		"id":   uint(669838),
		"form": uint(0),
		"type": "chaos",
	},
	"Fgobenienma": {
		"id":   uint(669839),
		"form": uint(0),
		"type": "food",
	},
	"Fgohector": {
		"id":   uint(669840),
		"form": uint(0),
		"type": "fighting",
	},
	"Ushiwakamaru": {
		"id":   uint(669841),
		"form": uint(0),
		"type": "wind",
	},
	"Scheherazade": {
		"id":   uint(669842),
		"form": uint(0),
		"type": "paper",
	},
	"Fgohimiko": {
		"id":   uint(669843),
		"form": uint(0),
		"type": "light",
	},
	"Fgoraikou": {
		"id":   uint(669844),
		"form": uint(0),
		"type": "chaos",
	},
	"Fgolancelot": {
		"id":   uint(669845),
		"form": uint(0),
		"type": "fairy",
	},
	"Fgodavid": {
		"id":   uint(669846),
		"form": uint(0),
		"type": "divine",
	},
	"Kiyolancer": {
		"id":   uint(669847),
		"form": uint(0),
		"type": "water",
	},
	"Mordrider": {
		"id":   uint(669848),
		"form": uint(0),
		"type": "water",
	},
	"Helena Blavatsky": {
		"id":   uint(669849),
		"form": uint(0),
		"type": "magic",
	},
	"Shutendoji": {
		"id":   uint(669850),
		"form": uint(0),
		"type": "poison",
	},
	"Dariuserker": {
		"id":   uint(669851),
		"form": uint(0),
		"type": "chaos",
	},
	"Avengergorgon": {
		"id":   uint(669852),
		"form": uint(0),
		"type": "blood",
	},
	"Fgobb": {
		"id":   uint(669853),
		"form": uint(0),
		"type": "cyber",
	},
	"Negarman": {
		"id":   uint(669854),
		"form": uint(0),
		"type": "chaos",
	},
	"Fgotiamat": {
		"id":   uint(669855),
		"form": uint(0),
		"type": "dragon",
	},
	"Irisviel": {
		"id":   uint(669856),
		"form": uint(0),
		"type": "divine",
	},
	"Emiyassin": {
		"id":   uint(669857),
		"form": uint(0),
		"type": "steel",
	},
	"Muramiya": {
		"id":   uint(669858),
		"form": uint(0),
		"type": "steel",
	},
	"Jaguar Taiga": {
		"id":   uint(669859),
		"form": uint(0),
		"type": "normal",
	},
	"Emiyalter": {
		"id":   uint(669860),
		"form": uint(0),
		"type": "steel",
	},
	"Sima Reines": {
		"id":   uint(669861),
		"form": uint(0),
		"type": "steel",
	},
	"Saberlion": {
		"id":   uint(669862),
		"form": uint(0),
		"type": "normal",
	},
	"Astraea": {
		"id":   uint(669863),
		"form": uint(0),
		"type": "divine",
	},
	"Richard I": {
		"id":   uint(669864),
		"form": uint(0),
		"type": "normal",
	},
	"Sfalcides": {
		"id":   uint(669865),
		"form": uint(0),
		"type": "fighting",
	},
	"Lolidusa": {
		"id":   uint(669866),
		"form": uint(0),
		"type": "divine",
	},
	"Fperseus": {
		"id":   uint(669867),
		"form": uint(0),
		"type": "divine",
	},
	"Alexandre Dumas": {
		"id":   uint(669868),
		"form": uint(0),
		"type": "paper",
	},
	"J Edgar Hoover": {
		"id":   uint(669869),
		"form": uint(0),
		"type": "dark",
	},
	"Kijyo Koyo": {
		"id":   uint(669870),
		"form": uint(0),
		"type": "dragon",
	},
	"Twilight Sparkle": {
		"id":   uint(669871),
		"form": uint(0),
		"type": "magic",
	},
	"Mlpapplejack": {
		"id":   uint(669872),
		"form": uint(0),
		"type": "ground",
	},
	"Mlprainbowdash": {
		"id":   uint(669873),
		"form": uint(0),
		"type": "flying",
	},
	"Mlprarity": {
		"id":   uint(669874),
		"form": uint(0),
		"type": "crystal",
	},
	"Mlppinkiepie": {
		"id":   uint(669875),
		"form": uint(0),
		"type": "fairy",
	},
	"Mlpfluttershy": {
		"id":   uint(669876),
		"form": uint(0),
		"type": "heart",
	},
	"Mlpspike": {
		"id":   uint(669877),
		"form": uint(0),
		"type": "dragon",
	},
	"Spowps": {
		"id":   uint(669878),
		"form": uint(0),
		"type": "dragon",
	},
	"Hgrowlithe": {
		"id":   uint(669879),
		"form": uint(0),
		"type": "fire",
	},
	"Harcanine": {
		"id":   uint(669880),
		"form": uint(0),
		"type": "fire",
	},
	"Hbraviary": {
		"id":   uint(669881),
		"form": uint(0),
		"type": "psychic",
	},
	"Btakua": {
		"id":   uint(669882),
		"form": uint(0),
		"type": "light",
	},
	"Takanuva": {
		"id":   uint(669883),
		"form": uint(0),
		"type": "light",
	},
	"Onua": {
		"id":   uint(669884),
		"form": uint(0),
		"type": "ground",
	},
	"Kopaka": {
		"id":   uint(669885),
		"form": uint(0),
		"type": "ice",
	},
	"Pohatu": {
		"id":   uint(669886),
		"form": uint(0),
		"type": "rock",
	},
	"Lewa": {
		"id":   uint(669887),
		"form": uint(0),
		"type": "wind",
	},
	"Gali": {
		"id":   uint(669888),
		"form": uint(0),
		"type": "water",
	},
	"Tahu": {
		"id":   uint(669889),
		"form": uint(0),
		"type": "fire",
	},
	"God": {
		"id":   uint(669890),
		"form": uint(0),
		"type": "divine",
	},
}

var moveData = map[string]map[string]interface{}{
	"Acupressure": {
		"id":             uint(367),
		"type":           "normal",
		"classification": uint(0),
	},
	"After You": {
		"id":             uint(495),
		"type":           "normal",
		"classification": uint(0),
	},
	"Assist": {
		"id":             uint(274),
		"type":           "normal",
		"classification": uint(0),
	},
	"Attract": {
		"id":             uint(213),
		"type":           "normal",
		"classification": uint(0),
	},
	"Barrage": {
		"id":             uint(140),
		"type":           "normal",
		"classification": uint(1),
	},
	"Baton Pass": {
		"id":             uint(226),
		"type":           "normal",
		"classification": uint(0),
	},
	"Belly Drum": {
		"id":             uint(187),
		"type":           "normal",
		"classification": uint(0),
	},
	"Bestow": {
		"id":             uint(516),
		"type":           "normal",
		"classification": uint(0),
	},
	"Bide": {
		"id":             uint(117),
		"type":           "normal",
		"classification": uint(1),
	},
	"Bind": {
		"id":             uint(20),
		"type":           "normal",
		"classification": uint(1),
	},
	"Block": {
		"id":             uint(335),
		"type":           "normal",
		"classification": uint(0),
	},
	"Body Slam": {
		"id":             uint(34),
		"type":           "normal",
		"classification": uint(1),
	},
	"Boomburst": {
		"id":             uint(586),
		"type":           "normal",
		"classification": uint(2),
	},
	"Camouflage": {
		"id":             uint(293),
		"type":           "normal",
		"classification": uint(0),
	},
	"Captivate": {
		"id":             uint(445),
		"type":           "normal",
		"classification": uint(0),
	},
	"Celebrate": {
		"id":             uint(606),
		"type":           "normal",
		"classification": uint(0),
	},
	"Chip Away": {
		"id":             uint(498),
		"type":           "normal",
		"classification": uint(1),
	},
	"Comet Punch": {
		"id":             uint(4),
		"type":           "normal",
		"classification": uint(1),
	},
	"Confide": {
		"id":             uint(590),
		"type":           "normal",
		"classification": uint(0),
	},
	"Constrict": {
		"id":             uint(132),
		"type":           "normal",
		"classification": uint(1),
	},
	"Conversion": {
		"id":             uint(160),
		"type":           "normal",
		"classification": uint(0),
	},
	"Conversion uint(2": {
		"id":             uint(176),
		"type":           "normal",
		"classification": uint(0),
	},
	"Copycat": {
		"id":             uint(383),
		"type":           "normal",
		"classification": uint(0),
	},
	"Covet": {
		"id":             uint(343),
		"type":           "normal",
		"classification": uint(1),
	},
	"Crush Claw": {
		"id":             uint(306),
		"type":           "normal",
		"classification": uint(1),
	},
	"Crush Grip": {
		"id":             uint(462),
		"type":           "normal",
		"classification": uint(1),
	},
	"Cut": {
		"id":             uint(15),
		"type":           "normal",
		"classification": uint(1),
	},
	"Defense Curl": {
		"id":             uint(111),
		"type":           "normal",
		"classification": uint(0),
	},
	"Disable": {
		"id":             uint(50),
		"type":           "normal",
		"classification": uint(0),
	},
	"Dizzy Punch": {
		"id":             uint(146),
		"type":           "normal",
		"classification": uint(1),
	},
	"Double Hit": {
		"id":             uint(458),
		"type":           "normal",
		"classification": uint(1),
	},
	"Double Slap": {
		"id":             uint(3),
		"type":           "normal",
		"classification": uint(1),
	},
	"Double Team": {
		"id":             uint(104),
		"type":           "normal",
		"classification": uint(0),
	},
	"Double-Edge": {
		"id":             uint(38),
		"type":           "normal",
		"classification": uint(1),
	},
	"Echoed Voice": {
		"id":             uint(497),
		"type":           "normal",
		"classification": uint(2),
	},
	"Egg Bomb": {
		"id":             uint(121),
		"type":           "normal",
		"classification": uint(1),
	},
	"Encore": {
		"id":             uint(227),
		"type":           "normal",
		"classification": uint(0),
	},
	"Endeavor": {
		"id":             uint(283),
		"type":           "normal",
		"classification": uint(1),
	},
	"Endure": {
		"id":             uint(203),
		"type":           "normal",
		"classification": uint(0),
	},
	"Entrainment": {
		"id":             uint(494),
		"type":           "normal",
		"classification": uint(0),
	},
	"Explosion": {
		"id":             uint(153),
		"type":           "normal",
		"classification": uint(1),
	},
	"Extreme Speed": {
		"id":             uint(245),
		"type":           "normal",
		"classification": uint(1),
	},
	"Facade": {
		"id":             uint(263),
		"type":           "normal",
		"classification": uint(1),
	},
	"Fake Out": {
		"id":             uint(252),
		"type":           "normal",
		"classification": uint(1),
	},
	"False Swipe": {
		"id":             uint(206),
		"type":           "normal",
		"classification": uint(1),
	},
	"Feint": {
		"id":             uint(364),
		"type":           "normal",
		"classification": uint(1),
	},
	"Flail": {
		"id":             uint(175),
		"type":           "normal",
		"classification": uint(1),
	},
	"Flash": {
		"id":             uint(148),
		"type":           "normal",
		"classification": uint(0),
	},
	"Focus Energy": {
		"id":             uint(116),
		"type":           "normal",
		"classification": uint(0),
	},
	"Follow Me": {
		"id":             uint(266),
		"type":           "normal",
		"classification": uint(0),
	},
	"Foresight": {
		"id":             uint(193),
		"type":           "normal",
		"classification": uint(0),
	},
	"Frustration": {
		"id":             uint(218),
		"type":           "normal",
		"classification": uint(1),
	},
	"Fury Attack": {
		"id":             uint(31),
		"type":           "normal",
		"classification": uint(1),
	},
	"Fury Swipes": {
		"id":             uint(154),
		"type":           "normal",
		"classification": uint(1),
	},
	"Giga Impact": {
		"id":             uint(416),
		"type":           "normal",
		"classification": uint(1),
	},
	"Glare": {
		"id":             uint(137),
		"type":           "normal",
		"classification": uint(0),
	},
	"Growl": {
		"id":             uint(45),
		"type":           "normal",
		"classification": uint(0),
	},
	"Growth": {
		"id":             uint(74),
		"type":           "normal",
		"classification": uint(0),
	},
	"Guillotine": {
		"id":             uint(12),
		"type":           "normal",
		"classification": uint(1),
	},
	"Happy Hour": {
		"id":             uint(603),
		"type":           "normal",
		"classification": uint(0),
	},
	"Harden": {
		"id":             uint(106),
		"type":           "normal",
		"classification": uint(0),
	},
	"Head Charge": {
		"id":             uint(543),
		"type":           "normal",
		"classification": uint(1),
	},
	"Headbutt": {
		"id":             uint(29),
		"type":           "normal",
		"classification": uint(1),
	},
	"Heal Bell": {
		"id":             uint(215),
		"type":           "normal",
		"classification": uint(0),
	},
	"Helping Hand": {
		"id":             uint(270),
		"type":           "normal",
		"classification": uint(0),
	},
	"Hidden Power": {
		"id":             uint(237),
		"type":           "normal",
		"classification": uint(2),
	},
	"Hidden Power Fighting": {
		"id":             uint(237),
		"type":           "fighting",
		"classification": uint(2),
	},
	"Hidden Power Flying": {
		"id":             uint(237),
		"type":           "flying",
		"classification": uint(2),
	},
	"Hidden Power Poison": {
		"id":             uint(237),
		"type":           "poison",
		"classification": uint(2),
	},
	"Hidden Power Ground": {
		"id":             uint(237),
		"type":           "ground",
		"classification": uint(2),
	},
	"Hidden Power Rock": {
		"id":             uint(237),
		"type":           "rock",
		"classification": uint(2),
	},
	"Hidden Power Bug": {
		"id":             uint(237),
		"type":           "bug",
		"classification": uint(2),
	},
	"Hidden Power Ghost": {
		"id":             uint(237),
		"type":           "ghost",
		"classification": uint(2),
	},
	"Hidden Power Steel": {
		"id":             uint(237),
		"type":           "steel",
		"classification": uint(2),
	},
	"Hidden Power Fire": {
		"id":             uint(237),
		"type":           "fire",
		"classification": uint(2),
	},
	"Hidden Power Water": {
		"id":             uint(237),
		"type":           "water",
		"classification": uint(2),
	},
	"Hidden Power Grass": {
		"id":             uint(237),
		"type":           "grass",
		"classification": uint(2),
	},
	"Hidden Power Electric": {
		"id":             uint(237),
		"type":           "electric",
		"classification": uint(2),
	},
	"Hidden Power Psychic": {
		"id":             uint(237),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Hidden Power Ice": {
		"id":             uint(237),
		"type":           "ice",
		"classification": uint(2),
	},
	"Hidden Power Dragon": {
		"id":             uint(237),
		"type":           "dragon",
		"classification": uint(2),
	},
	"Hidden Power Dark": {
		"id":             uint(237),
		"type":           "dark",
		"classification": uint(2),
	},
	"Hidden Power Fairy": {
		"id":             uint(237),
		"type":           "normal",
		"classification": uint(2),
	},
	"Hold Back": {
		"id":             uint(610),
		"type":           "normal",
		"classification": uint(1),
	},
	"Hold Hands": {
		"id":             uint(607),
		"type":           "normal",
		"classification": uint(0),
	},
	"Horn Attack": {
		"id":             uint(30),
		"type":           "normal",
		"classification": uint(1),
	},
	"Horn Drill": {
		"id":             uint(32),
		"type":           "normal",
		"classification": uint(1),
	},
	"Howl": {
		"id":             uint(336),
		"type":           "normal",
		"classification": uint(0),
	},
	"Hyper Beam": {
		"id":             uint(63),
		"type":           "normal",
		"classification": uint(2),
	},
	"Hyper Fang": {
		"id":             uint(158),
		"type":           "normal",
		"classification": uint(1),
	},
	"Hyper Voice": {
		"id":             uint(304),
		"type":           "normal",
		"classification": uint(2),
	},
	"Judgment": {
		"id":             uint(449),
		"type":           "normal",
		"classification": uint(2),
	},
	"Laser Focus": {
		"id":             uint(673),
		"type":           "normal",
		"classification": uint(0),
	},
	"Last Resort": {
		"id":             uint(387),
		"type":           "normal",
		"classification": uint(1),
	},
	"Leer": {
		"id":             uint(43),
		"type":           "normal",
		"classification": uint(0),
	},
	"Lock-On": {
		"id":             uint(199),
		"type":           "normal",
		"classification": uint(0),
	},
	"Lovely Kiss": {
		"id":             uint(142),
		"type":           "normal",
		"classification": uint(0),
	},
	"Lucky Chant": {
		"id":             uint(381),
		"type":           "normal",
		"classification": uint(0),
	},
	"Me First": {
		"id":             uint(382),
		"type":           "normal",
		"classification": uint(0),
	},
	"Mean Look": {
		"id":             uint(212),
		"type":           "normal",
		"classification": uint(0),
	},
	"Mega Kick": {
		"id":             uint(25),
		"type":           "normal",
		"classification": uint(1),
	},
	"Mega Punch": {
		"id":             uint(5),
		"type":           "normal",
		"classification": uint(1),
	},
	"Metronome": {
		"id":             uint(118),
		"type":           "normal",
		"classification": uint(0),
	},
	"Milk Drink": {
		"id":             uint(208),
		"type":           "normal",
		"classification": uint(0),
	},
	"Mimic": {
		"id":             uint(102),
		"type":           "normal",
		"classification": uint(0),
	},
	"Mind Blown": {
		"id":             uint(720),
		"type":           "fire",
		"classification": uint(2),
	},
	"Mind Reader": {
		"id":             uint(170),
		"type":           "normal",
		"classification": uint(0),
	},
	"Minimize": {
		"id":             uint(107),
		"type":           "normal",
		"classification": uint(0),
	},
	"Morning Sun": {
		"id":             uint(234),
		"type":           "normal",
		"classification": uint(0),
	},
	"Multi-Attack": {
		"id":             uint(718),
		"type":           "normal",
		"classification": uint(1),
	},
	"Natural Gift": {
		"id":             uint(363),
		"type":           "normal",
		"classification": uint(1),
	},
	"Nature Power": {
		"id":             uint(267),
		"type":           "normal",
		"classification": uint(0),
	},
	"Noble Roar": {
		"id":             uint(568),
		"type":           "normal",
		"classification": uint(0),
	},
	"Odor Sleuth": {
		"id":             uint(316),
		"type":           "normal",
		"classification": uint(0),
	},
	"Pain Split": {
		"id":             uint(220),
		"type":           "normal",
		"classification": uint(0),
	},
	"Pay Day": {
		"id":             uint(6),
		"type":           "normal",
		"classification": uint(1),
	},
	"Perish Song": {
		"id":             uint(195),
		"type":           "normal",
		"classification": uint(0),
	},
	"Photon Geyser": {
		"id":             uint(722),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Plasma Fists": {
		"id":             uint(721),
		"type":           "electric",
		"classification": uint(1),
	},
	"Play Nice": {
		"id":             uint(589),
		"type":           "normal",
		"classification": uint(0),
	},
	"Pound": {
		"id":             uint(1),
		"type":           "normal",
		"classification": uint(1),
	},
	"Present": {
		"id":             uint(217),
		"type":           "normal",
		"classification": uint(1),
	},
	"Protect": {
		"id":             uint(182),
		"type":           "normal",
		"classification": uint(0),
	},
	"Psych Up": {
		"id":             uint(244),
		"type":           "normal",
		"classification": uint(0),
	},
	"Quick Attack": {
		"id":             uint(98),
		"type":           "normal",
		"classification": uint(1),
	},
	"Rage": {
		"id":             uint(99),
		"type":           "normal",
		"classification": uint(1),
	},
	"Rapid Spin": {
		"id":             uint(229),
		"type":           "normal",
		"classification": uint(1),
	},
	"Razor Wind": {
		"id":             uint(13),
		"type":           "normal",
		"classification": uint(2),
	},
	"Recover": {
		"id":             uint(105),
		"type":           "normal",
		"classification": uint(0),
	},
	"Recycle": {
		"id":             uint(278),
		"type":           "normal",
		"classification": uint(0),
	},
	"Reflect Type": {
		"id":             uint(513),
		"type":           "normal",
		"classification": uint(0),
	},
	"Refresh": {
		"id":             uint(287),
		"type":           "normal",
		"classification": uint(0),
	},
	"Relic Song": {
		"id":             uint(547),
		"type":           "normal",
		"classification": uint(2),
	},
	"Retaliate": {
		"id":             uint(514),
		"type":           "normal",
		"classification": uint(1),
	},
	"Return": {
		"id":             uint(216),
		"type":           "normal",
		"classification": uint(1),
	},
	"Revelation Dance": {
		"id":             uint(686),
		"type":           "normal",
		"classification": uint(2),
	},
	"Roar": {
		"id":             uint(46),
		"type":           "normal",
		"classification": uint(0),
	},
	"Rock Climb": {
		"id":             uint(431),
		"type":           "normal",
		"classification": uint(1),
	},
	"Round": {
		"id":             uint(496),
		"type":           "normal",
		"classification": uint(2),
	},
	"Safeguard": {
		"id":             uint(219),
		"type":           "normal",
		"classification": uint(0),
	},
	"Scary Face": {
		"id":             uint(184),
		"type":           "normal",
		"classification": uint(0),
	},
	"Scratch": {
		"id":             uint(10),
		"type":           "normal",
		"classification": uint(1),
	},
	"Screech": {
		"id":             uint(103),
		"type":           "normal",
		"classification": uint(0),
	},
	"Secret Power": {
		"id":             uint(290),
		"type":           "normal",
		"classification": uint(1),
	},
	"Self-Destruct": {
		"id":             uint(120),
		"type":           "normal",
		"classification": uint(1),
	},
	"Sharpen": {
		"id":             uint(159),
		"type":           "normal",
		"classification": uint(0),
	},
	"Shell Smash": {
		"id":             uint(504),
		"type":           "normal",
		"classification": uint(0),
	},
	"Simple Beam": {
		"id":             uint(493),
		"type":           "normal",
		"classification": uint(0),
	},
	"Sing": {
		"id":             uint(47),
		"type":           "normal",
		"classification": uint(0),
	},
	"Sketch": {
		"id":             uint(166),
		"type":           "normal",
		"classification": uint(0),
	},
	"Skull Bash": {
		"id":             uint(130),
		"type":           "normal",
		"classification": uint(1),
	},
	"Slack Off": {
		"id":             uint(303),
		"type":           "normal",
		"classification": uint(0),
	},
	"Slam": {
		"id":             uint(21),
		"type":           "normal",
		"classification": uint(1),
	},
	"Slash": {
		"id":             uint(163),
		"type":           "normal",
		"classification": uint(1),
	},
	"Sleep Talk": {
		"id":             uint(214),
		"type":           "normal",
		"classification": uint(0),
	},
	"Smelling Salts": {
		"id":             uint(265),
		"type":           "normal",
		"classification": uint(1),
	},
	"Smokescreen": {
		"id":             uint(108),
		"type":           "normal",
		"classification": uint(0),
	},
	"Snore": {
		"id":             uint(173),
		"type":           "normal",
		"classification": uint(2),
	},
	"Soft-Boiled": {
		"id":             uint(135),
		"type":           "normal",
		"classification": uint(0),
	},
	"Sonic Boom": {
		"id":             uint(49),
		"type":           "normal",
		"classification": uint(2),
	},
	"Spectral Thief": {
		"id":             uint(712),
		"type":           "ghost",
		"classification": uint(1),
	},
	"Spike Cannon": {
		"id":             uint(131),
		"type":           "normal",
		"classification": uint(1),
	},
	"Spit Up": {
		"id":             uint(255),
		"type":           "normal",
		"classification": uint(2),
	},
	"Splash": {
		"id":             uint(150),
		"type":           "normal",
		"classification": uint(0),
	},
	"Spotlight": {
		"id":             uint(671),
		"type":           "normal",
		"classification": uint(0),
	},
	"Stockpile": {
		"id":             uint(254),
		"type":           "normal",
		"classification": uint(0),
	},
	"Stomp": {
		"id":             uint(23),
		"type":           "normal",
		"classification": uint(1),
	},
	"Strength": {
		"id":             uint(70),
		"type":           "normal",
		"classification": uint(1),
	},
	"Substitute": {
		"id":             uint(164),
		"type":           "normal",
		"classification": uint(0),
	},
	"Super Fang": {
		"id":             uint(162),
		"type":           "normal",
		"classification": uint(1),
	},
	"Supersonic": {
		"id":             uint(48),
		"type":           "normal",
		"classification": uint(0),
	},
	"Swagger": {
		"id":             uint(207),
		"type":           "normal",
		"classification": uint(0),
	},
	"Swallow": {
		"id":             uint(256),
		"type":           "normal",
		"classification": uint(0),
	},
	"Sweet Scent": {
		"id":             uint(230),
		"type":           "normal",
		"classification": uint(0),
	},
	"Swift": {
		"id":             uint(129),
		"type":           "normal",
		"classification": uint(2),
	},
	"Swords Dance": {
		"id":             uint(14),
		"type":           "normal",
		"classification": uint(0),
	},
	"Tackle": {
		"id":             uint(33),
		"type":           "normal",
		"classification": uint(1),
	},
	"Tail Slap": {
		"id":             uint(541),
		"type":           "normal",
		"classification": uint(1),
	},
	"Tail Whip": {
		"id":             uint(39),
		"type":           "normal",
		"classification": uint(0),
	},
	"Take Down": {
		"id":             uint(36),
		"type":           "normal",
		"classification": uint(1),
	},
	"Tearful Look": {
		"id":             uint(715),
		"type":           "normal",
		"classification": uint(0),
	},
	"Techno Blast": {
		"id":             uint(546),
		"type":           "normal",
		"classification": uint(2),
	},
	"Teeter Dance": {
		"id":             uint(298),
		"type":           "normal",
		"classification": uint(0),
	},
	"Thrash": {
		"id":             uint(37),
		"type":           "normal",
		"classification": uint(1),
	},
	"Tickle": {
		"id":             uint(321),
		"type":           "normal",
		"classification": uint(0),
	},
	"Transform": {
		"id":             uint(144),
		"type":           "normal",
		"classification": uint(0),
	},
	"Tri Attack": {
		"id":             uint(161),
		"type":           "normal",
		"classification": uint(2),
	},
	"Trump Card": {
		"id":             uint(376),
		"type":           "normal",
		"classification": uint(2),
	},
	"Uproar": {
		"id":             uint(253),
		"type":           "normal",
		"classification": uint(2),
	},
	"Vice Grip": {
		"id":             uint(11),
		"type":           "normal",
		"classification": uint(1),
	},
	"Weather Ball": {
		"id":             uint(311),
		"type":           "normal",
		"classification": uint(2),
	},
	"Whirlwind": {
		"id":             uint(18),
		"type":           "normal",
		"classification": uint(0),
	},
	"Wish": {
		"id":             uint(273),
		"type":           "normal",
		"classification": uint(0),
	},
	"Work Up": {
		"id":             uint(526),
		"type":           "normal",
		"classification": uint(0),
	},
	"Wrap": {
		"id":             uint(35),
		"type":           "normal",
		"classification": uint(1),
	},
	"Wring Out": {
		"id":             uint(378),
		"type":           "normal",
		"classification": uint(2),
	},
	"Yawn": {
		"id":             uint(281),
		"type":           "normal",
		"classification": uint(0),
	},
	"Arm Thrust": {
		"id":             uint(292),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Aura Sphere": {
		"id":             uint(396),
		"type":           "fighting",
		"classification": uint(2),
	},
	"Brick Break": {
		"id":             uint(280),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Bulk Up": {
		"id":             uint(339),
		"type":           "fighting",
		"classification": uint(0),
	},
	"Circle Throw": {
		"id":             uint(509),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Close Combat": {
		"id":             uint(370),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Counter": {
		"id":             uint(68),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Cross Chop": {
		"id":             uint(238),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Detect": {
		"id":             uint(197),
		"type":           "fighting",
		"classification": uint(0),
	},
	"Double Kick": {
		"id":             uint(24),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Drain Punch": {
		"id":             uint(409),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Dynamic Punch": {
		"id":             uint(223),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Final Gambit": {
		"id":             uint(515),
		"type":           "fighting",
		"classification": uint(2),
	},
	"Flying Press": {
		"id":             uint(560),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Focus Blast": {
		"id":             uint(411),
		"type":           "fighting",
		"classification": uint(2),
	},
	"Focus Punch": {
		"id":             uint(264),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Force Palm": {
		"id":             uint(395),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Hammer Arm": {
		"id":             uint(359),
		"type":           "fighting",
		"classification": uint(1),
	},
	"High Jump Kick": {
		"id":             uint(136),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Jump Kick": {
		"id":             uint(26),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Karate Chop": {
		"id":             uint(2),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Low Kick": {
		"id":             uint(67),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Low Sweep": {
		"id":             uint(490),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Mach Punch": {
		"id":             uint(183),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Mat Block": {
		"id":             uint(561),
		"type":           "fighting",
		"classification": uint(0),
	},
	"Power-Up Punch": {
		"id":             uint(612),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Quick Guard": {
		"id":             uint(501),
		"type":           "fighting",
		"classification": uint(0),
	},
	"Revenge": {
		"id":             uint(279),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Reversal": {
		"id":             uint(179),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Rock Smash": {
		"id":             uint(249),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Rolling Kick": {
		"id":             uint(27),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Sacred Sword": {
		"id":             uint(533),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Secret Sword": {
		"id":             uint(548),
		"type":           "fighting",
		"classification": uint(2),
	},
	"Seismic Toss": {
		"id":             uint(69),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Sky Uppercut": {
		"id":             uint(327),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Storm Throw": {
		"id":             uint(480),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Submission": {
		"id":             uint(66),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Superpower": {
		"id":             uint(276),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Triple Kick": {
		"id":             uint(167),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Vacuum Wave": {
		"id":             uint(410),
		"type":           "fighting",
		"classification": uint(2),
	},
	"Vital Throw": {
		"id":             uint(233),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Wake-Up Slap": {
		"id":             uint(358),
		"type":           "fighting",
		"classification": uint(1),
	},
	"Blast Burn": {
		"id":             uint(307),
		"type":           "fire",
		"classification": uint(2),
	},
	"Blaze Kick": {
		"id":             uint(299),
		"type":           "fire",
		"classification": uint(1),
	},
	"Blue Flare": {
		"id":             uint(551),
		"type":           "fire",
		"classification": uint(2),
	},
	"Burn Up": {
		"id":             uint(682),
		"type":           "fire",
		"classification": uint(2),
	},
	"Ember": {
		"id":             uint(52),
		"type":           "fire",
		"classification": uint(2),
	},
	"Eruption": {
		"id":             uint(284),
		"type":           "fire",
		"classification": uint(2),
	},
	"Fiery Dance": {
		"id":             uint(552),
		"type":           "fire",
		"classification": uint(2),
	},
	"Fire Blast": {
		"id":             uint(126),
		"type":           "fire",
		"classification": uint(2),
	},
	"Fire Fang": {
		"id":             uint(424),
		"type":           "fire",
		"classification": uint(1),
	},
	"Fire Lash": {
		"id":             uint(680),
		"type":           "fire",
		"classification": uint(1),
	},
	"Fire Pledge": {
		"id":             uint(519),
		"type":           "fire",
		"classification": uint(2),
	},
	"Fire Punch": {
		"id":             uint(7),
		"type":           "fire",
		"classification": uint(1),
	},
	"Fire Spin": {
		"id":             uint(83),
		"type":           "fire",
		"classification": uint(2),
	},
	"Flame Burst": {
		"id":             uint(481),
		"type":           "fire",
		"classification": uint(2),
	},
	"Flame Charge": {
		"id":             uint(488),
		"type":           "fire",
		"classification": uint(1),
	},
	"Flame Wheel": {
		"id":             uint(172),
		"type":           "fire",
		"classification": uint(1),
	},
	"Flamethrower": {
		"id":             uint(53),
		"type":           "fire",
		"classification": uint(2),
	},
	"Flare Blitz": {
		"id":             uint(394),
		"type":           "fire",
		"classification": uint(1),
	},
	"Fusion Flare": {
		"id":             uint(558),
		"type":           "fire",
		"classification": uint(2),
	},
	"Heat Crash": {
		"id":             uint(535),
		"type":           "fire",
		"classification": uint(1),
	},
	"Heat Wave": {
		"id":             uint(257),
		"type":           "fire",
		"classification": uint(2),
	},
	"Incinerate": {
		"id":             uint(510),
		"type":           "fire",
		"classification": uint(2),
	},
	"Inferno": {
		"id":             uint(517),
		"type":           "fire",
		"classification": uint(2),
	},
	"Lava Plume": {
		"id":             uint(436),
		"type":           "fire",
		"classification": uint(2),
	},
	"Magma Storm": {
		"id":             uint(463),
		"type":           "fire",
		"classification": uint(2),
	},
	"Mystical Fire": {
		"id":             uint(595),
		"type":           "fire",
		"classification": uint(2),
	},
	"Overheat": {
		"id":             uint(315),
		"type":           "fire",
		"classification": uint(2),
	},
	"Sacred Fire": {
		"id":             uint(221),
		"type":           "fire",
		"classification": uint(1),
	},
	"Searing Shot": {
		"id":             uint(545),
		"type":           "fire",
		"classification": uint(2),
	},
	"Shell Trap": {
		"id":             uint(704),
		"type":           "fire",
		"classification": uint(2),
	},
	"Sunny Day": {
		"id":             uint(241),
		"type":           "fire",
		"classification": uint(0),
	},
	"V-create": {
		"id":             uint(557),
		"type":           "fire",
		"classification": uint(1),
	},
	"Will-O-Wisp": {
		"id":             uint(261),
		"type":           "fire",
		"classification": uint(0),
	},
	"Aurora Beam": {
		"id":             uint(62),
		"type":           "ice",
		"classification": uint(2),
	},
	"Aurora Veil": {
		"id":             uint(694),
		"type":           "ice",
		"classification": uint(0),
	},
	"Avalanche": {
		"id":             uint(419),
		"type":           "ice",
		"classification": uint(1),
	},
	"Blizzard": {
		"id":             uint(59),
		"type":           "ice",
		"classification": uint(2),
	},
	"Freeze Shock": {
		"id":             uint(553),
		"type":           "ice",
		"classification": uint(1),
	},
	"Freeze-Dry": {
		"id":             uint(573),
		"type":           "ice",
		"classification": uint(2),
	},
	"Frost Breath": {
		"id":             uint(524),
		"type":           "ice",
		"classification": uint(2),
	},
	"Glaciate": {
		"id":             uint(549),
		"type":           "ice",
		"classification": uint(2),
	},
	"Hail": {
		"id":             uint(258),
		"type":           "ice",
		"classification": uint(0),
	},
	"Haze": {
		"id":             uint(114),
		"type":           "ice",
		"classification": uint(0),
	},
	"Ice Ball": {
		"id":             uint(301),
		"type":           "ice",
		"classification": uint(1),
	},
	"Ice Beam": {
		"id":             uint(58),
		"type":           "ice",
		"classification": uint(2),
	},
	"Ice Burn": {
		"id":             uint(554),
		"type":           "ice",
		"classification": uint(2),
	},
	"Ice Fang": {
		"id":             uint(423),
		"type":           "ice",
		"classification": uint(1),
	},
	"Ice Hammer": {
		"id":             uint(665),
		"type":           "ice",
		"classification": uint(1),
	},
	"Ice Punch": {
		"id":             uint(8),
		"type":           "ice",
		"classification": uint(1),
	},
	"Ice Shard": {
		"id":             uint(420),
		"type":           "ice",
		"classification": uint(1),
	},
	"Icicle Crash": {
		"id":             uint(556),
		"type":           "ice",
		"classification": uint(1),
	},
	"Icicle Spear": {
		"id":             uint(333),
		"type":           "ice",
		"classification": uint(1),
	},
	"Icy Wind": {
		"id":             uint(196),
		"type":           "ice",
		"classification": uint(2),
	},
	"Mist": {
		"id":             uint(54),
		"type":           "ice",
		"classification": uint(0),
	},
	"Powder Snow": {
		"id":             uint(181),
		"type":           "ice",
		"classification": uint(2),
	},
	"Sheer Cold": {
		"id":             uint(329),
		"type":           "ice",
		"classification": uint(2),
	},
	"Bolt Strike": {
		"id":             uint(550),
		"type":           "electric",
		"classification": uint(1),
	},
	"Charge": {
		"id":             uint(268),
		"type":           "electric",
		"classification": uint(0),
	},
	"Charge Beam": {
		"id":             uint(451),
		"type":           "electric",
		"classification": uint(2),
	},
	"Discharge": {
		"id":             uint(435),
		"type":           "electric",
		"classification": uint(2),
	},
	"Eerie Impulse": {
		"id":             uint(598),
		"type":           "electric",
		"classification": uint(0),
	},
	"Electric Terrain": {
		"id":             uint(604),
		"type":           "electric",
		"classification": uint(0),
	},
	"Electrify": {
		"id":             uint(582),
		"type":           "electric",
		"classification": uint(0),
	},
	"Electro Ball": {
		"id":             uint(486),
		"type":           "electric",
		"classification": uint(2),
	},
	"Electroweb": {
		"id":             uint(527),
		"type":           "electric",
		"classification": uint(2),
	},
	"Fusion Bolt": {
		"id":             uint(559),
		"type":           "electric",
		"classification": uint(1),
	},
	"Ion Deluge": {
		"id":             uint(569),
		"type":           "electric",
		"classification": uint(0),
	},
	"Magnet Rise": {
		"id":             uint(393),
		"type":           "electric",
		"classification": uint(0),
	},
	"Magnetic Flux": {
		"id":             uint(602),
		"type":           "electric",
		"classification": uint(0),
	},
	"Nuzzle": {
		"id":             uint(609),
		"type":           "electric",
		"classification": uint(1),
	},
	"Parabolic Charge": {
		"id":             uint(570),
		"type":           "electric",
		"classification": uint(2),
	},
	"Shock Wave": {
		"id":             uint(351),
		"type":           "electric",
		"classification": uint(2),
	},
	"Spark": {
		"id":             uint(209),
		"type":           "electric",
		"classification": uint(1),
	},
	"Thunder": {
		"id":             uint(87),
		"type":           "electric",
		"classification": uint(2),
	},
	"Thunder Fang": {
		"id":             uint(422),
		"type":           "electric",
		"classification": uint(1),
	},
	"Thunder Punch": {
		"id":             uint(9),
		"type":           "electric",
		"classification": uint(1),
	},
	"Thunder Shock": {
		"id":             uint(84),
		"type":           "electric",
		"classification": uint(2),
	},
	"Thunder Wave": {
		"id":             uint(86),
		"type":           "electric",
		"classification": uint(0),
	},
	"Thunderbolt": {
		"id":             uint(85),
		"type":           "electric",
		"classification": uint(2),
	},
	"Volt Switch": {
		"id":             uint(521),
		"type":           "electric",
		"classification": uint(2),
	},
	"Volt Tackle": {
		"id":             uint(344),
		"type":           "electric",
		"classification": uint(1),
	},
	"Wild Charge": {
		"id":             uint(528),
		"type":           "electric",
		"classification": uint(1),
	},
	"Zap Cannon": {
		"id":             uint(192),
		"type":           "electric",
		"classification": uint(2),
	},
	"Zing Zap": {
		"id":             uint(716),
		"type":           "electric",
		"classification": uint(1),
	},
	"Acrobatics": {
		"id":             uint(512),
		"type":           "flying",
		"classification": uint(1),
	},
	"Aerial Ace": {
		"id":             uint(332),
		"type":           "flying",
		"classification": uint(1),
	},
	"Aeroblast": {
		"id":             uint(177),
		"type":           "flying",
		"classification": uint(2),
	},
	"Air Cutter": {
		"id":             uint(314),
		"type":           "flying",
		"classification": uint(2),
	},
	"Air Slash": {
		"id":             uint(403),
		"type":           "flying",
		"classification": uint(2),
	},
	"Beak Blast": {
		"id":             uint(690),
		"type":           "flying",
		"classification": uint(1),
	},
	"Bounce": {
		"id":             uint(340),
		"type":           "flying",
		"classification": uint(1),
	},
	"Brave Bird": {
		"id":             uint(413),
		"type":           "flying",
		"classification": uint(1),
	},
	"Chatter": {
		"id":             uint(448),
		"type":           "flying",
		"classification": uint(2),
	},
	"Defog": {
		"id":             uint(432),
		"type":           "flying",
		"classification": uint(3),
	},
	"Dragon Ascent": {
		"id":             uint(620),
		"type":           "flying",
		"classification": uint(1),
	},
	"Drill Peck": {
		"id":             uint(65),
		"type":           "flying",
		"classification": uint(1),
	},
	"Feather Dance": {
		"id":             uint(297),
		"type":           "flying",
		"classification": uint(0),
	},
	"Fly": {
		"id":             uint(19),
		"type":           "flying",
		"classification": uint(1),
	},
	"Gust": {
		"id":             uint(16),
		"type":           "flying",
		"classification": uint(2),
	},
	"Hurricane": {
		"id":             uint(542),
		"type":           "flying",
		"classification": uint(2),
	},
	"Mirror Move": {
		"id":             uint(119),
		"type":           "flying",
		"classification": uint(0),
	},
	"Oblivion Wing": {
		"id":             uint(613),
		"type":           "flying",
		"classification": uint(2),
	},
	"Peck": {
		"id":             uint(64),
		"type":           "flying",
		"classification": uint(1),
	},
	"Pluck": {
		"id":             uint(365),
		"type":           "flying",
		"classification": uint(1),
	},
	"Roost": {
		"id":             uint(355),
		"type":           "flying",
		"classification": uint(0),
	},
	"Sky Attack": {
		"id":             uint(143),
		"type":           "flying",
		"classification": uint(1),
	},
	"Sky Drop": {
		"id":             uint(507),
		"type":           "flying",
		"classification": uint(1),
	},
	"Tailwind": {
		"id":             uint(366),
		"type":           "flying",
		"classification": uint(0),
	},
	"Wing Attack": {
		"id":             uint(17),
		"type":           "flying",
		"classification": uint(1),
	},
	"Absorb": {
		"id":             uint(71),
		"type":           "grass",
		"classification": uint(2),
	},
	"Aromatherapy": {
		"id":             uint(312),
		"type":           "grass",
		"classification": uint(0),
	},
	"Bullet Seed": {
		"id":             uint(331),
		"type":           "grass",
		"classification": uint(1),
	},
	"Cotton Guard": {
		"id":             uint(538),
		"type":           "grass",
		"classification": uint(0),
	},
	"Cotton Spore": {
		"id":             uint(178),
		"type":           "grass",
		"classification": uint(0),
	},
	"Energy Ball": {
		"id":             uint(412),
		"type":           "grass",
		"classification": uint(2),
	},
	"Forest's Curse": {
		"id":             uint(571),
		"type":           "grass",
		"classification": uint(0),
	},
	"Frenzy Plant": {
		"id":             uint(338),
		"type":           "grass",
		"classification": uint(2),
	},
	"Giga Drain": {
		"id":             uint(202),
		"type":           "grass",
		"classification": uint(2),
	},
	"Grass Knot": {
		"id":             uint(447),
		"type":           "grass",
		"classification": uint(2),
	},
	"Grass Pledge": {
		"id":             uint(520),
		"type":           "grass",
		"classification": uint(2),
	},
	"Grass Whistle": {
		"id":             uint(320),
		"type":           "grass",
		"classification": uint(0),
	},
	"Grassy Terrain": {
		"id":             uint(580),
		"type":           "grass",
		"classification": uint(0),
	},
	"Horn Leech": {
		"id":             uint(532),
		"type":           "grass",
		"classification": uint(1),
	},
	"Ingrain": {
		"id":             uint(275),
		"type":           "grass",
		"classification": uint(0),
	},
	"Leaf Blade": {
		"id":             uint(348),
		"type":           "grass",
		"classification": uint(1),
	},
	"Leaf Storm": {
		"id":             uint(437),
		"type":           "grass",
		"classification": uint(2),
	},
	"Leaf Tornado": {
		"id":             uint(536),
		"type":           "grass",
		"classification": uint(2),
	},
	"Leafage": {
		"id":             uint(670),
		"type":           "grass",
		"classification": uint(1),
	},
	"Leech Seed": {
		"id":             uint(73),
		"type":           "grass",
		"classification": uint(0),
	},
	"Magical Leaf": {
		"id":             uint(345),
		"type":           "grass",
		"classification": uint(2),
	},
	"Mega Drain": {
		"id":             uint(72),
		"type":           "grass",
		"classification": uint(2),
	},
	"Needle Arm": {
		"id":             uint(302),
		"type":           "grass",
		"classification": uint(1),
	},
	"Petal Blizzard": {
		"id":             uint(572),
		"type":           "grass",
		"classification": uint(1),
	},
	"Petal Dance": {
		"id":             uint(80),
		"type":           "grass",
		"classification": uint(2),
	},
	"Power Whip": {
		"id":             uint(438),
		"type":           "grass",
		"classification": uint(1),
	},
	"Razor Leaf": {
		"id":             uint(75),
		"type":           "grass",
		"classification": uint(1),
	},
	"Seed Bomb": {
		"id":             uint(402),
		"type":           "grass",
		"classification": uint(1),
	},
	"Seed Flare": {
		"id":             uint(465),
		"type":           "grass",
		"classification": uint(2),
	},
	"Sleep Powder": {
		"id":             uint(79),
		"type":           "grass",
		"classification": uint(0),
	},
	"Solar Beam": {
		"id":             uint(76),
		"type":           "grass",
		"classification": uint(2),
	},
	"Solar Blade": {
		"id":             uint(669),
		"type":           "grass",
		"classification": uint(1),
	},
	"Spiky Shield": {
		"id":             uint(596),
		"type":           "grass",
		"classification": uint(0),
	},
	"Spore": {
		"id":             uint(147),
		"type":           "grass",
		"classification": uint(0),
	},
	"Strength Sap": {
		"id":             uint(668),
		"type":           "grass",
		"classification": uint(0),
	},
	"Stun Spore": {
		"id":             uint(78),
		"type":           "grass",
		"classification": uint(0),
	},
	"Synthesis": {
		"id":             uint(235),
		"type":           "grass",
		"classification": uint(0),
	},
	"Trop Kick": {
		"id":             uint(688),
		"type":           "grass",
		"classification": uint(1),
	},
	"Vine Whip": {
		"id":             uint(22),
		"type":           "grass",
		"classification": uint(1),
	},
	"Wood Hammer": {
		"id":             uint(452),
		"type":           "grass",
		"classification": uint(1),
	},
	"Worry Seed": {
		"id":             uint(388),
		"type":           "grass",
		"classification": uint(0),
	},
	"Bone Club": {
		"id":             uint(125),
		"type":           "ground",
		"classification": uint(1),
	},
	"Bone Rush": {
		"id":             uint(198),
		"type":           "ground",
		"classification": uint(1),
	},
	"Bonemerang": {
		"id":             uint(155),
		"type":           "ground",
		"classification": uint(1),
	},
	"Bulldoze": {
		"id":             uint(523),
		"type":           "ground",
		"classification": uint(1),
	},
	"Dig": {
		"id":             uint(91),
		"type":           "ground",
		"classification": uint(1),
	},
	"Drill Run": {
		"id":             uint(529),
		"type":           "ground",
		"classification": uint(1),
	},
	"Earth Power": {
		"id":             uint(414),
		"type":           "ground",
		"classification": uint(2),
	},
	"Earthquake": {
		"id":             uint(89),
		"type":           "ground",
		"classification": uint(1),
	},
	"Fissure": {
		"id":             uint(90),
		"type":           "ground",
		"classification": uint(1),
	},
	"High Horsepower": {
		"id":             uint(667),
		"type":           "ground",
		"classification": uint(1),
	},
	"Land's Wrath": {
		"id":             uint(616),
		"type":           "ground",
		"classification": uint(1),
	},
	"Magnitude": {
		"id":             uint(222),
		"type":           "ground",
		"classification": uint(1),
	},
	"Mud Bomb": {
		"id":             uint(426),
		"type":           "ground",
		"classification": uint(2),
	},
	"Mud Shot": {
		"id":             uint(341),
		"type":           "ground",
		"classification": uint(2),
	},
	"Mud Sport": {
		"id":             uint(300),
		"type":           "ground",
		"classification": uint(0),
	},
	"Mud-Slap": {
		"id":             uint(189),
		"type":           "ground",
		"classification": uint(2),
	},
	"Precipice Blades": {
		"id":             uint(619),
		"type":           "ground",
		"classification": uint(1),
	},
	"Rototiller": {
		"id":             uint(563),
		"type":           "ground",
		"classification": uint(0),
	},
	"Sand Attack": {
		"id":             uint(28),
		"type":           "ground",
		"classification": uint(0),
	},
	"Sand Tomb": {
		"id":             uint(328),
		"type":           "ground",
		"classification": uint(1),
	},
	"Shore Up": {
		"id":             uint(659),
		"type":           "ground",
		"classification": uint(0),
	},
	"Spikes": {
		"id":             uint(191),
		"type":           "ground",
		"classification": uint(0),
	},
	"Stomping Tantrum": {
		"id":             uint(707),
		"type":           "ground",
		"classification": uint(1),
	},
	"Thousand Arrows": {
		"id":             uint(614),
		"type":           "ground",
		"classification": uint(1),
	},
	"Thousand Waves": {
		"id":             uint(615),
		"type":           "ground",
		"classification": uint(1),
	},
	"Acid": {
		"id":             uint(51),
		"type":           "poison",
		"classification": uint(2),
	},
	"Acid Armor": {
		"id":             uint(151),
		"type":           "poison",
		"classification": uint(0),
	},
	"Acid Spray": {
		"id":             uint(491),
		"type":           "poison",
		"classification": uint(2),
	},
	"Baneful Bunker": {
		"id":             uint(661),
		"type":           "poison",
		"classification": uint(0),
	},
	"Belch": {
		"id":             uint(562),
		"type":           "poison",
		"classification": uint(2),
	},
	"Clear Smog": {
		"id":             uint(499),
		"type":           "poison",
		"classification": uint(2),
	},
	"Coil": {
		"id":             uint(489),
		"type":           "poison",
		"classification": uint(0),
	},
	"Cross Poison": {
		"id":             uint(440),
		"type":           "poison",
		"classification": uint(1),
	},
	"Gastro Acid": {
		"id":             uint(380),
		"type":           "poison",
		"classification": uint(0),
	},
	"Gunk Shot": {
		"id":             uint(441),
		"type":           "poison",
		"classification": uint(1),
	},
	"Poison Fang": {
		"id":             uint(305),
		"type":           "poison",
		"classification": uint(1),
	},
	"Poison Gas": {
		"id":             uint(139),
		"type":           "poison",
		"classification": uint(0),
	},
	"Poison Jab": {
		"id":             uint(398),
		"type":           "poison",
		"classification": uint(1),
	},
	"Poison Powder": {
		"id":             uint(77),
		"type":           "poison",
		"classification": uint(0),
	},
	"Poison Sting": {
		"id":             uint(40),
		"type":           "poison",
		"classification": uint(1),
	},
	"Poison Tail": {
		"id":             uint(342),
		"type":           "poison",
		"classification": uint(1),
	},
	"Purify": {
		"id":             uint(685),
		"type":           "poison",
		"classification": uint(0),
	},
	"Sludge": {
		"id":             uint(124),
		"type":           "poison",
		"classification": uint(2),
	},
	"Sludge Bomb": {
		"id":             uint(188),
		"type":           "poison",
		"classification": uint(2),
	},
	"Sludge Wave": {
		"id":             uint(482),
		"type":           "poison",
		"classification": uint(2),
	},
	"Smog": {
		"id":             uint(123),
		"type":           "poison",
		"classification": uint(2),
	},
	"Toxic": {
		"id":             uint(92),
		"type":           "poison",
		"classification": uint(0),
	},
	"Toxic Spikes": {
		"id":             uint(390),
		"type":           "poison",
		"classification": uint(0),
	},
	"Toxic Thread": {
		"id":             uint(672),
		"type":           "poison",
		"classification": uint(0),
	},
	"Venom Drench": {
		"id":             uint(599),
		"type":           "poison",
		"classification": uint(0),
	},
	"Venoshock": {
		"id":             uint(474),
		"type":           "poison",
		"classification": uint(2),
	},
	"Attack Order": {
		"id":             uint(454),
		"type":           "bug",
		"classification": uint(1),
	},
	"Bug Bite": {
		"id":             uint(450),
		"type":           "bug",
		"classification": uint(1),
	},
	"Bug Buzz": {
		"id":             uint(405),
		"type":           "bug",
		"classification": uint(2),
	},
	"Defend Order": {
		"id":             uint(455),
		"type":           "bug",
		"classification": uint(0),
	},
	"Fell Stinger": {
		"id":             uint(565),
		"type":           "bug",
		"classification": uint(1),
	},
	"First Impression": {
		"id":             uint(660),
		"type":           "bug",
		"classification": uint(1),
	},
	"Fury Cutter": {
		"id":             uint(210),
		"type":           "bug",
		"classification": uint(1),
	},
	"Heal Order": {
		"id":             uint(456),
		"type":           "bug",
		"classification": uint(0),
	},
	"Infestation": {
		"id":             uint(611),
		"type":           "bug",
		"classification": uint(2),
	},
	"Leech Life": {
		"id":             uint(141),
		"type":           "bug",
		"classification": uint(1),
	},
	"Lunge": {
		"id":             uint(679),
		"type":           "bug",
		"classification": uint(1),
	},
	"Megahorn": {
		"id":             uint(224),
		"type":           "bug",
		"classification": uint(1),
	},
	"Pin Missile": {
		"id":             uint(42),
		"type":           "bug",
		"classification": uint(1),
	},
	"Pollen Puff": {
		"id":             uint(676),
		"type":           "bug",
		"classification": uint(2),
	},
	"Powder": {
		"id":             uint(600),
		"type":           "bug",
		"classification": uint(0),
	},
	"Quiver Dance": {
		"id":             uint(483),
		"type":           "bug",
		"classification": uint(0),
	},
	"Rage Powder": {
		"id":             uint(476),
		"type":           "bug",
		"classification": uint(0),
	},
	"Signal Beam": {
		"id":             uint(324),
		"type":           "bug",
		"classification": uint(2),
	},
	"Silver Wind": {
		"id":             uint(318),
		"type":           "bug",
		"classification": uint(2),
	},
	"Spider Web": {
		"id":             uint(169),
		"type":           "bug",
		"classification": uint(0),
	},
	"Steamroller": {
		"id":             uint(537),
		"type":           "bug",
		"classification": uint(1),
	},
	"Sticky Web": {
		"id":             uint(564),
		"type":           "bug",
		"classification": uint(0),
	},
	"String Shot": {
		"id":             uint(81),
		"type":           "bug",
		"classification": uint(0),
	},
	"Struggle Bug": {
		"id":             uint(522),
		"type":           "bug",
		"classification": uint(2),
	},
	"Tail Glow": {
		"id":             uint(294),
		"type":           "bug",
		"classification": uint(0),
	},
	"Twineedle": {
		"id":             uint(41),
		"type":           "bug",
		"classification": uint(1),
	},
	"U-turn": {
		"id":             uint(369),
		"type":           "bug",
		"classification": uint(1),
	},
	"X-Scissor": {
		"id":             uint(404),
		"type":           "bug",
		"classification": uint(1),
	},
	"Assurance": {
		"id":             uint(372),
		"type":           "dark",
		"classification": uint(1),
	},
	"Beat Up": {
		"id":             uint(251),
		"type":           "dark",
		"classification": uint(1),
	},
	"Bite": {
		"id":             uint(44),
		"type":           "dark",
		"classification": uint(1),
	},
	"Brutal Swing": {
		"id":             uint(693),
		"type":           "dark",
		"classification": uint(1),
	},
	"Crunch": {
		"id":             uint(242),
		"type":           "dark",
		"classification": uint(1),
	},
	"Dark Pulse": {
		"id":             uint(399),
		"type":           "dark",
		"classification": uint(2),
	},
	"Dark Void": {
		"id":             uint(464),
		"type":           "dark",
		"classification": uint(0),
	},
	"Darkest Lariat": {
		"id":             uint(663),
		"type":           "dark",
		"classification": uint(1),
	},
	"Embargo": {
		"id":             uint(373),
		"type":           "dark",
		"classification": uint(0),
	},
	"Fake Tears": {
		"id":             uint(313),
		"type":           "dark",
		"classification": uint(0),
	},
	"Feint Attack": {
		"id":             uint(185),
		"type":           "dark",
		"classification": uint(1),
	},
	"Flatter": {
		"id":             uint(260),
		"type":           "dark",
		"classification": uint(0),
	},
	"Fling": {
		"id":             uint(374),
		"type":           "dark",
		"classification": uint(1),
	},
	"Foul Play": {
		"id":             uint(492),
		"type":           "dark",
		"classification": uint(1),
	},
	"Hone Claws": {
		"id":             uint(468),
		"type":           "dark",
		"classification": uint(0),
	},
	"Hyperspace Fury": {
		"id":             uint(621),
		"type":           "dark",
		"classification": uint(1),
	},
	"Knock Off": {
		"id":             uint(282),
		"type":           "dark",
		"classification": uint(1),
	},
	"Memento": {
		"id":             uint(262),
		"type":           "dark",
		"classification": uint(0),
	},
	"Nasty Plot": {
		"id":             uint(417),
		"type":           "dark",
		"classification": uint(0),
	},
	"Night Daze": {
		"id":             uint(539),
		"type":           "dark",
		"classification": uint(2),
	},
	"Night Slash": {
		"id":             uint(400),
		"type":           "dark",
		"classification": uint(1),
	},
	"Parting Shot": {
		"id":             uint(575),
		"type":           "dark",
		"classification": uint(0),
	},
	"Payback": {
		"id":             uint(371),
		"type":           "dark",
		"classification": uint(1),
	},
	"Power Trip": {
		"id":             uint(681),
		"type":           "dark",
		"classification": uint(1),
	},
	"Punishment": {
		"id":             uint(386),
		"type":           "dark",
		"classification": uint(1),
	},
	"Pursuit": {
		"id":             uint(228),
		"type":           "dark",
		"classification": uint(1),
	},
	"Quash": {
		"id":             uint(511),
		"type":           "dark",
		"classification": uint(0),
	},
	"Snarl": {
		"id":             uint(555),
		"type":           "dark",
		"classification": uint(2),
	},
	"Snatch": {
		"id":             uint(289),
		"type":           "dark",
		"classification": uint(0),
	},
	"Sucker Punch": {
		"id":             uint(389),
		"type":           "dark",
		"classification": uint(1),
	},
	"Switcheroo": {
		"id":             uint(415),
		"type":           "dark",
		"classification": uint(3),
	},
	"Taunt": {
		"id":             uint(269),
		"type":           "dark",
		"classification": uint(0),
	},
	"Thief": {
		"id":             uint(168),
		"type":           "dark",
		"classification": uint(1),
	},
	"Throat Chop": {
		"id":             uint(675),
		"type":           "dark",
		"classification": uint(1),
	},
	"Topsy-Turvy": {
		"id":             uint(576),
		"type":           "dark",
		"classification": uint(0),
	},
	"Torment": {
		"id":             uint(259),
		"type":           "dark",
		"classification": uint(0),
	},
	"Aqua Jet": {
		"id":             uint(453),
		"type":           "water",
		"classification": uint(1),
	},
	"Aqua Ring": {
		"id":             uint(392),
		"type":           "water",
		"classification": uint(0),
	},
	"Aqua Tail": {
		"id":             uint(401),
		"type":           "water",
		"classification": uint(1),
	},
	"Brine": {
		"id":             uint(362),
		"type":           "water",
		"classification": uint(2),
	},
	"Bubble": {
		"id":             uint(145),
		"type":           "water",
		"classification": uint(2),
	},
	"Bubble Beam": {
		"id":             uint(61),
		"type":           "water",
		"classification": uint(2),
	},
	"Clamp": {
		"id":             uint(128),
		"type":           "water",
		"classification": uint(1),
	},
	"Crabhammer": {
		"id":             uint(152),
		"type":           "water",
		"classification": uint(1),
	},
	"Dive": {
		"id":             uint(291),
		"type":           "water",
		"classification": uint(1),
	},
	"Hydro Cannon": {
		"id":             uint(308),
		"type":           "water",
		"classification": uint(2),
	},
	"Hydro Pump": {
		"id":             uint(56),
		"type":           "water",
		"classification": uint(2),
	},
	"Liquidation": {
		"id":             uint(710),
		"type":           "water",
		"classification": uint(1),
	},
	"Muddy Water": {
		"id":             uint(330),
		"type":           "water",
		"classification": uint(2),
	},
	"Octazooka": {
		"id":             uint(190),
		"type":           "water",
		"classification": uint(2),
	},
	"Origin Pulse": {
		"id":             uint(618),
		"type":           "water",
		"classification": uint(2),
	},
	"Rain Dance": {
		"id":             uint(240),
		"type":           "water",
		"classification": uint(0),
	},
	"Razor Shell": {
		"id":             uint(534),
		"type":           "water",
		"classification": uint(1),
	},
	"Scald": {
		"id":             uint(503),
		"type":           "water",
		"classification": uint(2),
	},
	"Soak": {
		"id":             uint(487),
		"type":           "water",
		"classification": uint(0),
	},
	"Sparkling Aria": {
		"id":             uint(664),
		"type":           "water",
		"classification": uint(2),
	},
	"Steam Eruption": {
		"id":             uint(592),
		"type":           "water",
		"classification": uint(2),
	},
	"Surf": {
		"id":             uint(57),
		"type":           "water",
		"classification": uint(2),
	},
	"Water Gun": {
		"id":             uint(55),
		"type":           "water",
		"classification": uint(2),
	},
	"Water Pledge": {
		"id":             uint(518),
		"type":           "water",
		"classification": uint(2),
	},
	"Water Pulse": {
		"id":             uint(352),
		"type":           "water",
		"classification": uint(2),
	},
	"Water Shuriken": {
		"id":             uint(594),
		"type":           "water",
		"classification": uint(2),
	},
	"Water Sport": {
		"id":             uint(346),
		"type":           "water",
		"classification": uint(0),
	},
	"Water Spout": {
		"id":             uint(323),
		"type":           "water",
		"classification": uint(2),
	},
	"Waterfall": {
		"id":             uint(127),
		"type":           "water",
		"classification": uint(1),
	},
	"Whirlpool": {
		"id":             uint(250),
		"type":           "water",
		"classification": uint(2),
	},
	"Withdraw": {
		"id":             uint(110),
		"type":           "water",
		"classification": uint(0),
	},
	"Agility": {
		"id":             uint(97),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Ally Switch": {
		"id":             uint(502),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Amnesia": {
		"id":             uint(133),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Barrier": {
		"id":             uint(112),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Calm Mind": {
		"id":             uint(347),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Confusion": {
		"id":             uint(93),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Cosmic Power": {
		"id":             uint(322),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Dream Eater": {
		"id":             uint(138),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Extrasensory": {
		"id":             uint(326),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Future Sight": {
		"id":             uint(248),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Gravity": {
		"id":             uint(356),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Guard Split": {
		"id":             uint(470),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Guard Swap": {
		"id":             uint(385),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Heal Block": {
		"id":             uint(377),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Heal Pulse": {
		"id":             uint(505),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Healing Wish": {
		"id":             uint(361),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Heart Stamp": {
		"id":             uint(531),
		"type":           "psychic",
		"classification": uint(1),
	},
	"Heart Swap": {
		"id":             uint(391),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Hyperspace Hole": {
		"id":             uint(593),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Hypnosis": {
		"id":             uint(95),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Imprison": {
		"id":             uint(286),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Instruct": {
		"id":             uint(689),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Kinesis": {
		"id":             uint(134),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Light Screen": {
		"id":             uint(113),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Lunar Dance": {
		"id":             uint(461),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Luster Purge": {
		"id":             uint(295),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Magic Coat": {
		"id":             uint(277),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Magic Room": {
		"id":             uint(478),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Meditate": {
		"id":             uint(96),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Miracle Eye": {
		"id":             uint(357),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Mirror Coat": {
		"id":             uint(243),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Mist Ball": {
		"id":             uint(296),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Power Split": {
		"id":             uint(471),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Power Swap": {
		"id":             uint(384),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Power Trick": {
		"id":             uint(379),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Prismatic Laser": {
		"id":             uint(711),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Psybeam": {
		"id":             uint(60),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Psychic": {
		"id":             uint(94),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Psychic Fangs": {
		"id":             uint(706),
		"type":           "psychic",
		"classification": uint(1),
	},
	"Psychic Terrain": {
		"id":             uint(678),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Psycho Boost": {
		"id":             uint(354),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Psycho Cut": {
		"id":             uint(427),
		"type":           "psychic",
		"classification": uint(1),
	},
	"Psycho Shift": {
		"id":             uint(375),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Psyshock": {
		"id":             uint(473),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Psystrike": {
		"id":             uint(540),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Psywave": {
		"id":             uint(149),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Reflect": {
		"id":             uint(115),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Rest": {
		"id":             uint(156),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Role Play": {
		"id":             uint(272),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Skill Swap": {
		"id":             uint(285),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Speed Swap": {
		"id":             uint(683),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Stored Power": {
		"id":             uint(500),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Synchronoise": {
		"id":             uint(485),
		"type":           "psychic",
		"classification": uint(2),
	},
	"Telekinesis": {
		"id":             uint(477),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Teleport": {
		"id":             uint(100),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Trick": {
		"id":             uint(271),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Trick Room": {
		"id":             uint(433),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Wonder Room": {
		"id":             uint(472),
		"type":           "psychic",
		"classification": uint(0),
	},
	"Zen Headbutt": {
		"id":             uint(428),
		"type":           "psychic",
		"classification": uint(1),
	},
	"Clanging Scales": {
		"id":             uint(691),
		"type":           "dragon",
		"classification": uint(2),
	},
	"Core Enforcer": {
		"id":             uint(687),
		"type":           "dragon",
		"classification": uint(2),
	},
	"Draco Meteor": {
		"id":             uint(434),
		"type":           "dragon",
		"classification": uint(2),
	},
	"Dragon Breath": {
		"id":             uint(225),
		"type":           "dragon",
		"classification": uint(2),
	},
	"Dragon Claw": {
		"id":             uint(337),
		"type":           "dragon",
		"classification": uint(1),
	},
	"Dragon Dance": {
		"id":             uint(349),
		"type":           "dragon",
		"classification": uint(0),
	},
	"Dragon Hammer": {
		"id":             uint(692),
		"type":           "dragon",
		"classification": uint(1),
	},
	"Dragon Pulse": {
		"id":             uint(406),
		"type":           "dragon",
		"classification": uint(2),
	},
	"Dragon Rage": {
		"id":             uint(82),
		"type":           "dragon",
		"classification": uint(2),
	},
	"Dragon Rush": {
		"id":             uint(407),
		"type":           "dragon",
		"classification": uint(1),
	},
	"Dragon Tail": {
		"id":             uint(525),
		"type":           "dragon",
		"classification": uint(1),
	},
	"Dual Chop": {
		"id":             uint(530),
		"type":           "dragon",
		"classification": uint(1),
	},
	"Outrage": {
		"id":             uint(200),
		"type":           "dragon",
		"classification": uint(1),
	},
	"Roar of Time": {
		"id":             uint(459),
		"type":           "dragon",
		"classification": uint(2),
	},
	"Spacial Rend": {
		"id":             uint(460),
		"type":           "dragon",
		"classification": uint(2),
	},
	"Twister": {
		"id":             uint(239),
		"type":           "dragon",
		"classification": uint(2),
	},
	"Accelerock": {
		"id":             uint(709),
		"type":           "rock",
		"classification": uint(1),
	},
	"Ancient Power": {
		"id":             uint(246),
		"type":           "rock",
		"classification": uint(2),
	},
	"Diamond Storm": {
		"id":             uint(591),
		"type":           "rock",
		"classification": uint(1),
	},
	"Head Smash": {
		"id":             uint(457),
		"type":           "rock",
		"classification": uint(1),
	},
	"Power Gem": {
		"id":             uint(408),
		"type":           "rock",
		"classification": uint(2),
	},
	"Rock Blast": {
		"id":             uint(350),
		"type":           "rock",
		"classification": uint(1),
	},
	"Rock Polish": {
		"id":             uint(397),
		"type":           "rock",
		"classification": uint(0),
	},
	"Rock Slide": {
		"id":             uint(157),
		"type":           "rock",
		"classification": uint(1),
	},
	"Rock Throw": {
		"id":             uint(88),
		"type":           "rock",
		"classification": uint(1),
	},
	"Rock Tomb": {
		"id":             uint(317),
		"type":           "rock",
		"classification": uint(1),
	},
	"Rock Wrecker": {
		"id":             uint(439),
		"type":           "rock",
		"classification": uint(1),
	},
	"Rollout": {
		"id":             uint(205),
		"type":           "rock",
		"classification": uint(1),
	},
	"Sandstorm": {
		"id":             uint(201),
		"type":           "rock",
		"classification": uint(0),
	},
	"Smack Down": {
		"id":             uint(479),
		"type":           "rock",
		"classification": uint(1),
	},
	"Stealth Rock": {
		"id":             uint(446),
		"type":           "rock",
		"classification": uint(0),
	},
	"Stone Edge": {
		"id":             uint(444),
		"type":           "rock",
		"classification": uint(1),
	},
	"Wide Guard": {
		"id":             uint(469),
		"type":           "rock",
		"classification": uint(0),
	},
	"Astonish": {
		"id":             uint(310),
		"type":           "ghost",
		"classification": uint(1),
	},
	"Confuse Ray": {
		"id":             uint(109),
		"type":           "ghost",
		"classification": uint(0),
	},
	"Curse": {
		"id":             uint(174),
		"type":           "ghost",
		"classification": uint(0),
	},
	"Destiny Bond": {
		"id":             uint(194),
		"type":           "ghost",
		"classification": uint(0),
	},
	"Grudge": {
		"id":             uint(288),
		"type":           "ghost",
		"classification": uint(0),
	},
	"Hex": {
		"id":             uint(506),
		"type":           "ghost",
		"classification": uint(2),
	},
	"Lick": {
		"id":             uint(122),
		"type":           "ghost",
		"classification": uint(1),
	},
	"Moongeist Beam": {
		"id":             uint(714),
		"type":           "ghost",
		"classification": uint(2),
	},
	"Night Shade": {
		"id":             uint(101),
		"type":           "ghost",
		"classification": uint(2),
	},
	"Nightmare": {
		"id":             uint(171),
		"type":           "ghost",
		"classification": uint(0),
	},
	"Ominous Wind": {
		"id":             uint(466),
		"type":           "ghost",
		"classification": uint(2),
	},
	"Phantom Force": {
		"id":             uint(566),
		"type":           "ghost",
		"classification": uint(1),
	},
	"Shadow Ball": {
		"id":             uint(247),
		"type":           "ghost",
		"classification": uint(2),
	},
	"Shadow Bone": {
		"id":             uint(708),
		"type":           "ghost",
		"classification": uint(1),
	},
	"Shadow Claw": {
		"id":             uint(421),
		"type":           "ghost",
		"classification": uint(1),
	},
	"Shadow Force": {
		"id":             uint(467),
		"type":           "ghost",
		"classification": uint(1),
	},
	"Shadow Punch": {
		"id":             uint(325),
		"type":           "ghost",
		"classification": uint(1),
	},
	"Shadow Sneak": {
		"id":             uint(425),
		"type":           "ghost",
		"classification": uint(1),
	},
	"Spirit Shackle": {
		"id":             uint(662),
		"type":           "ghost",
		"classification": uint(1),
	},
	"Spite": {
		"id":             uint(180),
		"type":           "ghost",
		"classification": uint(0),
	},
	"Trick-or-Treat": {
		"id":             uint(567),
		"type":           "ghost",
		"classification": uint(0),
	},
	"Anchor Shot": {
		"id":             uint(677),
		"type":           "steel",
		"classification": uint(1),
	},
	"Autotomize": {
		"id":             uint(475),
		"type":           "steel",
		"classification": uint(0),
	},
	"Bullet Punch": {
		"id":             uint(418),
		"type":           "steel",
		"classification": uint(1),
	},
	"Doom Desire": {
		"id":             uint(353),
		"type":           "steel",
		"classification": uint(2),
	},
	"Flash Cannon": {
		"id":             uint(430),
		"type":           "steel",
		"classification": uint(2),
	},
	"Gear Grind": {
		"id":             uint(544),
		"type":           "steel",
		"classification": uint(1),
	},
	"Gear Up": {
		"id":             uint(674),
		"type":           "steel",
		"classification": uint(0),
	},
	"Gyro Ball": {
		"id":             uint(360),
		"type":           "steel",
		"classification": uint(1),
	},
	"Heavy Slam": {
		"id":             uint(484),
		"type":           "steel",
		"classification": uint(1),
	},
	"Iron Defense": {
		"id":             uint(334),
		"type":           "steel",
		"classification": uint(0),
	},
	"Iron Head": {
		"id":             uint(442),
		"type":           "steel",
		"classification": uint(1),
	},
	"Iron Tail": {
		"id":             uint(231),
		"type":           "steel",
		"classification": uint(1),
	},
	"King's Shield": {
		"id":             uint(588),
		"type":           "steel",
		"classification": uint(0),
	},
	"Magnet Bomb": {
		"id":             uint(443),
		"type":           "steel",
		"classification": uint(1),
	},
	"Metal Burst": {
		"id":             uint(368),
		"type":           "steel",
		"classification": uint(1),
	},
	"Metal Claw": {
		"id":             uint(232),
		"type":           "steel",
		"classification": uint(1),
	},
	"Metal Sound": {
		"id":             uint(319),
		"type":           "steel",
		"classification": uint(0),
	},
	"Meteor Mash": {
		"id":             uint(309),
		"type":           "steel",
		"classification": uint(1),
	},
	"Mirror Shot": {
		"id":             uint(429),
		"type":           "steel",
		"classification": uint(2),
	},
	"Shift Gear": {
		"id":             uint(508),
		"type":           "steel",
		"classification": uint(0),
	},
	"Smart Strike": {
		"id":             uint(684),
		"type":           "steel",
		"classification": uint(1),
	},
	"Steel Wing": {
		"id":             uint(211),
		"type":           "steel",
		"classification": uint(1),
	},
	"Sunsteel Strike": {
		"id":             uint(713),
		"type":           "steel",
		"classification": uint(1),
	},
	"Aromatic Mist": {
		"id":             uint(597),
		"type":           "fairy",
		"classification": uint(0),
	},
	"Baby-Doll Eyes": {
		"id":             uint(608),
		"type":           "fairy",
		"classification": uint(0),
	},
	"Charm": {
		"id":             uint(204),
		"type":           "fairy",
		"classification": uint(0),
	},
	"Crafty Shield": {
		"id":             uint(578),
		"type":           "fairy",
		"classification": uint(0),
	},
	"Dazzling Gleam": {
		"id":             uint(605),
		"type":           "fairy",
		"classification": uint(2),
	},
	"Disarming Voice": {
		"id":             uint(574),
		"type":           "fairy",
		"classification": uint(2),
	},
	"Draining Kiss": {
		"id":             uint(577),
		"type":           "fairy",
		"classification": uint(2),
	},
	"Fairy Lock": {
		"id":             uint(587),
		"type":           "fairy",
		"classification": uint(0),
	},
	"Fairy Wind": {
		"id":             uint(584),
		"type":           "fairy",
		"classification": uint(2),
	},
	"Fleur Cannon": {
		"id":             uint(705),
		"type":           "fairy",
		"classification": uint(2),
	},
	"Floral Healing": {
		"id":             uint(666),
		"type":           "fairy",
		"classification": uint(0),
	},
	"Flower Shield": {
		"id":             uint(579),
		"type":           "fairy",
		"classification": uint(0),
	},
	"Geomancy": {
		"id":             uint(601),
		"type":           "fairy",
		"classification": uint(0),
	},
	"Misty Terrain": {
		"id":             uint(581),
		"type":           "fairy",
		"classification": uint(0),
	},
	"Moonblast": {
		"id":             uint(585),
		"type":           "fairy",
		"classification": uint(2),
	},
	"Moonlight": {
		"id":             uint(236),
		"type":           "fairy",
		"classification": uint(0),
	},
	"Nature's Madness": {
		"id":             uint(717),
		"type":           "fairy",
		"classification": uint(2),
	},
	"Play Rough": {
		"id":             uint(583),
		"type":           "fairy",
		"classification": uint(1),
	},
	"Sweet Kiss": {
		"id":             uint(186),
		"type":           "fairy",
		"classification": uint(0),
	},
	"Dynamax Cannon": {
		"id":             0,
		"type":           "dragon",
		"classification": uint(2),
	},
	"Snipe Shot": {
		"id":             0,
		"type":           "water",
		"classification": uint(2),
	},
	"Jaw Lock": {
		"id":             0,
		"type":           "dark",
		"classification": uint(1),
	},
	"Stuff Cheeks": {
		"id":             0,
		"type":           "normal",
		"classification": uint(0),
	},
	"No Retreat": {
		"id":             0,
		"type":           "fighting",
		"classification": uint(0),
	},
	"Tar Shot": {
		"id":             0,
		"type":           "rock",
		"classification": uint(0),
	},
	"Magic Powder": {
		"id":             0,
		"type":           "psychic",
		"classification": uint(0),
	},
	"Dragon Darts": {
		"id":             0,
		"type":           "dragon",
		"classification": uint(1),
	},
	"Teatime": {
		"id":             0,
		"type":           "normal",
		"classification": uint(0),
	},
	"Octolock": {
		"id":             0,
		"type":           "fighting",
		"classification": uint(0),
	},
	"Bolt Beak": {
		"id":             0,
		"type":           "electric",
		"classification": uint(1),
	},
	"Fishious Rend": {
		"id":             0,
		"type":           "water",
		"classification": uint(1),
	},
	"Court Change": {
		"id":             0,
		"type":           "normal",
		"classification": uint(0),
	},
	"Clangorous Soul": {
		"id":             0,
		"type":           "dragon",
		"classification": uint(0),
	},
	"Body Press": {
		"id":             0,
		"type":           "fighting",
		"classification": uint(1),
	},
	"Decorate": {
		"id":             0,
		"type":           "fairy",
		"classification": uint(0),
	},
	"Drum Beating": {
		"id":             0,
		"type":           "grass",
		"classification": uint(1),
	},
	"Snap Trap": {
		"id":             0,
		"type":           "grass",
		"classification": uint(1),
	},
	"Pyro Ball": {
		"id":             0,
		"type":           "fire",
		"classification": uint(1),
	},
	"Behemoth Blade": {
		"id":             0,
		"type":           "steel",
		"classification": uint(1),
	},
	"Behemoth Bash": {
		"id":             0,
		"type":           "steel",
		"classification": uint(1),
	},
	"Aura Wheel": {
		"id":             0,
		"type":           "electric",
		"classification": uint(1),
	},
	"Breaking Swipe": {
		"id":             0,
		"type":           "dragon",
		"classification": uint(1),
	},
	"Branch Poke": {
		"id":             0,
		"type":           "grass",
		"classification": uint(1),
	},
	"Overdrive": {
		"id":             0,
		"type":           "electric",
		"classification": uint(2),
	},
	"Apple Acid": {
		"id":             0,
		"type":           "grass",
		"classification": uint(2),
	},
	"Grav Apple": {
		"id":             0,
		"type":           "grass",
		"classification": uint(1),
	},
	"Spirit Break": {
		"id":             0,
		"type":           "fairy",
		"classification": uint(1),
	},
	"Strange Steam": {
		"id":             0,
		"type":           "fairy",
		"classification": uint(2),
	},
	"Life Dew": {
		"id":             0,
		"type":           "water",
		"classification": uint(0),
	},
	"Obstruct": {
		"id":             0,
		"type":           "dark",
		"classification": uint(0),
	},
	"False Surrender": {
		"id":             0,
		"type":           "dark",
		"classification": uint(1),
	},
	"Meteor Assault": {
		"id":             0,
		"type":           "fighting",
		"classification": uint(1),
	},
	"Eternabeam": {
		"id":             0,
		"type":           "dragon",
		"classification": uint(2),
	},
	"Steel Beam": {
		"id":             0,
		"type":           "steel",
		"classification": uint(2),
	},
	"1000 Folds": {
		"id":             uint(0),
		"type":           "Steel",
		"classification": uint(1),
	},
	"An Attack": {
		"id":             uint(0),
		"type":           "???",
		"classification": uint(1),
	},
	"Avian Rush": {
		"id":             uint(0),
		"type":           "???",
		"classification": uint(1),
	},
	"Awaken": {
		"id":             uint(0),
		"type":           "Fighting",
		"classification": uint(0),
	},
	"Backdraft": {
		"id":             uint(0),
		"type":           "Flying",
		"classification": uint(1),
	},
	"Bad Eggs": {
		"id":             uint(0),
		"type":           "Dark",
		"classification": uint(1),
	},
	"Ban Hammer": {
		"id":             uint(0),
		"type":           "Normal",
		"classification": uint(1),
	},
	"Blobby Bop": {
		"id":             uint(0),
		"type":           "Ice",
		"classification": uint(1),
	},
	"Boil Over": {
		"id":             uint(0),
		"type":           "Normal",
		"classification": uint(2),
	},
	"Boltbeam": {
		"id":             uint(0),
		"type":           "Electric",
		"classification": uint(2),
	},
	"Branding Blade": {
		"id":             uint(0),
		"type":           "Steel",
		"classification": uint(1),
	},
	"Brutal Punishment": {
		"id":             uint(0),
		"type":           "Fairy",
		"classification": uint(1),
	},
	"Chaos Dunk": {
		"id":             uint(0),
		"type":           "Rock",
		"classification": uint(1),
	},
	"Check 'Em": {
		"id":             uint(0),
		"type":           "Psychic",
		"classification": uint(2),
	},
	"Cheese Claw": {
		"id":             uint(0),
		"type":           "Fighting",
		"classification": uint(1),
	},
	"Cloud Breaker": {
		"id":             uint(0),
		"type":           "Flying",
		"classification": uint(2),
	},
	"Cold Cutter": {
		"id":             uint(0),
		"type":           "Ice",
		"classification": uint(1),
	},
	"Come n' Go": {
		"id":             uint(0),
		"type":           "Water",
		"classification": uint(2),
	},
	"Cope": {
		"id":             uint(0),
		"type":           "Bug",
		"classification": uint(2),
	},
	"Crusader Crash": {
		"id":             uint(0),
		"type":           "Fighting",
		"classification": uint(1),
	},
	"Daily Dose": {
		"id":             uint(0),
		"type":           "Poison",
		"classification": uint(2),
	},
	"Decay Drain": {
		"id":             uint(0),
		"type":           "Poison",
		"classification": uint(1),
	},
	"Dildo Cannon": {
		"id":             uint(0),
		"type":           "Dragon",
		"classification": uint(1),
	},
	"Dragon Burst": {
		"id":             uint(0),
		"type":           "Dragon",
		"classification": uint(2),
	},
	"Dragon Fist": {
		"id":             uint(0),
		"type":           "Dragon",
		"classification": uint(1),
	},
	"Elbow Drop": {
		"id":             uint(0),
		"type":           "Ghost",
		"classification": uint(1),
	},
	"Enema": {
		"id":             uint(0),
		"type":           "Grass",
		"classification": uint(1),
	},
	"Erosion Wave": {
		"id":             uint(0),
		"type":           "Rock",
		"classification": uint(2),
	},
	"Falcon Punch": {
		"id":             uint(0),
		"type":           "Flying",
		"classification": uint(1),
	},
	"Final Hour": {
		"id":             uint(0),
		"type":           "Dark",
		"classification": uint(2),
	},
	"Fire Bomb": {
		"id":             uint(0),
		"type":           "Fire",
		"classification": uint(1),
	},
	"Fizzbitch": {
		"id":             uint(0),
		"type":           "Grass",
		"classification": uint(2),
	},
	"Focus Munch": {
		"id":             uint(0),
		"type":           "Fighting",
		"classification": uint(0),
	},
	"For You": {
		"id":             uint(0),
		"type":           "Fighting",
		"classification": uint(1),
	},
	"Fruit Juice": {
		"id":             uint(0),
		"type":           "Fairy",
		"classification": uint(2),
	},
	"Fruit Punch": {
		"id":             uint(0),
		"type":           "Fairy",
		"classification": uint(1),
	},
	"Futaba Break": {
		"id":             uint(0),
		"type":           "Grass",
		"classification": uint(1),
	},
	"Gay Agenda": {
		"id":             uint(0),
		"type":           "Fairy",
		"classification": uint(0),
	},
	"Gazer Beam": {
		"id":             uint(0),
		"type":           "Steel",
		"classification": uint(2),
	},
	"Genesis Boost": {
		"id":             uint(0),
		"type":           "Fairy",
		"classification": uint(0),
	},
	"Great Rage": {
		"id":             uint(0),
		"type":           "Grass",
		"classification": uint(1),
	},
	"Holy Duty": {
		"id":             uint(0),
		"type":           "Fire",
		"classification": uint(2),
	},
	"Homerun Bat": {
		"id":             uint(0),
		"type":           "Normal",
		"classification": uint(1),
	},
	"Hulk Up": {
		"id":             uint(0),
		"type":           "Fighting",
		"classification": uint(0),
	},
	"Inverse Room": {
		"id":             uint(0),
		"type":           "???",
		"classification": uint(0),
	},
	"It's Over": {
		"id":             uint(0),
		"type":           "Bug",
		"classification": uint(0),
	},
	"Lactose Shot": {
		"id":             uint(0),
		"type":           "Fairy",
		"classification": uint(2),
	},
	"Lick Clean": {
		"id":             uint(0),
		"type":           "Water",
		"classification": uint(1),
	},
	"Livewire": {
		"id":             uint(0),
		"type":           "Electric",
		"classification": uint(0),
	},
	"Mating Press": {
		"id":             uint(0),
		"type":           "Fairy",
		"classification": uint(1),
	},
	"Max Memeitude": {
		"id":             uint(0),
		"type":           "???",
		"classification": uint(1),
	},
	"Meds Now": {
		"id":             uint(0),
		"type":           "Poison",
		"classification": uint(2),
	},
	"Meme": {
		"id":             uint(0),
		"type":           "???",
		"classification": uint(2),
	},
	"Meme Jr.": {
		"id":             uint(0),
		"type":           "???",
		"classification": uint(1),
	},
	"Mop": {
		"id":             uint(0),
		"type":           "Fairy",
		"classification": uint(0),
	},
	"More Dakka": {
		"id":             uint(0),
		"type":           "Steel",
		"classification": uint(2),
	},
	"More Gun": {
		"id":             uint(0),
		"type":           "Normal",
		"classification": uint(2),
	},
	"Mud Maelstrom": {
		"id":             uint(0),
		"type":           "Ground",
		"classification": uint(2),
	},
	"Nuclear Winter": {
		"id":             uint(0),
		"type":           "Ice",
		"classification": uint(2),
	},
	"Overbite": {
		"id":             uint(0),
		"type":           "Dark",
		"classification": uint(1),
	},
	"Overenergize": {
		"id":             uint(0),
		"type":           "Electric",
		"classification": uint(2),
	},
	"Ow The Edge": {
		"id":             uint(0),
		"type":           "Dark",
		"classification": uint(1),
	},
	"Petrify": {
		"id":             uint(0),
		"type":           "Rock",
		"classification": uint(0),
	},
	"Phantom Fang": {
		"id":             uint(0),
		"type":           "Ghost",
		"classification": uint(1),
	},
	"Phase Through": {
		"id":             uint(0),
		"type":           "Ghost",
		"classification": uint(1),
	},
	"Pixie Pummel": {
		"id":             uint(0),
		"type":           "Fairy",
		"classification": uint(1),
	},
	"Please Don't Do That": {
		"id":             uint(0),
		"type":           "Psychic",
		"classification": uint(0),
	},
	"Plunder": {
		"id":             uint(0),
		"type":           "Water",
		"classification": uint(1),
	},
	"Psycho Fists": {
		"id":             uint(0),
		"type":           "Psychic",
		"classification": uint(1),
	},
	"Puke Blood": {
		"id":             uint(0),
		"type":           "Bug",
		"classification": uint(2),
	},
	"Punch Out": {
		"id":             uint(0),
		"type":           "Fighting",
		"classification": uint(1),
	},
	"Quick Sand": {
		"id":             uint(0),
		"type":           "Ground",
		"classification": uint(1),
	},
	"Regenerate": {
		"id":             uint(0),
		"type":           "Grass",
		"classification": uint(0),
	},
	"Riot Shield": {
		"id":             uint(0),
		"type":           "Fighting",
		"classification": uint(1),
	},
	"Rock Clock": {
		"id":             uint(0),
		"type":           "Rock",
		"classification": uint(1),
	},
	"Scorched Earth": {
		"id":             uint(0),
		"type":           "Ground",
		"classification": uint(2),
	},
	"Shadow Scales": {
		"id":             uint(0),
		"type":           "Ghost",
		"classification": uint(2),
	},
	"Shine Strike": {
		"id":             uint(0),
		"type":           "Steel",
		"classification": uint(1),
	},
	"Shitpost": {
		"id":             uint(0),
		"type":           "Ground",
		"classification": uint(1),
	},
	"Skull Cannon": {
		"id":             uint(0),
		"type":           "Dark",
		"classification": uint(2),
	},
	"Sleazy Spores": {
		"id":             uint(0),
		"type":           "Grass",
		"classification": uint(0),
	},
	"Slime Gulp": {
		"id":             uint(0),
		"type":           "Poison",
		"classification": uint(1),
	},
	"Soul Crusher": {
		"id":             uint(0),
		"type":           "Ghost",
		"classification": uint(2),
	},
	"Speed Weed": {
		"id":             uint(0),
		"type":           "Grass",
		"classification": uint(1),
	},
	"Spook Out": {
		"id":             uint(0),
		"type":           "Ghost",
		"classification": uint(1),
	},
	"Spooperpower": {
		"id":             uint(0),
		"type":           "Ghost",
		"classification": uint(2),
	},
	"Strato Blade": {
		"id":             uint(0),
		"type":           "Flying",
		"classification": uint(1),
	},
	"Strum": {
		"id":             uint(0),
		"type":           "Normal",
		"classification": uint(2),
	},
	"Sudoku": {
		"id":             uint(0),
		"type":           "Normal",
		"classification": uint(0),
	},
	"Super Snore": {
		"id":             uint(0),
		"type":           "Ice",
		"classification": uint(1),
	},
	"Swindle": {
		"id":             uint(0),
		"type":           "Dark",
		"classification": uint(2),
	},
	"Think Fast": {
		"id":             uint(0),
		"type":           "Psychic",
		"classification": uint(2),
	},
	"Toke": {
		"id":             uint(0),
		"type":           "Fire",
		"classification": uint(0),
	},
	"Tombstoner": {
		"id":             uint(0),
		"type":           "Rock",
		"classification": uint(1),
	},
	"Toxiravage": {
		"id":             uint(0),
		"type":           "Poison",
		"classification": uint(1),
	},
	"Trick Stab": {
		"id":             uint(0),
		"type":           "Dark",
		"classification": uint(1),
	},
	"Trigger": {
		"id":             uint(0),
		"type":           "Psychic",
		"classification": uint(2),
	},
	"Turnabout": {
		"id":             uint(0),
		"type":           "Ghost",
		"classification": uint(0),
	},
	"Villify": {
		"id":             uint(0),
		"type":           "Dark",
		"classification": uint(0),
	},
	"Voltaic Cyclone": {
		"id":             uint(0),
		"type":           "Electric",
		"classification": uint(2),
	},
	"Warhead": {
		"id":             uint(0),
		"type":           "Steel",
		"classification": uint(2),
	},
	"Weird Flex": {
		"id":             uint(0),
		"type":           "Fighting",
		"classification": uint(0),
	},
	"Wings Of Correction": {
		"id":             uint(0),
		"type":           "Flying",
		"classification": uint(2),
	},
	"Wow Wiener": {
		"id":             uint(0),
		"type":           "Fire",
		"classification": uint(2),
	},
}

var itemData = map[string]map[string]interface{}{
	"Ability Capsule": {
		"id": uint(645),
	},
	"Abomasite": {
		"id": uint(674),
	},
	"Absolite": {
		"id": uint(677),
	},
	"Absorb Bulb": {
		"id": uint(545),
	},
	"Adamant Orb": {
		"id": uint(135),
	},
	"Adrenaline Orb": {
		"id": uint(846),
	},
	"Aerodactylite": {
		"id": uint(672),
	},
	"Aggronite": {
		"id": uint(667),
	},
	"Aguav Berry": {
		"id": uint(162),
	},
	"Air Balloon": {
		"id": uint(541),
	},
	"Alakazite": {
		"id": uint(679),
	},
	"Aloraichium Z": {
		"id":   uint(803),
		"type": "electric",
	},
	"Altarianite": {
		"id": uint(755),
	},
	"Ampharosite": {
		"id": uint(658),
	},
	"Amulet Coin": {
		"id": uint(223),
	},
	"Antidote": {
		"id": uint(18),
	},
	"Apicot Berry": {
		"id": uint(205),
	},
	"Armor Fossil": {
		"id": uint(104),
	},
	"Aspear Berry": {
		"id": uint(153),
	},
	"Assault Vest": {
		"id": uint(640),
	},
	"Audinite": {
		"id": uint(757),
	},
	"Awakening": {
		"id": uint(21),
	},
	"Babiri Berry": {
		"id":   uint(199),
		"type": "steel",
	},
	"Balm Mushroom": {
		"id": uint(580),
	},
	"Banettite": {
		"id": uint(668),
	},
	"Beast Ball": {
		"id": uint(851),
	},
	"Beedrillite": {
		"id": uint(770),
	},
	"Berry Juice": {
		"id": uint(43),
	},
	"Big Malasada": {
		"id": uint(852),
	},
	"Big Mushroom": {
		"id": uint(87),
	},
	"Big Nugget": {
		"id": uint(581),
	},
	"Big Pearl": {
		"id": uint(89),
	},
	"Big Root": {
		"id": uint(296),
	},
	"Binding Band": {
		"id": uint(544),
	},
	"Black Belt": {
		"id":   uint(241),
		"type": "fighting",
	},
	"Black Glasses": {
		"id":   uint(240),
		"type": "dark",
	},
	"Black Sludge": {
		"id": uint(281),
	},
	"Blastoisinite": {
		"id": uint(661),
	},
	"Blazikenite": {
		"id": uint(664),
	},
	"Blue Orb": {
		"id": uint(535),
	},
	"Blue Shard": {
		"id": uint(73),
	},
	"Bluk Berry": {
		"id": uint(165),
	},
	"Bottle Cap": {
		"id": uint(795),
	},
	"Bright Powder": {
		"id": uint(213),
	},
	"Bug Gem": {
		"id":   uint(558),
		"type": "bug",
	},
	"Bug Memory": {
		"id":   uint(909),
		"type": "bug",
	},
	"Buginium Z": {
		"id":   uint(787),
		"type": "bug",
	},
	"Burn Drive": {
		"id":   uint(118),
		"type": "fire",
	},
	"Burn Heal": {
		"id": uint(19),
	},
	"Calcium": {
		"id": uint(49),
	},
	"Cameruptite": {
		"id": uint(767),
	},
	"Carbos": {
		"id": uint(48),
	},
	"Casteliacone": {
		"id": uint(591),
	},
	"Cell Battery": {
		"id": uint(546),
	},
	"Charcoal": {
		"id":   uint(249),
		"type": "fire",
	},
	"Charizardite X": {
		"id": uint(660),
	},
	"Charizardite Y": {
		"id": uint(678),
	},
	"Charti Berry": {
		"id":   uint(195),
		"type": "rock",
	},
	"Cheri Berry": {
		"id": uint(149),
	},
	"Chesto Berry": {
		"id": uint(150),
	},
	"Chilan Berry": {
		"id":   uint(200),
		"type": "normal",
	},
	"Chill Drive": {
		"id":   uint(119),
		"type": "ice",
	},
	"Choice Band": {
		"id": uint(220),
	},
	"Choice Scarf": {
		"id": uint(287),
	},
	"Choice Specs": {
		"id": uint(297),
	},
	"Chople Berry": {
		"id":   uint(189),
		"type": "fighting",
	},
	"Cleanse Tag": {
		"id": uint(224),
	},
	"Clever Wing": {
		"id": uint(569),
	},
	"Coba Berry": {
		"id":   uint(192),
		"type": "flying",
	},
	"Colbur Berry": {
		"id":   uint(198),
		"type": "dark",
	},
	"Comet Shard": {
		"id": uint(583),
	},
	"Cover Fossil": {
		"id": uint(572),
	},
	"Damp Rock": {
		"id":   uint(285),
		"type": "water",
	},
	"Dark Gem": {
		"id":   uint(562),
		"type": "dark",
	},
	"Dark Memory": {
		"id":   uint(919),
		"type": "dark",
	},
	"Darkinium Z": {
		"id":   uint(791),
		"type": "dark",
	},
	"Dawn Stone": {
		"id": uint(109),
	},
	"Decidium Z": {
		"id":   uint(798),
		"type": "ghost",
	},
	"Deep Sea Scale": {
		"id": uint(227),
	},
	"Deep Sea Tooth": {
		"id": uint(226),
	},
	"Destiny Knot": {
		"id": uint(280),
	},
	"Diancite": {
		"id": uint(764),
	},
	"Dire Hit": {
		"id": uint(56),
	},
	"Dive Ball": {
		"id": uint(7),
	},
	"Douse Drive": {
		"id":   uint(116),
		"type": "water",
	},
	"Draco Plate": {
		"id":   uint(311),
		"type": "dragon",
	},
	"Dragon Fang": {
		"id":   uint(250),
		"type": "dragon",
	},
	"Dragon Gem": {
		"id":   uint(561),
		"type": "dragon",
	},
	"Dragon Memory": {
		"id":   uint(918),
		"type": "dragon",
	},
	"Dragon Scale": {
		"id": uint(235),
	},
	"Dragonium Z": {
		"id":   uint(790),
		"type": "dragon",
	},
	"Dread Plate": {
		"id":   uint(312),
		"type": "dark",
	},
	"Dubious Disc": {
		"id": uint(324),
	},
	"Dusk Ball": {
		"id": uint(13),
	},
	"Dusk Stone": {
		"id": uint(108),
	},
	"Earth Plate": {
		"id":   uint(305),
		"type": "ground",
	},
	"Eevium Z": {
		"id":   uint(805),
		"type": "normal",
	},
	"Eject Button": {
		"id": uint(547),
	},
	"Electirizer": {
		"id": uint(322),
	},
	"Electric Gem": {
		"id":   uint(550),
		"type": "electric",
	},
	"Electric Memory": {
		"id":   uint(915),
		"type": "electric",
	},
	"Electric Seed": {
		"id": uint(881),
	},
	"Electrium Z": {
		"id":   uint(779),
		"type": "electric",
	},
	"Elixir": {
		"id": uint(40),
	},
	"Energy Powder": {
		"id": uint(34),
	},
	"Energy Root": {
		"id": uint(35),
	},
	"Escape Rope": {
		"id": uint(78),
	},
	"Ether": {
		"id": uint(38),
	},
	"Everstone": {
		"id": uint(229),
	},
	"Eviolite": {
		"id": uint(538),
	},
	"Expert Belt": {
		"id": uint(268),
	},
	"Fairium Z": {
		"id":   uint(793),
		"type": "fairy",
	},
	"Fairy Gem": {
		"id":   uint(715),
		"type": "fairy",
	},
	"Fairy Memory": {
		"id":   uint(920),
		"type": "fairy",
	},
	"Fast Ball": {
		"id": uint(492),
	},
	"Festival Ticket": {
		"id": uint(844),
	},
	"Fighting Gem": {
		"id":   uint(553),
		"type": "fighting",
	},
	"Fighting Memory": {
		"id": uint(904),
	},
	"Fightinium Z": {
		"id":   uint(782),
		"type": "fighting",
	},
	"Figy Berry": {
		"id": uint(159),
	},
	"Fire Gem": {
		"id":   uint(548),
		"type": "fire",
	},
	"Fire Memory": {
		"id":   uint(912),
		"type": "fire",
	},
	"Fire Stone": {
		"id": uint(82),
	},
	"Firium Z": {
		"id":   uint(777),
		"type": "fire",
	},
	"Fist Plate": {
		"id":   uint(303),
		"type": "fighting",
	},
	"Flame Orb": {
		"id": uint(273),
	},
	"Flame Plate": {
		"id":   uint(298),
		"type": "fire",
	},
	"Float Stone": {
		"id": uint(539),
	},
	"Flying Gem": {
		"id":   uint(556),
		"type": "flying",
	},
	"Flying Memory": {
		"id":   uint(905),
		"type": "flying",
	},
	"Flyinium Z": {
		"id":   uint(785),
		"type": "flying",
	},
	"Focus Band": {
		"id": uint(230),
	},
	"Focus Sash": {
		"id": uint(275),
	},
	"Fresh Water": {
		"id": uint(30),
	},
	"Friend Ball": {
		"id": uint(497),
	},
	"Full Heal": {
		"id": uint(27),
	},
	"Full Incense": {
		"id": uint(316),
	},
	"Full Restore": {
		"id": uint(23),
	},
	"Galladite": {
		"id": uint(756),
	},
	"Ganlon Berry": {
		"id": uint(202),
	},
	"Garchompite": {
		"id": uint(683),
	},
	"Gardevoirite": {
		"id": uint(657),
	},
	"Gengarite": {
		"id": uint(656),
	},
	"Genius Wing": {
		"id": uint(568),
	},
	"Ghost Gem": {
		"id":   uint(560),
		"type": "ghost",
	},
	"Ghost Memory": {
		"id":   uint(910),
		"type": "ghost",
	},
	"Ghostium Z": {
		"id":   uint(789),
		"type": "ghost",
	},
	"Glalitite": {
		"id": uint(763),
	},
	"Gold Bottle Cap": {
		"id": uint(796),
	},
	"Grass Gem": {
		"id":   uint(551),
		"type": "grass",
	},
	"Grass Memory": {
		"id":   uint(914),
		"type": "grass",
	},
	"Grassium Z": {
		"id":   uint(780),
		"type": "grass",
	},
	"Grassy Seed": {
		"id": uint(884),
	},
	"Great Ball": {
		"id": uint(3),
	},
	"Green Shard": {
		"id": uint(75),
	},
	"Grepa Berry": {
		"id": uint(173),
	},
	"Grip Claw": {
		"id": uint(286),
	},
	"Griseous Orb": {
		"id": uint(112),
	},
	"Ground Gem": {
		"id":   uint(555),
		"type": "ground",
	},
	"Ground Memory": {
		"id":   uint(907),
		"type": "ground",
	},
	"Groundium Z": {
		"id":   uint(784),
		"type": "ground",
	},
	"Guard Spec.": {
		"id": uint(55),
	},
	"Gyaradosite": {
		"id": uint(676),
	},
	"HP Up": {
		"id": uint(45),
	},
	"Haban Berry": {
		"id":   uint(197),
		"type": "dragon",
	},
	"Hard Stone": {
		"id":   uint(238),
		"type": "rock",
	},
	"Heal Ball": {
		"id": uint(14),
	},
	"Heal Powder": {
		"id": uint(36),
	},
	"Health Wing": {
		"id": uint(565),
	},
	"Heart Scale": {
		"id": uint(93),
	},
	"Heat Rock": {
		"id":   uint(284),
		"type": "fire",
	},
	"Heavy Ball": {
		"id": uint(495),
	},
	"Heracronite": {
		"id": uint(680),
	},
	"Hondew Berry": {
		"id": uint(172),
	},
	"Honey": {
		"id": uint(94),
	},
	"Houndoominite": {
		"id": uint(666),
	},
	"Hyper Potion": {
		"id": uint(25),
	},
	"Iapapa Berry": {
		"id": uint(163),
	},
	"Ice Gem": {
		"id":   uint(552),
		"type": "ice",
	},
	"Ice Heal": {
		"id": uint(20),
	},
	"Ice Memory": {
		"id":   uint(917),
		"type": "ice",
	},
	"Ice Stone": {
		"id": uint(849),
	},
	"Icicle Plate": {
		"id":   uint(302),
		"type": "ice",
	},
	"Icium Z": {
		"id":   uint(781),
		"type": "ice",
	},
	"Icy Rock": {
		"id":   uint(282),
		"type": "ice",
	},
	"Incinium Z": {
		"id":   uint(799),
		"type": "dark",
	},
	"Insect Plate": {
		"id":   uint(308),
		"type": "bug",
	},
	"Iron": {
		"id": uint(47),
	},
	"Iron Ball": {
		"id": uint(278),
	},
	"Iron Plate": {
		"id":   uint(313),
		"type": "steel",
	},
	"Kangaskhanite": {
		"id": uint(675),
	},
	"Kasib Berry": {
		"id":   uint(196),
		"type": "ghost",
	},
	"Kebia Berry": {
		"id":   uint(190),
		"type": "poison",
	},
	"Kee Berry": {
		"id": uint(687),
	},
	"Kelpsy Berry": {
		"id": uint(170),
	},
	"King's Rock": {
		"id": uint(221),
	},
	"Kommonium Z": {
		"id":   uint(926),
		"type": "dragon",
	},
	"Lagging Tail": {
		"id": uint(279),
	},
	"Lansat Berry": {
		"id": uint(206),
	},
	"Latiasite": {
		"id": uint(684),
	},
	"Latiosite": {
		"id": uint(685),
	},
	"Lava Cookie": {
		"id": uint(42),
	},
	"Lax Incense": {
		"id": uint(255),
	},
	"Leaf Stone": {
		"id": uint(85),
	},
	"Leftovers": {
		"id": uint(234),
	},
	"Lemonade": {
		"id": uint(32),
	},
	"Leppa Berry": {
		"id": uint(154),
	},
	"Level Ball": {
		"id": uint(493),
	},
	"Liechi Berry": {
		"id": uint(201),
	},
	"Life Orb": {
		"id": uint(270),
	},
	"Light Ball": {
		"id": uint(236),
	},
	"Light Clay": {
		"id": uint(269),
	},
	"Lopunnite": {
		"id": uint(768),
	},
	"Love Ball": {
		"id": uint(496),
	},
	"Lucarionite": {
		"id": uint(673),
	},
	"Luck Incense": {
		"id": uint(319),
	},
	"Lucky Egg": {
		"id": uint(231),
	},
	"Lucky Punch": {
		"id": uint(256),
	},
	"Lum Berry": {
		"id": uint(157),
	},
	"Luminous Moss": {
		"id": uint(648),
	},
	"Lumiose Galette": {
		"id": uint(708),
	},
	"Lunalium Z": {
		"id":   uint(922),
		"type": "ghost",
	},
	"Lure Ball": {
		"id": uint(494),
	},
	"Lustrous Orb": {
		"id": uint(136),
	},
	"Luxury Ball": {
		"id": uint(11),
	},
	"Lycanium Z": {
		"id":   uint(925),
		"type": "rock",
	},
	"Magmarizer": {
		"id": uint(323),
	},
	"Magnet": {
		"id":   uint(242),
		"type": "electric",
	},
	"Mago Berry": {
		"id": uint(161),
	},
	"Manectite": {
		"id": uint(682),
	},
	"Maranga Berry": {
		"id": uint(688),
	},
	"Marshadium Z": {
		"id":   uint(802),
		"type": "ghost",
	},
	"Master Ball": {
		"id": uint(1),
	},
	"Mawilite": {
		"id": uint(681),
	},
	"Max Elixir": {
		"id": uint(41),
	},
	"Max Ether": {
		"id": uint(39),
	},
	"Max Potion": {
		"id": uint(24),
	},
	"Max Repel": {
		"id": uint(77),
	},
	"Max Revive": {
		"id": uint(29),
	},
	"Meadow Plate": {
		"id":   uint(301),
		"type": "grass",
	},
	"Medichamite": {
		"id": uint(665),
	},
	"Mental Herb": {
		"id": uint(219),
	},
	"Metagrossite": {
		"id": uint(758),
	},
	"Metal Coat": {
		"id":   uint(233),
		"type": "steel",
	},
	"Metal Powder": {
		"id": uint(257),
	},
	"Metronome": {
		"id": uint(277),
	},
	"Mewnium Z": {
		"id":   uint(806),
		"type": "psychic",
	},
	"Mewtwonite X": {
		"id": uint(662),
	},
	"Mewtwonite Y": {
		"id": uint(663),
	},
	"Mimikium Z": {
		"id":   uint(924),
		"type": "fairy",
	},
	"Mind Plate": {
		"id":   uint(307),
		"type": "psychic",
	},
	"Miracle Seed": {
		"id":   uint(239),
		"type": "grass",
	},
	"Misty Seed": {
		"id": uint(883),
	},
	"Moomoo Milk": {
		"id": uint(33),
	},
	"Moon Ball": {
		"id": uint(498),
	},
	"Moon Stone": {
		"id": uint(81),
	},
	"Muscle Band": {
		"id": uint(266),
	},
	"Muscle Wing": {
		"id": uint(566),
	},
	"Mystic Water": {
		"id":   uint(243),
		"type": "water",
	},
	"Nest Ball": {
		"id": uint(8),
	},
	"Net Ball": {
		"id": uint(6),
	},
	"Never-Melt Ice": {
		"id":   uint(246),
		"type": "ice",
	},
	"Normal Gem": {
		"id":   uint(564),
		"type": "normal",
	},
	"Normalium Z": {
		"id":   uint(776),
		"type": "normal",
	},
	"Nugget": {
		"id": uint(92),
	},
	"Occa Berry": {
		"id":   uint(184),
		"type": "fire",
	},
	"Odd Incense": {
		"id":   uint(314),
		"type": "psychic",
	},
	"Old Gateau": {
		"id": uint(54),
	},
	"Oran Berry": {
		"id": uint(155),
	},
	"Oval Stone": {
		"id": uint(110),
	},
	"PP Max": {
		"id": uint(53),
	},
	"PP Up": {
		"id": uint(51),
	},
	"Paralyze Heal": {
		"id": uint(22),
	},
	"Passho Berry": {
		"id":   uint(185),
		"type": "water",
	},
	"Payapa Berry": {
		"id":   uint(193),
		"type": "psychic",
	},
	"Pearl": {
		"id": uint(88),
	},
	"Pearl String": {
		"id": uint(582),
	},
	"Pecha Berry": {
		"id": uint(151),
	},
	"Persim Berry": {
		"id": uint(156),
	},
	"Petaya Berry": {
		"id": uint(204),
	},
	"Pidgeotite": {
		"id": uint(762),
	},
	"Pikanium Z": {
		"id":   uint(794),
		"type": "electric",
	},
	"Pikashunium Z": {
		"id":   uint(835),
		"type": "electric",
	},
	"Pinap Berry": {
		"id": uint(168),
	},
	"Pink Nectar": {
		"id": uint(855),
	},
	"Pinsirite": {
		"id": uint(671),
	},
	"Pixie Plate": {
		"id":   uint(644),
		"type": "fairy",
	},
	"Plume Fossil": {
		"id": uint(573),
	},
	"Poison Barb": {
		"id":   uint(245),
		"type": "poison",
	},
	"Poison Gem": {
		"id":   uint(554),
		"type": "poison",
	},
	"Poison Memory": {
		"id":   uint(906),
		"type": "poison",
	},
	"Poisonium Z": {
		"id":   uint(783),
		"type": "poison",
	},
	"Pok\u00e9 Ball": {
		"id": uint(4),
	},
	"Pok\u00e9 Doll": {
		"id": uint(63),
	},
	"Pok\u00e9 Toy": {
		"id": uint(577),
	},
	"Pomeg Berry": {
		"id": uint(169),
	},
	"Potion": {
		"id": uint(17),
	},
	"Power Anklet": {
		"id": uint(293),
	},
	"Power Band": {
		"id": uint(292),
	},
	"Power Belt": {
		"id": uint(290),
	},
	"Power Bracer": {
		"id": uint(289),
	},
	"Power Herb": {
		"id": uint(271),
	},
	"Power Lens": {
		"id": uint(291),
	},
	"Power Weight": {
		"id": uint(294),
	},
	"Premier Ball": {
		"id": uint(12),
	},
	"Pretty Wing": {
		"id": uint(571),
	},
	"Primarium Z": {
		"id":   uint(800),
		"type": "water",
	},
	"Prism Scale": {
		"id": uint(537),
	},
	"Protective Pads": {
		"id": uint(880),
	},
	"Protector": {
		"id": uint(321),
	},
	"Protein": {
		"id": uint(46),
	},
	"Psychic Gem": {
		"id":   uint(557),
		"type": "psychic",
	},
	"Psychic Memory": {
		"id":   uint(916),
		"type": "psychic",
	},
	"Psychic Seed": {
		"id": uint(882),
	},
	"Psychium Z": {
		"id":   uint(786),
		"type": "psychic",
	},
	"Pure Incense": {
		"id": uint(320),
	},
	"Purple Nectar": {
		"id": uint(856),
	},
	"Qualot Berry": {
		"id": uint(171),
	},
	"Quick Ball": {
		"id": uint(15),
	},
	"Quick Claw": {
		"id": uint(217),
	},
	"Quick Powder": {
		"id": uint(274),
	},
	"Rage Candy Bar": {
		"id": uint(504),
	},
	"Rare Bone": {
		"id": uint(106),
	},
	"Rare Candy": {
		"id": uint(50),
	},
	"Rawst Berry": {
		"id": uint(152),
	},
	"Razor Claw": {
		"id": uint(326),
	},
	"Razor Fang": {
		"id": uint(327),
	},
	"Reaper Cloth": {
		"id": uint(325),
	},
	"Red Card": {
		"id": uint(542),
	},
	"Red Nectar": {
		"id": uint(853),
	},
	"Red Orb": {
		"id": uint(534),
	},
	"Red Shard": {
		"id": uint(72),
	},
	"Repeat Ball": {
		"id": uint(9),
	},
	"Repel": {
		"id": uint(79),
	},
	"Resist Wing": {
		"id": uint(567),
	},
	"Revival Herb": {
		"id": uint(37),
	},
	"Revive": {
		"id": uint(28),
	},
	"Rindo Berry": {
		"id":   uint(187),
		"type": "grass",
	},
	"Ring Target": {
		"id": uint(543),
	},
	"Rock Gem": {
		"id":   uint(559),
		"type": "rock",
	},
	"Rock Incense": {
		"id":   uint(315),
		"type": "rock",
	},
	"Rock Memory": {
		"id":   uint(908),
		"type": "rock",
	},
	"Rockium Z": {
		"id":   uint(788),
		"type": "rock",
	},
	"Rocky Helmet": {
		"id": uint(540),
	},
	"Rose Incense": {
		"id":   uint(318),
		"type": "grass",
	},
	"Roseli Berry": {
		"id":   uint(686),
		"type": "fairy",
	},
	"Sablenite": {
		"id": uint(754),
	},
	"Sachet": {
		"id": uint(647),
	},
	"Sacred Ash": {
		"id": uint(44),
	},
	"Safety Goggles": {
		"id": uint(650),
	},
	"Salac Berry": {
		"id": uint(203),
	},
	"Salamencite": {
		"id": uint(769),
	},
	"Sceptilite": {
		"id": uint(753),
	},
	"Scizorite": {
		"id": uint(670),
	},
	"Scope Lens": {
		"id": uint(232),
	},
	"Sea Incense": {
		"id":   uint(254),
		"type": "water",
	},
	"Shalour Sable": {
		"id": uint(709),
	},
	"Sharp Beak": {
		"id":   uint(244),
		"type": "flying",
	},
	"Sharpedonite": {
		"id": uint(759),
	},
	"Shed Shell": {
		"id": uint(295),
	},
	"Shell Bell": {
		"id": uint(253),
	},
	"Shiny Stone": {
		"id": uint(107),
	},
	"Shock Drive": {
		"id":   uint(117),
		"type": "electric",
	},
	"Shuca Berry": {
		"id":   uint(191),
		"type": "ground",
	},
	"Silk Scarf": {
		"id":   uint(251),
		"type": "normal",
	},
	"SilverPowder": {
		"id":   uint(222),
		"type": "bug",
	},
	"Silver Powder": {
		"id":   uint(222),
		"type": "bug",
	},
	"Sitrus Berry": {
		"id": uint(158),
	},
	"Skull Fossil": {
		"id": uint(105),
	},
	"Sky Plate": {
		"id":   uint(306),
		"type": "flying",
	},
	"Slowbronite": {
		"id": uint(760),
	},
	"Smoke Ball": {
		"id": uint(228),
	},
	"Smooth Rock": {
		"id":   uint(283),
		"type": "rock",
	},
	"Snorlium Z": {
		"id":   uint(804),
		"type": "normal",
	},
	"Snowball": {
		"id": uint(649),
	},
	"Soda Pop": {
		"id": uint(31),
	},
	"Soft Sand": {
		"id":   uint(237),
		"type": "ground",
	},
	"Solganium Z": {
		"id":   uint(921),
		"type": "steel",
	},
	"Soothe Bell": {
		"id": uint(218),
	},
	"Soul Dew": {
		"id": uint(225),
	},
	"Spell Tag": {
		"id":   uint(247),
		"type": "ghost",
	},
	"Splash Plate": {
		"id":   uint(299),
		"type": "water",
	},
	"Spooky Plate": {
		"id":   uint(310),
		"type": "ghost",
	},
	"Star Piece": {
		"id": uint(91),
	},
	"Stardust": {
		"id": uint(90),
	},
	"Starf Berry": {
		"id": uint(207),
	},
	"Steel Gem": {
		"id":   uint(563),
		"type": "steel",
	},
	"Steel Memory": {
		"id":   uint(911),
		"type": "steel",
	},
	"Steelium Z": {
		"id":   uint(792),
		"type": "steel",
	},
	"Steelixite": {
		"id": uint(761),
	},
	"Stick": {
		"id": uint(259),
	},
	"Sticky Barb": {
		"id": uint(288),
	},
	"Stone Plate": {
		"id":   uint(309),
		"type": "rock",
	},
	"Strange Souvenir": {
		"id": uint(704),
	},
	"Sun Stone": {
		"id": uint(80),
	},
	"Super Potion": {
		"id": uint(26),
	},
	"Super Repel": {
		"id": uint(76),
	},
	"Swampertite": {
		"id": uint(752),
	},
	"Sweet Heart": {
		"id": uint(134),
	},
	"Swift Wing": {
		"id": uint(570),
	},
	"Tamato Berry": {
		"id": uint(174),
	},
	"Tanga Berry": {
		"id":   uint(194),
		"type": "bug",
	},
	"Tapunium Z": {
		"id":   uint(801),
		"type": "fairy",
	},
	"Terrain Extender": {
		"id": uint(879),
	},
	"Thick Club": {
		"id": uint(258),
	},
	"Thunder Stone": {
		"id": uint(83),
	},
	"Timer Ball": {
		"id": uint(10),
	},
	"Tiny Mushroom": {
		"id": uint(86),
	},
	"Toxic Orb": {
		"id": uint(272),
	},
	"Toxic Plate": {
		"id":   uint(304),
		"type": "poison",
	},
	"Twisted Spoon": {
		"id":   uint(248),
		"type": "psychic",
	},
	"Tyranitarite": {
		"id": uint(669),
	},
	"Ultra Ball": {
		"id": uint(2),
	},
	"Ultranecrozium Z": {
		"id":   uint(923),
		"type": "psychic",
	},
	"Up-Grade": {
		"id": uint(252),
	},
	"Venusaurite": {
		"id": uint(659),
	},
	"Wacan Berry": {
		"id":   uint(186),
		"type": "electric",
	},
	"Water Gem": {
		"id":   uint(549),
		"type": "water",
	},
	"Water Memory": {
		"id":   uint(913),
		"type": "water",
	},
	"Water Stone": {
		"id": uint(84),
	},
	"Waterium Z": {
		"id":   uint(778),
		"type": "water",
	},
	"Wave Incense": {
		"id":   uint(317),
		"type": "water",
	},
	"Weakness Policy": {
		"id": uint(639),
	},
	"Whipped Dream": {
		"id": uint(646),
	},
	"White Herb": {
		"id": uint(214),
	},
	"Wide Lens": {
		"id": uint(265),
	},
	"Wiki Berry": {
		"id": uint(160),
	},
	"Wise Glasses": {
		"id": uint(267),
	},
	"X Accuracy": {
		"id": uint(60),
	},
	"X Attack": {
		"id": uint(57),
	},
	"X Defense": {
		"id": uint(58),
	},
	"X Sp. Atk": {
		"id": uint(61),
	},
	"X Sp. Def": {
		"id": uint(62),
	},
	"X Speed": {
		"id": uint(59),
	},
	"Yache Berry": {
		"id":   uint(188),
		"type": "ice",
	},
	"Yellow Nectar": {
		"id": uint(854),
	},
	"Yellow Shard": {
		"id": uint(74),
	},
	"Zap Plate": {
		"id":   uint(300),
		"type": "electric",
	},
	"Zinc": {
		"id": uint(52),
	},
	"Zoom Lens": {
		"id": uint(276),
	},
	"S Tier": {
		"id": uint(1),
	},
	"A Tier": {
		"id": uint(2),
	},
	"B Tier": {
		"id": uint(3),
	},
	"C Tier": {
		"id": uint(4),
	},
}
