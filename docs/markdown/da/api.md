Use any http client to fetch data from the 0G DA Scan.

#### Common Error Messages:
An API call that encounters an error will return non-zero as its status code. The error message will be returned in the message field and the detailed reason for the error will be returned in the data field.
```
{
"code": 1001,
"message":"No matching records found",
"data":"DA tx, txSeq 1000000"
}
```
| error code | error message             |
|:-----------|:--------------------------|
| 0          | Success                   |
| 1          | Invalid parameter         |
| 2          | Internal server error     |
| 3          | Too many requests         |
| 1001       | No matching records found |
| 1002       | Blockchain RPC failure    |


