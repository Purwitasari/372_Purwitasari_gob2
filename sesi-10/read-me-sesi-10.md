## Middleware

Middleware merupakan sebuah fungsi yang akan terkesekusi sesudah maupun sebelum mencapai sebuahendpoint. Biasa middleware digunakan untuk logging atau untuk mengamankan sebuah endpoint seperticontohnya proses autentikasi dan autorisasi.Untuk membuat middleware pada bahasa Go, kita akan menggunakan package net/http dengan menggunakanmultiplexer nya agar kita dapat melakukan kustomisasi.

## Securing Our App With JsonWebToken

JsonWebToken adalah sebuah token berbentuk string panjang yang digunakan untuk pertukaran informasi antara 2 belah pihak atau client dan server.

JsonWebToken biasa digunakan pada aplikasi web yang berbasis SPA atau Single Page Application dengan alur sebagai berikut:
- User melakukan login melalui client, biasanya user hanya perlu menginput email dan password saja.
- Setelah user memencet tombol submit, maka client akan mengirimkan data yang di input oleh user kepada server.
- Setelah data diterima oleh server, maka server akan memeriksa terlebih dahulu apakah data yang dikirimkan oleh client merupakan data yang valid atau tidak.
- Jika datanya valid, maka server akan mengirimkan data user tersebut kepada client, namun data user nya telahdijadikan sebuah token dengan menggunakan JWT ( JsonWebToken ).
- Lalu client akan menyimpan token yang dikirimkan oleh server pada local storage atau session storage, dan jikaclient hendak mengirimkan request kepada server kembali, maka client perlu mengirimkan token tersebut kepadaserver melalui request headers.
- Lalu ketika token tersebut sudah oleh server, maka server akan melakukan proses autentikasi terjadap tokentersebut yang dikirimkan oleh client.
