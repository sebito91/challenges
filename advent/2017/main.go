package main

import (
	"fmt"

	"github.com/sebito91/advent/2017/eight"
)

// 1122 produces a sum of 3 (1 + 2) because the first digit (1) matches the second digit and the third digit (2) matches the fourth digit.
// 1111 produces 4 because each digit (all 1) matches the next.
// 1234 produces 0 because no digit matches the next.
// 91212129 produces 9 because the only digit that matches the next one is the last digit, 9.

// main
func main() {
	fmt.Printf("Running each of the 2017 Advent of Code challenges...\n")

	fmt.Printf("*****************DAY EIGHT PART ONE*********************\n")

	id, val := eight.DayEightPartOne()
	fmt.Printf("codec in DayEightPartOne: %s, %d\n", id, val)

	fmt.Printf("*****************DAY EIGHT PART TWO*********************\n")

	id, val = eight.DayEightPartTwo()
	fmt.Printf("codec in DayEightPartTwo: %s, %d\n", id, val)

	//
	//	fmt.Printf("*****************DAY SEVEN PART ONE*********************\n")
	//
	//	fmt.Printf("codec in DaySevenPartOne: %s\n", seven.DaySevenPartOne())
	//
	//	fmt.Printf("*****************DAY SEVEN PART TWO*********************\n")
	//
	//	fmt.Printf("codec in DaySevenPartTwo: %d\n", seven.DaySevenPartTwo())

	//	fmt.Printf("*****************DAY SIX PART TWO*********************\n")
	//
	//	if err := six.DaySixPartTwo(); err != nil {
	//		fmt.Printf("error in DaySixPartTwo: %+v\n", err)
	//	}
	//
	//	fmt.Printf("*****************DAY SIX PART ONE*********************\n")
	//
	//	if err := six.DaySixPartOne(); err != nil {
	//		fmt.Printf("error in DaySixPartOne: %+v\n", err)
	//	}
	//
	//	fmt.Printf("*****************DAY SIX PART TWO*********************\n")
	//
	//	if err := six.DaySixPartTwo(); err != nil {
	//		fmt.Printf("error in DaySixPartTwo: %+v\n", err)
	//	}

	//	fmt.Printf("*****************DAY FIVE PART ONE********************\n")
	//
	//	if err := five.DayFivePartOne(); err != nil {
	//		fmt.Printf("error in DayFivePartOne: %+v\n", err)
	//	}
	//
	//	fmt.Printf("*****************DAY FIVE PART TWO********************\n")
	//
	//	if err := five.DayFivePartTwo(); err != nil {
	//		fmt.Printf("error in DayFivePartTwo: %+v\n", err)
	//	}
	//
	// fmt.Printf("*****************DAY FIVE PART TWO********************\n")

	//
	// fmt.Printf("*****************DAY FOUR PART ONE********************\n")
	// fmt.Printf("DayFourPartOne checksum is: %d\n", four.DayFourPartOne(keys))

	// fmt.Printf("*****************DAY FOUR PART TWO********************\n")

	// fmt.Printf("DayFourPartTwo checksum is: %d\n", four.DayFourPartTwo(keys))

	//	fmt.Printf("*****************DAY THREE*******************\n")
	//
	//	for _, x := range []int{1, 12, 23, 1024, 347991} {
	//		if err := three.DayThreePartOne(x); err != nil {
	//			fmt.Printf("error in DayThreePartOne: %+v\n", err)
	//		}
	//	}
	//
	//	fmt.Printf("*****************DAY THREE*******************\n")
	//
	//	for _, x := range []int{1, 10, 25, 806, 1024, 347991} {
	//		if err := three.DayThreePartTwo(x); err != nil {
	//			fmt.Printf("error in DayThreePartTwo: %+v\n", err)
	//		}
	//	}
	//
	//  fmt.Printf("***************DAY TWO***********************\n")
	//
	//	if err := two.DayTwoPartOne(keys); err != nil {
	//		fmt.Printf("error in DayTwoPartOne: %+v\n", err)
	//	}
	//
	//  fmt.Printf("***************DAY TWO***********************\n")
	//
	//	if err := two.DayTwoPartTwo(keys); err != nil {
	//		fmt.Printf("error in DayTwoPartTwo: %+v\n", err)
	//	}
	//
	//	for _, x := range []string{"1122", "1111", "1234", "91212129", str} {
	//		if err := one.DayOne(x); err != nil {
	//			fmt.Printf("failed DayOne (%s)! %+v\n", x, err)
	//		}
	//	}
	//
	//fmt.Printf("*********************************************\n")
	//
	//	for _, x := range []string{"1212", "1221", "123425", "123123", "12131415", str} {
	//		if err := one.DayOnePartTwo(x); err != nil {
	//			fmt.Printf("failed DayOnePartTwo (%s)! %+v\n", x, err)
	//		}
	//	}
	//

}

// DayThree key
//var key = 347991

// DayTwo key
//var keys = [][]int{
//	[]int{5806, 6444, 1281, 38, 267, 1835, 223, 4912, 5995, 230, 4395, 2986, 6048, 4719, 216, 1201},
//	[]int{74, 127, 226, 84, 174, 280, 94, 159, 198, 305, 124, 106, 205, 99, 177, 294},
//	[]int{1332, 52, 54, 655, 56, 170, 843, 707, 1273, 1163, 89, 23, 43, 1300, 1383, 1229},
//	[]int{5653, 236, 1944, 3807, 5356, 246, 222, 1999, 4872, 206, 5265, 5397, 5220, 5538, 286, 917},
//	[]int{3512, 3132, 2826, 3664, 2814, 549, 3408, 3384, 142, 120, 160, 114, 1395, 2074, 1816, 2357},
//	[]int{100, 2000, 112, 103, 2122, 113, 92, 522, 1650, 929, 1281, 2286, 2259, 1068, 1089, 651},
//	[]int{646, 490, 297, 60, 424, 234, 48, 491, 245, 523, 229, 189, 174, 627, 441, 598},
//	[]int{2321, 555, 2413, 2378, 157, 27, 194, 2512, 117, 140, 2287, 277, 2635, 1374, 1496, 1698},
//	[]int{101, 1177, 104, 89, 542, 2033, 1724, 1197, 474, 1041, 1803, 770, 87, 1869, 1183, 553},
//	[]int{1393, 92, 105, 1395, 1000, 85, 391, 1360, 1529, 1367, 1063, 688, 642, 102, 999, 638},
//	[]int{4627, 223, 188, 5529, 2406, 4980, 2384, 2024, 4610, 279, 249, 2331, 4660, 4350, 3264, 242},
//	[]int{769, 779, 502, 75, 1105, 53, 55, 931, 1056, 1195, 65, 292, 1234, 1164, 678, 1032},
//	[]int{2554, 75, 4406, 484, 2285, 226, 5666, 245, 4972, 3739, 5185, 1543, 230, 236, 3621, 5387},
//	[]int{826, 4028, 4274, 163, 5303, 4610, 145, 5779, 157, 4994, 5053, 186, 5060, 3082, 2186, 4882},
//	[]int{588, 345, 67, 286, 743, 54, 802, 776, 29, 44, 107, 63, 303, 372, 41, 810},
//	[]int{128, 2088, 3422, 111, 3312, 740, 3024, 1946, 920, 131, 112, 477, 3386, 2392, 1108, 2741},
//}
//
//// DayOne key
//var str = "3294199471327195994824832197564859876682638188889768298894243832665654681412886862234525991553276578641265589959178414218389329361496673991614673626344552179413995562266818138372393213966143124914469397692587251112663217862879233226763533911128893354536353213847122251463857894159819828724827969576432191847787772732881266875469721189331882228146576832921314638221317393256471998598117289632684663355273845983933845721713497811766995367795857965222183668765517454263354111134841334631345111596131682726196574763165187889337599583345634413436165539744188866156771585647718555182529936669683581662398618765391487164715724849894563314426959348119286955144439452731762666568741612153254469131724137699832984728937865956711925592628456617133695259554548719328229938621332325125972547181236812263887375866231118312954369432937359357266467383318326239572877314765121844831126178173988799765218913178825966268816476559792947359956859989228917136267178571776316345292573489873792149646548747995389669692188457724414468727192819919448275922166321158141365237545222633688372891451842434458527698774342111482498999383831492577615154591278719656798277377363284379468757998373193231795767644654155432692988651312845433511879457921638934877557575241394363721667237778962455961493559848522582413748218971212486373232795878362964873855994697149692824917183375545192119453587398199912564474614219929345185468661129966379693813498542474732198176496694746111576925715493967296487258237854152382365579876894391815759815373319159213475555251488754279888245492373595471189191353244684697662848376529881512529221627313527441221459672786923145165989611223372241149929436247374818467481641931872972582295425936998535194423916544367799522276914445231582272368388831834437562752119325286474352863554693373718848649568451797751926315617575295381964426843625282819524747119726872193569785611959896776143539915299968276374712996485367853494734376257511273443736433464496287219615697341973131715166768916149828396454638596713572963686159214116763"

