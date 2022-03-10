# School Bus Project API Document V0.01

## Notice

1. Every  interface has a response parameter : "error", it has "String" type Value

2. All error detail will be return by ZH-CN at this parameter
3. Every response Content-Type : "application/json"
4. Under this , every "Content-Type" Field means request body content-type
5. Tag “Logon Required” API work for “Mini Program”
6. Tag “Password Required” API work for “Embedded System”
7. Tag “Admin Required” API work for “Management System”

## Feedback API Group

### Feed Back

> Logon Required

URL : /feedback

Function : POST

Content-type : JSON

Request Parameter :

| Key  | Required | Type   | Note          |
| :--- | :------- | :----- | :------------ |
| type | true     | String | feedback type |
| text | true     | String | feedback body |

> Header Parameter:
>
> | key           | Required | Type   | Note                |
> | ------------- | -------- | ------ | ------------------- |
> | Authorization | true     | String | authorization token |

Response :

| Key   | Required | Type    | Note          |
| ----- | -------- | ------- | ------------- |
| error | true     | Strings | error message |

## User API Group

### Logon (code2session)

URL : /user/session

Function : POST

Content-Type : JSON

Request Parameter : 

| Key  | Required | Type   | Note        |
| :--- | :------- | :----- | :---------- |
| code | true     | String | Wechat Code |

Response : 

| Key     | Required | Type   | Note          |
| ------- | -------- | ------ | ------------- |
| error   | true     | String | Error Message |
| user_id | true     | String | User ID       |

> Header Parameter:
>
> | key           | Type   | Note                |
> | ------------- | ------ | ------------------- |
> | Authorization | String | Authorization token |

### Accept Agreement

> Logon Required

URL : /user/agreement

Function : GET

Content-Type : Query string

Request Parameter:

None

Response Parameter:

| Key   | Required | Type   | Note          |
| ----- | -------- | ------ | ------------- |
| error | true     | String | error message |

### Data Upload

> Logon Required

URL : /user/data

Function : POST

Content-Type : JSON

Request Parameter : 

| Key        | Required | Type   | Note                   |
| :--------- | :------- | :----- | :--------------------- |
| nick_name  | false    | String | nickname               |
| student_id | false    | String | user student ID number |

> Header Parameter:
>
> | key           | Required | Type   | Note                |
> | ------------- | -------- | ------ | ------------------- |
> | Authorization | true     | String | authorization token |

Response : 

| Key   | Required | Type   | Note          |
| ----- | -------- | ------ | ------------- |
| error | true     | String | Error Message |

### Data Get

> Logon Required

URL : /user/data

Function : GET

Content-type : JSON

Request Parameter :

> Header Parameter:
>
> | key           | Required | Type   | Note                |
> | ------------- | -------- | ------ | ------------------- |
> | Authorization | true     | String | authorization token |

Response :

| Key        | Required | Type   | Note                   |
| ---------- | -------- | ------ | ---------------------- |
| error      | true     | String | error message          |
| nick_name  | true     | String | nickname               |
| student_id | true     | String | user student ID number |

## Bus API Group

### Logon

> Admin Required

URL : /bus/logon

Function : POST

Content-Type : JSON

Request Parameter :

| Key    | Required | Type   | Note       |
| ------ | -------- | ------ | ---------- |
| key    | true     | String | bus key    |
| status | true     | String | bus status |

Response Parameter :

| Key   | Required | Type   | Note          |
| ----- | -------- | ------ | ------------- |
| error | true     | String | Error message |
| id    | true     | String | bus id        |

### Logout

> Admin Required

