package handlers

type ytClientMock struct{}

func NewYTStub() ytClientMock {
	return ytClientMock{}
}
