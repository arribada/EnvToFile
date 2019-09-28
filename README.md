A simple tool that reads all env variables and writes each one to a file.
Useful in environments where the only way to add configuratios is through env variables and
need to use docker images which accept configs only through files.

See the docker-compose file for a full example.


To write all current env variable and a custom `CONFIG` variable in the current dir:
```
VALUE=val CONFIG="var=$VALUE" go run main.go --dir=./
```
