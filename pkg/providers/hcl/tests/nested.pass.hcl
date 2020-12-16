flow "echo" {
    input {
        payload "object" {}
    }

    resource "get" {
        request "getter" "Get" {
            payload {
                message "nested" {
                    name = "{{ input:nested.name }}"

                    message "sub" {
                        message = "hello world"
                    }
                }
            }
        }
    }

    output {
        payload "object" {
            message "nested" {
                name = "<string>"
            }
        }
    }
}
