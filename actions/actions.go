package actions

import "errors"

// GetAction a factory method for Action
func GetAction(args map[string]interface{}) (interface {
	Action
}, error) {
	if args["start"].(bool) != false {
		if args["server"].(bool) != false {
			return &Server{args["<name>"].(string)}, nil
		} else if args["client"].(bool) != false {
			return &Client{args["<name>"].(string)}, nil
		}
	}
	return nil, errors.New("Unknown action")
}
