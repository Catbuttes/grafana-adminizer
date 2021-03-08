# Grafana Adminizer

This is a tool to boost any user up to Grafana server admin. Created because I am a dumbass and removed the admin user by mistake.

## Usage

Usage is pretty simple

```
$ grafana-adminizer --database grafana.db --user username
```

Options are as follows

```
Grafana Adminizer
For when you have lost your server admin account

  -database database
        The grafana database to update (default "/var/lib/grafana/grafana.db")
  -help
        Prints this message
  -user login
        The login of the user you wish to promote (default "admin")
```