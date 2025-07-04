// Code generated by http://github.com/gojuno/minimock (v3.4.5). DO NOT EDIT.

package create_comment

//go:generate minimock -i example/comments/internal/usecases/create-comment.UserService -o user_service_mock_test.go -n UserServiceMock -p create_comment

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// UserServiceMock implements UserService
type UserServiceMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCheckUserID          func(ctx context.Context, userID int64) (b1 bool, err error)
	funcCheckUserIDOrigin    string
	inspectFuncCheckUserID   func(ctx context.Context, userID int64)
	afterCheckUserIDCounter  uint64
	beforeCheckUserIDCounter uint64
	CheckUserIDMock          mUserServiceMockCheckUserID
}

// NewUserServiceMock returns a mock for UserService
func NewUserServiceMock(t minimock.Tester) *UserServiceMock {
	m := &UserServiceMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CheckUserIDMock = mUserServiceMockCheckUserID{mock: m}
	m.CheckUserIDMock.callArgs = []*UserServiceMockCheckUserIDParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mUserServiceMockCheckUserID struct {
	optional           bool
	mock               *UserServiceMock
	defaultExpectation *UserServiceMockCheckUserIDExpectation
	expectations       []*UserServiceMockCheckUserIDExpectation

	callArgs []*UserServiceMockCheckUserIDParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// UserServiceMockCheckUserIDExpectation specifies expectation struct of the UserService.CheckUserID
type UserServiceMockCheckUserIDExpectation struct {
	mock               *UserServiceMock
	params             *UserServiceMockCheckUserIDParams
	paramPtrs          *UserServiceMockCheckUserIDParamPtrs
	expectationOrigins UserServiceMockCheckUserIDExpectationOrigins
	results            *UserServiceMockCheckUserIDResults
	returnOrigin       string
	Counter            uint64
}

// UserServiceMockCheckUserIDParams contains parameters of the UserService.CheckUserID
type UserServiceMockCheckUserIDParams struct {
	ctx    context.Context
	userID int64
}

// UserServiceMockCheckUserIDParamPtrs contains pointers to parameters of the UserService.CheckUserID
type UserServiceMockCheckUserIDParamPtrs struct {
	ctx    *context.Context
	userID *int64
}

// UserServiceMockCheckUserIDResults contains results of the UserService.CheckUserID
type UserServiceMockCheckUserIDResults struct {
	b1  bool
	err error
}

// UserServiceMockCheckUserIDOrigins contains origins of expectations of the UserService.CheckUserID
type UserServiceMockCheckUserIDExpectationOrigins struct {
	origin       string
	originCtx    string
	originUserID string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmCheckUserID *mUserServiceMockCheckUserID) Optional() *mUserServiceMockCheckUserID {
	mmCheckUserID.optional = true
	return mmCheckUserID
}

// Expect sets up expected params for UserService.CheckUserID
func (mmCheckUserID *mUserServiceMockCheckUserID) Expect(ctx context.Context, userID int64) *mUserServiceMockCheckUserID {
	if mmCheckUserID.mock.funcCheckUserID != nil {
		mmCheckUserID.mock.t.Fatalf("UserServiceMock.CheckUserID mock is already set by Set")
	}

	if mmCheckUserID.defaultExpectation == nil {
		mmCheckUserID.defaultExpectation = &UserServiceMockCheckUserIDExpectation{}
	}

	if mmCheckUserID.defaultExpectation.paramPtrs != nil {
		mmCheckUserID.mock.t.Fatalf("UserServiceMock.CheckUserID mock is already set by ExpectParams functions")
	}

	mmCheckUserID.defaultExpectation.params = &UserServiceMockCheckUserIDParams{ctx, userID}
	mmCheckUserID.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmCheckUserID.expectations {
		if minimock.Equal(e.params, mmCheckUserID.defaultExpectation.params) {
			mmCheckUserID.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCheckUserID.defaultExpectation.params)
		}
	}

	return mmCheckUserID
}

// ExpectCtxParam1 sets up expected param ctx for UserService.CheckUserID
func (mmCheckUserID *mUserServiceMockCheckUserID) ExpectCtxParam1(ctx context.Context) *mUserServiceMockCheckUserID {
	if mmCheckUserID.mock.funcCheckUserID != nil {
		mmCheckUserID.mock.t.Fatalf("UserServiceMock.CheckUserID mock is already set by Set")
	}

	if mmCheckUserID.defaultExpectation == nil {
		mmCheckUserID.defaultExpectation = &UserServiceMockCheckUserIDExpectation{}
	}

	if mmCheckUserID.defaultExpectation.params != nil {
		mmCheckUserID.mock.t.Fatalf("UserServiceMock.CheckUserID mock is already set by Expect")
	}

	if mmCheckUserID.defaultExpectation.paramPtrs == nil {
		mmCheckUserID.defaultExpectation.paramPtrs = &UserServiceMockCheckUserIDParamPtrs{}
	}
	mmCheckUserID.defaultExpectation.paramPtrs.ctx = &ctx
	mmCheckUserID.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmCheckUserID
}

// ExpectUserIDParam2 sets up expected param userID for UserService.CheckUserID
func (mmCheckUserID *mUserServiceMockCheckUserID) ExpectUserIDParam2(userID int64) *mUserServiceMockCheckUserID {
	if mmCheckUserID.mock.funcCheckUserID != nil {
		mmCheckUserID.mock.t.Fatalf("UserServiceMock.CheckUserID mock is already set by Set")
	}

	if mmCheckUserID.defaultExpectation == nil {
		mmCheckUserID.defaultExpectation = &UserServiceMockCheckUserIDExpectation{}
	}

	if mmCheckUserID.defaultExpectation.params != nil {
		mmCheckUserID.mock.t.Fatalf("UserServiceMock.CheckUserID mock is already set by Expect")
	}

	if mmCheckUserID.defaultExpectation.paramPtrs == nil {
		mmCheckUserID.defaultExpectation.paramPtrs = &UserServiceMockCheckUserIDParamPtrs{}
	}
	mmCheckUserID.defaultExpectation.paramPtrs.userID = &userID
	mmCheckUserID.defaultExpectation.expectationOrigins.originUserID = minimock.CallerInfo(1)

	return mmCheckUserID
}

// Inspect accepts an inspector function that has same arguments as the UserService.CheckUserID
func (mmCheckUserID *mUserServiceMockCheckUserID) Inspect(f func(ctx context.Context, userID int64)) *mUserServiceMockCheckUserID {
	if mmCheckUserID.mock.inspectFuncCheckUserID != nil {
		mmCheckUserID.mock.t.Fatalf("Inspect function is already set for UserServiceMock.CheckUserID")
	}

	mmCheckUserID.mock.inspectFuncCheckUserID = f

	return mmCheckUserID
}

// Return sets up results that will be returned by UserService.CheckUserID
func (mmCheckUserID *mUserServiceMockCheckUserID) Return(b1 bool, err error) *UserServiceMock {
	if mmCheckUserID.mock.funcCheckUserID != nil {
		mmCheckUserID.mock.t.Fatalf("UserServiceMock.CheckUserID mock is already set by Set")
	}

	if mmCheckUserID.defaultExpectation == nil {
		mmCheckUserID.defaultExpectation = &UserServiceMockCheckUserIDExpectation{mock: mmCheckUserID.mock}
	}
	mmCheckUserID.defaultExpectation.results = &UserServiceMockCheckUserIDResults{b1, err}
	mmCheckUserID.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmCheckUserID.mock
}

// Set uses given function f to mock the UserService.CheckUserID method
func (mmCheckUserID *mUserServiceMockCheckUserID) Set(f func(ctx context.Context, userID int64) (b1 bool, err error)) *UserServiceMock {
	if mmCheckUserID.defaultExpectation != nil {
		mmCheckUserID.mock.t.Fatalf("Default expectation is already set for the UserService.CheckUserID method")
	}

	if len(mmCheckUserID.expectations) > 0 {
		mmCheckUserID.mock.t.Fatalf("Some expectations are already set for the UserService.CheckUserID method")
	}

	mmCheckUserID.mock.funcCheckUserID = f
	mmCheckUserID.mock.funcCheckUserIDOrigin = minimock.CallerInfo(1)
	return mmCheckUserID.mock
}

// When sets expectation for the UserService.CheckUserID which will trigger the result defined by the following
// Then helper
func (mmCheckUserID *mUserServiceMockCheckUserID) When(ctx context.Context, userID int64) *UserServiceMockCheckUserIDExpectation {
	if mmCheckUserID.mock.funcCheckUserID != nil {
		mmCheckUserID.mock.t.Fatalf("UserServiceMock.CheckUserID mock is already set by Set")
	}

	expectation := &UserServiceMockCheckUserIDExpectation{
		mock:               mmCheckUserID.mock,
		params:             &UserServiceMockCheckUserIDParams{ctx, userID},
		expectationOrigins: UserServiceMockCheckUserIDExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmCheckUserID.expectations = append(mmCheckUserID.expectations, expectation)
	return expectation
}

// Then sets up UserService.CheckUserID return parameters for the expectation previously defined by the When method
func (e *UserServiceMockCheckUserIDExpectation) Then(b1 bool, err error) *UserServiceMock {
	e.results = &UserServiceMockCheckUserIDResults{b1, err}
	return e.mock
}

// Times sets number of times UserService.CheckUserID should be invoked
func (mmCheckUserID *mUserServiceMockCheckUserID) Times(n uint64) *mUserServiceMockCheckUserID {
	if n == 0 {
		mmCheckUserID.mock.t.Fatalf("Times of UserServiceMock.CheckUserID mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmCheckUserID.expectedInvocations, n)
	mmCheckUserID.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmCheckUserID
}

func (mmCheckUserID *mUserServiceMockCheckUserID) invocationsDone() bool {
	if len(mmCheckUserID.expectations) == 0 && mmCheckUserID.defaultExpectation == nil && mmCheckUserID.mock.funcCheckUserID == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmCheckUserID.mock.afterCheckUserIDCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmCheckUserID.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// CheckUserID implements UserService
func (mmCheckUserID *UserServiceMock) CheckUserID(ctx context.Context, userID int64) (b1 bool, err error) {
	mm_atomic.AddUint64(&mmCheckUserID.beforeCheckUserIDCounter, 1)
	defer mm_atomic.AddUint64(&mmCheckUserID.afterCheckUserIDCounter, 1)

	mmCheckUserID.t.Helper()

	if mmCheckUserID.inspectFuncCheckUserID != nil {
		mmCheckUserID.inspectFuncCheckUserID(ctx, userID)
	}

	mm_params := UserServiceMockCheckUserIDParams{ctx, userID}

	// Record call args
	mmCheckUserID.CheckUserIDMock.mutex.Lock()
	mmCheckUserID.CheckUserIDMock.callArgs = append(mmCheckUserID.CheckUserIDMock.callArgs, &mm_params)
	mmCheckUserID.CheckUserIDMock.mutex.Unlock()

	for _, e := range mmCheckUserID.CheckUserIDMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.b1, e.results.err
		}
	}

	if mmCheckUserID.CheckUserIDMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCheckUserID.CheckUserIDMock.defaultExpectation.Counter, 1)
		mm_want := mmCheckUserID.CheckUserIDMock.defaultExpectation.params
		mm_want_ptrs := mmCheckUserID.CheckUserIDMock.defaultExpectation.paramPtrs

		mm_got := UserServiceMockCheckUserIDParams{ctx, userID}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmCheckUserID.t.Errorf("UserServiceMock.CheckUserID got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCheckUserID.CheckUserIDMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.userID != nil && !minimock.Equal(*mm_want_ptrs.userID, mm_got.userID) {
				mmCheckUserID.t.Errorf("UserServiceMock.CheckUserID got unexpected parameter userID, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCheckUserID.CheckUserIDMock.defaultExpectation.expectationOrigins.originUserID, *mm_want_ptrs.userID, mm_got.userID, minimock.Diff(*mm_want_ptrs.userID, mm_got.userID))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCheckUserID.t.Errorf("UserServiceMock.CheckUserID got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmCheckUserID.CheckUserIDMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCheckUserID.CheckUserIDMock.defaultExpectation.results
		if mm_results == nil {
			mmCheckUserID.t.Fatal("No results are set for the UserServiceMock.CheckUserID")
		}
		return (*mm_results).b1, (*mm_results).err
	}
	if mmCheckUserID.funcCheckUserID != nil {
		return mmCheckUserID.funcCheckUserID(ctx, userID)
	}
	mmCheckUserID.t.Fatalf("Unexpected call to UserServiceMock.CheckUserID. %v %v", ctx, userID)
	return
}

// CheckUserIDAfterCounter returns a count of finished UserServiceMock.CheckUserID invocations
func (mmCheckUserID *UserServiceMock) CheckUserIDAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCheckUserID.afterCheckUserIDCounter)
}

// CheckUserIDBeforeCounter returns a count of UserServiceMock.CheckUserID invocations
func (mmCheckUserID *UserServiceMock) CheckUserIDBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCheckUserID.beforeCheckUserIDCounter)
}

