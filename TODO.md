# TODOS

#### Urgent
* Change SetupShell for PTYMode 
* Create Pipe Mode
* Setup the basis for TUI Mode

---

* Fix shell unparsed characters
  * Improve SSH connection
  * Clean and comment the code

* Create the fs core code (sshfs)
* DO a research about TUI 

## Layer based architecture

* Transport 
* Session
* TUI
* Plugins (Maybe)

Everything above TUI layer (inclusive) will have access 
to an Event Bus to know triggered events

Events interface:
```
EVENT {
  name    string,
  key     string,
  payload byte[],
}
```
