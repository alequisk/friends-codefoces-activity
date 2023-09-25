# Friends Codeforces Activity

FCA is a notifier that informs you when your friends submit problems on the Codeforces platform. The idea behind this tracker is to gain insight into how active your peers are in competitive programming. This way, you can be notified of new submissions while using your computer as usual.

## How does it work?

The program makes periodic requests to the [API do codeforces](https://codeforces.com/apiHelp) for your friends' handles and checks if there have been any changes in the list of the most recent submissions. If there is a new submission, the program will send a notification to everyone who will be coding during that time interval.

<p align="center">
  <img src="./image-notification.png" alt="Example notification" width="700">
</p>

## Getting Started

To start using the program, you must have Go installed in version [v1.21.1](https://go.dev/doc/install) or higher. Clone or download this projects for you computer and start the configuration.

In the project folder download the dependencies by running the following command in the terminal:

```sh
  go mod tidy
```

If you're on Windows, you can run `.\install\run.ps1 [handle list separeted by space]` to build and run the project monitoring the handles described.

If you're on Linux, you can build the project run `go buil ./cmd/main.go` and after you can run the app .`/main [handle list separeted by space]` for starting monitoring your friends. (Script to run will be code soon)

You can also change the interval for each new check in the `cmd/main.go` file by modifying the value of the `period_for_check` variable.

## Limitations

Due to Codeforces API restrictions to prevent DDoS attacks, the program is limited to fetching submissions for handles every 4 seconds since the API's default limit is 1 request every 2 seconds. This prevents parallelism in the calls and may cause delays in checking if your friends list is very large.


## About

I'm not an expert in Golang, but this is a simple project that has been useful for both practice and constant use. The notifications are subtle, so they don't disrupt your workflow when a new notification arrives. The program is also quite small (around 10MB for the binary file) and has very low resource consumption. In some tests, it peaked at 4.5MB of RAM usage and decreased after the Garbage Collector kicked in.

Any suggestions or code improvements can be made through Pull Requests, and I'll be willing to incorporate them if they align with the program's purpose.
