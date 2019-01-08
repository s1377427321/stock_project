package pool

import (
	"runtime"
	"time"

	"github.com/hyper-carrot/go_lib"
	"github.com/hyper-carrot/go_lib/logging"
)

var logger logging.Logger = logging.GetSimpleLogger()

type InitFunc func() (interface{}, error)

type Pool struct {
	Id        string
	Size      int
	container chan interface{}
	rwSign    go_lib.RWSign
}

func (self *Pool) Init(initFunc InitFunc) error {
	if cap(self.container) != self.Size {
		logger.Infof("Initializing pool (Id=%s, Size=%d)...\n", self.Id, self.Size)
		self.container = make(chan interface{}, self.Size)
	}
	for i := 0; i < self.Size; i++ {
		element, err := initFunc()
		if err != nil {
			return err
		}
		if element == nil {
			logger.Warnf("The initialized element is NIL! (poolId=%s)", self.Id)
		}
		self.container <- element
	}
	logger.Infof("The pool (Id=%v, Size=%d) has been initialized.\n", self.Id, self.Size)
	return nil
}

func (self *Pool) Get(timeoutMs time.Duration) (element interface{}, ok bool) {
	// LogInfof("Getting! (Size: %v, Cap: %v)", len(self.container), cap(self.container))
	if self.Closed() {
		return nil, false
	}
	if timeoutMs > 0 {
		select {
		case element, ok = <-self.container:
			return
		case <-time.After(5 * time.Millisecond):
			logger.Infof("Getting Timeout! (Size: %v, Cap: %v)", len(self.container), cap(self.container))
			element, ok = nil, false
			return
		}
	} else {
		if len(self.container) == 0 {
			element, ok = nil, false
		} else {
			element, ok = <-self.container
		}
	}
	return
}

func (self *Pool) Put(element interface{}, timeoutMs time.Duration) bool {
	// logger.Infof("Putting! (Size: %v, Cap: %v)", len(self.container), cap(self.container))
	if self.Closed() {
		return false
	}
	result := false
	if timeoutMs > 0 {
		sign := make(chan bool)
		go func() {
			time.AfterFunc(5*time.Millisecond, func() {
				if !result {
					logger.Infof("Putting Timeout! (Size: %v, Cap: %v, Element: %v)", len(self.container), cap(self.container), element)
					sign <- result
				}
				runtime.Goexit()
			})
			self.container <- element
			result = true
			sign <- result
		}()
		return <-sign
	} else {
		if len(self.container) >= self.Size {
			result = false
		} else {
			self.container <- element
			result = true
		}
	}
	return result
}

func (self *Pool) Close() {
	close(self.container)
	self.container = nil
	return
}

func (self *Pool) Closed() bool {
	if self == nil || self.container == nil {
		return true
	}
	return false
}
