## courses-api

REST API application to manage courses. Written in Go with Echo Framework.

## Notes

There are two branch in this repository:

- `main`: implementation with clean architecture (WIP).
- `mvc`: implementation with MVC.

## How to use

1. Clone this repository.

2. Copy the configuration template.

```sh
cp .env.example .env
```

3. Fill the database configurations inside the `.env` file.

4. Create a new database.

```sql
CREATE DATABASE courses_api;
```

5. Run the application. Make sure the database is online.

```sh
go run main.go
```
