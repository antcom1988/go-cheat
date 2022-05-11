package sampletest

import (
	"errors"
	"io"
	"reflect"
	"testing"
)

type mockReadCloser struct {
	expectedData []byte
	expectedErr error
}

// allow test table to fill out expectedData and expectedErr field
func (mrc *mockReadCloser) Read(p []byte) (n int, err error) {
	copy(p, mrc.expectedData)
	return 0, mrc.expectedErr
}

// do nothing on close
func (mrc *mockReadCloser) Close() error {
	return nil
}

func TestReadContents(t *testing.T) {
	errRead := errors.New("error")
	tests := []struct {
		name    string
		readCloser io.ReadCloser
		numBytes int
		expectedData []byte
		expectedErr error
	}{
		{
			name: "Test Success",
			readCloser: &mockReadCloser{
				expectedData: []byte(`hello`),
				expectedErr: nil,
			},
			numBytes: 5,
			expectedData: []byte(`hello`),
			expectedErr: nil,
		},
		{
			name: "Test Error",
			readCloser: &mockReadCloser{
				expectedData: nil,
				expectedErr: errRead,
			},
			numBytes: 0,
			expectedData: nil,
			expectedErr: errRead,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := ReadContents(tt.readCloser, tt.numBytes)
			if !reflect.DeepEqual(data, tt.expectedData) {
				t.Errorf("ReadContents() got = %v, want %v", data, tt.expectedData)
			}
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("ReadContents() error = %v, wantErr %v", err, tt.expectedErr)
				return
			}
		})
	}
}
