package hook

const (
	TABLE_NAME = 1

	BEFORE_QUERY = 2
	AFTER_QUERY  = 3

	BEFORE_UPDATE = 4
	AFTER_UPDATE  = 5

	BEFORE_DELETE = 6
	AFTER_DELETE  = 7

	BEFORE_INSERT = 8
	AFTER_INSERT  = 9
)

type Hooks struct {
	actionHooks map[string]interface{}
}
