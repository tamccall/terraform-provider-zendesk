package zendesk

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nukosuke/go-zendesk/zendesk"
	"github.com/nukosuke/go-zendesk/zendesk/mock"
)

func TestSystemField(t *testing.T) {
	ctrl := gomock.NewController(t)
	c := mock.NewClient(ctrl)

	m := newIdentifiableGetterSetter()
	id := int(1234)
	m.Set("id", id)

	out := zendesk.TicketField{
		URL: "foobar",
	}

	c.EXPECT().GetTicketField(gomock.Eq(int64(id))).Return(out, nil)
	err := readSystemField(m, c)
	if err != nil {
		t.Fatalf("Read system field returned an error. %v", err)
	}

	if v, ok := m.GetOk("url"); !ok || v.(string) != out.URL {
		t.Fatalf("Read system field did not set URL field. Expected %v, Got %v", out.URL, v)
	}
}