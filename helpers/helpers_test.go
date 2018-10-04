package helpers_test

import (
	"reflect"
	"testing"

	"github.com/nii236/margin/helpers"
	"github.com/nii236/margin/models"
)

func TestUnpackCurrent(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.CurrentMessage
		wantErr bool
	}{
		{
			name: "Case 1",
			args: args{
				msg: "2~Bitstamp~BTC~USD~2~6542.79~1538663723~0.00109225~7.146362377499999~75381949~4047.342658139999~26360111.227924854~ce9",
			},
			want:    &models.CurrentMessage{},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := helpers.UnpackCurrent(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("helpers.UnpackTrade() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("helpers.UnpackTrade() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestPackTrade(t *testing.T) {
// 	type args struct {
// 		msg *models.CurrentMessage
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := helpers.PackTrade(tt.args.msg); got != tt.want {
// 				t.Errorf("helpers.PackTrade() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
