# Go Fake API
It's simple. You create a response file and when the endpoint is triggered the example api response will be served. And the best thing: *You can fully customize it!*

##Table of Contents
* [Configuration](#configuration)
* [Compiling the API](#compiling-the-api)
* [Launch the API](#launch-the-api)
* [TODO](#todo)

### Configuration
Everything is written in the main.go file. There are some variables you can change:
* **formats**:
    * *default:* 
        `[2]*Format{
            {Name: "json", FileSuffix: "json", MIMEType: "application/json"},
            {Name: "xml", FileSuffix: "xml", MIMEType: "application/xml"},
        }`
    * *description:* Here you can add formats that the api should resolve.
    * *notice:* You have to fulfill al fields or it wont work.      
* **errorMessage**: 
    * *default:* `{\"error\":\"%s\"}`
    * *description:* This will be the message, which will be served when the program errors.
* **errorFormat**:
    * *default:* `formats[0]`
    * *description:* The format in which the message is. You can select all formats you have defined in the *formats* array.
* **path**:
    * *default:* `"endpoints/${url}.${format}"`
    * *description:+ You can define the endpoint where your fake api files are stored. For better customization you can use some parameters.
    * *parameters:*
        * `${url}` - The URL you hit in the browser.
        * `${format}` - The FileSuffix from the *formats* array.
        
### Compiling the API
When you have configured the API you can compile it by using the go compiler. 
`go build main.go`

If you want to launch the API with the default settings you find the newest compiled version [here](https://github.com/AnguliNetworks/go-fake-api/releases)
        
### Launch the API
When you have compiled you own version of the api or downloaded a release, you can run it from the command line
* UNIX `./main`
* Windows `main`

### TODO
[ ] Move the configuration to an extra file.
