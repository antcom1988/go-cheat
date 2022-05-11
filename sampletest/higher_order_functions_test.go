package sampletest

import (
	"database/sql"
	"errors"
	"testing"
)

func TestOpenDBHighOrderFunc(t *testing.T) {
	type args struct {
		user     string
		password string
		addr     string
		db       string
		opener   sqlOpener
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Test Success",
			args:    args{
				user:     "u",
				password: "p",
				addr:     "a",
				db:       "db",
				opener: func(driver string, source string) (*sql.DB, error) {
					if source != "u:p@a/db" {
						return nil, errors.New("Wrong Connection String")
					}
					return nil, nil
				},
			},
			wantErr: false,
		},
		{
			name:    "Test Failed",
			args:    args{
				opener: func(driver string, source string) (*sql.DB, error) {
					return nil, errors.New("Wrong Connection String")
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := OpenDBHighOrderFunc(tt.args.user, tt.args.password, tt.args.addr, tt.args.db, tt.args.opener)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenDBHighOrderFunc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
