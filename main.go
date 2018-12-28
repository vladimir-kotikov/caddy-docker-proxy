package main

import (
	_ "github.com/lucaslorentz/caddy-docker-proxy/plugin"

	// DNS Providers
	_ "github.com/caddyserver/dnsproviders/azure"
	_ "github.com/caddyserver/dnsproviders/cloudflare"
	_ "github.com/caddyserver/dnsproviders/digitalocean"
	_ "github.com/caddyserver/dnsproviders/godaddy"
	_ "github.com/caddyserver/dnsproviders/googlecloud"
	_ "github.com/caddyserver/dnsproviders/linode"
	_ "github.com/caddyserver/dnsproviders/route53"
	_ "github.com/caddyserver/dnsproviders/vultr"

	// Plugins
	_ "github.com/BTBurke/caddy-jwt"
	_ "github.com/SchumacherFM/mailout"
	_ "github.com/caddyserver/forwardproxy"
	_ "github.com/captncraig/cors/caddy"
	_ "github.com/echocat/caddy-filter"
	_ "github.com/hacdias/caddy-minify"
	_ "github.com/miekg/caddy-prometheus"
	_ "github.com/nicolasazrak/caddy-cache"
	_ "github.com/tarent/loginsrv/caddy"
	_ "github.com/xuqingfeng/caddy-rate-limit"

	// Caddy
	"github.com/mholt/caddy/caddy/caddymain"
)

func main() {
	caddymain.Run()

	// Keep caddy running after main instance is stopped
	select {}
}
