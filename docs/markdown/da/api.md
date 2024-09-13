Use any http client to fetch data from the 0G DA Scan.

## Common Error Messages
### Response code and message
An API call that encounters an error will return non-zero as its response code. The error message will be returned in the message field and the detailed reason for the error will be returned in the data field.

| Error code | Error message                                      |
|:-----------|:---------------------------------------------------|
| 0          | Success                                            |
| 1          | Invalid parameter,see Data for detailed error.     |
| 2          | Internal server error,see Data for detailed error. |
| 3          | Too many requests,see Data for detailed error.     |

e.g.
```json
{
"code": 2,
"message":"Internal server error",
"data":"No matching DA-submit record found, txSeq 1000000"
}
```
### Http status code
To distinguish backend service error and gateway error, we only use `200` and `600` as HTTP response status code:
- 200: success, or known business error, e.g. entity not found.
- 600: unexpected system error, e.g. database error, blockchain rpc error, io error.

## Rate Limit
Here are references for various API tiers and their rate limits.

| API Tier	   | Price            | 	Rate Limit                           |
|:------------|:-----------------|:-----------------------------------------|
| Free	       | $0               | 5 calls/second, up to 100,000 calls/day  |
| Standard	   | To be determined | 20 calls/second, up to 500,000 calls/day |
| Enterprise  | To be determined | Customize on demand                      |





