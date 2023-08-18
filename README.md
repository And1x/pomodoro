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

- Run binary with default time(25min) and a task:

```
./pomodoro -t 'Learning Bitcoin from the Command Line chapter - ch. 6.1'
```

![progress bar](https://user-images.githubusercontent.com/92379947/198682630-fe65e2f7-c8a7-4434-a4f8-5b4b8851e13f.png)

- List daily runs:

```
./pomodoro -l d
```

![print daily](https://user-images.githubusercontent.com/92379947/198682656-b2ae0833-d128-4f8c-a8e2-f0c69d3f3c5e.png)

- To Pause press: ctrl + s
- To Resume press: ctrl + q

## Available flags:

- Duration with: `-d` arg`<int>`
- Task name: `-t` args`<string>`
- Print out: `-l` opts are:
  - `m` = monthly
  - `d` = daily
  - `all` = all months of the year
  - `January` , `February` ... = specific Month
