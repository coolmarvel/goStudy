package main

import "fmt"

func HasRichFriend() bool {
	return true
}
func GetFriendCount() int {
	return 3
}

func getMyAge() (int, bool) {
	return 33, true
}
func SwitchGetAge() int {
	return 30
}

type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

//함수이름은 GetDirection
//함수매개변수는 angle float64받음
//함수 결과는 Direction타입 반환

//angle이 300이상 north반환
//angle이 0이상  45 보다 작으면north반환
//angle이 45이상  100 보다 작으면east반환
//angle이 100이상  200 보다 작으면south반환
//모든 조건이 만족하지않으면 none반환
func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "undefined"

	}
}
func GetColor() ColorType {
	return Blue
}
func main() {

	color := "red"

	if color == "red" {
		fmt.Println("나는 빨간색")
	} else {
		fmt.Println("아무것도 아님")
	}

	var age int = 22

	if age >= 10 && age <= 15 {
		fmt.Println("dddd")
	}

	price := 35000

	if price > 50000 {
		if HasRichFriend() {
			fmt.Println("오잉 난 돈이 없어")
		} else {
			fmt.Println("뿐빠이")
		}
	} else if price >= 30000 && price <= 50000 {
		if GetFriendCount() > 3 {
			fmt.Println("크음...")
		} else {
			fmt.Println("제발 나눠 내자")
		}
	} else {
		fmt.Println("오늘 내지갑을 털어봐!")
	}

	//검사에 사용할 변수를 초기화 할때 사용
	//초기문은 어떤 함수를 실행하고 결과를(반환값에 따라)검사 할때 사용
	// if 초기문; 조건문 {
	// 	ㅇ너마ㅏ임넝
	// }
	//33, true
	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("젊음", age)
	} else if ok && age < 30 {
		fmt.Println("괜찮음", age)
	} else if ok {
		fmt.Println("굿굿굿", age)
	} else {
		fmt.Println("아무것도 아님")
	}

	//switch
	num := 3
	switch num {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("에러")
	}
	//하나이상의 값도 구분할수 있다.쉼표로 구분...
	day := "sunday"
	switch day {
	case "monday", "tuseday":
		fmt.Println("언제 금요일이 오냐?")
	case "wednesday", "friday", "sunday":
		fmt.Println("퇴근 시켜줘~")
	}
	num1 := 18
	//if문처럼 true가 되는 조건을 검사 할수 있다.
	switch true { //비교할 값이 true ->case가 true인 녀석을 찾는다.
	case num1 < 10:
		fmt.Println("dasdsa")
	case num1 > 30:
		fmt.Println("12323")
	case num1 >= 10 && num1 < 20:
		fmt.Println("하하하하하")
	}
	//if문처럼 초기문을 넣을수 있다.
	//초기값, 비교값(SwitchGetAge()의 리턴값과 비교)
	switch age := SwitchGetAge(); age {
	case 10:
		fmt.Println("해당안됨")
	case 30:
		fmt.Println("여기가 마찝 이여쒀", age)
	}

	switch age := SwitchGetAge(); true {
	case age < 10:
		fmt.Println("가위")
	case age > 30:
		fmt.Println("바위")
	case age > 20 && age <= 30:
		fmt.Println("보")
	}

	abc := 3
	//go는 break있든 없든 빠져나온다.
	switch abc {
	case 1:
		fmt.Println("안녕")
	case 2:
		fmt.Println("안녕")
	case 3:
		fmt.Println("나는 3이다")
		fallthrough //다음 case까지 실행한다.
	case 4:
		fmt.Println("안녕12321312")
	case 5:
		fmt.Println("안녕")

	}
	//const 열거형에 따라 로직을 변경할때 쓰면 유용하다.위에이씀
	fmt.Println("색깔", colorToString(GetColor()))

	//함수이름은 GetDirection
	//함수매개변수는 angle float64받음
	//함수 결과는 Direction타입 반환

	//angle이 300이상 north반환
	//angle이 0이상  45 보다 작으면north반환
	//angle이 45이상  100 보다 작으면east반환
	//angle이 100이상  200 보다 작으면south반환
	//모든 조건이 만족하지않으면 none반환

}
