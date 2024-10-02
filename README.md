# SimpleLog
A simple, infinite-scrolling digital logbook, supporting CRUD operations. The backend is written in `Go`, using `sqlc` to generate typesafe database code from raw Postgres queries. The `Postgres` database itself is deployed using `Docker`.

For the frontend, I used `htmx` and `hyperscript` to handle interactivity, and used basic `html` and `css` to display and style the application.

To render the `html` to the client, I used the `html/template` package from the standard Go library, and broke down various elments into component templates.



https://github.com/user-attachments/assets/ca942b87-f973-4679-9da3-1b5b65ceb1c8

