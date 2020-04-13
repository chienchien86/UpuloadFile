# API Example for Golangweb

## Install and Run
```shell
$ ./go run index.go
```
## Request limit
```
十秒內最多十個請求
```
## static
```
存放檔案的資料夾
```


## API Endpoint
- http://localhost:8080/filel
    - `GET`: get list of filename
    ```
    200 ok 
    {
    "data": [
        "test.png",
        "螢幕快照 2020-04-06 下午8.16.08.png"
    ],
    "message": "Successful",
    "status": true
    }
    ```
    
    - `DELETE`: remove file
    ```
    body - formdate
    {
        "filename": "XXX.XXX", // aaa.png
    }
    ```
    ```
    200 ok
    {
    "message": "delete Successful",
    "status": true
    }
    ```


    - `POST`: upload new file
    ```
    body - formdate
    {
       "filename"(file): "XXX.XXX", // aaa.png
    }
    ```
   ```
    200 ok 
    {
    "message": "Successful",
    "status": true,
    "uploadfilename": "螢幕快照 2020-04-09 下午5.16.39.png"
    }
    ```
    - `PUT`: rename file
    ```
    body - x-www-form-urlencoded
    {
        "oldfilename": "XXX.XXX", // aaa.png
        "newfilename": "XXX.XXX", // bbb.png
    }
    ```
    ```
    200 ok
    {
    "message": "rename file Successful",
    "status": true
    }       
    ```


