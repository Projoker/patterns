package viewmodel

import (
	"context"
	"fmt"
	"time"
)

type Model interface {
	GetData() []byte
}

type ViewModel struct {
	m        Model
	data     []byte
	notifyCh chan bool
}

func New(ctx context.Context, m Model) *ViewModel {
	vm := ViewModel{m: m, notifyCh: make(chan bool)}
	go vm.watch(ctx)

	return &vm
}

func (vm *ViewModel) GetNotifyCh() chan bool {
	return vm.notifyCh
}

func (vm *ViewModel) GetData() []byte {
	return vm.data
}

func (vm *ViewModel) watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			vm.checkModel()

			time.Sleep(500 * time.Millisecond)
		}
	}
}

func (vm *ViewModel) checkModel() {
	data := vm.m.GetData()

	if len(data) == len(vm.data) {
		return
	}

	fmt.Printf("[viewmodel] model update detected at %v\n", time.Now())
	vm.data = data
	vm.notifyCh <- true
}
