package sync

import (
	"context"
	"testing"
	"testing/synctest"
)

// func TestAfterFunc(t *testing.T) {

// 	ctx, cancel := context.WithCancel(context.Background())
// 	calledCh := make(chan struct{})

// 	// funcCalled reports whether the function was called.
// 	funcCalled := func() bool {
// 		select {
// 		case <-calledCh:
// 			return true
// 		case <-time.After(10 * time.Millisecond):
// 			return false
// 		}
// 	}

// 	context.AfterFunc(ctx, func() {
// 		close(calledCh)
// 	})

// 	// TODO: Assert that the AfterFunc has not been called yet
// 	if funcCalled() {
// 		t.Fatalf("AfterFunc function called before context is canceled")
// 	}

// 	cancel()

// 	// TODO: Assert that the AfterFunc has been called
// 	if !funcCalled() {
// 		t.Fatalf("AfterFunc function not called after context is canceled")
// 	}
// }

func TestAfterFuncSync(t *testing.T) {
	synctest.Run(func() {

		ctx, cancel := context.WithCancel(context.Background())
		funcCalled := false

		context.AfterFunc(ctx, func() {
			funcCalled = true
		})

		synctest.Wait()
		if funcCalled {
			t.Fatalf("AfterFunc function called before context is canceled")
		}
		cancel()

		synctest.Wait()

		if !funcCalled {
			t.Fatalf("AfterFunc function not called after context is canceled")
		}
	})
}
