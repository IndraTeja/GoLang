package main

import "fmt"

func main() {
	var response int
	fmt.Println("\n Welcome to ** Bank of Frisco! ** ")
	//fmt.Println("\n Enter Total number of accounts : ")
	//const total int
	//fmt.Scanf("%d", &total)
	var account[3]int
	var name[3]string
	var balance[3]float64

	for i:=0; i < 3; i++ {
		fmt.Println("Enter Account number, name, balance :")
		fmt.Scanf("%d", &account[i])
		//, &name[i], &balance[i])
		fmt.Println("Enter name : \n")
		fmt.Scanf("%s",&name[i])
		fmt.Println("Enter balance :\n")
		fmt.Scanf("%f \n",&balance[i])
	}
	fmt.Println("\n Make a selection from menu \n 1. Search an account \n 2. Display all accounts \n 3. Deposit \n 4. Withdrawal \n 5. Exit")
	fmt.Println("Enter your choice : ")
	fmt.Scanf("%d", &response)

	switch response {
	default:
		fmt.Println("You made an invalid selection. Try again. ")
	case 1 :
		fmt.Println("Enter an account number to search:")
		var f int
		fmt.Scanf("%d",f)
		for i := range account{
			if account[i] == f {
				fmt.Println("Account found.")
				fmt.Println(account[i], "," , name[i], "," , balance[i]  )
			} else{
				fmt.Println("Account not found. Try again.")
			}
		}
	case 2 :
		fmt.Println("All the accounts")
		for i:=0;i<3;i++{
			fmt.Println(account[i], "," , name[i], "," , balance[i]  )
		}
	case 3 :
		fmt.Println("Deposited the money")
	case 4 :
		fmt.Println("Done Withdrawing $$ ")
	case 5 :
		fmt.Println("Thank you! Have a nice day!")
	}
}

/*
func addAccount(acc []int)  {

	if acc[0] == 0 {
		fmt.Println("Adding new account", acc)
	}

} */