Biodata 
Kode Peserta : 149368582100-372
Nama		 : Purwitasari
Link Github  : https://github.com/Purwitasari/372_Purwitasari_gob2/tree/main/ass-02

Paduan : 
1. Download folder ass-02 dari Github diatas
2. Buka VS Code dan open folder ass-02
3. Buka aplikasi XAMPP dan jalankan MySql
4. Buka localhost/phpmyadmin/ dan buat database dengan nama "ass_godb"
5. Masuk ke terminal di vs code kemudian ketik "go run main.go"
6. Untuk mengakses data dapat menggunakan link Postman dibawah

Creat Data : http://localhost:1212/orders 
body: customer_name, item_code, description, quantity

Get All Data : http://localhost:1212/orders

Update Data : http://localhost:1212/orders/1            (update id:1)
body: customer_name, item_code, description, quantity

Delete Data : http://localhost:1212/orders/1            (delete id:1)