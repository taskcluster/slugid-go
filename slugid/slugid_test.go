// Licensed under the Mozilla Public Licence 2.0.
// https://www.mozilla.org/en-US/MPL/2.0

package slugid_test

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"github.com/taskcluster/slugid-go/slugid"
	"testing"
)

func ExampleNice() {
	slugid.Nice() // e.g. "eWIgwMgxSfeXQ36iPbOxiQ"
}

func ExampleDecode() {
	fmt.Printf("%s\n", slugid.Decode("eWIgwMgxSfeXQ36iPbOxiQ"))
	// Output: 796220c0-c831-49f7-9743-7ea23db3b189
}

func ExampleEncode() {
	fmt.Println(slugid.Encode(uuid.Parse("796220c0-c831-49f7-9743-7ea23db3b189")))
	// Output: eWIgwMgxSfeXQ36iPbOxiQ
}

func ExampleV4() {
	slugid.Nice() // e.g. "-9OpXaCORAaFh4sJRk7PUA"
}

// Test that we can correctly encode a "non-nice" uuid (with first bit set) to
// its known slug. The specific uuid was chosen since it has a slug which
// contains both `-` and `_` characters.
func TestEncode(t *testing.T) {

	// 10000000010011110011111111001000110111111100101101001011000001101000100111111011101011101111101011010101111000011000011101010100....
	// <8 ><0 ><4 ><f ><3 ><f ><c ><8 ><d ><f ><c ><b ><4 ><b ><0 ><6 ><8 ><9 ><f ><b ><a ><e ><f ><a ><d ><5 ><e ><1 ><8 ><7 ><5 ><4 >
	// < g  >< E  >< 8  >< _  >< y  >< N  >< _  >< L  >< S  >< w  >< a  >< J  >< -  >< 6  >< 7  >< 6  >< 1  >< e  >< G  >< H  >< V  >< A  >
	uuid_ := uuid.Parse("804f3fc8-dfcb-4b06-89fb-aefad5e18754")
	expectedSlug := "gE8_yN_LSwaJ-6761eGHVA"
	actualSlug := slugid.Encode(uuid_)

	if expectedSlug != actualSlug {
		t.Errorf("UUID not correctly encoded into slug: '" + expectedSlug + "' != '" + actualSlug + "'")
	}
}

// Test that we can decode a "non-nice" slug (first bit of uuid is set) that
// begins with `-`
func TestDecode(t *testing.T) {
	// 11111011111011111011111011111011111011111011111001000011111011111011111111111111111111111111111111111111111111111111111111111101....
	// <f ><b ><e ><f ><b ><e ><f ><b ><e ><f ><b ><e ><4 ><3 ><e ><f ><b ><f ><f ><f ><f ><f ><f ><f ><f ><f ><f ><f ><f ><f ><f ><d >
	// < -  >< -  >< -  >< -  >< -  >< -  >< -  >< -  >< Q  >< -  >< -  >< -  >< _  >< _  >< _  >< _  >< _  >< _  >< _  >< _  >< _  >< Q  >
	slug := "--------Q--__________Q"
	expectedUuid := uuid.Parse("fbefbefb-efbe-43ef-bfff-fffffffffffd")
	actualUuid := slugid.Decode(slug)

	if expectedUuid.String() != actualUuid.String() {
		t.Errorf("Slug not correctly decoded into uuid: '%s' != '%s'", expectedUuid, actualUuid)
	}
}

// Test that 10000 v4 uuids are unchanged after encoding and then decoding them
func TestUuidEncodeDecode(t *testing.T) {
	for i := 0; i < 10000; i++ {
		uuid1 := uuid.NewRandom()
		slug := slugid.Encode(uuid1)
		uuid2 := slugid.Decode(slug)
		if uuid1.String() != uuid2.String() {
			t.Errorf("Encode and decode isn't identity: '%s' != '%s'", uuid1, uuid2)
		}
	}
}

// Test that 10000 v4 slugs are unchanged after decoding and then encoding them.
func TestSlugDecodeEncode(t *testing.T) {
	for i := 0; i < 10000; i++ {
		slug1 := slugid.V4()
		uuid_ := slugid.Decode(slug1)
		slug2 := slugid.Encode(uuid_)
		if slug1 != slug2 {
			t.Errorf("Decode and encode isn't identity: '%s' != '%s'", slug1, slug2)
		}
	}
}
