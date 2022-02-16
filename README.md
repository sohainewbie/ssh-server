
# ssh-connect
> simple server ssh application

## How to Run:

> go run main.go

> file config config.json
```
{
    "ssh" : {
        "host": "localhost",
        "port": 4555,
        "username": "kucingliar",
        "password": "kucingliar123",
        "hostKeyFile" : "id_rsa",
        "clientAuth" : true,
        "textDisplay": "Welcome to Server"
    }
}
```

# Testing connect to SSH Server
>  ssh kucingliar@localhost -p 4555 
```
kucingliar@localhost's password:
Permission denied, please try again.
kucingliar@localhost's password:
Welcome to Server
kucingliar$ ls -la
total 64
drwxr-xr-x  12 sohai  staff   384 Feb 11 10:11 .
drwxr-xr-x   5 sohai  staff   160 Feb 11 10:11 ..
drwxr-xr-x  12 sohai  staff   384 Feb 11 10:12 .git
-rw-r--r--   1 sohai  staff     6 Feb 11 10:11 .gitignore
-rw-r--r--   1 sohai  staff   291 Feb 11 10:11 Makefile
-rw-r--r--   1 sohai  staff    68 Feb 11 10:12 README.md
-rw-r--r--   1 sohai  staff   241 Feb 11 10:11 config.json
-rw-r--r--   1 sohai  staff   431 Feb 11 10:11 go.mod
-rw-r--r--   1 sohai  staff  3389 Feb 11 10:11 go.sum
-rw-r--r--   1 sohai  staff   887 Feb 11 10:11 id_rsa
-rw-r--r--   1 sohai  staff   578 Feb 11 10:11 main.go
drwxr-xr-x   5 sohai  staff   160 Feb 11 10:11 modules

kucingliar$
```


#server test:
> ubuntu@3.134.219.221

# Testing Local Slave server
```
cd docker/ 
docker build -t metal-slave .
docker run --name metal_slave_1 -d -p 2201:22 metal-slave:latest
```

# Testing Local SFTP Slave Server
```
> sftp -P 2201 sohai@0.0.0.0
sohai@0.0.0.0's password:
Connected to 0.0.0.0.
sftp>
```

# Testing Local SSH Slave Server
```
> ssh sohai@0.0.0.0 -p 2201
sohai@0.0.0.0's password:
Permission denied, please try again.
sohai@0.0.0.0's password:
Welcome to Alpine!

The Alpine Wiki contains a large amount of how-to guides and general
information about administrating Alpine systems.
See <http://wiki.alpinelinux.org/>.

You can setup the system with the command: setup-alpine

You may change this message by editing /etc/motd.

4188dd09613d:~$ ls -la
total 24
drwxr-sr-x    1 sohai    sohai         4096 Feb 16 06:46 .
drwxr-xr-x    1 root     root          4096 Feb 16 06:05 ..
-rw-------    1 sohai    sohai           38 Feb 16 06:57 .ash_history
drwx--S---    2 sohai    sohai         4096 Feb 16 06:46 .ssh
drwxr-xr-x    1 sohai    sohai         4096 Feb 16 06:44 uploader
```

## User Default
| type  | username  | password     |
| ----- |:--------- |:------------:|
| Slave | sohai     | kucingliar   |
| Master| kucingliar| kucingliar123|