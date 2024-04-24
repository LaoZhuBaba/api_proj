package restclient

import (
	"reflect"
	"testing"

	"github.com/laozhubaba/api_proj/pkg/server"
)

func TestAddUser(t *testing.T) {
	type args struct {
		person server.Person
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test AddUser",
			args: args{
				person: server.Person{
					Name:    "John2",
					Address: "123 Horrid Lane",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddUser(tt.args.person); (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name       string
		args       args
		wantPerson server.Person
		wantErr    bool
	}{
		{
			name: "test GetUser",
			args: args{
				id: 1,
			},
			wantPerson: server.Person{
				Name:    "David",
				Address: "David's address",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPerson, err := GetUser(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotPerson, tt.wantPerson) {
				t.Errorf("GetUser() = %v, want %v", gotPerson, tt.wantPerson)
			}
		})
	}
}

// func TestGetUser2(t *testing.T) {
// 	type args struct {
// 		id int
// 	}
// 	tests := []struct {
// 		name       string
// 		args       args
// 		wantPerson server.Person
// 		wantErr    bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gotPerson, err := GetUser(tt.args.id)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(gotPerson, tt.wantPerson) {
// 				t.Errorf("GetUser() = %v, want %v", gotPerson, tt.wantPerson)
// 			}
// 		})
// 	}
// }
