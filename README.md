# SimpleLog 
> [!IMPORTANT]
> This project is still a work in progress!

A simple, infinite-scrolling digital logbook, supporting CRUD operations, following REST API guidelines. The backend is written in `Go`, using `sqlc` to generate typesafe database code from raw Postgres queries. The `Postgres` database itself is deployed using `Docker`.

For the frontend, I use `htmx` and `hyperscript` to handle interactivity, and used basic `html` and `css` to display and style the application.

To render the `html` to the client, I used the `html/template` package from the standard Go library, and broke down various elments into component templates.

https://github.com/user-attachments/assets/38af62bb-9f0e-4a3d-8d90-54da97de0037

# Deployment
1. Clone the repo:
    - `git clone git@github.com:Remi-Godin/simple-log.git`
2. Modify the `template.env` file with your own env variables, then rename it to `.env`.
3. Deploy the application:
    - `make dockerup`
    - `make migrateup`
4. Run the application:
    - `go run cmd/simple-log/simple-log.go`
