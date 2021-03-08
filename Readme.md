# Grafana Adminizer

This is a tool to boost any user up to Grafana server admin. Created because I am a dumbass and removed the admin user by mistake. If anybody else find it useful, enjoy!

## Disclaimer

Use this at your own risk. I am not responsible if you mangage to break your Grafana install with this. _Make sure you have a backup of your grafana database before you run it._

## Usage

Usage is pretty simple

1. Stop your grafana instance
2. Locate the sqlite config database (It usually lives at `/var/lib/grafana/grafana.db`)
3. **Make a backup** of that database
4. Run the tool as below

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
