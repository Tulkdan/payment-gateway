package constants

var lookupTable = map[string]struct{}{
	"AFN": {}, // Afghani
	"971": {}, // Afghani
	"EUR": {}, // Euro
	"978": {}, // Euro
	"ALL": {}, // Lek
	"008": {}, // Lek
	"DZD": {}, // Algerian Dinar
	"012": {}, // Algerian Dinar
	"USD": {}, // US Dollar
	"840": {}, // US Dollar
	"AOA": {}, // Kwanza
	"973": {}, // Kwanza
	"XCD": {}, // East Caribbean Dollar
	"951": {}, // East Caribbean Dollar
	"ARS": {}, // Argentine Peso
	"032": {}, // Argentine Peso
	"AMD": {}, // Armenian Dram
	"051": {}, // Armenian Dram
	"AWG": {}, // Aruban Florin
	"533": {}, // Aruban Florin
	"AUD": {}, // Australian Dollar
	"036": {}, // Australian Dollar
	"AZN": {}, // Azerbaijan Manat
	"944": {}, // Azerbaijan Manat
	"BSD": {}, // Bahamian Dollar
	"044": {}, // Bahamian Dollar
	"BHD": {}, // Bahraini Dinar
	"048": {}, // Bahraini Dinar
	"BDT": {}, // Taka
	"050": {}, // Taka
	"BBD": {}, // Barbados Dollar
	"052": {}, // Barbados Dollar
	"BYN": {}, // Belarusian Ruble
	"933": {}, // Belarusian Ruble
	"BZD": {}, // Belize Dollar
	"084": {}, // Belize Dollar
	"XOF": {}, // CFA Franc BCEAO
	"952": {}, // CFA Franc BCEAO
	"BMD": {}, // Bermudian Dollar
	"060": {}, // Bermudian Dollar
	"INR": {}, // Indian Rupee
	"356": {}, // Indian Rupee
	"BTN": {}, // Ngultrum
	"064": {}, // Ngultrum
	"BOB": {}, // Boliviano
	"068": {}, // Boliviano
	"BOV": {}, // Mvdol
	"984": {}, // Mvdol
	"BAM": {}, // Convertible Mark
	"977": {}, // Convertible Mark
	"BWP": {}, // Pula
	"072": {}, // Pula
	"NOK": {}, // Norwegian Krone
	"578": {}, // Norwegian Krone
	"BRL": {}, // Brazilian Real
	"986": {}, // Brazilian Real
	"BND": {}, // Brunei Dollar
	"096": {}, // Brunei Dollar
	"BGN": {}, // Bulgarian Lev
	"975": {}, // Bulgarian Lev
	"BIF": {}, // Burundi Franc
	"108": {}, // Burundi Franc
	"CVE": {}, // Cabo Verde Escudo
	"132": {}, // Cabo Verde Escudo
	"KHR": {}, // Riel
	"116": {}, // Riel
	"XAF": {}, // CFA Franc BEAC
	"950": {}, // CFA Franc BEAC
	"CAD": {}, // Canadian Dollar
	"124": {}, // Canadian Dollar
	"KYD": {}, // Cayman Islands Dollar
	"136": {}, // Cayman Islands Dollar
	"CLP": {}, // Chilean Peso
	"152": {}, // Chilean Peso
	"CLF": {}, // Unidad de Fomento
	"990": {}, // Unidad de Fomento
	"CNY": {}, // Yuan Renminbi
	"156": {}, // Yuan Renminbi
	"CNH": {}, // Yuan Fen
	"157": {}, // Yuan Fen
	"COP": {}, // Colombian Peso
	"170": {}, // Colombian Peso
	"COU": {}, // Unidad de Valor Real
	"970": {}, // Unidad de Valor Real
	"KMF": {}, // Comorian Franc
	"174": {}, // Comorian Franc
	"CDF": {}, // Congolese Franc
	"976": {}, // Congolese Franc
	"NZD": {}, // New Zealand Dollar
	"554": {}, // New Zealand Dollar
	"CRC": {}, // Costa Rican Colon
	"188": {}, // Costa Rican Colon
	"HRK": {}, // Kuna
	"191": {}, // Kuna
	"CUP": {}, // Cuban Peso
	"192": {}, // Cuban Peso
	"CUC": {}, // Peso Convertible
	"931": {}, // Peso Convertible
	"ANG": {}, // Netherlands Antillean Guilder
	"532": {}, // Netherlands Antillean Guilder
	"CZK": {}, // Czech Koruna
	"203": {}, // Czech Koruna
	"DKK": {}, // Danish Krone
	"208": {}, // Danish Krone
	"DJF": {}, // Djibouti Franc
	"262": {}, // Djibouti Franc
	"DOP": {}, // Dominican Peso
	"214": {}, // Dominican Peso
	"EGP": {}, // Egyptian Pound
	"818": {}, // Egyptian Pound
	"SVC": {}, // El Salvador Colon
	"222": {}, // El Salvador Colon
	"ERN": {}, // Nakfa
	"232": {}, // Nakfa
	"SZL": {}, // Lilangeni
	"748": {}, // Lilangeni
	"ETB": {}, // Ethiopian Birr
	"230": {}, // Ethiopian Birr
	"FKP": {}, // Falkland Islands Pound
	"238": {}, // Falkland Islands Pound
	"FJD": {}, // Fiji Dollar
	"242": {}, // Fiji Dollar
	"XPF": {}, // CFP Franc
	"953": {}, // CFP Franc
	"GMD": {}, // Dalasi
	"270": {}, // Dalasi
	"GEL": {}, // Lari
	"981": {}, // Lari
	"GHS": {}, // Ghana Cedi
	"936": {}, // Ghana Cedi
	"GIP": {}, // Gibraltar Pound
	"292": {}, // Gibraltar Pound
	"GTQ": {}, // Quetzal
	"320": {}, // Quetzal
	"GBP": {}, // Pound Sterling
	"826": {}, // Pound Sterling
	"GNF": {}, // Guinean Franc
	"324": {}, // Guinean Franc
	"GYD": {}, // Guyana Dollar
	"328": {}, // Guyana Dollar
	"HTG": {}, // Gourde
	"332": {}, // Gourde
	"HNL": {}, // Lempira
	"340": {}, // Lempira
	"HKD": {}, // Hong Kong Dollar
	"344": {}, // Hong Kong Dollar
	"HUF": {}, // Forint
	"348": {}, // Forint
	"ISK": {}, // Iceland Krona
	"352": {}, // Iceland Krona
	"IDR": {}, // Rupiah
	"360": {}, // Rupiah
	"XDR": {}, // SDR (Special Drawing Right)
	"960": {}, // SDR (Special Drawing Right)
	"IRR": {}, // Iranian Rial
	"364": {}, // Iranian Rial
	"IQD": {}, // Iraqi Dinar
	"368": {}, // Iraqi Dinar
	"ILS": {}, // New Israeli Sheqel
	"376": {}, // New Israeli Sheqel
	"JMD": {}, // Jamaican Dollar
	"388": {}, // Jamaican Dollar
	"JPY": {}, // Yen
	"392": {}, // Yen
	"JOD": {}, // Jordanian Dinar
	"400": {}, // Jordanian Dinar
	"KZT": {}, // Tenge
	"398": {}, // Tenge
	"KES": {}, // Kenyan Shilling
	"404": {}, // Kenyan Shilling
	"KPW": {}, // North Korean Won
	"408": {}, // North Korean Won
	"KRW": {}, // Won
	"410": {}, // Won
	"KWD": {}, // Kuwaiti Dinar
	"414": {}, // Kuwaiti Dinar
	"KGS": {}, // Som
	"417": {}, // Som
	"LAK": {}, // Lao Kip
	"418": {}, // Lao Kip
	"LBP": {}, // Lebanese Pound
	"422": {}, // Lebanese Pound
	"LSL": {}, // Loti
	"426": {}, // Loti
	"ZAR": {}, // Rand
	"710": {}, // Rand
	"LRD": {}, // Liberian Dollar
	"430": {}, // Liberian Dollar
	"LYD": {}, // Libyan Dinar
	"434": {}, // Libyan Dinar
	"CHF": {}, // Swiss Franc
	"756": {}, // Swiss Franc
	"MOP": {}, // Pataca
	"446": {}, // Pataca
	"MKD": {}, // Denar
	"807": {}, // Denar
	"MGA": {}, // Malagasy Ariary
	"969": {}, // Malagasy Ariary
	"MWK": {}, // Malawi Kwacha
	"454": {}, // Malawi Kwacha
	"MYR": {}, // Malaysian Ringgit
	"458": {}, // Malaysian Ringgit
	"MVR": {}, // Rufiyaa
	"462": {}, // Rufiyaa
	"MRU": {}, // Ouguiya
	"929": {}, // Ouguiya
	"MUR": {}, // Mauritius Rupee
	"480": {}, // Mauritius Rupee
	"XUA": {}, // ADB Unit of Account
	"965": {}, // ADB Unit of Account
	"MXN": {}, // Mexican Peso
	"484": {}, // Mexican Peso
	"MXV": {}, // Mexican Unidad de Inversion (UDI)
	"979": {}, // Mexican Unidad de Inversion (UDI)
	"MDL": {}, // Moldovan Leu
	"498": {}, // Moldovan Leu
	"MNT": {}, // Tugrik
	"496": {}, // Tugrik
	"MAD": {}, // Moroccan Dirham
	"504": {}, // Moroccan Dirham
	"MZN": {}, // Mozambique Metical
	"943": {}, // Mozambique Metical
	"MMK": {}, // Kyat
	"104": {}, // Kyat
	"NAD": {}, // Namibia Dollar
	"516": {}, // Namibia Dollar
	"NPR": {}, // Nepalese Rupee
	"524": {}, // Nepalese Rupee
	"NIO": {}, // Cordoba Oro
	"558": {}, // Cordoba Oro
	"NGN": {}, // Naira
	"566": {}, // Naira
	"OMR": {}, // Rial Omani
	"512": {}, // Rial Omani
	"PKR": {}, // Pakistan Rupee
	"586": {}, // Pakistan Rupee
	"PAB": {}, // Balboa
	"590": {}, // Balboa
	"PGK": {}, // Kina
	"598": {}, // Kina
	"PYG": {}, // Guarani
	"600": {}, // Guarani
	"PEN": {}, // Sol
	"604": {}, // Sol
	"PHP": {}, // Philippine Peso
	"608": {}, // Philippine Peso
	"PLN": {}, // Zloty
	"985": {}, // Zloty
	"QAR": {}, // Qatari Rial
	"634": {}, // Qatari Rial
	"RON": {}, // Romanian Leu
	"946": {}, // Romanian Leu
	"RUB": {}, // Russian Ruble
	"643": {}, // Russian Ruble
	"RWF": {}, // Rwanda Franc
	"646": {}, // Rwanda Franc
	"SHP": {}, // Saint Helena Pound
	"654": {}, // Saint Helena Pound
	"WST": {}, // Tala
	"882": {}, // Tala
	"STN": {}, // Dobra
	"930": {}, // Dobra
	"SAR": {}, // Saudi Riyal
	"682": {}, // Saudi Riyal
	"RSD": {}, // Serbian Dinar
	"941": {}, // Serbian Dinar
	"SCR": {}, // Seychelles Rupee
	"690": {}, // Seychelles Rupee
	"SLL": {}, // Leone
	"694": {}, // Leone
	"SGD": {}, // Singapore Dollar
	"702": {}, // Singapore Dollar
	"XSU": {}, // Sucre
	"994": {}, // Sucre
	"SBD": {}, // Solomon Islands Dollar
	"090": {}, // Solomon Islands Dollar
	"SOS": {}, // Somali Shilling
	"706": {}, // Somali Shilling
	"SSP": {}, // South Sudanese Pound
	"728": {}, // South Sudanese Pound
	"LKR": {}, // Sri Lanka Rupee
	"144": {}, // Sri Lanka Rupee
	"SDG": {}, // Sudanese Pound
	"938": {}, // Sudanese Pound
	"SRD": {}, // Surinam Dollar
	"968": {}, // Surinam Dollar
	"SEK": {}, // Swedish Krona
	"752": {}, // Swedish Krona
	"CHE": {}, // WIR Euro
	"947": {}, // WIR Euro
	"CHW": {}, // WIR Franc
	"948": {}, // WIR Franc
	"SYP": {}, // Syrian Pound
	"760": {}, // Syrian Pound
	"TWD": {}, // New Taiwan Dollar
	"901": {}, // New Taiwan Dollar
	"TJS": {}, // Somoni
	"972": {}, // Somoni
	"TZS": {}, // Tanzanian Shilling
	"834": {}, // Tanzanian Shilling
	"THB": {}, // Baht
	"764": {}, // Baht
	"TOP": {}, // Pa'anga
	"776": {}, // Pa'anga
	"TTD": {}, // Trinidad and Tobago Dollar
	"780": {}, // Trinidad and Tobago Dollar
	"TND": {}, // Tunisian Dinar
	"788": {}, // Tunisian Dinar
	"TRY": {}, // Turkish Lira
	"949": {}, // Turkish Lira
	"TMT": {}, // Turkmenistan New Manat
	"934": {}, // Turkmenistan New Manat
	"UGX": {}, // Uganda Shilling
	"800": {}, // Uganda Shilling
	"UAH": {}, // Hryvnia
	"980": {}, // Hryvnia
	"AED": {}, // UAE Dirham
	"784": {}, // UAE Dirham
	"USN": {}, // US Dollar (Next day)
	"997": {}, // US Dollar (Next day)
	"UYU": {}, // Peso Uruguayo
	"858": {}, // Peso Uruguayo
	"UYI": {}, // Uruguay Peso en Unidades Indexadas (UI)
	"940": {}, // Uruguay Peso en Unidades Indexadas (UI)
	"UYW": {}, // Unidad Previsional
	"927": {}, // Unidad Previsional
	"UZS": {}, // Uzbekistan Sum
	"860": {}, // Uzbekistan Sum
	"VUV": {}, // Vatu
	"548": {}, // Vatu
	"VES": {}, // Bolívar Soberano
	"928": {}, // Bolívar Soberano
	"VND": {}, // Dong
	"704": {}, // Dong
	"YER": {}, // Yemeni Rial
	"886": {}, // Yemeni Rial
	"ZMW": {}, // Zambian Kwacha
	"967": {}, // Zambian Kwacha
	"ZWL": {}, // Zimbabwe Dollar
	"932": {}, // Zimbabwe Dollar
	"ZWG": {}, // Zimbabwe Gold
	"924": {}, // Zimbabwe Gold
	"XBA": {}, // Bond Markets Unit European Composite Unit (EURCO)
	"955": {}, // Bond Markets Unit European Composite Unit (EURCO)
	"XBB": {}, // Bond Markets Unit European Monetary Unit (E.M.U.-6)
	"956": {}, // Bond Markets Unit European Monetary Unit (E.M.U.-6)
	"XBC": {}, // Bond Markets Unit European Unit of Account 9 (E.U.A.-9)
	"957": {}, // Bond Markets Unit European Unit of Account 9 (E.U.A.-9)
	"XBD": {}, // Bond Markets Unit European Unit of Account 17 (E.U.A.-17)
	"958": {}, // Bond Markets Unit European Unit of Account 17 (E.U.A.-17)
	"XTS": {}, // Codes specifically reserved for testing purposes
	"963": {}, // Codes specifically reserved for testing purposes
	"XXX": {}, // The codes assigned for transactions where no currency is involved
	"999": {}, // The codes assigned for transactions where no currency is involved
	"XAU": {}, // Gold
	"959": {}, // Gold
	"XPD": {}, // Palladium
	"964": {}, // Palladium
	"XPT": {}, // Platinum
	"962": {}, // Platinum
	"XAG": {}, // Silver
	"961": {}, // Silver
}

func Lookup(code string) bool {
	_, exists := lookupTable[code]
	return exists
}
