OpenInvoice
===========

Invoice system. You can create a Quotation and Invoice PDF with this system.

* Requirements

Docker and Docker compose(and internet connection).

* How to build docker image and run containers

step 1. build api server and web front.

```
$ cd go
$ build-with-container.sh
```

step 2. build docker images with `docker-compose`

```
$ docker-compose build
```

step 3. run containers

```
$ docker-compose run
```

step 4. Open `http://localhost:10080` in your browser.

step 5. Login with admin/admin (default password)

step 6(important!). update admin password!

* Use another port

modify `docker-compose.yaml`

```
    front:
        build: ./go
        image: fkmhrk/oiv:latest
        ports:
            - 10080:80 // here!
```




