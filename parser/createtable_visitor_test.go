package parser

import "testing"

func Test_onlyTableName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"foo", "foo"},
		{"foo`.`bar", "bar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := onlyTableName(tt.name); got != tt.want {
				t.Errorf("onlyTableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
