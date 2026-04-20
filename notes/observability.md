# Observability

Kemampuan untuk mengetahui kondisi internal system dari luar berdasarkan data dari system itu sendiri.

| Tools      | Fungsi                                                       |
| ---------- | ------------------------------------------------------------ |
| Logging    | mencatat apa saja yang terjadi di system                     |
| Monitoring | menggunakan metric untuk mengetahui status kesehatan system. |
| Alerting   | pemberitahuan jika ada kondisi yang membutuhkan tindakan     |
| Tracing    | kemampuan untuk menelusuri sumber masalah                    |

## Logging
Berikut contoh alur logging menggunakan zap.logger, filebeat, elastic search dan kibana.

```
Application.logger 
    -> stdout/container.log 
    -> filebeat baca log dr stdout atau container log 
    -> filebeat kirim ke elastic search 
    -> kibana baca dari elastic search
```