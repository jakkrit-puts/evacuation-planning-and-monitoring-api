﻿# evacuation-planning-and-monitoring-api

### Setup with docker
```
1.docker compose up -d 
```
```
2. create env file 
copy .env.example .env
copy .air.toml.example .air.toml
```

### Usage / Run
```
go run .\cmd   
```
```
==== for dev ====
1.go install github.com/air-verse/air@latest   // install package reload
2.air                                          // for dev mode = run app  
```
```
=== DB Web Tools ===
url adminer db: http://localhost:8080/
```

### Test APIs file
```
api.http   -> vs code install rest client extension
```

### password redis
```
strongpassword123  // หรือ ดูได้ที่ dir  /config/redis.conf  ->  find key requirepass
```
