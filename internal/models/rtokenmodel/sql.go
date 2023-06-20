package rtokenmodel

// token varchar(40) not null constraint refresh_tokens_unique unique,

const (
	RefreshTokenTable               = "refresh_tokens"
	RefreshTokenUniqueConstraint    = "refresh_tokens_unique"
	ForeignKey_RefreshTokens__Users = "fk_refresh_tokens__users"
)

type rtoken_column string

func (rt rtoken_column) String() string {
	return string(rt)
}

func (rt rtoken_column) TableName() string {
	return RefreshTokenTable
}

type refresh_token_sql struct {
	UserId  rtoken_column
	Expired rtoken_column
	Token   rtoken_column
}

var SQL = refresh_token_sql{
	UserId:  "user_id",
	Expired: "expired",
	Token:   "token",
}
