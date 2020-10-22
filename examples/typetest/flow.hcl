endpoint "typetest" "http" {
  endpoint = "/"
  method   = "POST"
  codec    = "json"
}

endpoint "typetest" "grpc" {
  package = "semaphore.typetest"
  service = "Typetest"
  method  = "Run"
}

flow "typetest" {
  input "semaphore.typetest.Request" {}

  // resource "users" {
  //   output
  // }

  output "semaphore.typetest.Response" {
    data = "{{ input:data }}"
    string = "{{ input:data.string }}"
  }
}
