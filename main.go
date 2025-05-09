package main

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID         int        `gorm:"primaryKey"`
	CreateTime *time.Time `gorm:"autoCreateTime"`
	UpdateTime *time.Time `gorm:"autoCreateTime"`
}

type Teacher struct {
	BaseModel
	Name   string `gorm:"type:varchar(32);unique;not null"`
	Tno    int
	Pwd    string `gorm:"type:varchar(100);not null"`
	Tel    string `gorm:"type:char(11)"`
	Birth  *time.Time
	Remark string `gorm:"type:varchar(255)"`
}

type Class struct {
	BaseModel
	Name string `gorm:"type:varchar(32);unique;not null"`
	Num  int
	// 将Teacher字段拼接ID作为teachers表的外键
	// Teacher Teacher `gorm:"foreignKey:TeacherID"`
	TeacherID int
	// 最终只有TeacherID字段
	Teacher Teacher `gorm:"constraint:onDelete:CASCADE"`
}

type Course struct {
	BaseModel
	Name      string `gorm:"type:varchar(32);unique;not null"`
	Credit    int
	Period    int
	TeacherID int
	// 最终只有TeacherID字段
	Teacher Teacher `gorm:"constraint:onDelete:CASCADE"`
}

type Student struct {
	BaseModel
	Name   string `gorm:"type:varchar(32);unique;not null"`
	Sno    int
	Pwd    string `gorm:"type:varchar(100);not null"`
	Tel    string `gorm:"type:char(11)"`
	Gender byte   `gorm:"default:1"`
	Birth  *time.Time
	Remark string `gorm:"type:varchar(255)"`
	// 多对一，学生（多）只有一个班级，students表与classes表是多对一的关系。
	ClassID int
	Class   Class `gorm:"foreignKey:ClassID;constraint:onDelete:CASCADE"`
	// 多对多，学生（多）有多个课程，students表与courses表是多对多的关系。
	Courses []Course `gorm:"many2many:student2course;constraint:onDelete:CASCADE"`
}

func addRecord(db *gorm.DB) {
	// 结构体的实例化对象和表记录产生映射

	// // 添加老师
	// t1 := Teacher{Name: "li",Tno:  1001,Pwd:  "123"}
	// db.Create(&t1)
	// t2 := Teacher{Name: "zhang", Tno: 1002, Pwd: "456"}
	// db.Create(&t2)
	// t3 := Teacher{Name: "wang", Tno: 1003, Pwd: "789"}
	// db.Create(&t3)

	// // 批量创建班级
	// c1 := Class{Name: "班级01", Num: 36, TeacherID: 1}
	// c2 := Class{Name: "班级02", Num: 40, TeacherID: 2}
	// c3 := Class{Name: "班级03", Num: 56, TeacherID: 3}
	// c4 := Class{Name: "班级04", Num: 32, TeacherID: 1}
	// classes := []Class{c1, c2, c3, c4}
	// db.Create(&classes)

	// // 批量创建课程
	// course1 := Course{Name: "go", Credit: 3, Period: 16, TeacherID: 1}
	// course2 := Course{Name: "java", Credit: 2, Period: 18, TeacherID: 1}
	// course3 := Course{Name: "python", Credit: 3, Period: 12, TeacherID: 2}
	// course4 := Course{Name: "c", Credit: 2, Period: 10, TeacherID: 2}
	// course5 := Course{Name: "js", Credit: 1, Period: 8, TeacherID: 3}
	// courses := []Course{course1, course2, course3, course4, course5}
	// db.Create(&courses)

	// 批量创建学生

}

func deleteRecord(db *gorm.DB) {
	// var course Course
	// db.Where("name = ? ", "python").Find(&course)
	// db.Delete(&course)
	// // DELETE FROM `courses` WHERE `courses`.`id` = 3;

	// db.Where("credit < ?", 2).Delete(&Course{})
	// // DELETE FROM `courses` WHERE credit < 2;

	// db.Where("1 = 1").Delete(&Course{})
	// // DELETE FROM `courses` WHERE 1 = 1;
}

func updateRecord(db *gorm.DB) {
	// db.Model(&Course{}).Where("period < ?", 16).Update("credit", 2)
	// // UPDATE `courses` SET `credit`= 2 WHERE period < 16;

	// db.Model(&Course{}).Where("period > ?", 16).Updates(&Course{Credit: 4, Period: 32})
	// // UPDATE `courses` SET `credit`= 4,`period`= 32 WHERE period > 16;

	// db.Model(&Course{}).Where("period > ?", 16).Updates(map[string]interface{}{"Credit": 4, "Period": 32})
	// // UPDATE `courses` SET `credit`= 4,`period`= 32 WHERE period > 16

	// // 更新表达式
	// db.Debug().Model(&Course{}).Where("period > ?", 16).Update("credit", gorm.Expr("credit+1"))
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold:             time.Second,   // Slow SQL threshold
	// 		LogLevel:                  logger.Silent, // Log level
	// 		IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
	// 		ParameterizedQueries:      true,          // Don't include params in the SQL log
	// 		Colorful:                  false,         // Disable color
	// 	},
	// )
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(&Teacher{})
	// db.AutoMigrate(&Class{})
	// db.AutoMigrate(&Course{})
	// db.AutoMigrate(&Student{})
	db.AutoMigrate(&Teacher{}, &Class{}, &Course{}, &Student{})

	//增加记录
	addRecord(db)

	//删除记录
	deleteRecord(db)

	//更新记录
	updateRecord(db)
}
