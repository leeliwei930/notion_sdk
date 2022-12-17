package enums

import "encoding/json"

type NumberFormatType uint8

const (
	Numeric NumberFormatType = iota
	NumberWithCommas
	Percent
	Dollar
	CanadianDollar
	Euro
	Pound
	Yen
	Ruble
	Rupee
	Won
	Yuan
	Real
	Lira
	Rupiah
	Franc
	HongKongDollar
	NewZealandDollar
	Krona
	NorwegianKrone
	MexicanPeso
	Rand
	NewTaiwanDollar
	DanishKrone
	Zloty
	Baht
	Forint
	Koruna
	Shekel
	ChileanPeso
	PhilippinePeso
	Dirham
	ColombianPeso
	Riyal
	Ringgit
	Leu
	ArgentinePeso
	UruguayanPeso
	SingaporeDollar
)

var numberFormatTypeMap = map[NumberFormatType]string{
	Numeric:          "number",
	NumberWithCommas: "number_with_commas",
	Percent:          "percent",
	Dollar:           "dollar",
	CanadianDollar:   "canadian_dollar",
	Euro:             "euro",
	Pound:            "pound",
	Yen:              "yen",
	Ruble:            "ruble",
	Rupee:            "rupee",
	Won:              "won",
	Yuan:             "yuan",
	Real:             "real",
	Lira:             "lira",
	Rupiah:           "rupiah",
	Franc:            "franc",
	HongKongDollar:   "hong_kong_dollar",
	NewZealandDollar: "new_zealand_dollar",
	Krona:            "krona",
	NorwegianKrone:   "norwegian_krone",
	MexicanPeso:      "mexican_peso",
	Rand:             "rand",
	NewTaiwanDollar:  "new_taiwan_dollar",
	DanishKrone:      "danish_krone",
	Zloty:            "zloty",
	Baht:             "baht",
	Forint:           "forint",
	Koruna:           "koruna",
	Shekel:           "shekel",
	ChileanPeso:      "chilean_peso",
	PhilippinePeso:   "philippine_peso",
	Dirham:           "dirham",
	ColombianPeso:    "colombian_peso",
	Riyal:            "riyal",
	Ringgit:          "ringgit",
	Leu:              "leu",
	ArgentinePeso:    "argentine_peso",
	UruguayanPeso:    "uruguayan_peso",
	SingaporeDollar:  "singapore_dollar",
}
var numberFormatTypeIndexes = map[string]NumberFormatType{
	"number":             Numeric,
	"number_with_commas": NumberWithCommas,
	"percent":            Percent,
	"dollar":             Dollar,
	"canadian_dollar":    CanadianDollar,
	"euro":               Euro,
	"pound":              Pound,
	"yen":                Yen,
	"ruble":              Ruble,
	"rupee":              Rupee,
	"won":                Won,
	"yuan":               Yuan,
	"real":               Real,
	"lira":               Lira,
	"rupiah":             Rupiah,
	"franc":              Franc,
	"hong_kong_dollar":   HongKongDollar,
	"new_zealand_dollar": NewZealandDollar,
	"krona":              Krona,
	"norwegian_krone":    NorwegianKrone,
	"mexican_peso":       MexicanPeso,
	"rand":               Rand,
	"new_taiwan_dollar":  NewTaiwanDollar,
	"danish_krone":       DanishKrone,
	"zloty":              Zloty,
	"baht":               Baht,
	"forint":             Forint,
	"koruna":             Koruna,
	"shekel":             Shekel,
	"chilean_peso":       ChileanPeso,
	"philippine_peso":    PhilippinePeso,
	"dirham":             Dirham,
	"colombian_peso":     ColombianPeso,
	"riyal":              Riyal,
	"ringgit":            Ringgit,
	"leu":                Leu,
	"argentine_peso":     ArgentinePeso,
	"uruguayan_peso":     UruguayanPeso,
	"singapore_dollar":   SingaporeDollar,
}

func ParseNumberFormat(s string) NumberFormatType {
	return numberFormatTypeIndexes[s]
}
func (n NumberFormatType) String() string {
	return numberFormatTypeMap[n]
}

func (n NumberFormatType) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.String())
}

func (n *NumberFormatType) UnmarshalJSON(input []byte) error {
	var j string
	err := json.Unmarshal(input, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*n = ParseNumberFormat(j)
	return nil
}
