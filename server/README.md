# API Reference
- Add new suggestion
    - Method: `POST`
    - Request example
    ```json
    {
	"content": "Some suggestion content here",
    }
    ```
    - Response example
    ```json
    {
	"id": 13246234123,
	"type": 1,
	"time": 11451441429366666,
	"content": "Some Suggestion content here",
    }
    ```
    - Comment
    `id` is randomly generate by server, it is a very long integer  
    `type` 1 means suggestion, 2 means reply  
    `time` is an integer, which is UTC unix time
- Reply an suggestion
    - Method: `POST`
    - Request example
    ```json
    {
	"id": 1145144142936666666,
	"content": "Some reply content here",
    }
    ```
    - Response example
    ```json
    {
	"title": "Success",
	"content": "Action success",
    }
    ```
- Get suggestion list by id
    - Method: `POST`
    - Request example
    ```json
    {
	"id": 114514414566666666
    }
    ```
    - Response example
    ```json
    {
	"suggestion_list": [
	    {
		"id": 39,
		"type": 1,
		"time": 100,
		"content": "give me more money!",
	    },
	    {
		"id": 39,
		"type": 2,
		"time": 200,
		"content": "go f*uc@!#! yourself!",
	    },
	]
    }
