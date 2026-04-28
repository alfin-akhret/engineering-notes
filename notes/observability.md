# Observability

Kemampuan untuk mengetahui kondisi internal system dari luar berdasarkan data dari system itu sendiri.

| Tools      | Fungsi                                                       | Stack                                                                                    |
| ---------- | ------------------------------------------------------------ | ---------------------------------------------------------------------------------------- |
| Logging    | mencatat apa saja yang terjadi di system                     | ELK (event-based, debugging)                                                             |
| Monitoring | menggunakan metric untuk mengetahui status kesehatan system. | Prometheus (time-series, monitoring) + Grafana                                           |
| Alerting   | pemberitahuan jika ada kondisi yang membutuhkan tindakan     | Prometheus (time-series, monitoring) + alert manager + alert channel (slack, email, etc) |
| Tracing    | kemampuan untuk menelusuri sumber masalah                    | Open Telemetry (tracer) + Jaeger (dashboard)                                             |

## 📦 1. Logging Flow (ELK stack)

Fungsi: untuk menjawab pertanyaan:
- apa yang terjadi?
- kenapa error?
- request detail apa?

Berikut contoh alur logging menggunakan zap.logger, filebeat/logstash, elastic search dan kibana.

```
app (Go service)
  ↓
structured logger (zap)
  ↓
stdout (docker container logs)
  ↓
filebeat (read /var/lib/docker/containers/*/*.log)
  ↓
decode_json_fields (parse log JSON)
  ↓
elasticsearch (store logs as documents)
  ↓
kibana
  ↓
search / filter / visualize logs
```

### Perubahan data
```
zap logger
  ↓
JSON log event
  ↓
{"level":"info","msg":"http request completed","request_id":"..."}
  ↓
filebeat harvest log line
  ↓
parse JSON → fields
  ↓
index: filebeat-*
  ↓
Kibana Discover / Logs UI
```

## 📊 2. Monitoring Flow (Prometheus + Grafana)

Fungsi:
untuk menjawab pertanyaan:
- system sehat atau tidak?
- latency naik atau tidak?
- traffic spike atau tidak?
  

```
app (Go service)
  ↓
/metrics endpoint (prometheus handler)
  ↓
prometheus scrape (HTTP pull)
  ↓
time series storage (Prometheus TSDB)
  ↓
Grafana (query PromQL)
  ↓
dashboard visualization
```

### 🧠 Detail metrics flow
```
http_requests_total (counter)
http_request_duration_seconds (histogram)
http_errors_total (counter)
http_requests_in_flight (gauge)
  ↓
Prometheus scrape interval (5s)
  ↓
rate(), histogram_quantile()
  ↓
Grafana panels (time series / bar chart)
```

## 3. Alert Flow
Fungsi: memberi tahu jika ada kondisi yang membutuhkan tindakan.

```
app (Go service)
  ↓
Alert Rules (p95 latency, error rate, etc.)
  ↓ 
/metrics endpoint (prometheus handler)
  ↓
prometheus scrape (HTTP pull)
  ↓
time series storage (Prometheus TSDB)
  ↓
prometheus alert manager
  ↓
notification channels: slack, email, etc
```

## 4. Tracing Flow
Fungsi: kemampuan untuk menelusuri sumber masalah

```
Go App (OpenTelemetry SDK)
        │
        ▼
OTLP Exporter
        │
        ▼
OpenTelemetry Collector
        │
        ├──────────────► Jaeger (debug UI)
        │
        └──────────────► Grafana Tempo (production)
```

## Full system view (LOGGING + METRICS TOGETHER)
```
                ┌────────────────────┐
                │   Go App (API)     │
                │--------------------│
                │ - zap logger       │
                │ - prometheus metrics│
                └───────┬────────────┘
                        │
        ┌───────────────┴────────────────┐
        │                                │
        ▼                                ▼

stdout logs                      /metrics endpoint
        │                                │
        ▼                                ▼
   filebeat                         prometheus scrape
        │                                │
        ▼                                ▼
elasticsearch                    prometheus TSDB
        │                                │
        ▼                                ▼
   kibana                        grafana dashboard

```

