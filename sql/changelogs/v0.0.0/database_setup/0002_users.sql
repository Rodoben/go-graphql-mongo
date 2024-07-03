--liquibase formatted sql

--changeset ronald.benjamin:1 context:dev and qa and rel splitStatements:false
-- create special user
DO
$do$
BEGIN
    IF EXISTS (SELECT FROM pg_user
        WHERE  usename = '${DB_USER}'
        ) THEN

        RAISE NOTICE 'User already exists!';
    ELSE
        CREATE USER ${DB_USER} WITH ENCRYPTED PASSWORD '${DB_PASS}';
    END IF;
END
$do$;
