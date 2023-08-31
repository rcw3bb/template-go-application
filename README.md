# Go Application Template

 A template for building go application.

## Usage

1. **Clone** this repository on your local.

2. Remove the **.git folder** from where you've cloned the repository.

3. Run the following command from where you've cloned the application:

   ```
   go run . -h
   ```

   > If it completes successfully your setup is good.
   
   Expect to see something similar to the following:
   
   ```
   go-app-template v1.0.0-SNAPSHOT
   
   Usage:
     go-app-template [flags]
     go-app-template [command]
   
   Available Commands:
     completion  Generate the autocompletion script for the specified shell
     hello       A hello to a name.
     help        Help about any command
   
   Flags:
         --debug                Show debug information
     -e, --environment string   Specify the environment
     -h, --help                 help for go-app-template
     -v, --version              Show the version
   
   Use "go-app-template [command] --help" for more information about a command.
   ```

## Packages

### The cmd Package

The cmd package holds the implementation of commands that are available to the application. This is based on **[cobra](https://github.com/spf13/cobra/tree/main)**.

The **root.go file** is the main application command where all other sub commands must be registered. The specific instance where do the registration is the following:

```
rootCmd
```

A sample sub command implementation is available at hello.go. The following is a sample registration to the rootCmd instance by the helloCmd instance:

```go
func init() {
	rootCmd.AddCommand(helloCmd)
	flags := helloCmd.Flags()
	flags.StringVarP(&objHello.name, "name", "",
		"", "Specify the name")
}
```

### The config Package

The config package holds all the things related to configuration *(e.g. reading configuration file.)*. This is based on **[viper](https://github.com/spf13/viper)**.

In this template, the config package reads the configuration files in the **conf directory**. The directory contains the following files sample configuration files:

* application.properties
* application-test.properties

By default only the application.properties file is loaded. But when the test is passed in as environment using the environment flag *(e.g. **-e test** or **--environment test**)*, it will also load the application-test.properties. This will override the values found in application.properties. 

Use the following syntax to create an environment based configuration:

```
application-<ENV>.properties
```

> Where ENV is anything that can be used with **-e** or **--environment** flag.

#### The appinfo Package

The appinfo package holds the information about the application like the application name and version.

### The logger Package

The logger package provides the logging capability using the [zap](https://github.com/uber-go/zap). The default is **info level**. Use the **--debug flag** to change it to debug level.

The log file will be generated in the **logs directory**.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## [Build](BUILD.md)

## [Changelog](CHANGELOG.md)

## Author

* Ronaldo Webb
