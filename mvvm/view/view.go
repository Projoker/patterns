package view

import (
	"context"
	"fmt"
	"time"
)

type ViewModel interface {
	GetData() []byte
	GetNotifyCh() chan bool
}

type View struct {
	vm   ViewModel
	data []byte
}

func New(ctx context.Context, vm ViewModel) *View {
	v := View{vm: vm}
	go v.watch(ctx)

	return &v
}

func (v *View) watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-v.vm.GetNotifyCh():
			v.checkViewModel()
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func (v *View) checkViewModel() {
	data := v.vm.GetData()

	if len(data) == len(v.data) {
		return
	}

	fmt.Printf("[view] viewmodel update detected at %v\n", time.Now())

	v.data = data
	fmt.Printf("[view] text:\n%v", string(v.data))
}
