// Package num2word holds Number to words converter
package num2word

import (
	"fmt"
	"strings"
	"unicode"
)

var repl = [][]string{
	// t - тысячи; m - милионы; M - миллиарды;
	{",,,,,,", "eM"},
	{",,,,,", "em"},
	{",,,,", "et"},
	// e - единицы; d - десятки; c - сотни;
	{",,,", "e"},
	{",,", "d"},
	{",", "c"},
	{"0c0d0et", ""},
	{"0c0d0em", ""},
	{"0c0d0eM", ""},
	// --
	{"0c", ""},
	{"1c", "сто "},
	{"2c", "двести "},
	{"3c", "триста "},
	{"4c", "четыреста "},
	{"5c", "пятьсот "},
	{"6c", "шестьсот "},
	{"7c", "семьсот "},
	{"8c", "восемьсот "},
	{"9c", "девятьсот "},
	{"1d0e", "десять "},
	{"1d1e", "одиннадцать "},
	{"1d2e", "двенадцать "},
	{"1d3e", "тринадцать "},
	{"1d4e", "четырнадцать "},
	{"1d5e", "пятнадцать "},
	{"1d6e", "шестнадцать "},
	{"1d7e", "семнадцать "},
	{"1d8e", "восемнадцать "},
	{"1d9e", "девятнадцать "},
	// --
	{"0d", ""},
	{"2d", "двадцать "},
	{"3d", "тридцать "},
	{"4d", "сорок "},
	{"5d", "пятьдесят "},
	{"6d", "шестьдесят "},
	{"7d", "семьдесят "},
	{"8d", "восемьдесят "},
	{"9d", "девяносто "},
	// --
	{"0e", ""},
	{"5e", "пять "},
	{"6e", "шесть "},
	{"7e", "семь "},
	{"8e", "восемь "},
	{"9e", "девять "},
	// --
	{"1et", "одна тысяча "},
	{"2et", "две тысячи "},
	{"3et", "три тысячи "},
	{"4et", "четыре тысячи "},
	{"1em", "один миллион "},
	{"2em", "два миллиона "},
	{"3em", "три миллиона "},
	{"4em", "четыре миллиона "},
	{"1eM", "один миллиард "},
	{"2eM", "два миллиарда "},
	{"3eM", "три миллиарда "},
	{"4eM", "четыре миллиарда "},
	{"1e", "один "},
	{"2e", "два "},
	{"3e", "три "},
	{"4e", "четыре "},
	//  блок для написания копеек без сокращения "коп"
	{"11k", "11"},
	{"12k", "12"},
	{"13k", "13"},
	{"14k", "14"},
	{"1k", "1"},
	{"2k", "2"},
	{"3k", "3"},
	{"4k", "4"},
	{"k", ""},
	// --
	{".", ""},
	{"t", "тысяч "},
	{"m", "миллионов "},
	{"M", "миллиардов "},
}

var mask = []string{",,,", ",,", ",", ",,,,", ",,", ",", ",,,,,", ",,", ",", ",,,,,,", ",,", ","}

// RuMoney - деньги прописью на русском
func RuMoney(number float64, upperFirstChar bool) (integer string, fractionalNumbers string) {

	s := fmt.Sprintf("%.2f", number)
	l := len(s)

	dest := s[l-3:l] + "k"
	s = s[:l-3]
	l = len(s)
	for i := l; i > 0; i-- {
		c := string(s[i-1])
		dest = c + mask[l-i] + dest
	}
	destSlice := strings.Split(dest, ".")
	integer = replaceToWord(destSlice[0], upperFirstChar)
	fractionalNumbers = replaceToWord(destSlice[1], upperFirstChar)
	return
}
func replaceToWord(dest string, upperFirstChar bool) string {
	for _, r := range repl {
		dest = strings.Replace(dest, r[0], r[1], -1)
	}
	if upperFirstChar {
		a := []rune(dest)
		a[0] = unicode.ToUpper(a[0])
		dest = string(a)
	}
	return dest
}
