package wtf

import (
	"github.com/gdamore/tcell"
)

var colors = map[string]tcell.Color{
	"aliceblue":            tcell.ColorAliceBlue,
	"antiquewhite":         tcell.ColorAntiqueWhite,
	"aqua":                 tcell.ColorAqua,
	"aquamarine":           tcell.ColorAquaMarine,
	"azure":                tcell.ColorAzure,
	"beige":                tcell.ColorBeige,
	"bisque":               tcell.ColorBisque,
	"black":                tcell.ColorBlack,
	"blanchedalmond":       tcell.ColorBlanchedAlmond,
	"blue":                 tcell.ColorBlue,
	"blueviolet":           tcell.ColorBlueViolet,
	"brown":                tcell.ColorBrown,
	"burlywood":            tcell.ColorBurlyWood,
	"cadetblue":            tcell.ColorCadetBlue,
	"chartreuse":           tcell.ColorChartreuse,
	"chocolate":            tcell.ColorChocolate,
	"coral":                tcell.ColorCoral,
	"cornflowerblue":       tcell.ColorCornflowerBlue,
	"cornsilk":             tcell.ColorCornsilk,
	"crimson":              tcell.ColorCrimson,
	"darkblue":             tcell.ColorDarkBlue,
	"darkcyan":             tcell.ColorDarkCyan,
	"darkgoldenrod":        tcell.ColorDarkGoldenrod,
	"darkgray":             tcell.ColorDarkGray,
	"darkgreen":            tcell.ColorDarkGreen,
	"darkkhaki":            tcell.ColorDarkKhaki,
	"darkmagenta":          tcell.ColorDarkMagenta,
	"darkolivegreen":       tcell.ColorDarkOliveGreen,
	"darkorange":           tcell.ColorDarkOrange,
	"darkorchid":           tcell.ColorDarkOrchid,
	"darkred":              tcell.ColorDarkRed,
	"darksalmon":           tcell.ColorDarkSalmon,
	"darkseagreen":         tcell.ColorDarkSeaGreen,
	"darkslateblue":        tcell.ColorDarkSlateBlue,
	"darkslategray":        tcell.ColorDarkSlateGray,
	"darkturquoise":        tcell.ColorDarkTurquoise,
	"darkviolet":           tcell.ColorDarkViolet,
	"deeppink":             tcell.ColorDeepPink,
	"deepskyblue":          tcell.ColorDeepSkyBlue,
	"dimgray":              tcell.ColorDimGray,
	"dodgerblue":           tcell.ColorDodgerBlue,
	"firebrick":            tcell.ColorFireBrick,
	"floralwhite":          tcell.ColorFloralWhite,
	"forestgreen":          tcell.ColorForestGreen,
	"fuchsia":              tcell.ColorFuchsia,
	"gainsboro":            tcell.ColorGainsboro,
	"ghostwhite":           tcell.ColorGhostWhite,
	"gold":                 tcell.ColorGold,
	"goldenrod":            tcell.ColorGoldenrod,
	"gray":                 tcell.ColorGray,
	"green":                tcell.ColorGreen,
	"greenyellow":          tcell.ColorGreenYellow,
	"grey":                 tcell.ColorGray,
	"honeydew":             tcell.ColorHoneydew,
	"hotpink":              tcell.ColorHotPink,
	"indianred":            tcell.ColorIndianRed,
	"indigo":               tcell.ColorIndigo,
	"ivory":                tcell.ColorIvory,
	"khaki":                tcell.ColorKhaki,
	"lavender":             tcell.ColorLavender,
	"lavenderblush":        tcell.ColorLavenderBlush,
	"lawngreen":            tcell.ColorLawnGreen,
	"lemonchiffon":         tcell.ColorLemonChiffon,
	"lightblue":            tcell.ColorLightBlue,
	"lightcoral":           tcell.ColorLightCoral,
	"lightcyan":            tcell.ColorLightCyan,
	"lightgoldenrodyellow": tcell.ColorLightGoldenrodYellow,
	"lightgray":            tcell.ColorLightGray,
	"lightgreen":           tcell.ColorLightGreen,
	"lightpink":            tcell.ColorLightPink,
	"lightsalmon":          tcell.ColorLightSalmon,
	"lightseagreen":        tcell.ColorLightSeaGreen,
	"lightskyblue":         tcell.ColorLightSkyBlue,
	"lightslategray":       tcell.ColorLightSlateGray,
	"lightsteelblue":       tcell.ColorLightSteelBlue,
	"lightyellow":          tcell.ColorLightYellow,
	"lime":                 tcell.ColorLime,
	"limegreen":            tcell.ColorLimeGreen,
	"linen":                tcell.ColorLinen,
	"maroon":               tcell.ColorMaroon,
	"mediumaquamarine":     tcell.ColorMediumAquamarine,
	"mediumblue":           tcell.ColorMediumBlue,
	"mediumorchid":         tcell.ColorMediumOrchid,
	"mediumpurple":         tcell.ColorMediumPurple,
	"mediumseagreen":       tcell.ColorMediumSeaGreen,
	"mediumslateblue":      tcell.ColorMediumSlateBlue,
	"mediumspringgreen":    tcell.ColorMediumSpringGreen,
	"mediumturquoise":      tcell.ColorMediumTurquoise,
	"mediumvioletred":      tcell.ColorMediumVioletRed,
	"midnightblue":         tcell.ColorMidnightBlue,
	"mintcream":            tcell.ColorMintCream,
	"mistyrose":            tcell.ColorMistyRose,
	"moccasin":             tcell.ColorMoccasin,
	"navajowhite":          tcell.ColorNavajoWhite,
	"navy":                 tcell.ColorNavy,
	"oldlace":              tcell.ColorOldLace,
	"olive":                tcell.ColorOlive,
	"olivedrab":            tcell.ColorOliveDrab,
	"orange":               tcell.ColorOrange,
	"orangered":            tcell.ColorOrangeRed,
	"orchid":               tcell.ColorOrchid,
	"palegoldenrod":        tcell.ColorPaleGoldenrod,
	"palegreen":            tcell.ColorPaleGreen,
	"paleturquoise":        tcell.ColorPaleTurquoise,
	"palevioletred":        tcell.ColorPaleVioletRed,
	"papayawhip":           tcell.ColorPapayaWhip,
	"peachpuff":            tcell.ColorPeachPuff,
	"peru":                 tcell.ColorPeru,
	"pink":                 tcell.ColorPink,
	"plum":                 tcell.ColorPlum,
	"powderblue":           tcell.ColorPowderBlue,
	"purple":               tcell.ColorPurple,
	"rebeccapurple":        tcell.ColorRebeccaPurple,
	"red":                  tcell.ColorRed,
	"rosybrown":            tcell.ColorRosyBrown,
	"royalblue":            tcell.ColorRoyalBlue,
	"saddlebrown":          tcell.ColorSaddleBrown,
	"salmon":               tcell.ColorSalmon,
	"sandybrown":           tcell.ColorSandyBrown,
	"seagreen":             tcell.ColorSeaGreen,
	"seashell":             tcell.ColorSeashell,
	"sienna":               tcell.ColorSienna,
	"silver":               tcell.ColorSilver,
	"skyblue":              tcell.ColorSkyblue,
	"slateblue":            tcell.ColorSlateBlue,
	"slategray":            tcell.ColorSlateGray,
	"snow":                 tcell.ColorSnow,
	"springgreen":          tcell.ColorSpringGreen,
	"steelblue":            tcell.ColorSteelBlue,
	"tan":                  tcell.ColorTan,
	"teal":                 tcell.ColorTeal,
	"thistle":              tcell.ColorThistle,
	"tomato":               tcell.ColorTomato,
	"turquoise":            tcell.ColorTurquoise,
	"violet":               tcell.ColorViolet,
	"wheat":                tcell.ColorWheat,
	"white":                tcell.ColorWhite,
	"whitesmoke":           tcell.ColorWhiteSmoke,
	"yellow":               tcell.ColorYellow,
	"yellowgreen":          tcell.ColorYellowGreen,
}

func ColorFor(label string) tcell.Color {
	if _, ok := colors[label]; ok {
		return colors[label]
	} else {
		return tcell.ColorGreen
	}
}
