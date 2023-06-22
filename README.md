# 📨 Maily

<br>

## Description

Maily is a simple GRPC/REST mailing list management micorservice created to learn more about google's GRPC protocol.


## Schema

The service serves 10 REST endpoints and 10 RPC.

### REST:

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

### GRPC:

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