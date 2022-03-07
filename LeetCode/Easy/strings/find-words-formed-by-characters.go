package main

/*
URL: https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/
Status: Success
Runtime: 40 ms, faster than 36.17% of Go online submissions for Find Words That Can Be Formed by Characters.
Memory Usage: 7.1 MB, less than 38.30% of Go online submissions for Find Words That Can Be Formed by Characters.
*/
import (
	"fmt"
)

func countCharacters(words []string, chars string) int {
	totalCount := 0
	charIdx := make(map[rune]int)
	charsLen := len(chars)
	for _, char := range chars {
		// get the count of each chars and store it in a map
		charIdx[char] = charIdx[char] + 1
	}
	for _, word := range words {
		if len(word) > charsLen {
			continue
		}
		// we cannot duplicate a map directly as the map is inmemeory and charIdx gets updated aswell
		// So create a new map again for the charIdx
		wordIdx := make(map[rune]int)
		for key, val := range charIdx {
			wordIdx[key] = val
		}
		// for each words loop thorugh the each chars and check the count of that char repetition
		completedWord := true
		for _, wordChar := range word {
			if _, ok := wordIdx[wordChar]; !ok {
				// if the char of word doesnot exist, break
				completedWord = false
				break
			} else {
				// if the char of the word is already exhausted, break
				if wordIdx[wordChar] <= 0 {
					completedWord = false
					break
				}
				wordIdx[wordChar] = wordIdx[wordChar] - 1
			}
		}
		if completedWord {
			totalCount = totalCount + len(word)
		}
	}
	// fmt.Println("charidx", charIdx)
	return totalCount
}

