### DOKUMENTASI PROGRAM

- Untuk get semua produk/data
```bash
curl -X GET http://localhost:3030/products
```

- Untuk get produk berdasarkan id
```bash
curl -X GET http://localhost:3030/products/1
```

- Untuk create produk
```bash
curl -X POST http://localhost:3030/products \
     -H "Content-Type: application/json" \
     -d '{"name":"Laptop Lenovo","price":8500000,"stock":5}'
```

- Untuk update produk
```bash
curl -X PUT http://localhost:3030/products/1 \
     -H "Content-Type: application/json" \
     -d '{"name":"Laptop Asus ROG","price":12000000,"stock":3}'
```

- Untuk hapus produk
```bash
curl -X DELETE http://localhost:3030/products/1
```
