Setup Go Fiber Backend

1) Install go binary, found at https://go.dev/dl/
2) Setup path environment variable on your laptop to point to Go binary file path
3) Create the directory where you want to save your project.
4) Inside the project directory, using terminal/command prompt
   a) Create go module using "go mod init <modulename>"
        - "go.mod" file will be created.
        This is where you manage your dependencies for your project
   b) Install Go Fiber dependencies using "go get github.com/gofiber/fiber/v2" in
        - You should see go.mod file populated with Fiber dependencies
5) Create "main.go" file in the project directory
   - You can now write server code to the file.
     Example fiber code can be found at the official doc [See References]
     or you can check out the branches


# References:
1) Official doc - https://docs.gofiber.io/
2) Fiber GitHub - https://github.com/gofiber/fiber
3) Fiber Middleware GitHub - https://github.com/gofiber/awesome-fiber


# General
1) LiveReloading - Air [https://github.com/air-verse/air]

2) Sample GoLang project
   a) gofiber/boilerplate
   b) thomasvvugt/fiber-boilerplate
   c) Youtube - Building a REST API using Gorm and Fiber
   d) embedmode/fiberseed


# Cloud Library Support for GoLang
1) Google Cloud Platform (Google Cloud, App Engine, Cloud Run, Cloud Functions)
Google Cloud Client Libraries - https://github.com/googleapis/google-cloud-go

2) Amazon Web Services
AWS SDK Go v1 - https://github.com/aws/aws-sdk-go [MAINTENANCE MODE]
AWS SDK Go v2 - https://aws.github.io/aws-sdk-go-v2/docs/
    E.g. AWS S3 - https://aws.github.io/aws-sdk-go-v2/docs/making-requests/