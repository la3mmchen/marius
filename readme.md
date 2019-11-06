# marius - a friendly prometheus rule test templater

Marius takes your prometheus rules and creates files for unittests from all of them.

## why

Ever broke your monitoring with an invalid alerting rule? Or tried to check if a rule is working with shutting down various services to get a signle from your prometheus? Or want to have a proper CI/CD pipeline for your prometheus setup?

Prometheus alert rules can tested (see 1). Quite uncool is that there is no way to bootstrap a test file for your multiple rules to start with. Its tedious to translate the myriad our your rules into proper testfiles. Marius can help you at least a little bit.

## how does it work

At the moment please build your binary for yourself:

```bash
make build
```

Afterwards you can run it: 

```bash
./bin/marius
```

## how to check created test files

At the moment there is a Makefile target that runs prometheus unit tests for your created test files

```bash
make prom
(...)

Unit Testing:  test-example-rules.yaml
  FAILED:
    alertname:ScrapeTargetDown, time:5m0s,
        exp:"[Labels:{alertname=\"ScrapeTargetDown\", instance=\"unittest\", layer=\"monitoring\", page=\"true\", runbook=\"https://example.example.de/confluence/display/Runbooks\", severity=\"warning\"} Annotations:{description=\"The instance has been not scraped for metrics more than 5 minutes. This indicates eithe
r connectivity issues or a problem on the instance.\", summary=\"Sracping from Instance {{ $labels.instance }} not possible\"}]",
        got:"[Labels:{alertname=\"ScrapeTargetDown\", instance=\"unittest\", layer=\"monitoring\", page=\"true\", runbook=\"https://example.example.de/confluence/display/Runbooks\", severity=\"warning\"} Annotations:{description=\"The instance has been not scraped for metrics more than 5 minutes. This indicates eithe
r connectivity issues or a problem on the instance.\", summary=\"Sracping from Instance unittest not possible\"}]"

make: *** [prom] Error 1

Unit Testing:  test-example-rules.yaml
  SUCCESS
```

## what should be created

Given a simple alerting rule marius should create something like this:

Input:

```yaml
    - alert: ScrapeTargetDown
      expr: 'up == 0'
      for: '5m'
      labels:
        severity: 'warning'
        layer: 'monitoring'
        runbook: 'https://docu.example.de/confluence/display/Runbooks'
        page: 'true'
      annotations:
        summary: 'Sracping from Instance {{ $labels.instance }} not possible'
        description: 'The instance has been not scraped for metrics more than 5 minutes. This indicates either connectivity issues or a problem on the instance.'
```

Output:

```yaml
tests:
- interval: 1m
  input_series:
    # replace this time series with time series matching the series you want to have

    - series: up{instance="unittest"}
      values: '0 0 0 0 0 0 0 0 0 0 0'

  alert_rule_test:
      - eval_time: 5m
        alertname: ScrapeTargetDown
        exp_alerts:
            - exp_labels:
                layer: monitoring
                page: true
                runbook: https://docu.example.de/confluence/display/Runbooks
                severity: warning
                instance: unittest
              exp_annotations:
                  summary: "Sracping from Instance unittest not possible"
                  description: "The instance has been not scraped for metrics more than 5 minutes. This indicates either connectivity issues or a problem on the instance."
```

## not working right now

Marius is pretty much alpha. Be kind if it does not work :P 

- marius should create the test files under a definable path but right now is hard-coded to write into `data/`
- marius does not replace alread integrated golang template within labels (e.g. annotations) - yet.
- the parsing of the `expr` field from the rules will not gather all metrics so `input_series` is not populated correctly.
- marius does not know what values a metric is providing so he just insert 0 values as the test data.
- the docker build is broken

## appendix

(1) <https://prometheus.io/docs/prometheus/latest/configuration/unit_testing_rules/>
