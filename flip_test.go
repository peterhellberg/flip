package flip

import "testing"

var flippersTests = []struct {
	name string
	in   string
	out  string
}{
	{"table", "foo", "(╯°□°）╯︵ooɟ"},
	{"gopher", "bar", "ʕ╯◔ϖ◔ʔ╯︵ɹɐq"},
}

func TestFlippers(t *testing.T) {
	for _, tt := range flippersTests {
		if got := Flippers[tt.name](tt.in); got != tt.out {
			t.Errorf("Flippers[%#v](%#v) = %v, want %v", tt.name, tt.in, got, tt.out)
		}
	}
}

var upsideDownTests = []struct {
	in  string
	out string
}{
	{"ab", "qɐ"},
	{"abcde", "ǝpɔqɐ"},
	{"abcdefghijk", "ʞɾıɥƃɟǝpɔqɐ"},
	{"abcdefghijklmnopqrstuvwxyz", "zʎxʍʌnʇsɹbdouɯןʞɾıɥƃɟǝpɔqɐ"},
}

func TestUpsideDown(t *testing.T) {
	for _, tt := range upsideDownTests {
		if got := UpsideDown(tt.in); got != tt.out {
			t.Errorf("UpsideDown(in) = %v, want %v", got, tt.out)
		}
	}
}

var tableTests = []struct {
	out string
	in  string
}{
	{"(╯°□°）╯︵q∀", "AB"},
	{"(╯°□°）╯︵ɾǝɥ", "hej"},
	{"(╯°□°）╯︵ʇxǝʇ", "text"},
	{"(╯°□°）╯︵ƃuıɹʇs ɹǝƃuoן ∀", "A longer string"},
}

func TestTable(t *testing.T) {
	for _, tt := range tableTests {
		if got := Table(tt.in); got != tt.out {
			t.Errorf("Table(in) = %v, want %v", got, tt.out)
		}
	}
}

var gopherTests = []struct {
	out string
	in  string
}{
	{"ʕ╯◔ϖ◔ʔ╯︵q∀", "AB"},
	{"ʕ╯◔ϖ◔ʔ╯︵ɾǝɥ", "hej"},
	{"ʕ╯◔ϖ◔ʔ╯︵ʇxǝʇ", "text"},
	{"ʕ╯◔ϖ◔ʔ╯︵ƃuıɹʇs ɹǝƃuoן ∀", "A longer string"},
}

func TestGopher(t *testing.T) {
	for _, tt := range gopherTests {
		if got := Gopher(tt.in); got != tt.out {
			t.Errorf("Gopher(in) = %v, want %v", got, tt.out)
		}
	}
}

var reverseTests = []struct {
	in  string
	out string
}{
	{"let", "tel"},
	{"open", "nepo"},
	{"enough", "hguone"},
	{"side", "edis"},
	{"case", "esac"},
	{"days", "syad"},
	{"yet", "tey"},
	{"better", "retteb"},
	{"nothing", "gnihton"},
	{"tell", "llet"},
	{"problem", "melborp"},
	{"toward", "drawot"},
	{"given", "nevig"},
	{"why", "yhw"},
	{"national", "lanoitan"},
	{"room", "moor"},
	{"young", "gnuoy"},
	{"social", "laicos"},
	{"light", "thgil"},
	{"business", "ssenisub"},
	{"president", "tnediserp"},
	{"help", "pleh"},
	{"power", "rewop"},
	{"country", "yrtnuoc"},
	{"next", "txen"},
	{"things", "sgniht"},
	{"word", "drow"},
	{"looked", "dekool"},
	{"real", "laer"},
	{"John", "nhoJ"},
}

func TestReverse(t *testing.T) {
	for _, tt := range reverseTests {
		if got := Reverse(tt.in); got != tt.out {
			t.Errorf("Reverse(in) = %v, want %v", got, tt.out)
		}
	}
}
