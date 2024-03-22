package parser

import (
	"testing"
)

func TestParseVlessLink(t *testing.T) {
	t.Parallel()
	link := "vless://9247e317-7ef0-4d9f-86db-d394378d7402@c.xf.free.hr:443?encryption=none&security=tls&sni=v01.2023abc1.xyz&fp=random&type=ws&host=v01.20210401.xyz&path=%2F%3Fed%3D2048#%E8%8A%82%E7%82%B9%E4%BD%BF%E7%94%A8%E6%95%99%E7%A8%8B%2Fguide https://t.me/edtunnel/7462"
	if u, err := ParseVlessLink(link); err == nil {
		t.Logf("%v", u)
	} else {
		t.Error(err)
	}
}
