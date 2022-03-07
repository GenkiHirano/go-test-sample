package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Multiplication(t *testing.T) {

	// args 引数を定義
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string // テスト名
		args    args   // 引数
		want    int    // 欲しい結果
		wantErr bool   // 成功か失敗か
	}{
		{
			name:    "success",        // テスト名
			args:    args{a: 2, b: 3}, // 引数
			want:    6,                // 欲しい結果
			wantErr: false,            // 成功か失敗か
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := multiplication(tt.args.a, tt.args.b)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
