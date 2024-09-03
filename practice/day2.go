package practice

import "fmt"

func Day2(){

	 //local variable
	 var age int = 25;

	 var salary float32 = 100.00;
	 //first declearation and later assign
	 var name string;

	 var boolean bool = false;
	 fmt.Println(boolean)

	 //assigning later
	 name = "john";

	 fmt.Println("Name, Age, Salary :  ", name, age, salary);

	 //dynamic variable
	 city := "Bengaluru";

	 fmt.Println(city);

	 //Constant variables

	 const ERROR_MESSAGE string = "Not found";
	 const ERROR_CODE int = 404;

	 fmt.Println(ERROR_MESSAGE);
	 fmt.Println(ERROR_CODE);

	 // ERROR_CODE = 405 give the compile time error

	 var a, b, c = 5, 4, "foo";

	 // varify data type
	 fmt.Printf("a is of type %T\n", a);
	 fmt.Printf("b is of type %T\n", b);
	 fmt.Printf("c is of type %T\n", c);

	 //conditions
	 if(a>b){
		fmt.Println("A is greater than B", a);
	 }else{
		fmt.Println("B is greater than A",b);
	 }

	 //switch single case
	 day := "Monday";

	 switch (day) {
	case "Monday":
        fmt.Println("It's Monday!");
    case "Tuesday":
        fmt.Println("It's Tuesday!");
    case "Wednesday", "Thursday":
        fmt.Println("It's mid-week.");
    default:
        fmt.Println("It's some other day.");
	 }

	 //switch with multiple case
        month := 28;
	 	switch month{
		case 31,29 : fmt.Println("Odd month");
		case 30, 28: fmt.Println("Even month");
		default : fmt.Println("No month for the input");
		}

	 //loop
	 for i := 0; i < 5; i++ {
        fmt.Println("Iteration:", i);
    }

	cars :=[3] string{"BMW","AUDI","TATA"};
	for idx, val := range cars{
		fmt.Printf("%v\t%v\n", idx,val);
	}

	car :=[...] string{"BMW","AUDI","TATA","MARUTI"};
	car[3] = "NewCar"
	for _, val := range car{
		fmt.Printf("%v\n",val);
	}

	j := 0
    for {
        fmt.Println(j)
        j++;
        if j == 2 {
            break;// Exits the loop when j reaches 2
        }
    }

	//finding min and max value from the array
	arr := [5]int{2,4,1,-1,12};
	min, max := findMaxMini(arr[:]);
	fmt.Println("Min and Max value : ", min, max)
}

func MaxValue(num1 , num2 int) int {
	var result int;
	if(num1 > num2){
		result = num1;
	}else {
		result =  num2;
	}
	return result;
}

func Swap(x,y string) (string, string){
	return y,x;
}

func findMaxMini(arr []int) (int, int){
	min := 0;
	max := 0;
	for i := 0; i < len(arr); i++ {
		if(arr[i] <= min){
			min = arr[i]
		}
		if(arr[i] >= max){
			max = arr[i]
		}
	}
	return min, max;
}

