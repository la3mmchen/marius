---
{{- $rulefile := .}}
rule_files:
- {{$rulefile.FileName}}

evaluation_interval: 1m

group_eval_order:
{{range $group := $rulefile.Groups -}}
  - {{$group.Name}}
{{end}}

tests:
- interval: 1m
  input_series:
    # replace this time series with time series matching the series you want to have
{{range $key, $value := .TmplSeries}}
    - series: {{$key -}}{{- "{" -}}{{- $value -}}="unittest"{{- "}"}}
      values: '0 0 0 0 0 0 0 0 0 0 0'
{{- end}}

  alert_rule_test:
{{- range $group := .Groups}}{{range $rule := $group.Rules}}
      - eval_time: {{$rule.For}}
        alertname: {{$rule.Alert}}
        exp_alerts:
            - exp_labels:
{{- range $key, $value := .AllLabels}}
                {{$key}}: {{$value}}
{{- end}}
{{- range $label,$value := $rulefile.Labels}}
                {{$label}}: unittest
{{- end}}
              exp_annotations:
                  summary: "{{$rule.Annotations.Summary}}"
                  description: "{{$rule.Annotations.Description}}"
{{- end}}
{{- end}}


