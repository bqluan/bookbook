bookbook
======
bookbook backend

### Build dev env ###
```
docker build -t bookbook-dev .
```

### Dev shell ###
```
docker run --rm -it --network bookbook_default -v `pwd`:/go/src/github.com/bqluan/bookbook bookbook-dev bash
```
