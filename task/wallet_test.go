package task

import (
	"sync"
	"testing"
)

func TestRefill(t *testing.T) {
	userWallet := &Wallet{}
	var wg sync.WaitGroup  
	wg.Add(100000)          

	for i := 0; i < 100000; i++ {
		go func() {
			defer wg.Done()       
			Refill(userWallet, 1)
		}()
	}

	wg.Wait() 

	expectedBalance := 100000 
	if got := GetBalance(userWallet); got != expectedBalance {
		t.Errorf("expected balance %d, got %d", expectedBalance, got) 
	}
}

func TestWithdrawal(t *testing.T) {
	userWallet := &Wallet{}   
	Refill(userWallet, 10000) 
	var wg sync.WaitGroup    
	wg.Add(5000)            

	for i := 0; i < 5000; i++ {
		go func() {
			defer wg.Done() 
			if err := Withdrawal(userWallet, 1); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		}()
	}

	wg.Wait() 

	expectedBalance := 10000 - 5000 
	if got := GetBalance(userWallet); got != expectedBalance {
		t.Errorf("expected balance %d, got %d", expectedBalance, got) 
	}
}