//// input for DayFourPartOne
//var keys = [][]string{
//	[]string{"una", "bokpr", "ftz", "ryw", "nau", "yknf", "fguaczl", "anu"},
//	[]string{"tvay", "wvco", "bcoblpt", "fwzg", "sfsys", "zvuqll", "mcbhwz", "ovcw", "fgdy"},
//	[]string{"ynsocz", "vid", "rfmsy", "essqt", "fpbjvvq", "sldje", "qfpvjvb"},
//	[]string{"yvh", "nxc", "kla", "vhy", "vkbq", "cxfzgr"},
//	[]string{"kljv", "ymobo", "jlvk", "ebst", "cse", "cmlub", "tavputz", "omoby", "psif"},
//	[]string{"ebfuvvu", "paqgzri", "uvvubef", "hglk", "jvn", "gzeln", "bdl", "ziqgpra", "bzcfa"},
//	[]string{"tclq", "ytw", "tclq", "tezcqys"},
//	[]string{"qae", "eofr", "yzwcwqf", "wqm", "ytcdnc", "pxnmkw"},
//	[]string{"krryi", "irykr", "ycp", "lbeed", "ykrir", "lhq", "rdd", "tyujwd"},
//	[]string{"hms", "pii", "lxoa", "dchsvz", "bepjwst", "bllxkqg", "hsm", "yrdj", "myzvju", "msh", "lwnnc"},
//	[]string{"yxqh", "hqxy", "xkn", "ljjsqjh", "jjljshq"},
//	[]string{"mhgsehj", "urbvnvf", "gbz", "ykxsd", "hsmgehj", "wtoc", "ujpfaos", "eir", "vifog", "tsy", "kdo"},
//	[]string{"wfruf", "wwijme", "pxbbsvf", "asmgs", "ccbn", "vwtc", "mkhah", "oxxfh"},
//	[]string{"lxqy", "jzvcvd", "cfgg", "uahxrwr", "dqmaqr", "bwzm", "wruxhra", "lrb", "lmk"},
//	[]string{"jspgxo", "yuhgni", "njzqtn", "zglkzz", "ybc", "itj", "orqr", "zgqwuoo", "mjagh", "erll", "srqrk"},
//	[]string{"cbrtnbx", "ukblei", "ienmdm", "sinzq", "lytyliz", "mma", "lizylty", "zeumwgu"},
//	[]string{"aeggz", "eljcry", "buqdeog", "dvjzn", "ilvw", "arz", "vep", "kxdzm", "mvh", "szkf"},
//	[]string{"imn", "sfty", "ugg", "flqq", "nydky", "mni", "bkqzlok", "wye", "lehwlmp", "xeyfmj"},
//	[]string{"beyv", "oamrpkc", "tebdkwv", "zlq", "jes", "mqvif", "sej", "qpsnmjz", "edvtbkw"},
//	[]string{"hylmacl", "wwlra", "xntnvg", "ppvb", "bzof", "cymllha"},
//	[]string{"qktxomf", "ngfpuz", "qqz", "malc", "zxuqz", "szg", "zis", "vzro", "rfpgk"},
//	[]string{"phru", "sxlg", "qzqlw", "uej", "vmd", "omzga", "jue"},
//	[]string{"drzgojf", "ttqdqar", "weikik", "wvrjtxi", "gbj", "jramqh", "nlwoj", "drzgojf", "bgabmn", "xqlaeun"},
//	[]string{"aiuohu", "pca", "apkmv", "cpa", "kmvpa", "nmdn"},
//	[]string{"gelymv", "eto", "itcnuhn", "ote", "teo"},
//	[]string{"oxiz", "xzio", "kqu", "wwgow"},
//	[]string{"picoyb", "coibpy", "acsw", "ehlirq", "deyz", "gymqvz", "opicyb", "vuet", "lrl"},
//	[]string{"zerg", "rezg", "miwtjgw", "gezr", "cui"},
//	[]string{"mlh", "qlu", "ktil", "tnlgnrk", "bfqbk", "pgg", "qxeyd", "noadmjo", "nonlsh", "eqxdy"},
//	[]string{"yqqaq", "yqqqa", "xod", "oss"},
//	[]string{"mkotw", "bznvs", "xowoofq", "sebp", "wsgpsmn", "fytcpc", "vvmzr", "vmzrv", "xwtxz", "zrvvm"},
//	[]string{"dvs", "twhz", "teqmlow", "oqg", "sjetxd", "aku", "cwl", "qfvrkex", "mndsio", "hfg"},
//	[]string{"itdl", "qwdagkk", "wli", "tdil", "vlgur", "dyl", "xvfm"},
//	[]string{"wlmyd", "dwmlx", "zhmqd", "zqmhd", "edzsvmz", "yclg", "umpq"},
//	[]string{"petva", "gazw", "byrca", "pvaet", "epkoqh", "nlp", "vqfl", "vatpe"},
//	[]string{"rykn", "ckr", "dijkme", "kmiedj", "ozlh", "deikmj"},
//	[]string{"kuecjh", "sfqv", "pojfzf", "fjopzf", "fjpfzo", "amxtc"},
//	[]string{"hri", "gglmial", "lrwbyc", "memgszu", "hir", "cfwlg", "ylcrwb"},
//	[]string{"rxrfbtv", "pwpra", "fngt", "auh", "rapwp", "zrruuq", "uah"},
//	[]string{"cevc", "utfd", "ysrfcw", "nnlg", "fnqtx", "aua", "htteunu", "mrjhhj"},
//	[]string{"tvnxznj", "mvpagt", "nqmxvnl", "mutn", "ntmu", "eybh", "vkqeaj", "frayclp"},
//	[]string{"ygktzxo", "lzwwy", "qsipu", "nwb", "jdmw", "pahao", "paow", "mwjd", "uivqbnj", "woap", "nyyogcc"},
//	[]string{"log", "zihz", "rsmx", "ixfr", "xwvd", "osg", "zxc", "gol", "ufnbig"},
//	[]string{"dogve", "cnb", "osa", "xbafl", "vefr", "nxlw", "yjgquui"},
//	[]string{"ucyl", "aaoae", "ktjvi", "fyps", "xvjhpbh", "iiur", "tuc"},
//	[]string{"pqt", "jasxg", "ehhs", "lzjzzzl", "sqmmj", "vwyva", "eklbtv", "hksanb", "fuesnd", "oyw", "fuesdn"},
//	[]string{"uankv", "wesi", "frhpau", "qiucu", "lim", "uzbdapf", "ciwjd", "tifbqxh", "tfbtsdi"},
//	[]string{"vgjd", "izpofu", "fqjpcct", "phuz"},
//	[]string{"cfg", "cfg", "rgahl", "frm", "emn", "pbfsmgy", "frm", "jemwqgn", "sfpm", "azunntj", "igl"},
//	[]string{"daps", "hpe", "fqg", "err", "sapd", "dci", "vbzlqx", "gsigq", "eyp", "rre"},
//	[]string{"iuqyqdy", "djprdj", "mgtkdxr", "pwmkzv", "wmkvzp", "hppisd", "pidphs"},
//	[]string{"rlr", "rrl", "vhevh", "cucprc", "xki", "urte", "lrr", "zfc", "xrqydzk", "ipjkyxj", "kytul"},
//	[]string{"jwbkaee", "rgyjl", "rjljbwe", "ppq", "plp", "pfts", "ijd", "ckpvmw", "mbdrqh", "zolt", "lzmr"},
//	[]string{"alw", "law", "awl", "wknavtb", "esklsbj", "wvssyai"},
//	[]string{"aqy", "ldf", "qdht", "soxkg", "qtfipe", "nsdm", "aqe", "rtlc", "fbqrk", "ius", "gybbhxr"},
//	[]string{"xteov", "wgqoqim", "nlz", "szlj", "oxevt", "xwb"},
//	[]string{"tmgdst", "fyn", "oul", "tsohzbq", "ltmxji", "fgops", "gatssx", "zxdzfc", "talb"},
//	[]string{"zkvjpu", "jnhtc", "nxs", "jqv", "pyoqz", "zsj", "ckwd", "xot", "ykai"},
//	[]string{"fxfarre", "yjbxvj", "lqfaglo", "mbxuv", "bmuxv", "bxumv"},
//	[]string{"yrhi", "uosldj", "hryi", "fwres", "ycygw", "ycvxze", "zevxyc", "iyk"},
//	[]string{"yphev", "xisbai", "xdb", "hzrbg", "ayxbhdx", "qnvbus", "pwc"},
//	[]string{"wytqraw", "yvurj", "erqckl", "rvrvda", "xsh", "gsd", "bxtm", "acxle", "gpndk"},
//	[]string{"kpvjtu", "vacggk", "oabcuoq", "qusf", "zitqpgn", "pbyq", "ocabouq", "ntpgizq", "gaiiry", "dke"},
//	[]string{"frz", "ceozajo", "ljltawq", "tjepzp", "iflv"},
//	[]string{"zntejm", "dkfgc", "ysno", "noys", "sony", "muy"},
//	[]string{"qdnyvvw", "oykq", "bnmldt", "zjgauw", "pviurd", "cbcnl", "tnkhq", "ebsey", "bccln", "arvwe"},
//	[]string{"iqazt", "xidjgom", "wcrdz", "itht", "lfh", "ifzwkj", "rwqxhy", "ervcgmt", "vojg", "lzntz", "ezg"},
//	[]string{"tlcxioh", "qvvkan", "wpi", "ody"},
//	[]string{"mhsy", "dwm", "hyms", "yegvvc"},
//	[]string{"hbkvi", "wvemc", "uwxgqf", "pwrgu", "wcy", "wxqfgu", "qkzppc", "vxcwdwd", "rcgp", "ivjd", "wmevc"},
//	[]string{"upc", "ucp", "cpu", "unmr", "pyod"},
//	[]string{"bqfew", "ebwfq", "paccwh", "phgc", "fchhr", "jrnio"},
//	[]string{"abh", "bjalls", "bdtac", "zzvt", "totdlc", "yptqmgu", "rpcin", "bch", "cqklqly"},
//	[]string{"bsnss", "qcsgi", "tegyz", "lqoqkpf", "qvatlyu", "ibza", "mzfotsk", "lye", "oqqf", "mnor"},
//	[]string{"lzrxca", "stkbn", "axhr", "wckbip", "bsntk", "ahrx"},
//	[]string{"oricdw", "cnpte", "dbz", "nqx", "xloxc", "bdz", "fdsl", "uyvgi", "nvoy", "ezbi"},
//	[]string{"nlcqye", "ofta", "jcvqvtg", "yxduyh", "btawc", "tjgvqvc", "tcvqjvg"},
//	[]string{"nji", "znctdfm", "kkmp", "pmt", "ikhl", "jjoubc", "xnp", "zdctnmf", "covvmsh", "ydh", "ircplcm"},
//	[]string{"yafkuk", "yasqsyt", "ytqayss", "nusgb", "ukfyka"},
//	[]string{"yogcf", "emg", "jlkd", "blupemf", "axg", "wihhrb", "ysernt", "yznhz"},
//	[]string{"gmc", "yyqtgp", "use", "lohoit"},
//	[]string{"lclwa", "ojjkm", "rxr", "rrx"},
//	[]string{"punyfv", "iphw", "ddbc", "jghx", "lrssszc", "bepexv", "sicpy", "syicp", "lszrscs", "vrqjb"},
//	[]string{"rjanra", "juh", "hljdmad", "scu", "usc", "baifu", "ijs", "suc", "bgdbbv"},
//	[]string{"ogs", "hehi", "lgiwowc", "wwezznc", "ogs", "lorcl", "naunpll", "wlmal", "ajcbj", "ciujw"},
//	[]string{"slenm", "xxod", "vhf", "amhts"},
//	[]string{"mrrduda", "mrrduda", "lwecwjv", "lwecwjv", "vvzmjla", "cjipv"},
//	[]string{"ixnv", "invx", "inmzz", "aoxghpv"},
//	[]string{"ulyvfrf", "zsup", "zfryjy", "xoo", "agdsd", "giw", "papa", "ljtu", "rzbjiq", "wrex"},
//	[]string{"bzpv", "svpuyov", "mklrct", "uzox"},
//	[]string{"fxs", "iypd", "uaqch", "kxazj", "ksjyv"},
//	[]string{"uxaommf", "xtq", "uubbfk", "bmlq", "kdhgjzg", "oxwpag", "itfij", "irmkjx", "ggod", "sddcyo", "bopn"},
//	[]string{"lch", "plmvfni", "qbjr", "dnu", "zjcod", "qlwax", "gejmyj", "gxjqm", "mfzkb", "gejmyj"},
//	[]string{"yoa", "thrfbto", "wposvrm", "amulogu", "mcqzfax", "fjquli", "oay"},
//	[]string{"kywj", "kqqhney", "ozuljxz", "wqwfte", "ubo", "mjj", "anhhg", "aphy", "ocfnef", "yhin", "ywnx"},
//	[]string{"vxuledm", "wllj", "obqtsrr", "jwll", "uvmelxd", "xvj", "gbwte"},
//	[]string{"hevc", "bitwu", "ydw", "ywd", "btiwu", "iuether", "gfe"},
//	[]string{"dzn", "ssmfpel", "wbbdeyt", "xge", "hrfi"},
//	[]string{"zebz", "ifhq", "euto", "ppfnrwc", "winkkj"},
//	[]string{"utuly", "wtdt", "iahpe", "ihtxwmh", "zxun", "bqluj", "hsaxgcs", "ytluu", "jlfnnuv", "drxlctr", "myhp"},
//	[]string{"kwxgy", "hreul", "rsnh", "wdmsx", "kkajywb"},
//	[]string{"bond", "djq", "kccazc", "zvzcie", "hndm", "myx", "cmhyhkc", "ove", "ord", "dqj"},
//	[]string{"zcong", "tekgn", "pbzs", "ywqgqgu", "eizrx", "ypydeku", "yqyjdjp", "dwsu", "zxire", "zcgon", "iggnvf"},
//	[]string{"tkioh", "hflkm", "fsjz", "gisjbi", "otikh"},
//	[]string{"ccdqqm", "fdabbne", "fyd", "lbyqm", "cyzgtd", "puitvjz", "nluf", "hirrpxd", "tgxrg", "vvl"},
//	[]string{"hjnygbz", "fnu", "twwbp", "xqw", "pfdlt", "uoalyme", "rsd", "muayeol", "ezcq"},
//	[]string{"kubeooi", "bxgwoun", "paogjs", "twvwlji", "opsajg", "higbdfi", "aazo", "evmj"},
//	[]string{"sfipxe", "mqbkmrn", "plwqd", "zvq", "nmvom", "fyfbs", "nbs", "qkrbmmn", "eym", "kqnrmbm"},
//	[]string{"ruhsp", "hurps", "mqws", "umm", "sphru"},
//	[]string{"ksjs", "pgpxh", "qozg", "enplxbn", "oqzg", "rvjnaje", "sjsk"},
//	[]string{"rbwbvog", "mhgtgg", "uld", "twrqz", "rbf", "kpop"},
//	[]string{"lwho", "lohw", "ylhd", "dej", "lydh", "vsfffsm"},
//	[]string{"icljgu", "gluijc", "vthqx", "orynv", "xhvqt"},
//	[]string{"biav", "elxkct", "mtaw", "nlafk", "snyr", "cbqdwim", "blim", "rtrqmc", "wze", "cxktel"},
//	[]string{"fgpeia", "ebkhga", "azlfsr", "bsj", "pipvwsd", "nry", "bayrjzl", "ftth", "ynr", "mfhd"},
//	[]string{"ymlowur", "nighqgk", "yjv", "pyxf", "nan", "xamb", "ohm", "jvy", "owrxbg", "icbz"},
//	[]string{"iyyt", "linaqu", "httt", "zyfeo", "udap", "mllq", "pdxo", "lpl", "djhqaou", "zkit", "llp"},
//	[]string{"dxspk", "yge", "kcqjqpz", "ulb", "hoe", "mfx", "nwayo"},
//	[]string{"rdnmmh", "gyqd", "qhxrzj", "dgizu", "lyittbv", "txngpdg", "fiu", "mwd", "ndp", "oks", "vxnxybi"},
//	[]string{"eul", "ztpe", "evnz", "yxx", "iuwon", "rkbbsw", "liy", "mqhxt"},
//	[]string{"qahp", "zwn", "ualtk", "txbt", "cbthj", "xchqy", "pirucp", "povdwq"},
//	[]string{"mqwu", "mwww", "muiafa", "miaafu", "hzripx", "wmww"},
//	[]string{"auw", "hyyi", "ylewfi", "ihva", "jknbrry", "voxzooq", "xaos", "xymv", "qzzjw", "hjc"},
//	[]string{"enpb", "jqa", "ajciy", "cbeopfs", "tqrlqj", "ecudb", "gso", "cyjai", "gxoet"},
//	[]string{"yohkjj", "yohjjk", "xapawgo", "rtgnjj"},
//	[]string{"lnlxxn", "nxllnx", "vhjrev", "uoew", "zts", "smkd", "kynlrg"},
//	[]string{"bkblpr", "vgafcy", "alju", "aiyqe", "eebtsyu", "buve", "hdesodl", "pspbohw"},
//	[]string{"aacmw", "qpndwo", "tcwsfqy", "qecnms", "wondpq", "sto"},
//	[]string{"wdsyxe", "edsxyw", "jnzruiw", "pfqdrhi"},
//	[]string{"pfgxekl", "vswgxhb", "qyn", "mykn", "jimiatq", "zkcz", "jimiatq", "kaexgxm", "mykn"},
//	[]string{"xegwn", "upudt", "dew", "uqjrcl", "abyof", "hbtiko", "wxgne", "sorgno", "etm"},
//	[]string{"xzojs", "zxsjo", "szjox", "gumjtwi"},
//	[]string{"gttngkk", "bwdgce", "bhuw", "fgo", "rcbf", "byw"},
//	[]string{"ngtzwqx", "ptx", "xodmy", "ctmtf", "oioahmm", "qajlhje", "jzilpk", "cvypp", "ijaefz"},
//	[]string{"arlx", "slcugvm", "hyuo", "zoptsh", "emyr", "tndx", "rroecp", "tdnx", "xea", "rtkpd"},
//	[]string{"sfckdx", "ktyrjju", "ruwjtp", "zhqznj", "vncun"},
//	[]string{"oqqh", "xpc", "itrdg", "gtrid", "hoqq", "tuo", "yijh", "ncp", "suvck", "jic"},
//	[]string{"brrlqu", "twdt", "urblrq", "twtd"},
//	[]string{"brfuh", "arwtkpu", "mzlj", "wdyqk"},
//	[]string{"pmag", "dtwnva", "nect", "azld", "axqrwy", "apgm", "xbv", "gdq", "ienubsy", "andvwt"},
//	[]string{"hqb", "qbh", "gxdl", "mwjn", "cmfsmik"},
//	[]string{"yiwma", "utlor", "qxjfjsn", "aomlvu", "gxp", "ryj", "rfkdsw", "kuguhyi", "qxxpg"},
//	[]string{"ifq", "wcvcgly", "jdalgzx", "lgcycwv", "rajmnqw"},
//	[]string{"latakk", "yxxbw", "evy", "vey"},
//	[]string{"odkvw", "ojgveb", "jhg", "qwhkyoh", "btvu", "qbfot", "nouv", "owgtsi", "pdwbmfn", "pmvcv", "dnqbo"},
//	[]string{"tmka", "qqnty", "knz", "swi", "fyvmt", "ikcdu", "jfjzsfu", "dshgi", "cosacuj", "szjjuff"},
//	[]string{"eanet", "uua", "fbztk", "bzkft"},
//	[]string{"jepi", "hyo", "jgzplr", "jyqk", "zzcs", "iepj", "hfamvu", "bfgbz", "sjsnf", "lprgzj"},
//	[]string{"mlca", "ywko", "mtg", "vkfv", "ojka", "zbelq", "qkaujs", "simt", "kafq", "qtycfzo"},
//	[]string{"sqh", "omv", "llvra", "ecvxmtx", "dpnafv", "llvszx", "xzlsvl", "quj", "ufnhvod", "faraf", "fivmnl"},
//	[]string{"pvxlls", "fzeoav", "ahgv", "uhq", "nodcr", "cohy", "vqisgaj", "jsfcyur", "dbohh"},
//	[]string{"ztccbwk", "okv", "vok", "kov", "ywel"},
//	[]string{"xyu", "cmaikc", "jgqu", "ozccdzk", "ybn", "yoeq", "fky", "aetrkj", "eyoyvla", "laiu", "cermo"},
//	[]string{"sssnb", "kxly", "mgvaz", "jpffkq", "bysbwwu", "rfnkm", "eycp", "ipyd", "hyi", "wjew"},
//	[]string{"obdfqmh", "flkm", "semednj", "iafewg", "lvh", "uwa", "ciepe"},
//	[]string{"zmldp", "ykffe", "gtehz", "qlmvule", "edrtzg", "prund", "oagwto", "qia", "bvyxur"},
//	[]string{"kjok", "neno", "qbxh", "wqgkkt", "ympclso", "poyclsm", "cajgnnn"},
//	[]string{"paqili", "kur", "sfit", "jbqchzx", "bhjqzxc"},
//	[]string{"fhghm", "ubtaana", "qbn", "autnaab", "aaaunbt", "vmz"},
//	[]string{"exlrl", "hfnpq", "zgdwx", "smix", "nyg", "ogccrhj", "iimhhwc", "uhcldo", "oydwxp", "kqc", "jxxpycv"},
//	[]string{"wtdqjfh", "ialoqr", "zeej", "ipoh", "qtjdwhf", "wdhqftj"},
//	[]string{"jcgug", "cmtvmu", "ynhnilj", "txlv", "uemowyu", "cvrool", "oolcvr", "njr", "cxqntdh"},
//	[]string{"uhtwtp", "tgnc", "jmmjl", "utiu", "jfxtsoz", "cxwqcz"},
//	[]string{"qntxl", "lyownp", "tsp", "tps", "mixyge", "rqfqumc", "bxjiry", "zmaj", "azjm"},
//	[]string{"abt", "bat", "tftvm", "nyjs", "jyns"},
//	[]string{"hzsdh", "pwthfvm", "cedg", "hzsdh", "rsxtehn", "owh", "cedg"},
//	[]string{"hcoiny", "aqbeme", "eeuigt", "pocpvox", "tiugee", "rwb", "tvgmyc", "ojg", "hgdaf"},
//	[]string{"mzlwcfc", "uoobo", "bouoo", "tvgvmiu", "evsfkm", "popgm", "evmfsk", "ehxvits", "vebxbmd", "qhmz", "jzj"},
//	[]string{"mypgg", "jbpx", "vgeb", "ahvjl", "zbm", "ancdzfy", "wytkcq"},
//	[]string{"bivscw", "zmzmjxu", "jzm", "fwb", "ujefxp", "jzsiskp", "cgx", "atcj", "sszikjp", "cxg", "nqvxga"},
//	[]string{"vvurbxp", "iry", "zlz", "gfnlpuy", "npyugfl"},
//	[]string{"fpmee", "mhj", "iul", "lui", "liu"},
//	[]string{"xjoesn", "ggsdc", "vnisnmw", "kyxmmv", "xphfq"},
//	[]string{"zcsml", "ngzlpy", "udaoab", "eotbv", "ylca", "bfmums", "izx"},
//	[]string{"pdi", "bpyoep", "cofsy", "qazl", "oaovek", "fvfbe", "sotc", "lfdmaea", "smvs"},
//	[]string{"zajm", "bskaqhj", "qxyiptb", "bdyeuqr", "dqjrekn", "iywj"},
//	[]string{"hzhky", "hukvpve", "iqcbwju", "nyiwb", "rvutxlb", "hyuah", "urnhxg", "savicaw", "hexr", "ptedk"},
//	[]string{"qndji", "wrr", "sin", "ljogf", "ncrb", "zvt", "tvz"},
//	[]string{"kvfke", "tjpzhrl", "zvd", "doq", "kxyw", "fdgr", "oqd", "egkybdh", "rqpfxks", "nja"},
//	[]string{"escstpv", "ccc", "ryzdv", "gxkjuyt", "gkhw", "jxnfda", "awpzg", "csestpv"},
//	[]string{"cpcd", "onxeae", "nimbrpt", "zyji", "qnuo", "ktxgwbj", "vtjfglz", "skcozd", "zgii", "zgii", "nimbrpt"},
//	[]string{"lwq", "iue", "hfbv", "hgbg", "aeqc"},
//	[]string{"vzgbod", "yjkoc", "ckt", "bpiaz"},
//	[]string{"eyco", "ecoy", "uzousjp", "faxj", "utu", "yoec"},
//	[]string{"fhqdi", "myd", "tvex", "bzizkcx", "pifcfhz", "fczhpif", "eofzv", "bqzd", "knbhbgj", "dok", "ffcizhp"},
//	[]string{"qlqlgmz", "hofmnc", "cwtk", "ahgnpne", "acn", "prwdh", "hwdrp", "rfofhl", "atavrf", "afkcbk"},
//	[]string{"sgl", "apyfr", "pwxzptv", "osuwy", "vmqqh", "soyuw", "lqilked", "oinhh"},
//	[]string{"eglqdox", "gcxfxql", "ejtnwu", "wvho", "can", "eyu", "uetwnj", "elgdxqo", "atvpkk", "eailsnn", "cwosyn"},
//	[]string{"mylxhuj", "kbc", "apnllw", "qbmtj", "sqy", "hxtnvoe", "ins", "iyudo", "aagezrq", "nsi", "ikvn"},
//	[]string{"lpmzo", "tkdeg", "zilkm", "vdkmtf", "yulbdd", "dkfmtv"},
//	[]string{"fzrv", "grq", "zfvr", "ychga", "gqr"},
//	[]string{"vdjxx", "wew", "pdxgp", "cjywsc", "meoffrj", "pgpdx", "chxmw", "eizgz", "ruea"},
//	[]string{"iaklyhx", "btqqik", "tbiqqk", "ynmq", "laxykhi", "qatrnsh", "lnmtm", "plz"},
//	[]string{"sfogup", "jtdsx", "tsxjd", "wwzkyy", "wzywky", "vgdsgr"},
//	[]string{"paupqb", "kyy", "bccklmr", "vopviup", "lctcxza", "yyk", "yky"},
//	[]string{"gduuia", "doek", "hqcr", "upvb", "szeewnu", "rrrdz"},
//	[]string{"lhnsaf", "lut", "kzf", "allu", "dvj", "tyngx", "zkf", "aqsgz", "rtkzzdz"},
//	[]string{"xxqj", "msg", "xxqj", "ezmm", "tmyji", "msg", "cco", "tain", "ctpel"},
//	[]string{"pvcfzv", "rhn", "hlhxyu", "bghzzpx", "dlhvr", "hrvdl", "lhuxhy"},
//	[]string{"npzhkp", "voghdv", "rvezqff", "hvgvdo", "jndf", "gpa", "wevrwpu"},
//	[]string{"faixq", "aecax", "hxdouer", "yqlngzd", "grf", "wid", "iwd", "cce", "xnmmr"},
//	[]string{"ifqwiah", "dib", "ibd", "dtvkwqj", "mpn", "dtwjkqv", "kyntq", "xwlv"},
//	[]string{"rwoiz", "dja", "cvv", "lvza", "kfdblq", "bgtwgh", "ongdub", "wggthb", "lvaz"},
//	[]string{"xajf", "eyasx", "rupsyqx", "wubjwx", "bsrqi", "ripghci", "sbzxp", "sbz", "dhooax"},
//	[]string{"ydnv", "tvhwgp", "uvrh", "yryhl", "yxdlwa", "ladwxy", "awi", "mkwyn", "ghvpwt"},
//	[]string{"qrl", "vwdvwic", "ocbhpi", "bcmz", "dor", "lrq", "hokg", "gokh"},
//	[]string{"adz", "echnlju", "ebnmw", "twjl", "cfw", "loq", "fqklyyq", "clgqq", "jtgpsu", "wltj"},
//	[]string{"vwlgisb", "murtsw", "ldkacqv", "wxfcke", "vcqkald", "ldhh", "gsl", "kpzn"},
//	[]string{"itnvo", "lyddd", "saewfse", "spnri", "vtmst", "iblx"},
//	[]string{"qsgv", "qni", "wvqiih", "mneg", "lkpb", "quhbkyi"},
//	[]string{"efwaaa", "huu", "fslzwpc", "uuh", "szflwpc"},
//	[]string{"sgmj", "ajh", "vcwpcua", "enltaks", "aquooh", "bwoda", "txbuve"},
//	[]string{"vbe", "astgezx", "xqbxkdj", "evb", "bev", "yuuesdc", "fvohzq"},
//	[]string{"gpn", "oqxfz", "pbwibjw", "gljdbf", "gbldfj", "sis", "dpk", "iss"},
//	[]string{"pha", "ebybvye", "ntxhs", "wcuce"},
//	[]string{"odnnywv", "qpcjll", "aslxqjm", "injfs", "vkbturz", "atxi"},
//	[]string{"zzncfj", "kbhk", "jzzvnwf", "kqipx", "qkxpi", "rzb", "czfnzj"},
//	[]string{"ygu", "vqpnxkw", "trdtv", "rrte"},
//	[]string{"hrutley", "ljxuuq", "yonbpmk", "hmpc", "etyrhlu"},
//	[]string{"odxp", "fpvizzx", "dxop", "jjbr", "skhxq", "mpzawhe", "zptdxuu", "erxk", "adbbuk", "zfzipvx"},
//	[]string{"qjovasi", "yutjpg", "rcp", "bykpctm", "fqmmg", "pckbymt", "hqj"},
//	[]string{"ssqc", "cype", "tkowxb", "fbh", "rsluu", "zjk", "scrukwv", "pkuexk", "qlgjtdq", "oulrke"},
//	[]string{"bkcd", "nnf", "hdj", "sdlweyr", "uyf", "kmvzq"},
//	[]string{"sgeg", "moy", "png", "blv", "bblree", "ufe", "uqknuqd", "lnjwbh"},
//	[]string{"snpex", "hrbcfok", "pffv", "cwrvhcs", "fpk", "uprhn", "gbpy", "zkxyi", "esug", "ccnnj"},
//	[]string{"bmwue", "dugcrdu", "uifiy", "clid", "rdmodec", "jodp", "hti", "xptj", "tpopl", "vuwhdyi", "hnoq"},
//	[]string{"cwlkg", "qsz", "nnp", "lfyk", "pwn", "dpe", "oeuzp", "jusxxkq", "qlnchc"},
//	[]string{"tsmkvge", "pxauyc", "cxypua", "boi", "hybq", "rzf", "iioyi", "rtedkc", "gjmk", "iob", "mqb"},
//	[]string{"cvip", "wgbjhe", "ulwg", "jckkwd", "gdu", "bmaoisp"},
//	[]string{"drpl", "xbliszf", "rpld", "ngnvgxl", "xnrd", "xsmd", "oetrcmn", "xvfohp", "mtocren"},
//	[]string{"habmf", "dmfxq", "qitw", "xxtybla", "cxvb", "colcvpj", "sowoeuh", "bhmfa", "plccvjo", "naftjgw"},
//	[]string{"cix", "soo", "icx", "ahx", "cdrjyxe", "htcnp"},
//	[]string{"acoksaf", "cgahlg", "tdj", "euchwnj", "bdn", "lunouq", "aewrk", "uktre", "kklwqy", "lnouuq"},
//	[]string{"ibsel", "hwjbah", "vxuk", "bjxa", "dvzbpq", "tffqvo", "bjax", "qfoftv"},
//	[]string{"iynms", "tzv", "jakuuw", "cmz", "cjnyr", "ddibtd", "gcb"},
//	[]string{"tmgerk", "pvwizc", "lxoma", "ergmtk", "xmg", "loaxm"},
//	[]string{"ajazon", "yjwt", "viyemnk", "wobzwwm", "jcucn", "nopymyq", "znaajo", "qcjtmlq", "ccjun", "ywvt", "oqczod"},
//	[]string{"kdhgnv", "kdnvgh", "rpyrxx", "xpyrxr"},
//	[]string{"qftmshx", "hrbr", "kcggxw", "jwtuk", "qxbzkn"},
//	[]string{"ddi", "fjekwxs", "xxua", "cmmkrso"},
//	[]string{"ptrsv", "favvfh", "innnnx", "nninnx"},
//	[]string{"kzsnd", "pnf", "evtazgw", "wmjk", "gvxp", "bslajo"},
//	[]string{"nphqtka", "umlxu", "ymw", "whqiab", "whqiab", "vwigkz", "pevtps"},
//	[]string{"vhje", "cnu", "uzfzum", "rwucy", "mxr", "wyrdqfi", "gnkuwz", "dkrwc", "otfc", "vbfc"},
//	[]string{"ubtzio", "vlijsst", "anruj", "qyntadb", "fnic", "klz", "ywqq", "fnic", "vlijsst"},
//	[]string{"rprj", "ybyqawb", "tgeieih", "nzcr", "rjpr", "bjfpozh", "tpevsx", "flvjdq"},
//	[]string{"kvqqzvm", "cfdm", "wzjmkq", "sbcfx", "vzmkvqq"},
//	[]string{"zqtt", "elpg", "eglp", "uoe", "glep"},
//	[]string{"lqv", "madhtch", "xevl", "fal", "ijmx", "chcpra", "lzl", "afl", "cndbvnq"},
//	[]string{"yjx", "jyx", "xjy", "otwklfj"},
//	[]string{"cur", "egsdzaz", "ocbx", "wvky", "coxb", "pgiysbh", "lclhvy", "gfu", "oxbc", "vqyjvhh"},
//	[]string{"gtd", "pytdaz", "kecsku", "nkiud", "ytt", "bmgobx", "tyt", "pfleji", "ebqlifv", "lqp", "ytzadp"},
//	[]string{"bopfdvy", "eovszvw", "krgu", "jhuhyqi", "kaczafr", "krgu", "zlfxtl"},
//	[]string{"yixprs", "zqai", "oprgw", "vcsjoc", "pgorw", "ypx", "ijo", "urjcjqv"},
//	[]string{"estg", "oqnhw", "xgwajp", "mpbsag", "ibzi"},
//	[]string{"zggbt", "jmmtkru", "sye", "ybd", "ztggb"},
//	[]string{"tzryuqb", "blyxnnu", "sjpmf", "yxe", "zimf", "uyzqtbr", "qbyrtzu"},
//	[]string{"rzivz", "rxn", "invxqd", "nazw", "efzun", "bwnw", "ywx", "rfuda", "jhnoww", "mketav"},
//	[]string{"zxfw", "zcaqi", "qaciz", "ktefiwk", "xwzf"},
//	[]string{"ntougl", "fvyaxfr", "obml", "obml", "bjkm", "lgsqj", "yfcggdu", "rqcpgt", "ntougl", "nveto"},
//	[]string{"rma", "dakifg", "pvfc", "ticvff", "iffctv", "difkga"},
//	[]string{"wpnt", "eclov", "vmmoqh", "qvw", "mljlvnj", "hxrx"},
//	[]string{"ijnpo", "uhgxrfe", "xxopw", "xuzwyd", "powlpo", "ctduj", "eepw", "gubnepv"},
//	[]string{"rxcmve", "myxckxk", "ecid", "nxe", "xevrmc", "juzaonr", "ilkx", "zpb", "pbz", "mvgos", "yzr"},
//	[]string{"yfecm", "wqkh", "defqqa", "mnzgx", "nwe", "ixxg", "rjdhe", "new"},
//	[]string{"awztgx", "vqgnfd", "iwcakr", "ajaiwn", "jiwnaa", "uqfrim", "wrgbjon", "ufqrmi", "vdu", "yjwy", "gwkdc"},
//	[]string{"vrqf", "yzmvnr", "jkjji", "ghya", "pdlye", "ygha", "qlcm", "twmkex", "frqv"},
//	[]string{"hjb", "xgypw", "vtr", "mgj", "asa", "owcuks", "fnllp", "ommrhng", "senv", "iix"},
//	[]string{"usw", "iyuatv", "ondiwh", "neac", "ttge", "tzw", "bvjkfe", "neac", "usw"},
//	[]string{"qwehwhj", "idrwo", "vex", "zopkjd", "lrcc", "sfqyz", "smte", "qrfh", "lccr", "qwjhewh", "vlb"},
//	[]string{"efnlhsj", "ltm", "xirn", "nctwio", "cin"},
//	[]string{"zocc", "cpfo", "czoc", "oiz", "tftk"},
//	[]string{"rlzvqe", "inivone", "kptyumn", "eatw", "urjxc", "aetw"},
//	[]string{"qavvqa", "jvvc", "yux", "cvvj"},
//	[]string{"bfr", "fcpc", "xpkphcf", "irak", "bfr", "nuhsooj", "nniugf", "bfr", "gllq", "ipo"},
//	[]string{"ekd", "dydxs", "rnmgb", "dek", "yowk"},
//	[]string{"ipdll", "wdgx", "gjiigd", "uleiwg", "buvv", "vdhuzg", "gigidj", "gkyigmx", "lpdli", "lzyode", "fqdpvms"},
//	[]string{"ycna", "rhyz", "bsipz", "lit", "rmc", "mrofb", "cyan", "mrc", "wujk"},
//	[]string{"tjrk", "cwdsvf", "srkdjy", "lsyvryj", "nko", "syjvlry", "fgqq", "srdykj", "pgb", "koh", "dyle"},
//	[]string{"sflcxt", "wmgdgej", "akllaoa", "bbsvve", "nsxnt", "nsxnt", "kgm", "akllaoa", "btqbez"},
//	[]string{"bzmoci", "agemx", "mdtiol", "pyohvf", "zwtx", "aqblx", "oibmcz", "ypcmz", "lfg", "ckssn", "ozx"},
//	[]string{"cuojke", "joekcu", "eusr", "dxqk", "xxwob", "klpsm"},
//	[]string{"byspz", "lyunt", "eojrx", "ubh", "rjxoe", "ypzsb"},
//	[]string{"ibok", "bkrtut", "wzcdk", "ppm", "qekhvy", "aupaic", "vswwul", "lmlxrv", "ainigy", "sasurbx"},
//	[]string{"jeigcyc", "cycgjie", "resio", "ncz"},
//	[]string{"xvxr", "lmlaje", "ebmtn", "cvms", "xrvx", "vsmc"},
//	[]string{"cfjbffj", "xvo", "hpsbu", "nfm", "jhlsk", "mnf", "fmn"},
//	[]string{"xncxo", "iwuon", "voj", "aebv", "jks", "nynzl", "hwjwo", "womejo", "ugzmr", "tdfaep", "olpdiaf"},
//	[]string{"jlnusc", "hgpprf", "nulcjs", "kwiwypu", "kitjjbj", "snlcju"},
//	[]string{"buqzm", "kpuoxel", "ajlqke", "qqdur", "jecuibn", "leajqk", "qudrq", "usi"},
//	[]string{"ahbnjf", "uuzecdv", "yfyrsxv", "eoqey", "oonue", "vyyrxfs", "jixmvao"},
//	[]string{"wjdi", "cfgurdl", "usdnlk", "jmao", "qnus", "cicxnux", "vtdxxkx", "nusq"},
//	[]string{"mlvfz", "twu", "unj", "mar", "qpiz", "fhjczpz", "ytl", "npwjys", "ppq", "koa"},
//	[]string{"ippdky", "pvwthzj", "qlkhl", "pvwthzj"},
//	[]string{"kfm", "lcedomr", "xgdkrzo", "hfxyoe", "rafcby", "uwe", "pzjyuja", "weu", "nso", "erdwc", "fjvc"},
//	[]string{"peep", "oufzlb", "fsgn", "kxj", "uge", "xvufb", "zsnrxss", "lere", "gfsn", "gvwajkj", "fmh"},
//	[]string{"mexvi", "kgkktz", "kgkktz", "auyln", "ghvqc", "mexvi"},
//	[]string{"wit", "qxtewrk", "qdaz", "oljeb", "wowb", "suergyt", "hxq", "pfnfbei", "rdu", "qrxkwte", "fyw"},
//	[]string{"qjjzkd", "oxedeu", "uoxbehs", "zder", "vvjnn", "ocxkiz", "wkblzy", "eyzksc", "waiiqo", "fut", "raak"},
//	[]string{"dhojys", "qkusuxs", "kzicui", "dcsxo"},
//	[]string{"hsnexb", "yoz", "inza", "gqxtbc", "rya", "jqfe", "ufzlqva"},
//	[]string{"scpquf", "gbloz", "ceol", "eclo", "qpvzavo", "rwfnxa"},
//	[]string{"jyg", "edqf", "vdpsihl", "edqf"},
//	[]string{"mbyjg", "yjgbm", "mgybj", "mhgi", "grw"},
//	[]string{"ler", "oxssrel", "jhw", "jwh", "sfa", "hdhlo", "gng", "auzoan"},
//	[]string{"hmkuis", "iaxf", "muhisk", "ikshum"},
//	[]string{"eodbpo", "prajasi", "zsu", "hyb", "npr", "tbcntup", "uzs", "bxchne"},
//	[]string{"zpyr", "kxmvz", "frlzwnb", "tzqrw", "vdh", "ndbwqmu", "cadwsv", "adq", "bzfnrwl", "qfgf"},
//	[]string{"dybnn", "dmarc", "mknr", "fovokgj", "ukrp", "cuwag"},
//	[]string{"khweq", "eljs", "rllijp", "pizevm", "lwws", "kehqw", "mkjcu", "otqodz"},
//	[]string{"fvsrb", "kzbjbcc", "kzbjbcc", "mee", "dhyedb", "kzbjbcc"},
//	[]string{"cnlmjd", "dvnd", "vhlvsa", "rsrtc", "scrrt", "tqhx", "vke", "jqmxpd", "udkjqc", "qxriw", "pfqpk"},
//	[]string{"tyureg", "urteyg", "rutyge", "rmoihs"},
//	[]string{"pccxeak", "tkmyane", "qqggpr", "tbdmpi", "ieb"},
//	[]string{"wjucbi", "wjm", "hais", "pcxjd", "kkzh", "qns", "hmf", "mhf", "mnsv", "ifigsc"},
//	[]string{"lqeyr", "pnopcig", "cpgoinp", "pncpigo", "mjfkjus", "cko", "zedvvyq"},
//	[]string{"ofsnspv", "goj", "wqm", "ynolb", "qczly", "brl", "lrupzg", "buof", "zladwte"},
//	[]string{"xzn", "zxn", "yaseulw", "qwhxb", "easyluw", "vvlmh"},
//	[]string{"aiybip", "ydfczwh", "wkl", "rjsu", "xreokl", "qov", "mko", "pna", "fkfu"},
//	[]string{"zjlwozs", "nxke", "ozwlzjs", "jybx", "jpsxp", "qtkll", "idsrad", "savpoe"},
//	[]string{"xph", "lpvkmvy", "afq", "uqhg", "qqjgm", "smg", "tmhem", "mxdyqx", "bvhot", "lpvmkyv"},
//	[]string{"jxkkzil", "pkjheow", "fewr", "ggbfy", "fuol", "cuzud", "wnx", "fxujfwh", "srjsmic"},
//	[]string{"lzbjx", "vfx", "sncis", "xuv", "unoa", "nlgs", "stdhf", "oibxsgk", "uhogsb"},
//	[]string{"hfqzkms", "bzyfnz", "npcd", "yhfdo", "myqukug", "pjq", "adle", "sqkfhmz"},
//	[]string{"czhainb", "iazcnhb", "hhaqr", "cyrwj", "zzow", "luuvt", "zdyhnh", "uppysr"},
//	[]string{"fyw", "dulbxa", "drewqsr", "tldlaac", "kyaw", "datclal", "ylplup", "hdzbj"},
//	[]string{"kiiv", "tly", "gua", "lfg"},
//	[]string{"gphbvwc", "lqdd", "jqja", "ffpkz", "hafal", "eiapksw", "wsikpea", "vphgbcw", "lkcpm", "zjxcx"},
//	[]string{"viapp", "rxt", "vdgbm", "ezphp", "pcqr", "bll", "mklgx", "epzhp"},
//	[]string{"favz", "bwmczgg", "zoyns", "pens", "wpgi", "mrwxel"},
//	[]string{"ozwjjn", "kbzaozc", "cuaa", "igbfyq", "swi", "uypx", "bczaozk", "pyux", "odvawqx"},
//	[]string{"npnpw", "nwppn", "egnpj", "fkzh", "wppnn"},
//	[]string{"asu", "mlqmwa", "npewa", "cjluw", "qmalmw", "newpa", "fznx", "dzypi", "yiy"},
//	[]string{"hleh", "usw", "bgmgscg", "cqc", "fijfuw", "idtyh", "cgmsbgg", "zjhr", "wus", "hymbju"},
//	[]string{"tmre", "fvm", "cgowgb", "eduyfla", "ttez", "vdzp", "vtmtaog", "ezxsfi", "gyxgzi", "pvzd"},
//	[]string{"acvarlu", "hkmfzdg", "jsnil", "hpv", "wjj", "rljpk", "pygl", "wjhhohk", "rkplj", "spvgx", "psgvx"},
//	[]string{"wyz", "rvuobq", "kbmhaf", "bec", "bec"},
//	[]string{"zosyz", "psuo", "lgihdo", "mjtftk", "fjkttm", "ddmcd"},
//	[]string{"pqm", "qpswpb", "opviwxg", "ppqsbw", "waco", "jpx"},
//	[]string{"yzgumgq", "kqv", "hqjghnl", "jixhoyg", "ufer", "kvq", "lzi", "rojm", "gbhvsd", "urd", "tuy"},
//	[]string{"sxc", "jndqc", "ozpex", "wkchpu", "tmwv", "utcxzd", "piecpma", "cmppeia"},
//	[]string{"ifjc", "lttj", "tltj", "rxmgxqa", "jcif", "lzhxkg", "zqb", "mdq", "kbjavv"},
//	[]string{"isyxn", "zjbj", "uiw", "avziqxf", "zpezazx", "iuw"},
//	[]string{"rjaudu", "amtpuk", "gufogys", "xiqs"},
//	[]string{"gau", "sndrkv", "cmiolti", "cdxm", "ikkcisu", "xusnfbp", "jxffy", "ffcizj", "psye", "sgd"},
//	[]string{"mvx", "onzmy", "oynzm", "mwfgvs"},
//	[]string{"mrdg", "roanty", "dljs", "jlil", "gzcj", "dqitnfb", "gxb", "vzzqf", "ooweeh", "pgs", "oyntra"},
//	[]string{"yna", "xgok", "fvbdl", "xgko", "udqp", "sorfo", "hmhl", "yan"},
//	[]string{"kycl", "ule", "blupejp", "kycl", "fhpkoe", "pqkptw", "cfzpv", "nkncl"},
//	[]string{"snugkyw", "zfdbsfs", "aehb", "olq", "kkoi", "xpsvy", "jqcspf", "lajjyu", "jtm"},
//	[]string{"hifhfa", "mii", "clukcbc", "fhhifa", "wcts", "tgai", "vvqsyr", "kclcbcu"},
//	[]string{"ordjftj", "dokk", "hdhytwc", "fjodrtj", "ojrjfdt", "san", "ajxrwy", "ked", "jfrqc"},
//	[]string{"eylx", "cohd", "biswq", "xgiibz", "gzcptbf", "eylx", "icunv", "bfg", "jqanbv", "rntp", "sbfkiey"},
//	[]string{"kub", "gdpbdms", "qnnto", "bit", "visqop"},
//	[]string{"tywk", "clicj", "daics", "cbuewkx", "yyjjjka", "vxzk", "afsdyqg"},
//	[]string{"bmxzll", "wqjnyr", "mxlzbl", "yutkaa", "qmpz", "hiqkf", "lqqqle", "jagj", "qlqelq"},
//	[]string{"jgdeatg", "qubj", "jsu", "bhgbm", "swmgy", "lwgnuh", "qjbu", "dqwiikp", "mgwys"},
//	[]string{"ryisldg", "asrrhz", "vxvrnk", "ibjr", "kebyx", "dwbx", "qdrpa", "tgakt"},
//	[]string{"dfvgzk", "hifan", "dpjdnqc", "ncnvf", "xmqjjao", "npjq", "vobt", "evtaety", "kvufds", "pcugx", "oyjo"},
//	[]string{"ezionjg", "ioznegj", "cykxy", "igeojzn", "ezm"},
//	[]string{"dkv", "dkv", "vfqyl", "dkv", "dtjhrem"},
//	[]string{"xfgh", "brsjcdw", "auvq", "fibb", "gcbecl"},
//	[]string{"eet", "qdnrymr", "ndqmyrr", "tpfqxoi", "kbkxw"},
//	[]string{"qhmaj", "maukf", "uygg", "hqmaj", "tfmtv", "irao", "wsari"},
//	[]string{"ofoywus", "wxs", "leemkn", "wfko", "dwzqv", "txg", "qsiiiss", "aiiffe", "fadmdq", "zjtaovt"},
//	[]string{"fgfms", "oazi", "sam", "msgff", "qif", "evxca", "reho"},
//	[]string{"gpzhy", "qxh", "sco", "qeax", "wtabxwv", "sjd", "oev"},
//	[]string{"xsmpi", "wskvku", "xspmi", "smipx"},
//	[]string{"ghgf", "tbpeun", "qdivuvq", "dump", "umdp", "khxcxtx", "laljpv", "lownp", "egovve"},
//	[]string{"vhmu", "eziabt", "hnz", "neko", "nkz", "hfmizn"},
//	[]string{"vqhb", "lax", "zzyf", "lxa", "lik", "jrqynr", "rgcos"},
//	[]string{"zjgbfzv", "ieufyz", "kjxad", "qxeuewx"},
//	[]string{"ufl", "drkaac", "hoyic", "pqutop", "wqzdk", "eewabsr", "mqspcr", "ewarbse", "dzqkw"},
//	[]string{"bgatanj", "xmddwv", "mwlmw", "scgzboo", "myignm", "lkfl", "fsqr"},
//	[]string{"xkrx", "otjzfk", "wek", "dpbwk", "cromg", "fclmhg", "pkvw", "wsln"},
//	[]string{"yyqjs", "lifg", "gifl", "cfv", "lfig", "fycza"},
//	[]string{"dfup", "fkfeiqq", "rcmuv", "dly", "iforzi", "lngkjc", "rzifio", "oiifrz", "mlhor", "puwm", "qrthoa"},
//	[]string{"nzfaeto", "punt", "rtmlg", "dwdk", "hyig"},
//	[]string{"mds", "ueoyclh", "lxb", "axgea", "wqt", "wwqqglf", "tqw", "yvzji", "ryr", "dst", "bojf"},
//	[]string{"ayoj", "yzj", "lyctgnc", "woxz", "gqltw", "lkzkwte", "wysb", "mjyeu", "hrwso"},
//	[]string{"gymmvtt", "lhskza", "lsb", "nhlijnt", "men", "zphurrw", "oftksy", "zxs", "ykerwue", "hnijltn", "iknqyz"},
//	[]string{"xuaxkc", "lgzeef", "fwb", "nlzzhjj", "lsieg", "qdr", "ews", "rue", "rdq"},
//	[]string{"xnf", "lljcmod", "kyuercp", "kvlvd", "lkvh", "ncn", "afaq"},
//	[]string{"bjytofa", "ltz", "mkyy", "bwt", "uxca", "somiz", "rhgdg", "keaqu", "ybr"},
//	[]string{"aipljn", "qpq", "nilajp", "udgkchc", "dirvxg", "rrbmi", "mxfxkk", "kmfxkx"},
//	[]string{"sfzjemk", "hjsvnmb", "hfd", "hprfmvg", "pbhkc"},
//	[]string{"cvxi", "srj", "ucy", "yuc", "euswuns", "jpox"},
//	[]string{"tajlnn", "ivuecv", "fdfce", "rakjq", "bfuxirh", "eibde", "tajnln", "nlajtn"},
//	[]string{"ndvm", "mlnhy", "bfqlzn", "nmdv", "ohto"},
//	[]string{"jysyvwy", "xbcyt", "lbbm", "osoye", "swwtwa", "emfznci", "qnzc", "qsgk"},
//	[]string{"xcm", "jbqsuo", "xmc", "mtrk", "bojuqs"},
//	[]string{"ukshrrh", "xhfl", "ooxgq", "vadlcrg", "ydva", "hugplg", "mnqbd", "wkyouq"},
//	[]string{"mnmqys", "bhws", "megar", "krgoke", "modxe", "krgoke", "clovh", "dlo"},
//	[]string{"ejl", "qzc", "agxph", "jcn", "zcq", "zqc"},
//	[]string{"jgh", "yhh", "hjg", "jhg"},
//	[]string{"tarm", "jboyg", "gbyjo", "pgalg", "xugzt", "bph", "mart"},
//	[]string{"yur", "wrrahr", "fnnfqu", "rwhrar", "cdq"},
//	[]string{"mukod", "gueg", "guge", "epalg", "xjkctt"},
//	[]string{"hub", "hbu", "hbzul", "buh", "sqfl"},
//	[]string{"xyrly", "lvpitr", "xfzn", "jjcl", "uvcnz", "dnpdyzq", "ifaiwe", "zlvzcx"},
//	[]string{"wxzsf", "tgd", "khvwp", "cmd", "mzv", "bsvjvjm", "wvhpk", "ublnqyz", "mvbjvjs", "peishcb"},
//	[]string{"zunmk", "hxpney", "nphxey", "znmku"},
//	[]string{"bfxlgur", "wftcw", "xfkf", "fsik", "xkff", "ffxk"},
//	[]string{"sxyjzr", "ugcscx", "uiovqx", "ldzhord", "xgxbfph", "ldzhord", "prdhg", "rhdhzd", "ugcscx"},
//	[]string{"udg", "drb", "apyjq", "dgyevo", "fuxjhg"},
//	[]string{"qshbe", "aigfdp", "wyvz", "xfcos", "wve", "dgfrufw", "dwimmb", "jfh", "wfrjbzk", "nwgrigr", "sbrpbb"},
//	[]string{"ahpn", "xnzeof", "wxbv", "chxpcu", "jmso", "age", "ojsm", "bqonfki", "hqhrkw"},
//	[]string{"mfupm", "vvig", "ndqbbm", "jlw"},
//	[]string{"ejqh", "ebcrytj", "zpiqtpp", "ogznd"},
//	[]string{"wkwkae", "odq", "rsrnqk", "nslczz", "hiyyhur", "kuw", "mjbuwll", "vduvod", "ryhuhiy", "swo", "tsos"},
//	[]string{"znkufyx", "jejrdno", "knr", "wkx", "ikrlly", "tnxtj"},
//	[]string{"iizdiw", "iizdiw", "hukep", "rwj", "eaq", "ptm", "klta", "rwj", "onaz"},
//	[]string{"znb", "egqy", "qqnc", "igqr", "ikza", "ojgzkr", "xaen", "kurb", "pyckxvt", "wqx"},
//	[]string{"pbohpw", "bphowp", "dajwdpp", "kotevs"},
//	[]string{"hmuvxu", "zdnguk", "jhcmj", "gccyxiu", "cxgiycu", "uyxcgic", "akxi", "demeff"},
//	[]string{"zjr", "lupzwcy", "puva", "rzj"},
//	[]string{"cdn", "wee", "iqkbhl", "jwxo", "nhl", "cqd", "mqgqf", "ltdfg"},
//	[]string{"phwco", "ggcj", "cggj", "ergpqmc", "kcz"},
//	[]string{"aryjl", "wqwmkc", "aklktpz", "kptnroc", "mckqww"},
//	[]string{"jadydt", "atjdyd", "tajdyd", "owswsgm"},
//	[]string{"dshqt", "kacoge", "sdqth", "ylobih"},
//	[]string{"kdnik", "knkdi", "dinkk", "xwvqa", "gvii"},
//	[]string{"cifbkpt", "zye", "xhwnrhm", "fctibpk", "sbn", "pdqry", "emkye", "kzyjpa", "plzkc", "btkfcip", "gcchi"},
//	[]string{"kekfr", "fufp", "dfy", "eqebgko", "obbn", "nuh"},
//	[]string{"zixfbus", "skuf", "bea", "gimbqq", "caubhto", "eba", "uvkovz", "xisfzub", "peukmyn"},
//	[]string{"okihcgh", "gazrc", "slee", "vlnwyg", "geviex", "eesl", "nmnvk", "rcbv", "ycupyw"},
//	[]string{"tcvlgqs", "wxe", "lusvwzy", "unr", "yzluwvs", "wsylvzu", "zkwth", "qdykv"},
//	[]string{"hyenkj", "ugao", "vlwgb", "putcepr", "lyeccs", "fqdotx", "burf", "aqew", "fje", "rfbu"},
//	[]string{"uhmnc", "zgnkarz", "gylqawm", "abl", "zimcz", "qbs", "zzmic"},
//	[]string{"pxkbpn", "zuxlwtt", "rlbhegv", "zlxuwtt", "ooxpr", "pgjx"},
//	[]string{"leg", "wavgps", "fcplfpc", "xvxih", "ueskmi", "dvu", "wbiq", "nvtia", "pwjojw", "usiemk", "ojwwjp"},
//	[]string{"zmrpknx", "xrfpq", "avque", "tvoyqp"},
//	[]string{"lyposyj", "mckyoub", "sqbl", "olpsjyy", "hjafzi", "wmojb", "nvezofd"},
//	[]string{"yflxrg", "egi", "aij", "qvc", "yflxrg", "typbs", "yflxrg", "kliexy", "eqnj", "jqrr"},
//	[]string{"gggt", "sht", "kdajvz", "sht", "gkqwaot", "sht", "vout"},
//	[]string{"ahl", "aucpih", "feig", "man", "umtchcv", "ceqabr", "tfptb"},
//	[]string{"ftlywun", "voaorf", "kde", "ilwt", "hlpoe", "pksqxyh", "vpg", "cxo", "xgq", "fdkkl", "sgxhnq"},
//	[]string{"zzekhfi", "akb", "lupta", "sgtd", "qapznzf", "lgidsx", "lidsgx", "akgmq", "ettuwjq", "xyumf"},
//	[]string{"dxhpku", "lwoxpim", "gwb", "lhjmoh", "gxqapd", "ntmvc", "rvwwszg", "pvin", "lwoxpim", "coubc"},
//	[]string{"qia", "bxmremo", "rjf", "vaeio", "eqexwz", "wuoz", "sguf", "bsbusof", "xqeewz"},
//	[]string{"iczzz", "krf", "hbq", "tsgrih", "mzjabrt", "sfnwrm", "djignf", "zwac", "cwaz", "dgc", "nsrfmw"},
//	[]string{"yvarsva", "zzpbp", "yai", "und", "kkbinr", "zlyj", "nyxxof", "ggrgu", "vyk", "eib"},
//	[]string{"nepzm", "yrrgr", "vrlhbv", "hykmiog", "natrqx", "jvpq", "nbyhe", "zuo", "grx", "nwl"},
//	[]string{"cfboeev", "hcn", "yfobyx", "cqvlo", "obctww", "xxaux", "ofybxy", "wouguq", "avuztl", "xmgqq", "xyofby"},
//	[]string{"tikv", "uvzp", "oxopqy", "reh", "uzvp", "wefka", "vli", "kied", "gngks", "vbz", "thfsxyt"},
//	[]string{"exxvknp", "pucbdyl", "dboto", "tzat", "qze", "xyinygz", "mhzl", "ubahr", "ekxbtk"},
//	[]string{"jcz", "ufszbi", "pknsfgb", "ivok", "ijau", "okxolj", "etecn", "aurun", "zsa", "gbxs", "uryx"},
//	[]string{"ypnb", "ousd", "osg", "mvset", "ipffzdn", "dfinfpz", "ltescx"},
//	[]string{"taeoct", "aoetct", "aocett", "ttda", "fcdqnxv"},
//	[]string{"bimtlk", "ssax", "bmrifkr", "vfxdmq", "hglp", "rgzr", "zpvk", "zhxtq", "rndwy", "mmr", "arkr"},
//	[]string{"bwvdb", "axxbhzk", "nwfmbbu", "kzuc", "sahv", "cvas", "wdac", "acsv"},
//	[]string{"xavkwou", "yvx", "ouwkxva", "otbe", "uzr", "mmw", "atq", "yiy", "ghavd", "qta", "pqlhv"},
//	[]string{"omzht", "vsdsc", "zhtmo", "hmotz"},
//	[]string{"eqt", "wtveez", "syift", "shtfnc", "hmckjxa", "apwbvn", "yme", "okdl", "hbihdtv", "hxahns", "eetvwz"},
//	[]string{"rokdg", "ndjw", "hprxjc", "viys", "mbcctod", "dbvd"},
//	[]string{"lhzb", "fyxf", "xaslmi", "sjd", "vqp", "grxhmfe", "snetfv", "mgivd", "uaknj", "givkdi"},
//	[]string{"gxkxl", "kqcdnl", "rna", "jhcuepd", "npiedg", "djcpheu", "huce", "njryw", "bjluhq", "bvedvl", "kqxu"},
//	[]string{"sogh", "uym", "atpzuwx", "vjgbe", "xgrvkg", "thgbyn", "mptcebt", "rkro"},
//	[]string{"tnpxw", "uxrqxd", "lajmsmr", "tnnlt", "vrvbf", "deret", "hkmvrs", "eubvkn", "kks", "hjq"},
//	[]string{"rcdoa", "zfja", "vod", "supip", "dvo"},
//	[]string{"zbxdo", "xglqv", "how", "mgoq", "jqrdou", "pwrminc", "lidi", "nfs", "xglqv", "lidi"},
//	[]string{"ldmnp", "dnqn", "ahhr", "tha", "mvsbsfh", "rpm", "rgus", "faf", "tjash"},
//	[]string{"ewrdol", "jqhfpf", "rckj", "mrxlwj", "redjg", "bmxprx", "grla"},
//	[]string{"rhr", "jksebwa", "vtu", "skwaejb", "vut"},
//	[]string{"wrx", "iqvrjh", "atrt", "xrw", "vtqo", "tkcasd", "xedjh", "zkqrh", "vvhj"},
//	[]string{"owc", "qlzygar", "uajwwe", "obzl", "inxawur"},
//	[]string{"crbtrf", "phvy", "nzipo", "rctbfr", "trrcbf"},
//	[]string{"vwuun", "wcfhhzo", "vxxjdt", "fbf", "bqtmmhs", "bffqcna"},
//	[]string{"wkxfxmv", "zmrkyh", "sggw", "whwcw", "zukynw"},
//	[]string{"lsdiy", "lnbn", "kblxi", "qfyib", "irfl", "mymrr", "zqbl"},
//	[]string{"gwdkeu", "ghn", "afry", "zxoz", "fary", "uzntlnk", "kee", "xtnop", "ptnox", "zngoran"},
//	[]string{"lgs", "lsg", "sgeseiz", "gsl"},
//	[]string{"erpoqpi", "svtnv", "vsogl", "uym", "amzxbs"},
//	[]string{"jtmodqx", "yjcsfcl", "zosugm", "jrekfdw", "xxbdqnx", "fcha"},
//	[]string{"vio", "tlfxokx", "xaoq", "pydeiq", "glxsuzm", "honifvf", "maiuzsy", "uizsyam", "eco"},
//	[]string{"ophcui", "saka", "qyt", "ltti", "syw"},
//	[]string{"qff", "qff", "sde", "ryv"},
//	[]string{"eiii", "jazx", "nlehtx", "tnhvxl", "rzvsjo", "qkupif", "feypppe", "tefxr", "wdjmlc"},
//	[]string{"pdrr", "mwuy", "wccd", "rxla", "drpr", "enbbap"},
//	[]string{"pclml", "ubwdbz", "hfudj", "gdpujfm", "ovabv"},
//	[]string{"uink", "ffebi", "wdvhqzs", "qiympf", "lqxihty", "vnsp", "wdvhqzs", "hutxkcs", "lxfuos", "hutxkcs"},
//	[]string{"fycoaw", "palkpz", "yrkg", "kappzl", "ncjym", "mergg", "kryg"},
//	[]string{"eqy", "npvgh", "ghafkro", "piqnogb", "polacs", "qye", "hnvpg"},
//	[]string{"dxyy", "udhmz", "jij", "tqsuic", "qxz", "erctv"},
//	[]string{"urum", "nmbr", "cgek", "eetmhj", "gxr", "oxgukf", "wzdmvi", "oibzt", "fxkoug", "rcrywcr", "rglx"},
//	[]string{"jkp", "ofej", "waibl", "opqhmww", "ifnczcg", "jdtkbc", "lil", "isc", "ill", "mylvuv"},
//	[]string{"vqbcosk", "yhhsy", "gasmj", "bspx", "peakt", "cjtekw", "hvzo", "ywe", "qcvbosk", "ohzv", "qddt"},
//	[]string{"edq", "llbvsx", "vedyvlm", "gou", "wkecank", "rkgf", "ziyrr", "belgo", "tbz"},
//	[]string{"wbome", "vhzf", "ztk", "zaxiu", "ywjb", "supuf", "beq", "sxq", "spuuf", "pufus"},
//	[]string{"femu", "ymkdoew", "kjynct", "aia"},
//	[]string{"yjymr", "orovqj", "aremii", "licw", "bdtnc"},
//	[]string{"uyade", "fbx", "duaye", "ujtvpn"},
//	[]string{"yzvp", "pvzgjp", "yofcvya", "gvkkoh", "cafyvoy", "mhsm", "okhkvg"},
//	[]string{"xuh", "qkaf", "dmi", "imd", "tzmlce", "mqkxj", "qilrc", "dim", "cadotvy"},
//	[]string{"azpqgb", "kyc", "aflgyaf", "laagffy", "kesmk", "jzyzaer", "taf", "bpkbzdg"},
//	[]string{"ogd", "dbdlh", "dqt", "zaaloh"},
//	[]string{"exal", "vgnfx", "omu", "omepvwf", "szcwq", "snz", "bptite", "bzqyxl", "khmblyc", "sse", "emg"},
//	[]string{"yqcbwsn", "aihhf", "tqck", "tcqk", "wqwqy", "cfez", "xahpn"},
//	[]string{"qqbuf", "lil", "ies", "tqu", "pyxhqp", "mnfuk", "azj"},
//	[]string{"vwma", "rzdtgl", "mxbasw", "nwgjav", "mwav"},
//	[]string{"itpjfq", "rrgyt", "hralwm", "fqrig", "btwcod"},
//	[]string{"ydjd", "kmk", "fvwr", "wrfv", "yvhw", "mkk"},
//	[]string{"xbsxub", "yhsj", "xzbuf", "ace", "xubbsx", "fzuxb", "vxk"},
//	[]string{"ttsist", "vubpf", "mhwkmtx", "vlj", "hdsva", "kmmhtwx", "ukxr", "upfvb", "tbma", "fxsrnxl", "hzwufho"},
//	[]string{"wckjvz", "unmtev", "egxts", "ihw", "topvw", "ptowv", "rnihhmq"},
//	[]string{"gpdtl", "kcric", "nwg", "ssbs", "qah", "aarp", "ydsdty", "ngw"},
//	[]string{"lzhxbbq", "oktvcw", "xbasqe", "owtmwgp", "koa", "gumjie", "sodwrp", "hqsw", "aqh", "dtgsbb"},
//	[]string{"xjbyy", "mxfxa", "ogvk", "nqiy", "qyni", "ldqwryj", "niyq", "jjixc"},
//	[]string{"uhbul", "daccgva", "xtiz", "dim", "uhbul", "yjmakv", "yjmakv"},
//	[]string{"huo", "esajup", "ouj", "oju", "ujo"},
//	[]string{"eeeu", "hwvsk", "jfkmds", "okhi", "pogskfm", "itdlbll"},
//	[]string{"lpyubo", "dylpfb", "iehwug", "decj", "ntidy", "cuygyg", "lalkb", "iutu", "oxgm", "imn"},
//}
