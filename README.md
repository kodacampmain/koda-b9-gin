# KODA B9 GIN
[![License](https://img.shields.io/badge/License-MIT-green)](https://opensource.org/license/mit)
[![Go](https://img.shields.io/badge/Go-1.27.1-blue?logo=Go)](https://go.dev/)
![Gin-Gonic](https://img.shields.io/badge/Gin%20Gonic-1.12.0-blue?logo=Gin&logoColor=fff300)

<br>

![Logo Koda](https://kodacampmain.github.io/assets-library/img/logo-black.png)

<br>

Project for gin exercise by Koda batch 9

<br>

## Technologies
- Go
- Gin-gonic
- PostgreSQL
- Redis
- Swagger
- etc.

<br>

## Features

- Authentication (Register, Login, Logout, Forgot Password)
- User Profile
- Event List
- Community List
- etc.

<br>

## How to use

### Setup

#### Database

Install menggunakan docker
```bash
$ docker pull postgresql
$ docker run ....
```
#### Redis
1. Siapkan config redis
```
## redis.conf
bind 0.0.0.0
port 6379

dst.
```
2. Install menggunakan docker
```bash
$ docker pull redis:version
$ docker run ....
```

### Run
1. Clone project
```bash
$ git clone https://github.com/kodacampmain/koda-b9-gin.git
```
2. Siapkan env ([.env.example](https://github.com/kodacampmain/koda-b9-gin/blob/master/.env.example))

3. Install dependencies
```bash
$ go mod download
```

4. dst.