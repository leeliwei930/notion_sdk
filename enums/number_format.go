package enums

type NumberFormatType int

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

var NumberFormatTypeMap = map[NumberFormatType]string{
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
