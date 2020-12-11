package db

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	type args struct {
		protocol string
		host     string
		port     int
		dbname   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{protocol: "mongodb", host: "localhost", port: 27017, dbname: "testdb"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.protocol, tt.args.host, tt.args.port, tt.args.dbname)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			defer got.Close()
		})
	}
}
