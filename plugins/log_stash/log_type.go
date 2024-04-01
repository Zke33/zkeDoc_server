package log_stash

import "encoding/json"

type LogType int

const (
	LoginType   LogType = 1
	ActionType  LogType = 2
	RuntimeType LogType = 3
)

func (t LogType) String() string {
	switch t {
	case LoginType:
		return "loginType"
	case ActionType:
		return "actionType"
	case RuntimeType:
		return "runtimeType"
	}
	return ""
}

func (t LogType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}
