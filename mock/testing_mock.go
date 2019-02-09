package mock

// The TMock is a struct used to test testing.T mock method calls
type TMock struct {
	calls map[string]int
}

// CallsTo returns the number of calls for the method passed as parameter.
// In case of a not existing mock method, the ok result parameter will be false.
func (m *TMock) CallsTo(name string) (numCalls int, ok bool) {
	numCalls, ok = m.calls[name]
	return
}

// Fail method is the fake implementation of the original method in the testing.T
func (m *TMock) Fail() {
	m.calls["Fail"]++
	return
}

// FailNow is the mock for the same method in testing.T
func (m *TMock) FailNow() {
	m.calls["FailNow"]++
	return
}

// Name returns the name of the Test as in the testing.T interface.
func (m *TMock) Name() string {
	m.calls["Name"]++
	return "Test"
}

// NewTMock creates a new testing.T mock object useful to test assertions.
func NewTMock() *TMock {
	return &TMock{
		calls: make(map[string]int),
	}
}