func main() {
	fmt.Println("input: <>, output:", countCharacters([]string{"hello", "world", "leetcode"}, "welldonehoneyr"))
	fmt.Println("input: <>, output:", countCharacters([]string{"qobxtxjzdngkrzamsxzdvbvkiwijokwdyndqllhqpaoxzwonriclsm", "fahtqqnuzhhhrcblquaosdfdcqoskxcsdnhtwvvvzsxkpjprytieieniafnltxmuzwkdnttofpibi", "xedhntgrqegfs", "wrssfvsbcehbahbvojnzgacbgveitirkjmmyiorwykynqddgydzgfvlvplfnyumgxturxraonpchd", "hpmdzhpgijirecxzkuyhptiytnuscu", "xljgysrjjukphjgzbpn", "gmwbirxt", "yhvpsvsnhfmbmcpihnqtodspbvexnpgcqrrughbakbufyjispkquvfppkaffypdpnvikjygdaarcigipfhuvzzzbgw", "gvvbgpjolobpbxcnhnzuqrsqgrkanwmnnkqynakkooailoafunopsrlijqhaqmszssxikftcfoqsw", "naktgxfyuvuoh", "muui", "ghzqskipqprrzeligdjyowypeerogxonvztsq", "onosezgutawtzteoagbftclsqpfsxtwetourxjxurdqevrir",
		"fcskqxwkzsldsjihahalnw", "lsstgzxjxabcbspjwelqmhtnurgfmdtpaihxnxad", "nwrwtwetgjwooevhxhkzlvsyghtkldioyjhkkniepktqs", "utohzbqvkqkq", "vdoqnhtidgicevprrrwlbtldtaxpsxdhxhgbwlkbeglkbhrujtafswjkozdmdpvwhsuskofmxxtprtpopacslinwozth", "zqlzxpemomnbzxlorvlkxt", "kubshvnldnkofitnnjere", "czxmqpowzzhdbhgtlqdokrzxowsvwyavfhcycctgdvrsfzmanrlktfaetnuayrqkvhcbxezfahkrxgaaztovrxuhnll", "rrcesnfbxglhjawldnnuiiepdsofbrsbjczlemusqwvenicxmtdmpwfrnizymfmqynvtkbrmablcugroshc", "thholqebekkysstqzlbbdapktxumygplqganucwnahmrihiraxdnvbiaqhykcti", "fagqnapzeeglbdzsbneosxmptmwopjcxztukhpjkqjmjbkpbzrfaqskctehyboeddmilkwlurcb", "wtjdykncubkduhxiwwusoxvzpnaxpnzdjmwddnonsmmvwmuaxghetgrwpoeqbprxgviwzagyqopfdakrqjgiy", "nsftehpgzstetbganbtozdopptseucdqkhzdmujnzjdvufqtikgeepttnrabb", "ozihktgnvljzhqwanxemtzxphzqvmoblvi", "iwdsrekqllbbyarzzmbqbvldvxctdeswiyahiwfoefhfsfwktdzaulnalewbusazjhcbtxjuck", "dylhds", "idnaddnzbodhjrpshhahnbbnrskruxszxeeyanumazvahktizectmmvjbhnhbrohsyqhrgq", "scyzsykrwzuozmbrbenfiqpxchtpmcxepjiglaeionsmbwrzeidupayusczcooudpcgkgspbuyzcdfymsejucb", "otpcfrhckxmnseayhwoyjjfkceaoznmmkikpdsuueyqmbsuelmhqnmdsjs", "xoghnhpypfiibqrpgtyux", "cxhpbcziptgandiwxtctdjpboiqwv", "gyjosuhwgbqpcdsdqadopdqozjxpzwxbqaffnxaddhgrxmunpidvpnijxnbawshcznkmprpgkhvzxmzbaftedgtfgjhaisdnva", "jlnxshfthqgzlnvjzztrnjswwsotpybxecyqhnfkbfwvmxpcs", "mtddspmqwbnoiajpexsuhxogarduzsoammqqelh", "sixkvxgnbalipwmkbcwpugpvolsvvlmaaeeobmoeogspbkbsskwjdqkite", "xhsutriuynfrldjbhexhxjgcfwcujvwsqfiaqwvjnkjkswdpaynelhryrzfeqjhajezmolk", "zfiredtxenzgtrkiamuoubetexzbnwkdtxbtihhtsfubnmryxq", "sot", "szgkkhuhkcnr", "firxkgvkzqlnallzwjispsizoawobemuhqrhpyvknasjzwctamfuonder", "vyei", "fidqszoicndwifns", "xkickycwzj", "gmybflbjunudxrwguzditaxmyadgdjeengut", "gvlwwnmrddhzwewugdtobauecealfhasyftgxkiqeluysxxmroukfziifpryvddggttojhcszeyjetbsldmorqnbuqreuxfw", "vbhuerxkcjlkamgruturkbrubbscmvzqruwvlrlyylcvuiothkhpznjzsfnyfoaqkziyjqzdreeygmtbdljwnaojexfgmjlup", "bxjvop", "aswdmctosieicucsjwxych", "lxbtlhnrfyaznrqgxrltmxkjmhdqjjgnvstxaygmuxznjfiiukm", "npncdabkmbgfhnedcmbfyjiplzwymgvsfrifvvjayobsygwolqbwkblqutakcshnlsqadtfiqmpauslvqd", "kuyemuuymacyvmahtagmcewkspverpjtjscccnrczchgkjzppdxcalxxcxrwnepk", "jxwouobfvzttz", "yucpsdhqvzboeezcqpxsepuuk", "iwbdmxdcbypnzqmgkrjhfivkrmnv", "lpqklgcwdvgbghfyoyejnpexrelywqfdtczflwetbxvzigtvisi", "dlasodatffcreungntdmezgfqklrabyymsnhdberufcrgpxgsziklznhdprbczhbxgtuslufycjronozdqumzlnidkvaydg", "oejzlmrrbdysqlezifcvgjnmvodfvmrnjmnfkejdbnnyljzjaxfecrfefdgarqbtwoijuktacywmsubtxtgzkbnstgrsrjpkdfe", "mlegvhxielwlfadlnqvnipcuizpdxgtvro", "vwgmwkbturocovaykdsjaftbtgmtwknnmhexfgcfchpwwgcgecfobbencotjizxbtdrijwfjxdsxanwfjyjamrxrdaiusgr", "xjmkcsekcorldyrjiavrhuhqtndujymc", "rmxwcaorznumwxgwnibnxwzvnxjhzwqzgmkgifnnnnzpgtsvprycjtdeirtpqbduursabbidzkfbexgthkoacagkb", "tkrsxhztwgvqxamjtexklnooaloydjhejlbasavskttwxiabarogvmfdfjttaxhgqljlbfnrvrwwbxkurhufiknoxfmemcv", "cojlyayladyrhofzetaddteqrjbyxtvyszgdorexqmgznqmuvemegbwki", "zktqnurtpttlcjgkmnprqeyeepqunfqqvjwuevwbvnaupyofwiqwhpyumyfwpjrruhleromrmpvczlkxqiuq", "gyxl", "qmhwlymfsjixvvjhkczylqcsnbjxliasetxciggtfl", "gizysqkqbyhzeagzsscvdigtcfiupyhqovaaioxfrphugxzrcjvwqwc", "hpgkulrmbaixnjiapmjiwhwsgromfqpxqkkiscjwpiicslwcijginxfiwqakeezeohhskxgvgdkezmqys", "vibswdyqaxjvqsffwmcismooheyhwzqvyzezluejztlgddshvwcryzcllaisxrygwqyyoiaivfvgtzicycyrkckv", "oyclwdejlifmocfjsbgmernsyitkfaahjxfnwnusrdypsziawlpsjgnavytdihpxcmugshnqbkyfts", "xwcbiporfbktpvqhznjuaxfcdykifuzwdsdhxirwwxinoffywgxscrtuwhvuwjejzqirsfijgguouapqpmfdnpsfm", "qffrjfkj", "rpynimubaxdgbrkdawgugnhobaowxdnzkiidworcsnejgqctftyksvbhiwkcnyffmsbxwptqn", "oeqdvldvfbklukstxwomapaauaozblhxudkdxihepqagndnlkvbqhluscvczhrsrhodnftoszhjdymuywdtjgzbmkrdf", "xviupppkeswkctwlqwyuhokbhqqjshmaeiouhununbhrkabacenkg", "habomjnlznqvckmowgelhnglfizogckplbymkdowfpicxydzgoskckraxpdysksvzezcpg", "zemawxwjeowraaqsqytsshtavnvoyiyollelxqabmjwhxtrdijiacbbjiiyzwkxboot", "jcnpsxnkbqdbh", "bftrbaurtzkftodotjguzjmwxrrmswxwclohotuanypmhtemmsaicwckowmcdmpnhcfzptekpgljfvwpqjgilxu", "bbaigjqxdmuchkkb", "effbptpwafzqbsebbjmvdcxdbmnlfqjklongafzkvaqmkehefedjxgxlbdhltihtgfqjjsdis", "lhuxgqpwcgpujfvvnlrxhgbiwxmxzhglyhkdkafafojtlkuxkmjwlxrd", "ussuyjqsxwsdhkjeipwycrkcxxjatoqhygzymgikqdnqiyyzlbcdlgtmneyickucbkpkza", "ykmsksjorovmdymlbgprvprnyxwppvwgmzfjsqouvgara", "wduqkbncayo", "qdfkyomjlhfozonwpdxllqdnvpohyijqmyymuclnydzmlqksntdfj", "ynzekkkljzrvnwclzcfgtvcmstxgadxpztofckypbgqgbnumnkeaqclaspzxjbygtkjznxeduhbksr", "aw", "jbremnahqoiqktpbkteefdkbrerrxmhqbbselpnfzapgmxidrhbixetaetjcroufa", "wtufuqgljvxzsurhviikdxyuythezjnvwrxqrykmotkhlphlyfljjsfugzwxxfqkc", "ytjuaagqnfxqwiopolnejmoxow", "oqoskpzkwxsffgeuuhdklidtmjobzayatyaqefgwgqbo", "qkhtpuhhkspffkpryvdjasbxhtfrmptpljszvtgvhfvqpxypxfdaphfqdmigzgfg",
		"goddugelwdvemxwureitzwqmbqeqtymrlrzahuxnpgufaagbpexwpoahvdnsyvgxw", "dfhxrctagxkuasofoejsalcerkbtsjwnbnoahsumzfyiiskhzzwyykeefszrzrbztbrzhxxgaajb", "tjsardsbhmhefysdqtbhzsxukbwdpioqaohloaancxdkkmofniojrvxj", "jzjozqfqiwyfadplibubtpcfxxfvjtbizxlmpaccjpihvnrtvfqtfiaztvfbednyemfmahbljdvykddawaujdksenm"},
		"fcxpzkzkqeyeijquyfixvlzjpzomujrqzxeoynjiflnmdrpdkrm"))
}
