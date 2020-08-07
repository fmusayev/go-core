# go-core

# 1. Environment Parser
You can now simple create env.yaml file with profile names (default, dev, local, etc.) 
```
default:
    URL: http://test.com
    PORT: 80

dev:
    PORT: 8080

    #database
    DB_HOST: <url>
    DB_PORT: 1234
    DB_NAME: dbname
    DB_USER: dbuser
    DB_PASS: dbpass
```
And you can simple load env variables by: \
`core.LoadEnv(profiles...)`, example `core.LoadEnv("dev")` \
Please note that `default` profile will be loaded automatically, no need to specify in LoadEnv as a parameter. \
\
By default method will try to load file with name `env.yaml`. If your file name is different then use method \
`core.LoadEnvFile(filename, profiles...)`

### Limitations
Currently we support only one level of depth
```
profilename:           # 1st level is profilename
    URL: <url>         # 2nd level is env key:value property
    DATABASE:  
        URL: <url>     # 3rd level not allowed
```