URL : /bus/logout/**{id}**	

Function : POST

Content-Type : Query string

Request Parameter :

| Key  | Required | Type   | Note     |
| ---- | -------- | ------ | -------- |
| id   | true     | String | bus code |

Response Parameter :

| Key   | Required | Type   | Note          |
| ----- | -------- | ------ | ------------- |
| error | true     | String | error message |

### Offline

> Admin Required

URL : /bus

Function : PUT

Content-Type : JSON

Request Parameter :

| Key    | Required | Type   | Note       |
| ------ | -------- | ------ | ---------- |
| id     | true     | String | bus code   |
| status | true     | String | bus status |

Response Parameter :

| Key   | Required | Type   | Note          |
| ----- | -------- | ------ | ------------- |
| error | true     | String | error message |

### Position Upload 

> Password Required

URL : /bus/position

Function : POST

Content-Type : JSON

Request Parameter :

| Key      | Required | Type   | Note         |
| -------- | -------- | ------ | ------------ |
| id       | true     | String | bus code     |
| position | true     | Dict   | bus position |

> position sample :
> ``` { latitude : "xxx", longitude : "xxx"} ```

Response Parameter:

| Key   | Required | Type   | Note          |
| ----- | -------- | ------ | ------------- |
| error | true     | String | error message |

### Seat Upload

> Password Required

URL : /bus/seat

Function : POST

Content-Type : JSON

Request Parameter:

| Key  | Required | Type   | Note       |
| ---- | -------- | ------ | ---------- |
| id   | true     | String | bus code   |
| seat | true     | Int    | empty seat |

Response Parameter:

| Key   | Required | Type   | Note          |
| ----- | -------- | ------ | ------------- |
| error | true     | String | error message |

### Get All Bus

> Logon Required

URL : /bus/all

Function : GET

Content-Type : Query string

Request Parameter :

> Header Parameter:
>
> | key           | Required | Type   | Note                |
> | ------------- | -------- | ------ | ------------------- |
> | Authorization | true     | String | authorization token |

Response Parameter:

| Key   | Required | Type   | Note          |
| ----- | -------- | ------ | ------------- |
| error | true     | String | error message |
| buses | true     | List   | all bus list  |

> station sample :
> ``` [{ id : "xx",  : "xx", position : { latitude : "xx",,longitude : "xx"} }, ...] ```

## Station API Group

### Login 

> Admin Required

URL : /station

Function : POST

Content-Type : JSON

Request Parameter :

| Key      | Required | Type   | Note             |
| -------- | -------- | ------ | ---------------- |
| id       | true     | String | station code     |
| status   | true     | String | station status   |
| area     | true     | String | station area     |
| position | true     | Dict   | station position |

>position sample :
>`{ latitude : "xxx", longitude : "xxx"}`

Response Parameter :

| Key   | Required | Type   | Note          |
| ----- | -------- | ------ | ------------- |
| error | true     | String | error message |

### Get All Station

> Logon Required

URL : /station/all

Function : GET

Content-Type : JSON

Request Parameter :

> Header Parameter:
>
> | key           | Required | Type   | Note                |
> | ------------- | -------- | ------ | ------------------- |
> | Authorization | true     | String | authorization token |

Response Parameter:

| Key     | Required | Type   | Note             |
| ------- | -------- | ------ | ---------------- |
| error   | true     | String | error message    |
| station | true     | List   | all station list |

> station sample :
> ``` [{ id : "xx", area : "xx", position : { latitude : "xx",,longitude : "xx"} }, ...] ```

### Queue Append

> Logon Required

URL : /station/ **{id}** 

Function : POST

Content-Type : Query string

Request Parameter:

> Header Parameter:
>
> | key           | Required | Type   | Note                |
> | ------------- | -------- | ------ | ------------------- |
> | Authorization | true     | String | authorization token |

Response Parameter:

| Key   | Required | Type   | Note          |
| ----- | -------- | ------ | ------------- |
| error | true     | String | error message |

### Queue Pop

> Logon Required

URL : /station/ **{id}** 

Function : DELETE

Content-Type : Query string

Request Parameter:

> Header Parameter:
>
> | key           | Required | Type   | Note                |
> | ------------- | -------- | ------ | ------------------- |
> | Authorization | true     | String | authorization token |

Response Parameter:

| Key   | Required | Type   | Note          |
| ----- | -------- | ------ | ------------- |
| error | true     | String | error message |