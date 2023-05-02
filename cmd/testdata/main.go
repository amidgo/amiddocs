package main

import (
	"context"
	"fmt"
	"os"

	"github.com/amidgo/amiddocs/pkg/postgres"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("./config/.env")
	p, _ := postgres.New(os.Getenv("DATABASEURL"))
	defer p.Close()
	dep_id := uint64(0)
	group_id := uint64(0)
	user_id := uint64(0)
	doc_id := uint64(0)
	err := p.Pool.QueryRow(
		context.Background(),
		`INSERT INTO departments (name, short_name) VALUES('test','test') RETURNING id`,
	).Scan(&dep_id)
	fmt.Printf("\n Department err is %v", err)
	err = p.Pool.QueryRow(
		context.Background(),
		`INSERT INTO groups 
		(name,is_budget,education_form,education_start_date,education_year,education_finish_date,department_id)
		VALUES ('test',true, 'FULL_TIME','2020-01-01',1,'2024-01-01',$1) 
		RETURNING id`, dep_id,
	).Scan(&group_id)
	fmt.Printf("\n Group err is %v", err)
	err = p.Pool.QueryRow(
		context.Background(),
		`INSERT INTO student_documents 
		(doc_number,order_number,order_date,study_start_date)
		VALUES ('10101','10101','2020-01-01','2020-01-01')
		RETURNING id`,
	).Scan(&doc_id)
	fmt.Printf("\n Document err is %v", err)
	err = p.Pool.QueryRow(
		context.Background(),
		`INSERT INTO users
		(name,surname,father_name,login,email,password)
		VALUES ('amidman','amidman','amidman','amidman','amidman@mail.ru','$2a$10$K/NPcXlswDBNr1QsJkT/rOIYPv2JAZbKxBBPTK4jjFX1dbX2gtxlm')
		RETURNING id`,
	).Scan(&user_id)
	fmt.Printf("\n User err is %v", err)
	_, err = p.Pool.Exec(
		context.Background(),
		`INSERT INTO user_roles (role_id,user_id) VALUES ((SELECT roles.id FROM roles WHERE roles.role = 'STUDENT'),$1), ((SELECT roles.id FROM roles WHERE roles.role = 'ADMIN'),$2)`, user_id, user_id,
	)
	fmt.Printf("\n Roles err is %v", err)
	_, err = p.Pool.Exec(
		context.Background(),
		`INSERT INTO students (user_id,student_document_id,group_id) VALUES($1,$2,$3)`,
		user_id, doc_id, group_id,
	)
	fmt.Printf("\n Student err is %v", err)
	fmt.Println(user_id, doc_id, group_id, dep_id)
}
