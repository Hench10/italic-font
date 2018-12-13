package italic

import (
	"fmt"
	"strings"
)

var space = map[int]string{1: "   ", 2: "   ", 3: "   ", 4: "   ", 5: "   ", 6: "   "}
var mapList = map[string]map[int]string{
	" ": {1: "   ", 2: "   ", 3: "   ", 4: "   ", 5: "   ", 6: "   "},
	"a": {1: "       ", 2: "____   ", 3: "/ __ \\ ", 4: "/ /_/ /", 5: " \\__/_/", 6: "       "},
	"b": {1: " __    ", 2: "/ /_   ", 3: "/ __ \\ ", 4: "/ /_/ /", 5: " \\____/", 6: "       "},
	"c": {1: "       ", 2: "____   ", 3: "/ ___\\ ", 4: "/ /___ ", 5: " \\____/", 6: "       "},
	"d": {1: "     __", 2: "____/ /", 3: "/ __  /", 4: "/ /_/ /", 5: " \\__/_/", 6: "       "},
	"e": {1: "       ", 2: "____   ", 3: "/ __ \\ ", 4: "/  ___/", 5: " \\____/", 6: "       "},
	"f": {1: "      ", 2: "_____ ", 3: "/ ___/", 4: "/ /__ ", 5: "/ ___/", 6: "/_/   "},
	"g": {1: "       ", 2: "____   ", 3: "/ __ \\ ", 4: " \\/_/ /", 5: "/\\__/ /", 6: " \\____/"},
	"h": {1: " __    ", 2: "/ /_   ", 3: "/ __ \\ ", 4: "/ / / /", 5: "/_/ /_/", 6: "       "},
	"i": {1: "  __   ", 2: "_/_/   ", 3: "/_  /  ", 4: "  / /_ ", 5: "  /___/", 6: "       "},
	"j": {1: "   __ ", 2: " _/_/ ", 3: " /_  /", 4: "/\\_/ /", 5: " \\___/", 6: "      "},
	"k": {1: " __    ", 2: "/ /__  ", 3: "/  __/ ", 4: "/ / \\  ", 5: "/_/ \\_\\", 6: "       "},
	"l": {1: " __   ", 2: "/ /   ", 3: "/ /   ", 4: "/ /__ ", 5: " \\___/", 6: "      "},
	"m": {1: "         ", 2: "______   ", 3: "/ ____ \\ ", 4: "/ / / / /", 5: "/_/ / /_/", 6: "         "},
	"n": {1: "       ", 2: "____   ", 3: "/ __ \\ ", 4: "/ / / /", 5: "/_/ /_/", 6: "       "},
	"o": {1: "       ", 2: "____   ", 3: "/ __ \\ ", 4: "/ /_/ /", 5: " \\____/", 6: "       "},
	"p": {1: "       ", 2: "____   ", 3: "/ __ \\ ", 4: "/ /_/ /", 5: "/ ____/", 6: "/_/    "},
	"q": {1: "        ", 2: "____    ", 3: "/  _ \\  ", 4: "/ /_/ / ", 5: " \\__/ / ", 6: "    /__/"},
	"r": {1: "      ", 2: "  ___ ", 3: "/\\/ _/", 4: " |  / ", 5: "  |_/ ", 6: "      "},
	"s": {1: "       ", 2: "____   ", 3: "/ __ \\ ", 4: "/\\__ \\/", 5: " \\____/", 6: "       "},
	"t": {1: "   __  ", 2: "__/ /_ ", 3: "/_  __/", 4: "  / /_ ", 5: "  /___/", 6: "       "},
	"u": {1: "       ", 2: "__  __ ", 3: "/ / / /", 4: "/ /_/ /", 5: " \\____/", 6: "       "},
	"v": {1: "      ", 2: "_  __ ", 3: "| |/ /", 4: " | / /", 5: "  |__/", 6: "      "},
	"w": {1: "        ", 2: "_    __ ", 3: "| |/\\/ /", 4: " | / / /", 5: "  |____/", 6: "        "},
	"x": {1: "        ", 2: "_  __   ", 3: "\\ \\/ /  ", 4: "  /  /  ", 5: "  /_/\\_\\", 6: "        "},
	"y": {1: "       ", 2: "__  __ ", 3: "/ / / /", 4: " \\/_/ /", 5: "/\\__/ /", 6: " \\____/"},
	"z": {1: "         ", 2: "_____    ", 3: "/__  /   ", 4: "   / /__ ", 5: "   /____/", 6: "         "},
}

