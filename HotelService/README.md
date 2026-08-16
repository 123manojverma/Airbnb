### models --> we write the TS/JS representation (Classes) of our tables
### seeders --> in this folder we can write code to put dummy/seed data in our tables
### migrations --> is used to create versions of our db

- migration --> up | down
up->contains the code which will make new changes in the db when we run the migration
down -> contains the code which revert the changes made by the migration if we want to rollback

## Commands
- to generate

npx sequelize-cli migration:generate --name (name)

- to migrate

npx sequelize-cli db:migrate

- to revert the changes

npx sequelize-cli db:migrate:undo