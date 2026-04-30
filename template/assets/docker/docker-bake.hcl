variable "VITE_API_URL" {
  default = ""
}

group "default" {
  targets = [
    "backend",
    "frontend",
  ]
}

target "backend" {
  context    = "."
  dockerfile = "apps/backend/Dockerfile"
  tags = ["backend"]
}

target "frontend" {
  context    = "."
  dockerfile = "apps/frontend/Dockerfile"
  tags = ["frontend"]

  args = {
    VITE_API_URL = VITE_API_URL
  }
}
