package main

import "fmt"

//구조체 : 여러 필드를 묶어서 사용한다.
//연관된 여러 데이터를 하나의 이름으로 묶어서 사용한다.

//type:이 키워드를 통해 새로운 사용자 정의 타입을 정의한다.
//name : 타입명을 적는다.
//struct : 타입의 종류
type Student struct {
	Name  string
	Age   int
	Score float32
	No    int
}
type User struct {
	Name     string
	ID       string
	Password int
}

//구조체 안에 구조체를 넣을수 있다.
type HardUser struct {
	userInfo User //위에 있는 User구조체
	Level    int
	Age      int
}

//필드명 생략
type HardUser1 struct {
	User
	Level int
	Age   int
}

func main() {

	// var Student_name string = "홍길동"
	// var Student_age int =10
	// var Student_phoneNumber int

	var st Student
	st.Age = 15
	st.No = 10
	st.Name = "캡틴 아메리카노"
	st.Score = 3.14

	fmt.Println(st)
	fmt.Println("학생이름 : ", st.Name)
	fmt.Println("학생나이 : ", st.Age)
	fmt.Println("학생번호 : ", st.No)
	fmt.Println("학생점수 : ", st.Score)

	//초기화 방법
	var st1 Student = Student{"달마대사", 1000, 30.5, 20} //한줄에 전부
	var st2 Student = Student{                        //여러줄에..
		"원효대사",
		2000,
		10.1,
		50}
	var st3 Student = Student{Name: "신사임당", Age: 1234} //일부만
	fmt.Println(st1)
	fmt.Println(st2)
	fmt.Println(st3)

	var normalUser = User{"홍길동", "오늘만 산다", 100}
	var hardUser = HardUser{User{"놀부", "흥부가 기가막혀", 200}, 10, 20}
	fmt.Println(normalUser)
	fmt.Println(hardUser)

	fmt.Printf("과금 유저 이름 : %s 아이디 :%s 비밀번호 : %d 레벨  : %d 나이  : %d",
		hardUser.userInfo.Name, //UserInfo안에 Name
		hardUser.userInfo.ID,
		hardUser.userInfo.Password,
		hardUser.Level,
		hardUser.Age,
	)
	//구조체안에 구조체에 접근 하려면 .을 두번찍어 접근해야함
	//하지만 포함하는 구조체 필드명 생략하면 한번만 사용해서 접근 가능함.
	var hardUser1 = HardUser1{User{"하하하하", "바빠", 200}, 10, 20}
	fmt.Printf("과금 유저 이름 : %s 아이디 :%s 비밀번호 : %d 레벨  : %d 나이  : %d",
		hardUser1.Name,
		hardUser1.ID,
		hardUser1.Password,
		hardUser1.Level,
		hardUser1.Age,
	)

}
