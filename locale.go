package main

import (
	"github.com/petrjahoda/database"
	"gorm.io/gorm"
)

func createLocales(db *gorm.DB) {
	logInfo("MAIN", "Creating locales")
	//live := database.Locale{
	//	Name: "navbar-live",
	//	CsCZ: "živě",
	//	DeDE: "live",
	//	EnUS: "live",
	//	EsES: "vivo",
	//	FrFR: "direct",
	//	ItIT: "diretta",
	//	PlPL: "żywo",
	//	PtPT: "vivo",
	//	SkSK: "živo",
	//	RuRU: "прямое",
	//}
	//db.Create(&live)
	//charts := database.Locale{
	//	Name: "navbar-charts",
	//	CsCZ: "grafy",
	//	DeDE: "diagramme",
	//	EnUS: "charts",
	//	EsES: "gráficos",
	//	FrFR: "graphiques",
	//	ItIT: "grafici",
	//	PlPL: "wykresy",
	//	PtPT: "gráficos",
	//	SkSK: "grafy",
	//	RuRU: "графики",
	//}
	//db.Create(&charts)
	//statistics := database.Locale{
	//	Name: "navbar-statistics",
	//	CsCZ: "statistiky",
	//	DeDE: "statistiken",
	//	EnUS: "statistics",
	//	EsES: "estadísticas",
	//	FrFR: "statistiques",
	//	ItIT: "statistiche",
	//	PlPL: "statystyka",
	//	PtPT: "estatisticas",
	//	SkSK: "štatistiky",
	//	RuRU: "статистика",
	//}
	//db.Create(&statistics)
	//comparisons := database.Locale{
	//	Name: "navbar-comparisons",
	//	CsCZ: "porovnání",
	//	DeDE: "vergleiche",
	//	EnUS: "comparisons",
	//	EsES: "comparaciones",
	//	FrFR: "comparaisons",
	//	ItIT: "confronti",
	//	PlPL: "porównania",
	//	PtPT: "comparações",
	//	SkSK: "porovnania",
	//	RuRU: "сравнения",
	//}
	//db.Create(&comparisons)
	//trends := database.Locale{
	//	Name: "navbar-trends",
	//	CsCZ: "trendy",
	//	DeDE: "trends",
	//	EnUS: "trends",
	//	EsES: "tendencias",
	//	FrFR: "tendances",
	//	ItIT: "tendenze",
	//	PlPL: "trendy",
	//	PtPT: "tendências",
	//	SkSK: "trendy",
	//	RuRU: "тенденции",
	//}
	//db.Create(&trends)
	//data := database.Locale{
	//	Name: "navbar-data",
	//	CsCZ: "data",
	//	DeDE: "daten",
	//	EnUS: "data",
	//	EsES: "datos",
	//	FrFR: "données",
	//	ItIT: "dati",
	//	PlPL: "dane",
	//	PtPT: "dados",
	//	SkSK: "dáta",
	//	RuRU: "данные",
	//}
	//db.Create(&data)
	//settings := database.Locale{
	//	Name: "navbar-settings",
	//	CsCZ: "nastavení",
	//	DeDE: "einstellungen",
	//	EnUS: "settings",
	//	EsES: "ajustes",
	//	FrFR: "réglages",
	//	ItIT: "impostazioni",
	//	PlPL: "ustawienia",
	//	PtPT: "definições",
	//	SkSK: "nastavenie",
	//	RuRU: "настройки",
	//}
	//db.Create(&settings)
	//chartsStandard := database.Locale{
	//	Name: "navbar-charts-standard",
	//	CsCZ: "standardní",
	//	DeDE: "standard",
	//	EnUS: "standard",
	//	EsES: "estándar",
	//	FrFR: "standard",
	//	ItIT: "standard",
	//	PlPL: "standard",
	//	PtPT: "padrão",
	//	SkSK: "štandardné",
	//	RuRU: "стандарт",
	//}
	//db.Create(&chartsStandard)
	//chartsSpecial := database.Locale{
	//	Name: "navbar-charts-special",
	//	CsCZ: "speciální",
	//	DeDE: "speziell",
	//	EnUS: "special",
	//	EsES: "especial",
	//	FrFR: "spécial",
	//	ItIT: "speciale",
	//	PlPL: "specjalne",
	//	PtPT: "especial",
	//	SkSK: "špeciálné",
	//	RuRU: "специальный",
	//}
	//db.Create(&chartsSpecial)
	//liveCompany := database.Locale{
	//	Name: "navbar-live-company",
	//	CsCZ: "společnost",
	//	DeDE: "firma",
	//	EnUS: "company",
	//	EsES: "empresa",
	//	FrFR: "firme",
	//	ItIT: "azienda",
	//	PlPL: "firma",
	//	PtPT: "companhia",
	//	SkSK: "spoločnosť",
	//	RuRU: "Компания",
	//}
	//db.Create(&liveCompany)
	//liveGroup := database.Locale{
	//	Name: "navbar-live-group",
	//	CsCZ: "skupina",
	//	DeDE: "gruppe",
	//	EnUS: "group",
	//	EsES: "grupo",
	//	FrFR: "groupe",
	//	ItIT: "gruppo",
	//	PlPL: "grupa",
	//	PtPT: "grupo",
	//	SkSK: "skupina",
	//	RuRU: "группа",
	//}
	//db.Create(&liveGroup)
	//liveWorkplace := database.Locale{
	//	Name: "navbar-live-workplace",
	//	CsCZ: "pracoviště",
	//	DeDE: "arbeitsplatz",
	//	EnUS: "workplace",
	//	EsES: "lugar de trabajo",
	//	FrFR: "lieu de travail",
	//	ItIT: "posto di lavoro",
	//	PlPL: "miejsce pracy",
	//	PtPT: "local de trabalho",
	//	SkSK: "pracovisko",
	//	RuRU: "рабочее место",
	//}
	//db.Create(&liveWorkplace)
	//statisticCompany := database.Locale{
	//	Name: "navbar-statistics-company",
	//	CsCZ: "společnost",
	//	DeDE: "firma",
	//	EnUS: "company",
	//	EsES: "empresa",
	//	FrFR: "firme",
	//	ItIT: "azienda",
	//	PlPL: "firma",
	//	PtPT: "companhia",
	//	SkSK: "spoločnosť",
	//	RuRU: "Компания",
	//}
	//db.Create(&statisticCompany)
	//statisticsGroup := database.Locale{
	//	Name: "navbar-statistics-group",
	//	CsCZ: "skupina",
	//	DeDE: "gruppe",
	//	EnUS: "group",
	//	EsES: "grupo",
	//	FrFR: "groupe",
	//	ItIT: "gruppo",
	//	PlPL: "grupa",
	//	PtPT: "grupo",
	//	SkSK: "skupina",
	//	RuRU: "группа",
	//}
	//db.Create(&statisticsGroup)
	//statisticsWorkplace := database.Locale{
	//	Name: "navbar-statistics-workplace",
	//	CsCZ: "pracoviště",
	//	DeDE: "arbeitsplatz",
	//	EnUS: "workplace",
	//	EsES: "lugar de trabajo",
	//	FrFR: "lieu de travail",
	//	ItIT: "posto di lavoro",
	//	PlPL: "miejsce pracy",
	//	PtPT: "local de trabalho",
	//	SkSK: "pracovisko",
	//	RuRU: "рабочее место",
	//}
	//db.Create(&statisticsWorkplace)
	//
	//statisticsUser := database.Locale{
	//	Name: "navbar-statistics-user",
	//	CsCZ: "uživatel",
	//	DeDE: "nutzer",
	//	EnUS: "user",
	//	EsES: "usuario",
	//	FrFR: "utilisateur",
	//	ItIT: "utente",
	//	PlPL: "użytkownik",
	//	PtPT: "utilizador",
	//	SkSK: "užívateľ",
	//	RuRU: "пользователь",
	//}
	//db.Create(&statisticsUser)
	//
	//statisticsDowntime := database.Locale{
	//	Name: "navbar-statistics-downtime",
	//	CsCZ: "prostoj",
	//	DeDE: "ausfallzeit",
	//	EnUS: "downtime",
	//	EsES: "inactividad",
	//	FrFR: "inactivité",
	//	ItIT: "inattività",
	//	PlPL: "przestoj",
	//	PtPT: "inatividade",
	//	SkSK: "prestoj",
	//	RuRU: "простой",
	//}
	//db.Create(&statisticsDowntime)
	//dataBreakdowns := database.Locale{
	//	Name: "navbar-statistics-breakdown",
	//	CsCZ: "porucha",
	//	DeDE: "ausfall",
	//	EnUS: "breakdown",
	//	EsES: "trastorno",
	//	FrFR: "désordre",
	//	ItIT: "disturbo",
	//	PlPL: "nieład",
	//	PtPT: "transtorno",
	//	SkSK: "porucha",
	//	RuRU: "беспорядок",
	//}
	//db.Create(&dataBreakdowns)
	//statisticsAlarm := database.Locale{
	//	Name: "navbar-statistics-alarm",
	//	CsCZ: "alarm",
	//	DeDE: "alarm",
	//	EnUS: "alarm",
	//	EsES: "alarma",
	//	FrFR: "alarme",
	//	ItIT: "allarme",
	//	PlPL: "alarm",
	//	PtPT: "alarme",
	//	SkSK: "alarm",
	//	RuRU: "тревога",
	//}
	//db.Create(&statisticsAlarm)
	//
	//trendsCompany := database.Locale{
	//	Name: "navbar-trends-company",
	//	CsCZ: "společnost",
	//	DeDE: "firma",
	//	EnUS: "company",
	//	EsES: "empresa",
	//	FrFR: "firme",
	//	ItIT: "azienda",
	//	PlPL: "firma",
	//	PtPT: "companhia",
	//	SkSK: "spoločnosť",
	//	RuRU: "Компания",
	//}
	//db.Create(&trendsCompany)
	//trendsGroup := database.Locale{
	//	Name: "navbar-trends-group",
	//	CsCZ: "skupina",
	//	DeDE: "gruppe",
	//	EnUS: "group",
	//	EsES: "grupo",
	//	FrFR: "groupe",
	//	ItIT: "gruppo",
	//	PlPL: "grupa",
	//	PtPT: "grupo",
	//	SkSK: "skupina",
	//	RuRU: "группа",
	//}
	//db.Create(&trendsGroup)
	//trendsWorkplace := database.Locale{
	//	Name: "navbar-trends-workplace",
	//	CsCZ: "pracoviště",
	//	DeDE: "arbeitsplatz",
	//	EnUS: "workplace",
	//	EsES: "lugar de trabajo",
	//	FrFR: "lieu de travail",
	//	ItIT: "posto di lavoro",
	//	PlPL: "miejsce pracy",
	//	PtPT: "local de trabalho",
	//	SkSK: "pracovisko",
	//	RuRU: "рабочее место",
	//}
	//db.Create(&trendsWorkplace)
	//
	//trendsUser := database.Locale{
	//	Name: "navbar-trends-user",
	//	CsCZ: "uživatel",
	//	DeDE: "nutzer",
	//	EnUS: "user",
	//	EsES: "usuario",
	//	FrFR: "utilisateur",
	//	ItIT: "utente",
	//	PlPL: "użytkownik",
	//	PtPT: "utilizador",
	//	SkSK: "užívateľ",
	//	RuRU: "пользователь",
	//}
	//db.Create(&trendsUser)
	//
	//trendsDowntime := database.Locale{
	//	Name: "navbar-trends-downtime",
	//	CsCZ: "prostoj",
	//	DeDE: "ausfallzeit",
	//	EnUS: "downtime",
	//	EsES: "inactividad",
	//	FrFR: "inactivité",
	//	ItIT: "inattività",
	//	PlPL: "przestoj",
	//	PtPT: "inatividade",
	//	SkSK: "prestoj",
	//	RuRU: "простой",
	//}
	//db.Create(&trendsDowntime)
	//trendsBreakdown := database.Locale{
	//	Name: "navbar-trends-breakdown",
	//	CsCZ: "porucha",
	//	DeDE: "ausfall",
	//	EnUS: "breakdown",
	//	EsES: "trastorno",
	//	FrFR: "désordre",
	//	ItIT: "disturbo",
	//	PlPL: "nieład",
	//	PtPT: "transtorno",
	//	SkSK: "porucha",
	//	RuRU: "беспорядок",
	//}
	//db.Create(&trendsBreakdown)
	//trendsAlarm := database.Locale{
	//	Name: "navbar-trends-alarm",
	//	CsCZ: "alarm",
	//	DeDE: "alarm",
	//	EnUS: "alarm",
	//	EsES: "alarma",
	//	FrFR: "alarme",
	//	ItIT: "allarme",
	//	PlPL: "alarm",
	//	PtPT: "alarme",
	//	SkSK: "alarm",
	//	RuRU: "тревога",
	//}
	//db.Create(&trendsAlarm)
	//
	//trendsTop := database.Locale{
	//	Name: "navbar-trends-top10",
	//	CsCZ: "10 nejlepších",
	//	DeDE: "beste 10",
	//	EnUS: "top 10",
	//	EsES: "10 mejores",
	//	FrFR: "meilleur 10",
	//	ItIT: "migliori",
	//	PlPL: "10 najlepszych",
	//	PtPT: "melhores 10",
	//	SkSK: "10 najlepších",
	//	RuRU: "10 лучших",
	//}
	//db.Create(&trendsTop)
	//
	//trendsWorst := database.Locale{
	//	Name: "navbar-trends-worst10",
	//	CsCZ: "10 nejhorších",
	//	DeDE: "schlimmste 10",
	//	EnUS: "worst 10",
	//	EsES: "10 peores",
	//	FrFR: "pire 10",
	//	ItIT: "peggiore 10",
	//	PlPL: "10 najgorszych",
	//	PtPT: "piores 10",
	//	SkSK: "10 najhorších",
	//	RuRU: "10 худшие",
	//}
	//db.Create(&trendsWorst)
	//
	//dataCompany := database.Locale{
	//	Name: "navbar-data-company",
	//	CsCZ: "společnost",
	//	DeDE: "firma",
	//	EnUS: "company",
	//	EsES: "empresa",
	//	FrFR: "firme",
	//	ItIT: "azienda",
	//	PlPL: "firma",
	//	PtPT: "companhia",
	//	SkSK: "spoločnosť",
	//	RuRU: "Компания",
	//}
	//db.Create(&dataCompany)
	//dataGroup := database.Locale{
	//	Name: "navbar-data-group",
	//	CsCZ: "skupina",
	//	DeDE: "gruppe",
	//	EnUS: "group",
	//	EsES: "grupo",
	//	FrFR: "groupe",
	//	ItIT: "gruppo",
	//	PlPL: "grupa",
	//	PtPT: "grupo",
	//	SkSK: "skupina",
	//	RuRU: "группа",
	//}
	//db.Create(&dataGroup)
	//dataWorkplace := database.Locale{
	//	Name: "navbar-data-workplace",
	//	CsCZ: "pracoviště",
	//	DeDE: "arbeitsplatz",
	//	EnUS: "workplace",
	//	EsES: "lugar de trabajo",
	//	FrFR: "lieu de travail",
	//	ItIT: "posto di lavoro",
	//	PlPL: "miejsce pracy",
	//	PtPT: "local de trabalho",
	//	SkSK: "pracovisko",
	//	RuRU: "рабочее место",
	//}
	//db.Create(&dataWorkplace)
	//
	//dataUser := database.Locale{
	//	Name: "navbar-data-user",
	//	CsCZ: "uživatel",
	//	DeDE: "nutzer",
	//	EnUS: "user",
	//	EsES: "usuario",
	//	FrFR: "utilisateur",
	//	ItIT: "utente",
	//	PlPL: "użytkownik",
	//	PtPT: "utilizador",
	//	SkSK: "užívateľ",
	//	RuRU: "пользователь",
	//}
	//db.Create(&dataUser)
	//
	//dataDowntime := database.Locale{
	//	Name: "navbar-data-downtime",
	//	CsCZ: "prostoj",
	//	DeDE: "ausfallzeit",
	//	EnUS: "downtime",
	//	EsES: "inactividad",
	//	FrFR: "inactivité",
	//	ItIT: "inattività",
	//	PlPL: "przestoj",
	//	PtPT: "inatividade",
	//	SkSK: "prestoj",
	//	RuRU: "простой",
	//}
	//db.Create(&dataDowntime)
	//dataBreakdown := database.Locale{
	//	Name: "navbar-data-breakdown",
	//	CsCZ: "porucha",
	//	DeDE: "ausfall",
	//	EnUS: "breakdown",
	//	EsES: "trastorno",
	//	FrFR: "désordre",
	//	ItIT: "disturbo",
	//	PlPL: "nieład",
	//	PtPT: "transtorno",
	//	SkSK: "porucha",
	//	RuRU: "беспорядок",
	//}
	//db.Create(&dataBreakdown)
	//dataAlarm := database.Locale{
	//	Name: "navbar-data-alarm",
	//	CsCZ: "alarm",
	//	DeDE: "alarm",
	//	EnUS: "alarm",
	//	EsES: "alarma",
	//	FrFR: "alarme",
	//	ItIT: "allarme",
	//	PlPL: "alarm",
	//	PtPT: "alarme",
	//	SkSK: "alarm",
	//	RuRU: "тревога",
	//}
	//db.Create(&dataAlarm)
	//
	//dataOrder := database.Locale{
	//	Name: "navbar-data-order",
	//	CsCZ: "zakázka",
	//	DeDE: "bestellung",
	//	EnUS: "order",
	//	EsES: "orden",
	//	FrFR: "ordre",
	//	ItIT: "ordinazione",
	//	PlPL: "zamówienie",
	//	PtPT: "ordem",
	//	SkSK: "zakázka",
	//	RuRU: "заказ",
	//}
	//db.Create(&dataOrder)
	//
	//dataProduct := database.Locale{
	//	Name: "navbar-data-product",
	//	CsCZ: "produkt",
	//	DeDE: "produkt",
	//	EnUS: "product",
	//	EsES: "producto",
	//	FrFR: "produit",
	//	ItIT: "prodotto",
	//	PlPL: "produkt",
	//	PtPT: "produtos",
	//	SkSK: "produkt",
	//	RuRU: "товар",
	//}
	//db.Create(&dataProduct)
	//
	//settingsCompany := database.Locale{
	//	Name: "navbar-settings-company",
	//	CsCZ: "společnost",
	//	DeDE: "firma",
	//	EnUS: "company",
	//	EsES: "empresa",
	//	FrFR: "firme",
	//	ItIT: "azienda",
	//	PlPL: "firma",
	//	PtPT: "companhia",
	//	SkSK: "spoločnosť",
	//	RuRU: "Компания",
	//}
	//db.Create(&settingsCompany)
	//settingsGroup := database.Locale{
	//	Name: "navbar-settings-group",
	//	CsCZ: "skupina",
	//	DeDE: "gruppe",
	//	EnUS: "group",
	//	EsES: "grupo",
	//	FrFR: "groupe",
	//	ItIT: "gruppo",
	//	PlPL: "grupa",
	//	PtPT: "grupo",
	//	SkSK: "skupina",
	//	RuRU: "группа",
	//}
	//db.Create(&settingsGroup)
	//settingsWorkplace := database.Locale{
	//	Name: "navbar-settings-workplace",
	//	CsCZ: "pracoviště",
	//	DeDE: "arbeitsplatz",
	//	EnUS: "workplace",
	//	EsES: "lugar de trabajo",
	//	FrFR: "lieu de travail",
	//	ItIT: "posto di lavoro",
	//	PlPL: "miejsce pracy",
	//	PtPT: "local de trabalho",
	//	SkSK: "pracovisko",
	//	RuRU: "рабочее место",
	//}
	//db.Create(&settingsWorkplace)
	//
	//settingsUser := database.Locale{
	//	Name: "navbar-settings-user",
	//	CsCZ: "uživatel",
	//	DeDE: "nutzer",
	//	EnUS: "user",
	//	EsES: "usuario",
	//	FrFR: "utilisateur",
	//	ItIT: "utente",
	//	PlPL: "użytkownik",
	//	PtPT: "utilizador",
	//	SkSK: "užívateľ",
	//	RuRU: "пользователь",
	//}
	//db.Create(&settingsUser)
	//
	//settingsDowntime := database.Locale{
	//	Name: "navbar-settings-downtime",
	//	CsCZ: "prostoj",
	//	DeDE: "ausfallzeit",
	//	EnUS: "downtime",
	//	EsES: "inactividad",
	//	FrFR: "inactivité",
	//	ItIT: "inattività",
	//	PlPL: "przestoj",
	//	PtPT: "inatividade",
	//	SkSK: "prestoj",
	//	RuRU: "простой",
	//}
	//db.Create(&settingsDowntime)
	//settingsBreakdown := database.Locale{
	//	Name: "navbar-settings-breakdown",
	//	CsCZ: "porucha",
	//	DeDE: "ausfall",
	//	EnUS: "breakdown",
	//	EsES: "trastorno",
	//	FrFR: "désordre",
	//	ItIT: "disturbo",
	//	PlPL: "nieład",
	//	PtPT: "transtorno",
	//	SkSK: "porucha",
	//	RuRU: "беспорядок",
	//}
	//db.Create(&settingsBreakdown)
	//settingsAlarm := database.Locale{
	//	Name: "navbar-settings-alarm",
	//	CsCZ: "alarm",
	//	DeDE: "alarm",
	//	EnUS: "alarm",
	//	EsES: "alarma",
	//	FrFR: "alarme",
	//	ItIT: "allarme",
	//	PlPL: "alarm",
	//	PtPT: "alarme",
	//	SkSK: "alarm",
	//	RuRU: "тревога",
	//}
	//db.Create(&settingsAlarm)
	//
	//settingsDevice := database.Locale{
	//	Name: "navbar-settings-device",
	//	CsCZ: "zařízení",
	//	DeDE: "gerät",
	//	EnUS: "device",
	//	EsES: "aparato",
	//	FrFR: "l'appareil",
	//	ItIT: "dispositivo",
	//	PlPL: "urządzenie",
	//	PtPT: "dispositivo",
	//	SkSK: "zariadenie",
	//	RuRU: "устройство",
	//}
	//db.Create(&settingsDevice)
	//
	//comparisonsCompany := database.Locale{
	//	Name: "navbar-comparisons-company",
	//	CsCZ: "společnost",
	//	DeDE: "firma",
	//	EnUS: "company",
	//	EsES: "empresa",
	//	FrFR: "firme",
	//	ItIT: "azienda",
	//	PlPL: "firma",
	//	PtPT: "companhia",
	//	SkSK: "spoločnosť",
	//	RuRU: "Компания",
	//}
	//db.Create(&comparisonsCompany)
	//comparisonsGroup := database.Locale{
	//	Name: "navbar-comparisons-group",
	//	CsCZ: "skupina",
	//	DeDE: "gruppe",
	//	EnUS: "group",
	//	EsES: "grupo",
	//	FrFR: "groupe",
	//	ItIT: "gruppo",
	//	PlPL: "grupa",
	//	PtPT: "grupo",
	//	SkSK: "skupina",
	//	RuRU: "группа",
	//}
	//db.Create(&comparisonsGroup)
	//comparisonsWorkplace := database.Locale{
	//	Name: "navbar-comparisons-workplace",
	//	CsCZ: "pracoviště",
	//	DeDE: "arbeitsplatz",
	//	EnUS: "workplace",
	//	EsES: "lugar de trabajo",
	//	FrFR: "lieu de travail",
	//	ItIT: "posto di lavoro",
	//	PlPL: "miejsce pracy",
	//	PtPT: "local de trabalho",
	//	SkSK: "pracovisko",
	//	RuRU: "рабочее место",
	//}
	//db.Create(&comparisonsWorkplace)
	//
	//comparisonsUser := database.Locale{
	//	Name: "navbar-comparisons-user",
	//	CsCZ: "uživatel",
	//	DeDE: "nutzer",
	//	EnUS: "user",
	//	EsES: "usuario",
	//	FrFR: "utilisateur",
	//	ItIT: "utente",
	//	PlPL: "użytkownik",
	//	PtPT: "utilizador",
	//	SkSK: "užívateľ",
	//	RuRU: "пользователь",
	//}
	//db.Create(&comparisonsUser)
	//
	//comparisonsDowntime := database.Locale{
	//	Name: "navbar-comparisons-downtime",
	//	CsCZ: "prostoj",
	//	DeDE: "ausfallzeit",
	//	EnUS: "downtime",
	//	EsES: "inactividad",
	//	FrFR: "inactivité",
	//	ItIT: "inattività",
	//	PlPL: "przestoj",
	//	PtPT: "inatividade",
	//	SkSK: "prestoj",
	//	RuRU: "простой",
	//}
	//db.Create(&comparisonsDowntime)
	//comparisonsBreakdown := database.Locale{
	//	Name: "navbar-comparisons-breakdown",
	//	CsCZ: "porucha",
	//	DeDE: "ausfall",
	//	EnUS: "breakdown",
	//	EsES: "trastorno",
	//	FrFR: "désordre",
	//	ItIT: "disturbo",
	//	PlPL: "nieład",
	//	PtPT: "transtorno",
	//	SkSK: "porucha",
	//	RuRU: "беспорядок",
	//}
	//db.Create(&comparisonsBreakdown)
	//comparisonsAlarm := database.Locale{
	//	Name: "navbar-comparisons-alarm",
	//	CsCZ: "alarm",
	//	DeDE: "alarm",
	//	EnUS: "alarm",
	//	EsES: "alarma",
	//	FrFR: "alarme",
	//	ItIT: "allarme",
	//	PlPL: "alarm",
	//	PtPT: "alarme",
	//	SkSK: "alarm",
	//	RuRU: "тревога",
	//}
	//db.Create(&comparisonsAlarm)
	//
	//companyTodayLive := database.Locale{
	//	Name: "navbar-live-company-1-today",
	//	CsCZ: "dnes",
	//	DeDE: "heute",
	//	EnUS: "today",
	//	EsES: "hoy",
	//	FrFR: "de nos jours",
	//	ItIT: "oggi",
	//	PlPL: "dzisiaj",
	//	PtPT: "hoje",
	//	SkSK: "dnes",
	//	RuRU: "сегодня",
	//}
	//db.Create(&companyTodayLive)
	//companyYesterdayLive := database.Locale{
	//	Name: "navbar-live-company-1-yesterday",
	//	CsCZ: "včera",
	//	DeDE: "gestern",
	//	EnUS: "yesterday",
	//	EsES: "ayer",
	//	FrFR: "hier",
	//	ItIT: "ieri",
	//	PlPL: "wczoraj",
	//	PtPT: "ontem",
	//	SkSK: "včera",
	//	RuRU: "вчера",
	//}
	//db.Create(&companyYesterdayLive)
	//companyThisWeek := database.Locale{
	//	Name: "navbar-live-company-1-this-week",
	//	CsCZ: "tento týden",
	//	DeDE: "diese woche",
	//	EnUS: "this week",
	//	EsES: "esta semana",
	//	FrFR: "cette semaine",
	//	ItIT: "questa settimana",
	//	PlPL: "ten tydzień",
	//	PtPT: "esta semana",
	//	SkSK: "tento týždeň",
	//	RuRU: "эта неделя",
	//}
	//db.Create(&companyThisWeek)
	//companyLastWeek := database.Locale{
	//	Name: "navbar-live-company-1-last-week",
	//	CsCZ: "minulý týden",
	//	DeDE: "letzte woche",
	//	EnUS: "last week",
	//	EsES: "semana passada",
	//	FrFR: "semaine dernière",
	//	ItIT: "settimana scorsa",
	//	PlPL: "zeszły tydzień",
	//	PtPT: "semana passsa",
	//	SkSK: "minulý týždeň",
	//	RuRU: "на прошлой неделе",
	//}
	//db.Create(&companyLastWeek)
	//
	//companyThisMonth := database.Locale{
	//	Name: "navbar-live-company-1-this-month",
	//	CsCZ: "tento měsíc",
	//	DeDE: "diesen monat",
	//	EnUS: "this month",
	//	EsES: "este mes",
	//	FrFR: "ce mois-ci",
	//	ItIT: "questo mese",
	//	PlPL: "ten miesiąc",
	//	PtPT: "este mes",
	//	SkSK: "tento mesiac",
	//	RuRU: "этот месяц",
	//}
	//db.Create(&companyThisMonth)
	//
	//companyLastMonth := database.Locale{
	//	Name: "navbar-live-company-1-last-month",
	//	CsCZ: "minulý měsíc",
	//	DeDE: "letzten monat",
	//	EnUS: "last month",
	//	EsES: "el mes pasado",
	//	FrFR: "le mois dernier",
	//	ItIT: "le scorso mese",
	//	PlPL: "zeszły miesiąc",
	//	PtPT: "este mes",
	//	SkSK: "minulý mesiac",
	//	RuRU: "прошлый месяц",
	//}
	//db.Create(&companyLastMonth)
	//
	//groupTodayLive := database.Locale{
	//	Name: "navbar-live-group-1-today",
	//	CsCZ: "dnes",
	//	DeDE: "heute",
	//	EnUS: "today",
	//	EsES: "hoy",
	//	FrFR: "de nos jours",
	//	ItIT: "oggi",
	//	PlPL: "dzisiaj",
	//	PtPT: "hoje",
	//	SkSK: "dnes",
	//	RuRU: "сегодня",
	//}
	//db.Create(&groupTodayLive)
	//groupYesterdayLive := database.Locale{
	//	Name: "navbar-live-group-1-yesterday",
	//	CsCZ: "včera",
	//	DeDE: "gestern",
	//	EnUS: "yesterday",
	//	EsES: "ayer",
	//	FrFR: "hier",
	//	ItIT: "ieri",
	//	PlPL: "wczoraj",
	//	PtPT: "ontem",
	//	SkSK: "včera",
	//	RuRU: "вчера",
	//}
	//db.Create(&groupYesterdayLive)
	//groupThisWeek := database.Locale{
	//	Name: "navbar-live-group-1-this-week",
	//	CsCZ: "tento týden",
	//	DeDE: "diese woche",
	//	EnUS: "this week",
	//	EsES: "esta semana",
	//	FrFR: "cette semaine",
	//	ItIT: "questa settimana",
	//	PlPL: "ten tydzień",
	//	PtPT: "esta semana",
	//	SkSK: "tento týždeň",
	//	RuRU: "эта неделя",
	//}
	//db.Create(&groupThisWeek)
	//groupLastWeek := database.Locale{
	//	Name: "navbar-live-group-1-last-week",
	//	CsCZ: "minulý týden",
	//	DeDE: "letzte woche",
	//	EnUS: "last week",
	//	EsES: "semana passada",
	//	FrFR: "semaine dernière",
	//	ItIT: "settimana scorsa",
	//	PlPL: "zeszły tydzień",
	//	PtPT: "semana passsa",
	//	SkSK: "minulý týždeň",
	//	RuRU: "на прошлой неделе",
	//}
	//db.Create(&groupLastWeek)
	//
	//groupThisMonth := database.Locale{
	//	Name: "navbar-live-group-1-this-month",
	//	CsCZ: "tento měsíc",
	//	DeDE: "diesen monat",
	//	EnUS: "this month",
	//	EsES: "este mes",
	//	FrFR: "ce mois-ci",
	//	ItIT: "questo mese",
	//	PlPL: "ten miesiąc",
	//	PtPT: "este mes",
	//	SkSK: "tento mesiac",
	//	RuRU: "этот месяц",
	//}
	//db.Create(&groupThisMonth)
	//
	//groupLastMonth := database.Locale{
	//	Name: "navbar-live-group-1-last-month",
	//	CsCZ: "minulý měsíc",
	//	DeDE: "letzten monat",
	//	EnUS: "last month",
	//	EsES: "el mes pasado",
	//	FrFR: "le mois dernier",
	//	ItIT: "le scorso mese",
	//	PlPL: "zeszły miesiąc",
	//	PtPT: "este mes",
	//	SkSK: "minulý mesiac",
	//	RuRU: "прошлый месяц",
	//}
	//db.Create(&groupLastMonth)
	//
	//workplaceTodayLive := database.Locale{
	//	Name: "navbar-live-workplace-1-today",
	//	CsCZ: "dnes",
	//	DeDE: "heute",
	//	EnUS: "today",
	//	EsES: "hoy",
	//	FrFR: "de nos jours",
	//	ItIT: "oggi",
	//	PlPL: "dzisiaj",
	//	PtPT: "hoje",
	//	SkSK: "dnes",
	//	RuRU: "сегодня",
	//}
	//db.Create(&workplaceTodayLive)
	//workplaceYesterdayLive := database.Locale{
	//	Name: "navbar-live-workplace-1-yesterday",
	//	CsCZ: "včera",
	//	DeDE: "gestern",
	//	EnUS: "yesterday",
	//	EsES: "ayer",
	//	FrFR: "hier",
	//	ItIT: "ieri",
	//	PlPL: "wczoraj",
	//	PtPT: "ontem",
	//	SkSK: "včera",
	//	RuRU: "вчера",
	//}
	//db.Create(&workplaceYesterdayLive)
	//workplaceThisWeek := database.Locale{
	//	Name: "navbar-live-workplace-1-this-week",
	//	CsCZ: "tento týden",
	//	DeDE: "diese woche",
	//	EnUS: "this week",
	//	EsES: "esta semana",
	//	FrFR: "cette semaine",
	//	ItIT: "questa settimana",
	//	PlPL: "ten tydzień",
	//	PtPT: "esta semana",
	//	SkSK: "tento týždeň",
	//	RuRU: "эта неделя",
	//}
	//db.Create(&workplaceThisWeek)
	//workplaceLastWeek := database.Locale{
	//	Name: "navbar-live-workplace-1-last-week",
	//	CsCZ: "minulý týden",
	//	DeDE: "letzte woche",
	//	EnUS: "last week",
	//	EsES: "semana passada",
	//	FrFR: "semaine dernière",
	//	ItIT: "settimana scorsa",
	//	PlPL: "zeszły tydzień",
	//	PtPT: "semana passsa",
	//	SkSK: "minulý týždeň",
	//	RuRU: "на прошлой неделе",
	//}
	//db.Create(&workplaceLastWeek)
	//
	//workplaceThisMonth := database.Locale{
	//	Name: "navbar-live-workplace-1-this-month",
	//	CsCZ: "tento měsíc",
	//	DeDE: "diesen monat",
	//	EnUS: "this month",
	//	EsES: "este mes",
	//	FrFR: "ce mois-ci",
	//	ItIT: "questo mese",
	//	PlPL: "ten miesiąc",
	//	PtPT: "este mes",
	//	SkSK: "tento mesiac",
	//	RuRU: "этот месяц",
	//}
	//db.Create(&workplaceThisMonth)
	//
	//workplaceLastMonth := database.Locale{
	//	Name: "navbar-live-workplace-1-last-month",
	//	CsCZ: "minulý měsíc",
	//	DeDE: "letzten monat",
	//	EnUS: "last month",
	//	EsES: "el mes pasado",
	//	FrFR: "le mois dernier",
	//	ItIT: "le scorso mese",
	//	PlPL: "zeszły miesiąc",
	//	PtPT: "este mes",
	//	SkSK: "minulý mesiac",
	//	RuRU: "прошлый месяц",
	//}
	//db.Create(&workplaceLastMonth)
	//
	//companyLiveHeader := database.Locale{
	//	Name: "navbar-live-company-1-header",
	//	CsCZ: "produktivita",
	//	DeDE: "produktivität",
	//	EnUS: "productivity",
	//	EsES: "productividad",
	//	FrFR: "productivité",
	//	ItIT: "produttività",
	//	PlPL: "wydajność",
	//	PtPT: "produtividade",
	//	SkSK: "produktivita",
	//	RuRU: "продуктивность",
	//}
	//db.Create(&companyLiveHeader)
	//
	//groupLiveHeader := database.Locale{
	//	Name: "navbar-live-group-1-header",
	//	CsCZ: "produktivita",
	//	DeDE: "produktivität",
	//	EnUS: "productivity",
	//	EsES: "productividad",
	//	FrFR: "productivité",
	//	ItIT: "produttività",
	//	PlPL: "wydajność",
	//	PtPT: "produtividade",
	//	SkSK: "produktivita",
	//	RuRU: "продуктивность",
	//}
	//db.Create(&groupLiveHeader)
	//
	//workplaceLiveHeader := database.Locale{
	//	Name: "navbar-live-workplace-1-header",
	//	CsCZ: "produktivita",
	//	DeDE: "produktivität",
	//	EnUS: "productivity",
	//	EsES: "productividad",
	//	FrFR: "productivité",
	//	ItIT: "produttività",
	//	PlPL: "wydajność",
	//	PtPT: "produtividade",
	//	SkSK: "produktivita",
	//	RuRU: "продуктивность",
	//}
	//db.Create(&workplaceLiveHeader)
	//
	//companyProductionLive := database.Locale{
	//	Name: "navbar-live-company-3-production",
	//	CsCZ: "v produkci",
	//	DeDE: "in produktion",
	//	EnUS: "in production",
	//	EsES: "en producción",
	//	FrFR: "en production",
	//	ItIT: "in produzione",
	//	PlPL: "w produkcji",
	//	PtPT: "em produção",
	//	SkSK: "v produkcii",
	//	RuRU: "в производстве",
	//}
	//db.Create(&companyProductionLive)
	//companyDowntimeLive := database.Locale{
	//	Name: "navbar-live-company-3-downtime",
	//	CsCZ: "v prostoji",
	//	DeDE: "in Ausfallzeiten",
	//	EnUS: "in downtime",
	//	EsES: "en tiempo de inactividad",
	//	FrFR: "en temps d'arrêt",
	//	ItIT: "nei tempi di inattività",
	//	PlPL: "w czasie przestoju",
	//	PtPT: "em tempo de inatividade",
	//	SkSK: "v prostoji",
	//	RuRU: "во время простоя",
	//}
	//db.Create(&companyDowntimeLive)
	//
	//companyPoweroffLive := database.Locale{
	//	Name: "navbar-live-company-3-poweroff",
	//	CsCZ: "vypnuto",
	//	DeDE: "ausgeschaltet",
	//	EnUS: "powered off",
	//	EsES: "apagado",
	//	FrFR: "éteint",
	//	ItIT: "spento ",
	//	PlPL: "wyłączony",
	//	PtPT: "sesligado",
	//	SkSK: "vypnuto",
	//	RuRU: "выключен",
	//}
	//db.Create(&companyPoweroffLive)
	//
	//groupProductionLive := database.Locale{
	//	Name: "navbar-live-group-3-production",
	//	CsCZ: "v produkci",
	//	DeDE: "in produktion",
	//	EnUS: "in production",
	//	EsES: "en producción",
	//	FrFR: "en production",
	//	ItIT: "in produzione",
	//	PlPL: "w produkcji",
	//	PtPT: "em produção",
	//	SkSK: "v produkcii",
	//	RuRU: "в производстве",
	//}
	//db.Create(&groupProductionLive)
	//groupDowntimeLive := database.Locale{
	//	Name: "navbar-live-group-3-downtime",
	//	CsCZ: "v prostoji",
	//	DeDE: "in Ausfallzeiten",
	//	EnUS: "in downtime",
	//	EsES: "en tiempo de inactividad",
	//	FrFR: "en temps d'arrêt",
	//	ItIT: "nei tempi di inattività",
	//	PlPL: "w czasie przestoju",
	//	PtPT: "em tempo de inatividade",
	//	SkSK: "v prostoji",
	//	RuRU: "во время простоя",
	//}
	//db.Create(&groupDowntimeLive)
	//
	//groupPoweroffLive := database.Locale{
	//	Name: "navbar-live-group-3-poweroff",
	//	CsCZ: "vypnuto",
	//	DeDE: "ausgeschaltet",
	//	EnUS: "powered off",
	//	EsES: "apagado",
	//	FrFR: "éteint",
	//	ItIT: "spento ",
	//	PlPL: "wyłączony",
	//	PtPT: "sesligado",
	//	SkSK: "vypnuto",
	//	RuRU: "выключен",
	//}
	//db.Create(&groupPoweroffLive)
	//
	//production := database.Locale{
	//	Name: "production",
	//	CsCZ: "výroba",
	//	DeDE: "produktion",
	//	EnUS: "production",
	//	EsES: "producción",
	//	FrFR: "production",
	//	ItIT: "produzione",
	//	PlPL: "produkcja",
	//	PtPT: "Produção",
	//	SkSK: "výroba",
	//	RuRU: "производство",
	//}
	//db.Create(&production)
	//
	//downtime := database.Locale{
	//	Name: "downtime",
	//	CsCZ: "prostoj",
	//	DeDE: "ausfallzeit",
	//	EnUS: "downtime",
	//	EsES: "inactividad",
	//	FrFR: "inactivité",
	//	ItIT: "inattività",
	//	PlPL: "przestoj",
	//	PtPT: "inatividade",
	//	SkSK: "prestoj",
	//	RuRU: "простой",
	//}
	//db.Create(&downtime)
	//
	//poweroff := database.Locale{
	//	Name: "poweroff",
	//	CsCZ: "vypnuto",
	//	DeDE: "ausgeschaltet",
	//	EnUS: "poweroff",
	//	EsES: "apagado",
	//	FrFR: "éteint",
	//	ItIT: "spento",
	//	PlPL: "wyłączony",
	//	PtPT: "desligado",
	//	SkSK: "vypnuté",
	//	RuRU: "выключен",
	//}
	//db.Create(&poweroff)
	//
	//liveWorkplaceUser := database.Locale{
	//	Name: "navbar-live-workplace-3-user",
	//	CsCZ: "operátor",
	//	DeDE: "operator",
	//	EnUS: "operator",
	//	EsES: "operador",
	//	FrFR: "opérateur",
	//	ItIT: "operatore",
	//	PlPL: "operator",
	//	PtPT: "operador",
	//	SkSK: "operátor",
	//	RuRU: "оператор",
	//}
	//db.Create(&liveWorkplaceUser)
	//
	//liveWorkplaceOrder := database.Locale{
	//	Name: "navbar-live-workplace-3-order",
	//	CsCZ: "zakázka",
	//	DeDE: "bestellung",
	//	EnUS: "order",
	//	EsES: "orden",
	//	FrFR: "ordre",
	//	ItIT: "ordinazione",
	//	PlPL: "zamówienie",
	//	PtPT: "ordem",
	//	SkSK: "zakázka",
	//	RuRU: "заказ",
	//}
	//db.Create(&liveWorkplaceOrder)
	//
	//liveWorkplaceDowntime := database.Locale{
	//	Name: "navbar-live-workplace-3-downtime",
	//	CsCZ: "prostoj",
	//	DeDE: "ausfallzeit",
	//	EnUS: "downtime",
	//	EsES: "inactividad",
	//	FrFR: "inactivité",
	//	ItIT: "inattività",
	//	PlPL: "przestoj",
	//	PtPT: "inatividade",
	//	SkSK: "prestoj",
	//	RuRU: "простой",
	//}
	//db.Create(&liveWorkplaceDowntime)
	//liveWorkplaceBreakdown := database.Locale{
	//	Name: "navbar-live-workplace-3-breakdown",
	//	CsCZ: "porucha",
	//	DeDE: "ausfall",
	//	EnUS: "breakdown",
	//	EsES: "trastorno",
	//	FrFR: "désordre",
	//	ItIT: "disturbo",
	//	PlPL: "nieład",
	//	PtPT: "transtorno",
	//	SkSK: "porucha",
	//	RuRU: "беспорядок",
	//}
	//db.Create(&liveWorkplaceBreakdown)
	//liveWorkplaceAlarm := database.Locale{
	//	Name: "navbar-live-workplace-3-alarm",
	//	CsCZ: "alarm",
	//	DeDE: "alarm",
	//	EnUS: "alarm",
	//	EsES: "alarma",
	//	FrFR: "alarme",
	//	ItIT: "allarme",
	//	PlPL: "alarm",
	//	PtPT: "alarme",
	//	SkSK: "alarm",
	//	RuRU: "тревога",
	//}
	//db.Create(&liveWorkplaceAlarm)

	menuOverview := database.Locale{
		Name: "menu-overview",
		CsCZ: "přehled",
		DeDE: "überblick",
		EnUS: "overview",
		EsES: "visión general",
		FrFR: "aperçu",
		ItIT: "panoramica",
		PlPL: "przegląd",
		PtPT: "visão geral",
		SkSK: "prehľad",
		RuRU: "обзор",
	}
	db.Create(&menuOverview)

	menuWorkplaces := database.Locale{
		Name: "menu-workplaces",
		CsCZ: "skupina",
		DeDE: "gruppe",
		EnUS: "group",
		EsES: "grupo",
		FrFR: "groupe",
		ItIT: "gruppo",
		PlPL: "grupa",
		PtPT: "grupo",
		SkSK: "skupina",
		RuRU: "группа",
	}
	db.Create(&menuWorkplaces)

	menuCharts := database.Locale{
		Name: "menu-charts",
		CsCZ: "grafy",
		DeDE: "diagramme",
		EnUS: "charts",
		EsES: "gráficos",
		FrFR: "graphiques",
		ItIT: "grafici",
		PlPL: "wykresy",
		PtPT: "gráficos",
		SkSK: "grafy",
		RuRU: "графики",
	}
	db.Create(&menuCharts)

	menuStatistics := database.Locale{
		Name: "menu-statistics",
		CsCZ: "statistiky",
		DeDE: "statistiken",
		EnUS: "statistics",
		EsES: "estadísticas",
		FrFR: "statistiques",
		ItIT: "statistiche",
		PlPL: "statystyka",
		PtPT: "estatisticas",
		SkSK: "štatistiky",
		RuRU: "статистика",
	}
	db.Create(&menuStatistics)

	menuData := database.Locale{
		Name: "menu-data",
		CsCZ: "data",
		DeDE: "daten",
		EnUS: "data",
		EsES: "datos",
		FrFR: "données",
		ItIT: "dati",
		PlPL: "dane",
		PtPT: "dados",
		SkSK: "dáta",
		RuRU: "данные",
	}
	db.Create(&menuData)
	menuSettings := database.Locale{
		Name: "menu-settings",
		CsCZ: "nastavení",
		DeDE: "einstellungen",
		EnUS: "settings",
		EsES: "ajustes",
		FrFR: "réglages",
		ItIT: "impostazioni",
		PlPL: "ustawienia",
		PtPT: "definições",
		SkSK: "nastavenie",
		RuRU: "настройки",
	}
	db.Create(&menuSettings)

	dataAlarms := database.Locale{
		Model: gorm.Model{},
		Name:  "alarms",
		CsCZ:  "alarmy",
		DeDE:  "alarme",
		EnUS:  "alarms",
		EsES:  "alarmas",
		FrFR:  "alarmes",
		ItIT:  "allarmi",
		PlPL:  "alarmy",
		PtPT:  "alarmes",
		SkSK:  "alarmy",
		RuRU:  "будильники",
	}
	db.Create(&dataAlarms)

	dataBreakdowns := database.Locale{
		Name: "breakdowns",
		CsCZ: "poruchy",
		DeDE: "störungen",
		EnUS: "breakdowns",
		EsES: "averías",
		FrFR: "pannes",
		ItIT: "disturbi",
		PlPL: "zaburzenia",
		PtPT: "desordens",
		SkSK: "poruchy",
		RuRU: "расстройства",
	}
	db.Create(&dataBreakdowns)

	dataDowntimes := database.Locale{
		Name: "downtimes",
		CsCZ: "prostoje",
		DeDE: "ausfallzeiten",
		EnUS: "downtimes",
		EsES: "inactividades",
		FrFR: "inactivités",
		ItIT: "inattività",
		PlPL: "przestoje",
		PtPT: "inatividades",
		SkSK: "prestoje",
		RuRU: "просто́я",
	}
	db.Create(&dataDowntimes)

	dataFaults := database.Locale{
		Name: "faults",
		CsCZ: "chyby",
		DeDE: "fehler",
		EnUS: "faults",
		EsES: "fallas",
		FrFR: "défauts",
		ItIT: "difetti",
		PlPL: "wady",
		PtPT: "falhas",
		SkSK: "chyby",
		RuRU: "недостатки",
	}
	db.Create(&dataFaults)

	dataOrders := database.Locale{
		Name: "orders",
		CsCZ: "zakázky",
		DeDE: "bestellungen",
		EnUS: "orders",
		EsES: "pedidos",
		FrFR: "ordres",
		ItIT: "ordini",
		PlPL: "zamówienia",
		PtPT: "ordens",
		SkSK: "zakázky",
		RuRU: "заказы",
	}
	db.Create(&dataOrders)

	dataPackages := database.Locale{
		Name: "packages",
		CsCZ: "balení",
		DeDE: "pakete",
		EnUS: "packages",
		EsES: "paquetes",
		FrFR: "paquets",
		ItIT: "pacchetti",
		PlPL: "pakiety",
		PtPT: "pacotes",
		SkSK: "balenie",
		RuRU: "пакеты",
	}
	db.Create(&dataPackages)

	dataParts := database.Locale{
		Name: "parts",
		CsCZ: "díly",
		DeDE: "teile",
		EnUS: "parts",
		EsES: "partes",
		FrFR: "pièces",
		ItIT: "parti",
		PlPL: "części",
		PtPT: "partes",
		SkSK: "diely",
		RuRU: "части",
	}
	db.Create(&dataParts)

	dataStates := database.Locale{
		Name: "states",
		CsCZ: "stavy",
		DeDE: "zustände",
		EnUS: "states",
		EsES: "estados",
		FrFR: "états",
		ItIT: "stati",
		PlPL: "stany",
		PtPT: "estados",
		SkSK: "stavy",
		RuRU: "состояния",
	}
	db.Create(&dataStates)

	dataUsers := database.Locale{
		Name: "users",
		CsCZ: "uživatelé",
		DeDE: "benutzer",
		EnUS: "users",
		EsES: "usuarias",
		FrFR: "utilisateurs",
		ItIT: "utenti",
		PlPL: "użytkowników",
		PtPT: "comercial",
		SkSK: "užívatelia",
		RuRU: "пользователи",
	}
	db.Create(&dataUsers)

	systemData := database.Locale{
		Name: "system-statistics",
		CsCZ: "systémová data",
		DeDE: "systemdaten",
		EnUS: "system data",
		EsES: "datos de sistema",
		FrFR: "données système",
		ItIT: "dati di sistema",
		PlPL: "dane systemowe",
		PtPT: "dados do sistema",
		SkSK: "systémové údaje",
		RuRU: "системные данные",
	}
	db.Create(&systemData)

}
