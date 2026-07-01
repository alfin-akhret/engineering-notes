# Context (go)

3 Context penting di golang:
- WithCancel
- WithTimeout
- WithDeadline

## 1. WithCancel
```
ctx, cancel := context.WithCancel(parent)
```
-> bikin konteks turunan yang bisa di-cancel kapan saja

contoh:
```
func WillDoSomething() error {

    ctx, cancel := context.WithCancel(parent)

    do.StartWorker(ctx)

    cancel()
    return nil
}
```
saat `cancel(`) dipanggil maka semua worker akan stop.

cocok untuk:
- shutdown worker
- stop goroutine
- propagate cancel manual

## 2.WithTimeout
```
ctx, cancel := context.WithTimeout(parent, 2*time.Second)
```
-> bikin konteks turunan yang otomatis cancel setelah 2 detik.

contoh:
```
func WillDoSomething() error {

    ctx, cancel := context.WithTimeout(parent, 2*time.Second)
    defer cancel()

    do.Something(ctx)

    return nil
}
```
jika `do.Something` lewat dari 2 second, maka otomatis cancel di panggil, jika tidak cancel akan dipanggil saat fungsi `WillDoSomething` return.

cocok untuk:
- worker operation timeout
- http timeout
- db operation timeout

## WithDeadline
```
ctx, cancel := context.WithDeadline(parent, time.Now().Add(2*time.Second))
```
mirip dengan WithTimeout bedanya ini timeout nya menggunakan TimeStamp (waktu absolute), sedangkan WithTimeout menggunakan durasi.
WithTimeout sebenernya wrapper dari WithDeadline.





