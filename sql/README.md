# sql

This directory is meant to contain all SQL and helper scripts for creating and updating the Time service database


### Directory Structure

Generated using `tree (Resolve-Path .\sql).Path /a` command

```
+---changelogs
|   \---v0.0.0
|       +---database_setup
|       +---permissions
|       +---structure
|       +---tables
|       \---views
\---scripts
```

| Folder                           | Usage                     |
|----------------------------------|---------------------------|
| scripts                          | database scripts          |
| changelogs\vX.X.X\database_setup | create schema, roles, etc |
| changelogs\vX.X.X\permissions    | Set user permissions      |
| changelogs\vX.X.X\structure      | Define foreign keys       |
| changelogs\vX.X.X\tables         | schema tables             |
| changelogs\vX.X.X\views          | table views               |
