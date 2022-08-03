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
		{
			name:    "example",
			args:    args{s: "42"},
			want:    42,
			wantErr: false,
		},
		{
			name:    "minus",
			args:    args{s: "-42"},
			want:    0,
			wantErr: true,
		},
		{
			name:    "non_numeric",
			args:    args{s: "a"},
			want:    0,
			wantErr: true,
		},
		{
			name:    "empty",
			args:    args{s: ""},
			want:    0,
			wantErr: true,
		},
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

func TestConvertLimit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "example",
			args:    args{s: "42"},
			want:    42,
			wantErr: false,
		},
		{
			name:    "minus",
			args:    args{s: "-42"},
			want:    0,
			wantErr: true,
		},
		{
			name:    "non_numeric",
			args:    args{s: "a"},
			want:    0,
			wantErr: true,
		},
		{
			name:    "empty",
			args:    args{s: ""},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertLimit(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertLimit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertLimit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertPageQuery(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "example",
			args:    args{s: "42"},
			want:    42,
			wantErr: false,
		},
		{
			name:    "minus",
			args:    args{s: "-42"},
			want:    0,
			wantErr: true,
		},
		{
			name:    "non_numeric",
			args:    args{s: "a"},
			want:    0,
			wantErr: true,
		},
		{
			name:    "empty",
			args:    args{s: ""},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertPageQuery(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertPageQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertPageQuery() got = %v, want %v", got, tt.want)
			}
		})
	}
}
