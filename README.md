
# Mangamee-Api

Open source manga api


## Run Locally

Clone the project

```bash
  git clone https://github.com/Dhiiyaur/mangamee-api
```

Go to the project directory

```bash
  cd my-project
```

Install dependencies

```bash
  go mod download
```

Setup Environment (.env)
```bash
  see .env.example
```

Start the server

```bash
  make run

  or
  
  go run cmd/server/main.go
```

## Usage

| Name              | Method   | URL  | Detail |
| ----------------- | -------- | ---- | ---    |
| Source            | `GET`    | `/manga/source` | Retrieve all source ID
| Index             | `GET`    | `/manga/index/:sourceId/:page` | Retrieve manga from source "x" with page "x"
| Search            | `GET`    | `/manga/search/:sourceId?title=a` | Retrieve manga from source "x" with title a
| Detail            | `GET`    | `/manga/detail/:sourceId/:mangaId` | Retrieve manga detail manga from source "x" with mangaID "x"
| Chapter           | `GET`    | `/manga/read-chapter/:sourceId/:mangaId`|  Retrieve manga chapter manga from source "x" with mangaID "x"
| Read              | `GET`    | `/manga/read/:sourceId/:mangaId/:chapterId` |  Retrieve manga from source "x" with mangaID "x" and chapterID "x"

