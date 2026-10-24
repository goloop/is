package is

import (
	"os"
	"strings"
	"testing"
)

func TestNewFormatVectors(t *testing.T) {
	tests := []struct {
		name      string
		check     func(string) bool
		good, bad []string
	}{
		{"date", Date, []string{"2024-02-29", "2000-02-29", "9999-12-31"}, []string{"1900-02-29", "2023-02-29", "2024-2-01", "0000-01-01", "2024-01-01Z"}},
		{"timestamp", RFC3339, []string{"2024-02-29T23:59:59Z", "2024-02-29t12:00:00.123456789z", "2024-01-01T00:00:00-00:00"}, []string{"2024-01-01T00:00:00+24:00", "2024-01-01T00:00:00+01:60", "2024-01-01T1:00:00Z", "2024-01-01T00:00:00,1Z", "2024-01-01T00:00:60Z", "2024-02-30T12:00:00Z"}},
		{"offset", UTCOffset, []string{"Z", "+00:00", "-00:00", "+23:59", "-12:00"}, []string{"+24:00", "+01:60", "+1:00", "z", " +01:00"}},
		{"duration", Duration, []string{"0", "-1h30m", "1.5s", "2562047h47m16.854775807s"}, []string{"P1D", "1d", "2562047h47m16.854775808s", ""}},
		{"cidr", CIDR, []string{"192.0.2.1/24", "2001:db8::/32", "0.0.0.0/0"}, []string{"192.0.2.1/33", "2001:db8::/129", "192.0.2.1", "fe80::1%eth0/64"}},
		{"port", Port, []string{"0", "65535", "00080"}, []string{"65536", "-1", "+80", " 80", "８０", ""}},
		{"ipport", IPPort, []string{"192.0.2.1:80", "[2001:db8::1]:65535"}, []string{"example.com:80", "2001:db8::1:80", "[fe80::1%eth0]:80", "192.0.2.1:+80"}},
		{"json", JSON, []string{`{"a":[1,true,null]}`, `"text"`, `42`}, []string{`{"a":}`, `NaN`, `{} {}`, "\"\xff\""}},
		{"mime", MIME, []string{"application/json", "text/plain; charset=utf-8", "application/vnd.example+json"}, []string{"text", "text/", "text/*", " application/json", "text/plain\r\nX: y"}},
		{"semver", SemVer, []string{"1.2.3", "0.0.0-alpha.1+001", "1.0.0-0", "1.0.0-01a", "999999999999999999999.0.0"}, []string{"v1.2.3", "01.2.3", "1.2", "1.2.3-01", "1.2.3+", "1.2.3-a..b"}},
		{"base32", func(s string) bool { return Base32(s) }, []string{"MY======", "MZXW6===", "MZXW6YTB"}, []string{"MZ======", "my======", "MY", "MY======\n", ""}},
		{"base32raw", func(s string) bool { return Base32(s, false) }, []string{"MY", "MZXW6YTB"}, []string{"MY======", "MZ", "A"}},
		{"sha384", SHA384, []string{"38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b"}, []string{strings.Repeat("f", 95), strings.Repeat("g", 96)}},
		{"uuid4", UUIDv4, []string{"919108f7-52d1-4320-9bac-f847db4148a8"}, []string{"919108f7-52d1-7320-9bac-f847db4148a8", "919108f7-52d1-4320-7bac-f847db4148a8"}},
		{"uuid7", UUIDv7, []string{"017f22e2-79b0-7cc3-98c4-dc0c0c07398f"}, []string{"017f22e2-79b0-4cc3-98c4-dc0c0c07398f", "017f22e2-79b0-7cc3-c8c4-dc0c0c07398f"}},
		{"isbn", ISBN, []string{"9780306406157", "978-0-306-40615-7", "0-306-40615-2", "080442957X", "978 0 306 40615 7"}, []string{"9780306406158", "978--0306406157", "-9780306406157", "9780306406157-", "9790060115615", "4006381333931", "080442957x"}},
		{"ean", EAN, []string{"4006381333931", "96385074"}, []string{"4006381333932", "036000291452", "9638507"}},
		{"gtin", GTIN, []string{"036000291452", "10012345678902", "4006381333931"}, []string{"036000291453", "10012345678903", "0000000X"}},
		{"country", CountryCode, []string{"UA", "GB", "US", "AX", "BQ"}, []string{"uk", "UK", "XK", "ZZ", "EU", "USA", ""}},
		{"currency", CurrencyCode, []string{"USD", "UAH", "EUR", "XAU", "XXX", "XTS", "XAD"}, []string{"usd", "BTC", "ZZZ", "BGN", ""}},
		{"bic", BIC, []string{"DEUTDEFF", "NBUAUAUX", "DEUTDEFF500", "A1BCGB2L"}, []string{"DEUTZZFF", "DEUTDEFFF", "deutdeff", "DEUTDEFF!00"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, s := range tt.good {
				if !tt.check(s) {
					t.Errorf("rejected %q", s)
				}
			}
			for _, s := range tt.bad {
				if tt.check(s) {
					t.Errorf("accepted %q", s)
				}
			}
		})
	}
}

func TestPostalCodeVectors(t *testing.T) {
	for cc, code := range map[string]string{"US": "78701-1234", "GB": "SW1A 1AA", "DE": "10115", "PL": "00-001", "UA": "01001", "FR": "75001", "ES": "08001", "IT": "00118", "PT": "1000-001", "NL": "1012 AB", "CZ": "110 00"} {
		if !PostalCode(code, cc) {
			t.Errorf("rejected %s %s", cc, code)
		}
		if PostalCode(code+"!", cc) {
			t.Errorf("accepted corrupt %s", cc)
		}
	}
	for _, v := range [][2]string{{"SW1A1AA", "GB"}, {"QV1 1AA", "GB"}, {"53001", "ES"}, {"0000 AA", "NL"}, {"800 00", "CZ"}, {"12345", "ZZ"}} {
		if PostalCode(v[0], v[1]) {
			t.Errorf("accepted %v", v)
		}
	}
}

func TestSWIFTRegistryExamples(t *testing.T) {
	raw, err := os.ReadFile("testdata/iban_swift_101.txt")
	if err != nil {
		t.Fatal(err)
	}
	examples := strings.Fields(string(raw))
	if len(examples) != 89 {
		t.Fatalf("got %d examples", len(examples))
	}
	for _, s := range examples {
		if !IBAN(s, true) {
			t.Errorf("rejected SWIFT example %s", s)
		}
	}
	// MOD-97 alone cannot detect a numeric bank identifier where letters are required.
	if CalculateIBANChecksum("474663316052549159GB54") != 1 {
		t.Fatal("bad adversarial fixture")
	}
	for _, s := range []string{"GB54474663316052549159", "FK4912345678901234", "NI31123456789012345678901234"} {
		if IBAN(s) {
			t.Errorf("accepted invalid national structure %s", s)
		}
	}
}

func FuzzStructuredValidators(f *testing.F) {
	for _, s := range []string{"", "2024-02-29", "GB82WEST12345698765432", "!", "017f22e2-79b0-7cc3-98c4-dc0c0c07398f"} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, s string) {
		Date(s)
		RFC3339(s)
		UTCOffset(s)
		Duration(s)
		IPPort(s)
		CIDR(s)
		JSON(s)
		MIME(s)
		SemVer(s)
		Base32(s)
		UUIDv4(s)
		UUIDv7(s)
		IBAN(s)
		ISBN(s)
		GTIN(s)
		BIC(s)
	})
}
