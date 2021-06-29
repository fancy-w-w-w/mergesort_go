package main

/*import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var DbErr error
// 用户表结构体
type User struct {
	Id uint `db:"id"`
	Name string  `db:"name"`
	Age int `db:"age"`
}


func init(){
	Db,DbErr=sql.Open("mysql","root:123456@/test?charset=utf8")
	if DbErr!=nil{
		fmt.Println(DbErr)
	}

	defer Db.Close()
}

func StructQueryField() {

	user := new(User)
	row := Db.QueryRow("select id, name, age from mydb where id=?",1)
	if err :=row.Scan(&user.Id,&user.Name,&user.Age); err != nil{
		fmt.Printf("scan failed, err:%v",err)
		return
	}
	fmt.Println(user.Id,user.Name,user.Age)
}
func InsertQuryField(){

	ret,_:=Db.Exec("insert into mydb(id,name,price) values(1,'dog',100)")
	insID,_:=ret.LastInsertId()
	fmt.Println("插入成功，finalID：",insID)

	//影响行数
	rowsaffected,_ := ret.RowsAffected()
	fmt.Println("插入行数：",rowsaffected)

}
// 更新数据
func StructUpdate() {

	ret,_ := Db.Exec("UPDATE users set price=? where id=?","100",1)
	upd_nums,_ := ret.RowsAffected()
	fmt.Println("RowsAffected:",upd_nums)
}
// 删除数据
func StructDel() {

	ret,_ := Db.Exec("delete from users where id=?",1)
	del_nums,_ := ret.RowsAffected()

	fmt.Println("RowsAffected:",del_nums)
}

//预处理
func PrepareExec(){
	st,_:=Db.Prepare("select id,name from mydb where id=? or name=?")
	row,_:=st.Query(3,"dog")
	var id int
	var name string
	for row.Next(){
		row.Scan(&id,&name)
		fmt.Println(id,name)
	}
}

// 事务处理,结合预处理
func StructTx() {

	//事务处理
	tx, _ := Db.Begin();


	// 新增
	userAddPre, _ := Db.Prepare("insert into users(name, price) values(?, ?)");
	addRet, _ := userAddPre.Exec("zhaoliu", 15);
	ins_nums, _ := addRet.RowsAffected();


	// 更新
	userUpdatePre1, _ := tx.Exec("update mydb set name = 'zhansan'  where name=?", "张三");
	upd_nums1, _ := userUpdatePre1.RowsAffected();
	userUpdatePre2, _ := tx.Exec("update mydb set name = 'lisi'  where name=?", "李四");
	upd_nums2, _ := userUpdatePre2.RowsAffected();

	fmt.Println(ins_nums);
	fmt.Println(upd_nums1);
	fmt.Println(upd_nums2);

	if ins_nums > 0 && upd_nums1 > 0 && upd_nums2 > 0 {
		tx.Commit();
	}else{
		tx.Rollback();
	}

}*/







