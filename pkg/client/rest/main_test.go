package restclient

import (
	"context"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/laozhubaba/api_proj/cmd/server/common"
	"github.com/laozhubaba/api_proj/cmd/server/rest/start"
	"github.com/laozhubaba/api_proj/pkg/server"
)

func TestRestClient(t *testing.T) {
	log.Printf("creating context with cancel...")
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(func() {
		log.Printf("running cancel function...")
		cancel()
	})
	log.Printf("starting the REST server...")
	go common.Start(ctx, start.StartRest)
	log.Printf("sleeping for 10 milliseconds to give the server time to start...")
	time.Sleep(10 * time.Millisecond)

	t.Run("TestAddUser", _testAddUser)
	t.Run("TestGetUser", _testGetUser)
	t.Run("TestGetAllUsers", _testGetAllUsers)
}
func _testAddUser(t *testing.T) {
	type args struct {
		person server.Person
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Add John",
			args: args{
				person: server.Person{
					ID:      1,
					Name:    "John",
					Address: "123 Horrid Lane",
				},
			},
			wantErr: false,
		},
		{
			name: "Add Mary",
			args: args{
				person: server.Person{
					ID:      2,
					Name:    "Mary",
					Address: "234 Lovely Lane",
				},
			},
			wantErr: false,
		},
		{
			name: "Add Ann",
			args: args{
				person: server.Person{
					ID:      3,
					Name:    "Ann",
					Address: "789 Somewhere Else",
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

func _testGetUser(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name       string
		args       args
		wantPerson *server.Person
		wantErr    bool
	}{
		{
			name: "Test getting user 1",
			args: args{
				id: 1,
			},
			wantPerson: &server.Person{
				ID:      1,
				Name:    "John",
				Address: "123 Horrid Lane",
			},
			wantErr: false,
		},
		{
			name: "Test getting user 2",
			args: args{
				id: 2,
			},
			wantPerson: &server.Person{
				ID:      2,
				Name:    "Mary",
				Address: "234 Lovely Lane",
			},
			wantErr: false,
		},
		{
			name: "Test getting user 3",
			args: args{
				id: 3,
			},
			wantPerson: &server.Person{
				ID:      3,
				Name:    "Ann",
				Address: "789 Somewhere Else",
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

func _testGetAllUsers(t *testing.T) {
	tests := []struct {
		name        string
		wantPersons []server.Person
		wantErr     bool
	}{
		{
			name: "test GetAllUsers",
			wantPersons: []server.Person{
				{
					ID:      1,
					Name:    "John",
					Address: "123 Horrid Lane",
				},
				{
					ID:      2,
					Name:    "Mary",
					Address: "234 Lovely Lane",
				},
				{
					ID:      3,
					Name:    "Ann",
					Address: "789 Somewhere Else",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPersons, err := GetAllUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPersons, tt.wantPersons) {
				t.Errorf("GetAllUsers() = %v, want %v", gotPersons, tt.wantPersons)
			}
		})
	}
}
