Postman Collection
==================================

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/23338632-9791196b-28ef-40e8-accd-ea988535b9c5?action=collection%2Ffork&collection-url=entityId%3D23338632-9791196b-28ef-40e8-accd-ea988535b9c5%26entityType%3Dcollection%26workspaceId%3D375a5178-dc50-49c8-9ed3-5295f7908d43#?env%5Bnode-go-mono-app%5D=W3sia2V5Ijoibm9kZSIsInZhbHVlIjoiaHR0cDovL2xvY2FsaG9zdDo1MDAxLyIsImVuYWJsZWQiOnRydWUsInR5cGUiOiJkZWZhdWx0Iiwic2Vzc2lvblZhbHVlIjoiaHR0cDovL2xvY2FsaG9zdDo1MDAxLyIsInNlc3Npb25JbmRleCI6MH0seyJrZXkiOiJnbyIsInZhbHVlIjoiaHR0cDovL2xvY2FsaG9zdDo1MDAyLyIsImVuYWJsZWQiOnRydWUsInR5cGUiOiJkZWZhdWx0Iiwic2Vzc2lvblZhbHVlIjoiaHR0cDovL2xvY2FsaG9zdDo1MDAyLyIsInNlc3Npb25JbmRleCI6MX0seyJrZXkiOiJ0b2tlbiIsInZhbHVlIjoiZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SnVZVzFsSWpvaWMybHRjMmx0SWl3aWNHaHZibVVpT2lJd09EazVNVEl6TVRFeElpd2ljbTlzWlNJNkltRmtiV2x1SWl3aWFXRjBJam94TmpjNU5UWTROREE0TENKbGVIQWlPakUyTnprMU56VTJNRGg5LkdSLXZxOUJSNXk2RzNSS0tLaTJKXzF2aDktVkdUUFJTX2wzUnlXRkIwQ2siLCJlbmFibGVkIjp0cnVlLCJ0eXBlIjoic2VjcmV0Iiwic2Vzc2lvblZhbHVlIjoiZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SnVZVzFsSWpvaWMybHRjMmx0SWl3aWNHaHZibVVpT2lJd09EazVNVEl6TVRFeElpd2ljbTlzWlNJNkltRmtiV2x1SWl3aWFXRjBJam94TmpjNU5UWTROREE0TENKbGVIQWkuLi4iLCJzZXNzaW9uSW5kZXgiOjJ9XQ==)

Getting Started
==================================

Welcome to my monorepo Golang and NodeJs App! This is a simple application that consists of a Golang and a NodeJs.

Prerequisites
-------------

Before you begin, you will need to have the following installed on your system:

*   Golang
*   NodeJs and npm
*   Docker
*   Makefile

Installation
------------

To install App, follow these steps:

1.  Clone this repository from GitHub:
    
    ```bash
    $ git clone https://github.com/hilmanfjrsy/node-go-mono-app.git
    ```

2.  Change into the directory:
    
    ```bash
    $ cd node-go-mono-app
    ```

Running with docker
------------------------------

-   Run `docker-compose` using `makefile` with the following command:
    ```bash
    # Run and build container and image docker 
    $ make fetch-up
    $ make auth-up
    #This will start the `fetch-app` and it will be accessible at http://localhost:5002 for the `auth-app` it will be accessible at http://localhost:5001

    # Down container and image docker
    $ make fetch-down
    $ make auth-down

    # Start docker container
    $ make fetch-start
    $ make auth-start

    # Stop docker container
    $ make fetch-stop
    $ make auth-stop
    ```
Running locally
------------------------------
### Auth App
1. Navigate to the `auth-app` directory.
2. Install dependencies using this command
    ```bash
    $ npm install
    ```
3. Run the following command to run the `auth-app`:
    ```bash
    $ npm start
    ```
    This will start the `auth-app` and it will be accessible at http://localhost:5001.

### Fetch App
1. Navigate to the `fetch-app` directory.
2. Install dependencies using this command
    ```bash
    $ go mod download
    ```
3. Run the following command to run the `fetch-app`:
    ```bash
    $ go run main.go
    ```
    This will start the `fetch-app` and it will be accessible at http://localhost:5002.

Running Unit Test
------------------------------
Run unit test for `fetch-app` with the following command:
```bash
$ make test
```

## Diagram C4 Model

![C4 Diagram](c4.png)