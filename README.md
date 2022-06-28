# Pomodoro 
## Features:
- [ ] set time frame for work & break 
  - at the begining just use 4x25min á 5min break after 4phases 15min break
- [ ] Add what topic/job you're going to do
- [x] set pomodoro phases
- [x] UI in cli
- [x] pop up or sound when finished
  - sound works
- [x] create Logs that saves individual times and overall times when pomodoro was running

## ToDo:
### High Prio:

- **Settings File** for
  - Phase & Break duration
  - which sound
  - Styling (color...) in UI 
### Low Prio:
  - **Frames** probably not needed to create as array
  - **Logger** - ~~Use CLI timeframe argument to add to TOTAL TIME instead fixed 25~~
  - --
  - **Logger** - Create Weekly logs -> smaller files; exec: open file -> read File -> Get Runs -> write complete new file or append??
  - use an init function in case logPhases.txt does not exists and/ or has no initial values
  - --
---
## Tipps while developing:
use branches for new features [BRANCHES:](https://git-scm.com/book/en/v2/Git-Branching-Branches-in-a-Nutshell)

---
# How to use: 
Pause/Resume in the cli: ctrl + s / ctrl + q
