# GO QR BACKEND   

Ini merupakan repo backend dari aplikasi GO QR. Agar berjalan dengan baik, anda perlu menjalankan dua repo sekaligus. Yaitu repo [Frontend](https://github.com/rezairfanwijaya/go-qr-frontend) dan [Backend](https://github.com/rezairfanwijaya/go-qr-backend). Untuk containerization dapat dipull melalui [image-backend](https://hub.docker.com/repository/docker/rezairfanwijaya/go-qr-backend/general)

#### Step menjalankan repo [Backend](https://github.com/rezairfanwijaya/go-qr-backend) di lokal


#### Clone the project

```bash
  git clone https://github.com/rezairfanwijaya/go-qr-backend.git
```

#### Go to the project directory

```bash
  cd go-qr-backend
```

#### Get Dependency
```bash
  go mod tidy
```

#### Run application
```bash
  go run main.go
```

##### Untuk menjalankan aplikasi GO-QR pastikan repo [Frontend](https://github.com/rezairfanwijaya/go-qr-frontend) dan [Backend](https://github.com/rezairfanwijaya/go-qr-backend) dijalakan secara bersamaan

##### Atau anda dapat langsung menjalankan aplikasi ini dengan command ```docker compose up``` dan akses aplikasi pada alamat ```http://localhost:3000/```




![image](https://user-images.githubusercontent.com/87264553/211035057-ce501d3f-6b08-4a29-9099-c17acec2a944.png)






