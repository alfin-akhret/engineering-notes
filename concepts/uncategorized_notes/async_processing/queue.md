# Queue
Catatan berisi konsep-konsep yang berhubungan dengan queue atau antrian.

## 1. At-least-once delivery
Sebuah mekanisme dalam dalam antrian yang memastikan setidaknya satu pesan terkirim (terproses). Pesan bisa gagal dikarenakan oleh error yg sifatnya sementara seperti network error, host server tidak merespon, dll.

contoh: pengiriman email.
notes: q=queue/broker, c=consumer
```
q:mail_queue 
-> c: kirim mail 
-> c: gagal (eg: network error, etc)
-> q: NACK 
-> q: masukin main_queue lagi 
-> c: kirim lagi 
-> berhasil
```

## 2. Poison message
message yang valid masuk queue tapi `tidak akan pernah berhasil diproses` sehingga menyebabkan retry terus menerus. contoh: mengirim email ke alamat yg salah.

```
q:mail_queue 
-> c: kirim mail 
-> c: gagal (invalid mail address)
-> q: NACK 
-> q: masukin main_queue lagi 
-> c: kirim lagi
-> c: gagal (invalid mail address)
-> q: NACK 
-> q: masukin main_queue lagi  
-> ...
-> ...
```

Akibat:
- CPU kepake terus
- Log penuh
- Queue macet

## 3. DLQ (Dead Letter Queue)
Masalah diatas dapat ditanggulangi lewat DLQ, yaitu queue khusus tempat menyimpan message2 yg gagal diproses.

```
q:mail_queue 
-> c: kirim mail 
-> c: gagal (invalid mail address)
-> c: RETRY 3x
-> c: kalau masih gagal, masukin ke DLQ
```
