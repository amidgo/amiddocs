package rtokenmodel

// create table if not exists refresh_tokens (
//     user_id bigserial not null references users(id) on delete cascade,
//     expired timestamp not null,
//     token varchar(40) not null unique,
//     primary key(user_id)
// );

const RefreshTokenTable = "refresh_tokens"

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
