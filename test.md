## MPQ can do create or update operations for print queue.
---
##### frontend sign off;
##### backend sign off: `Sun Alex`  
&emsp;
>## parameter
```
@host = host name
@port = host port
@qname = queue name
@tenant_id = tenant id 
@qdescription = queue description
@qstatus = queue status
@cleanupPrd = clean up period
@techUserName = technical user name
```
&emsp;
>## Request
```
PUT http://{{host}}:{{port}}/api/v1/rest/queues/{{qname}} HTTP/1.1
Content-Type: application/json
zid: {{tenant_id}}

{
    "qname": "{{qname}}",
    "qdescription": "{{qdescription}}",
    "qstatus": "{{qstatus}}",
    "qformat": "{{qformat}}",
    "cleanupPrd": "{{cleanupPrd}}",
    "techUserName": "{{techUserName}}"
}
```
&emsp;
>## Response
- When `techUserName` cannot be found in the table.
  >**Response Code:** 400  
  >**Response body:** User doesn't exist, please enter an existing user.


- When `qname` is empty or qname field is not provided in the body.
  >**Response Code:** 400  
  >**Response body:** Please enter a print queue name.

- When length of `qname` is more than 32 characters or not be A-Z, a-z, 0-9 and underscore characters.
  > **Response Code:** 400  
  > **Response body:** The queue name hasn't been validated. Only A-Z, a-z, 0-9 and underscore are valid.

- When `qdescription` is empty or qdescription field is not provided in the body
  > **Response Code:** 400  
  > **Response body:** Please enter a print queue description.

- When `qformat` is empty or qformat field is not provided in the body
  > **Response Code:** 400  
  > **Response body:** Please enter a print queue format.

- When `technicalUserName` is empty or technicalUserName field is not provided in the body
  > **Response Code:** 400  
  > **Response body:** Please enter a technical user name.

- When `cleanupPrd` is empty or cleanupPrd field is not provided in the body
  > **Response Code:** 400  
  > **Response body:** Please enter a clean-up period.

- When `qname` is exist in Print Queue but has been logical deleted
  > **Response Code:** 409  
  > **Response body:** Sorry, print queue already exists, please use another one.

- When `qname` does not exist in Print Queue
  > **Response Code:** 204  
  > **Response body:** null
