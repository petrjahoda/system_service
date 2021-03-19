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
		CsCZ:  "Alarmy",
		DeDE:  "Alarme",
		EnUS:  "Alarms",
		EsES:  "Alarmas",
		FrFR:  "Alarmes",
		ItIT:  "Allarmi",
		PlPL:  "Alarmy",
		PtPT:  "Alarmes",
		SkSK:  "Alarmy",
		RuRU:  "Будильники",
	}
	db.Create(&dataAlarms)

	dataBreakdowns := database.Locale{
		Name: "breakdowns",
		CsCZ: "Poruchy",
		DeDE: "Störungen",
		EnUS: "Breakdowns",
		EsES: "Averías",
		FrFR: "Pannes",
		ItIT: "Disturbi",
		PlPL: "Zaburzenia",
		PtPT: "Desordens",
		SkSK: "Poruchy",
		RuRU: "Pасстройства",
	}
	db.Create(&dataBreakdowns)

	dataDowntimes := database.Locale{
		Name: "downtimes",
		CsCZ: "Prostoje",
		DeDE: "Ausfallzeiten",
		EnUS: "Downtimes",
		EsES: "Inactividades",
		FrFR: "Inactivités",
		ItIT: "Inattività",
		PlPL: "Przestoje",
		PtPT: "Inatividades",
		SkSK: "Prestoje",
		RuRU: "Просто́я",
	}
	db.Create(&dataDowntimes)

	dataFaults := database.Locale{
		Name: "faults",
		CsCZ: "Chyby",
		DeDE: "Fehler",
		EnUS: "Faults",
		EsES: "Fallas",
		FrFR: "Défauts",
		ItIT: "Difetti",
		PlPL: "Wady",
		PtPT: "Falhas",
		SkSK: "Chyby",
		RuRU: "Недостатки",
	}
	db.Create(&dataFaults)

	dataOrders := database.Locale{
		Name: "orders",
		CsCZ: "Zakázky",
		DeDE: "Bestellungen",
		EnUS: "Orders",
		EsES: "Pedidos",
		FrFR: "Ordres",
		ItIT: "Ordini",
		PlPL: "Zamówienia",
		PtPT: "Ordens",
		SkSK: "Zakázky",
		RuRU: "Заказы",
	}
	db.Create(&dataOrders)

	dataPackages := database.Locale{
		Name: "packages",
		CsCZ: "Balení",
		DeDE: "Pakete",
		EnUS: "Packages",
		EsES: "Paquetes",
		FrFR: "Paquets",
		ItIT: "Pacchetti",
		PlPL: "Pakiety",
		PtPT: "Pacotes",
		SkSK: "Balenie",
		RuRU: "Пакеты",
	}
	db.Create(&dataPackages)

	dataParts := database.Locale{
		Name: "parts",
		CsCZ: "Díly",
		DeDE: "Teile",
		EnUS: "Parts",
		EsES: "Partes",
		FrFR: "Pièces",
		ItIT: "Parti",
		PlPL: "Części",
		PtPT: "Partes",
		SkSK: "Diely",
		RuRU: "Запчасти",
	}
	db.Create(&dataParts)

	dataStates := database.Locale{
		Name: "states",
		CsCZ: "Stavy",
		DeDE: "Zustände",
		EnUS: "States",
		EsES: "Estados",
		FrFR: "États",
		ItIT: "Stati",
		PlPL: "Stany",
		PtPT: "Estados",
		SkSK: "Stavy",
		RuRU: "Cостояния",
	}
	db.Create(&dataStates)

	dataUsers := database.Locale{
		Name: "users",
		CsCZ: "Uživatelé",
		DeDE: "Benutzer",
		EnUS: "Users",
		EsES: "Usuarias",
		FrFR: "Utilisateurs",
		ItIT: "Utenti",
		PlPL: "Użytkowników",
		PtPT: "Comercial",
		SkSK: "Užívatelia",
		RuRU: "Пользователи",
	}
	db.Create(&dataUsers)

	systemData := database.Locale{
		Name: "system-statistics",
		CsCZ: "Systémová data",
		DeDE: "Systemdaten",
		EnUS: "System data",
		EsES: "Datos de sistema",
		FrFR: "Données système",
		ItIT: "Dati di sistema",
		PlPL: "Dane systemowe",
		PtPT: "Dados do sistema",
		SkSK: "Systémové údaje",
		RuRU: "Cистемные данные",
	}
	db.Create(&systemData)

	workplaceName := database.Locale{
		Name: "workplace-name",
		CsCZ: "Pracoviště",
		DeDE: "Arbeitsplatz",
		EnUS: "Workplace",
		EsES: "Lugar de trabajo",
		FrFR: "Lieu de travail",
		ItIT: "Posto di lavoro",
		PlPL: "Miejsce pracy",
		PtPT: "Local de trabalho",
		SkSK: "Pracovisko",
		RuRU: "Pабочее место",
	}
	db.Create(&workplaceName)

	orderStart := database.Locale{
		Name: "order-start",
		CsCZ: "Začátek",
		DeDE: "Anfang",
		EnUS: "Beginning",
		EsES: "Comenzando",
		FrFR: "Début",
		ItIT: "Inizio",
		PlPL: "Początek",
		PtPT: "Começo",
		SkSK: "Začiatok",
		RuRU: "Начало",
	}
	db.Create(&orderStart)

	orderEnd := database.Locale{
		Name: "order-end",
		CsCZ: "Konec",
		DeDE: "Ende",
		EnUS: "End",
		EsES: "Fin",
		FrFR: "Finir",
		ItIT: "Fine",
		PlPL: "Koniec",
		PtPT: "Fim",
		SkSK: "Koniec",
		RuRU: "Конец",
	}
	db.Create(&orderEnd)

	workplaceModeName := database.Locale{
		Name: "workplacemode-name",
		CsCZ: "Pracovní režim",
		DeDE: "Arbeitsplatzmodus",
		EnUS: "Workplace mode",
		EsES: "Modo de trabajo",
		FrFR: "Mode de travail",
		ItIT: "Modalità di lavoro",
		PlPL: "Tryb pracy",
		PtPT: "Modo de trabalho",
		SkSK: "Režim pracoviska",
		RuRU: "Pабочий режим",
	}
	db.Create(&workplaceModeName)

	workshiftName := database.Locale{
		Name: "workshift-name",
		CsCZ: "Pracovní směna",
		DeDE: "Arbeitsschicht",
		EnUS: "Workshift",
		EsES: "Turno de trabajo",
		FrFR: "Un poste de travail",
		ItIT: "Turno di lavoro",
		PlPL: "Zmiany roboczej",
		PtPT: "Turno de trabalho",
		SkSK: "Pracovná zmena",
		RuRU: "Pабочая смена",
	}
	db.Create(&workshiftName)

	userName := database.Locale{
		Name: "user-name",
		CsCZ: "Uživatel",
		DeDE: "Nutzer",
		EnUS: "User",
		EsES: "Usuario",
		FrFR: "Utilisateur",
		ItIT: "Utente",
		PlPL: "Użytkownik",
		PtPT: "Utilizador",
		SkSK: "Užívateľ",
		RuRU: "Пользователь",
	}
	db.Create(&userName)

	orderName := database.Locale{
		Name: "order-name",
		CsCZ: "Zakázka",
		DeDE: "Bestellung",
		EnUS: "Order",
		EsES: "Orden",
		FrFR: "Ordre",
		ItIT: "Ordinazione",
		PlPL: "Zamówienie",
		PtPT: "Ordem",
		SkSK: "Zakázka",
		RuRU: "Заказ",
	}
	db.Create(&orderName)

	operationName := database.Locale{
		Name: "operation-name",
		CsCZ: "Operace",
		DeDE: "Operation",
		EnUS: "Operation",
		EsES: "Operación",
		FrFR: "Opération",
		ItIT: "Operazione",
		PlPL: "Operacja",
		PtPT: "Operação",
		SkSK: "Operácia",
		RuRU: "Oперация",
	}
	db.Create(&operationName)

	cycleName := database.Locale{
		Name: "cycle-name",
		CsCZ: "Cyklus",
		DeDE: "Zyklus",
		EnUS: "Cycle",
		EsES: "Ciclo",
		FrFR: "Cycle",
		ItIT: "Ciclo",
		PlPL: "Cykl",
		PtPT: "Ciclo",
		SkSK: "Cyklus",
		RuRU: "Цикл",
	}
	db.Create(&cycleName)

	cavityName := database.Locale{
		Name: "cavity-name",
		CsCZ: "Násobič",
		DeDE: "Multiplikator",
		EnUS: "Multiplier",
		EsES: "Multiplicador",
		FrFR: "Multiplicateur",
		ItIT: "Moltiplicatore",
		PlPL: "Mnożnik",
		PtPT: "Multiplicador",
		SkSK: "Másobiteľ",
		RuRU: "Мультипликатор",
	}
	db.Create(&cavityName)

	goodName := database.Locale{
		Name: "good-pieces-name",
		CsCZ: "Dobré kusy",
		DeDE: "Gute stücke",
		EnUS: "Good pieces",
		EsES: "Buenas piezas",
		FrFR: "Bons morceaux",
		ItIT: "Buoni pezzi",
		PlPL: "Dobre kawałki",
		PtPT: "Doas peças",
		SkSK: "Dobré kusy",
		RuRU: "Добрэ кусы",
	}
	db.Create(&goodName)

	badName := database.Locale{
		Name: "bad-pieces-name",
		CsCZ: "Špatné kusy",
		DeDE: "Schlechte stücke",
		EnUS: "Bad pieces",
		EsES: "Malas piezas",
		FrFR: "Mauvais morceaux",
		ItIT: "Pezzi cattivi",
		PlPL: "Złe kawałki",
		PtPT: "Peças ruins",
		SkSK: "Zlé kusy",
		RuRU: "Плохие штуки",
	}
	db.Create(&badName)

	noteName := database.Locale{
		Name: "note-name",
		CsCZ: "Poznámka",
		DeDE: "Note",
		EnUS: "Note",
		EsES: "Nota",
		FrFR: "La note",
		ItIT: "La nota",
		PlPL: "Notatka",
		PtPT: "Anotação",
		SkSK: "Poznámka",
		RuRU: "Примечание",
	}
	db.Create(&noteName)

	showEntries := database.Locale{
		Name: "data-table-rows-count-title",
		CsCZ: "Zobrazit záznamy: ",
		DeDE: "Einträge anzeigen: ",
		EnUS: "Show entries: ",
		EsES: "Mostrar entradas: ",
		FrFR: "Afficher les entrées: ",
		ItIT: "Mostra voci: ",
		PlPL: "Pokaż wpisy: ",
		PtPT: "Mostrar entradas: ",
		SkSK: "Zobraziť záznamy: ",
		RuRU: "показать записи: ",
	}
	db.Create(&showEntries)

	search := database.Locale{
		Name: "data-table-search-title",
		CsCZ: "Hledat: ",
		DeDE: "Suche: ",
		EnUS: "Search: ",
		EsES: "Buscar: ",
		FrFR: "Chercher: ",
		ItIT: "Ricerca: ",
		PlPL: "Szukaj: ",
		PtPT: "Procurar: ",
		SkSK: "Vyhľadávanie: ",
		RuRU: "Поиск: ",
	}
	db.Create(&search)

	tableInfo := database.Locale{
		Name: "data-table-info-title",
		CsCZ: "Zobrazeno $1 až $2 z $3 záznamů",
		DeDE: "Anzeigen von $1 bis $2 von $3 Einträgen",
		EnUS: "Showing $1 to $2 of $3 entries",
		EsES: "Mostrando $1 a $2 de $3 entradas",
		FrFR: "Affichage des entrées de $1 à $2 sur $3",
		ItIT: "Visualizzazione di $1 a $2 di $3 voci",
		PlPL: "Wyświetlam pozycje od $1 do $2 z $3",
		PtPT: "Mostrando $1 a $2 de $3 entradas",
		SkSK: "Zobrazuje sa $1 až $2 z $3 záznamov",
		RuRU: "Показаны записи от $1 до $2 из $3",
	}
	db.Create(&tableInfo)

	alarmName := database.Locale{
		Name: "Alarm-name",
		CsCZ: "Alarm",
		DeDE: "Alarm",
		EnUS: "Alarm",
		EsES: "Alarma",
		FrFR: "Alarme",
		ItIT: "Allarme",
		PlPL: "Alarm",
		PtPT: "Alarme",
		SkSK: "Alarm",
		RuRU: "Тревога",
	}
	db.Create(&alarmName)

	alarmStart := database.Locale{
		Name: "alarm-start",
		CsCZ: "Začátek",
		DeDE: "Anfang",
		EnUS: "Beginning",
		EsES: "Comenzando",
		FrFR: "Début",
		ItIT: "Inizio",
		PlPL: "Początek",
		PtPT: "Começo",
		SkSK: "Začiatok",
		RuRU: "Начало",
	}
	db.Create(&alarmStart)

	alarmEnd := database.Locale{
		Name: "alarm-end",
		CsCZ: "Konec",
		DeDE: "Ende",
		EnUS: "End",
		EsES: "Fin",
		FrFR: "Finir",
		ItIT: "Fine",
		PlPL: "Koniec",
		PtPT: "Fim",
		SkSK: "Koniec",
		RuRU: "Конец",
	}
	db.Create(&alarmEnd)

	alarmProcessed := database.Locale{
		Name: "alarm-processed",
		CsCZ: "Zpracováno",
		DeDE: "Verarbeitet",
		EnUS: "Processed",
		EsES: "Procesada",
		FrFR: "Traité",
		ItIT: "Elaborato",
		PlPL: "Obrobiony",
		PtPT: "Processado",
		SkSK: "Spracované",
		RuRU: "Обработанный",
	}
	db.Create(&alarmProcessed)

	breakdownStart := database.Locale{
		Name: "breakdown-start",
		CsCZ: "Začátek",
		DeDE: "Anfang",
		EnUS: "Beginning",
		EsES: "Comenzando",
		FrFR: "Début",
		ItIT: "Inizio",
		PlPL: "Początek",
		PtPT: "Começo",
		SkSK: "Začiatok",
		RuRU: "Начало",
	}
	db.Create(&breakdownStart)

	breakdownEnd := database.Locale{
		Name: "breakdown-end",
		CsCZ: "Konec",
		DeDE: "Ende",
		EnUS: "End",
		EsES: "Fin",
		FrFR: "Finir",
		ItIT: "Fine",
		PlPL: "Koniec",
		PtPT: "Fim",
		SkSK: "Koniec",
		RuRU: "Конец",
	}
	db.Create(&breakdownEnd)

	breakdownName := database.Locale{
		Name: "breakdown-name",
		CsCZ: "Porucha",
		DeDE: "Störung",
		EnUS: "Breakdown",
		EsES: "Trastorno",
		FrFR: "Désordre",
		ItIT: "Disturbo",
		PlPL: "Nieład",
		PtPT: "Transtorno",
		SkSK: "Porucha",
		RuRU: "Беспорядок",
	}
	db.Create(&breakdownName)

	downtimeStart := database.Locale{
		Name: "downtime-start",
		CsCZ: "Začátek",
		DeDE: "Anfang",
		EnUS: "Beginning",
		EsES: "Comenzando",
		FrFR: "Début",
		ItIT: "Inizio",
		PlPL: "Początek",
		PtPT: "Começo",
		SkSK: "Začiatok",
		RuRU: "Начало",
	}
	db.Create(&downtimeStart)

	downtimeEnd := database.Locale{
		Name: "downtime-end",
		CsCZ: "Konec",
		DeDE: "Ende",
		EnUS: "End",
		EsES: "Fin",
		FrFR: "Finir",
		ItIT: "Fine",
		PlPL: "Koniec",
		PtPT: "Fim",
		SkSK: "Koniec",
		RuRU: "Конец",
	}
	db.Create(&downtimeEnd)

	downtimeName := database.Locale{
		Name: "downtime-name",
		CsCZ: "Prostoj",
		DeDE: "Ausfallzeit",
		EnUS: "Downtime",
		EsES: "Inactividad",
		FrFR: "Inactivité",
		ItIT: "Inattività",
		PlPL: "Przestoj",
		PtPT: "Inatividade",
		SkSK: "Prestoj",
		RuRU: "Простой",
	}
	db.Create(&downtimeName)

	faultDate := database.Locale{
		Name: "fault-date",
		CsCZ: "Datum",
		DeDE: "Datum",
		EnUS: "Date",
		EsES: "Fecha",
		FrFR: "Date",
		ItIT: "Data",
		PlPL: "Data",
		PtPT: "Data",
		SkSK: "Dátum",
		RuRU: "Дата",
	}
	db.Create(&faultDate)

	faultName := database.Locale{
		Name: "fault-name",
		CsCZ: "Chyba",
		DeDE: "Fehler",
		EnUS: "Fault",
		EsES: "Falla",
		FrFR: "Défaut",
		ItIT: "Difetto",
		PlPL: "Wada",
		PtPT: "Falha",
		SkSK: "Chyba",
		RuRU: "Недостаток",
	}
	db.Create(&faultName)

	faultCount := database.Locale{
		Name: "fault-count",
		CsCZ: "Množství",
		DeDE: "Menge",
		EnUS: "Amount",
		EsES: "Cantidad",
		FrFR: "Quantité",
		ItIT: "Quantità",
		PlPL: "Ilość",
		PtPT: "Quantia",
		SkSK: "Množstvo",
		RuRU: "Количество",
	}
	db.Create(&faultCount)

	packageDate := database.Locale{
		Name: "package-date",
		CsCZ: "Datum",
		DeDE: "Datum",
		EnUS: "Date",
		EsES: "Fecha",
		FrFR: "Date",
		ItIT: "Data",
		PlPL: "Data",
		PtPT: "Data",
		SkSK: "Dátum",
		RuRU: "Дата",
	}
	db.Create(&packageDate)

	packageName := database.Locale{
		Name: "package-name",
		CsCZ: "Balení",
		DeDE: "Paket",
		EnUS: "Package",
		EsES: "Paquete",
		FrFR: "Paquet",
		ItIT: "Pacchetto",
		PlPL: "Pakiet",
		PtPT: "Pacote",
		SkSK: "Baleni",
		RuRU: "Пакет",
	}
	db.Create(&packageName)

	packageCount := database.Locale{
		Name: "package-count",
		CsCZ: "Množství",
		DeDE: "Menge",
		EnUS: "Amount",
		EsES: "Cantidad",
		FrFR: "Quantité",
		ItIT: "Quantità",
		PlPL: "Ilość",
		PtPT: "Quantia",
		SkSK: "Množstvo",
		RuRU: "Количество",
	}
	db.Create(&packageCount)

	partDate := database.Locale{
		Name: "part-date",
		CsCZ: "Datum",
		DeDE: "Datum",
		EnUS: "Date",
		EsES: "Fecha",
		FrFR: "Date",
		ItIT: "Data",
		PlPL: "Data",
		PtPT: "Data",
		SkSK: "Dátum",
		RuRU: "Дата",
	}
	db.Create(&partDate)

	partName := database.Locale{
		Name: "part-name",
		CsCZ: "Díl",
		DeDE: "Teil",
		EnUS: "Part",
		EsES: "Parte",
		FrFR: "Partie",
		ItIT: "Parte",
		PlPL: "Część",
		PtPT: "Papel",
		SkSK: "Diel",
		RuRU: "Часть",
	}
	db.Create(&partName)

	partCount := database.Locale{
		Name: "part-count",
		CsCZ: "Množství",
		DeDE: "Menge",
		EnUS: "Amount",
		EsES: "Cantidad",
		FrFR: "Quantité",
		ItIT: "Quantità",
		PlPL: "Ilość",
		PtPT: "Quantia",
		SkSK: "Množstvo",
		RuRU: "Количество",
	}
	db.Create(&partCount)

	stateStart := database.Locale{
		Name: "state-start",
		CsCZ: "Začátek",
		DeDE: "Anfang",
		EnUS: "Beginning",
		EsES: "Comenzando",
		FrFR: "Début",
		ItIT: "Inizio",
		PlPL: "Początek",
		PtPT: "Começo",
		SkSK: "Začiatok",
		RuRU: "Начало",
	}
	db.Create(&stateStart)

	stateEnd := database.Locale{
		Name: "state-end",
		CsCZ: "Konec",
		DeDE: "Ende",
		EnUS: "End",
		EsES: "Fin",
		FrFR: "Finir",
		ItIT: "Fine",
		PlPL: "Koniec",
		PtPT: "Fim",
		SkSK: "Koniec",
		RuRU: "Конец",
	}
	db.Create(&stateEnd)

	stateName := database.Locale{
		Name: "state-name",
		CsCZ: "Stav",
		DeDE: "Zustand",
		EnUS: "State",
		EsES: "Expresar",
		FrFR: "État",
		ItIT: "Stato",
		PlPL: "Stan",
		PtPT: "Estado",
		SkSK: "Stav",
		RuRU: "Состояние",
	}
	db.Create(&stateName)

	userStart := database.Locale{
		Name: "user-start",
		CsCZ: "Začátek",
		DeDE: "Anfang",
		EnUS: "Beginning",
		EsES: "Comenzando",
		FrFR: "Début",
		ItIT: "Inizio",
		PlPL: "Początek",
		PtPT: "Começo",
		SkSK: "Začiatok",
		RuRU: "Начало",
	}
	db.Create(&userStart)

	userEnd := database.Locale{
		Name: "user-end",
		CsCZ: "Konec",
		DeDE: "Ende",
		EnUS: "End",
		EsES: "Fin",
		FrFR: "Finir",
		ItIT: "Fine",
		PlPL: "Koniec",
		PtPT: "Fim",
		SkSK: "Koniec",
		RuRU: "Конец",
	}
	db.Create(&userEnd)

	systemDate := database.Locale{
		Name: "system-date",
		CsCZ: "Datum",
		DeDE: "Datum",
		EnUS: "Date",
		EsES: "Fecha",
		FrFR: "Date",
		ItIT: "Data",
		PlPL: "Data",
		PtPT: "Data",
		SkSK: "Dátum",
		RuRU: "Дата",
	}
	db.Create(&systemDate)

	databaseSize := database.Locale{
		Name: "database-size",
		CsCZ: "Velikost databáze",
		DeDE: "Datenbankgröße",
		EnUS: "Database size",
		EsES: "Tamaño de la base de datos",
		FrFR: "Taille de la base de données",
		ItIT: "Dimensioni del database",
		PlPL: "Rozmiar bazy danych",
		PtPT: "Tamanho do banco de dados",
		SkSK: "Veľkosť databázy",
		RuRU: "Размер базы данных",
	}
	db.Create(&databaseSize)

	databaseGrowth := database.Locale{
		Name: "database-growth",
		CsCZ: "Nárůst databáze",
		DeDE: "Datenbankwachstum",
		EnUS: "Database growth",
		EsES: "Crecimiento de la base de datos",
		FrFR: "Croissance de la base de données",
		ItIT: "Crescita del database",
		PlPL: "Rozwój bazy danych",
		PtPT: "Crescimento do banco de dados",
		SkSK: "Rast databázy",
		RuRU: "Рост базы данных",
	}
	db.Create(&databaseGrowth)

	discFreeSpace := database.Locale{
		Name: "disc-free-space",
		CsCZ: "Volné místo na disku",
		DeDE: "Disc freier Speicherplatz",
		EnUS: "Disc free space",
		EsES: "Espacio libre en disco",
		FrFR: "Espace libre sur le disque",
		ItIT: "Spazio libero su disco",
		PlPL: "Wolne miejsce na dysku",
		PtPT: "Espaço livre no disco",
		SkSK: "Voľné miesto na disku",
		RuRU: "Свободное место на диске",
	}
	db.Create(&discFreeSpace)

	estimatedFreeSpace := database.Locale{
		Name: "estimated-free-space",
		CsCZ: "Odhad dnů volného místa",
		DeDE: "Geschätzte Tage freier Speicherplatz",
		EnUS: "Estimated days free space",
		EsES: "Días estimados de espacio libre",
		FrFR: "Nombre estimé de jours d'espace libre",
		ItIT: "Giorni stimati di spazio libero",
		PlPL: "Szacunkowa liczba dni wolnego miejsca",
		PtPT: "Dias estimados de espaço livre",
		SkSK: "Odhadovaný počet dní voľného miesta",
		RuRU: "Приблизительное количество дней свободного места",
	}
	db.Create(&estimatedFreeSpace)

	combinedChart := database.Locale{
		Name: "combined-chart",
		CsCZ: "Kombinovaný graf",
		DeDE: "Kombiniertes Diagramm",
		EnUS: "Combined chart",
		EsES: "Gráfico combinado",
		FrFR: "Graphique combiné",
		ItIT: "Grafico combinato",
		PlPL: "Wykres łączony",
		PtPT: "Gráfico combinado",
		SkSK: "Kombinovaný graf",
		RuRU: "Комбинированная диаграмма",
	}
	db.Create(&combinedChart)

	timelineChart := database.Locale{
		Name: "timeline-chart",
		CsCZ: "Časový graf",
		DeDE: "Zeitdiagramm",
		EnUS: "Timeline chart",
		EsES: "Gráfico de línea de tiempo",
		FrFR: "Graphique chronologique",
		ItIT: "Grafico cronologico",
		PlPL: "Wykres z osią czasu",
		PtPT: "Gráfico de linha do tempo",
		SkSK: "Časový graf",
		RuRU: "График времени",
	}
	db.Create(&timelineChart)

	analogData := database.Locale{
		Name: "analog-data",
		CsCZ: "Analogová data",
		DeDE: "Analoge Daten",
		EnUS: "Analog data",
		EsES: "Datos analógicos",
		FrFR: "Données analogiques",
		ItIT: "Dati analogici",
		PlPL: "Dane analogowe",
		PtPT: "Dados analógicos",
		SkSK: "Analógové dáta",
		RuRU: "Аналоговые данные",
	}
	db.Create(&analogData)

	digitalData := database.Locale{
		Name: "digital-data",
		CsCZ: "Digitální data",
		DeDE: "Digitale Daten",
		EnUS: "Digital data",
		EsES: "Información digital",
		FrFR: "Données numériques",
		ItIT: "Dati digitali",
		PlPL: "Dane cyfrowe",
		PtPT: "Dados digitais",
		SkSK: "Digitálné data",
		RuRU: "Цифровые данные",
	}
	db.Create(&digitalData)

	productionChart := database.Locale{
		Name: "production-chart",
		CsCZ: "Graf produkce",
		DeDE: "Produktionsdiagramm",
		EnUS: "Production chart",
		EsES: "Tabla de producción",
		FrFR: "Tableau de production",
		ItIT: "Grafico di produzione",
		PlPL: "Wykres produkcyjny",
		PtPT: "Gráfico de produção",
		SkSK: "Graf výroby",
		RuRU: "График производства",
	}
	db.Create(&productionChart)

	consumptionChart := database.Locale{
		Name: "consumption-chart",
		CsCZ: "Graf spotřeby",
		DeDE: "Verbrauchsdiagramm",
		EnUS: "Consumption chart",
		EsES: "Gráfico de consumo",
		FrFR: "Graphique de consommation",
		ItIT: "Grafico dei consumi",
		PlPL: "Wykres zużycia",
		PtPT: "Gráfico de consumo",
		SkSK: "Graf spotreby",
		RuRU: "График потребления",
	}
	db.Create(&consumptionChart)

	devices := database.Locale{
		Name: "devices",
		CsCZ: "Zařízení",
		DeDE: "Geräte",
		EnUS: "Devices",
		EsES: "Dispositivos",
		FrFR: "Dispositifs",
		ItIT: "Dispositivi",
		PlPL: "Urządzenia",
		PtPT: "Dispositivos",
		SkSK: "Zariadenie",
		RuRU: "Устройства",
	}
	db.Create(&devices)

	workplaces := database.Locale{
		Name: "workplaces",
		CsCZ: "Pracoviště",
		DeDE: "Arbeitsplätze",
		EnUS: "Workplaces",
		EsES: "Lugares de trabajo",
		FrFR: "Lieux de travail",
		ItIT: "Ambienti di lavoro",
		PlPL: "Miejsca pracy",
		PtPT: "Locais de trabalho",
		SkSK: "Pracoviska",
		RuRU: "Рабочие места",
	}
	db.Create(&workplaces)

	products := database.Locale{
		Name: "products",
		CsCZ: "Produkty",
		DeDE: "Produkte",
		EnUS: "Products",
		EsES: "Productos",
		FrFR: "Produits",
		ItIT: "Prodotti",
		PlPL: "Produkty",
		PtPT: "Produtos",
		SkSK: "Produkty",
		RuRU: "Товары",
	}
	db.Create(&products)

	operations := database.Locale{
		Name: "operations",
		CsCZ: "Operace",
		DeDE: "Operationen",
		EnUS: "Operations",
		EsES: "Operaciones",
		FrFR: "Opérations",
		ItIT: "Operazioni",
		PlPL: "Operacje",
		PtPT: "Operações",
		SkSK: "Operácie",
		RuRU: "Операции",
	}
	db.Create(&operations)

	workshifts := database.Locale{
		Name: "workshifts",
		CsCZ: "Pracovní směny",
		DeDE: "Arbeitsschichten",
		EnUS: "Workshifts",
		EsES: "Turnos de trabajo",
		FrFR: "Quarts de travail",
		ItIT: "Turni lavorativi",
		PlPL: "Zmiany robocze",
		PtPT: "Turnos de trabalho",
		SkSK: "Pracovné zmeny",
		RuRU: "Рабочие смены",
	}
	db.Create(&workshifts)

	systemSettings := database.Locale{
		Name: "system-settings",
		CsCZ: "Nastavení",
		DeDE: "Einstellungen",
		EnUS: "Settings",
		EsES: "Ajustes",
		FrFR: "Réglages",
		ItIT: "Impostazioni",
		PlPL: "Ustawienia",
		PtPT: "Definições",
		SkSK: "Nastavenie",
		RuRU: "Hастройки",
	}
	db.Create(&systemSettings)

	sqlCommand := database.Locale{
		Name: "sql-command",
		CsCZ: "Sql příkaz",
		DeDE: "SQL-Befehl",
		EnUS: "Sql command",
		EsES: "Comando sql",
		FrFR: "Commande SQL",
		ItIT: "Comando SQL",
		PlPL: "Sql polecenie",
		PtPT: "Comando Sql",
		SkSK: "Sql príkaz",
		RuRU: "Команда sql",
	}
	db.Create(&sqlCommand)

	messageHeader := database.Locale{
		Name: "message-header",
		CsCZ: "Předmět zprávy",
		DeDE: "Betreff der Nachricht",
		EnUS: "Message subject",
		EsES: "Asunto del mensaje",
		FrFR: "Objet du message",
		ItIT: "Soggetto del messaggio",
		PlPL: "Temat wiadomości",
		PtPT: "Assunto da mensagem",
		SkSK: "Predmet správy",
		RuRU: "Тема сообщения",
	}
	db.Create(&messageHeader)

	messageText := database.Locale{
		Name: "message-text",
		CsCZ: "Text zprávy",
		DeDE: "Nachrichtentext",
		EnUS: "Message text",
		EsES: "Mensaje de texto",
		FrFR: "Texte du message",
		ItIT: "Messaggio di testo",
		PlPL: "Wiadomość tekstowa",
		PtPT: "Mensagem de texto",
		SkSK: "Text správy",
		RuRU: "Текст сообщения",
	}
	db.Create(&messageText)

	recipients := database.Locale{
		Name: "recipients",
		CsCZ: "Příjemci",
		DeDE: "Empfänger",
		EnUS: "Recipients",
		EsES: "Destinatarios",
		FrFR: "Bénéficiaires",
		ItIT: "Destinatari",
		PlPL: "Odbiorcy",
		PtPT: "Destinatários",
		SkSK: "Príjemcovia",
		RuRU: "Получатели",
	}
	db.Create(&recipients)

	url := database.Locale{
		Name: "url",
		CsCZ: "Url",
		DeDE: "Url",
		EnUS: "Url",
		EsES: "URL",
		FrFR: "URL",
		ItIT: "Url",
		PlPL: "Url",
		PtPT: "Url",
		SkSK: "Url",
		RuRU: "URL",
	}
	db.Create(&url)

	pdf := database.Locale{
		Name: "pdf",
		CsCZ: "Pdf",
		DeDE: "Pdf",
		EnUS: "Pdf",
		EsES: "PDF",
		FrFR: "PDF",
		ItIT: "Pdf",
		PlPL: "Pdf",
		PtPT: "Pdf",
		SkSK: "Pdf",
		RuRU: "PDF",
	}
	db.Create(&pdf)

	createdAt := database.Locale{
		Name: "created-at",
		CsCZ: "Vytvořeno v",
		DeDE: "Hergestellt am",
		EnUS: "Created at",
		EsES: "Creado en",
		FrFR: "Créé à",
		ItIT: "Creato a",
		PlPL: "Utworzono w",
		PtPT: "Criado em",
		SkSK: "Vytvorené v",
		RuRU: "Создано на",
	}
	db.Create(&createdAt)

	updatedAt := database.Locale{
		Name: "updated-at",
		CsCZ: "Aktualizováno v",
		DeDE: "Aktualisiert am",
		EnUS: "Updated at",
		EsES: "Actualizado en",
		FrFR: "Mis à jour à",
		ItIT: "Aggiornato a",
		PlPL: "Zaktualizowano o godz",
		PtPT: "Atualizado em",
		SkSK: "Aktualizované o",
		RuRU: "Обновлено в",
	}
	db.Create(&updatedAt)

	productName := database.Locale{
		Name: "product-name",
		CsCZ: "Produkt",
		DeDE: "Produkt",
		EnUS: "Product",
		EsES: "Producto",
		FrFR: "Produit",
		ItIT: "Prodotto",
		PlPL: "Produkt",
		PtPT: "Produtos",
		SkSK: "Produkt",
		RuRU: "Tовар",
	}
	db.Create(&productName)

	deviceName := database.Locale{
		Name: "device-name",
		CsCZ: "Zařízení",
		DeDE: "Gerät",
		EnUS: "Device",
		EsES: "Aparato",
		FrFR: "L'appareil",
		ItIT: "Dispositivo",
		PlPL: "Urządzenie",
		PtPT: "Dispositivo",
		SkSK: "Zariadenie",
		RuRU: "Yстройство",
	}
	db.Create(&deviceName)

	typeName := database.Locale{
		Name: "type-name",
		CsCZ: "Typ",
		DeDE: "Sorte",
		EnUS: "Type",
		EsES: "Escribe",
		FrFR: "Catégorie",
		ItIT: "Genere",
		PlPL: "Rodzaj",
		PtPT: "Modelo",
		SkSK: "Typ",
		RuRU: "Тип",
	}
	db.Create(&typeName)

	name := database.Locale{
		Name: "name",
		CsCZ: "Jméno",
		DeDE: "Name",
		EnUS: "Name",
		EsES: "Nombre",
		FrFR: "Nom",
		ItIT: "Nome",
		PlPL: "Nazva",
		PtPT: "Nome",
		SkSK: "Meno",
		RuRU: "Имя",
	}
	db.Create(&name)

	barcode := database.Locale{
		Name: "barcode",
		CsCZ: "Čárový kód",
		DeDE: "Barcode",
		EnUS: "Barcode",
		EsES: "Código de barras",
		FrFR: "Code à barre",
		ItIT: "Codice a barre",
		PlPL: "Kod kreskowy",
		PtPT: "Código de barras",
		SkSK: "Čiarový kód",
		RuRU: "Штрих-код",
	}
	db.Create(&barcode)

	dateRequest := database.Locale{
		Name: "date-requested",
		CsCZ: "Požadovaný datum výroby",
		DeDE: "Erforderliches Herstellungsdatum",
		EnUS: "Required date of manufacture",
		EsES: "Fecha requerida de fabricación",
		FrFR: "Date de fabrication requise",
		ItIT: "Data di produzione richiesta",
		PlPL: "Wymagana data produkcji",
		PtPT: "Data de fabricação exigida",
		SkSK: "Požadovaný dátum výroby",
		RuRU: "Требуемая дата изготовления",
	}
	db.Create(&dateRequest)

	countRequest := database.Locale{
		Name: "count-requested",
		CsCZ: "Požadované množství",
		DeDE: "Angeforderter Betrag",
		EnUS: "Count requested",
		EsES: "Monto requerido",
		FrFR: "Quantité exigée",
		ItIT: "Importo richiesto",
		PlPL: "Wnioskowana kwota",
		PtPT: "Quantidade solicitada",
		SkSK: "Požadované množstvo",
		RuRU: "Запрошенная сумма",
	}
	db.Create(&countRequest)

	downtimeDuration := database.Locale{
		Name: "downtime-duration",
		CsCZ: "Doba prostoje",
		DeDE: "Ausfallzeitdauer",
		EnUS: "Downtime duration",
		EsES: "Duración del tiempo de inactividad",
		FrFR: "Durée des temps d'arrêt",
		ItIT: "Durata del tempo di inattività",
		PlPL: "Czas przestoju",
		PtPT: "Duração do tempo de inatividade",
		SkSK: "Trvanie prestoje",
		RuRU: "Продолжительность простоя",
	}
	db.Create(&downtimeDuration)

	color := database.Locale{
		Name: "color",
		CsCZ: "Barva",
		DeDE: "Farbe",
		EnUS: "Color",
		EsES: "Color",
		FrFR: "Couleur",
		ItIT: "Colore",
		PlPL: "Kolor",
		PtPT: "Cor",
		SkSK: "Farba",
		RuRU: "Цвет",
	}
	db.Create(&color)

	workshiftStart := database.Locale{
		Name: "workshift-start",
		CsCZ: "Začátek",
		DeDE: "Anfang",
		EnUS: "Beginning",
		EsES: "Comenzando",
		FrFR: "Début",
		ItIT: "Inizio",
		PlPL: "Początek",
		PtPT: "Começo",
		SkSK: "Začiatok",
		RuRU: "Начало",
	}
	db.Create(&workshiftStart)

	workshiftEnd := database.Locale{
		Name: "workshift-end",
		CsCZ: "Konec",
		DeDE: "Ende",
		EnUS: "End",
		EsES: "Fin",
		FrFR: "Finir",
		ItIT: "Fine",
		PlPL: "Koniec",
		PtPT: "Fim",
		SkSK: "Koniec",
		RuRU: "Конец",
	}
	db.Create(&workshiftEnd)

	firstName := database.Locale{
		Name: "first-name",
		CsCZ: "Jméno",
		DeDE: "Vorname",
		EnUS: "First name",
		EsES: "Nombre de pila",
		FrFR: "Prénom",
		ItIT: "Nome di battesimo",
		PlPL: "Imię",
		PtPT: "Primeiro nome",
		SkSK: "Meno",
		RuRU: "Имя",
	}
	db.Create(&firstName)

	secondName := database.Locale{
		Name: "second-name",
		CsCZ: "Příjmení",
		DeDE: "Nachname",
		EnUS: "First name",
		EsES: "Apellido",
		FrFR: "Nom de famille",
		ItIT: "Cognome",
		PlPL: "Nazwisko",
		PtPT: "Sobrenome",
		SkSK: "Priezvisko",
		RuRU: "Фамилия",
	}
	db.Create(&secondName)

	roleName := database.Locale{
		Name: "role-name",
		CsCZ: "Role",
		DeDE: "Rolle",
		EnUS: "Role",
		EsES: "Oficio",
		FrFR: "Rôle",
		ItIT: "Ruolo",
		PlPL: "Rola",
		PtPT: "Função",
		SkSK: "Rola",
		RuRU: "Роль",
	}
	db.Create(&roleName)

	email := database.Locale{
		Name: "email",
		CsCZ: "E-mail",
		DeDE: "E-mail",
		EnUS: "E-mail",
		EsES: "E-mail",
		FrFR: "E-mail",
		ItIT: "E-mail",
		PlPL: "E-mail",
		PtPT: "E-mail",
		SkSK: "E-mail",
		RuRU: "E-мейл",
	}
	db.Create(&email)

	password := database.Locale{
		Name: "password",
		CsCZ: "Heslo",
		DeDE: " Kennwort",
		EnUS: "Password",
		EsES: "Contraseña",
		FrFR: "Mot de passe",
		ItIT: "Parola d'ordine",
		PlPL: "Hasło",
		PtPT: "Senha",
		SkSK: "Heslo",
		RuRU: "Пароль",
	}
	db.Create(&password)

	phone := database.Locale{
		Name: "phone",
		CsCZ: "Telefon",
		DeDE: "Telefon",
		EnUS: "Phone",
		EsES: "Teléfono",
		FrFR: "Téléphone",
		ItIT: "Telefono",
		PlPL: "Telefon",
		PtPT: "Telefone",
		SkSK: "Telefón",
		RuRU: "Телефон",
	}
	db.Create(&phone)

	position := database.Locale{
		Name: "position",
		CsCZ: "Pozice",
		DeDE: "Funktion",
		EnUS: "Position",
		EsES: "Posición",
		FrFR: "Positionner",
		ItIT: "Posizione",
		PlPL: "Pozycja",
		PtPT: "Posição",
		SkSK: "Pozícia",
		RuRU: "Позиция",
	}
	db.Create(&position)

	pin := database.Locale{
		Name: "pin",
		CsCZ: "Osobní identifikační číslo",
		DeDE: "Persönliche Identifikationsnummer",
		EnUS: "Personal identification number",
		EsES: "Número de identificación personal",
		FrFR: "Numéro d'identification personnel",
		ItIT: "Numero di identificazione personale",
		PlPL: "Osobisty numer identyfikacyjny",
		PtPT: "Número de identificação pessoal",
		SkSK: "Osobné identifikačné číslo",
		RuRU: "Персональный идентификационный номер",
	}
	db.Create(&pin)

	rfid := database.Locale{
		Name: "rfid",
		CsCZ: "RFID číslo",
		DeDE: "RFID-Nummer.",
		EnUS: "RFID number",
		EsES: "Número de RFID",
		FrFR: "Nombre RFID",
		ItIT: "Numero RFID",
		PlPL: "Numer RFID.",
		PtPT: "Número do RFID.",
		SkSK: "Číslo RFID",
		RuRU: "RFID номер",
	}
	db.Create(&rfid)

	locale := database.Locale{
		Name: "locale",
		CsCZ: "Jazyk",
		DeDE: "Sprache",
		EnUS: "Language",
		EsES: "Idioma",
		FrFR: "Langue",
		ItIT: "Linguaggio",
		PlPL: "Język",
		PtPT: "Língua",
		SkSK: "Jazyk",
		RuRU: "Язык",
	}
	db.Create(&locale)

	value := database.Locale{
		Name: "value",
		CsCZ: "Hodnota",
		DeDE: "Wert",
		EnUS: "Value",
		EsES: "Valor",
		FrFR: "Évaluer",
		ItIT: "Valore",
		PlPL: "Wartość",
		PtPT: "Valor",
		SkSK: "Hodnota",
		RuRU: "Значение",
	}
	db.Create(&value)

	enabled := database.Locale{
		Name: "enabled",
		CsCZ: "Povoleno",
		DeDE: "Aktiviert",
		EnUS: "Enabled",
		EsES: "Activado",
		FrFR: "Activé",
		ItIT: "Abilitato",
		PlPL: "Włączony",
		PtPT: "Habilitado",
		SkSK: "Povolené",
		RuRU: "Включено",
	}
	db.Create(&enabled)

	ipAddress := database.Locale{
		Name: "ip-address",
		CsCZ: "IP adresa",
		DeDE: "IP Adresse",
		EnUS: "IP address",
		EsES: "Dirección IP",
		FrFR: "Adresse IP",
		ItIT: "Indirizzo IP",
		PlPL: "Adres IP",
		PtPT: "Endereço de IP",
		SkSK: "IP adresa",
		RuRU: "IP адрес",
	}
	db.Create(&ipAddress)

	macAddress := database.Locale{
		Name: "mac-address",
		CsCZ: "MAC adresa",
		DeDE: "MAC Adresse",
		EnUS: "MAC address",
		EsES: "Dirección MAC",
		FrFR: "Adresse MAC",
		ItIT: "Indirizzo MAC",
		PlPL: "Adres MAC",
		PtPT: "Endereço de MAC",
		SkSK: "MAC adresa",
		RuRU: "MAC-адрес",
	}
	db.Create(&macAddress)

	deviceVersion := database.Locale{
		Name: "device-version",
		CsCZ: "Verze",
		DeDE: "Ausführung",
		EnUS: "Version",
		EsES: "Versión",
		FrFR: "Version",
		ItIT: "Versione",
		PlPL: "Wersja",
		PtPT: "Versão",
		SkSK: "Verzia",
		RuRU: "Версия",
	}
	db.Create(&deviceVersion)

	deviceSettings := database.Locale{
		Name: "device-settings",
		CsCZ: "Nastavení",
		DeDE: "Einstellungen",
		EnUS: "Settings",
		EsES: "Ajustes",
		FrFR: "Réglages",
		ItIT: "Impostazioni",
		PlPL: "Ustawienia",
		PtPT: "Definições",
		SkSK: "Nastavenie",
		RuRU: "Hастройки",
	}
	db.Create(&deviceSettings)

	portName := database.Locale{
		Name: "port-name",
		CsCZ: "Jméno",
		DeDE: "Name",
		EnUS: "Name",
		EsES: "Nombre",
		FrFR: "Nom",
		ItIT: "Nome",
		PlPL: "Nazva",
		PtPT: "Nome",
		SkSK: "Meno",
		RuRU: "Имя",
	}
	db.Create(&portName)

	portTypeName := database.Locale{
		Name: "port-type-name",
		CsCZ: "Typ",
		DeDE: "Sorte",
		EnUS: "Type",
		EsES: "Escribe",
		FrFR: "Catégorie",
		ItIT: "Genere",
		PlPL: "Rodzaj",
		PtPT: "Modelo",
		SkSK: "Typ",
		RuRU: "Тип",
	}
	db.Create(&portTypeName)

	portUnit := database.Locale{
		Name: "port-unit",
		CsCZ: "Jednotka",
		DeDE: "Einheit",
		EnUS: "Unit",
		EsES: "Unidad",
		FrFR: "Unité",
		ItIT: "Unità",
		PlPL: "Jednostka",
		PtPT: "Unidade",
		SkSK: "Jednotka",
		RuRU: "Единица измерения",
	}
	db.Create(&portUnit)

	plcDataType := database.Locale{
		Name: "plc-data-type",
		CsCZ: "PLC data typ",
		DeDE: "SPS-Datentyp",
		EnUS: "PLC data type",
		EsES: "Tipo de datos PLC",
		FrFR: "Type de données PLC",
		ItIT: "Tipo di dati PLC",
		PlPL: "Typ danych PLC",
		PtPT: "Tipo de dados PLC",
		SkSK: "Typ údajov PLC",
		RuRU: "Тип данных PLC",
	}
	db.Create(&plcDataType)

	plcDataAddress := database.Locale{
		Name: "plc-data-address",
		CsCZ: "PLC adresa",
		DeDE: "SPS-Datenadresse",
		EnUS: "PLC data address",
		EsES: "Dirección de datos PLC",
		FrFR: "Adresse de données PLC",
		ItIT: "Indirizzo dei dati del PLC",
		PlPL: "Adres danych PLC",
		PtPT: "Endereço de dados do PLC.",
		SkSK: "Adresa údajov PLC",
		RuRU: "Адрес данных PLC",
	}
	db.Create(&plcDataAddress)

	portVirtual := database.Locale{
		Name: "port-virtual",
		CsCZ: "Virtuální",
		DeDE: "Virtuell",
		EnUS: "Virtual",
		EsES: "Virtual",
		FrFR: "Virtuel",
		ItIT: "Virtuale",
		PlPL: "Wirtualny",
		PtPT: "Virtual",
		SkSK: "Virtuálne",
		RuRU: "Виртуальный",
	}
	db.Create(&portVirtual)

	portNumber := database.Locale{
		Name: "port-number",
		CsCZ: "Číslo",
		DeDE: "Nummer",
		EnUS: "Number",
		EsES: "Número",
		FrFR: "Nombre",
		ItIT: "Numero",
		PlPL: "Numer",
		PtPT: "Número",
		SkSK: "Číslo",
		RuRU: "Число",
	}
	db.Create(&portNumber)

	poweroffDuration := database.Locale{
		Name: "poweroff-duration",
		CsCZ: "Doba vypnutí",
		DeDE: "Ausschaltdauer",
		EnUS: "Poweroff duration",
		EsES: "Duración del apagado",
		FrFR: "Durée de la mise hors tension",
		ItIT: "Durata dello spegnimento",
		PlPL: "Czas wyłączenia zasilania",
		PtPT: "Duração do desligamento",
		SkSK: "Trvanie vypnutia",
		RuRU: "Продолжительность выключения питания",
	}
	db.Create(&poweroffDuration)

	sectionName := database.Locale{
		Name: "section-name",
		CsCZ: "Skupina",
		DeDE: "Gruppe",
		EnUS: "Group",
		EsES: "Grupo",
		FrFR: "Groupe",
		ItIT: "Gruppo",
		PlPL: "Grupa",
		PtPT: "Grupo",
		SkSK: "Skupina",
		RuRU: "Группа",
	}
	db.Create(&sectionName)

	code := database.Locale{
		Name: "code-name",
		CsCZ: "Kód",
		DeDE: "Code",
		EnUS: "Code",
		EsES: "Código",
		FrFR: "Code",
		ItIT: "Codice",
		PlPL: "Kod",
		PtPT: "Código",
		SkSK: "Kód",
		RuRU: "Код",
	}
	db.Create(&code)
}
