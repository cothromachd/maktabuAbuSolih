package main

import (
	"fmt"
	"github.com/cothromachd/maktabuAbuSolih/repo"
	"log"
	"os"
	"time"

	"github.com/cothromachd/maktabuAbuSolih/migrations"

	tele "gopkg.in/telebot.v3"
)

var (
	welcomeText = `السلام عليكم ورحمة الله وبركات
	какой раздел тебе нужен?`
	m = make(map[string][2]any)
	//mFile = make(map[string]*tele.Document)
	pgUser     = os.Getenv("POSTGRES_USER")
	pgPassword = os.Getenv("POSTGRES_PASSWORD")
	pgDb       = os.Getenv("POSTGRES_DB")
	token      = os.Getenv("TOKEN") // postgres://postgres:@localhost:5432/maktabu_bot
)

func main() {
	pgConn := fmt.Sprintf("postgres://%s:%s@db:5432/%s", pgUser, pgPassword, pgDb)
	log.Println(pgConn)
	errFile, err := os.OpenFile("errLogs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer errFile.Close()

	logger := log.New(errFile, "", log.Lshortfile|log.LstdFlags)

	pref := tele.Settings{
		Token:     token,
		Poller:    &tele.LongPoller{Timeout: 20 * time.Second},
		ParseMode: tele.ModeHTML,
		OnError: func(err error, c tele.Context) {
			log.Printf("%+v\n", err)
			logger.Println(err)
		},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		logger.Println(err)
		log.Fatal(err)
	}

	storage := repo.New(pgConn, logger)

	err = migrations.Migrate(pgConn)
	if err != nil {
		logger.Println(err)
		log.Fatal(err)
	}

	var (
		mainMarkup = &tele.ReplyMarkup{ResizeKeyboard: true}

		mainAdab       = mainMarkup.Text("Адаб")
		adabMarkup     = &tele.ReplyMarkup{ResizeKeyboard: true}
		adabZhppp      = adabMarkup.Text("Женщина придерживающаяся правильного пути")
		adabKb         = adabMarkup.Text("Книга благопристойности")
		adabMNielvsKiS = adabMarkup.Text("Мусульманин и его личность в свете Корана и Сунны Карима Cорокоумова")
		adabMKielvKiS  = adabMarkup.Text("Мусульманка и ее личность в свете Корана и Сунны Карима Cорокоумова")
		adabPsm        = adabMarkup.Text("Проблемы современной молодёжи")
		adabSpezs      = adabMarkup.Text("Соблюдение правил этикета залог счастья Салих ибн Абд аль Азиз ибн")

		mainDifferent         = mainMarkup.Text("Разное")
		differentMarkup       = &tele.ReplyMarkup{ResizeKeyboard: true}
		different20smsdevz    = differentMarkup.Text("20 советов моей сестре до ее выхода замуж")
		differentVdvi         = differentMarkup.Text("Выбор друзей в Исламе")
		differentGotChpvsoib  = differentMarkup.Text("Глава о том, что при возникновении смут обязательно искать безопасность")
		differentGopcaChablmr = differentMarkup.Text("Грубая ошибка — перед цитированием аята читать «А узу Би Лляхи минаш»")
		differentSui          = differentMarkup.Text("Способы увеличения имана")
		differentTpppksgi     = differentMarkup.Text("Табукское послание Провизия переселяющегося к своему Господу Ибн Къаййим")
		differentFik          = differentMarkup.Text("Фаваид Ибн Каййим")
		differentIkvkpn       = differentMarkup.Text("Ибн Къаййим. Всем, кого постигло несчастье")
		differentItpaipSh     = differentMarkup.Text("Ибн Теймия. Приближение Аллаха и приближение шайтана")

		mainHadithCollection    = mainMarkup.Text("Сборники хадисов")
		hadithCollectionMarkup  = &tele.ReplyMarkup{ResizeKeyboard: true}
		hadithCollection40Khiu  = hadithCollectionMarkup.Text("40 хадисов ибн Усаймин")
		hadithCollectionBam     = hadithCollectionMarkup.Text("Булуг аль-маррам")
		hadithCollectionIhab    = hadithCollectionMarkup.Text("Избранные хадисы Аль-Бухари")
		hadithCollectionIh      = hadithCollectionMarkup.Text("Избранные хадисы")
		hadithCollectionPptKh   = hadithCollectionMarkup.Text("Пособие по терминологии хадисов")
		hadithCollectionSp      = hadithCollectionMarkup.Text("Сады праведных")
		hadithCollectionSab     = hadithCollectionMarkup.Text("Сахих аль-Бухари")
		hadithCollectionSaDzhas = hadithCollectionMarkup.Text("Сахих аль-Джами’ Ас-Саг1ир")
		hadithCollectionSim     = hadithCollectionMarkup.Text("Сахих ибн Маджа")
		hadithCollectionSm      = hadithCollectionMarkup.Text("Сахих Муслим")
		hadithCollectionSad     = hadithCollectionMarkup.Text("Сунан Абу Давуд")
		hadithCollectionSabki   = hadithCollectionMarkup.Text("Сахих аль-Бухари краткое изложение")

		mainSira     = mainMarkup.Text("Сира")
		siraMarkup   = &tele.ReplyMarkup{ResizeKeyboard: true}
		siraDiaim    = siraMarkup.Text("Достоверная история Али и Муавии")
		siraZhpm     = siraMarkup.Text("Жизнеописание Пророка Мухаммада")
		siraIZhs     = siraMarkup.Text("Из жизни сподвижниц")
		siraKZhp     = siraMarkup.Text("Краткое жизнеописание Пророка")
		siraRop      = siraMarkup.Text("Рассказы о пророках")
		siraRiZhsmab = siraMarkup.Text("Рассказы из жизни сподвижников Мухаммада Аль Баша")
		siraSpiKh    = siraMarkup.Text("Сира Пророка ﷺ Ибн Хишам")
		siraSpk      = siraMarkup.Text("Сира Пророка ﷺ Кахтани")
		siraSpm      = siraMarkup.Text("Сира Пророка ﷺ Мубаракфури")
		siraTiu      = siraMarkup.Text("Тальха ибн Убайдуллах")

		mainTafsir       = mainMarkup.Text("Тафсир")
		tafsirMarkup     = &tele.ReplyMarkup{ResizeKeyboard: true}
		tafsirPsaKaa     = tafsirMarkup.Text("Перевод смыслов аятов Корана Абу Адель")
		tafsirPsaKek     = tafsirMarkup.Text("Перевод смыслов аятов Корана Эльмир Кулиев")
		tafsirSsonT      = tafsirMarkup.Text("Сияющее слово о науке Тафсира")
		tafsirTiasak     = tafsirMarkup.Text("Тафсир Ибн Аббаса (сура аль-Кадр)")
		tafsirTsakik     = tafsirMarkup.Text("Тафсир суры аль-Кадр — ибн Касир")
		tafsirTsabau     = tafsirMarkup.Text("Тафсир суры «Аль_Бакара» аль Усаймин")
		tafsirTsaaShuiit = tafsirMarkup.Text("Тафсир Суры Аль Ахзаб шейх уль Ислам Ибн Таймия")
		tafsirFtia       = tafsirMarkup.Text("Фатиха Тафсир Ибн Аббаса")

		mainKnowledgeNeeding   = mainMarkup.Text("Требование знаний")
		knowledgeNeedingMarkup = &tele.ReplyMarkup{ResizeKeyboard: true}
		knowledgeNeedingKnpz   = knowledgeNeedingMarkup.Text("Как начать приобретать знания")
		knowledgeNeedingRKhptz = knowledgeNeedingMarkup.Text("Разъяснение хадиса про требование знаний")
		knowledgeNeedingRCh    = knowledgeNeedingMarkup.Text("Рекомендации читателю")
		knowledgeNeedingUiz    = knowledgeNeedingMarkup.Text("Украшение искателя знаний")
		knowledgeNeedingUpdiz  = knowledgeNeedingMarkup.Text("Уникальное пособие для ищущих знания")

		mainFikkh        = mainMarkup.Text("Фикх")
		fikkhMarkup      = &tele.ReplyMarkup{ResizeKeyboard: true}
		fikkhBrak        = fikkhMarkup.Text("Брак")
		brakMarkup       = &tele.ReplyMarkup{ResizeKeyboard: true}
		brakVt           = brakMarkup.Text("Важность потомства")
		brakPm           = brakMarkup.Text("Подчинение мужу")
		brakSvadbasi     = brakMarkup.Text("Свадьба согласно Исламу")
		brakSvatovstvosi = brakMarkup.Text("Сватовство согласно исламу")
		brakSo           = brakMarkup.Text("Супружеские ошибки")
		brakEba          = brakMarkup.Text("Этикет бракосочетания Альбани")

		fikkhJihad  = fikkhMarkup.Text("Джихад")
		jihadMarkup = &tele.ReplyMarkup{ResizeKeyboard: true}
		jihadVvab   = jihadMarkup.Text("Вопросы войны аль-Бадр")

		fikkhClearing  = fikkhMarkup.Text("Очищение")
		clearingMarkup = &tele.ReplyMarkup{ResizeKeyboard: true}
		clearingMabko  = clearingMarkup.Text("Мухтасар аль Бувейты книга омовения")

		fikkhWorship = fikkhMarkup.Text("Поклонение")

		worshipMarkup = &tele.ReplyMarkup{ResizeKeyboard: true}
		worshipZakyat = worshipMarkup.Text("Закят")
		zakyatMarkup  = &tele.ReplyMarkup{ResizeKeyboard: true}
		zakyatZaf     = zakyatMarkup.Text("Закят аль-Фитр")
		zakyatZ       = zakyatMarkup.Text("Закят Умар Абуль Хасан")

		worshipMolitva = worshipMarkup.Text("Молитва")
		molitvaMarkup  = &tele.ReplyMarkup{ResizeKeyboard: true}
		molitvaDnm     = molitvaMarkup.Text("Добровольные ночные молитвы")
		molitvaDm      = molitvaMarkup.Text("Достоинство молитвы")
		molitvaI       = molitvaMarkup.Text("Истихара")
		molitvaKdpm    = molitvaMarkup.Text("Когда дозволено прерывать молитву")
		molitvaPn      = molitvaMarkup.Text("Праздничный намаз")
		molitvaPnaia   = molitvaMarkup.Text("Праздничный намаз — адабы и ахкамы")
		molitvaT       = molitvaMarkup.Text("Таравих")
		molitvaUmSha   = molitvaMarkup.Text("Условия молитвы шарх Аббад")

		worshipMolba = worshipMarkup.Text("Мольба")
		molbaMarkup  = &tele.ReplyMarkup{ResizeKeyboard: true}
		molbaPm      = molbaMarkup.Text("Положение мольбы")

		worshipPokayanie      = worshipMarkup.Text("Покаяние")
		pokayanieMarkup       = &tele.ReplyMarkup{ResizeKeyboard: true}
		pokayanieKpzKhzetokzn = pokayanieMarkup.Text("Как покаяться за хулу злословие, если тот, о ком злословили, ничего")
		pokayaniePivssn       = pokayanieMarkup.Text("Покаяние и все связанное с ним")
		pokayanieYKhp         = pokayanieMarkup.Text("Я хочу покаяться")

		worshipPost    = worshipMarkup.Text("Пост")
		postMarkup     = &tele.ReplyMarkup{ResizeKeyboard: true}
		postPv6dSh     = postMarkup.Text("Пост в 6 дней шавваля")
		postOpikuaa    = postMarkup.Text("О посте из книги Умдат аль-Ахкам")
		postKpppaaiaar = postMarkup.Text("Краткое пояснение положений поста Абдуль Азиз ибн Абдуллах Ар Раджихи")
		postDpinmvraib = postMarkup.Text("Достоинство поста и ночных молитв в Рамадан Абдульазиз ибн Баз")

		worshipKhidzhra = worshipMarkup.Text("Хиджра")
		khidzhraMarkup  = &tele.ReplyMarkup{ResizeKeyboard: true}
		khidzhraPKhdi   = khidzhraMarkup.Text("Положение хиджры в исламе")

		fikkhFzh      = fikkhMarkup.Text("Фикх женщин")
		fzhMarkup     = &tele.ReplyMarkup{ResizeKeyboard: true}
		fzhOkZhv      = fzhMarkup.Text("О каждодневных женских выделениях")
		fzhTKhoma     = fzhMarkup.Text("Твой хиджаб, о мусульманка! 'Абдур")
		fzhKhniipvKhf = fzhMarkup.Text("Хайд, нифас и истихада положение в ханбалитском фикхе")
		fzhKhuau      = fzhMarkup.Text("Хиджаб (‘Умар аль-‘Умар)")
		fzhEpkvZhaf   = fzhMarkup.Text("EQAMO Постановления касающиеся верующих женщин аль Фаузан")
		fzhTKhomArab  = fzhMarkup.Text("EQAMO Твой хиджаб, о мусульманка! Абдур Разак аль Бадр")

		fikkhFinance  = fikkhMarkup.Text("Финансы")
		financeMarkup = &tele.ReplyMarkup{ResizeKeyboard: true}
		financeFfo1t  = financeMarkup.Text("Фикх финансовых отношений - 1 том")
		financeFfo2t  = fikkhMarkup.Text("Фикх финансовых отношений — 2 том")

		fikkhBammab = fikkhMarkup.Text("Булуг аль-марам Мухсин аль-Бармауи")
		fikkhVpShvi = fikkhMarkup.Text("Все про шутки в исламе")
		fikkhZhf    = fikkhMarkup.Text("Жемчужина фикха")
		fikkhZzaSh  = fikkhMarkup.Text("Законоположения зимы аш Шувей’ир")
		fikkhI      = fikkhMarkup.Text("Ихтилят")
		fikkhKvhfu  = fikkhMarkup.Text("Ключ в ханбалитском фикхе Усойми")
		fikkhluknp  = fikkhMarkup.Text("Лайлат уль Кадр – Ночь Предопределения!!!")
		fikkhM      = fikkhMarkup.Text("Махрамы")
		fikkhOpn    = fikkhMarkup.Text("О праздниках неверующих")
		fikkhOpssp  = fikkhMarkup.Text("О положениях, связанных с поздравлениями")
		fikkhOp     = fikkhMarkup.Text("Обряды похорон")
		fikkhOkn    = fikkhMarkup.Text("Отношение к неверующим")
		fikkhPtvi   = fikkhMarkup.Text("Положение тазкии в исламе")
		fikkhPsss   = fikkhMarkup.Text("Положения, связанные со сновидениями")
		fikkhRpk    = fikkhMarkup.Text("Рукья посредством Корана")
		fikkhSm     = fikkhMarkup.Text("Следование мазхабу")
		fikkhSoobib = fikkhMarkup.Text("Суждение относительно отпускания бороды ибн Баз")
		fikkhUfikam = fikkhMarkup.Text("Умдатуль фикх ибн кудама аль-макдиси")
		fikkhHihvp  = fikkhMarkup.Text("Харам и халяль в пище")

		mainAkida   = mainMarkup.Text("Акыда")
		akidaMarkup = &tele.ReplyMarkup{ResizeKeyboard: true}

		akidaTo  = akidaMarkup.Text("Три основы")
		ToMarkup = &tele.ReplyMarkup{ResizeKeyboard: true}
		ToToir   = ToMarkup.Text("Три основы избранные разъяснения")
		ToTom    = ToMarkup.Text("ТРИ ОСНОВЫ МАТН")
		ToShib   = ToMarkup.Text("Три основы шарх ибн Баз")
		ToShik   = ToMarkup.Text("Три основы шарх ибн Касим")
		ToShiu   = ToMarkup.Text("Три основы шарх ибн Усаймин")
		ToShu    = ToMarkup.Text("Три основы шарх Усойми")
		ToShf    = ToMarkup.Text("Три основы шарх Фаузан")

		akidaChp  = akidaMarkup.Text("Четыре правила")
		ChpMarkup = &tele.ReplyMarkup{ResizeKeyboard: true}
		Chp4paSh  = ChpMarkup.Text("4 правила - аль-Шейх")
		Chp4pu    = ChpMarkup.Text("4 правила - Усайми")
		ChpShal   = ChpMarkup.Text("4 правила шарх аль-Люхайдан")
		ChpISh    = ChpMarkup.Text("Четыре правила избранные шурухи")
		ChpM      = ChpMarkup.Text("Четыре правила матн")
		ChpSharab = ChpMarkup.Text("Четыре правила шарх АбдуРРахман аль-Баррок")
		ChpShaDzh = ChpMarkup.Text("Четыре правила шарх Абу Джабир")
		ChpShib   = ChpMarkup.Text("Четыре правила шарх ибн Баз")
		ChpShsaSh = ChpMarkup.Text("Четыре правила шарх Солих али-Шейх")
		ChpShf    = ChpMarkup.Text("Четыре правила шарх Фаузан")

		akidaSho   = akidaMarkup.Text("Шесть основ")
		ShoMarkup  = &tele.ReplyMarkup{ResizeKeyboard: true}
		ShoShom    = ShoMarkup.Text("Шесть основ матн")
		ShoShoShad = ShoMarkup.Text("Шесть основ шарх Абу Джабир")
		ShoSharab  = ShoMarkup.Text("Шесть основ шарх АдбуРразак аль-Бадр")
		ShoShoshf  = ShoMarkup.Text("Шесть основ шарх Фаузан")

		akidaOs    = akidaMarkup.Text("Отведение сомнений")
		OsMarkup   = &tele.ReplyMarkup{ResizeKeyboard: true}
		OsOru      = OsMarkup.Text("Отведение разъяснения ученых")
		OsOsm      = OsMarkup.Text("Отведение сомнений матн")
		OsOsShg    = OsMarkup.Text("Отведение сомнений шарх Гунейман")
		OsOsShsaSh = OsMarkup.Text("Отведение сомнений шарх Солих али-Шейх")

		akidaKe     = akidaMarkup.Text("Книга единобожия")
		KeMarkup    = &tele.ReplyMarkup{ResizeKeyboard: true}
		KeKeab      = KeMarkup.Text("Книга единобожия аль-Бадр")
		KeKeas      = KeMarkup.Text("Книга единобожия ас-Саади")
		KeKekShsaSh = KeMarkup.Text("Книга единобожия короткий шарх Солиха али-Шейх")
		KeKekShf    = KeMarkup.Text("Книга единобожия короткий шарх Фаузана")
		KeKem       = KeMarkup.Text("Книга единобожия матн")
		KeKeShAraSh = KeMarkup.Text("Книга единобожия шарх АбдуРрахман али-Шейх")
		KeKeshad    = KeMarkup.Text("Книга единобожия шарх Абу Джабир")
		KeShia      = KeMarkup.Text("Книга единобожия шарх ибн Атик")

		akidaNi    = akidaMarkup.Text("Навакыдуль ислям")
		NiMarkup   = &tele.ReplyMarkup{ResizeKeyboard: true}
		Ni10paiShb = NiMarkup.Text("10 пунктов аннулирующих ислам шарх Баррак")
		Ni10paiShr = NiMarkup.Text("10 пунктов аннулирующих ислам шарх Роджихи")
		Ni10paiShf = NiMarkup.Text("10 пунктов аннулирующих ислам — Шейх Фаузан")
		NiNf       = NiMarkup.Text("Науакыд - Фаузан")

		akidaAviia     = akidaMarkup.Text("Акыда в именах и атрибутах")
		AviiaMarkup    = &tele.ReplyMarkup{ResizeKeyboard: true}
		AviiaAvsf      = AviiaMarkup.Text("АКЪИДА ВАСАТИЯ САЛИХ ФАУЗАН")
		AviiaAatShiaai = AviiaMarkup.Text("Акыда ат-Тахавия шарх ибн аби аль-Изз")
		AviiaBuShShau  = AviiaMarkup.Text("Блеск убеждений Шарх шейха аль Усеймин")
		AviiaIt        = AviiaMarkup.Text("Идеология тафуида")
		AviiaKmpaAaSh  = AviiaMarkup.Text("Коранический метод познания атрибутов Аллаха Аш Шанкыти")
		AviiaPik       = AviiaMarkup.Text("Прекрасные имена Кахтани")
		AviiaRvsiod    = AviiaMarkup.Text("Разногласия в словах и опровержение джахмита")
		AviiaSvv       = AviiaMarkup.Text("Середийность в вероубеждений")
		AviiaTaavsvo   = AviiaMarkup.Text("Таухид аль-асма ва-сыфат в общем")

		akidaVi        = akidaMarkup.Text("Вопрос имана")
		ViMarkup       = &tele.ReplyMarkup{ResizeKeyboard: true}
		ViAf           = ViMarkup.Text("аль-Фурук")
		ViKiit         = ViMarkup.Text("Книга Имана ибн Таймия")
		ViUiatShShsaSh = ViMarkup.Text("Усулю иман ат-Тамими шарх Шейх Солих али-Шейха")
		ViUkpm         = ViMarkup.Text("Учёные комитета против мурджиитов")

		akidaLin  = akidaMarkup.Text("Любовь и непричастность")
		LinMarkup = &tele.ReplyMarkup{ResizeKeyboard: true}
		LinDinvif = LinMarkup.Text("Дружба и непричастность в исламе Фаузан")
		LinOlin   = LinMarkup.Text("Основы любви и непричастности")
		LinPuvi   = LinMarkup.Text("Принцип Уаля в Исламе")

		akidaRvubd  = akidaMarkup.Text("Разбор вопроса узр биль джахль")
		RvubdMarkup = &tele.ReplyMarkup{ResizeKeyboard: true}
		RvubdAooo   = RvubdMarkup.Text("Аяты об отсутствии оправдания")
		RvubdRubd   = RvubdMarkup.Text("Разбор узр биль джахль")
		RvubdRubd2  = RvubdMarkup.Text("Разбор узр биль джахль 2")
		RvubdRubd3  = RvubdMarkup.Text("Разбор узр биль джахль 3")
		RvubdRubd4  = RvubdMarkup.Text("Разбор узр биль джахль 4")

		akidaKpoaim  = akidaMarkup.Text("Книги по основам акыды и манхаджа")
		KpoaimMarkup = &tele.ReplyMarkup{ResizeKeyboard: true}
		KpoaimVu     = KpoaimMarkup.Text("Важные уроки")
		KpoaimVef    = KpoaimMarkup.Text("Вероубеждение единобожия, Фаузан")
		KpoaimDsid   = KpoaimMarkup.Text("Два свидетельства ибн Джибрин")
		KpoaimDeab   = KpoaimMarkup.Text("Доказательства Единобожия аль Бадр")
		KpoaimOvl    = KpoaimMarkup.Text("Основы вероучения Лялякаи")
		KpoaimRopv   = KpoaimMarkup.Text("Разъяснение основых постулатов веры")
		KpoaimRn     = KpoaimMarkup.Text("Религиозные новшества")
		KpoaimSeAab  = KpoaimMarkup.Text("Слова единобожия АбдуРразак аль-Бадр")
		KpoaimUpsieo = KpoaimMarkup.Text("Убеждения приверженцев сунны и единой общины")
		KpoaimUiikf  = KpoaimMarkup.Text("Уроки извлекаемые из Корана Фаузан")
		KpoaimShasb  = KpoaimMarkup.Text("Шарх ас-Сунна Барбахари")
		KpoaimShasb2 = KpoaimMarkup.Text("Шарх ас-сунна (Барбахари) (2)")

		akidaPkpa       = akidaMarkup.Text("Полезные книги по акыде")
		PkpaMarkup      = &tele.ReplyMarkup{ResizeKeyboard: true}
		PkpaPisaf       = PkpaMarkup.Text("«Пользы из суры «аль-Фатиха»»")
		PkpaAsas        = PkpaMarkup.Text("Акыда Суфьяна ас-Саури")
		PkpaAsu         = PkpaMarkup.Text("Акыда Суфьяна Уйейна")
		PkpaKasKhak     = PkpaMarkup.Text("Китаб ас-Сунна Харб аль-Кирмани")
		PkpaMDzhf       = PkpaMarkup.Text("МАСАИЛЬ ДЖАХИЛИЯ ФАУЗАН")
		PkpaMamat       = PkpaMarkup.Text("Муфид аль-Мустафид ат-Тамими")
		PkpaR           = PkpaMarkup.Text("Различия")
		PkpaUasiaShaDzh = PkpaMarkup.Text("Усуль ас-Сунна Имама Ахмада шарх Абу Джабир")
		PkpaFpsi        = PkpaMarkup.Text("Фетвы по столпам Ислама")
		PkpaKhiadShf    = PkpaMarkup.Text("Хаия ибн Аби Давуда шарх Фаузан")

		selector             = &tele.ReplyMarkup{ResizeKeyboard: true}
		selectorBackAtAllBtn = selector.Text("В начало")
		selectorBack         = selector.Text("⬅️")
	)

	m["Адаб"] = [2]any{"Main", mainMarkup}
	m["Акыда"] = [2]any{"Main", mainMarkup}
	m["Разное"] = [2]any{"Main", mainMarkup}
	m["Сборники хадисов"] = [2]any{"Main", mainMarkup}
	m["Сира"] = [2]any{"Main", mainMarkup}
	m["Тафсир"] = [2]any{"Main", mainMarkup}
	m["Требование знаний"] = [2]any{"Main", mainMarkup}
	m["Фикх"] = [2]any{"Main", mainMarkup}

	m["Три основы"] = [2]any{"Акыда", akidaMarkup}
	m["Четыре правила"] = [2]any{"Акыда", akidaMarkup}
	m["Шесть основ"] = [2]any{"Акыда", akidaMarkup}
	m["Отведение сомнений"] = [2]any{"Акыда", akidaMarkup}
	m["Книга единобожия"] = [2]any{"Акыда", akidaMarkup}
	m["Навакыдуль ислям"] = [2]any{"Акыда", akidaMarkup}
	m["Акыда в именах и атрибутах"] = [2]any{"Акыда", akidaMarkup}
	m["Вопрос имана"] = [2]any{"Акыда", akidaMarkup}
	m["Любовь и непричастность"] = [2]any{"Акыда", akidaMarkup}
	m["Разбор вопроса узр биль джахль"] = [2]any{"Акыда", akidaMarkup}
	m["Книги по основам акыды и манхаджа"] = [2]any{"Акыда", akidaMarkup}
	m["Полезные книги по акыде"] = [2]any{"Акыда", akidaMarkup}

	m["Брак"] = [2]any{"Фикх", fikkhMarkup}
	m["Джихад"] = [2]any{"Фикх", fikkhMarkup}
	m["Очищение"] = [2]any{"Фикх", fikkhMarkup}
	m["Поклонение"] = [2]any{"Фикх", fikkhMarkup}

	m["Закят"] = [2]any{"Поклонение", fikkhMarkup}
	m["Молитва"] = [2]any{"Поклонение", fikkhMarkup}
	m["Мольба"] = [2]any{"Поклонение", fikkhMarkup}
	m["Покаяние"] = [2]any{"Поклонение", fikkhMarkup}
	m["Пост"] = [2]any{"Поклонение", fikkhMarkup}
	m["Хиджра"] = [2]any{"Поклонение", fikkhMarkup}

	m["Фикх женщин"] = [2]any{"Фикх", fikkhMarkup}
	m["Финансы"] = [2]any{"Фикх", fikkhMarkup}

	mainMarkup.Reply(
		mainMarkup.Row(mainAdab, mainAkida),
		mainMarkup.Row(mainDifferent, mainHadithCollection),
		mainMarkup.Row(mainSira, mainTafsir),
		mainMarkup.Row(mainKnowledgeNeeding),
		mainMarkup.Row(mainFikkh),
	)

	adabMarkup.Reply(
		adabMarkup.Row(adabZhppp),
		adabMarkup.Row(adabKb),
		adabMarkup.Row(adabMNielvsKiS),
		adabMarkup.Row(adabMKielvKiS),
		adabMarkup.Row(adabPsm),
		adabMarkup.Row(adabSpezs),
		selector.Row(selectorBack, selectorBackAtAllBtn),
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

	ToMarkup.Reply(
		ToMarkup.Row(ToToir),
		ToMarkup.Row(ToTom),
		ToMarkup.Row(ToShib),
		ToMarkup.Row(ToShik),
		ToMarkup.Row(ToShiu),
		ToMarkup.Row(ToShu),
		ToMarkup.Row(ToShf),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	ChpMarkup.Reply(
		ChpMarkup.Row(Chp4paSh),
		ChpMarkup.Row(Chp4pu),
		ChpMarkup.Row(ChpShal),
		ChpMarkup.Row(ChpISh),
		ChpMarkup.Row(ChpM),
		ChpMarkup.Row(ChpSharab),
		ChpMarkup.Row(ChpShaDzh),
		ChpMarkup.Row(ChpShib),
		ChpMarkup.Row(ChpShsaSh),
		ChpMarkup.Row(ChpShf),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	ShoMarkup.Reply(
		ShoMarkup.Row(ShoShom),
		ShoMarkup.Row(ShoShoShad),
		ShoMarkup.Row(ShoSharab),
		ShoMarkup.Row(ShoShoshf),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	KeMarkup.Reply(
		KeMarkup.Row(KeKeab),
		KeMarkup.Row(KeKeas),
		KeMarkup.Row(KeKekShsaSh),
		KeMarkup.Row(KeKekShf),
		KeMarkup.Row(KeKem),
		KeMarkup.Row(KeKeShAraSh),
		KeMarkup.Row(KeKeshad),
		KeMarkup.Row(KeShia),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	OsMarkup.Reply(
		OsMarkup.Row(OsOru),
		OsMarkup.Row(OsOsm),
		OsMarkup.Row(OsOsShg),
		OsMarkup.Row(OsOsShsaSh),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	NiMarkup.Reply(
		NiMarkup.Row(Ni10paiShb),
		NiMarkup.Row(Ni10paiShr),
		NiMarkup.Row(Ni10paiShf),
		NiMarkup.Row(NiNf),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	AviiaMarkup.Reply(
		AviiaMarkup.Row(AviiaAvsf),
		AviiaMarkup.Row(AviiaAatShiaai),
		AviiaMarkup.Row(AviiaBuShShau),
		AviiaMarkup.Row(AviiaIt),
		AviiaMarkup.Row(AviiaKmpaAaSh),
		AviiaMarkup.Row(AviiaPik),
		AviiaMarkup.Row(AviiaRvsiod),
		AviiaMarkup.Row(AviiaSvv),
		AviiaMarkup.Row(AviiaTaavsvo),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	ViMarkup.Reply(
		ViMarkup.Row(ViAf),
		ViMarkup.Row(ViKiit),
		ViMarkup.Row(ViUiatShShsaSh),
		ViMarkup.Row(ViUkpm),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	LinMarkup.Reply(
		LinMarkup.Row(LinDinvif),
		LinMarkup.Row(LinOlin),
		LinMarkup.Row(LinPuvi),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	RvubdMarkup.Reply(
		RvubdMarkup.Row(RvubdAooo),
		RvubdMarkup.Row(RvubdRubd),
		RvubdMarkup.Row(RvubdRubd2),
		RvubdMarkup.Row(RvubdRubd3),
		RvubdMarkup.Row(RvubdRubd4),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	KpoaimMarkup.Reply(
		KpoaimMarkup.Row(KpoaimVu),
		KpoaimMarkup.Row(KpoaimVef),
		KpoaimMarkup.Row(KpoaimDsid),
		KpoaimMarkup.Row(KpoaimDeab),
		KpoaimMarkup.Row(KpoaimOvl),
		KpoaimMarkup.Row(KpoaimRopv),
		KpoaimMarkup.Row(KpoaimRn),
		KpoaimMarkup.Row(KpoaimSeAab),
		KpoaimMarkup.Row(KpoaimUpsieo),
		KpoaimMarkup.Row(KpoaimUiikf),
		KpoaimMarkup.Row(KpoaimShasb),
		KpoaimMarkup.Row(KpoaimShasb2),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	PkpaMarkup.Reply(
		PkpaMarkup.Row(PkpaPisaf),
		PkpaMarkup.Row(PkpaAsas),
		PkpaMarkup.Row(PkpaAsu),
		PkpaMarkup.Row(PkpaKasKhak),
		PkpaMarkup.Row(PkpaMDzhf),
		PkpaMarkup.Row(PkpaMamat),
		PkpaMarkup.Row(PkpaR),
		PkpaMarkup.Row(PkpaUasiaShaDzh),
		PkpaMarkup.Row(PkpaFpsi),
		PkpaMarkup.Row(PkpaKhiadShf),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	fikkhMarkup.Reply(
		fikkhMarkup.Row(fikkhBrak, fikkhJihad),
		fikkhMarkup.Row(fikkhClearing, fikkhWorship),
		fikkhMarkup.Row(fikkhFzh, fikkhFinance),
		fikkhMarkup.Row(fikkhBammab),
		fikkhMarkup.Row(fikkhVpShvi),
		fikkhMarkup.Row(fikkhZhf),
		fikkhMarkup.Row(fikkhZzaSh),
		fikkhMarkup.Row(fikkhI),
		fikkhMarkup.Row(fikkhKvhfu),
		fikkhMarkup.Row(fikkhluknp),
		fikkhMarkup.Row(fikkhM),
		fikkhMarkup.Row(fikkhOpn),
		fikkhMarkup.Row(fikkhOpssp),
		fikkhMarkup.Row(fikkhOp),
		fikkhMarkup.Row(fikkhOkn),
		fikkhMarkup.Row(fikkhPtvi),
		fikkhMarkup.Row(fikkhPsss),
		fikkhMarkup.Row(fikkhRpk),
		fikkhMarkup.Row(fikkhSm),
		fikkhMarkup.Row(fikkhSoobib),
		fikkhMarkup.Row(fikkhUfikam),
		fikkhMarkup.Row(fikkhHihvp),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	brakMarkup.Reply(
		brakMarkup.Row(brakVt),
		brakMarkup.Row(brakPm),
		brakMarkup.Row(brakSvadbasi),
		brakMarkup.Row(brakSvatovstvosi),
		brakMarkup.Row(brakSo),
		brakMarkup.Row(brakEba),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	jihadMarkup.Reply(
		jihadMarkup.Row(jihadVvab),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	clearingMarkup.Reply(
		clearingMarkup.Row(clearingMabko),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	worshipMarkup.Reply(
		worshipMarkup.Row(worshipZakyat, worshipMolitva),
		worshipMarkup.Row(worshipMolba, worshipPokayanie),
		worshipMarkup.Row(worshipPost, worshipKhidzhra),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	zakyatMarkup.Reply(
		zakyatMarkup.Row(zakyatZaf, zakyatZ),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	molitvaMarkup.Reply(
		molitvaMarkup.Row(molitvaDnm),
		molitvaMarkup.Row(molitvaDm),
		molitvaMarkup.Row(molitvaI),
		molitvaMarkup.Row(molitvaKdpm),
		molitvaMarkup.Row(molitvaPn),
		molitvaMarkup.Row(molitvaPnaia),
		molitvaMarkup.Row(molitvaT),
		molitvaMarkup.Row(molitvaUmSha),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	molbaMarkup.Reply(
		molbaMarkup.Row(molbaPm),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	pokayanieMarkup.Reply(
		pokayanieMarkup.Row(pokayanieKpzKhzetokzn),
		pokayanieMarkup.Row(pokayaniePivssn),
		pokayanieMarkup.Row(pokayanieYKhp),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	postMarkup.Reply(
		postMarkup.Row(postPv6dSh),
		postMarkup.Row(postOpikuaa),
		postMarkup.Row(postKpppaaiaar),
		postMarkup.Row(postDpinmvraib),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	khidzhraMarkup.Reply(
		khidzhraMarkup.Row(khidzhraPKhdi),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	fzhMarkup.Reply(
		fzhMarkup.Row(fzhOkZhv),
		fzhMarkup.Row(fzhTKhoma),
		fzhMarkup.Row(fzhKhniipvKhf),
		fzhMarkup.Row(fzhKhuau),
		fzhMarkup.Row(fzhEpkvZhaf),
		fzhMarkup.Row(fzhTKhomArab),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	financeMarkup.Reply(
		financeMarkup.Row(financeFfo1t),
		financeMarkup.Row(financeFfo2t),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	differentMarkup.Reply(
		differentMarkup.Row(different20smsdevz),
		differentMarkup.Row(differentVdvi),
		differentMarkup.Row(differentGotChpvsoib),
		differentMarkup.Row(differentGopcaChablmr),
		differentMarkup.Row(differentSui),
		differentMarkup.Row(differentTpppksgi),
		differentMarkup.Row(differentFik),
		differentMarkup.Row(differentIkvkpn),
		differentMarkup.Row(differentItpaipSh),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	hadithCollectionMarkup.Reply(
		hadithCollectionMarkup.Row(hadithCollection40Khiu),
		hadithCollectionMarkup.Row(hadithCollectionBam),
		hadithCollectionMarkup.Row(hadithCollectionIhab),
		hadithCollectionMarkup.Row(hadithCollectionIh),
		hadithCollectionMarkup.Row(hadithCollectionPptKh),
		hadithCollectionMarkup.Row(hadithCollectionSp),
		hadithCollectionMarkup.Row(hadithCollectionSab),
		hadithCollectionMarkup.Row(hadithCollectionSaDzhas),
		hadithCollectionMarkup.Row(hadithCollectionSim),
		hadithCollectionMarkup.Row(hadithCollectionSm),
		hadithCollectionMarkup.Row(hadithCollectionSad),
		hadithCollectionMarkup.Row(hadithCollectionSabki),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	siraMarkup.Reply(
		siraMarkup.Row(siraDiaim),
		siraMarkup.Row(siraZhpm),
		siraMarkup.Row(siraIZhs),
		siraMarkup.Row(siraKZhp),
		siraMarkup.Row(siraRop),
		siraMarkup.Row(siraRiZhsmab),
		siraMarkup.Row(siraSpiKh),
		siraMarkup.Row(siraSpk),
		siraMarkup.Row(siraSpm),
		siraMarkup.Row(siraTiu),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	tafsirMarkup.Reply(
		tafsirMarkup.Row(tafsirPsaKaa),
		tafsirMarkup.Row(tafsirPsaKek),
		tafsirMarkup.Row(tafsirSsonT),
		tafsirMarkup.Row(tafsirTiasak),
		tafsirMarkup.Row(tafsirTsakik),
		tafsirMarkup.Row(tafsirTsabau),
		tafsirMarkup.Row(tafsirTsaaShuiit),
		tafsirMarkup.Row(tafsirFtia),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	knowledgeNeedingMarkup.Reply(
		knowledgeNeedingMarkup.Row(knowledgeNeedingKnpz),
		knowledgeNeedingMarkup.Row(knowledgeNeedingRKhptz),
		knowledgeNeedingMarkup.Row(knowledgeNeedingRCh),
		knowledgeNeedingMarkup.Row(knowledgeNeedingUiz),
		knowledgeNeedingMarkup.Row(knowledgeNeedingUpdiz),
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	selector.Reply(
		selector.Row(selectorBack, selectorBackAtAllBtn),
	)

	b.Handle("/start", func(c tele.Context) error {
		err := storage.NewUser(c.Sender().ID)
		if err != nil {
			return err
		}

		return c.Send(welcomeText, mainMarkup)
	})

	b.Handle(&mainAkida, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Акыда")
		if err != nil {
			return err
		}

		return c.Send("Акыда", akidaMarkup)
	})

	b.Handle(&akidaTo, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Три основы")
		if err != nil {
			return err
		}

		return c.Send("Три основы", ToMarkup)
	})

	b.Handle(&ToToir, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Три Основы/Три основы избранные разъяснения.pdf")}
		pdf.FileName = "Три основы избранные разъяснения.pdf"
		return c.Send(pdf)
	})

	b.Handle(&ToTom, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Три Основы/ТРИ ОСНОВЫ МАТН.pdf")}
		pdf.FileName = "ТРИ ОСНОВЫ МАТН.pdf"
		return c.Send(pdf)
	})

	b.Handle(&ToShib, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Три Основы/Три основы шарх ибн Баз.pdf")}
		pdf.FileName = "Три основы шарх ибн Баз.pdf"
		return c.Send(pdf)
	})

	b.Handle(&ToShik, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Три Основы/Три основы шарх ибн Касим.pdf")}
		pdf.FileName = "Три основы шарх ибн Касим.pdf"
		return c.Send(pdf)
	})

	b.Handle(&ToShiu, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Три Основы/Три основы шарх ибн Усаймин.pdf")}
		pdf.FileName = "Три основы шарх ибн Усаймин.pdf"
		return c.Send(pdf)
	})

	b.Handle(&ToShu, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Три Основы/Три основы шарх Усойми.pdf")}
		pdf.FileName = "Три основы шарх Усойми.pdf"
		return c.Send(pdf)
	})

	b.Handle(&ToShf, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Три Основы/Три основы шарх Фаузан.pdf")}
		pdf.FileName = "Три основы шарх Фаузан.pdf"
		return c.Send(pdf)
	})

	b.Handle(&akidaChp, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Четыре правила")
		if err != nil {
			return err
		}

		return c.Send("Четыре правила", ChpMarkup)
	})

	b.Handle(&Chp4paSh, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		docx := &tele.Document{File: tele.FromDisk("./media/Акыда/Четыре правила/4 правила - аль-Шейх.docx")}
		docx.FileName = "4 правила - аль-Шейх.docx"
		return c.Send(docx)
	})

	b.Handle(&Chp4pu, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		docx := &tele.Document{File: tele.FromDisk("./media/Акыда/Четыре правила/4 правила - Усайми.docx")}
		docx.FileName = "4 правила - Усайми.docx"
		return c.Send(docx)
	})

	b.Handle(&ChpShal, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Четыре правила/4 правила шарх аль-Люхайдан.pdf")}
		pdf.FileName = "4 правила шарх аль-Люхайдан.pdf"
		return c.Send(pdf)
	})

	b.Handle(&ChpISh, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Четыре правила/Четыре правила избранные шурухи.pdf")}
		pdf.FileName = "Четыре правила избранные шурухи.pdf"
		return c.Send(pdf)
	})

	b.Handle(&ChpM, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Четыре правила/Четыре правила матн.pdf")}
		pdf.FileName = "Четыре правила матн.pdf"
		return c.Send(pdf)
	})

	b.Handle(&ChpSharab, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Четыре правила/Четыре правила шарх АбдуРРахман аль-Баррок.pdf")}
		pdf.FileName = "Четыре правила шарх АбдуРРахман аль-Баррок.pdf"
		return c.Send(pdf)
	})

	b.Handle(&ChpShaDzh, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Четыре правила/Четыре правила шарх Абу Джабир.pdf")}
		pdf.FileName = "Четыре правила шарх Абу Джабир.pdf"
		return c.Send(pdf)
	})

	b.Handle(&ChpShib, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Четыре правила/Четыре правила шарх ибн Баз.pdf")}
		pdf.FileName = "Четыре правила шарх ибн Баз.pdf"
		return c.Send(pdf)
	})

	b.Handle(&ChpShsaSh, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Четыре правила/Четыре правила шарх Солих али-Шейх.pdf")}
		pdf.FileName = "Четыре правила шарх Солих али-Шейх.pdf"
		return c.Send(pdf)
	})

	b.Handle(&ChpShf, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Четыре правила/Четыре правила шарх Фаузан.pdf")}
		pdf.FileName = "Четыре правила шарх Фаузан.pdf"
		return c.Send(pdf)
	})

	b.Handle(&akidaSho, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Шесть основ")
		if err != nil {
			return err
		}

		return c.Send("Шесть основ", ShoMarkup)
	})

	b.Handle(&ShoShom, func(c tele.Context) error {
		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/шесть основ/Шесть основ матн.pdf")}
		pdf.FileName = "Шесть основ матн.pdf"
		return c.Send(pdf)
	})

	b.Handle(&ShoShoShad, func(c tele.Context) error {
		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/шесть основ/Шесть основ шарх Абу Джабир.pdf")}
		pdf.FileName = "Шесть основ шарх Абу Джабир.pdf"
		return c.Send(pdf)
	})

	b.Handle(&ShoSharab, func(c tele.Context) error {
		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/шесть основ/Шесть основ шарх АдбуРразак аль-Бадр.pdf")}
		pdf.FileName = "Шесть основ шарх АдбуРразак аль-Бадр.pdf"
		return c.Send(pdf)
	})

	b.Handle(&ShoShoshf, func(c tele.Context) error {
		docx := &tele.Document{File: tele.FromDisk("./media/Акыда/шесть основ/Шесть основ шарх Фаузан.doc")}
		docx.FileName = "Шесть основ шарх Фаузан.doc"
		return c.Send(docx)
	})

	b.Handle(&akidaOs, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Отведение сомнений")
		if err != nil {
			return err
		}

		return c.Send("Отведение сомнений", OsMarkup)
	})

	b.Handle(&OsOru, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Отведение сомнений/Отведение разъяснения ученых.pdf")}
		pdf.FileName = "Отведение разъяснения ученых.pdf"
		return c.Send(pdf)
	})

	b.Handle(&OsOsm, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Отведение сомнений/Отведение сомнений матн.pdf")}
		pdf.FileName = "Отведение сомнений матн.pdf"
		return c.Send(pdf)
	})

	b.Handle(&OsOsShg, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Отведение сомнений/Отведение сомнений шарх Гунейман.pdf")}
		pdf.FileName = "Отведение сомнений шарх Гунейман.pdf"
		return c.Send(pdf)
	})

	b.Handle(&OsOsShsaSh, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Отведение сомнений/Отведение сомнений шарх Солих али-Шейх.pdf")}
		pdf.FileName = "Отведение сомнений шарх Солих али-Шейх.pdf"
		return c.Send(pdf)
	})

	b.Handle(&akidaKe, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Книга единобожия")
		if err != nil {
			return err
		}

		return c.Send("Книга единобожия", KeMarkup)
	})

	b.Handle(&KeKeab, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Книга единобожия/Книга единобожия аль-Бадр.pdf")}
		pdf.FileName = "Книга единобожия аль-Бадр.pdf"
		return c.Send(pdf)
	})

	b.Handle(&KeKeas, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Книга единобожия/Книга единобожия ас-Саади.pdf")}
		pdf.FileName = "Книга единобожия ас-Саади.pdf"
		return c.Send(pdf)
	})

	b.Handle(&KeKekShsaSh, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Книга единобожия/Книга единобожия короткий шарх Солиха али-Шейх.pdf")}
		pdf.FileName = "Книга единобожия короткий шарх Солиха али-Шейх.pdf"
		return c.Send(pdf)
	})

	b.Handle(&KeKekShf, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Книга единобожия/Книга единобожия короткий шарх Фаузана .pdf")}
		pdf.FileName = "Книга единобожия короткий шарх Фаузана.pdf"
		return c.Send(pdf)
	})

	b.Handle(&KeKem, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Книга единобожия/Книга единобожия матн.pdf")}
		pdf.FileName = "Книга единобожия матн.pdf"
		return c.Send(pdf)
	})

	b.Handle(&KeKeShAraSh, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Книга единобожия/Книга единобожия шарх АбдуРрахман али-Шейх.pdf")}
		pdf.FileName = "Книга единобожия шарх АбдуРрахман али-Шейх.pdf"
		return c.Send(pdf)
	})

	b.Handle(&KeKeshad, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Книга единобожия/Книга единобожия шарх Абу Джабир.pdf")}
		pdf.FileName = "Книга единобожия шарх Абу Джабир.pdf"
		return c.Send(pdf)
	})

	b.Handle(&KeShia, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Книга единобожия/Книга единобожия шарх ибн Атик.pdf")}
		pdf.FileName = "Книга единобожия шарх ибн Атик.pdf"
		return c.Send(pdf)
	})

	b.Handle(&akidaNi, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Навакыдуль ислям")
		if err != nil {
			return err
		}

		return c.Send("Навакыдуль ислям", NiMarkup)
	})

	b.Handle(&Ni10paiShb, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Навакыдуль ислям/10 пунктов аннулирующих ислам шарх Баррак.pdf")}
		pdf.FileName = "10 пунктов аннулирующих ислам шарх Баррак.pdf"
		return c.Send(pdf)
	})

	b.Handle(&Ni10paiShr, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Навакыдуль ислям/10 пунктов аннулирующих ислам шарх Роджихи.pdf")}
		pdf.FileName = "10 пунктов аннулирующих ислам шарх Роджихи.pdf"
		return c.Send(pdf)
	})

	b.Handle(&Ni10paiShf, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Навакыдуль ислям/10 пунктов аннулирующих ислам — Шейх Фаузан.pdf")}
		pdf.FileName = "10 пунктов аннулирующих ислам — Шейх Фаузан.pdf"
		return c.Send(pdf)
	})

	b.Handle(&NiNf, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		docx := &tele.Document{File: tele.FromDisk("./media/Акыда/Навакыдуль ислям/Науакыд - Фаузан.docx")}
		docx.FileName = "Науакыд - Фаузан.docx"
		return c.Send(docx)
	})

	b.Handle(&akidaAviia, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Акыда в именах и атрибутах")
		if err != nil {
			return err
		}

		return c.Send("Акыда в именах и атрибутах", AviiaMarkup)
	})

	b.Handle(&AviiaAvsf, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Акыда в именах и атрибутах/АКЪИДА ВАСАТИЯ САЛИХ ФАУЗАН.pdf")}
		pdf.FileName = "АКЪИДА ВАСАТИЯ САЛИХ ФАУЗАН.pdf"
		return c.Send(pdf)
	})

	b.Handle(&AviiaAatShiaai, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Акыда в именах и атрибутах/Акыда ат-Тахавия шарх ибн аби аль-Изз.pdf")}
		pdf.FileName = "Акыда ат-Тахавия шарх ибн аби аль-Изз.pdf"
		return c.Send(pdf)
	})

	b.Handle(&AviiaBuShShau, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Акыда в именах и атрибутах/Блеск убеждений Шарх шейха аль Усеймин.pdf")}
		pdf.FileName = "Блеск убеждений Шарх шейха аль Усеймин.pdf"
		return c.Send(pdf)
	})

	b.Handle(&AviiaIt, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Акыда в именах и атрибутах/Идеология тафуида.pdf")}
		pdf.FileName = "Идеология тафуида.pdf"
		return c.Send(pdf)
	})

	b.Handle(&AviiaKmpaAaSh, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Акыда в именах и атрибутах/Коранический метод познания атрибутов Аллаха Аш Шанкыти.pdf")}
		pdf.FileName = "Коранический метод познания атрибутов Аллаха Аш Шанкыти.pdf"
		return c.Send(pdf)
	})

	b.Handle(&AviiaPik, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Акыда в именах и атрибутах/Прекрасные имена Кахтани.pdf")}
		pdf.FileName = "Прекрасные имена Кахтани.pdf"
		return c.Send(pdf)
	})

	b.Handle(&AviiaRvsiod, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Акыда в именах и атрибутах/Разногласия в словах и опровержение джахмита.pdf")}
		pdf.FileName = "Разногласия в словах и опровержение джахмита.pdf"
		return c.Send(pdf)
	})

	b.Handle(&AviiaSvv, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Акыда в именах и атрибутах/Середийность в вероубеждений.pdf")}
		pdf.FileName = "Середийность в вероубеждений.pdf"
		return c.Send(pdf)
	})

	b.Handle(&AviiaTaavsvo, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Акыда в именах и атрибутах/Таухид аль-асма ва-сыфат в общем.pdf")}
		pdf.FileName = "Таухид аль-асма ва-сыфат в общем.pdf"
		return c.Send(pdf)
	})

	b.Handle(&akidaVi, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Вопрос имана")
		if err != nil {
			return err
		}

		return c.Send("Вопрос имана", ViMarkup)
	})

	b.Handle(&ViAf, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Вопрос имана/аль-Фурук.pdf")}
		pdf.FileName = "аль-Фурук.pdf"
		return c.Send(pdf)
	})

	b.Handle(&ViKiit, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Вопрос имана/Книга Имана ибн Таймия.pdf")}
		pdf.FileName = "Книга Имана ибн Таймия.pdf"
		return c.Send(pdf)
	})

	b.Handle(&ViUiatShShsaSh, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Вопрос имана/Усулю иман ат-Тамими шарх Шейх Солих али-Шейха.pdf")}
		pdf.FileName = "Усулю иман ат-Тамими шарх Шейх Солих али-Шейха.pdf"
		return c.Send(pdf)
	})

	b.Handle(&ViUkpm, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Вопрос имана/Учёные_комитета_против_мурджиитов.pdf")}
		pdf.FileName = "Учёные комитета против мурджиитов.pdf"
		return c.Send(pdf)
	})

	b.Handle(&akidaLin, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Любовь и непричастность")
		if err != nil {
			return err
		}

		return c.Send("Любовь и непричастность", LinMarkup)
	})

	b.Handle(&LinDinvif, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Любовь и непричастность/Дружба и непричастность в исламе Фаузан.pdf")}
		pdf.FileName = "Дружба и непричастность в исламе Фаузан.pdf"
		return c.Send(pdf)
	})

	b.Handle(&LinOlin, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Любовь и непричастность/Основы любви и непричастности.pdf")}
		pdf.FileName = "Основы любви и непричастности.pdf"
		return c.Send(pdf)
	})

	b.Handle(&LinPuvi, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Любовь и непричастность/Принцип Уаля в Исламе.pdf")}
		pdf.FileName = "Принцип Уаля в Исламе.pdf"
		return c.Send(pdf)
	})

	b.Handle(&akidaRvubd, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Разбор вопроса узр биль джахль")
		if err != nil {
			return err
		}

		return c.Send("Разбор вопроса узр биль джахль", RvubdMarkup)
	})

	b.Handle(&RvubdAooo, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Разбор вопроса узр биль джахль/Аяты об отсутствии оправдания.pdf")}
		pdf.FileName = "Аяты об отсутствии оправдания.pdf"
		return c.Send(pdf)
	})

	b.Handle(&RvubdRubd, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Разбор вопроса узр биль джахль/Разбор узр биль джахль.pdf")}
		pdf.FileName = "Разбор узр биль джахль.pdf"
		return c.Send(pdf)
	})

	b.Handle(&RvubdRubd2, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Разбор вопроса узр биль джахль/Разбор узр биль джахль 2.pdf")}
		pdf.FileName = "Разбор узр биль джахль 2.pdf"
		return c.Send(pdf)
	})

	b.Handle(&RvubdRubd3, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Разбор вопроса узр биль джахль/Разбор узр биль джахль 3.pdf")}
		pdf.FileName = "Разбор узр биль джахль 3.pdf"
		return c.Send(pdf)
	})

	b.Handle(&RvubdRubd4, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Разбор вопроса узр биль джахль/Разбор узр биль джахль 4.pdf")}
		pdf.FileName = "Разбор узр биль джахль 4.pdf"
		return c.Send(pdf)
	})

	b.Handle(&akidaKpoaim, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Книги по основам акыды и манхаджа")
		if err != nil {
			return err
		}

		return c.Send("Книги по основам акыды и манхаджа", KpoaimMarkup)
	})

	b.Handle(&KpoaimVu, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Книги по основам акыды и манхаджа/Важные уроки.pdf")}
		pdf.FileName = "Важные уроки.pdf"
		return c.Send(pdf)
	})

	b.Handle(&KpoaimVef, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Книги по основам акыды и манхаджа/Вероубеждение единобожия, Фаузан.pdf")}
		pdf.FileName = "Вероубеждение единобожия, Фаузан.pdf"
		return c.Send(pdf)
	})

	b.Handle(&KpoaimDsid, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Книги по основам акыды и манхаджа/Два свидетельства ибн Джибрин.pdf")}
		pdf.FileName = "Два свидетельства ибн Джибрин.pdf"
		return c.Send(pdf)
	})

	b.Handle(&KpoaimDeab, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Книги по основам акыды и манхаджа/Доказательства Единобожия аль Бадр.pdf")}
		pdf.FileName = "Доказательства Единобожия аль Бадр.pdf"
		return c.Send(pdf)
	})

	b.Handle(&KpoaimOvl, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Книги по основам акыды и манхаджа/Основы вероучения Лялякаи.pdf")}
		pdf.FileName = "Основы вероучения Лялякаи.pdf"
		return c.Send(pdf)
	})

	b.Handle(&KpoaimRopv, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Книги по основам акыды и манхаджа/Разъяснение основых постулатов веры.pdf")}
		pdf.FileName = "Разъяснение основых постулатов веры.pdf"
		return c.Send(pdf)
	})

	b.Handle(&KpoaimRn, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Книги по основам акыды и манхаджа/Религиозные новшества.pdf")}
		pdf.FileName = "Религиозные новшества.pdf"
		return c.Send(pdf)
	})

	b.Handle(&KpoaimSeAab, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Книги по основам акыды и манхаджа/Слова единобожия АбдуРразак аль-Бадр.pdf")}
		pdf.FileName = "Слова единобожия АбдуРразак аль-Бадр.pdf"
		return c.Send(pdf)
	})

	b.Handle(&KpoaimUpsieo, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Книги по основам акыды и манхаджа/Убеждения_приверженцев_сунны_и_единой_общины.pdf")}
		pdf.FileName = "Убеждения приверженцев сунны и единой общины.pdf"
		return c.Send(pdf)
	})

	b.Handle(&KpoaimUiikf, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Книги по основам акыды и манхаджа/Уроки извлекаемые из Корана Фаузан.pdf")}
		pdf.FileName = "Уроки извлекаемые из Корана Фаузан.pdf"
		return c.Send(pdf)
	})

	b.Handle(&KpoaimShasb, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Книги по основам акыды и манхаджа/Шарх ас-Сунна Барбахари.pdf")}
		pdf.FileName = "Шарх ас-Сунна Барбахари.pdf"
		return c.Send(pdf)
	})

	b.Handle(&KpoaimShasb2, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Книги по основам акыды и манхаджа/Шарх ас-сунна (Барбахари) (2).pdf")}
		pdf.FileName = "Шарх ас-сунна (Барбахари) (2).pdf"
		return c.Send(pdf)
	})

	b.Handle(&akidaPkpa, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Полезные книги по акыде")
		if err != nil {
			return err
		}

		return c.Send("Полезные книги по акыде", PkpaMarkup)
	})

	b.Handle(&PkpaPisaf, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Полезные книги по акыде/«Пользы из суры «аль-Фатиха»».pdf")}
		pdf.FileName = "«Пользы из суры \"аль-Фатиха\".pdf»"
		return c.Send(pdf)
	})

	b.Handle(&PkpaAsas, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Полезные книги по акыде/Акыда Суфьяна ас-Саури.pdf")}
		pdf.FileName = "Акыда Суфьяна ас-Саури.pdf"
		return c.Send(pdf)
	})

	b.Handle(&PkpaAsu, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Полезные книги по акыде/Акыда Суфьяна Уйейна.pdf")}
		pdf.FileName = "Акыда Суфьяна Уйейна.pdf"
		return c.Send(pdf)
	})

	b.Handle(&PkpaKasKhak, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Полезные книги по акыде/Китаб ас-Сунна Харб аль-Кирмани.pdf")}
		pdf.FileName = "Китаб ас-Сунна Харб аль-Кирмани.pdf"
		return c.Send(pdf)
	})

	b.Handle(&PkpaMDzhf, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Полезные книги по акыде/МАСАИЛЬ ДЖАХИЛИЯ ФАУЗАН.pdf")}
		pdf.FileName = "МАСАИЛЬ ДЖАХИЛИЯ ФАУЗАН.pdf"
		return c.Send(pdf)
	})

	b.Handle(&PkpaMamat, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Полезные книги по акыде/Муфид аль-Мустафид ат-Тамими.pdf")}
		pdf.FileName = "Муфид аль-Мустафид ат-Тамими.pdf"
		return c.Send(pdf)
	})

	b.Handle(&PkpaR, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Полезные книги по акыде/Различия.pdf")}
		pdf.FileName = "Различия.pdf"
		return c.Send(pdf)
	})

	b.Handle(&PkpaUasiaShaDzh, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Полезные книги по акыде/Усуль ас-Сунна Имама Ахмада шарх Абу Джабир.pdf")}
		pdf.FileName = "Усуль ас-Сунна Имама Ахмада шарх Абу Джабир.pdf"
		return c.Send(pdf)
	})

	b.Handle(&PkpaFpsi, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Полезные книги по акыде/Фетвы по столпам Ислама.pdf")}
		pdf.FileName = "Фетвы по столпам Ислама.pdf"
		return c.Send(pdf)
	})

	b.Handle(&PkpaKhiadShf, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Акыда/Полезные книги по акыде/Хаия ибн Аби Давуда шарх Фаузан.pdf")}
		pdf.FileName = "Хаия ибн Аби Давуда шарх Фаузан.pdf"
		return c.Send(pdf)
	})

	b.Handle(&mainAdab, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Адаб")
		if err != nil {
			return err
		}

		return c.Send("Адаб", adabMarkup)
	})

	b.Handle(&adabZhppp, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Адаб/Женщина придерживающаяся правильного пути.pdf")}
		pdf.FileName = "Женщина придерживающаяся правильного пути.pdf"
		return c.Send(pdf)
	})

	b.Handle(&adabKb, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Адаб/Книга благопристойности.pdf")}
		pdf.FileName = "Книга благопристойности.pdf"
		return c.Send(pdf)
	})

	b.Handle(&adabMNielvsKiS, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Адаб/Мусульманин_и_его_личность_в_свете_Корана_и_Сунны_Карима_Cорокоумова.pdf")}
		pdf.FileName = "Мусульманин и его личность в свете Корана и Сунны Карима Cорокоумова.pdf"
		return c.Send(pdf)
	})

	b.Handle(&adabMKielvKiS, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Адаб/Мусульманка_и_ее_личность_в_свете_Корана_и_Сунны_Карима_Cорокоумова.pdf")}
		pdf.FileName = "Мусульманка и ее личность в свете Корана и Сунны Карима Cорокоумова.pdf"
		return c.Send(pdf)
	})

	b.Handle(&adabPsm, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Адаб/Проблемы современной молодёжи.pdf")}
		pdf.FileName = "Проблемы современной молодёжи.pdf"
		return c.Send(pdf)
	})

	b.Handle(&adabSpezs, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Адаб/Соблюдение_правил_этикета_залог_счастья_Салих_ибн_Абд_аль_Азиз_ибн.pdf")}
		pdf.FileName = "Соблюдение правил этикета залог счастья Салих ибн Абд аль Азиз ибн.pdf"
		return c.Send(pdf)
	})

	b.Handle(&mainDifferent, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Разное")
		if err != nil {
			return err
		}

		return c.Send("Разное", differentMarkup)
	})

	b.Handle(&different20smsdevz, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Разное/20_советов_моей_сестре_до_ее_выхода_замуж.pdf")}
		pdf.FileName = "20 советов моей сестре до ее выхода замуж.pdf"
		return c.Send(pdf)
	})

	b.Handle(&differentVdvi, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Разное/Выбор друзей в Исламе.pdf")}
		pdf.FileName = "Выбор друзей в Исламе.pdf"
		return c.Send(pdf)
	})

	b.Handle(&differentGotChpvsoib, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		docx := &tele.Document{File: tele.FromDisk("./media/Разное/Глава_о_том,_что_при_возникновении_смут_обязательно_искать_безопасность.docx")}
		docx.FileName = "Глава о том, что при возникновении смут обязательно искать безопасность.docx"
		return c.Send(docx)
	})

	b.Handle(&differentGopcaChablmr, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Разное/Грубая_ошибка_—_перед_цитированием_аята_читать_\"А_узу_Би_Лляхи_минаш\".pdf")}
		pdf.FileName = "Грубая ошибка — перед цитированием аята читать «А узу Би Лляхи минаш».pdf"
		return c.Send(pdf)
	})

	b.Handle(&differentSui, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Разное/Способы увеличения имана.pdf")}
		pdf.FileName = "Способы увеличения имана.pdf"
		return c.Send(pdf)
	})

	b.Handle(&differentTpppksgi, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Разное/Табукское_послание_Провизия_переселяющегося_к_своему_Господу_Ибн_Къаййим.pdf")}
		pdf.FileName = "Табукское послание Провизия переселяющегося к своему Господу Ибн Къаййим.pdf"
		return c.Send(pdf)
	})

	b.Handle(&differentFik, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Разное/Фаваид Ибн Каййим.pdf")}
		pdf.FileName = "Фаваид Ибн Каййим.pdf"
		return c.Send(pdf)
	})

	b.Handle(&differentIkvkpn, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Разное/Ibn_Kayyim_-_Vsem_kogo_postiglo_neschastye.pdf")}
		pdf.FileName = "Ибн Къаййим. Всем, кого постигло несчастье.pdf"
		return c.Send(pdf)
	})

	b.Handle(&differentItpaipSh, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Разное/ibn_Taymia_-_Priblizhennye_Allakha_i_priblizhennye_shaytana.pdf")}
		pdf.FileName = "Ибн Теймия. Приближение Аллаха и приближение шайтана.pdf"
		return c.Send(pdf)
	})

	b.Handle(&mainHadithCollection, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Сборники хадисов")
		if err != nil {
			return err
		}

		return c.Send("Сборники хадисов", hadithCollectionMarkup)
	})

	b.Handle(&hadithCollection40Khiu, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Сборники хадисов/40 хадисов ибн Усаймин.pdf")}
		pdf.FileName = "40 хадисов ибн Усаймин.pdf"
		return c.Send(pdf)
	})

	b.Handle(&hadithCollectionBam, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Сборники хадисов/Булуг аль-маррам.pdf")}
		pdf.FileName = "Булуг аль-маррам.pdf"
		return c.Send(pdf)
	})

	b.Handle(&hadithCollectionIhab, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Сборники хадисов/Избранные хадисы Аль-Бухари.pdf")}
		pdf.FileName = "Избранные хадисы Аль-Бухари.pdf"
		return c.Send(pdf)
	})

	b.Handle(&hadithCollectionIh, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Сборники хадисов/Избранные хадисы.pdf")}
		pdf.FileName = "Избранные хадисы.pdf"
		return c.Send(pdf)
	})

	b.Handle(&hadithCollectionPptKh, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Сборники хадисов/Пособие по терминологии хадисов.pdf")}
		pdf.FileName = "Пособие по терминологии хадисов.pdf"
		return c.Send(pdf)
	})

	b.Handle(&hadithCollectionSp, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Сборники хадисов/Сады праведных.pdf")}
		pdf.FileName = "Сады праведных.pdf"
		return c.Send(pdf)
	})

	b.Handle(&hadithCollectionSab, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Сборники хадисов/Сахих аль-Бухари.pdf")}
		pdf.FileName = "Сахих аль-Бухари.pdf"
		return c.Send(pdf)
	})

	b.Handle(&hadithCollectionSaDzhas, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Сборники хадисов/Сахих аль-Джами’ Ас-Саг1ир.pdf")}
		pdf.FileName = "Сахих аль-Джами’ Ас-Саг1ир.pdf"
		return c.Send(pdf)
	})

	b.Handle(&hadithCollectionSim, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Сборники хадисов/Сахих ибн Маджа.pdf")}
		pdf.FileName = "Сахих ибн Маджа.pdf"
		return c.Send(pdf)
	})

	b.Handle(&hadithCollectionSm, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Сборники хадисов/Сахих Муслим.pdf")}
		pdf.FileName = "Сахих Муслим.pdf"
		return c.Send(pdf)
	})

	b.Handle(&hadithCollectionSad, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Сборники хадисов/Сунан Абу Давуд.pdf")}
		pdf.FileName = "Сунан Абу Давуд.pdf"
		return c.Send(pdf)
	})

	b.Handle(&hadithCollectionSabki, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Сборники хадисов/Sakhikh_al-Bukhari_Kratkoe_ikhlozhenie.pdf")}
		pdf.FileName = "Сахих аль-Бухари краткое изложение.pdf"
		return c.Send(pdf)
	})

	b.Handle(&mainSira, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Сира")
		if err != nil {
			return err
		}

		return c.Send("Сира", siraMarkup)
	})

	b.Handle(&siraDiaim, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Сира/Достоверная история Али и Муавии.pdf")}
		pdf.FileName = "Достоверная история Али и Муавии.pdf"
		return c.Send(pdf)
	})

	b.Handle(&siraZhpm, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Сира/Жизнеописание Пророка Мухаммада.pdf")}
		pdf.FileName = "Жизнеописание Пророка Мухаммада.pdf"
		return c.Send(pdf)
	})

	b.Handle(&siraIZhs, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Сира/Из жизни сподвижниц.pdf")}
		pdf.FileName = "Из жизни сподвижниц.pdf"
		return c.Send(pdf)
	})

	b.Handle(&siraKZhp, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Сира/КРАТКОЕ_ЖИЗНЕОПИСАНИЕ_ПРОРОКА.pdf")}
		pdf.FileName = "Краткое жизнеописание Пророка.pdf"
		return c.Send(pdf)
	})

	b.Handle(&siraRop, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Сира/Рассказы о пророках.pdf")}
		pdf.FileName = "Рассказы о пророках.pdf"
		return c.Send(pdf)
	})

	b.Handle(&siraRiZhsmab, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Сира/Рассказы_из_жизни_сподвижников_Мухаммада_Аль_Баша.pdf")}
		pdf.FileName = "Рассказы из жизни сподвижников Мухаммада Аль Баша.pdf"
		return c.Send(pdf)
	})

	b.Handle(&siraSpiKh, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Сира/Сира Пророка ﷺ Ибн Хишам.pdf")}
		pdf.FileName = "Сира Пророка ﷺ Ибн Хишам.pdf"
		return c.Send(pdf)
	})

	b.Handle(&siraSpk, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Сира/Сира Пророка ﷺ Кахтани.pdf")}
		pdf.FileName = "Сира Пророка ﷺ Кахтани.pdf"
		return c.Send(pdf)
	})

	b.Handle(&siraSpm, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Сира/Сира Пророка ﷺ Мубаракфури.pdf")}
		pdf.FileName = "Сира Пророка ﷺ Мубаракфури.pdf"
		return c.Send(pdf)
	})

	b.Handle(&siraTiu, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Сира/Тальха ибн Убайдуллах.pdf")}
		pdf.FileName = "Тальха ибн Убайдуллах.pdf"
		return c.Send(pdf)
	})

	b.Handle(&mainTafsir, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Тафсир")
		if err != nil {
			return err
		}

		return c.Send("Тафсир", tafsirMarkup)
	})

	b.Handle(&tafsirPsaKaa, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Тафсир/Перевод_смыслов_аятов_Корана_Абу_Адель.pdf")}
		pdf.FileName = "Перевод смыслов аятов Корана Абу Адель.pdf"
		return c.Send(pdf)
	})

	b.Handle(&tafsirPsaKek, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Тафсир/Перевод_смыслов_аятов_Корана_Эльмир_Кулиев.pdf")}
		pdf.FileName = "Перевод смыслов аятов Корана Эльмир Кулиев.pdf"
		return c.Send(pdf)
	})

	b.Handle(&tafsirSsonT, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		docx := &tele.Document{File: tele.FromDisk("./media/Тафсир/Сияющее слово о науке Тафсира.docx")}
		docx.FileName = "Сияющее слово о науке Тафсира.docx"
		return c.Send(docx)
	})

	b.Handle(&tafsirTiasak, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Тафсир/Тафсир Ибн Аббаса (сура аль-Кадр).pdf")}
		pdf.FileName = "Тафсир Ибн Аббаса (сура аль-Кадр).pdf"
		return c.Send(pdf)
	})

	b.Handle(&tafsirTsakik, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Тафсир/Тафсир суры аль-Кадр — ибн Касир.pdf")}
		pdf.FileName = "Тафсир суры аль-Кадр — ибн Касир.pdf"
		return c.Send(pdf)
	})

	b.Handle(&tafsirTsabau, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Тафсир/Тафсир_суры_«Аль_Бакара»_аль_Усаймин.pdf")}
		pdf.FileName = "Тафсир суры «Аль Бакара» аль Усаймин.pdf"
		return c.Send(pdf)
	})

	b.Handle(&tafsirTsaaShuiit, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Тафсир/Тафсир_Суры_Аль_Ахзаб_шейх_уль_Ислам_Ибн_Таймия.pdf")}
		pdf.FileName = "Тафсир Суры Аль Ахзаб шейх уль Ислам Ибн Таймия.pdf"
		return c.Send(pdf)
	})

	b.Handle(&tafsirFtia, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Тафсир/Фатиха Тафсир Ибн Аббаса.pdf")}
		pdf.FileName = "Фатиха Тафсир Ибн Аббаса.pdf"
		return c.Send(pdf)
	})

	b.Handle(&mainKnowledgeNeeding, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Требование знаний")
		if err != nil {
			return err
		}

		return c.Send("Требование знаний", knowledgeNeedingMarkup)
	})

	b.Handle(&knowledgeNeedingKnpz, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		docx := &tele.Document{File: tele.FromDisk("./media/Требование знаний/Как начать приобретать знания.docx")}
		docx.FileName = "Как начать приобретать знания.docx"
		return c.Send(docx)
	})

	b.Handle(&knowledgeNeedingRKhptz, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Требование знаний/Разъяснение хадиса про требование знаний.pdf")}
		pdf.FileName = "Разъяснение хадиса про требование знаний.pdf"
		return c.Send(pdf)
	})

	b.Handle(&knowledgeNeedingRCh, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Требование знаний/Рекомендации читателю.pdf")}
		pdf.FileName = "Рекомендации читателю.pdf"
		return c.Send(pdf)
	})

	b.Handle(&knowledgeNeedingUiz, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Требование знаний/Украшение искателя знаний.pdf")}
		pdf.FileName = "Украшение искателя знаний.pdf"
		return c.Send(pdf)
	})

	b.Handle(&knowledgeNeedingUpdiz, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Требование знаний/уникальное_пособие_для_ищущих_знания.pdf")}
		pdf.FileName = "Уникальное пособие для ищущих знания.pdf"
		return c.Send(pdf)
	})

	b.Handle(&mainFikkh, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Фикх")
		if err != nil {
			return err
		}

		return c.Send("Фикх", fikkhMarkup)
	})

	b.Handle(&fikkhBrak, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Брак")
		if err != nil {
			return err
		}

		return c.Send("Брак", brakMarkup)
	})

	b.Handle(&brakVt, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		doc := &tele.Document{File: tele.FromDisk("./media/Фикх/Брак/Важность потомства.doc")}
		doc.FileName = "Важность потомства.doc"
		return c.Send(doc)
	})

	b.Handle(&brakPm, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Брак/Подчинение мужу.pdf")}
		pdf.FileName = "Подчинение мужу.pdf"
		return c.Send(pdf)
	})

	b.Handle(&brakSvadbasi, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Брак/Свадьба согласно Исламу.pdf")}
		pdf.FileName = "Свадьба согласно Исламу.pdf"
		return c.Send(pdf)
	})

	b.Handle(&brakSvatovstvosi, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Брак/Сватовство согласно исламу.pdf")}
		pdf.FileName = "Сватовство согласно исламу.pdf"
		return c.Send(pdf)
	})

	b.Handle(&brakSo, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Брак/Супружеские ошибки.pdf")}
		pdf.FileName = "Супружеские ошибки.pdf"
		return c.Send(pdf)
	})

	b.Handle(&brakEba, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Брак/Этикет бракосочетания Альбани.pdf")}
		pdf.FileName = "Этикет бракосочетания Альбани.pdf"
		return c.Send(pdf)
	})

	b.Handle(&fikkhJihad, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Джихад")
		if err != nil {
			return err
		}

		return c.Send("Джихад", jihadMarkup)
	})

	b.Handle(&jihadVvab, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Джихад/Вопросы войны аль-Бадр.pdf")}
		pdf.FileName = "Вопросы войны аль-Бадр.pdf"
		return c.Send(pdf)
	})

	b.Handle(&fikkhClearing, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Очищение")
		if err != nil {
			return err
		}

		return c.Send("Очищение", clearingMarkup)
	})

	b.Handle(&clearingMabko, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		docx := &tele.Document{File: tele.FromDisk("./media/Фикх/Очищение/Мухтасар аль Бувейты книга омовения.docx")}
		docx.FileName = "Мухтасар аль Бувейты книга омовения.docx"
		return c.Send(docx)
	})

	b.Handle(&fikkhWorship, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Поклонение")
		if err != nil {
			return err
		}

		return c.Send("Поклонение", worshipMarkup)
	})

	b.Handle(&worshipZakyat, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Закят")
		if err != nil {
			return err
		}

		return c.Send("Закят", zakyatMarkup)
	})

	b.Handle(&zakyatZaf, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Поклонение/Закят/Закят аль-Фитр.pdf")}
		pdf.FileName = "Закят аль-Фитр.pdf"
		return c.Send(pdf)
	})

	b.Handle(&zakyatZ, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Поклонение/Закят/Закят Умар Абуль Хасан.pdf")}
		pdf.FileName = "Закят.pdf"
		return c.Send(pdf)
	})

	b.Handle(&worshipMolitva, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Молитва")
		if err != nil {
			return err
		}

		return c.Send("Молитва", molitvaMarkup)
	})

	b.Handle(&molitvaDnm, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		docx := &tele.Document{File: tele.FromDisk("./media/Фикх/Поклонение/Молитва/Добровольные ночные молитвы.docx")}
		docx.FileName = "Добровольные ночные молитвы.docx"
		return c.Send(docx)
	})

	b.Handle(&molitvaDm, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Поклонение/Молитва/Достоинство молитвы.pdf")}
		pdf.FileName = "Достоинство молитвы.pdf"
		return c.Send(pdf)
	})

	b.Handle(&molitvaI, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Поклонение/Молитва/Истихара.pdf")}
		pdf.FileName = "Истихара.pdf"
		return c.Send(pdf)
	})

	b.Handle(&molitvaKdpm, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Поклонение/Молитва/Когда дозволено прерывать молитву.pdf")}
		pdf.FileName = "Когда дозволено прерывать молитву.pdf"
		return c.Send(pdf)
	})

	b.Handle(&molitvaPn, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Поклонение/Молитва/Праздничный намаз.pdf")}
		pdf.FileName = "Праздничный намаз.pdf"
		return c.Send(pdf)
	})

	b.Handle(&molitvaPnaia, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Поклонение/Молитва/Праздничный намаз — адабы и ахкамы.pdf")}
		pdf.FileName = "Праздничный намаз — адабы и ахкамы.pdf"
		return c.Send(pdf)
	})

	b.Handle(&molitvaT, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		docx := &tele.Document{File: tele.FromDisk("./media/Фикх/Поклонение/Молитва/Таравих.docx")}
		docx.FileName = "Таравих.docx"
		return c.Send(docx)
	})

	b.Handle(&molitvaUmSha, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Поклонение/Молитва/Условия молитвы шарх Аббад.pdf")}
		pdf.FileName = "Условия молитвы шарх Аббад.pdf"
		return c.Send(pdf)
	})

	b.Handle(&worshipMolba, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Мольба")
		if err != nil {
			return err
		}

		return c.Send("Мольба", molbaMarkup)
	})

	b.Handle(&molbaPm, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Поклонение/Мольба/Положение мольбы.pdf")}
		pdf.FileName = "Положение мольбы.pdf"
		return c.Send(pdf)
	})

	b.Handle(&worshipPokayanie, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Покаяние")
		if err != nil {
			return err
		}

		return c.Send("Покаяние", pokayanieMarkup)
	})

	b.Handle(&pokayanieKpzKhzetokzn, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Поклонение/Покаяние/Как покаяться за хулу злословие, если тот, о ком злословили, ничего.pdf")}
		pdf.FileName = "Как покаяться за хулу злословие, если тот, о ком злословили, ничего.pdf"
		return c.Send(pdf)
	})

	b.Handle(&pokayaniePivssn, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Поклонение/Покаяние/Покаяние и все связанное с ним.pdf")}
		pdf.FileName = "Покаяние и все связанное с ним.pdf"
		return c.Send(pdf)
	})

	b.Handle(&pokayanieYKhp, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Поклонение/Покаяние/Я хочу покаяться.pdf")}
		pdf.FileName = "Я хочу покаяться.pdf"
		return c.Send(pdf)
	})

	b.Handle(&worshipPost, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Пост")
		if err != nil {
			return err
		}

		return c.Send("Пост", postMarkup)
	})

	b.Handle(&postPv6dSh, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		docx := &tele.Document{File: tele.FromDisk("./media/Фикх/Поклонение/Пост/Пост в 6 дней шавваля.docx")}
		docx.FileName = "Пост в 6 дней шавваля.docx"
		return c.Send(docx)
	})

	b.Handle(&postOpikuaa, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Поклонение/Пост/О посте из книги Умдат аль-Ахкам.pdf")}
		pdf.FileName = "О посте из книги Умдат аль-Ахкам.pdf"
		return c.Send(pdf)
	})

	b.Handle(&postKpppaaiaar, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Поклонение/Пост/Краткое пояснение положений поста Абдуль Азиз ибн Абдуллах Ар Раджихи.pdf")}
		pdf.FileName = "Краткое пояснение положений поста Абдуль Азиз ибн Абдуллах Ар Раджихи.pdf"
		return c.Send(pdf)
	})

	b.Handle(&postDpinmvraib, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Поклонение/Пост/Достоинство поста и ночных молитв в Рамадан Абдульазиз ибн Баз.pdf")}
		pdf.FileName = "Достоинство поста и ночных молитв в Рамадан Абдульазиз ибн Баз.pdf"
		return c.Send(pdf)
	})

	b.Handle(&worshipKhidzhra, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Хиджра")
		if err != nil {
			return err
		}

		return c.Send("Хиджра", khidzhraMarkup)
	})

	b.Handle(&khidzhraPKhdi, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Поклонение/Хиджра/Положение хиджры в исламе.pdf")}
		pdf.FileName = "Положение хиджры в исламе.pdf"
		return c.Send(pdf)
	})

	b.Handle(&fikkhFzh, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Фикх женщин")
		if err != nil {
			return err
		}

		return c.Send("Фикх женщин", fzhMarkup)
	})

	b.Handle(&fzhOkZhv, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Фикх женщин/О каждодневных женских выделениях.pdf")}
		pdf.FileName = "О каждодневных женских выделениях.pdf"
		return c.Send(pdf)
	})

	b.Handle(&fzhTKhoma, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Фикх женщин/Твой хиджаб, о мусульманка! 'Абдур.pdf")}
		pdf.FileName = "Твой хиджаб, о мусульманка! 'Абдур.pdf"
		return c.Send(pdf)
	})

	b.Handle(&fzhKhniipvKhf, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Фикх женщин/Хайд, нифас и истихада положение в ханбалитском фикхе.pdf")}
		pdf.FileName = "Хайд, нифас и истихада положение в ханбалитском фикхе.pdf"
		return c.Send(pdf)
	})

	b.Handle(&fzhKhuau, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		docx := &tele.Document{File: tele.FromDisk("./media/Фикх/Фикх женщин/Хиджаб (‘Умар аль-‘Умар).docx")}
		docx.FileName = "Хиджаб (‘Умар аль-‘Умар).docx"
		return c.Send(docx)
	})

	b.Handle(&fzhEpkvZhaf, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Фикх женщин/EQAMO Постановления касающиеся верующих женщин аль Фаузан.pdf")}
		pdf.FileName = "EQAMO Постановления касающиеся верующих женщин аль Фаузан"
		return c.Send(pdf)
	})

	b.Handle(&fikkhFinance, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Финансы")
		if err != nil {
			return err
		}

		return c.Send("Финансы", financeMarkup)
	})

	b.Handle(&financeFfo1t, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Финансы/Фикх финансовых отношений - 1 том.pdf")}
		pdf.FileName = "Фикх финансовых отношений — 1 том.pdf"
		return c.Send(pdf)
	})

	b.Handle(&financeFfo2t, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		docx := &tele.Document{File: tele.FromDisk("./media/Фикх/Финансы/Фикх финансовых отношений — 2 том.docx")}
		docx.FileName = "Фикх финансовых отношений — 2 том.docx"
		return c.Send(docx)
	})

	b.Handle(&fzhTKhomArab, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Фикх женщин/EQAMO Твой хиджаб, о мусульманка! Абдур Разак аль Бадр.pdf")}
		pdf.FileName = "EQAMO Твой хиджаб, о мусульманка! Абдур Разак аль Бадр.pdf"
		return c.Send(pdf)
	})

	b.Handle(&fikkhBammab, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Булуг аль-марам Мухсин аль-Бармауи.pdf")}
		pdf.FileName = "Булуг аль-марам Мухсин аль-Бармауи.pdf"
		return c.Send(pdf)
	})

	b.Handle(&fikkhVpShvi, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Все про шутки в исламе.pdf")}
		pdf.FileName = "Все про шутки в исламе.pdf"
		return c.Send(pdf)
	})

	b.Handle(&fikkhZhf, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Жемчужина фикха.pdf")}
		pdf.FileName = "Жемчужина фикха.pdf"
		return c.Send(pdf)
	})

	b.Handle(&fikkhZzaSh, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		docx := &tele.Document{File: tele.FromDisk("./media/Фикх/Законоположения зимы аш Шувей’ир.docx")}
		docx.FileName = "Законоположения зимы аш Шувей’ир.docx"
		return c.Send(docx)
	})

	b.Handle(&fikkhI, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Ихтилят.pdf")}
		pdf.FileName = "Ихтилят.pdf"
		return c.Send(pdf)
	})

	b.Handle(&fikkhKvhfu, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Ключ в ханбалитском фикхе Усойми.pdf")}
		pdf.FileName = "Ключ в ханбалитском фикхе Усойми.pdf"
		return c.Send(pdf)
	})

	b.Handle(&fikkhluknp, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Лайлат уль Кадр – Ночь Предопределения!!!.pdf")}
		pdf.FileName = "Лайлат уль Кадр – Ночь Предопределения!!!.pdf"
		return c.Send(pdf)
	})

	b.Handle(&fikkhM, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Махрамы.pdf")}
		pdf.FileName = "Махрамы.pdf"
		return c.Send(pdf)
	})

	b.Handle(&fikkhOpn, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/О праздниках неверующих.pdf")}
		pdf.FileName = "О праздниках неверующих.pdf"
		return c.Send(pdf)
	})

	b.Handle(&fikkhOpssp, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/О положениях, связанных с поздравлениями.pdf")}
		pdf.FileName = "О положениях, связанных с поздравлениями.pdf"
		return c.Send(pdf)
	})

	b.Handle(&fikkhOp, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Обряды похорон.pdf")}
		pdf.FileName = "Обряды похорон.pdf"
		return c.Send(pdf)
	})

	b.Handle(&fikkhOkn, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Отношение к неверующим.pdf")}
		pdf.FileName = "Отношение к неверующим.pdf"
		return c.Send(pdf)
	})

	b.Handle(&fikkhPtvi, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Положение тазкии в Исламе.pdf")}
		pdf.FileName = "Положение тазкии в Исламе.pdf"
		return c.Send(pdf)
	})

	b.Handle(&fikkhPsss, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Положения, связанные со сновидениями.pdf.pdf")}
		pdf.FileName = "Положения, связанные со сновидениями.pdf"
		return c.Send(pdf)
	})

	b.Handle(&fikkhRpk, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Рукья посредством Корана.pdf")}
		pdf.FileName = "Рукья посредством Корана.pdf"
		return c.Send(pdf)
	})

	b.Handle(&fikkhSm, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Следование мазхабу.pdf")}
		pdf.FileName = "Следование мазхабу.pdf"
		return c.Send(pdf)
	})

	b.Handle(&fikkhSoobib, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		docx := &tele.Document{File: tele.FromDisk("./media/Фикх/Суждение относительно отпускания бороды ибн Баз.docx")}
		docx.FileName = "Суждение относительно отпускания бороды ибн Баз.docx"
		return c.Send(docx)
	})

	b.Handle(&fikkhUfikam, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Умдатуль фикх ибн кудама аль-макдиси.pdf")}
		pdf.FileName = "Умдатуль фикх ибн кудама аль-макдиси.pdf"
		return c.Send(pdf)
	})

	b.Handle(&fikkhHihvp, func(c tele.Context) error {
		err = c.Notify(tele.UploadingDocument)
		if err != nil {
			return err
		}

		pdf := &tele.Document{File: tele.FromDisk("./media/Фикх/Харам и халяль в пище.pdf")}
		pdf.FileName = "Харам и халяль в пище.pdf"
		return c.Send(pdf)
	})

	b.Handle(&selectorBackAtAllBtn, func(c tele.Context) error {
		err := storage.SetUserState(c.Sender().ID, "Main")
		if err != nil {
			return err
		}

		return c.Send(welcomeText, mainMarkup)
	})

	b.Handle(&selectorBack, func(c tele.Context) error {
		id := c.Sender().ID
		keyToPreviousStep, err := storage.GetUserState(id)
		if err != nil {
			return err
		}

		if keyToPreviousStep == "Main" {
			return nil
		}

		previousStepState, previousStepMarkup := m[keyToPreviousStep][0].(string), m[keyToPreviousStep][1].(*tele.ReplyMarkup)
		err = storage.SetUserState(id, previousStepState)
		if err != nil {
			return err
		}

		return c.Send("Назад", previousStepMarkup)
	})

	log.Printf("%s started...", b.Me.Username)
	b.Start()
}
