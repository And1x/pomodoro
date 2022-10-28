# Pomodoro 
## CLI-Tool to run a timer and track your tasks
- Timer runs like a progress bar in the CLI - bell rings after completion.
- Print stats on how many runs you've already completed e.g. on a daily or monthly basis.
- Tasks saved in monthly JSON files.
---
## How to get
- Download binary or build from source. 
## How to use
### Esample usage
- Run binary with default time(25min) and specified a task:
``` 
./pomodoro -t 'Learning Bitcoin from the Command Line chapter - ch. 6.1'
```
- Print daily runs:
```
./pomodoro -print d
```
- To Pause press: ctrl + s 
- To Resume press: ctrl + q

## Available flags:
- Duration with: `-d` arg`<int>`
- Task name: `-t` args`<string>`
- Print out: `-print` opts are:
  - `m` = monthly
  - `d` = daily

