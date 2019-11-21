# GO-API-HACK



## Blog Posts - More Information About This Repo

You can find more information about this project/repository and how to use it in following bl
- [Setting Up Swagger Docs for Golang API](https://towardsdatascience.com/setting-up-swagger-docs-for-golang-api-8d0442263641?source=friends_link&sk=224695e91ff6867e3718095f85880b79)


### Setting Up
- Replace All Occurrences of `eftakhairul/go-api-hack/cmd` with your username repository name
- Replace All Occurrences of `blueprint` with your desired image name


### Adding New Libraries/Dependencies
```bash
go mod vendor
```



### Setup CodeClimate
- Go to <https://codeclimate.com/github/repos/new>
- Add Repository
- Go to Test Coverage Tab
- Copy Test Reporter ID
- Go to Travis and Open Settings for Your Repository
- Add Environment Variable: name: `CC_TEST_REPORTER_ID`, value: _Copied from CodeClimate_

## Swagger

- Application uses [gin-swagger](https://github.com/swaggo/gin-swagger).
- To generate/update docs use `swag init` (from `/go-project-blueprint/cmd/blueprint`)
- You can find generated docs in `docs` package

To view docs, navigate to <http://localhost:1234/swagger/index.html> or to <http://localhost:1234/swagger/doc.json> for raw _JSON_
