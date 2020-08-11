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
	"type": true,
	"time": 11451441429366666,
	"content": "Some Suggestion content here",
    }
    ```
    - Comment
    `id` is randomly generate by server, it is a very long integer  
    `type` true means suggestion, false means reply  
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
		"type": false,
		"time": 100,
		"content": "give me more money!",
	    },
	    {
		"id": 39,
		"type": false,
		"time": 200,
		"content": "go f*uc@!#! yourself!",
	    },
	]
    }
