{{/*
Compute if the server serviceaccount is enabled.
*/}}
{{- define "vault.serverServiceAccountEnabled" -}}
{{- $_ := set . "serverServiceAccountEnabled"
  (and
    (eq (.Values.server.serviceAccount.create | toString) "true" )
    (or
      (eq (.Values.server.enabled | toString) "true")
      (eq (.Values.global.enabled | toString) "true"))) -}}
{{- end -}}