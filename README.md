# simple-loan-apps

# Setup your machine - Linux
1. Install golang - `https://go.dev/doc/install`
2. Install mysql - `https://www.digitalocean.com/community/tutorials/how-to-install-mysql-on-ubuntu-20-04`
3. Install docker - `sudo snap install docker`
4. Create your database and table - `docker-compose up` 
5. Run your application - `go build cmd/main.go && ./main`

# API Documentation
#### 1. Create New Loans - `POST - localhost:8900/loan/create`
###### Request
```
{
    "name": "alfin lie",
    "identityNumber": "5371041403950004",
    "dateOfBirth": "14-03-1995",
    "sex": "MALE",
    "amount": 1000000,
    "period": 6,
    "purpose": "beli car"
}
```
###### Response 
```
{
    "data": {
        "isSuccess": true
    },
    "error": null
}
```

#### 2. Get Loan Details - `GET - localhost:8900/loan/get-by-id/{id}`
###### Response 
```
{
    "data": {
        "ID": "5f606b7e-1bbc-11ed-a0df-0242ac130002",
        "UserID": 2,
        "Amount": 1000000,
        "Period": 6,
        "Purpose": "beli car",
        "CreatedAt": "2022-08-14T10:32:18Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": "0001-01-01T00:00:00Z"
    },
    "error": null
}
```
#### 3. Get Loans by KTP - `GET - localhost:8900/loan/get-by-ktp/{ktp}`
###### Response 
```
{
    "data": [
        {
            "ID": "5f606b7e-1bbc-11ed-a0df-0242ac130002",
            "UserID": 2,
            "Amount": 1000000,
            "Period": 6,
            "Purpose": "beli car",
            "CreatedAt": "2022-08-14T10:32:18Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": "0001-01-01T00:00:00Z"
        },
        {
            "ID": "90d76161-1bbc-11ed-a0df-0242ac130002",
            "UserID": 2,
            "Amount": 1000000,
            "Period": 6,
            "Purpose": "beli car",
            "CreatedAt": "2022-08-14T10:33:41Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": "0001-01-01T00:00:00Z"
        }
    ],
    "error": null
}
```
