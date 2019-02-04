package mock

type TMock struct {
	calls map[string]int
}

func (m *TMock) CallsTo(name string) (numCalls int, ok bool) {
	numCalls, ok = m.calls[name]
	return
}

func (m *TMock) Fail() {
	m.calls["Fail"]++
	return
}

func (m *TMock) FailNow() {
	m.calls["FailNow"]++
	return
}

func (m *TMock) Name() string {
	m.calls["Name"]++
	return "Test"
}

func NewTMock() *TMock {
	return &TMock{
		calls: make(map[string]int),
	}
}
