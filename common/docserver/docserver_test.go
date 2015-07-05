package docserver

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestGetAuthContent(t *testing.T) {
	text := "test"
	hash := common.Hash{}
	copy(hash[:], crypto.Sha3([]byte(text)))
	ioutil.WriteFile("/tmp/test.content", []byte(text), os.ModePerm)

	ds := New("/tmp/")
	content, err := ds.GetAuthContent("file:///test.content", hash)
	if err != nil {
		t.Errorf("no error expected, got %v", err)
	}
	if string(content) != text {
		t.Errorf("incorrect content. expected %v, got %v", text, string(content))
	}

	hash = common.Hash{}
	content, err = ds.GetAuthContent("file:///test.content", hash)
	expected := "content hash mismatch 0000000000000000000000000000000000000000000000000000000000000000 != 9c22ff5f21f0b81b113e63f7db6da94fedef11b2119b4088b89664fb9a3cb658 (exp)"
	if err == nil {
		t.Errorf("expected error, got nothing")
	} else {
		if err.Error() != expected {
			t.Errorf("expected error '%s' got '%v'", expected, err)
		}
	}

}

type rt struct{}

func (rt) RoundTrip(req *http.Request) (resp *http.Response, err error) { return }

func TestRegisterScheme(t *testing.T) {
	ds := New("/tmp/")
	if ds.HasScheme("scheme") {
		t.Errorf("expected scheme not to be registered")
	}
	ds.RegisterScheme("scheme", rt{})
	if !ds.HasScheme("scheme") {
		t.Errorf("expected scheme to be registered")
	}
}
