package olclient

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTags(t *testing.T) {
	jsonStr := `{"publishers": ["A\u0304ra. A\u0304ra. S\u0301et\u0323hani\u0304 Kampani\u0304"], "series": ["Sam\u0323ska\u0304ra grantha\u0304vali ;", "pustaka 418 mum\u0323"], "lc_classifications": ["MLCSA 86/1661 (P)"], "key": "/books/OL2163622M", "authors": [{"key": "/authors/OL875222A"}], "publish_places": ["Mumbai\u0304"], "languages": [{"key": "/languages/guj"}], "pagination": "20, 602 p. ;", "source_records": ["marc:marc_records_scriblio_net/part20.dat:112583293:622", "marc:marc_loc_updates/v37.i24.records.utf8:2426129:669", "marc:marc_loc_2016/BooksAll.2016.part19.utf8:69312919:669"], "title": "Asu\u0304ryaloka", "lccn": ["88900775"], "notes": {"type": "/type/text", "value": "Novel.\nIn Gujarati."}, "number_of_pages": 602, "edition_name": "1. a\u0304vr\u0325tti.", "publish_date": "1987", "publish_country": "ii ", "by_statement": "Bhagavati\u0304kuma\u0304ra S\u0301arma\u0304.", "works": [{"key": "/works/OL4420976W"}], "type": {"key": "/type/edition"}, "covers": [10581023], "latest_revision": 5, "revision": 5, "created": {"type": "/type/datetime", "value": "2008-04-01T03:28:50.625462"}, "last_modified": {"type": "/type/datetime", "value": "2021-01-26T04:51:35.127961"}}`
	testWorkJSON(t, jsonStr)

	jsonStr2 := `{"number_of_pages": 1193, "table_of_contents": [{"level": 0, "label": "", "title": "Pt. 1. The fellowship of the ring --", "pagenum": ""}, {"level": 0, "label": "", "title": "pt. 2. The two towers --", "pagenum": ""}, {"level": 0, "label": "", "title": "pt. 3. The return of the king.", "pagenum": ""}], "contributors": [{"role": "Illustrator", "name": "Alan Lee"}], "series": ["The Lord of the Rings"], "covers": [10327719, 9255566], "ia_loaded_id": ["lordofrings00tolk_1"], "lc_classifications": ["PR6039.O32 L6 1993"], "latest_revision": 22, "ocaid": "lordofrings00tolk_1", "description": "This special edition of the Lord of the Rings celebrates the birth of J.R.R. Tolkien with fifty paintings specially commissioned from the noted English artist Alan Lee.\r\n\r\nSince its first publication in the 1950s, The Lord of the Rings has been accepted as a unique creation, becoming, as Tolkien envisioned, one of the living 'myths' of the English-speaking peoples.\r\n\r\nIn this one volume the three parts of The Lord of the Rings, The Fellowship of the Rings, The Two Towers, and The Return of the King are enhanced by the work of a dedicated artist whose vision matches Tolkien's own.\r\n\r\nIn ancient times the Rings of Power were crafted by the Elven-smiths, and Sauron, the Dark Lord, forged the One Ring to rule all the others. But the One Ring was taken from him, and remained lost until after many ages it fell by chance into the hands of the hobbit, Bilbo, and was then bequeathed to his young nephew Frodo. The Lord of the Rings tells of the quest undertaken by Frodo together with the fellowship of the ring: Gandalf the wizard, Merry, Pippin, and Sam the hobbits, Gimli the dwarf, Legolas the elf, Boromir mand of Gondor, and a tall mysterious stranger called Strider.\r\n\r\nTheir quest was perilous to journey across Middle-earth, deep into the shadow of the Dark Lord, and destroy the Ring by casting it into the Cracks of Doom.\r\n\r\n-jacket flap", "languages": [{"key": "/languages/eng"}], "source_records": ["marc:marc_laurentian/openlibrary.mrc:321070512:736", "ia:lordofrings00tolk_1"], "title": "The Lord of the Rings", "edition_name": "1965 edition; Centenary Edition", "publish_country": "nyu", "by_statement": "by J. R. H. Tolkien ; illustrated by Alan Lee.", "oclc_numbers": ["1057087697"], "type": {"key": "/type/edition"}, "revision": 22, "publishers": ["Houghton Mifflin Company"], "ia_box_id": ["IA177101"], "physical_format": "Hardcover", "last_modified": {"type": "/type/datetime", "value": "2020-08-04T23:27:39.053471"}, "key": "/books/OL21058613M", "authors": [{"key": "/authors/OL26320A"}], "publish_places": ["Boston, USA"], "pagination": "1193p.", "classifications": {"lccn_permalink": ["https://lccn.loc.gov/91010298"]}, "created": {"type": "/type/datetime", "value": "2008-10-31T15:17:48.646102"}, "lccn": ["91010298"], "identifiers": {"librarything": ["1386651"], "goodreads": ["44172570"]}, "isbn_13": ["9780395595114"], "dewey_decimal_class": ["823/.912"], "isbn_10": ["0395595118"], "publish_date": "1993", "copyright_date": "1993", "works": [{"key": "/works/OL27448W"}]}`
	testWorkJSON(t, jsonStr2)
}

func testWorkJSON(t *testing.T, jsonStr string) {
	a := Work{}
	err := json.Unmarshal([]byte(jsonStr), &a)
	assert.NoError(t, err)

	m := map[string]interface{}{}
	err = json.Unmarshal([]byte(jsonStr), &m)
	assert.NoError(t, err)

	encodedJSONStr1, err := json.MarshalIndent(a, "", "\t")
	assert.NoError(t, err)

	encodedJSONStr2, err := json.MarshalIndent(m, "", "\t")
	assert.NoError(t, err)

	assert.Equal(t, string(encodedJSONStr1), string(encodedJSONStr2))
}
