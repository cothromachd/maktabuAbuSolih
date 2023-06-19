package main

import (
	"log"
	"strings"
	"time"

	tele "gopkg.in/telebot.v3"
)

var (
	welcomeText = `السلام عليكم ورحمة الله وبركات
	какой раздел тебе нужен?`
	m = make(map[string][2]any)
)

func main() {
	pref := tele.Settings{
		Token:     "6088172429:AAE3DZF6abqLO-T3k50H5FODg6yi__NWqeQ",
		Poller:    &tele.LongPoller{Timeout: 20 * time.Second},
		ParseMode: tele.ModeHTML,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	var (
		mainMarkup           = &tele.ReplyMarkup{}
		mainAdab             = mainMarkup.Data("Адаб", "adabBtn")
		mainAkida            = mainMarkup.Data("Акыда", "akidaBtn")
		mainDifferent        = mainMarkup.Data("Разное", "differentBtn")
		mainHadithCollection = mainMarkup.Data("Сборники хадисов", "hadithCollectionBtn")
		mainSira             = mainMarkup.Data("Сира", "siraBtn")
		mainTafsir           = mainMarkup.Data("Тафсир", "tafsirBtn")
		mainKnowledgeNeeding = mainMarkup.Data("Требование знаний", "knowledgeNeedingBtn")
		mainFikkh            = mainMarkup.Data("Фикх", "fikkhBtn")

		akidaMarkup = &tele.ReplyMarkup{}
		akidaTo     = akidaMarkup.Data("Три основы", "toBtn")
		akidaChp    = akidaMarkup.Data("Четыре правила", "chpBtn")
		akidaSho    = akidaMarkup.Data("Шесть основ", "shoBtn")
		akidaOs     = akidaMarkup.Data("Отведение сомнений", "osBtn")
		akidaKe     = akidaMarkup.Data("Книга единобожия", "keBtn")
		akidaNi     = akidaMarkup.Data("Навакыдуль ислям", "niBtn")
		akidaAviia  = akidaMarkup.Data("Акыда в именах и атрибутах", "aviiaBtn")
		akidaVi     = akidaMarkup.Data("Вопрос имана", "viBtn")
		akidaLin    = akidaMarkup.Data("Любовь и непричастность", "linBtn")
		akidaRvubd  = akidaMarkup.Data("Разбор вопроса узр биль джахль", "rvubdBtn")
		akidaKpoaim = akidaMarkup.Data("Книги по основам акыды и манхаджа", "kpoaimBtn")
		akidaPkpa   = akidaMarkup.Data("Полезные книги по акыде", "pkpaBtn")

		fikkhMarkup   = &tele.ReplyMarkup{}
		fikkhBrak     = fikkhMarkup.Data("Брак", "brakBtn")
		fikkhJihad    = fikkhMarkup.Data("Джихад", "jihadBtn")
		fikkhClearing = fikkhMarkup.Data("Очищение", "clearingBtn")
		fikkhWorship  = fikkhMarkup.Data("Поклонение", "worshipBtn")
		fikkhFow      = fikkhMarkup.Data("Фикх женщин", "fowBtn")
		fikkhFinance  = fikkhMarkup.Data("Финансы", "financeBtn")

		selector             = &tele.ReplyMarkup{}
		selectorBackAtAllBtn = selector.Data("В начало", "inMainBtn")
		selectorBack         = selector.Data("⬅️", "backBtn")
	)

	m["Адаб"] = [2]any{`السلام عليكم ورحمة الله وبركاته
	какой раздел тебе нужен?`, mainMarkup}
	m["Акыда"] = [2]any{`السلام عليكم ورحمة الله وبركاته
	какой раздел тебе нужен?`, mainMarkup}
	m["Разное"] = [2]any{`السلام عليكم ورحمة الله وبركاته
	какой раздел тебе нужен?`, mainMarkup}
	m["Сборники хадисов"] = [2]any{`السلام عليكم ورحمة الله وبركاته
	какой раздел тебе нужен?`, mainMarkup}
	m["Сира"] = [2]any{`السلام عليكم ورحمة الله وبركاته
	какой раздел тебе нужен?`, mainMarkup}
	m["Тафсир"] = [2]any{`السلام عليكم ورحمة الله وبركاته
	какой раздел тебе нужен?`, mainMarkup}
	m["Требование знаний"] = [2]any{`السلام عليكم ورحمة الله وبركاته
	какой раздел тебе нужен?`, mainMarkup}
	m["Фикх"] = [2]any{`السلام عليكم ورحمة الله وبركاته
	какой раздел тебе нужен?`, mainMarkup}

	m["Три основы"] = [2]any{"Акыда\n", akidaMarkup}
	m["Четыре правила"] = [2]any{"Акыда\n", akidaMarkup}
	m["Шесть основ"] = [2]any{"Акыда\n", akidaMarkup}
	m["Отведение сомнений"] = [2]any{"Акыда\n", akidaMarkup}
	m["Книга единобожия"] = [2]any{"Акыда\n", akidaMarkup}
	m["Навакыдуль ислям"] = [2]any{"Акыда\n", akidaMarkup}
	m["Акыда в именах и атрибутах"] = [2]any{"Акыда\n", akidaMarkup}
	m["Вопрос имана"] = [2]any{"Акыда\n", akidaMarkup}
	m["Любовь и непричастность"] = [2]any{"Акыда\n", akidaMarkup}
	m["Разбор вопроса узр биль джахль"] = [2]any{"Акыда\n", akidaMarkup}
	m["Книги по основам акыды и манхаджа"] = [2]any{"Акыда\n", akidaMarkup}
	m["Полезные книги по акыде"] = [2]any{"Акыда\n", akidaMarkup}

	mainMarkup.Inline(
		mainMarkup.Row(mainAdab, mainAkida),
		mainMarkup.Row(mainDifferent, mainHadithCollection),
		mainMarkup.Row(mainSira, mainTafsir),
		mainMarkup.Row(mainKnowledgeNeeding),
		mainMarkup.Row(mainFikkh),
	)

	akidaMarkup.Reply(
		akidaMarkup.Row(akidaTo, akidaChp),
		akidaMarkup.Row(akidaSho),
		akidaMarkup.Row(akidaOs),
		akidaMarkup.Row(akidaKe),
		akidaMarkup.Row(akidaNi),
		akidaMarkup.Row(akidaAviia),
		akidaMarkup.Row(akidaVi),
		akidaMarkup.Row(akidaLin),
		akidaMarkup.Row(akidaRvubd),
		akidaMarkup.Row(akidaKpoaim),
		akidaMarkup.Row(akidaPkpa),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	fikkhMarkup.Reply(
		fikkhMarkup.Row(fikkhBrak, fikkhJihad),
		fikkhMarkup.Row(fikkhClearing, fikkhWorship),
		fikkhMarkup.Row(fikkhFow, fikkhFinance),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	selector.Reply(
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	b.Handle("/start", func(c tele.Context) error {
		return c.Send(welcomeText, mainMarkup)
	})

	b.Handle(&mainAkida, func(c tele.Context) error {
		return c.Edit("Акыда\n", akidaMarkup)
	})

	b.Handle(&akidaTo, func(c tele.Context) error {
		return c.Edit(`Три основы

Книга "Три основы избранные разъяснения": <a href="https://drive.google.com/uc?export=download&id=1P8UXBJXq69QuyhSZcUP6QfuwYWUO3F8G">[Скачать]</a>

Книга "ТРИ ОСНОВЫ МАТН": <a href="https://drive.google.com/uc?export=download&id=1y2rPgnngzZpx86Z6-LpCAK5G9GcIAOcC">[Скачать]</a>
		
Книга "Три основы шарх ибн Баз": <a href="https://drive.google.com/uc?export=download&id=1bVtxCwpXd-mDIc3UE91rrW5cQj15QEDP">[Скачать]</a>

Книга "Три основы шарх ибн Касим": <a href="https://drive.google.com/uc?export=download&id=1QtiNTxvDimo3-An07jpo1pBsEAj1vWiU">[Скачать]</a>

Книга "Три основы шарх ибн Усаймин": <a href="https://drive.google.com/uc?export=download&id=1m4dGn5eQh7O70QBIpsNQfzuDa7yIWGAH">[Скачать]</a>
		
Книга "Три основы шарх Усойми": <a href="https://drive.google.com/uc?export=download&id=1Yb0PWKHznJiC20hBu_YFh9Huy6vUgm9j">[Скачать]</a>
		
Книга "Три основы шарх Фаузан": <a href="https://drive.google.com/uc?export=download&id=1Y4zZMd_XalbNuXY2D1IP1jFPhLJTxqb4">[Скачать]</a>
		`, selector, tele.NoPreview)
	})

	b.Handle(&akidaChp, func(c tele.Context) error {
		return c.Edit(`Четыре правила

Книга "4 правила - аль-Шейх": <a href="https://drive.google.com/uc?export=download&id=110SCaaU3-Xf-Nn2ha2-8N-XLcxhh77xM">[Скачать]</a>
		
Книга "4 правила - Усайми": <a href="https://drive.google.com/uc?export=download&id=1zWG-dmnt9Cr0xpvDhVK8pXql2zV914Dj">[Скачать]</a>
		
Книга "4 правила шарх аль-Люхайдан": <a href="https://drive.google.com/uc?export=download&id=1SNsgrcJo43-vEqwmFqYY37OWA6NDIFWb">[Скачать]</a>
		
Книга "Четыре правила избранные шурухи": <a href="https://drive.google.com/uc?export=download&id=1xldtiDsrcOOtv4iF8H2v5-lOlZKL1CCr">[Скачать]</a>
		
Книга "Четыре правила матн ": <a href="https://drive.google.com/uc?export=download&id=19ZhSD8p-zv9v46wRrM8OpeCWW-UkSSRf">[Скачать]</a>
		
Книга "Четыре правила шарх АбдуРРахман аль-Баррок": <a href="https://drive.google.com/uc?export=download&id=10dH7g_f18i9W50X-0W36WBe9m1QBtuz5">[Скачать]</a>
		
Книга "Четыре правила шарх Абу Джабир": <a href="https://drive.google.com/uc?export=download&id=1-3msbxXqw5GJa0UYwsAcjsbNeqx_hpwn">[Скачать]</a>
		
Книга "Четыре правила шарх ибн Баз": <a href="https://drive.google.com/uc?export=download&id=1GgCw_wRFFlfE05j1nLyd1UpFRvWoh_OI">[Скачать]</a>
		
Книга "Четыре правила шарх Солих али-Шейх": <a href="https://drive.google.com/uc?export=download&id=17nv4SbXLtc2v58ClFjMHcucGu4oJqviz">[Скачать]</a>
		
Книга "Четыре правила шарх Фаузан": <a href="https://drive.google.com/uc?export=download&id=1RTel62opBVHaEo8nWtxWiyhnHmFHyRO3">[Скачать]</a>`, selector, tele.NoPreview)
	})

	b.Handle(&akidaSho, func(c tele.Context) error {
		return c.Edit(`Шесть основ

Книга "Шесть основ матн": <a href="https://drive.google.com/uc?export=download&id=1aTiCPvZwVb0Kxlt_uM_5bDvzQR8R_n_8">[Скачать]</a>
		
Книга "Шесть основ шарх Абу Джабир": <a href="https://drive.google.com/uc?export=download&id=1l3VTb2SsH-Itv4rc1PBo5U93b6m6P1zV">[Скачать]</a>
		
Книга "Шесть основ шарх АдбуРразак аль-Бадр": <a href="https://drive.google.com/uc?export=download&id=1VgMa6rAMFxo_oQzGGlC4Phzqkxl7yrT1">[Скачать]</a>
		
Книга "Шесть основ шарх Фаузан": <a href="https://drive.google.com/uc?export=download&id=1bo3jYVpIMHqEjsj45wiggdL2SJOBlwFJ">[Скачать]</a>`, selector, tele.NoPreview)
	})

	b.Handle(&akidaOs, func(c tele.Context) error {
		return c.Edit(`Отведение сомнений

Книга "Отведение разъяснения ученых": <a href="https://drive.google.com/uc?export=download&id=1Gwl0YEsXaR5YpyaySmO2MYmO5nlkOODN">[Скачать]</a>
		
Книга "Отведение сомнений матн": <a href="https://drive.google.com/uc?export=download&id=1gud6a4mHdM0kppkxmhkAEsf5RO8ObLo2">[Скачать]</a>
		
Книга "Отведение сомнений шарх Гунейман": <a href="https://drive.google.com/uc?export=download&id=1Aqr-G6hN9IQdlcAe3Gey14soytlCvnEB">[Скачать]</a>
		
Книга "Отведение сомнений шарх Солих али-Шейх": <a href="https://drive.google.com/uc?export=download&id=16mwBiVpXmzT36yfKc-JwlkKMn6HUNI0z">[Скачать]</a>
		`, selector, tele.NoPreview)
	})

	b.Handle(&akidaKe, func(c tele.Context) error {
		return c.Edit(`Книга единобожия

Книга "Книга единобожия аль-Бадр": <a href="https://drive.google.com/uc?export=download&id=1UsIOKDE3BjBqvMo64OXuAFz2ZkhFHhFM">[Скачать]</a>
		
Книга "Книга единобожия ас-Саади": <a href="https://drive.google.com/uc?export=download&id=1SUhyIxyj8VMQT3SbQJa69hSrlYZjUWqY">[Скачать]</a>
		
Книга "Книга единобожия короткий шарх Солиха али-Шейх": <a href="https://drive.google.com/uc?export=download&id=1PR9-ww2lW5-6QNqqdgeIWF6ZE-wKuHRv">[Скачать]</a>
		
Книга "Книга единобожия короткий шарх Фаузана ": <a href="https://drive.google.com/uc?export=download&id=1RBtadvWTIMlPMSy4IXFgGMu59T4hZKE1">[Скачать]</a>
		
Книга "Книга единобожия матн": <a href="https://drive.google.com/uc?export=download&id=19A_v12_y7hOJ6x777zsKrzXtfHjefD5V">[Скачать]</a>
		
Книга "Книга единобожия шарх АбдуРрахман али-Шейх": <a href="https://drive.google.com/uc?export=download&id=19yis7n-T9HYPiUfKpUjWJLRya8cCDtGx">[Скачать]</a>
		
Книга "Книга единобожия шарх Абу Джабир": <a href="https://drive.google.com/uc?export=download&id=1vbBwzSHMq9r80BIk96R_akrBwA25f1u0">[Скачать]</a>
		
Книга "Книга единобожия шарх ибн Атик": <a href="https://drive.google.com/uc?export=download&id=1uxKNBzoUB0sQR-5qvf8-wjlTd9bQZGD3">[Скачать]</a>`, selector, tele.NoPreview)
	})

	b.Handle(&akidaNi, func(c tele.Context) error {
		return c.Edit(`Навакыдуль ислям

Книга "10 пунктов аннулирующих ислам — Шейх Фаузан": <a href="https://drive.google.com/uc?export=download&id=1v7hCrnTCwnwJXbwOaCrGU3aUDiznIynC">[Скачать]</a>
		
Книга "10 пунктов аннулирующих ислам шарх Роджихи": <a href="https://drive.google.com/uc?export=download&id=1EqIkQB9D9E9tYZ0YO0tA1B_24f3wDYbs">[Скачать]</a>
		
Книга "10 пунктов аннулирующих исламшарх Баррак": <a href="https://drive.google.com/uc?export=download&id=1_0yoE67bSfU0H-BEAN_5OYaPoD926RVY">[Скачать]</a>
		
Книга "Науакыд - Фаузан": <a href="https://drive.google.com/uc?export=download&id=1uYmLakg5qpBzyWIpJTn8Qf61K4njA8Zo">[Скачать]</a>`, selector, tele.NoPreview)
	})

	b.Handle(&akidaAviia, func(c tele.Context) error {
		return c.Edit(`Акыда в именах и атрибутах

Книга "АКЪИДА ВАСАТИЯ САЛИХ ФАУЗАН": <a href="https://drive.google.com/uc?export=download&id=1-NrJIZUCD3P5iVFyFlDKJbkRPvQHvf3I">[Скачать]</a>
		
Книга "Акыда ат-Тахавия шарх ибн аби аль-Изз": <a href="https://drive.google.com/uc?export=download&id=1B3p_qTx_5Y4dKx5PggeyGXW2cgpxQFeV">[Скачать]</a>
		
Книга "Блеск_убеждений_Шарх_шейха_аль_Усеймин": <a href="https://drive.google.com/uc?export=download&id=17My8d125p8s5i7Iq5NEjd5rzm6QS0zF3">[Скачать]</a>
		
Книга "Идеология тафуида": <a href="https://drive.google.com/uc?export=download&id=1DJxbXmWe__u3E0ORU0PPtR5cP3tLpmyj">[Скачать]</a>
		
Книга "Коранический метод познания атрибутов Аллаха Аш Шанкыти": <a href="https://drive.google.com/uc?export=download&id=13ibRXNTFUsiXDH7Ll0nM7ajzsubBGiR7">[Скачать]</a>
		
Книга "Прекрасные имена Кахтани": <a href="https://drive.google.com/uc?export=download&id=144yN70FjkV64AvEw36B1ppyf6fQkOw0c">[Скачать]</a>
		
Книга "Разногласия в словах и опровержение джахмита": <a href="https://drive.google.com/uc?export=download&id=13LPdUtJy4L5DAmLoVMTlAX1hLzB_BykU">[Скачать]</a>
		
Книга "Середийность в вероубеждений": <a href="https://drive.google.com/uc?export=download&id=1_sJU7y6XfhKOiSCcsNSMd9Pxm9LOIrRv">[Скачать]</a>
		
Книга "Таухид аль-асма ва-сыфат в общем": <a href="https://drive.google.com/uc?export=download&id=1BIFJdCYeDogW7CLWjOF9jvHN_fgtI9X_">[Скачать]</a>`, selector, tele.NoPreview)
	})

	b.Handle(&akidaVi, func(c tele.Context) error {
		return c.Edit(`Вопрос имана

Книга "аль-Фурук": <a href="https://drive.google.com/uc?export=download&id=12cfw1TTcxZTXdr2hYKwtH_TReRXqbGoF">[Скачать]</a>
		
Книга "Книга Имана ибн Таймия": <a href="https://drive.google.com/uc?export=download&id=1h6YYsiP7aCUK-ovvtgi-VVND0ggYFOwb">[Скачать]</a>
		
Книга "Усулю иман ат-Тамими шарх Шейх Солих али-Шейха": <a href="https://drive.google.com/uc?export=download&id=1vTwA9Aac7q4EFAaEocSJrgJGzLVqTf-_">[Скачать]</a>
		
Книга "Учёные комитета против мурджиитов": <a href="https://drive.google.com/uc?export=download&id=1x1e6KhuDMmd2GPMzNB8QCIbgwEEBO73w">[Скачать]</a>`, selector, tele.NoPreview)
	})

	b.Handle(&akidaLin, func(c tele.Context) error {
		return c.Edit(`Любовь и непричастность

Книга "Дружба и непричастность в исламе Фаузан": <a href="https://drive.google.com/uc?export=download&id=1bTAmWQLC6XYgvr7jbPw1QxaGNPKvHEPf">[Скачать]</a>
		
Книга "Основы любви и непричастности": <a href="https://drive.google.com/uc?export=download&id=13Fy1QSGlxxVwLGyouDPt1dhuY_S9M0ui">[Скачать]</a>
		
Книга "Принцип Уаля в Исламе": <a href="https://drive.google.com/uc?export=download&id=1OJI7kU9y6D50ts96SfnPXw4alqN6G2Lx">[Скачать]</a>`, selector, tele.NoPreview)
	})

	b.Handle(&akidaRvubd, func(c tele.Context) error {
		return c.Edit(`Разбор вопроса узр биль джахль

Книга "Аяты об отсутствии оправдания": <a href="https://drive.google.com/uc?export=download&id=12Oj33ZLCZmlL8devtEcrlNOdE4T5-xDz">[Скачать]</a>
		
Книга "Разбор узр биль джахль": <a href="https://drive.google.com/uc?export=download&id=1uU_4XY6wC8Og2_8PzrJjRYASAyADZ2Vn">[Скачать]</a>)
		
Книга "Разбор узр биль джахль 2": <a href="https://drive.google.com/uc?export=download&id=1cF725zdRn_onwqpQADmvWRYm4spFS5Jd">[Скачать]</a>
		
Книга "Разбор узр биль джахль 3": <a href="https://drive.google.com/uc?export=download&id=1gVxBKUxfAE0N0j-mZztwTB63B4TL1gBM">[Скачать]</a>
		
Книга "Разбор узр биль джахль 4": <a href="https://drive.google.com/uc?export=download&id=1Jmvw7d87gnFxbbQISgolUvVshqm6LefF">[Скачать]</a>`, selector, tele.NoPreview)
	})

	b.Handle(&akidaKpoaim, func(c tele.Context) error {
		return c.Edit(`Книги по основам акыды и манхаджа

Книга "Важные уроки": <a href="https://drive.google.com/uc?export=download&id=19Pime4yAMyZ_l0MprV7l3eLVsO4XFvF9">[Скачать]</a>
		
Книга "Вероубеждение единобожия, Фаузан": <a href="https://drive.google.com/uc?export=download&id=1chxAY_rsT9x4phms6T5pp4qsapsMXhJg">[Скачать]</a>
		
Книга "Два свидетельства ибн Джибрин": <a href="https://drive.google.com/uc?export=download&id=1GREagt_2fO_tWzcfVUO3eSNSKIugbJrw">[Скачать]</a>
		
Книга "Доказательства Единобожия аль Бадр": <a href="https://drive.google.com/uc?export=download&id=1mNMUaayWGf3j9R8DBz6xb6x-YhRRtmUS">[Скачать]</a>
		
Книга "Основы вероучения Лялякаи": <a href="https://drive.google.com/uc?export=download&id=1gWUyDd5kV3d8mYsPL3cSWCeckBp523Hu">[Скачать]</a>
		
Книга "Разъяснение основых постулатов веры": <a href="https://drive.google.com/uc?export=download&id=1_r5erIjIZJUPuruuX11WdDMKF69BZpqY">[Скачать]</a>
		
Книга "Религиозные новшества": <a href="https://drive.google.com/uc?export=download&id=1n-TJOKrJ8Ja-XTo7RmraFShc9qU2miKB">[Скачать]</a>
		
Книга "Слова единобожия АбдуРразак аль-Бадр": <a href="https://drive.google.com/uc?export=download&id=1cXWXJIpgSW85FTtZFWPbIkoYNJyDE8Ls">[Скачать]</a>
		
Книга "Убеждения приверженцев сунны и единой общины": <a href="https://drive.google.com/uc?export=download&id=1AvvylkZT6_PbYBsfLkQUcLJSbjEsW7J3">[Скачать]</a>
		
Книга "Уроки извлекаемые из Корана Фаузан": <a href="https://drive.google.com/uc?export=download&id=1aqH57cIr2l8Qjtv6pv4-w6faC8mIeszM">[Скачать]</a>
		
Книга "Шарх ас-Сунна Барбахари": <a href="https://drive.google.com/uc?export=download&id=1WFo2FQ8uK3gkCtNPL_sCnJyK4Wbaxgaa">[Скачать]</a>
		
Книга "Шарх ас-сунна (Барбахари) (2)": <a href="https://drive.google.com/uc?export=download&id=15NE42i21dDbezYOc3Gf55icjPejF5L7H">[Скачать]</a>`, selector, tele.NoPreview)
	})

	b.Handle(&akidaPkpa, func(c tele.Context) error {
		return c.Edit(`Полезные книги по акыде

Книга "Пользы из суры «аль-Фатиха»": <a href="https://drive.google.com/uc?export=download&id=1YnraeXlIk0oZfifRTnUei6LY76QKdcj-">[Скачать]</a>
		
Книга "Акыда Суфьяна ас-Саури": <a href="https://drive.google.com/uc?export=download&id=1wgep3eAi37zegxlQHaIDOmWb5qpZuKfl">[Скачать]</a>
		
Книга "Акыда Суфьяна Уйейна ": <a href="https://drive.google.com/uc?export=download&id=1xTThYm7QKuEMMxkuma9S1vXCsTEojXOP">[Скачать]</a>
		
Книга "Китаб ас-Сунна Харб аль-Кирмани": <a href="https://drive.google.com/uc?export=download&id=1mJNAFaEvsiozGzf1ujmPM2FgyY563tQ8">[Скачать]</a>
		
Книга "МАСАИЛЬ ДЖАХИЛИЯ ФАУЗАН": <a href="https://drive.google.com/uc?export=download&id=1kcVNqyr1DsClR7tFHlpY-PBbN_FF_huP">[Скачать]</a>
		
Книга "Муфид аль-Мустафид ат-Тамими": <a href="https://drive.google.com/uc?export=download&id=1AKOFiXP7xXvBFxAVGXKBk_YBFBgjDesC">[Скачать]</a>
		
Книга "Различия": <a href="https://drive.google.com/uc?export=download&id=1Zf6QCbWpy9r2oRN6whF2pVTDQoFFIFKI">[Скачать]</a>
		
Книга "Усуль ас-Сунна Имама Ахмада шарх Абк Джабир": <a href="https://drive.google.com/uc?export=download&id=1XsW0fpGNgivqRymaTEQBU3_GOyeVos9W">[Скачать]</a>
		
Книга "Фетвы по столпам Ислама": <a href="https://drive.google.com/uc?export=download&id=1z6pO70pEmtrKnetTBKnHyWILN_oKUhSF">[Скачать]</a>
		
Книга "Хаия ибн Аби Давуда шарх Фаузан": <a href="https://drive.google.com/uc?export=download&id=1TFSxj9wNKX6-s3kpQPWe9VxrX3QKnp6x">[Скачать]</a>`, selector, tele.NoPreview)
	})

	b.Handle(&mainAdab, func(c tele.Context) error {
		return c.Edit(`Адаб

Книга "Женщина придерживающаяся правильного пути": <a href="https://drive.google.com/uc?export=download&id=1hBhoqwvtGx_B0Uc7CRtcsNNaNV-5UpcN">[Скачать]</a>
		
Книга "Книга благопристойности": <a href="https://drive.google.com/uc?export=download&id=1an5KdU8onbUy-b3OFqxTKHiNjGMuvnO-">[Скачать]</a>
		
Книга "Мусульманин и его личность в свете Корана и Сунны Карима Cорокоумова": <a href="https://drive.google.com/uc?export=download&id=1OvdmGv2qJM3X_kRvQ5Ilc6oXd6ubnuAW">[Скачать]</a>
		
Книга "Мусульманка и ее личность в свете Корана и Сунны Карима Cорокоумова": <a href="https://drive.google.com/uc?export=download&id=1Lz6ikBPESeYtKtElre0Co7ZgcaZDtdPO">[Скачать]</a>
		
Книга "Проблемы современной молодёжи": <a href="https://drive.google.com/uc?export=download&id=1iM8taeO3pne0omCwEXv6D5Vc5RE4mL8y">[Скачать]</a>
		
Книга "Соблюдение правил этикета залог счастья Салих ибн Абд аль Азиз ибн": <a href="https://drive.google.com/uc?export=download&id=1HR1e7eGA-8iNKmOF6DNFIjLO-omVYhMh">[Скачать]</a>`, selector, tele.NoPreview)
	})

	b.Handle(&mainDifferent, func(c tele.Context) error {
		return c.Edit(`Разное

Книга "Выбор друзей в Исламе": <a href="https://drive.google.com/uc?export=download&id=1wbcMbQ-5kE9v-jtDc6v8LLXAf4n8p9mA">[Скачать]</a>
		
Книга "Способы увеличения имана": <a href="https://drive.google.com/uc?export=download&id=1hs35HAEGAd5qyp273cydIPGSaBl8D8Ea">[Скачать]</a>
		
Книга "20 советов моей сестре до ее выхода замуж": <a href="https://drive.google.com/uc?export=download&id=1_rFIMcAVPrJjuRWHdlYun5HrA3UQYZp5">[Скачать]</a>
		
Книга "Глава о том, что при возникновении смут обязательно искать безопасность": <a href="https://drive.google.com/uc?export=download&id=11ZwDWfYY3Ku0xfDIvlzJFDIrrwApVpCG">[Скачать]</a>
		
Книга "Грубая ошибка — перед цитированием аята читать «А узу Би Лляхи минаш": <a href="https://drive.google.com/uc?export=download&id=1K4H0iQcloWHWK7AjqYkNRDQlSqbbnkg2">[Скачать]</a>
		
Книга "Табукское послание Провизия переселяющегося к своему Господу Ибн": <a href="https://drive.google.com/uc?export=download&id=1YzRg9Q0xk6OBFFjg1trgNfgWhX2dvbUe">[Скачать]</a>
		
Книга "Фаваид Ибн Каййим": <a href="https://drive.google.com/uc?export=download&id=1GuSD-uosy7L0qKX21E19l7Cz5e_XRp5d">[Скачать]</a>
		
Книга "Ibn_Kayyim_-_Vsem_kogo_postiglo_neschastye": <a href="https://drive.google.com/uc?export=download&id=18KnESqYSF5YD3GVlmJo4nhb7S1kcb4eo">[Скачать]</a>
		
Книга "ibn_Taymia_-_Priblizhennye_Allakha_i_priblizhennye_shaytana": <a href="https://drive.google.com/uc?export=download&id=1Z07opjokAsMmrSWLXtfN322NWfbEEQI-">[Скачать]</a>`, selector, tele.NoPreview)
	})

	b.Handle(&mainHadithCollection, func(c tele.Context) error {
		return c.Edit(`Сборники хадисов

Книга "40 хадисов ибн Усаймин": <a href="https://drive.google.com/uc?export=download&id=1XSMys4l1Yhqzyfv8ChZZp1R3qNOERy_x">[Скачать]</a>
		
Книга "Булуг аль-маррам": <a href="https://drive.google.com/uc?export=download&id=1Bxs0nM2myXaulw2wsNG1IjM1MfZlIcfO">[Скачать]</a>
		
Книга "Избранные хадисы Аль-Бухари": <a href="https://drive.google.com/uc?export=download&id=1y_b6MEh9IhCdCw26YQBBVPLxAMMZ3d6-">[Скачать]</a>
		
Книга "Избранные хадисы": <a href="https://drive.google.com/uc?export=download&id=1M-nteft7GAwd_EqlD3jB8QGIJF_5xsQ-">[Скачать]</a>
		
Книга "Пособие по терминологии хадисов": <a href="https://drive.google.com/uc?export=download&id=1G2VNvmes7Luz6LDhKSRofyeWrV0Z8BBC">[Скачать]</a>
		
Книга "Сады праведных": <a href="https://drive.google.com/uc?export=download&id=1YJt0LBs8aKK32Nd2QDR628IzW2yTT1JN">[Скачать]</a>
		
Книга "Сахих аль-Бухари": <a href="https://drive.google.com/uc?export=download&id=1J_i_tEfGuT5Rrao5SWq5S116a0ePWbNA">[Скачать]</a>
		
Книга "Сахих аль-Джами’ Ас-Саг1ир": <a href="https://drive.google.com/uc?export=download&id=1whvVbykHThCM59OfzFRB42cSxOKbCI9B">[Скачать]</a>
		
Книга "Сахих ибн Маджа": <a href="https://drive.google.com/uc?export=download&id=1oRBjWHZlPi6hq-Ogg4bfcd-6ZNPwsSfS">[Скачать]</a>
		
Книга "Сахих Муслим": <a href="https://drive.google.com/uc?export=download&id=1u77kE4p180lgP72FxESottztCdjvF-Ff">[Скачать]</a>
		
Книга "Сунан Абу Давуд": <a href="https://drive.google.com/uc?export=download&id=1XZwNVSI9otpeLEsHSgbQ6RpWjfEn3oez">[Скачать]</a>
		
Книга "Sakhikh_al-Bukhari_Kratkoe_ikhlozhenie": <a href="https://drive.google.com/uc?export=download&id=10E3OSWyiGLvHbbH4Oy7e7PLH-kIAuDKF">[Скачать]</a>
		`, selector, tele.NoPreview)
	})

	b.Handle(&mainSira, func(c tele.Context) error {
		return c.Edit(`Сира

Книга "Достоверная история Али и Муавии": <a href="https://drive.google.com/uc?export=download&id=1v_uktF-7bvB0imnl-P-Z0O-QeZffnOfE">[Скачать]</a>
		
Книга "Жизнеописание Пророка Мухаммада": <a href="https://drive.google.com/uc?export=download&id=1ThOTQazXgciIlA0CmgkWV8QPIcCnYZ0c">[Скачать]</a>
		
Книга "Из жизни сподвижниц": <a href="https://drive.google.com/uc?export=download&id=1ODp7LVfzakbHl_pR--xZMxOTg3paxabc">[Скачать]</a>
		
Книга "КРАТКОЕ_ЖИЗНЕОПИСАНИЕ_ПРОРОКА": <a href="https://drive.google.com/uc?export=download&id=1FRN9TZWfRKO9yBcpbP_pIQoyDZZZvpUR">[Скачать]</a>
		
Книга "Рассказы о пророках": <a href="https://drive.google.com/uc?export=download&id=1x1P-cCTKbPnU7OvknuMJrDP_D5HKQS9g">[Скачать]</a>
		
Книга "Рассказы из жизни сподвижников Мухаммада Аль Баша": <a href="https://drive.google.com/uc?export=download&id=1dw8xlVkK-eSADwprms4ORU4yvhVXz8_y">[Скачать]</a>
		
Книга "Сира Пророка ﷺ Ибн Хишам": <a href="https://drive.google.com/uc?export=download&id=1LfdT0qr845cUU3MctN3U7fqsNaKbEh2C">[Скачать]</a>
		
Книга "Сира Пророка ﷺ Кахтани": <a href="https://drive.google.com/uc?export=download&id=1GDuWd6D2ghi1corgxcMGptEAS58dGFox">[Скачать]</a>
		
Книга "Сира Пророка ﷺ Мубаракфури": <a href="https://drive.google.com/uc?export=download&id=1IvjlWPUGIouL7MNt2XB3vYSCi9YGGJv9">[Скачать]</a>
		
Книга "Тальха ибн Убайдуллах": <a href="https://drive.google.com/uc?export=download&id=1VACRJfiZGVLZjMoSRxi2ouiNPER-IH-B">[Скачать]</a>`, selector, tele.NoPreview)
	})

	// Книга "": <a href="https://drive.google.com/uc?export=download&id=">[Скачать]</a>
	b.Handle(&mainTafsir, func(c tele.Context) error {
		return c.Edit(`Тафсир

Книга "Перевод смыслов аятов Корана Абу Адель": <a href="https://drive.google.com/uc?export=download&id=1PrySZ7zUMVuxv5B9ltw3URNxq_6jpU_u">[Скачать]</a>
		
Книга "Перевод смыслов аятов Корана Эльмир Кулиев": <a href="https://drive.google.com/uc?export=download&id=1J3-BxhkUocgOQTLGeVphBY72FTV8BE20">[Скачать]</a>
		
Книга "Сияющее слово о науке Тафсира": <a href="https://drive.google.com/uc?export=download&id=1jd2yJXXWMQAOs3ONMNZvIXe_ERD4wHzj">[Скачать]</a>
		
Книга "Тафсир Ибн Аббаса (сура аль-Кадр)": <a href="https://drive.google.com/uc?export=download&id=13_rS6_xHrrHluqbr2jpdPG2Cg8Yl0__j">[Скачать]</a>
		
Книга "Тафсир суры аль-Кадр — ибн Касир": <a href="https://drive.google.com/uc?export=download&id=1LKBEdJ3X9dar0XezcRAvtfs0mhGpZmK8">[Скачать]</a>
		
Книга "Тафсир суры «Аль Бакара» аль Усаймин": <a href="https://drive.google.com/uc?export=download&id=1cv4bc-wkGUBCi-WrwPLJkR3UpCwz6M5I">[Скачать]</a>
		
Книга "Тафсир Суры Аль Ахзаб шейх уль Ислам Ибн Таймия": <a href="https://drive.google.com/uc?export=download&id=1hFre9Perca24kZThDKMwqISFeHqTo6ZL">[Скачать]</a>
		
Книга "Фатиха Тафсир Ибн Аббаса": <a href="https://drive.google.com/uc?export=download&id=1lkbb2ZwQLFXp_t_rMK1NuWyBs7BXq2Zh">[Скачать]</a>`, selector, tele.NoPreview)
	})

	b.Handle(&mainKnowledgeNeeding, func(c tele.Context) error {
		return c.Edit(`Требование знаний
		
Книга "Как начать приобретать знания": <a href="https://drive.google.com/uc?export=download&id=1fCW7Cqgf3IHsbU2uFzyhp4mP834aD1Za">[Скачать]</a>
		
Книга "Разъяснение хадиса про требование знаний": <a href="https://drive.google.com/uc?export=download&id=14pPhPNnW8fLFB82_hDipTMknCSqUrshe">[Скачать]</a>
		
Книга "Рекомендации читателю": <a href="https://drive.google.com/uc?export=download&id=1ElwByCuxb07vnnDTbQwWzfPpeJLOgWki">[Скачать]</a>
		
Книга "Украшение искателя знаний": <a href="https://drive.google.com/uc?export=download&id=1SbaovZHMKKLa_OV2gBvz83DtvsbYB9q6">[Скачать]</a>
		
Книга "Уникальное пособие для ищущих знания": <a href="https://drive.google.com/uc?export=download&id=1xzzMrirLooNC1BjgOc2XZoIMyZxhm1Yt">[Скачать]</a>`, selector, tele.NoPreview)
	})

	b.Handle(&mainFikkh, func(c tele.Context) error {
		return c.Edit(`Фикх

Книга "Булуг аль-марам Мухсин аль-Бармауи": <a href="https://drive.google.com/uc?export=download&id=1wXtR75X0YV-iUBppPvCiHHKyQyYhuP7X">[Скачать]</a>

Книга "Все про шутки в исламе": <a href="https://drive.google.com/uc?export=download&id=1FMgaLNkUZpvwitb-Ynvnu-_gSDEyUo5M">[Скачать]</a>

Книга "Жемчужина фикха": <a href="https://drive.google.com/uc?export=download&id=1Wyz9p2JqOQ2NOEHjVWGP2Fou7YJraX4x">[Скачать]</a>

Книга "Законоположения зимы аш Шувей’ир": <a href="https://drive.google.com/uc?export=download&id=13ETG_k4aI9Tw1DTL5fssfN85cSZe9QD4">[Скачать]</a>

Книга "Ихтилят": <a href="https://drive.google.com/uc?export=download&id=1FONm3IpiNSQ6S6aXCHi_-pbUNBFjBHLA">[Скачать]</a>

Книга "Ключ в ханбалитском фикхе Усойми": <a href="https://drive.google.com/uc?export=download&id=1PajCu9mPZdejiiNob1L5Uhg9f12SrP36">[Скачать]</a>

Книга "Лайлат уль Кадр – Ночь_Предопределения!": <a href="https://drive.google.com/uc?export=download&id=1_LLzjv1mq2X0LHAK8WFhddRaa3X_XWxK">[Скачать]</a>

Книга "Махрамы": <a href="https://drive.google.com/uc?export=download&id=1RkNwCDB6hTkYYKYrACmgYJdOwViGn-Nb">[Скачать]</a>

Книга "О праздниках неверующих": <a href="https://drive.google.com/uc?export=download&id=1WSEwdMJZrLJG4fd6LYgL8pfB9I5jjHjZ">[Скачать]</a>

Книга "О положениях, связанных с поздравлениями": <a href="https://drive.google.com/uc?export=download&id=178Ys8fmSHyNNbZRz6Wm3xa1G2dT6lhNk">[Скачать]</a>

Книга "Обряды похорон": <a href="https://drive.google.com/uc?export=download&id=1-xCCogLSlix4SAgP71QjOh960n4gnF9A">[Скачать]</a>

Книга "Отношение к неверующим": <a href="https://drive.google.com/uc?export=download&id=19mPeIdOpFrthV2--pYshLOYzIjfvlHtI">[Скачать]</a>

Книга "Положение тазкии в исламе": <a href="https://drive.google.com/uc?export=download&id=1_JAp6dvmnHZbxVFm0W1XlFsTW4N0GPft">[Скачать]</a>

Книга "Положения, связанные со сновидениями": <a href="https://drive.google.com/uc?export=download&id=1G0M-zmIWxybDTqxtANx6_5enXR0EEupd">[Скачать]</a>

Книга "Рукья посредством Корана": <a href="https://drive.google.com/uc?export=download&id=19UFo0KAKOjM6yAb4p3ESyxQPm3DVWZ8K">[Скачать]</a>

Книга "Следование мазхабу": <a href="https://drive.google.com/uc?export=download&id=1h4Ys3voJnm0KZtyFKBMdGMi3cHPYauY0">[Скачать]</a>

Книга "Суждение относительно отпускания бороды ибн Баз": <a href="https://drive.google.com/uc?export=download&id=1hr9v5ETfOsTDv-MWEukZGqdrrxk4J0JY">[Скачать]</a>

Книга "Умдатуль фикх ибн кудама аль-макдиси": <a href="https://drive.google.com/uc?export=download&id=1OBhUao46UHbhFlHrsmkTmXkWktvpMfbB">[Скачать]</a>

Книга "Харам и халяль в пище": <a href="https://drive.google.com/uc?export=download&id=1wYAhr_yD6cgfu9z__HVCED9WdB71dKDG">[Скачать]</a>`, fikkhMarkup, tele.NoPreview)
	})

	b.Handle(&selectorBackAtAllBtn, func(c tele.Context) error {
		return c.Edit(welcomeText, mainMarkup)
	})

	b.Handle(&selectorBack, func(c tele.Context) error {
		keyToPreviousStep := strings.Split(c.Message().Text, "\n")[0]
		previousText, previousMarkup := m[keyToPreviousStep][0].(string), m[keyToPreviousStep][1].(*tele.ReplyMarkup)
		return c.Edit(previousText, previousMarkup)
	})

	log.Printf("%s started...", b.Me.Username)
	b.Start()
}
