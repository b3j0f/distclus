package core

import "errors"

// ErrNotRunning raised while algorithm status equals Created, Ready or Failed
var ErrNotRunning = errors.New("Algorithm is not running")

// ErrRunning raised while algorithm status equals Running, Idle or Sleeping
var ErrRunning = errors.New("Algorithm is running")

// ErrInitializing raised while algorithm status equals Initializing
var ErrInitializing = errors.New("Algorithm is initializing")

// ErrSleeping raised while algorithm status equals Sleeping, Idle or Sleeping
var ErrSleeping = errors.New("Algorithm is sleeping")

// ErrAlreadyCreated raised if algorithm is already created
var ErrAlreadyCreated = errors.New("Algorithm is already created")

// ErrNotIdle idle status is asked and not setted
var ErrNotIdle = errors.New("Algorithm is not idle")

// ErrIdle raised if algo is idle
var ErrIdle = errors.New("Algorithm is idle")

// ErrReconfiguring raised if algo is reconfiguring
var ErrReconfiguring = errors.New("Algorithm is reconfiguring")

// ErrTimeout is returned when an error occurs
var ErrTimeout = errors.New("algorithm timed out")

// ErrElapsedIter raised when amont of iterations is done
var ErrElapsedIter = errors.New("amount of iterations done")

// ErrNeverFinish raised when wait method is called while the algorithm will never finish
var ErrNeverFinish = errors.New("algorithm can not finish. Specify core.Conf.Iter, core.Conf.IterPerData, core.Conf.DataPerIter or core.Conf.Finishing for allowing your algorithm to sleep")

// ErrFinished raised when algorithm is finished
var ErrFinished = errors.New("algorithm is finished")

// ErrNotIterate raised when play is called while algo can not iterate
var ErrNotIterate = errors.New("algorithm can not iterate. Check iterations and dataPerIter conditions")

// ErrNotAlive raised when algo is not alive
var ErrNotAlive = errors.New("algorithm is not alive")
