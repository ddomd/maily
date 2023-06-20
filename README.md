# 📨 Maily

<br>

Maily is a simple GRPC/REST mailing list management micorservice created to learn more about google's GRPC protocol.


## Schema

The service serves 5 REST endpoints and 5 RPC.

### REST:

| METHOD | ROUTE                        | PAYLOAD                                                     | ACTION                                            |
| ------ | ---------------------------- | ----------------------------------------------------------- | ------------------------------------------------- |
| GET    | /rest/get/{id}               | {}                                                          | Returns a single email entry                      |
| GET    | /rest/batch/{limit}&{offset} | {}                                                          | Returns a batch of email entries                  |
| POST   | /rest/create                 | { "email_address": "example@example.com" }                  | Creates a new email entry                         |
| PUT    | /rest/update/{id}            | { "opt_out": bool }                                         | Updates the subscription status of a single entry |
| DELETE | /rest/delete/{id}            | {}                                                          | Deletes a single email entry                      |

### GRPC:

| CALL            | PAYLOAD                                                     | ACTION                                            |
| --------------- | ----------------------------------------------------------- | ------------------------------------------------- |
| /GetEmail       | { "id": int64 }                                             | Returns a single email entry                      |
| /GetBatchEmails | { "limit": int32, "offset": int32 }                         | Returns a batch of email entries                  |
| /CreateEmail    | { "email_address": "example@example.com" }                  | Creates a new email entry                         |
| /UpdateEmail    | { "id": int64, "opt_out": bool }                            | Updates the subscription status of a single entry |
| /DeleteEmail    | { "id": int64 }                                             | Deletes a single email entry                      |