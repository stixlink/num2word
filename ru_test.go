package num2word

import "testing"

var samples = []struct {
	amount   float64
	upper    bool
	expected string
}{
	{1148.23, true, "Одна тысяча сто сорок восемь 23"},
	{100.21, false, "сто 21"},
	{222, false, "двести двадцать два 00"},
	{569710234.1245, false, "пятьсот шестьдесят девять миллионов семьсот десять тысяч двести тридцать четыре 12"},
}

func Test_RuMoney(t *testing.T) {
	for _, tt := range samples {
		integer, fractionalNumbers := RuMoney(tt.amount, tt.upper)
		if integer+fractionalNumbers != tt.expected {
			t.Errorf("RuMoney(%.2f): expected '%s', got '%s'", tt.amount, tt.expected, integer+fractionalNumbers)
		}
	}
}