// Calls returns a list of arguments used in each call to UserServiceMock.CheckUserID.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCheckUserID *mUserServiceMockCheckUserID) Calls() []*UserServiceMockCheckUserIDParams {
	mmCheckUserID.mutex.RLock()

	argCopy := make([]*UserServiceMockCheckUserIDParams, len(mmCheckUserID.callArgs))
	copy(argCopy, mmCheckUserID.callArgs)

	mmCheckUserID.mutex.RUnlock()

	return argCopy
}

// MinimockCheckUserIDDone returns true if the count of the CheckUserID invocations corresponds
// the number of defined expectations
func (m *UserServiceMock) MinimockCheckUserIDDone() bool {
	if m.CheckUserIDMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.CheckUserIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.CheckUserIDMock.invocationsDone()
}

// MinimockCheckUserIDInspect logs each unmet expectation
func (m *UserServiceMock) MinimockCheckUserIDInspect() {
	for _, e := range m.CheckUserIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserServiceMock.CheckUserID at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterCheckUserIDCounter := mm_atomic.LoadUint64(&m.afterCheckUserIDCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.CheckUserIDMock.defaultExpectation != nil && afterCheckUserIDCounter < 1 {
		if m.CheckUserIDMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to UserServiceMock.CheckUserID at\n%s", m.CheckUserIDMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to UserServiceMock.CheckUserID at\n%s with params: %#v", m.CheckUserIDMock.defaultExpectation.expectationOrigins.origin, *m.CheckUserIDMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCheckUserID != nil && afterCheckUserIDCounter < 1 {
		m.t.Errorf("Expected call to UserServiceMock.CheckUserID at\n%s", m.funcCheckUserIDOrigin)
	}

	if !m.CheckUserIDMock.invocationsDone() && afterCheckUserIDCounter > 0 {
		m.t.Errorf("Expected %d calls to UserServiceMock.CheckUserID at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.CheckUserIDMock.expectedInvocations), m.CheckUserIDMock.expectedInvocationsOrigin, afterCheckUserIDCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *UserServiceMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCheckUserIDInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *UserServiceMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *UserServiceMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCheckUserIDDone()
}
