
resource "http-b64" "example" {
  provider = http2b64
  url      = "http://localhost/test.txt"
}