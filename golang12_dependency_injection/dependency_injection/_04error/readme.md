# Error

- Google Wire juga bisa mendeteksi jika terjadi error pada Provider kita
- Jika terdapat error, secara otomatis akan mengembalikan error ketika kita melakukan dependency injection
- Caranya sederhana, kita cukup buat di Provider return value kedua berupa error, dan di Injector nya juga perlu kita
  tambahkan return value kedua berupa error 