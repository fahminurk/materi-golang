package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	RwMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RwMutex.Lock()
	account.Balance = account.Balance + amount
	account.RwMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RwMutex.RLock()
	balance := account.Balance
	account.RwMutex.RUnlock()
    return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0;i < 100; i++ {
		go func(){
			for j := 0;j < 100; j++ {
				account.AddBalance(1)
				println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance", account.GetBalance())
}

type UserBalance struct{
	Mutex sync.Mutex
	Balance int
	Name string
}

func (user *UserBalance) Lock(){
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock(){
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int){
	user.Balance += amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int){
	user1.Lock()
	user1.Change(-amount)
	fmt.Println("Lock user1", user1.Name)

	time.Sleep(1 * time.Second)

	user2.Lock()
	user2.Change(amount)
	fmt.Println("Lock user2", user2.Name)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T){
	user1 := UserBalance{
		Name: "fahmi",
		Balance: 10000,
	}
	user2 := UserBalance{
		Name: "andi",
		Balance: 10000,
	}

	go Transfer(&user1, &user2, 5000)
	go Transfer(&user2, &user1, 1000)

	time.Sleep(2 * time.Second)

	fmt.Println("user :", user1.Name, ", balance :", user1.Balance)
	fmt.Println("user :", user2.Name, ", balance :", user2.Balance)
}