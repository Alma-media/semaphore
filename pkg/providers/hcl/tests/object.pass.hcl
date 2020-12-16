flow "echo" {
    input {
        payload "object" {}
    }

    resource "get" {
        request "getter" "Get" {
            payload {
                array = [
                    "{{ input:id }}",
                    "{{ input:name }}",
                    "static",
                ]
            }
        }
    }

    output {
        payload "object" {
            object = {
                "message": "hello world",
                "meta": {
                    "id": "{{ getter:output }}"
                }
            }
        }
    }
}
