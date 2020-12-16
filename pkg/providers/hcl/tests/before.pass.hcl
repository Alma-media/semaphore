flow "mock" {
    before {
        resource "check" {
            request "com.semaphore" "Fetch" {
                payload {
                    key = "value"
                }
            }
        }

        resources {
            sample = "key"
        }
    }
}

proxy "mock" {
    before {
        resource "check" {
            request "com.semaphore" "Fetch" {
                payload {
                    key = "value"
                }
            }
        }

        resources {
            sample = "key"
        }
    }

    forward "" {}
}
