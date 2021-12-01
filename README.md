<h1> <img src="https://covergates.com/logo.png" alt="logo" width="48" height=48> Covergates - Portal Gates to Coverage Reports</h1>

[![badge](https://covergates.com/api/v1/reports/bsi5dvi23akg00a0tgl0/badge?)](https://covergates.com/report/github/laojianzi/covergates)
![CI](https://github.com/laojianzi/covergates/workflows/CI/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/laojianzi/covergates)](https://goreportcard.com/report/github.com/laojianzi/covergates)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/laojianzi/covergates)](https://pkg.go.dev/github.com/laojianzi/covergates)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![CLA assistant](https://cla-assistant.io/readme/badge/covergates/covergates)](https://cla-assistant.io/laojianzi/covergates)
[![Twitter Follow](https://img.shields.io/twitter/follow/covergates_tw.svg?style=social)](https://twitter.com/covergates_tw)

[![card](https://covergates.com/api/v1/reports/bsi5dvi23akg00a0tgl0/card)](https://covergates.com/report/github/laojianzi/covergates)

## Purpose

**Covergates** is to make the easiest way to setup a self-hosted coverage report service.
It's an alternative to services, such as:

- [Code Climate](https://codeclimate.com/)
- [Codecov](https://codecov.io/)
- [Coveralls](https://coveralls.io/)

The reason why this project is because managing coverage reports for private repositories should not be hard!
It is able to link with your self-hosted Git service.
Use it to improve coding review and quality management flow for your internal projects.
Want to try? Visit [covergates.com](https://covergates.com) before you starting.

## Using

To get started, please download prebuilt binary [covergates-**version**-**platform**-**architecture**.zip](https://github.com/laojianzi/covergates/releases) and try:

```sh
unzip covergates-<version>-<platform>-<architecture>.zip
./covergates-server
```

Visit [http://localhost:8080](http://localhost:8080) for your **covergates** service.

To upload report, run `covergate` cli:

```sh
export API_URL=http://localhost:8080/api/v1
covergates upload -report <report id> -type go coverage.out
```

## Configure

`covergates-server` uses environment variables to change configurations.
Below is the list of variables for basic configuration:

- `GATES_SERVER_ADDR` Default `http://localhost:8080`
- `GATES_SERVER_BASE` Default `/`
- `GATES_DB_DRIVER` Default `sqlite3`. Other options are `postgres` and `cloudrun`
- `GATES_DB_HOST` Required host for `postgres` and `cloudrun`
- `GATES_DB_PORT` Required port for `postgres` and `cloudrun`
- `GATES_DB_USER` Required user for`postgres` and `cloudrun`
- `GATES_DB_NAME` Required database name for `postgres` and `cloudrun`
- `GATES_DB_PASSWORD` Required password for `postgres` and `cloudrun`
- `GATES_GITEA_SERVER` Default `https://try.gitea.io/`, gitea server address
- `GATES_GITEA_CLIENT_ID` Required for Gitea OAuth login
- `GATES_GITEA_CLIENT_SECRET` Required for Gitea OAuth login
- `GATES_GITHUB_SERVER` Default `https://github.com`
- `GATES_GITHUB_API_SERVER` Default `https://api.github.com`
- `GATES_GITHUB_CLIENT_ID` Required for GitHub OAuth login
- `GATES_GITHUB_CLIENT_SECRET` Required for GitHub OAuth login

## Supported SCM and Language

| SCM       | Supported          |
| --------- | ------------------ |
| GitHub    | :heavy_check_mark: |
| Gitea     | :heavy_check_mark: |
| GitLab    | :heavy_check_mark: |
| Gogs      | :x:                |
| Bitbucket | :x:                |

| Language                  | Supported          | Tutorial                                               |
| ------------------------- | ------------------ | ------------------------------------------------------ |
| Go                        | :heavy_check_mark: | [go-example](https://github.com/covergates/go-example) |
| Perl                      | :heavy_check_mark: | :wrench:, ongoing                                      |
| Python                    | :heavy_check_mark: | :wrench:, ongoing                                      |
| Ruby (SimpleCov: RSpec)   | :heavy_check_mark: | :heavy_minus_sign:                                     |
| lcov (C, C++, Javascript) | :heavy_check_mark: | :heavy_minus_sign:                                     |
| Clover (PHP)              | :heavy_check_mark: | :heavy_minus_sign:                                     |
| Java (Jacoco)             | :wrench:, ongoing  | :heavy_minus_sign:                                     |

**Covergates** is at an early development stage.
Other languages and SCM support is ongoing!
If you would like to assist with development, please refer to [Contributing Section](#contributing).

## Development

The build is split into `backend`, `cli` and `frontend`. To build backend, run:

```sh
go build -o covergates-server ./cmd/server
```

To build CLI, run:

```sh
export SERVER_API_URL=http://localhost:8080/api/v1
go build -o covergates -ldflags="-X main.CoverGatesAPI=$SERVER_API_URL" ./cmd/cli
```

You may change `SERVER_API_URL` to your self-hosted **covergates-server** address.

If your are behind firewall or proxy,
you may also download source package with `vendor` modules from [covergates.**version**.src.zip
](https://github.com/laojianzi/covergates/releases). To build with `vendor` modules, run:

```
go build -o covergates-server -mod vendor ./cmd/server
```

To build frontend, it requires:

1. [Node.js v12](https://nodejs.org/en/download/)
2. [togo](https://github.com/bradrydzewski/togo)

Read [web/README.md](https://github.com/laojianzi/covergates/blob/main/web/README.md) for more details.

## Contributing

It would be highly appreciated if you could contribute to the project.
There are many ways in which you can participate in the project:

1. Contributing directly to the code base

   The expected workflow is [GitHub flow](https://guides.github.com/introduction/flow/).
   Read [CONTRIBUTING.md](https://github.com/laojianzi/covergates/blob/main/CONTRIBUTING.md) before getting start.

2. [Submit feature requests and bugs](https://github.com/laojianzi/covergates/issues)

   Especially for the new language support.
   It would be great if you could provide coverage report examples and how to produce coverage for other languages.

3. Testing, both unit testing and e2e testing are welcome.

## Further Information

For more information and tutorial about self-hosted Covergates server, please refer to our [documentation](https://docs.covergates.com/)

## Milestones

Refer to [TODO.md](https://github.com/laojianzi/covergates/blob/main/TODO.md) for details.

## License

This project is licensed under the GNU General Public License v3.0. See the [LICENSE](https://github.com/laojianzi/covergates/blob/main/LICENSE) file for the full license text.

## Screenshots

![report](https://raw.githubusercontent.com/covergates/brand/master/screenshots/covergates.png)

![files](https://raw.githubusercontent.com/covergates/brand/master/screenshots/covergates_code.png)

![setting](https://raw.githubusercontent.com/covergates/brand/master/screenshots/covergates_setting.png)

![pull request](https://raw.githubusercontent.com/covergates/brand/master/screenshots/covergates_pr.png)
