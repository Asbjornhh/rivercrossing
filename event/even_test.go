package event

import (
	"testing"
)

func TestPutIn(t *testing.T) {
	wanted := "[ Rev Korn ---V \\ \\_HS+Kylling_/ _____________/ Ø---]"
	got := PutIn()
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}

func TestCrossRiver(t *testing.T) {
	wanted := "[ Rev Korn ---V \\ _____________\\_HS+Kylling_/ / Ø---]"
	got := CrossRiver()
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}

func TestTakeOut(t *testing.T) {
	wanted := "[ Rev Korn ---V \\ _____________\\_HS_/ / Ø--- Kylling]"
	got := TakeOut()
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}
