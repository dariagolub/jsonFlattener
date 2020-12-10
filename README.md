# JSON flattener

JSON flattener is a command line tool written in Golang. Tool takes a JSON object as input and outputs a flattened version of the JSON object, with keys as the path to every terminal value in the JSON structure. 

Input example:

```json
{
    "a": 1,
    "b": true,
    "c": {
        "d": 3,
        "e": "test"
    }
}
```
The output for the above input would be:

```json
{
    "a": 1,
    "b": true,
    "c.d": 3,
    "c.e": "test"
}
```
## Assumptions

* The input should be a JSON object
* All keys named in the original object will be simple strings without ‘.’ characters
* The input JSON will not contain arrays

# How to run

## Prerequisits

  - Go should be installed. Follow this [guide](https://golang.org/doc/install) for installation

## Run

Use the following command:

```sh
$ cat {filePath} | go run main.go
```

Example of how to run the tool:
```sh
$ cat resources/testData.json | go run main.go
```

# Improvement points
* add scenarios for errors with mocking stdin and stdout with errors
