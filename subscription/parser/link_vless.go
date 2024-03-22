package parser

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"

	C "github.com/sagernet/sing-box/constant"
	"github.com/sagernet/sing-box/option"
	E "github.com/sagernet/sing/common/exceptions"
)

func ParseVlessLink(link string) (option.Outbound, error) {
	outbound := option.Outbound{}
	if u, err := url.Parse(link); err == nil {
		if u.Scheme != C.TypeVLESS {
			return outbound, E.New("unsupported scheme: ", u.Scheme)
		}

		outbound.Type = C.TypeVLESS
		outbound.Tag = fmt.Sprintf("serenity-%s-%s", u.Fragment, generateUniqueString(6))
		port, err := strconv.ParseUint(u.Port(), 10, 16)
		if err != nil {
			return outbound, err
		}
		options := option.VLESSOutboundOptions{
			ServerOptions: option.ServerOptions{
				Server:     u.Hostname(),
				ServerPort: uint16(port),
			},
			UUID: u.User.Username(),
		}
		host := u.Query().Get("sni")
		if host == "" {
			host = u.Query().Get("host")
		}
		//if u.Query().Get("security") != "tls" {
		//	return option.Outbound{}, E.New("unsupported security: ", u.Query().Get("security"))
		//}

		options.OutboundTLSOptionsContainer = option.OutboundTLSOptionsContainer{
			TLS: &option.OutboundTLSOptions{
				Enabled:    true,
				ServerName: host,
				Insecure:   true,
			},
		}

		if u.Query().Get("type") == "ws" {
			headers := make(map[string]option.Listable[string])
			headers["Host"] = []string{host}
			path, _ := url.QueryUnescape(u.Query().Get("path"))
			options.Transport = &option.V2RayTransportOptions{
				Type: C.V2RayTransportTypeWebsocket,

				WebsocketOptions: option.V2RayWebsocketOptions{
					Path:    path,
					Headers: headers,
				},
			}
		}
		outbound.VLESSOptions = options
		return outbound, nil
	} else {
		return outbound, err
	}
}

func generateUniqueString(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)[:length]
}
