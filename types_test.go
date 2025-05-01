package guestline_test

import (
	"encoding/xml"
	"testing"

	"github.com/omniboost/go-guestline"
)

func TestXMLMapStringStructUnmarshalSlice(t *testing.T) {
	b := `<root>
		<header1>
			<test>1</test>
		</header1>
		<header1>
			<test>2</test>
		</header1>
		<header1>
			<test>3</test>
		</header1>
	</root>`

	target := struct {
		guestline.XMLMapStringStruct[[]struct {
			Test string `xml:"test"`
		}]
	}{}

	err := xml.Unmarshal([]byte(b), &target)
	if err != nil {
		t.Error(err)
	}

	if len(target.XMLMapStringStruct["header1"]) != 3 {
        t.Errorf("expected 3, got %d", len(target.XMLMapStringStruct["header1"]))
    }
}
