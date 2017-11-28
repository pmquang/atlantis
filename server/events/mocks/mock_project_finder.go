// Automatically generated by pegomock. DO NOT EDIT!
// Source: github.com/hootsuite/atlantis/server/events (interfaces: ProjectFinder)

package mocks

import (
	"reflect"

	models "github.com/hootsuite/atlantis/server/events/models"
	logging "github.com/hootsuite/atlantis/server/logging"
	pegomock "github.com/petergtz/pegomock"
)

type MockProjectFinder struct {
	fail func(message string, callerSkip ...int)
}

func NewMockProjectFinder() *MockProjectFinder {
	return &MockProjectFinder{fail: pegomock.GlobalFailHandler}
}

func (mock *MockProjectFinder) FindModified(log *logging.SimpleLogger, modifiedFiles []string, repoFullName string) []models.Project {
	params := []pegomock.Param{log, modifiedFiles, repoFullName}
	result := pegomock.GetGenericMockFrom(mock).Invoke("FindModified", params, []reflect.Type{reflect.TypeOf((*[]models.Project)(nil)).Elem()})
	var ret0 []models.Project
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]models.Project)
		}
	}
	return ret0
}

func (mock *MockProjectFinder) VerifyWasCalledOnce() *VerifierProjectFinder {
	return &VerifierProjectFinder{mock, pegomock.Times(1), nil}
}

func (mock *MockProjectFinder) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierProjectFinder {
	return &VerifierProjectFinder{mock, invocationCountMatcher, nil}
}

func (mock *MockProjectFinder) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierProjectFinder {
	return &VerifierProjectFinder{mock, invocationCountMatcher, inOrderContext}
}

type VerifierProjectFinder struct {
	mock                   *MockProjectFinder
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierProjectFinder) FindModified(log *logging.SimpleLogger, modifiedFiles []string, repoFullName string) *ProjectFinder_FindModified_OngoingVerification {
	params := []pegomock.Param{log, modifiedFiles, repoFullName}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "FindModified", params)
	return &ProjectFinder_FindModified_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ProjectFinder_FindModified_OngoingVerification struct {
	mock              *MockProjectFinder
	methodInvocations []pegomock.MethodInvocation
}

func (c *ProjectFinder_FindModified_OngoingVerification) GetCapturedArguments() (*logging.SimpleLogger, []string, string) {
	log, modifiedFiles, repoFullName := c.GetAllCapturedArguments()
	return log[len(log)-1], modifiedFiles[len(modifiedFiles)-1], repoFullName[len(repoFullName)-1]
}

func (c *ProjectFinder_FindModified_OngoingVerification) GetAllCapturedArguments() (_param0 []*logging.SimpleLogger, _param1 [][]string, _param2 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*logging.SimpleLogger, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(*logging.SimpleLogger)
		}
		_param1 = make([][]string, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.([]string)
		}
		_param2 = make([]string, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
	}
	return
}
