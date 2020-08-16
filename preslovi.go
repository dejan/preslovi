package preslovi

import "strings"

var replacer = strings.NewReplacer(
	"а", "a",
	"А", "A",
	"б", "b",
	"Б", "B",
	"в", "v",
	"В", "V",
	"г", "g",
	"Г", "G",
	"д", "d",
	"Д", "D",
	"ђ", "đ",
	"Ђ", "Đ",
	"е", "e",
	"Е", "E",
	"ж", "ž",
	"Ж", "Ž",
	"з", "z",
	"З", "Z",
	"и", "i",
	"И", "I",
	"ј", "j",
	"Ј", "J",
	"к", "k",
	"К", "K",
	"л", "l",
	"Л", "L",
	"љ", "lj",
	"Љ", "LJ",
	"м", "m",
	"М", "M",
	"н", "n",
	"Н", "N",
	"њ", "nj",
	"Њ", "NJ",
	"о", "o",
	"О", "O",
	"п", "p",
	"П", "P",
	"р", "r",
	"Р", "R",
	"с", "s",
	"С", "S",
	"т", "t",
	"Т", "T",
	"ћ", "ć",
	"Ћ", "Ć",
	"ц", "c",
	"Ц", "C",
	"у", "u",
	"У", "U",
	"ф", "f",
	"Ф", "F",
	"х", "h",
	"Х", "H",
	"ч", "č",
	"Ч", "Č",
	"џ", "dž",
	"Џ", "DŽ",
	"ш", "š",
	"Ш", "Š",
)

var asciiReplacer = strings.NewReplacer(
	"а", "a",
	"А", "A",
	"б", "b",
	"Б", "B",
	"в", "v",
	"В", "V",
	"г", "g",
	"Г", "G",
	"д", "d",
	"Д", "D",
	"ђ", "dj",
	"Ђ", "Dj",
	"е", "e",
	"Е", "E",
	"ж", "z",
	"Ж", "Z",
	"з", "z",
	"З", "Z",
	"и", "i",
	"И", "I",
	"ј", "j",
	"Ј", "J",
	"к", "k",
	"К", "K",
	"л", "l",
	"Л", "L",
	"љ", "lj",
	"Љ", "Lj",
	"м", "m",
	"М", "M",
	"н", "n",
	"Н", "N",
	"њ", "nj",
	"Њ", "Nj",
	"о", "o",
	"О", "O",
	"п", "p",
	"П", "P",
	"р", "r",
	"Р", "R",
	"с", "s",
	"С", "S",
	"т", "t",
	"Т", "T",
	"ћ", "c",
	"Ћ", "C",
	"ц", "c",
	"Ц", "C",
	"у", "u",
	"У", "U",
	"ф", "f",
	"Ф", "F",
	"х", "h",
	"Х", "H",
	"ч", "c",
	"Ч", "C",
	"џ", "dz",
	"Џ", "Dz",
	"ш", "s",
	"Ш", "S",
)

// Latinicom converts Serbian cyrillic to latin
func Latinicom(s string) string {
	return replacer.Replace(s)
}

// Latinicom converts Serbian cyrillic to ascii latin
func LatinicomAscii(s string) string {
	return asciiReplacer.Replace(s)
}
