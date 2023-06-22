# 📨 Maily

A GRPC/REST mailing list management micorservice.

## Description

Maily was created as a simple project to learn more about google's GRPC protocol implementation.
The service serves both legacy JSON via REST and protocol buffers via GRPC and supports server reflection.
The database uses offset pagination to efficiently provide entry batches. 

Configuration options are loaded from an environment file with structure:

```
JSON_PORT=":<port>"
GRPC_PORT=":<port>"
DBPATH="/path/to/db"
```

### Schema 

Considering the small size of the project it uses sqlite to build the database using the following schema:

| COLUMN       | TYPE                    | EXAMPLE             |
| ------------ | ----------------------- | ------------------- |
| id           | INTEGER                 | Primary Key         |
| email        | TEXT UNIQUE             | "email@address.com" |
| confirmed_at | INTEGER(UNIX TIMESTAMP) | 	1687436941         |
| opt_out      | INTEGER(BOOL)           |  1/0(true/false)    |

### REST Endpoints

| METHOD | ROUTE                                              | JSON PAYLOAD                               | ACTION                                                                      |
| ------ | -------------------------------------------------- | ------------------------------------------ | --------------------------------------------------------------------------- |
| GET    | /rest/get/{id}                                     | {}                                         | Returns a single email entry                                                |
| GET    | /rest/all                                          | {}                                         | Returns all email entries                                                   |
| GET    | /rest/subs                                         | {}                                         | Returns all subscribed email entries                                        |
| GET    | /rest/batch/limit={int}&offset={int}               | {}                                         | Returns a batch of email entries                                            |
| GET    | /rest/batch_subs/limit={int}&offset={int}          | {}                                         | Returns a batch of subscribed email entries                                 |
| POST   | /rest/create                                       | { "email_address": "example@example.com" } | Creates a new email entry                                                   |
| PUT    | /rest/update/{id}                                  | { "opt_out": bool }                        | Updates the subscription status of a single entry                           |
| DELETE | /rest/delete/{id}                                  | {}                                         | Deletes a single email entry                                                |
| DELETE | /rest/delete_unsub                                 | {}                                         | Deletes all unsubscribed email entries                                      |
| DELETE | /rest/delete_before                                | { "date": int(unix time) }                 | Deletes all email entries that unsubscribed before a given date             |

### GRPC Calls

| CALL                          | REQUEST                                                       | ACTION                                                                      |
| ----------------------------- | ------------------------------------------------------------- | --------------------------------------------------------------------------- |
| GetEmail                      | { "id": int64 }                                               | Returns a single email entry                                                |
| GetAll                        | {}                                                            | Returns all email entries                                                   |
| GetAllSubscribed              | {}                                                            | Returns all subscribed email entries                                        |
| GetBatch                      | { "limit": int32, "offset": int32 }                           | Returns a batch of email entries                                            |
| GetBatchSubscribed            | { "limit": int32, "offset": int32 }                           | Returns a batch of subscribed email entries                                 |
| CreateEmail                   | { "email_address": "example@example.com" }                    | Creates a new email entry                                                   |
| UpdateEmail                   | { "id": int64, "opt_out": bool }                              | Updates the subscription status of a single entry                           |
| DeleteEmail                   | { "id": int64 }                                               | Deletes a single email entry                                                |
| DeleteUnsubscribed            | {}                                                            | Deletes all unsubscribed email entries                                      |
| DeleteUnsubscribedBefore      | { "date": int64(unix time) }                                  | Deletes all email entries that unsubscribed before a given date             |