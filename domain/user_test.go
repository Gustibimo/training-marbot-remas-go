package domain

import (
	"reflect"
	"testing"

	"github.com/Gustibimo/training-marbot-remas-go.git/persistence"
)

func TestUserService_FindUserByFirstName(t *testing.T) {
	type fields struct {
		Users *persistence.Users
	}
	type args struct {
		firstName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []persistence.Users
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				Users: tt.fields.Users,
			}
			got, err := s.FindUserByFirstName(tt.args.firstName)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.FindUserByFirstName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.FindUserByFirstName() = %v, want %v", got, tt.want)
			}
		})
	}
}
