package sqlutils

type Column interface {
	TableName() string
	String() string
}

func Full(c Column) string {
	return c.TableName() + "." + c.String()
}
