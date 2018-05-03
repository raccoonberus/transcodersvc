# transcodersvc

Alternative free software, in contrast to Amazon Transcoder.

### RESR API description

| URI  | Method  | Params  | Description  |
|---|---|---|---|
| /file/  | POST  | file  | Upload file to server and return info about resource  |
| /task/  | POST  | filename  | Create conversion task  |
| /task/{id}  | GET  | --  | Get info about task  |

