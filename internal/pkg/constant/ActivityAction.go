package constant

const ACTIVITY_ACTION_CREATE = "create"
const ACTIVITY_ACTION_UPDATE = "update"
const ACTIVITY_ACTION_DELETE = "delete"
const ACTIVITY_ACTION_GENERAL = "general"

type ActivityAction struct{}

func (srv ActivityAction) OptionCodeNames() []string {
	return []string{
		ACTIVITY_ACTION_CREATE,
		ACTIVITY_ACTION_UPDATE,
		ACTIVITY_ACTION_DELETE,
		ACTIVITY_ACTION_GENERAL,
	}
}
