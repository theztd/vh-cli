{{- range . }}
# {{ . }}
resource "cloudflare_record" "{{ .Type }}_{{ Replace .Name "." "_" }}" {
  zone_id = __ZONE_ID__
  name    = "{{ .Name }}"
  value   = "{{ .Content }}"
  type    = "{{ .Type }}"
  ttl     = {{ .TTL }}
  comment = "{{ . }}"
  proxied = false
}
{{ end }}