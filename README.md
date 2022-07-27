# KUMPARAN BE TEST - ARTICLE CASE

## Identity
- Bagas Tri Wibowo
- S1 Informatika
- Universitas Telkom
- IPK: 3.98

## Tech stack
- Golang v1.18.2
- Echo V4
- Sqlx with mysql driver
- Squirrel for query builder
- Golang Playground Validator v10 
- uber Zap for logger
- New Relic for APM
- Mockery for mock generator
- Docker and docker compose for deployment
- Postman API Documentation

## First Set Up

1. copy .env.example to .env
2. Install [DBmate](https://github.com/amacneil/dbmate)
3. run migration script, ex: for linux run `make db-migrate`
4. run main app with `make run`

## Live API Demo

Saya melakukan deployment terhadap service ini dan bisa diakses melalui `https://cloud.krobot.my.id/kumparan-api/{ENDPOINT}`

## Documentation
[Documentation](https://documenter.getpostman.com/view/10876762/UzXNTHEX)

## My Solution

### Query builder
Disini sayan menggunakan Query Builder yaitu squirell, karena query builder dapat menghasilkan native SQL yang baik dan tidak menghasilkan query yang tidak perlu seperti ORM.

### Date Atribute
Saya menambahkan atribut created_at, updated_at, deleted_at sebagai standarisasi dari setiap instance yang dibuat dan mempermudah user dalam mengidentifikasikan kapan instance tersebut dibuat, diupdate, dan dihapus.

### APM (New Relic)
Untuk memonitor performa dari service yang dibangun saya menggunakan APM New Relic. Alasan penggunaan new relic adalah resource yang ditawarkan cukup besar untuk akun gratis. 

## What you need to consider

1. What if there are thousands of articles in the database?
    
    untuk mengatasi hal ini terdapat beberapa solusi, yang pertama adalah dengan Mmnggunakan Pagination, agar data yang ditampilkan lebih dibatasi, sehingga kerja dari service akan lebih ringan karena tidak harus menampilkan ribuan data sekaligus. Kemudian kita dapat pula membatasi hasil dengan menggunakan filter range waktu.

    Jika solusi diatas  dirasa belum cukup, maka kita dapat melakukan horizontal scaling ataupun vertical scaling.

2. What if many users are accessing your API at the same time?

   Jika kemungkinan yang terburuk adalah service down, maka untuk mengatasi hal tersebut adalah kita dapat melakukan caching pada service-service yang cenderung mgenghasilkalkan response yang sama atau jarang berubah. Kemudian kita juga dapat melakukan horizontal scaling dengan melakukan load balancing. Dan solusi terkakhir adalah jika masih dirasa kurang adalah melakukan vertical scaling

## Testing

### Unit test
command unit test
```
make test
```
![make test](https://i.im.ge/2022/07/27/FisnbW.png)
command make covarage
```
make cover
```
![make cover](https://i.im.ge/2022/07/27/Fiaowx.png)
![make cover](https://i.im.ge/2022/07/27/FiarXJ.png)
