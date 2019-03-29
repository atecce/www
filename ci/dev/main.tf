provider "cloudflare" {
    email = "root@atec.pub"
    token = "${file("/keybase/private/atec/etc/cloudflare/token")}"
}

resource "cloudflare_record" "CNAME" {
    name = "dev"
    domain = "atec.pub"
    type = "CNAME"
    value = "kbp.keybaseapi.com"
}

resource "cloudflare_record" "TXT" {
    name = "_keybase_pages.dev.atec.pub"
    domain = "atec.pub"
    type = "TXT"
    value = "kbp=/keybase/private/atec,kbpbot/dev"
}
