package medium_0424

import "testing"

func TestCharacterReplacement(t *testing.T) {

	var crs = []struct {
		s        string
		k        int
		expected int
	}{
		{"ABAB", 2, 4},
		{"AABABBA", 1, 4},
		{"CCDAEEEDDEBCCDBACFCEDFAECCCCFDEADFECCDEBBEAFEBEEAFDEACBCECEDFBBDECCAAFCDDCABDBAEFABCBAEDCBFAAECEECAEABCBCCABAABECAFEAACEEBCFFFDDCDBCDCACCFFBDBEDBDEDFFECDBDBDABBBDAEACEDFBBAFBABFAEBADCEFDEAFCFDDFFCAFAEEAAEEECBFADCEFADAEEADCBACCDDEFCCCACEBCCEEDBBFFBCEDEFEACBDBBABDCBADDFEADCBEBBCFBCBAECECCCDDADCDBEAFCEADEBBFFACEAAFFDDCEFEEACDFDACFCEAFDCABBBBEBEBDDABEDCAEFCBFEFFABBCFAFCCBDECDBCABFABFABEECEDFFDFFDEBDBBBEEAEFFEDBADAABEEACFCCBEDADFCBDBDDBDCABDFCDECEFBFDFCDCFEBCCDDFABCBBCCCBCABEECBCBCBEADECAFFFCFAEFFFCDADBADFDFFDCBABFADFAAEFBADCCAFEEBFCFEBCFCEDACCADAAEEEBDCADABFBADBAECACBEFAECFFBCFBDEAECEBCFFDCEEEEEAFEBFFAEBBBDBBFFDBFADEECFDEBDBBCEEDBBADDBBDEDDAFEBCEDCAEBEEEDCFEDDACCDDCAEFDECBDCBBEFCFCECAABCFBBACACCEACDCBDBFECBDFFEBBFBABDECAECAFEEDAABEAFBECAACCACEEAADCEFEEBFFCCBDBCCBADDAAAADBFAECFBFACFDECFACDCCBBCADAEBEDDEEDCBDAFDDAAAFECDADBDAACCBDFAEEACFBBDEDDCCEAAEFEBFCAAEDDCFBDEFECBEFACCCDEDAAABEDAECDEEFCDACCBCDFBCADBCFAACCBBDFFBFBCACEABACAFEDCECCBDFAEDCBDCBAAFEDBCCBBFAFAEABFBEEBDCDFECDDDBDDCCBEFCCEDCECFCBEFFFBFBFBFDDCBEDADECCFFAFAEBEAABECFECABDABFBCEDFCCCCADEBCFBECCDFAAFEFEADACDEFFCBACAACFFECAAFDEDECCDFEDADBDAECDFFECAFFFEDEBEDBBBBBABDAEFFFDDDDCCEEEADEBFEAEFDACBFDBFDADEEACECFADACEEDBCECBCEFCBFBDEADAFCDACCFAEFEFAFDFECAABACFFFDCCBBCCDFECEDECDFDBBFACCDFBADFABBFADDFFCAAABEBDFAAEDAACFCCAEECCAFCEEBEEBCCFFEFFDEFBEEFBBFECFADAFFEBAAFBCFBAFBACDFDAEDBEBEEBBFDEDCEFCBBBABCDAEBEFAEBEAFDECABBADBFCEFFFDBAFDBBEBBCAACBFDDCEFEEEDDADBCFCEAACDBFBFDEACDCDFABDACBCAEEDECCDBDBCBACEBBBBADCBFCDDBDFECEECFDFFBDCBAAEBFCFCDFDBBBFCDECBADAFCBFECABACBFDBAFDCDBBBECCDFDFBECEDBCCAEFDAFACEBFDBDFAFCFFACABCAAECCBBEEBDAFAEFFFEBDECBEAECEAECACDBCFEAADBFEADCCEFAFCAADAEEDFBAFACBAFCEACAADDECFBFDBFCCFAABDABFBEEDFDBAFFECBDDCDCBDEBEFDCCBDBFECBFEAFDCDECAFEAAEAABDFDBBFECAECFEBADCFEEDCBDBFDBBECECFBDDBEFAACEABEFCAAFBDACDDFAEEFDAACBEAEAFDAFEECEFECFBABABBCCBFECBBCAFDCEDCBCAEBFBBCCCBBAEFEFBEEDEFCABEFEECCFBAADEFCEFDDDEDDDBBDABEEFECAEEBCCDDFBAEDCBCCEFFBECFBDECAFEEDBECEBACEBFCEFBECEDFABFAADEBACECDEEBDDFAAEFCCFBFCCAEAAEFAEDEEFBADFCDCFAFEDBDADCFFBDBBEDFEDEFDFCAAAFBADDBDCCCCACDDBBFBECCACDAACCBBDFEEBAACAEDEBBDCDEBDADCAADEFFDDAEADDDFBDBAACFDCEEAAAFDEBBFFEBACADCACDEECCCABFDDCAADEAFABCCCEEFEDDDCBEEFFEDFBBAECDEAAADAECEAACBCBBFCFBEEAFFFDEFDDCBDDDAFAFDEBCEEEACACACCBBBCBADDCECBDEDAAADEFDDDEAEBBBEAEFEEBEAABCFCBCAFDCFDDCFDFAABEFCAFEFADBCDDEADEFAFDAFDECBCECEADFCDAECBDBFEAEAACEBFFFAFEEDABBEBCCBAFACEAFFFABCAEAFEDFCBFEAAEDCCBDBFEFFFCDEAEBEBCDCDEEFDBEEAFBBAEEEAABBEDBFACFCACFABAEBAFADBBEBCEBCFECBFDABBAAEDCDECEAADEDBBFAEECCADFCBAFAAFEDAEFDACBDEAABFFAACEFDFCFFCEFACFBDDABCACACDACACDEBACDCDCADEADFBFACEAABCFDCDDAAEFCFEAFCFFAFDCFDEDAEFCEDAFDFFFBECCFEAADEFAACFDEFBEEBCDADDFCDBADBADFFCCCFBBBDDDEEADDADEFDDECFFACCDFCCBBFACFDFBBEFBCEDACACBAECDECBCBCCFEBEEAEECDCDBAADEABFFCFECBBEFACDBFEBAAFBBDFDEEFCCCFEBBBFECCEBDBADBFCDFCDBADCEACCBDCBAAFEEDDFEAEDFEEEEBCAFCDDBCBDDDCABCBFDCBEFCBCEBFDFBCCDBEAABEAFCACDEBEDCAADEDDFFBFEDAFEFACDDCAECEAEBBECBAACADDCBEAFFBCFAFBBAFCDDCEBCAEFEFBABFCBFDDAEFDBFCDFDACFAAAFCECEAEDDBFFBBBEBAAAAADABEDEEBCECCBCBCEEBECDCFDCCDEADFBDACAFFAADCBDAEEABEDBBADAEFCBBAEFEDBBFEAADFAFBABADBBDCDECCDAEBFACDDEECEEBEFFBCEAEDEBDEACCFAEDDDFADCCFDFACFECEFDDCCEECBFDABDFDCBEDDABEFFCDDEDCEBEFAEBFCFBEACCDABDDACBEFBCFCEEEEFEBABBEFCAEEECEBBDBEBDDABEBCCABCFFCDBBAEBBFAABCEBADCCEEADBBEBCCDBAAAFBFBCFFACCFFFDDBFAAEBACACFEBEDACAABAEECFCAABCDABBCEDADFFAABBDDBFDCFEEFADBCCBDAAFCFCBEAECEBFFEDDFCBCCFDEEAABCDFEBFEACDAFCCDCFBCDDAEDFBAAFBCDFFBAAFBCEFAFCBEAEECEDDBBDFFDAFBBFFECAFDAAFDCFDCBBBFCEDCFADFAEEDFEFDCCACACAACCBAEDFDDBEECFDEBFCACBDACBAFFCCEDBBEECBACFCDBEEBAADBCABCFFAFFCBABDBEBCDDEEDEBBDECAEEEBFFFEDFDFECEAFAADDAEAAAACBBDFFACAFFEEFAECDDFDBCBCDEFDFFDDACEEBBFDBBEBAEFDACEEEACAFACCBDFEABEDCADFCFFEEBFFBFAEAFBDEDABEACADEACDDADCCFBCDBFCBDCADBFBBBFECCCEBFDEDEEDBCDCBEEBCEFFEADDFDCABAAEAFCAEACCBDFCEECDCFCABCCDDAFDFACCCFBDCFFFDDBAFCECCCBFCAECABCBFEEDADCFFCFBDAEEDBCABBCFFDBCECBAABEEFFDDAFDEBAAFBDBDEEBACFAEBEECFAEDDFDBBFEDACBADECEBCFCCFCCBEEEBCCFFFEFFAEDBECAFBFBFDDFEBDCECAFFCCBAECCAAFEDDDAFAABADBAFEDAFFEEAFFFBEFBEADEDDFAFEBDFAEBCDDFBCFEEAFDDAAFAAEACBCEACBCBEECFAAEFBBBEFDBBDDCEBBABFEFCABEBDAFDDDFDCCFFFECDFACCAABCDABDABFFBACBAADDCEBBBDFEEFAAADEEEFEDDDFCBCDFDFEFCEEFCBCAEFFCEDAADBDEFEABBEEDBFBECDCACAACCAAAFCEFBBCCEADBFBCFCCCEFABBDADBACACDABBADAFDECFBDABDEBEBFDFCBADAEDFCFFCEDEBBADFDBCFFBFAFDCFEADFFAEDFEBDCCEAABDBDCADABACCAAFEDDFCADAFDFADDCEEAFAAACAFCAAAFEBDBEBACCABCBFCCBCEDADAEAFDAEFDEABCACCCECFDFFCAEBCECAEEDFFEDAEFEBCAAECCFBDDACABCBCEBBADCDDECDBBEFAFCDAFBEDCEDBAEBCADEDFFEECEFDDEFCBBBBEDBBBBCBAEFDDEBDBCCEAAAFADEFACFBABCCFDCBBADCBAEEEEBAEAFCEEFBCFCBFEFDDCBDCDCEAFEAADEEADAEFBBEDBBCECABFCCEBADECACDAECEBECBAAFCBEABCBDFECDECFDBAFDBCBDFEDBEACDBDEBEABFBBEDDBAFBACCDEABDBEDEEFEDDDCDACEBEDBBBFBEFFBCAFBAFFAEAFACCFDFDFCBCDAABDFBBECBFFAABCFCFCAEEBFCCDFFEFEADBCDDDFCEDBFEABDDFBCBEABECBCCACCCCABAAFCABCAFFCCCFCBBEFEEBEECEEBDDEBABAFFEAAADACDFDBEAABFAECADAEBFFBECBEBCEEEBFEAECEEFAFAEFCDECDECFEBFFBAAFADCAAABDBCDABCAAAFBAEDEFFEDABCDBBECBCCDEFBDACFFBFCEAEFBDCADAEEDCBFEDAEEDBBFFFCFEDFDDAADDCAECBBECDDDFBAFEABFDBBBEEDBFAFAADFDCFCBEAFAEBBDDFDFFCBBCFCCBCCEBAFBAFDEAFCAFFCBDDABDFDBAFDEADCBDFCAACAEEFDEABAEDAFBBBFDACBDABFBFBEAEEFEFCCDCECAECECEABBEFABDFAFDABBDFEFFCBDFCFFEEBBFEDFCADDBEADFFDDACCDCDEBADDABBDADAFCACAEDFDFFDCABAABCCEBDEAFBDCBECFFCAAABBECDBFDFBDEDCACDEAADCDFEAEBBBADFECECEAEBEBFDCABEDDBEECEEDFCEDDADAEBCEEDBAFDFACDEDDABDEBDDDAECBFEFCEEDBBBEBADFFCABDFABFBBFCABFFCCADFCBFBABCDFFDCFBACBCCBBDCDCABFDAFBCFAFBDAFAACBFACCCDBEDFADFEEEDCADEFFFDDCECADEFCDEDCCDDFEDFEBCBCABDDFEADFDFFECFFCCBCABEFBDBDAEDACBBFFBDFEEEACDEBDFFCDEFEBFBBEFABBDFDADDBACEEEBAEECEDACCCDDDFBCEEECCDAEDCCCDBCEBDDADAACCFAEEDFBAEECACFFDEBADDDAABBCDACEEBEAFCADEAFFDBAFFEDAFECCAAEBAFACACEDEBCFFFFDEDDEAFEBCEFDCBCBCBCCBADBAACFDFFBFDABDACCBFEEADFECBBCAFCFDADAACFDBAAAEEACCDAFFCFBCAFBDBAECCCFBDFBFCEBEABFEADAADAEAABEDDACBCCDBAFCCFEAAFBBBBFFCBADEBDECFCCADDCCFFBEDCAADDDFDCFBABBEBBFFDDDBCDBCDDAEDFDBCADEDDECBAAAFCDACDEAAEFEACEBCEEAEBEFCEDEDCAAAFBFCDCEFFBCFDCBAFEDACCBFFAFCBADAEFDBDFAABABDCDABDEAAFEBFFCFABBDECEDDBECEEACCDBEEAEBEECDBBBDEEADFDEFDADDBAACBCCDCCFCFEFEBCDABEFBABCDCBFDAACDDFCCFACEBDCFEABACFEBAFECEAABFFDBFFFDACEAECBFDECFDAFCECEBCBFFAEDBEEDFDEBBBCFFEBFFDDCDBDFCCBADCDDBDEACEEDDBECCEFDAFEAABCEDFAEAFCEFDAAECEFDEDDBACDBDFCAECBEFEDADBBBDFFADDCCADBFBAECBACDDBEFDCEDBCFDDFCAEDBCBBDCFBDACFBBEEECFCEBFCFDCAACBBEBEACFCCECCCEFEAEBEBADDACCFCFBCADEEBCBDBFAAFEFBFBBAFAACDDFFBDCECAABEAADAEDCECDEDADFBEACDBBEAACDAABEEDBDADCACFEEABBCBFEDDCDBFDCADCABBECBBCCBDBFEBEBFCCEDBFABDCFFFECFCEEFFBECCCAAFBCBDBDDADCFCEDBEEDAEBCABDFEDFACBFCBEABEFCDBDCFACBBACDFEBCCDBBBCFEADACAEDCCFEAECDAACAFFDFFABFBDCBAEBABDBCFABAAFEFECBAFBCEDCFEAFEDCBCCDDDCEBDFADCEEBABBFAEEACADCABCFAFBBEAFDFBDABBCECAEFEDACCCAEFFACCDADACCFAFEBABACFEFACCAACFFFDCDDDADCEEAABEEBBEFFFBFEECFEBCBFDEBDAAFCBFABFFCBDBDEADACEACBCDCEEFABAFBDDADDADDDAEBDDABBEAFDEDFDACCCFEBCACDDEFBBDACCBCACBDFFCFAABEFFDBFAFDADBDFEFABEFFFCBDEBFDBEECDACFEDCBFDBFEABEEEBADBADDFEAECBCEBAEABBBEEFDCEFBDEADCFBBAFEACBDEDAAEBDABBAAFCBDEFDEABCAFCDECABFDAFBEEDDECBCABBACFEBBACEEEFADEDECFBABECECEBCABAAFBFDADFDEDCBEEAEEAFBFFCCAACCEBCBEEBBABBDAEBDBBEFDEDCABEAAADFCDFDEDFBFDFFDEEEBEACBDDDFDCACAAABBDBFDCDEDFDDFBDDABEFFBCCDFFAFAACCEACEFAEFFCBBEECCBDEADCDBDADEACBFFDFDFDACFDABBEFFDFEDEDEEACDFEEBFCBBDADAEEFFFBFDACBFDBBBDDDDCDFECBDBFAFDEEEDBFDBDEFAFEFDACBEDDAAEBEDCFCABFDAADFEFBACCEBEECECDDBCCDECFEABABEEDDAEBEDCEBABEBBFAEEBCEDEFCDBEEBBCBCEEEEFFFFEDFABFBDECFFFDDBDFCCDBFFDDDEDEBAEFABFCDEDCCDEBBAECCAEEFFFEFCACACDEECBBBBEFBAEFAAFBBCDABEDEEECBDDCFEDAFCFFEAAFEEFBFDCBDDEDBEFECBDFEBFEBDFDEEBCAEBDFDFAEAACEADCCECDECCDFAAFBCEDFEFAEDCFFDAECDAAEBEFEFDCBDCCDCDADBEDECADBBCBEBBABEDEEDEDBBAABCCBBABADEADFCCFDBCACAAFAFCDDFDFFAAAFDFFBDDEFAFFFCAEEEFFACECEDDEDCEFACFEDBDDFABEBFCEBECFFACACAFFBDBCBDBFCEFBDFBEFAEDFCFBCDCCABFEABECFCACBEFFAEAAAEDBFAAEDECBBBEAFDCACDCFDAFCECBCBCCEFFDCACDEEADAADDACCDAFECDFAEEFCDBAEAFFCDBADDFAAFACFEDDADBDFECAECBCFCAEACDDFEABFFEDADCECDFADFCBDFAFCFFEFECABACADCFCCFCDFDFBBCEFDACBBCCDDBADEDDFBACFFAFBBFBECBCCFEBCFABBBDCCCBDBEBDDCBBDFDDBEFCDADAECDEBFDCEDEBBECAABABEDAAECAEEFFDDFBCADCFFCAFCFECFCEBFDECBEEBCCFCCACFAFEBBADAEECCEEECBBDAEEEDCCDCECAEFBAFDABBECFDFBDEFDFAFDEABEFBAAAACCEAFEFABAAEEAAFABFAACBDCCBDFCEBBECECFEBCEBEFACCFEEBDBAEBDAAFBFFBADADECCABFEEABADEFBBFFEFCFFDEEDECAECABBBAFAFDCFCDADACACDDDBFADBDCEADEAAAACDBDADCCEEEBBFADAEEEDABDCEACBBEFEFEADBDEABEEFABDAFBFAEBBDDFCECECDFBEBFBFBCCBABBBAFDBECBADBACDBAACDEBDFFEDEBCACACADBCFEEDDCFBCABACAADCDAAFEADEDDBCCDFFFCAEAFBEBBCEFCEABCCEBBAAFCCDDFDEBEAECEEADDADFDFADDAABDCADEDDAACFCFADDCEDBDCAEFACFBDADECCCFDEFFDCADDEAACFCCEADCAEECCBAABFBDFADDDBBFDDCAEDDBFCCEBFADFBAFCBEECEAFBFADEBCCADCAAEDCCCFACECDADFAEDEBDDAADBBEECEEFEFEEBDBBBACBBBEEFECEFDAAFEDECCAEABCBCAFBFABFBAFDCDEFDDFFCADBDABFEFCCBDEEFBEDDBABACABBAABCBCBCEDAAABFABFDFAACCFDCDFFCFBAFCFDECEBBBBCFEAFABDEDEBECCCDBEABAAAFFEBAFAEDDBEFAFCBAACCABDCCCEFADFCAEFBFABBBECFCCFDCFDDCECFFCBDACCDACDEAEDFCDADCDDCEBCBEBDDAEBBEDDDCBBEBADECCFEAABFDDFECCCCCDEADFBAACCBAEADBDADFABCDDCECDFEDBFFBDFDAEECDBBAABAAFBECFCDDCBFDBFACDBACAFEBFBFFEEDBEABEFEBABDFBBDDEDFABBDBFAFBEAFAEDCABCFCDEFABAFDDABBDAADCEBACEAEADBEBCBBDAFABCEEEADFCABFFCBFCBACEBEACEAABDCCAFFDABFADDCFBDDAFFABAEDADFFBCBABDCDEACCDAAADEAEDEBFDDBADDBACDEABFFCAEBEBDCFDABACECFBBFBEACCDACBDBCEECFFAECEBADEDDFBEEACECBAFDBFFAFBBDCCCEACEEACFABECDBBAAFCDABFADFFACCFDECAFDCEFDDBDFBAFAEBEFEADFCFBDEBCACDBEFDEBAACCADFDDFEDEABBDDAEEADEBDFEBAFDCBCBABCDBAFBCDDEEDDCCEAFFCCADAEDFCAFABEDFECFBBAEEABBDBECCCCCEABFFFBBBFBEDDDEAFFFEAEFFDAADDEABEDDFDBBEFABFECFEECCCCDBAECFCEACEDEBAFCCFDEDCFCCDBDEBCCFBFACCDDAFFBCDEAFEEDCDABBFEAADDECDBFBBFAEDAACABEFFFFECFADEEAEDEEEDFFBCBFABEDABBDBBEDBFEEAAACFFDBBDEEBBBDDDCCEBBCDDFEEFDFBFCBCEFEDCEFFECCBDEEFFECABFACBFCFFEFAECDEABAFECEDEEBAFBBCFFBDAACFBCBFBEFABBBCECCDAACBFBAEABCBBEFEFCDDBBAEFEFBDBAECDBCEAFECEDDECAFAEABCCEAAEACEFFFADDCDEBFDAAACFCCEBEBABCFBDDEECEDBFFFADBDFABDFFCEDCAAEFCDCFECBBCCCDAAABDBAECBABDABEAAFABBFFCDECFCAEFCAFEECBFCAEDCFFDBBFDEBBCBCFBFFFFFECDBEFDEEDFBAFADBBBFFEEFFECFEEEDABBCBDEDEFDCBEBCCAEFFADDAAECCEFEFCFADBAADAFDEBBEECABDEEDCEBCABFBCECDDCCDBBABDABCEDBDCBBDDACECBDEFBEEEEBACBBADFDBACDCDCEBFDEBDAAFCDDBCCCBACAECCAADFBCBACBFBDFFCCCAFBBBBACEFABBEFAAFFDCEEDDDDFBBBADDEEDFBFADCDBDAAAEEFAFFAFCAACFCCFFDAFFCCDCECDEBADACCFFEBCAFAAFBB", 4567, 5526},
	}

	for _, val := range crs {
		actual := characterReplacement(val.s, val.k)

		if actual != val.expected {
			t.Error("Test Failed!")
		}
	}

}