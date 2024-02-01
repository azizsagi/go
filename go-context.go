package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func isUserActivate(ctx context.Context, userId int) (bool, error) {
	time.Sleep(time.Millisecond * 400)
	if ctx.Err() == context.DeadlineExceeded {
		return false, errors.New("Time limit exceeds")
	}
	return true, nil
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*400)
	defer cancel()

	userId := 1234

	userActivated, error := isUserActivate(ctx, userId)

	if error != nil {
		fmt.Printf("Error in returning user :%d Error is: %s", userId, error)
	} else {
		fmt.Printf("The user is activated %d", userActivated)
	}

}
