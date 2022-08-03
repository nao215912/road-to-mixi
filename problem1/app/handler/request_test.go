package handler

import "testing"

func TestConvertUserID(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertUserID(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
