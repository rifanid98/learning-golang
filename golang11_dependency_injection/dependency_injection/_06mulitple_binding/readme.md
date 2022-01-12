# Multiple Binding

- Saat melakukan dependency injection, kadang ada kasus kita membuat beberapa Provider dengan tipe yang sama
- Hal ini akan membuat error proses auto generate kode dependency injection, karena Google Wire tidak mendukung multiple
  binding dengan tipe yang sama
- Pada kasus ini, kita bisa membuat tipe alias untuk multiple binding 