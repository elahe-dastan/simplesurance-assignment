# [simplesurance](https://www.simplesurance.com/) by Allianz Golang Coding Challenge

## Description

Using only the standard library, create a Go HTTP server that on each request responds with a
counter of the total number of requests that it has received during the previous 60 seconds
(moving window). The server should continue to return the correct numbers after restarting it, by
persisting data to a file.
When you’re done, please share your solution with me, and I’ll ask our Tech Team to review it.

- **Deadline**: Feel free to do it at your own pace, but, to give you an overview, candidates usually
  take up to 3 days to complete the task.
- **Submission**: We accept your task via GitHub, GitLab or any Git-based Software. Please also
  share how much time was needed to finish the task.
- **Review**: We ask for 2 to 3 days to review your results.

## How to Build and Run

This guide assumes you have Go installed on your system. If not, please [install Go](https://golang.org/doc/install)
first.

### Build

1. Open a terminal.
2. Navigate to the directory containing the source code of the project.
   If the source code is in your current directory, you can skip this step.
3. Compile the project by executing the following command:

```bash
go build
```

This will build the executable from the source files. If successful, an executable file
will be created in the current directory.

### Run

After building the project,
you can run the executable directly from the command line:

```bash
./simplesurance-assignment
```

If you are using Windows, you can run the executable without the `./`:

```powershell
simplesurance-assignment.exe
```

## Implementation

There are different algorithms to implement with their own pros and cons. The simplest approach is implementing
a deque in which I was concerned about the size of the deque that could grow a lot based on number of requests and also
though the amortized time complexity will be O(1) but some requests may take much longer than the others so, I
implemented another algorithm which I'm going to explain besides the rest of code and take a look at why each decision
was made.

In the current implementation we are using a fixed-window-size array in which each room indicates the number of requests
queried in that second (I descreteized the time by seconds, if this code was going to be used in a use like
microprocessor we had to change our idea). To return the response we simply sum up the numbers in this array. To write
each request to this array we follow below precedure.

```mermaid
graph TD;
    NewRequest("new request") -->|if time of the request is in the first server up time window| ArrayTimeOfRequest("array[time of request % window]++");
    NewRequest("new request") -->|if time of the request has passed the first server up time window| Empty("empty expired cells");
    Empty("empty expired cells") --> ArrayTimeOfRequest("array[time of request % window]++");
    
    Empty("empty expired cells\n cells to remove = (time of new request) - (time of last request)") -->| cells to remove > window| ReCreate("re-create the array");
    Empty("empty expired cells\n cells to remove = (time of new request) - (time of last request)") -->|(time of new request) - (time of last request) < window| CalculateExpired("time of new");
```

This solution has constant memory and time completexity which is good.
For each request we locked the array because we want to read and also
make some rooms sets to zero (so we don't use `RWLock` here).

Writing into a file using the current solution is easy. We write
the array using JSON format into a file which is named `state.json` by default.
Then read the file in the start phase and if there was any error we fallback to use
any empty state.

Writing to file is time consuming, so we start another goroutine which writes the file
priodically (period is configurable) by locking the counter which seems more configurable and have better performance.
