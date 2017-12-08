package main

import (
	"reflect"
	"testing"
)

func Test_parse(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want Register
	}{
		{"test", args{str: "es inc -936 if qen == 139"},
			Register{Name: "es", Action: "inc", Value: -936, Expression: "qen == 139", Condition: "==", ConditionRegister: "qen", ConditionValue: 139}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