func FormatPrint(s string) {
	s = strings.ToLower(strings.TrimSpace(s))
	l := len(s)
	var last = map[int]string{1: " ", 2: " ", 3: " ", 4: " ", 5: " ", 6: " "}
	var result = map[int]string{1: "    ", 2: "    ", 3: "   ", 4: "  ", 5: " ", 6: ""}
	var i int
	for i := 0; i < l; i++ {
		ch := string(s[i]) // 依据下标取字符串中的字符，类型为byte
		m := mapList[ch]
		mixStr(&last, &result, &m, i == l-1)
	}
	for i = 1; i <= 6; i++ {
		fmt.Println(result[i])
	}
}

func mixStr(last *map[int]string, result *map[int]string, apd *map[int]string, lst bool) {
	var l = len((*apd)[1])
	if l == 0 {
		apd = &space
		l = len((*apd)[1])
	}
	for i := 1; i <= 6; i++ {
		if (*last)[i] != " " {
			(*result)[i] += (*last)[i]
			(*result)[i] += (*apd)[i][1 : l-1]
		} else {
			(*result)[i] += (*apd)[i][0 : l-1]
		}
		(*last)[i] = (*apd)[i][l-1 : l]
		if lst {
			(*result)[i] += (*last)[i]
		}
	}
}

func FmtFormatPrint(s string) {
	s = strings.ToLower(strings.TrimSpace(s))
	l := len(s)
	var last = map[int]string{1: "  ", 2: "  ", 3: "  ", 4: "  ", 5: "  ", 6: "  "}
	var result = map[int]string{1: "    ", 2: "    ", 3: "   ", 4: "  ", 5: " ", 6: ""}
	for i := 0; i < l; i++ {
		ch := string(s[i]) // 依据下标取字符串中的字符，类型为byte
		m := mapList[ch]
		fmtMixStr(&last, &result, &m, i == l-1)
	}
	for i := 1; i <= 6; i++ {
		fmt.Println(result[i])
	}
}

func fmtMixStr(last *map[int]string, result *map[int]string, apd *map[int]string, lst bool) {
	var l = len((*apd)[1])
	if l == 0 {
		apd = &space
		l = len((*apd)[1])
	}

	var two = true
	var is_space = false
	space_byte := " "[0]
	for i := 1; i <= 6; i++ {
		if len((*apd)[i]) < 4 {
			two = false
			is_space = true
			break
		}
		// 一版：缩进同符号可重叠
		// if ((*last)[i][0] != space_byte &&  (*apd)[i][0] != space_byte && (*last)[i][0] != (*apd)[i][0]) ||
		// 	((*last)[i][1] != space_byte &&  (*apd)[i][1] != space_byte && (*last)[i][1] != (*apd)[i][1]) {
		// 	two = false
		// 	break
		// }
		// 二版：缩进不重叠
		if ((*last)[i][0] != space_byte && (*apd)[i][0] != space_byte) ||
			((*last)[i][1] != space_byte && (*apd)[i][1] != space_byte) {
			two = false
			break
		}
	}

	if two {
		for i := 1; i <= 6; i++ {
			if (*last)[i][0] != space_byte {
				(*result)[i] += (*last)[i][0:1]
			} else {
				(*result)[i] += (*apd)[i][0:1]
			}
			if (*last)[i][1] != space_byte {
				(*result)[i] += (*last)[i][1:2]
			} else {
				(*result)[i] += (*apd)[i][1:2]
			}
			(*result)[i] += (*apd)[i][2 : l-2]
			(*last)[i] = (*apd)[i][l-2 : l]
			if lst {
				(*result)[i] += (*last)[i]
			}
		}
	} else if is_space {
		for i := 1; i <= 6; i++ {
			(*result)[i] += (*last)[i]
			(*result)[i] += " "
			(*last)[i] = "  "
			if lst {
				(*result)[i] += (*last)[i]
			}
		}
	} else {
		for i := 1; i <= 6; i++ {
			if (*last)[i][1] != space_byte {
				(*result)[i] += (*last)[i]
				(*result)[i] += (*apd)[i][1 : l-2]
			} else {
				(*result)[i] += (*last)[i][0:1]
				(*result)[i] += (*apd)[i][0 : l-2]
			}
			(*last)[i] = (*apd)[i][l-2 : l]
			if lst {
				(*result)[i] += (*last)[i]
			}
		}
	}
}
